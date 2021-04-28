package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-ticketing/apps"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Custom Validator
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Trailing slash
	e.Pre(middleware.RemoveTrailingSlash())

	// Static folder images
	e.Static("/", "public")

	// Main root
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.M{"message": "success"})
	})

	// App initialization
	apps.AppInit(e)
	// Run server

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"

// 	"github.com/hpazk/go-ticketing/helper"
// )

// // Queue holds name, list of jobs and context with cancel.
// type Queue struct {
// 	name   string
// 	jobs   chan Job
// 	ctx    context.Context
// 	cancel context.CancelFunc
// }

// // Job - holds logic to perform some operations during queue execution.
// type Job struct {
// 	Name   string
// 	Action func() error // A function that should be executed when the job is running.
// }

// // NewQueue instantiates new
// func NewQueue(name string) *Queue {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	return &Queue{
// 		jobs:   make(chan Job),
// 		name:   name,
// 		ctx:    ctx,
// 		cancel: cancel,
// 	}
// }

// // AddJobs adds jobs to the queue and cancels channel.
// func (q *Queue) AddJobs(jobs []Job) {
// 	var wg sync.WaitGroup
// 	wg.Add(len(jobs))

// 	for _, job := range jobs {
// 		// Goroutine which adds job to the
// 		go func(job Job) {
// 			q.AddJob(job)
// 			wg.Done()
// 		}(job)
// 	}

// 	go func() {
// 		wg.Wait()
// 		// Cancel queue channel, when all goroutines were done.
// 		q.cancel()
// 	}()
// }

// // AddJob sends job to the channel.
// func (q *Queue) AddJob(job Job) {
// 	q.jobs <- job
// 	log.Printf("New job %s added to %s queue", job.Name, q.name)
// }

// // Run performs job execution.
// func (j Job) Run() error {
// 	// log.Printf("Job running: %s", j.GetName())

// 	err := j.Action()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Worker responsible for queue serving.
// type Worker struct {
// 	Queue *Queue
// }

// // NewWorker initializes a new Worker.
// func NewWorker(queue *Queue) *Worker {
// 	return &Worker{
// 		Queue: queue,
// 	}
// }

// // DoWork processes jobs from the queue (jobs channel).
// func (w *Worker) DoWork() bool {
// 	for {
// 		select {
// 		// if context was canceled.
// 		case <-w.Queue.ctx.Done():
// 			log.Printf("Work done in queue %s: %s!", w.Queue.name, w.Queue.ctx.Err())
// 			return true
// 		// if job received.
// 		case job := <-w.Queue.jobs:
// 			err := job.Run()
// 			if err != nil {
// 				log.Print(err)
// 				continue
// 			}
// 		}
// 	}
// }

// // Our products storage.
// var products = []string{
// 	"books",
// 	"computers",
// }

// func main() {
// 	// New products, which we need to add to our products storage.
// 	emails := []string{
// 		"muhammadiqbalali167@gmail.com",
// 		"refinerydev@gmail.com",
// 		"hpazk.zkh@gmail.com",
// 		"djajan.project@gmail.com",
// 		"djajan.checker@gmail.com",
// 	}
// 	// New queue initialization.
// 	productsQueue := NewQueue("NewProducts")
// 	var jobs []Job

// 	// Range over new products.
// 	for _, email := range emails {
// 		// We need to do this, because variables declared in for loops are passed by reference.
// 		// Otherwise, our closure will always receive the last item from the newProducts.
// 		// Defining of the closure, where we add a new product to our simple storage (products slice)
// 		action := func() error {
// 			go helper.SendEmail(email, "Tes", "tes tes tes")
// 			return nil
// 		}
// 		// Append job to jobs slice.
// 		jobs = append(jobs, Job{
// 			Name:   fmt.Sprintf("Importing new product: %s", email),
// 			Action: action,
// 		})
// 	}

// 	// Adds jobs to the
// 	productsQueue.AddJobs(jobs)

// 	// Defines a queue worker, which will execute our
// 	worker := NewWorker(productsQueue)
// 	// Execute jobs in
// 	worker.DoWork()

// 	// Prints products storage after queue execution.
// 	defer fmt.Print(products)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"

// 	"github.com/gocraft/work"
// 	"github.com/gomodule/redigo/redis"
// 	"github.com/hpazk/go-ticketing/helper"
// )

// // Make a redis pool
// var redisPool = &redis.Pool{
// 	MaxActive: 5,
// 	MaxIdle:   5,
// 	Wait:      true,
// 	Dial: func() (redis.Conn, error) {
// 		return redis.Dial("tcp", ":6379")
// 	},
// }

// // Make an enqueuer with a particular namespace
// var enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)

// func main() {
// 	// Enqueue a job named "send_email" with the specified parameters.
// 	_, err := enqueuer.Enqueue("send_email", work.Q{"address": "muhammadiqbalali167@gmail.com", "subject": "hello world", "customer_id": 4})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Make a new pool. Arguments:
// 	// Context{} is a struct that will be the context for the request.
// 	// 10 is the max concurrency
// 	// "my_app_namespace" is the Redis namespace
// 	// redisPool is a Redis pool
// 	pool := work.NewWorkerPool(Context{}, 10, "my_app_namespace", redisPool)

// 	// Add middleware that will be executed for each job
// 	pool.Middleware((*Context).Log)
// 	pool.Middleware((*Context).FindCustomer)

// 	// Map the name of jobs to handler functions
// 	pool.Job("send_email", (*Context).SendEmail)

// 	// Customize options:
// 	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

// 	// Start processing jobs
// 	pool.Start()

// 	// Wait for a signal to quit:
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, os.Interrupt, os.Kill)
// 	<-signalChan

// 	// Stop the pool
// 	pool.Stop()
// }

// func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
// 	fmt.Println("Starting job: ", job.Name)
// 	return next()
// }

// func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
// 	// If there's a customer_id param, set it in the context for future middleware and handlers to use.
// 	if _, ok := job.Args["customer_id"]; ok {
// 		c.customerID = job.ArgInt64("customer_id")
// 		if err := job.ArgError(); err != nil {
// 			return err
// 		}
// 	}

// 	return next()
// }

// type Context struct {
// 	customerID int64
// }

// func (c *Context) SendEmail(job *work.Job) error {
// 	// Extract arguments:
// 	addr := job.ArgString("address")
// 	subject := job.ArgString("subject")
// 	if err := job.ArgError(); err != nil {
// 		return err
// 	}

// 	// Go ahead and send the email...
// 	helper.SendEmail(addr, subject, "test")

// 	return nil
// }

// func (c *Context) Export(job *work.Job) error {
// 	return nil
// }

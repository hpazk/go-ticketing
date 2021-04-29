package scheduler

import (
	"fmt"

	"github.com/hpazk/go-ticketing/apps/user"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/jasonlvhit/gocron"
)

func PaymentNotification() {

}

func SendPromotion() {
	// queue := list.New()

	userService := user.UserService()
	usersList, _ := userService.FetchUsers()

	messages := make(chan string, 2)

	go func() {
		for {
			i := <-messages

			helper.SendEmail(i, "Test", "test test tes")
			fmt.Println("receive data", i)
		}
	}()

	for _, u := range usersList {
		messages <- u.Email
	}

}

func Scheduler() {
	s := gocron.NewScheduler()
	//s.Every(1).Second().Do(taskWithParams, 1, "hello")
	// err := s.Every(1).Seconds().Do(SendPromotion)
	err := s.Every(1).Day().At("06:00").Do(SendPromotion)
	if err != nil {
		fmt.Println(err)
	}
	<-s.Start()
}

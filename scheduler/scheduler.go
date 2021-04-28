package scheduler

import (
	"fmt"

	"github.com/hpazk/go-ticketing/apps/user"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/jasonlvhit/gocron"
)

func task() {
	userService := user.UserService()
	users, _ := userService.FetchUsers()
	for _, u := range users {
		helper.SendEmail(u.Email, "Subject", "tes")
	}
}

// func taskWithParams(a int, b string) {
// 	fmt.Println(a, b)
// }

func Scheduler() {
	s := gocron.NewScheduler()
	//s.Every(1).Second().Do(taskWithParams, 1, "hello")
	// err := s.Every(1).Seconds().Do(task)
	err := s.Every(1).Day().At("06:00").Do(task)
	if err != nil {
		fmt.Println(err)
	}
	<-s.Start()
}

package app

import (
	"fmt"

	"github.com/u8slvn/raspigosms/gsm"
	"github.com/u8slvn/raspigosms/web/controllers"
)

//StartDispatcher function
func StartDispatcher() {
	fmt.Println("Starting the dispatcher")
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue := make(chan gsm.Sms)

	// Create the sms worker.
	fmt.Println("Starting worker")
	worker := NewWorker(WorkerQueue)
	worker.Start()

	go func() {
		for {
			select {
			case sms := <-controllers.SmsQueue:
				fmt.Println("Received sms request")
				worker.WorkerQueue <- sms
			}
		}
	}()
}

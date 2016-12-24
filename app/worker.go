package app

import (
	"fmt"
	"time"

	"github.com/u8slvn/raspigosms/gsm"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(workerQueue chan gsm.Sms) Worker {
	worker := Worker{
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}

	return worker
}

// Worker struct
type Worker struct {
	WorkerQueue chan gsm.Sms
	QuitChan    chan bool
}

// Start function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case sms := <-w.WorkerQueue:
				// Receive a work request.
				fmt.Printf("worker: Received sms, for %s\n", sms.Phone)
				// Here will be the SMS sender.
				time.Sleep(4)
				fmt.Printf("worker: => to : %s, message : %s\n", sms.Phone, sms.Message)

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker stopping\n")
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

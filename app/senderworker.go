package app

import (
	"fmt"
	"time"

	"github.com/u8slvn/raspigosms/database"
	"github.com/u8slvn/raspigosms/gsm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewSenderWorker creates, and returns a new SenderWorker object.
func NewSenderWorker(workerQueue chan gsm.Sms) SenderWorker {
	worker := SenderWorker{
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}

	return worker
}

// SenderWorker struct
type SenderWorker struct {
	WorkerQueue chan gsm.Sms
	QuitChan    chan bool
}

// Start function "starts" an infinite loop which consume the SmsQueue.
func (w SenderWorker) Start() {
	fmt.Printf("SenderWorker starting...\n")
	go func() {
		for {
			select {
			case sms := <-w.WorkerQueue:
				fmt.Printf("worker: Received sms, for %s\n", sms.Phone)
				// Here will be the SMS sender.
				time.Sleep(4)
				fmt.Printf("worker: => to : %s, message : %s\n", sms.Phone, sms.Message)
				go func() {
					change := mgo.Change{
						Update:    bson.M{"$set": bson.M{"status": gsm.SmsSent}},
						ReturnNew: true,
					}
					database.DBConnection.DB("raspi_go_sms").C("sms").FindId(sms.UUID).Apply(change, &sms)
				}()

			case <-w.QuitChan:
				fmt.Printf("worker stopping\n")
				return
			}
		}
	}()
}

// Stop the worker to stop listening for Sms requests.
// The worker will only stop *after* it has finished its work.
func (w SenderWorker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

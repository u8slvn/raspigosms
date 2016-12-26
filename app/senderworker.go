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
func NewSenderWorker() SenderWorker {
	worker := SenderWorker{
		WorkerQueue: make(chan SmsRequest),
		QuitChan:    make(chan bool),
	}

	return worker
}

// SenderWorker struct
type SenderWorker struct {
	WorkerQueue chan SmsRequest
	QuitChan    chan bool
}

// Start function "starts" an infinite loop which consume the SmsQueue.
func (w SenderWorker) Start() {
	fmt.Printf("SenderWorker starting...\n")
	go func() {
		for {
			select {
			case smsr := <-w.WorkerQueue:
				fmt.Printf("worker: Received sms, for %s\n", smsr.Sms.Phone)
				// Here will be the SMS sender.
				time.Sleep(4)
				fmt.Printf("worker: => to : %s, message : %s\n", smsr.Sms.Phone, smsr.Sms.Message)
				go func() {
					change := mgo.Change{
						Update:    bson.M{"$set": bson.M{"status": gsm.SmsSent}},
						ReturnNew: true,
					}
					database.DBConnection.C("sms").FindId(smsr.Sms.UUID).Apply(change, &smsr.Sms)
					smsr.RemainingAttempts = 0
					if smsr.RemainingAttempts != 0 {
						SmsRequestQueue <- smsr
					}
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

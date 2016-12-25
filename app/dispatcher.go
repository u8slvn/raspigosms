package app

import "github.com/u8slvn/raspigosms/gsm"

// SmsQueue is a buffered channel used to send sms requests on.
var SmsQueue = make(chan gsm.Sms, 100)

//StartWorking function
func StartWorking() {

	senderWorker := NewSenderWorker()
	senderWorker.Start()

	go func() {
		for {
			select {
			case sms := <-SmsQueue:
				go func() {
					senderWorker.WorkerQueue <- sms
				}()
			}
		}
	}()
}

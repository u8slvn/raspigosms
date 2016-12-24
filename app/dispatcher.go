package app

import "github.com/u8slvn/raspigosms/gsm"

// SmsQueue is a buffered channel used to send sms requests on.
var SmsQueue = make(chan gsm.Sms, 100)

//StartWorking function
func StartWorking() {
	workerList := [2]Worker{
		NewSenderWorker(SmsQueue),
		NewReceiverWorker(),
	}

	for _, worker := range workerList {
		worker.Start()
	}
}

package app

// SmsRequestQueue is a buffered channel used to send sms requests on.
var SmsRequestQueue = make(chan SmsRequest, 100)

//StartWorking function
func StartWorking() {

	senderWorker := NewSenderWorker()
	senderWorker.Start()

	go func() {
		for {
			select {
			case smsr := <-SmsRequestQueue:
				go func() {
					senderWorker.WorkerQueue <- smsr
				}()
			}
		}
	}()
}

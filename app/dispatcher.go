package app

import "github.com/u8slvn/raspigosms/gsm"

// SmsRequestQueue is a buffered channel used to send sms requests on.
var SmsRequestQueue = make(chan SmsRequest, 100)

//StartWorking function
func StartWorking() {
	LoadConfig()

	modem := gsm.NewModem(Conf.Modem.Serial, Conf.Modem.Baud)
	modem.Connect()

	senderWorker := NewSenderWorker(&modem)
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

package raspigosms

import "github.com/u8slvn/raspigosms/gsm"

// SmsRequestQueue is a buffered channel used to send sms requests on.
var SmsRequestQueue = make(chan smsRequest, 100)

//Start raspigosms
func Start() {
	dbConnect()
	loadConfig()

	modem := gsm.NewModem(Conf.Modem.Serial, Conf.Modem.Baud)
	modem.Connect()

	sw := newSenderWorker(modem)
	sw.Start()

	go func() {
		for {
			select {
			case smsr := <-SmsRequestQueue:
				go func() {
					sw.WorkerQueue <- smsr
				}()
			}
		}
	}()
}

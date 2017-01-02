package raspigosms

import "github.com/u8slvn/raspigosms/gsm"
import "fmt"

// SmsRequestQueue is a buffered channel used to send sms requests on.
var SmsRequestQueue = make(chan SmsRequest, 100)

//Start function
func Start() {
	databaseConnect()
	loadConfig()

	modem := gsm.NewModem(Conf.Modem.Serial, Conf.Modem.Baud)
	fmt.Println(modem)
	modem.Connect()

	senderWorker := NewSenderWorker(modem)
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

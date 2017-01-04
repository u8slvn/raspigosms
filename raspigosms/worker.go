package raspigosms

import (
	"fmt"

	"github.com/u8slvn/raspigosms/raspigosms/gsm"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type senderWorker struct {
	WorkerQueue chan smsRequest
	Modem       *gsm.Modem
	QuitChan    chan bool
}

func newSenderWorker(modem *gsm.Modem) senderWorker {
	worker := senderWorker{
		WorkerQueue: make(chan smsRequest),
		Modem:       modem,
		QuitChan:    make(chan bool),
	}

	return worker
}

// Start function "starts" an infinite loop which consume the SmsQueue.
func (w senderWorker) Start() {
	fmt.Printf("SenderWorker starting...\n")
	go func() {
		for {
			select {
			case smsr := <-w.WorkerQueue:
				fmt.Printf("worker: => to : %s, message : %s\n", smsr.sms.Phone, smsr.sms.Message)
				go func() {
					status := gsm.SmsStatusFailed
					err := w.Modem.SendSms(smsr.sms)
					if err == nil {
						smsr.remainingAttempts = 0
						status = gsm.SmsStatusSent
						fmt.Printf("Success\n")
					} else {
						smsr.remainingAttempts--
						fmt.Printf("Failed\n")
					}

					if smsr.remainingAttempts > 0 {
						smsRequestQueue <- smsr
						return
					}

					change := mgo.Change{
						Update:    bson.M{"$set": bson.M{"status": status}},
						ReturnNew: true,
					}
					db.C("sms").FindId(smsr.sms.UUID).Apply(change, &smsr.sms)
				}()
			case <-w.QuitChan:
				fmt.Printf("worker stopping\n")
				return
			}
		}
	}()
}

// The worker will only stop *after* it has finished its work.
func (w senderWorker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

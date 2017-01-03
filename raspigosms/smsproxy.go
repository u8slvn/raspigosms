package raspigosms

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/u8slvn/raspigosms/gsm"
)

type smsRequest struct {
	remainingAttempts int
	sms               gsm.Sms
}

// SubmitSms Build and add a SmsRequest to the SmsRequestQueue.
func SubmitSms(phone string, message string) (gsm.Sms, error) {
	sms, err := gsm.NewSms(phone, message, gsm.SmsStatusPending)
	if err != nil {
		return sms, err
	}

	if err := db.C("sms").Insert(sms); err != nil {
		return sms, err
	}

	smsr := smsRequest{
		remainingAttempts: Conf.RemainingAttempts,
		sms:               sms,
	}

	SmsRequestQueue <- smsr
	return sms, nil
}

func FindSmsByUUID(id string) (gsm.Sms, error) {
	var sms gsm.Sms

	UUID, err := uuid.ParseHex(id)
	if err != nil {
		return sms, err
	}

	if err := db.C("sms").FindId(UUID).One(&sms); err != nil {
		return sms, err
	}

	return sms, nil
}

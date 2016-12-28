package app

import "github.com/u8slvn/raspigosms/gsm"

// SmsRequest model
type SmsRequest struct {
	RemainingAttempts int
	Sms               gsm.Sms
}

// NewSmsRequest Build and return a new SmsRequest.
func NewSmsRequest(sms gsm.Sms) SmsRequest {
	return SmsRequest{
		RemainingAttempts: Conf.RemainingAttempts, // todo : move this var in config
		Sms:               sms,
	}
}

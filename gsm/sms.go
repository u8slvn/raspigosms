package gsm

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/nu7hatch/gouuid"
)

const (
	SmsPending = iota
	SmsSent
	SmsFailed
)

// NewSms creates and return an Sms object with automatic
func NewSms(phone string, message string) (Sms, error) {
	var sms Sms

	if phone == "" {
		return sms, errors.New("The phone number is required")
	}

	if message == "" {
		return sms, errors.New("The sms message body is required")
	}

	u4, err := uuid.NewV4()
	if err != nil {
		return sms, errors.New("The sms message body is required")
	}

	sms = Sms{
		UUID:         u4,
		Phone:        phone,
		Message:      message,
		Status:       SmsPending,
		ReceivedDate: time.Now(),
	}

	return sms, nil
}

// Sms struct
type Sms struct {
	UUID         *uuid.UUID `json:"uuid"`
	Phone        string     `json:"phone"`
	Message      string     `json:"message"`
	Status       int        `json:"status"`
	ReceivedDate time.Time  `json:"received_date"`
}

// MarshalJSON overwrite the standard json.Marshal function to
// add a custom format to the Sms struct
func (sms *Sms) MarshalJSON() ([]byte, error) {
	type Alias Sms
	return json.Marshal(&struct {
		UUID string `json:"uuid"`
		*Alias
	}{
		UUID:  sms.UUID.String(),
		Alias: (*Alias)(sms),
	})
}

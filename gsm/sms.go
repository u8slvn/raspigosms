package gsm

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/nu7hatch/gouuid"
)

// Sms available status
const (
	SmsPending = iota
	SmsSent
	SmsFailed
)

// Sms struct model
type Sms struct {
	UUID         *uuid.UUID `json:"uuid" bson:"_id"`
	Phone        string     `json:"phone" bson:"phone"`
	Message      string     `json:"message" bson:"message"`
	Status       int        `json:"status" bson:"status"`
	ReceivedDate time.Time  `json:"received_date" bson:"received_date"`
}

// NewSms creates and return an Sms object with automatic
// UUID and ReceivedDate generation
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

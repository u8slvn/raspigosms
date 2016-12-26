package gsm

import (
	"fmt"
	"math/rand"
	"time"
)

// Modem model
type Modem struct {
	Serial string
	Baud   int
}

// NewModem build and return a new (fake) Modem
func NewModem(serial string, baud int) Modem {
	return Modem{
		Serial: serial,
		Baud:   baud,
	}
}

func (m *Modem) Connect() error {
	return nil
}

func (m *Modem) SendSms(sms Sms) error {
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(10)
	if r > 7 {
		return fmt.Errorf("Sending sms failed")
	}
	return nil
}

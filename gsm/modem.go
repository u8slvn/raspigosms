package gsm

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

// Modem model
type Modem struct {
	Serial string
	Baud   int
	Port   *serial.Port
}

// NewModem build and return a new (fake) Modem
func NewModem(serial string, baud int) *Modem {
	return &Modem{
		Serial: serial,
		Baud:   baud,
	}
}

func (m *Modem) Connect() (err error) {
	config := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	m.Port = s
	return err
}

func (m *Modem) SendSms(sms Sms) (err error) {
	_, err = m.Port.Write([]byte("AT+CMGF=1\r"))
	if err != nil {
		fmt.Println(err.Error())
	}
	_, _ = m.Port.Write([]byte("AT+CMGS=\"" + sms.Phone + "\"\r"))
	_, _ = m.Port.Write([]byte(sms.Message + string(26)))

	buf := make([]byte, 128)
	n, err := m.Port.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("%q", buf[:n])

	return err
}

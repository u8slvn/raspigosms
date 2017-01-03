package gsm

import (
	"fmt"
	"log"
	"sync"

	"github.com/tarm/serial"
)

// Modem model
type Modem struct {
	Serial string
	Baud   int
	Port   *serial.Port
	mu     *sync.Mutex
}

// NewModem build and return a new Modem.
func NewModem(serial string, baud int) *Modem {
	return &Modem{
		Serial: serial,
		Baud:   baud,
		mu:     &sync.Mutex{},
	}
}

// Connect the usb gsm modem.
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
	m.mu.Lock()
	_, err = m.Port.Write([]byte("AT+CMGF=1\r"))
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = m.Port.Write([]byte("AT+CMGS=\"" + sms.Phone + "\"\r"))
	_, err = m.Port.Write([]byte(sms.Message + string(26)))

	buf := make([]byte, 128)
	n, err := m.Port.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("%q", buf[:n])
	m.mu.Unlock()
	return err
}

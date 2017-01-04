package gsm

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/tarm/serial"
)

const (
	atcmgf = "AT+CMGF=1\r"                       // AT command : sms format (1 = text mode)
	atcmgs = "AT+CMGS=\"%s\"\r%s" + string(0x1A) // AT command : send message
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
	config := &serial.Config{Name: "/dev/ttyUSB1", Baud: 115200, ReadTimeout: time.Second}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	m.Port = s
	return err
}

// SendSms Sends an sms in text mode.
func (m *Modem) SendSms(sms Sms) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if err := m.Port.Flush(); err != nil {
		return err
	}

	// switch the modem to text mode
	if err := m.writeOnPort(atcmgf); err != nil {
		return err
	}
	ret, err := m.readOnPort(1)
	if err != nil {
		return err
	}
	if strings.Contains(ret, "ERROR") {
		return errors.New("modem can't switch to text mode")
	}
	log.Println("retour 1")
	log.Printf("%q", ret)

	// send sms
	if err = m.writeOnPort(fmt.Sprintf(atcmgs, sms.Phone, sms.Message)); err != nil {
		return err
	}
	ret2, err := m.readOnPort(7)
	log.Println("retour 2")
	log.Printf("%q", ret2)

	return nil
}

func (m *Modem) writeOnPort(str string) error {
	if _, err := m.Port.Write([]byte(str)); err != nil {
		return err
	}
	return nil
}

func (m *Modem) readOnPort(timeOut int) (string, error) {
	var buffer bytes.Buffer
	buf := make([]byte, 64)
	for {
		n, _ := m.Port.Read(buf)
		if n > 0 {
			ret := string(buf[:n])
			fmt.Printf("%q", ret)
			buffer.WriteString(ret)
		}

		if n == 0 && timeOut <= 0 {
			return buffer.String(), nil
		}
		buf = make([]byte, 64)
		timeOut--
	}
}

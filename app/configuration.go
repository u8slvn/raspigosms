package app

import (
	"encoding/json"
	"os"

	"github.com/u8slvn/raspigosms/gsm"
)

var Conf configuration

type configuration struct {
	HTTPAddr          string
	RemainingAttempts int
	Modem             gsm.Modem
}

func LoadConfig() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	Conf = configuration{}
	err := decoder.Decode(&Conf)
	if err != nil {
		panic(err)
	}
}

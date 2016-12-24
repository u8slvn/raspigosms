package app

import "fmt"

func NewReceiverWorker() ReceiverWorker {
	worker := ReceiverWorker{
		QuitChan: make(chan bool),
	}

	return worker
}

type ReceiverWorker struct {
	QuitChan chan bool
}

func (w ReceiverWorker) Start() {
	fmt.Printf("ReceiverWorker starting...\n")
}

func (w ReceiverWorker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

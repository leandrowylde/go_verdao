package execute

import (
	"fmt"
	"time"

	"github.com/cogny/go_verdao/application/model"
	"github.com/cogny/go_verdao/application/usecase"
)

type Execute struct {
	TestAPIUseCase usecase.TestAPIUseCase
}

func (e *Execute) Start(config model.Config) {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("Check schedule...: %v", t)
				fmt.Printf("Start: %v", config.Start)
				fmt.Printf("Stop: %v", config.Stop)
				fmt.Printf("URIs: %v", config.URLs)
			}
		}
	}()
}

package execute

import (
	"fmt"
	"time"

	"github.com/cogny/go_verdao/application/faktory"
	"github.com/cogny/go_verdao/application/model"
	"github.com/cogny/go_verdao/application/usecase"
	"github.com/jinzhu/gorm"
)

type Execute struct {
	TestAPIUseCase usecase.TestAPIUseCase
}

func (e *Execute) Start(config model.Config) {
	tickerStart := time.NewTicker(5 * time.Second)
	tickerStop := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	fmt.Printf("Vai verdão (Agora o mundial vem)...\n")

	execRequest := func(uri model.URI) {
		code, data, err := e.TestAPIUseCase.MakeRequest(uri)
		if err != nil {
			panic(err)
		}
		var logData string
		if len(data) > 101 {
			logData = data[0:100] + "..."
		} else {
			logData = data
		}
		fmt.Printf("Call API: %v - %v", code, logData)
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-tickerStart.C:
				timeStart, err := time.Parse(time.RFC3339, config.Start)
				if err != nil {
					panic(err)
				}
				timeStop, err := time.Parse(time.RFC3339, config.Stop)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Config start time: %v | Time Now: %v\n", timeStart, time.Now().Format(time.RFC3339))
				if time.Now().After(timeStart) && time.Now().Before(timeStop) {
					fmt.Printf("GAME TIME\n")
					for _, u := range e.TestAPIUseCase.APIs {
						go execRequest(u)
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case <-tickerStop.C:
				timeStop, err := time.Parse(time.RFC3339, config.Stop)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Config stop time: %v | Time Now: %v\n", timeStop, time.Now().Format(time.RFC3339))
				if time.Now().After(timeStop) {
					fmt.Printf("END GAME...sending stop ticker\n")
					done <- true
				}
			}
		}
	}()
	<-done
}

func NewExecuteProcess(database *gorm.DB) (Execute, error) {
	config, err := model.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	testAPIsUseCase := faktory.TestAPIUseCaseFaktory(config.URLs, database)

	return Execute{
		TestAPIUseCase: testAPIsUseCase,
	}, nil
}

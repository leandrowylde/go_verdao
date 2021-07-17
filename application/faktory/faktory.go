package faktory

import "github.com/cogny/go_verdao/application/model"

func TestAPIUseCaseFaktory(uris []model.URI) model.Config {
	config, err := model.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	return config
}

package faktory

import (
	"github.com/cogny/go_verdao/application/model"
	"github.com/cogny/go_verdao/application/usecase"
	"github.com/cogny/go_verdao/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TestAPIUseCaseFaktory(uris []model.URI, database *gorm.DB) usecase.TestAPIUseCase {
	repo := repository.ResultRespositoryDB{DB: database}
	testAPIUseCase := usecase.TestAPIUseCase{
		APIs:             uris,
		ResultRepository: repo,
	}
	return testAPIUseCase
}

package respository

import (
	"github.com/cogny/go_verdao/application/model"
	"github.com/jinzhu/gorm"
)

type ResultRespositoryDB struct {
	DB *gorm.DB
}

func (r *ResultRespositoryDB) SaveResult(result model.Result) error {
	err := r.DB.Create(result).Error
	return err
}

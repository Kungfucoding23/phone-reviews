package gateway

import (
	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/models"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/logs"
)

//SmartphoneStorageGateway ..
type SmartphoneStorageGateway interface {
	create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

//SmartphoneStorage ...
type SmartphoneStorage struct {
	*database.MySQLClient
}

func (s *SmartphoneStorage) create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := s.MySQLClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, os) 
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.OS)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		ID:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		OS:            cmd.OS,
	}, nil
}

package gateway

import (
	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/models"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
)

//SmartphoneCreateGateway ...
type SmartphoneCreateGateway interface {
	Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

//SmartphoneCreateGtw ...
type SmartphoneCreateGtw struct {
	SmartphoneStorageGateway
}

//NewSmartphoneCreateGateway ...
func NewSmartphoneCreateGateway(client *database.MySQLClient) SmartphoneCreateGateway {
	return &SmartphoneCreateGtw{&SmartphoneStorage{client}}
}

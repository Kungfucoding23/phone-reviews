package gateway

import "github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/models"

//Create ..
func (s *SmartphoneCreateGtw) Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	return s.create(cmd)
}

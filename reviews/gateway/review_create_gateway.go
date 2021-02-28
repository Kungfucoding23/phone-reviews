package gateway

import (
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
	"github.com/Kungfucoding23/API-Go-mysql-chi/reviews/models"
)

//ReviewGateway ..
type ReviewGateway interface {
	AddReview(cmd *models.CreateReviewCMD) (string, error)
}

//ReviewGtw ..
type ReviewGtw struct {
	ReviewStorage
}

//NewReviewGateway ..
func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}

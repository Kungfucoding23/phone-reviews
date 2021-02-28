package gateway

import (
	"context"
	"time"

	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/logs"
	"github.com/Kungfucoding23/API-Go-mysql-chi/reviews/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReviewStorage ...
type ReviewStorage interface {
	Add(cmd *models.CreateReviewCMD) (string, error)
}

//ReviewStg ...
type ReviewStg struct {
	*database.Mongo
}

//Add --
func (s *ReviewStg) Add(cmd *models.CreateReviewCMD) (string, error) {
	coll := s.Client.Database("reviewDB").Collection("reviews")

	res, err := coll.InsertOne(context.Background(),
		bson.D{
			{"comment", cmd.Comment},
			{"stars", cmd.Stars},
			{"createdAt", time.Now()},
			{"gadgetID", cmd.GadgetID},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}

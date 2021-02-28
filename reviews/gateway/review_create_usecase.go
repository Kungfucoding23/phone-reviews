package gateway

import "github.com/Kungfucoding23/API-Go-mysql-chi/reviews/models"

//AddReview ..
func (g *ReviewGtw) AddReview(cmd *models.CreateReviewCMD) (string, error) {
	return g.Add(cmd)
}

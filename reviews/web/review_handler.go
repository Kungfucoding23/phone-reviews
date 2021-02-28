package web

import (
	"encoding/json"
	"net/http"

	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
	"github.com/Kungfucoding23/API-Go-mysql-chi/reviews/gateway"
	"github.com/Kungfucoding23/API-Go-mysql-chi/reviews/models"
)

//AddReviewHandler ..
func (h *ReviewHandler) AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	cmd := params(r)
	res, err := h.AddReview(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(res))
}

func params(r *http.Request) *models.CreateReviewCMD {
	var cmd models.CreateReviewCMD

	_ = json.NewDecoder(r.Body).Decode(&cmd)

	return &cmd
}

//ReviewHandler ..
type ReviewHandler struct {
	gateway.ReviewGateway
}

//NewReviewHandler ..
func NewReviewHandler(mongo *database.Mongo) *ReviewHandler {
	return &ReviewHandler{gateway.NewReviewGateway(mongo)}
}

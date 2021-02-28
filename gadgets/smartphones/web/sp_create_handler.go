package web

import (
	"encoding/json"
	"net/http"

	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/gateway"
	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/models"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
)

//SaveSmartphoneHandler ..
func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create smartphone"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

//CreateSmartphoneHandler ..
type CreateSmartphoneHandler struct {
	gateway.SmartphoneCreateGateway
}

//NewCreateSmartphoneHandler ..
func NewCreateSmartphoneHandler(client *database.MySQLClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{gateway.NewSmartphoneCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateSmartphoneCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateSmartphoneCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

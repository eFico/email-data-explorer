package web

import (
	"encoding/json"
	"net/http"

	"github.com/eFico/email-api/emails/gateway"
	"github.com/eFico/email-api/emails/models"
)

func (h *EmailSearchHandler) EmailSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	emailReq := parseRequest(r)
	res, err := h.Search(emailReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in EmailSearchHandler"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type EmailSearchHandler struct {
	gateway.EmailApiGateway
}

func NewEmailSearchHandler(URL string) *EmailSearchHandler {
	return &EmailSearchHandler{gateway.NewEmailSearchGateway(URL)}
}

func parseRequest(r *http.Request) *models.EmailRequest {
	body := r.Body
	defer body.Close()
	var emailReq models.EmailRequest

	_ = json.NewDecoder(body).Decode(&emailReq)

	return &emailReq
}

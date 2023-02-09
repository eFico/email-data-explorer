package gateway

import (
	"github.com/eFico/email-api/emails/models"
)

type EmailApiGateway interface {
	Search(cmd *models.EmailRequest) (*models.EmailResponse, error)
}

type EmailSearchGtw struct {
	EmailSearchGateway
}

func NewEmailSearchGateway(client string) EmailApiGateway {
	return &EmailSearchGtw{&EmailSearch{client}}
}

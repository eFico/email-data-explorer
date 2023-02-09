package gateway

import "github.com/eFico/email-api/emails/models"

func (s *EmailSearchGtw) Search(cmd *models.EmailRequest) (*models.EmailResponse, error) {
	return s.search(cmd)
}

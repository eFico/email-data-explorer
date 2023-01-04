package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/eFico/email-api/emails/models"
)

type EmailSearchGateway interface {
	search(cmd *models.EmailRequest) (*models.EmailResponse, error)
}

type EmailSearch struct {
	URL string
}

func (s *EmailSearch) search(cmd *models.EmailRequest) (*models.EmailResponse, error) {

	data := `{
		"query": {
			"bool": {
				"must": [
					{
						"query_string": {
							"query": "%s"
						}
					}
				]
			}
		},
		"sort": [
			"-@timestamp"
		],
		"from": %d,
		"size": %d
	}`

	//fmt.Println(string(nombre))
	data = fmt.Sprintf(data, cmd.Query, cmd.Page, cmd.Size)

	req, err := http.NewRequest("POST", s.URL, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respJSON models.ZincResponse
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		// Error al convertir la cadena a JSON
		log.Fatal(err)
	}

	respEmail := models.EmailResponse{
		Records: respJSON.Hits.Hits,
		Total:   respJSON.Hits.Total.Value,
	}
	return &respEmail, nil

}

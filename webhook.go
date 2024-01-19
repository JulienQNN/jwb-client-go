package jwb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
  "time"
)

// CreateWebhook - Create new webhook
func CreateTeamsWebhook(webhookUrl string, payloadWebhook TeamsPayloadWebhook) (*APIResponse, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

  rb, err := json.Marshal(payloadWebhook)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(webhookUrl),
		strings.NewReader(string(rb)),
	)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	webhook := APIResponse{}
	err = json.Unmarshal(body, &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

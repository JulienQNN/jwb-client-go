package jwb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateWebhook - Create new webhook
func (c *Client) CreateTeamsWebhook(webhookUrl string, payloadWebhook TeamsPayloadWebhook) (*TeamsPayloadWebhook, error) {
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

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	webhook := TeamsPayloadWebhook{}

	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

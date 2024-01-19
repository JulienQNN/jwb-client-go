package jwb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateWebhook - Create new webhook
func (c *Client) CreateTeamsWebhook(webhookUrl string, payloadWebhook TeamsPayloadWebhook) (*APIResponse, error) {
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

package jwb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateWebhook - Create new webhook
func (c *Client) CreateSlackWebhook(webhookUrl string, payloadWebhook SlackPayloadWebhook) (*SlackPayloadWebhook, error) {
	rb, err := json.Marshal(payloadWebhook)
	if err != nil {
		return nil, err
	}

	fmt.Println("rb", string(rb))

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

	webhook := SlackPayloadWebhook{}

	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

package jwb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateSlackAttachmentWebhook(webhookUrl string, payloadWebhook SlackAttachmentPayloadWebhook) (*SlackAttachmentPayloadWebhook, error) {
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

	webhook := SlackAttachmentPayloadWebhook{}

	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (c *Client) CreateSlackMessageWebhook(webhookUrl string, payloadWebhook SlackMessagePayloadWebhook) (*SlackMessagePayloadWebhook, error) {
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

	webhook := SlackMessagePayloadWebhook{}

	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

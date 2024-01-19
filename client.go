package jwb

import (
	"fmt"
	"io"
	"net/http"
)

// Client -
type Client struct {
	HTTPClient *http.Client
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	
  res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

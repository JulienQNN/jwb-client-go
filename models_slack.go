package jwb

type SlackPayloadWebhook struct {
	Block []SlackBlock `json:"blocks"`
}

type SlackBlock struct {
	Type      string          `json:"type"`
	Text      *SlackText      `json:"text,omitempty"`
	Title     *SlackTitle     `json:"title,omitempty"`
	ImageUrl  string          `json:"image_url,omitempty"`
	AltText   string          `json:"alt_text,omitempty"`
	Fields    []SlackField    `json:"fields,omitempty"`
	Accessory *SlackAccessory `json:"accessory,omitempty"`
	Elements  []SlackElement  `json:"elements,omitempty"`
}

type SlackField struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SlackText struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type SlackTitle struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

type SlackAccessory struct {
	Type string              `json:"type"`
	Text *SlackAccessoryText `json:"text,omitempty"`
}

type SlackAccessoryText struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

type SlackElement struct {
	Type  string           `json:"type"`
	Text  SlackElementText `json:"text"`
	Style string           `json:"style"`
	Url   string           `json:"value"`
}

type SlackElementText struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

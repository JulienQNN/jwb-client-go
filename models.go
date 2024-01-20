package jwb

// TeamsPayloadWebhook Types for Teams webhook
type TeamsPayloadWebhook struct {
	ID              string                        `json:"id"`
  Type            string                        `json:"type"`
	Context         string                        `json:"context"`
	ThemeColor      string                        `json:"theme_color"`
	Summary         string                        `json:"summary"`
	Sections        []Section                     `json:"sections"`
	PotentialAction []PotentialActionIntermediate `json:"potentialAction"`
}

type Section struct {
	ActivityTitle    *string `json:"title"`
	ActivitySubtitle *string `json:"subtitle"`
	ActivityImage    *string `json:"image"`
	Text             *string `json:"text"`
	Facts            []Fact  `json:"facts"`
	Markdown         *bool   `json:"markdown"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PotentialActionIntermediate struct {
	Type            string            `json:"@type"`
	Name            string            `json:"name"`
  PotentialAction []PotentialAction `json:"potentialAction"`
	Targets         []Target          `json:"targets"`
}

type PotentialAction struct {
  Name    string   `json:"name"`
  Targets []Target `json:"targets"`
}

type Target struct {
	Os  string `json:"os"`
	Uri string `json:"uri"`
}

package jwb

// TeamsPayloadWebhook Types for Teams webhook
type TeamsPayloadWebhook struct {
	Type                 string                 `json:"type"`
	Context              string                 `json:"context"`
	ThemeColor           string                 `json:"themeColor"`
	Summary              string                 `json:"summary"`
	Sections             []TeamsSection         `json:"sections"`
	TeamsPotentialAction []TeamsPotentialAction `json:"potentialAction"`
}

type TeamsSection struct {
	ActivityTitle    string      `json:"activityTitle"`
	ActivitySubtitle string      `json:"activitySubtitle"`
	ActivityImage    string      `json:"activityImage"`
	Text             string      `json:"text"`
	Facts            []TeamsFact `json:"facts,omitempty"`
	Markdown         bool        `json:"markdown,omitempty"`
}

type TeamsFact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TeamsPotentialAction struct {
	Type    string        `json:"@type"`
	Name    string        `json:"name"`
	Targets []TeamsTarget `json:"targets"`
}

type TeamsTarget struct {
	Os  string `json:"os"`
	Uri string `json:"uri"`
}

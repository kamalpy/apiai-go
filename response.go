package apiaigo

type Metadata struct {
	IntentID                  string `json:"intentId"`
	IntentName                string `json:"intentName"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
	WebhookUsed               string `json:"webhookUsed"`
}

type Status struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorID      string `json:"errorId"`
	ErrorDetails string `json:"errorDetails"`
}

type Message struct {
	Speech string `json:"speech"`
	Type   int    `json:"type"`
}

type Fulfillment struct {
	Messages []Message `json:"messages"`
	Speech   string    `json:"speech"`
}

type Result struct {
	Action           string            `json:"action"`
	ActionIncomplete bool              `json:"actionIncomplete"`
	Contexts         []Context         `json:"contexts"`
	Fulfillment      Fulfillment       `json:"fulfillment"`
	Metadata         Metadata          `json:"metadata"`
	Parameters       map[string]string `json:"parameters"`
	ResolvedQuery    string            `json:"resolvedQuery"`
	Score            float32           `json:"score"`
	Source           string            `json:"source"`
}

type ResponseStruct struct {
	ID        string `json:"id"`
	Language  string `json:"lang"`
	Result    Result `json:"result"`
	SessionID string `json:"sessionId"`
	Status    Status `json:"status"`
	Timestamp string `json:"timestamp"`
}

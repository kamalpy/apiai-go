package apiaigo

// Metadata contains data on intents and contexts.
type Metadata struct {
	IntentID                  string `json:"intentId"`
	IntentName                string `json:"intentName"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
	WebhookUsed               string `json:"webhookUsed"`
}

// Status wraps API.ai response status
type Status struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorID      string `json:"errorId"`
	ErrorDetails string `json:"errorDetails"`
}

// Message is the returned output. Type is for various message types. The will be added soon.
type Message struct {
	Speech string `json:"speech"`
	Type   int    `json:"type"`
}

// Fulfillment is data about text response(s), rich messages, response received from webhook.
type Fulfillment struct {
	Messages []Message `json:"messages"`
	Speech   string    `json:"speech"`
}

// Result is the result of NLP
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

// ResponseStruct wraps the response from API.ai. Please see. https://docs.api.ai/docs/query#response
type ResponseStruct struct {
	ID        string `json:"id"`
	Language  string `json:"lang"`
	Result    Result `json:"result"`
	SessionID string `json:"sessionId"`
	Status    Status `json:"status"`
	Timestamp string `json:"timestamp"`
}

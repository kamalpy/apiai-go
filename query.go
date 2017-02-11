package apiaigo

// Context object
type Context struct {
	Lifespan   int               `json:"lifespan"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
}

// Location takes latitude and longitude
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Event can be used instead of query text
type Event struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// Entity wrapper of API.ai entities
type Entity struct {
	Count   int    `json:"count"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Preview string `json:"preview"`
}

// OrigReq ...
type OrigReq struct {
	Source string            `json:"source"`
	Data   map[string]string `json:"data"`
}

// QueryStruct wraps API.ai query. To know about all the parameters: https://docs.api.ai/docs/query#query-parameters-and-json-fields
type QueryStruct struct {
	Contexts        []Context `json:"contexts"`
	Language        string    `json:"lang"`
	Location        Location  `json:"location"`
	Query           string    `json:"query"`
	Event           Event     `json:"event"`
	SessionID       string    `json:"sessionId"`
	Timezone        string    `json:"timezone"`
	OriginalRequest OrigReq   `json:"originalRequest"`
	ResetContexts   bool      `json:"resetContexts"`
	Entities        []Entity  `json:"entities"`
}

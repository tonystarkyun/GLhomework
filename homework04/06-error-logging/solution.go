package q06errorlogging

import "encoding/json"

// ErrorPayload is the unified API error shape.
type ErrorPayload struct {
	Error string `json:"error"`
}

// ParseErrorBody decodes unified error response.
func ParseErrorBody(body []byte) (ErrorPayload, error) {
	var out ErrorPayload
	err := json.Unmarshal(body, &out)
	return out, err
}

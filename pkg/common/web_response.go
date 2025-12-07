package common

type WebResponse struct {
	Code       int    `json:"code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	MessageDev string `json:"message_dev"`
	Data       any    `json:"data,omitempty"`
}

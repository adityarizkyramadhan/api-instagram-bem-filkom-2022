package entity

type Response struct {
	Success interface{}            `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

type ResponseKedua struct {
	Body       map[string]interface{} `json:"body"`
	Status     string                 `json:"status"`
	StatusCode int                    `json:"status_code"`
}

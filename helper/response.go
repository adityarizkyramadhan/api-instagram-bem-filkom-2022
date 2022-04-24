package helper

type meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type response struct {
	Meta meta        `json:"meta"`
	Body interface{} `json:"body"`
}

func ResponseAPI(message string, status bool, data interface{}) response {
	meta := meta{
		Message: message,
		Status:  status,
	}
	response := response{
		Meta: meta,
		Body: data,
	}
	return response
}

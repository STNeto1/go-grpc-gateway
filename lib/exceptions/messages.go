package exceptions

type ValidationError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

type BadRequest struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type BadValidation struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Errors     []ValidationError `json:"errors"`
}

type NotFound struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type InternalServerError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Unauthorized struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

package schemas

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegistrationResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

type Response struct {
	Status int   `json:"status"`
	Error  error `json:"error,omitempty"`
}

type ErrorResponse struct {
	Status int   `json:"status"`
	Error  error `json:"error,omitempty"`
}

func ReturnErrResponse(status int, err error) ErrorResponse {
	return ErrorResponse{Status: status, Error: err}
}

package response

import (
	"encoding/json"
	"net/http"
)

// Response api object
type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code,omitempty"`
	Error      *Error      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Count      *int        `json:"count,omitempty"`
}

// Error api response Object
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewJSONResponse(w http.ResponseWriter, message string, status int, respStatus int, success bool, payload interface{}, httpErr interface{}) {

	response, err := json.Marshal(
		&Response{
			Success:    success,
			Message:    message,
			StatusCode: respStatus,
			Data:       &payload,
		},
	)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, status int, message string, details interface{}) {
	err := Error{
		Code:    status,
		Message: message,
	}

	NewJSONResponse(w, message, status, 0, false, details, err)
}

func JSONResponse(w http.ResponseWriter, message string, status int, respStatus int, payload interface{}) {

	NewJSONResponse(w, message, status, respStatus, true, payload, nil)
}

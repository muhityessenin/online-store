package response

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

}

func ParseResponse(resp *http.Response) (Response, error) {
	var res Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, &res)
	return res, err
}

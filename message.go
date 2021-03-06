package zend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type MessageService struct {
	client *Client
}

func (svc *MessageService) Send(recipient string, message string) (*Message, error) {

	/**
	* We have to marshal struct with information containing for
	* message endpoint to provide it as the request body to the
	* upstream.
	 */
	payload := map[string]interface{}{
		"to":      recipient,
		"from":    svc.client.Sender,
		"message": message,
	}

	json_object, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v%s/message", svc.client.URI, svc.client.Version), bytes.NewBuffer(json_object))
	if err != nil {
		return nil, err
	}

	/**
	* We can now safely dispatch the request to the upstream server
	* for furthur processing and execution of the task.
	 */
	response, err := svc.client.Dispatch(request)
	if err != nil {
		return nil, err
	}

	/**
	* Handle the return response according to the API and return the
	* response data back to the application.
	 */
	if response["status"] == "success" {
		return &Message{
			ID: uint64(response["data"].(map[string]interface{})["id"].(float64)),
		}, nil
	} else if response["status"] == "failed" {
		return nil, errors.New(response["error"].(map[string]interface{})["message"].(string))
	}

	/**
	* We got unexpected state and we have to return general error
	* back to the application.
	 */
	return nil, errors.New("unknown error")

}

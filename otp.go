package zend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Otp(recipient string) (*Otp, error) {

	/**
	* We have to marshal struct with information containing for
	* message endpoint to provide it as the request body to the
	* upstream.
	 */
	payload := map[string]string{
		"to":     recipient,
		"sender": c.Sender,
	}
	json_object, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v%s/otp/send", c.URI, c.Version), bytes.NewBuffer(json_object))
	if err != nil {
		return nil, err
	}

	/**
	* We can now safely dispatch the request to the upstream server
	* for furthur processing and execution of the task.
	 */
	response, err := c.Dispatch(request)
	if err != nil {
		return nil, err
	}

	/**
	* Handle the return response according to the API and return the
	* response data back to the application.
	 */
	fmt.Println()
	if response["status"] == "success" {
		return &Otp{
			ID: uint64(response["data"].(map[string]interface{})["otp"].(map[string]interface{})["id"].(float64)),
			Time: OtpTime{
				Expire:  uint64(response["data"].(map[string]interface{})["otp"].(map[string]interface{})["expire"].(float64)),
				Created: uint64(response["data"].(map[string]interface{})["otp"].(map[string]interface{})["created"].(float64)),
			},
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

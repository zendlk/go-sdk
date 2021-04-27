package zend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type OtpService struct {
	client *Client
}

func (svc *OtpService) Send(recipient string) (*Otp, error) {

	/**
	* We have to marshal struct with information containing for
	* message endpoint to provide it as the request body to the
	* upstream.
	 */
	payload, err := json.Marshal(map[string]string{
		"to":     recipient,
		"sender": svc.client.Sender,
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v%s/otp/send", svc.client.URI, svc.client.Version), bytes.NewBuffer(payload))
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
		return &Otp{
			ID: uint64(response["data"].(map[string]interface{})["otp"].(map[string]interface{})["id"].(float64)),
			Time: Times{
				Expires: uint64(response["data"].(map[string]interface{})["otp"].(map[string]interface{})["expire"].(float64)),
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

func (svc *OtpService) Verify(otp_id string, code string) (bool, error) {

	/**
	* We have to marshal struct with information containing for
	* message endpoint to provide it as the request body to the
	* upstream.
	 */
	payload, err := json.Marshal(map[string]string{
		"id":   otp_id,
		"code": code,
	})
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v%s/otp/verify", svc.client.URI, svc.client.Version), bytes.NewBuffer(payload))
	if err != nil {
		return false, err
	}

	/**
	* We can now safely dispatch the request to the upstream server
	* for furthur processing and execution of the task.
	 */
	response, err := svc.client.Dispatch(request)
	if err != nil {
		return false, err
	}

	/**
	* Handle the return response according to the API and return the
	* response data back to the application.
	 */
	if response["status"] == "success" {
		return true, nil
	} else if response["status"] == "failed" {
		return false, errors.New(response["error"].(map[string]interface{})["message"].(string))
	}

	/**
	* We got unexpected state and we have to return general error
	* back to the application.
	 */
	return false, errors.New("unknown error")
}

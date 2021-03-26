package zend

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func (c *Client) Dispatch(request *http.Request) ([]byte, error) {

	/**
	* We have to set headers before we dispatch the http request
	* to the origin to append authentication token, etc.
	*/
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "zend/go-sdk")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { return nil, err }
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil { return nil, err }

	return body, nil
}

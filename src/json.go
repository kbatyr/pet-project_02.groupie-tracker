package grpt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Gets info about URL page (headers, body(content), etc.)
func GetRequestApi(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return body
}

// Decoding json file to golang structures
func DecodeAPI(body []byte, data interface{}) error {

	if err := json.Unmarshal(body, data); err != nil {
		fmt.Println(err.Error())
		return errors.New("can not unmarshal JSON")
	}
	return nil
}

package client

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func setup() {
	fmt.Println("Max your test character here!")
}

// TODO: if we are in dev mode, register new user, and set the key,token,client in the apiClient
//  https://habitica.com/api/v4/user/auth/local/register
// http://localhost:8080/api/v4/user/auth/local/register

func RegisterNewUser(client *HabiticaApiClient) (*http.Response, error) {
	// TODO: check if test server is running, and if user already exists?
	//   an error in the response might just do this for us

	httpClient := http.Client{}
	jsonData := []byte(`{"username":"test","email":"test@test.com","password":"test","confirmPassword":"test"}`)
	request, err := http.NewRequest(http.MethodPost, client.BuildAddress("v4", "user/auth/local/register"), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("[ERROR] http.NewRequest:", err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("[ERROR] client.Do:", err)
		return nil, err
	}

	if response.StatusCode == http.StatusGatewayTimeout {
		return nil, errors.New("[ERROR] dev server not available")
	}
	return response, err

	// curl -v 'http://localhost:8080/api/v4/user/auth/local/register' -d '{"username":"test","email":"test@test.com","password":"test","confirmPassword":"test"}' -H "Content-Type: application/json"
	//
	//Success 200
	//Field	Type	Description
	//data	Object
	//The user object, if local auth was just attached to a social user then only user.auth.local
}

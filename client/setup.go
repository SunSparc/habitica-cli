package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"habitica-cli/database"
)

type HabiticaApiClient struct {
	Mode      string
	User      string
	Key       string
	XClientID string // https://habitica.fandom.com/wiki/Guidance_for_Comrades#X-Client_Header
	Host      string

	HabiticaUser database.HabiticaUser
}

func NewHabiticaApiClient(mode string) *HabiticaApiClient {
	hbc := &HabiticaApiClient{
		Mode: mode,
	}
	if err := hbc.Config(); err != nil {
		log.Println("[ERROR] in the config:", err)
	}
	if err := hbc.GetHabiticaUser(); err != nil {
		log.Println("[ERROR] in the gethabiticauser:", err)
	}

	return hbc
}

func (this *HabiticaApiClient) asdf() {
	// TODO: decide if we need separate live and dev modes
	//       things that are done on live can also be done on dev
	//       things that are done on dev may or may not be able to be done on live
	//       setup could take care of all the nuances between the two different modes
	//             and then just return a single object which can be do everything
	live := NewLiveMode(this.HabiticaUser)
	// or
	dev := NewDevMode(this.HabiticaUser)
}

// TODO: apiClient.storeHabiticaUser()
//       store user locally (?? every time? overwrite previous copy?) why?

func (this *HabiticaApiClient) GetHabiticaUser() error {
	response, err := this.doAPIGetRequest("user")
	if err != nil {
		return errors.New("unable to do Habitica API request")
	}
	if response.StatusCode != http.StatusOK {
		log.Println("got a bad response:", response.StatusCode, response.Status)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[ERROR] ioutil.ReadAll:", err)
	}

	this.HabiticaUser = database.HabiticaUser{}
	err = json.Unmarshal(responseBytes, &this.HabiticaUser)
	if err != nil {
		log.Println("[ERROR] json.Unmarshal:", err)
	}

	//log.Printf("habitica user: %#v\n", this.HabiticaUser)
	return nil
}

func (this *HabiticaApiClient) doAPIGetRequest(action string) (*http.Response, error) {
	client := http.Client{}
	// #live# request, err := http.NewRequest(http.MethodGet, "https://habitica.com/api/v3/user", nil)
	request, err := http.NewRequest(http.MethodGet, this.BuildAddress("v3", action), nil)
	if err != nil {
		log.Println("[ERROR] http.NewRequest:", err)
	}

	request.Header.Set("x-api-user", this.User)
	request.Header.Set("x-api-key", this.Key)
	request.Header.Set("x-client", this.XClientID)

	response, err := client.Do(request)
	if err != nil {
		log.Println("[ERROR] client.Do:", err)
	}
	return response, err
}

func (this *HabiticaApiClient) BuildAddress(version, action string) string {
	return this.Host + path.Join(API_PATH, version, action)
}

func (this *HabiticaApiClient) Config() error {
	if this.Mode == "live" {
		return this.ConfigLive()
	} else {
		return this.ConfigDev()
	}
}

func (this *HabiticaApiClient) ConfigLive() error {
	this.Host = LIVE_HOST
	this.User = os.Getenv("HABITICA_API_USER")
	this.Key = os.Getenv("HABITICA_API_KEY")
	this.XClientID = os.Getenv("HABITICA_API_CLIENT")
	if this.User == "" || this.Key == "" || this.XClientID == "" {
		return errors.New("[ERROR] user id, key, and client are required")
	}
	return nil
}
func (this *HabiticaApiClient) ConfigDev() error {
	this.Host = DEV_HOST

	this.User = "HABITICA_API_USER"
	this.Key = "HABITICA_API_KEY"
	this.XClientID = "HABITICA_API_CLIENT"

	response, err := RegisterNewUser(this)
	if err != nil {
		log.Println("[ERROR] devmode.RegisterNewUser:", TranslateError(err.Error()))
		return err
	}
	log.Println("[INFO] devmode.RegisterNewUser response:", response)
	return nil
}

// TODO: move RegisterNewUser here for Dev setup

const (
	LIVE_HOST string = "https://habitica.com/"
	DEV_HOST  string = "http://localhost:8080/"
	API_PATH  string = "api"

	DEV_USERNAME string = "test"
	DEV_PASSWORD string = "test"
	DEV_EMAIL    string = "test@test.com"
)

///////////////

func TranslateError(errorMessage string) string {
	switch errorMessage {
	case "connection refused":
		return "Either the requested server does not exist, or the connection is being blocked by a firewall."
	default:
		// We do not know what this error means.
	}
	return ""
}

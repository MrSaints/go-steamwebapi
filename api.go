package godoto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const API_HOST = "api.steampowered.com"
const DOTA_ID = "570"

type WebAPI struct {
	Endpoint string
	Params   url.Values
	Version  int
	Language string
	Key      string
}

type Result struct {
	Data json.RawMessage `json:"result"`
}

func DotaAPI(method string, match bool) (webAPI *WebAPI) {
	endpoint := "/"

	if !match {
		endpoint += "IEconDOTA2"
	} else {
		endpoint += "IDOTA2Match"
	}

	webAPI = new(WebAPI)
	webAPI.Endpoint = endpoint + "_" + DOTA_ID + "/" + method
	return
}

func (this *WebAPI) GetResult() (result Result) {
	// Set defaults
	if this.Version == 0 {
		this.Version = 1
	}

	if this.Language == "" {
		this.Language = "en"
	}

	if this.Key == "" {
		this.Key = os.Getenv("STEAM_API_KEY")
	}

	// Build Web API URL
	apiURL := url.URL{}
	apiURL.Scheme = "https"
	apiURL.Host = API_HOST
	apiURL.Path = this.Endpoint + "/v" + strconv.Itoa(this.Version)

	if this.Params == nil {
		this.Params = url.Values{}
	}

	this.Params.Set("key", this.Key)
	this.Params.Set("language", this.Language)
	apiURL.RawQuery = this.Params.Encode()

	// GET request
	res, err := http.Get(apiURL.String())
	failOnError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	failOnError(err)

	err = json.Unmarshal(body, &result)
	failOnError(err)

	return
}

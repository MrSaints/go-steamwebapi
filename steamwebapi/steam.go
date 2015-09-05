package steamwebapi

import (
	"encoding/json"
	//"log"
	"net/http"
	"net/url"
	"os"
)

// See https://wiki.teamfortress.com/wiki/WebAPI for more information.
const (
	defaultBaseURL         = "https://api.steampowered.com"
	dotaId                 = "570"
	baseDOTA2Endpoint      = "IEconDOTA2_" + dotaId
	baseDOTA2MatchEndpoint = "IDOTA2Match_" + dotaId
)

type Client struct {
	BaseURL  *url.URL
	Language string
	Key      string

	DOTA2        *DOTA2Services
	DOTA2Matches *DOTA2MatchesServices
}

type Result struct {
	Data interface{} `json:"result"`
}

/*
 Returns a new Steam Web API client.
 You can specify your API key as a parameter (k) or leave it blank
 to use the `STEAM_API_KEY` environment variable.
*/
func NewClient(k string) *Client {
	c := new(Client)
	c.BaseURL, _ = url.Parse(defaultBaseURL)
	c.Language = "en"

	c.Key = k
	if k == "" {
		c.Key = os.Getenv("STEAM_API_KEY")
	}

	c.DOTA2 = &DOTA2Services{client: c}
	c.DOTA2Matches = &DOTA2MatchesServices{client: c}

	return c
}

/*
 Creates and sends an API request using the specified endpoint (e),
 params (p), and interface (v). The latter will be used to deconstruct,
 and store the JSON result.
*/
func (c *Client) Get(e string, p url.Values, v interface{}) (*http.Response, error) {
	rel, err := url.Parse(e)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	if p == nil {
		p = url.Values{}
	}

	p.Set("key", c.Key)
	p.Set("language", c.Language)
	u.RawQuery = p.Encode()

	//log.Printf("API requested: %v", u.String())

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&Result{v})
	if err != nil {
		return nil, err
	}

	return res, err
}

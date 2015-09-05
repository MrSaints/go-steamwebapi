package steamwebapi

import (
	"net/url"
	"strconv"
)

// A service to handle methods related to Dota 2.
type DOTA2Services struct {
	client *Client
}

type Heroes struct {
	Heroes []Hero `json:"heroes"`
	Count  int    `json:"count"`
}

type Hero struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	LocalizedName string `json:"localized_name,omitempty"`
}

/*
 Returns a list of heroes within Dota 2.
 See https://wiki.teamfortress.com/wiki/WebAPI/GetHeroes for more information.
*/
func (s *DOTA2Services) GetHeroes() (Heroes, error) {
	heroes := new(Heroes)
	_, err := s.client.Get(baseDOTA2Endpoint+"/GetHeroes/v1", nil, heroes)
	return *heroes, err
}

/*
 Returns the current prize pool for specific tournaments.
 See https://wiki.teamfortress.com/wiki/WebAPI/GetTournamentPrizePool for more information.
*/
func (s *DOTA2Services) GetTournamentPrizePool(leagueId int) (float64, error) {
	params := url.Values{}
	params.Set("leagueid", strconv.Itoa(leagueId))

	var data map[string]interface{}
	_, err := s.client.Get(baseDOTA2Endpoint+"/GetTournamentPrizePool/v1", params, &data)
	if err != nil {
		return 0, err
	}

	return data["prize_pool"].(float64), nil
}

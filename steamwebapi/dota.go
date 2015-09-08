package steamwebapi

import (
	"net/url"
	"strconv"
)

// DOTA2Services handles communication with the
// IDOTA2 related methods of the Steam Web API.
//
// Steam Web API docs: https://wiki.teamfortress.com/wiki/WebAPI#Dota_2
type DOTA2Services struct {
	client *Client
}

// Heroes represents a list of heroes within Dota 2
type Heroes struct {
	Heroes []Hero `json:"heroes"`
	Count  int    `json:"count"`
}

// Hero represents information about a single hero.
type Hero struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name,omitempty"`
}

// GetHeroes returns a list of heroes within Dota 2.
// See https://wiki.teamfortress.com/wiki/WebAPI/GetHeroes for more information.
func (s *DOTA2Services) GetHeroes() (Heroes, error) {
	heroes := new(Heroes)
	_, err := s.client.Get(baseDOTA2Endpoint+"/GetHeroes/v1", nil, heroes)
	return *heroes, err
}

// GetTournamentPrizePool returns the current prize pool for specific tournaments.
// See https://wiki.teamfortress.com/wiki/WebAPI/GetTournamentPrizePool for more information.
func (s *DOTA2Services) GetTournamentPrizePool(leagueID int) (float64, error) {
	params := url.Values{}
	params.Set("leagueid", strconv.Itoa(leagueID))

	var data map[string]interface{}
	_, err := s.client.Get(baseDOTA2Endpoint+"/GetTournamentPrizePool/v1", params, &data)
	if err != nil {
		return 0, err
	}

	return data["prize_pool"].(float64), nil
}

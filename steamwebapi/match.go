package steamwebapi

import (
	"net/url"
	"strconv"
)

// A service to handle methods related to Dota 2 matches.
type DOTA2MatchesServices struct {
	client *Client
}

type MatchHistory struct {
	Status    int     `json:"status"`
	Limit     int     `json:"num_results"`
	Total     int     `json:"total_results"`
	Remaining int     `json:"results_remaining"`
	Matches   []Match `json:"matches"`
}

type Match struct {
	Id       int      `json:"match_id"`
	Sequence int      `json:"match_seq_num"`
	Start    int      `json:"start_time"`
	Type     int      `json:"lobby_type"`
	Players  []Player `json:"players"`
}

type Player struct {
	Id       int `json:"account_id"`
	Position int `json:"player_slot"`
	Hero     int `json:"hero_id"`
}

type Matches []MatchDetails

type MatchDetails struct {
	Players         []PlayerDetails `json:"players"`
	RadiantWin      bool            `json:"radiant_win"`
	Duration        int             `json:"duration"`
	Start           int             `json:"start_time"`
	Id              int             `json:"match_id"`
	Sequence        int             `json:"match_seq_num"`
	RadiantTower    int             `json:"tower_status_radiant"`
	DireTower       int             `json:"tower_status_dire"`
	RadiantBarracks int             `json:"barracks_status_radiant"`
	DireBarracks    int             `json:"barracks_status_dire"`
	Cluster         int             `json:"cluster"`
	FirstBlood      int             `json:"first_blood_time"`
	Type            int             `json:"lobby_type"`
	HumanPlayers    int             `json:"human_players"`
	League          int             `json:"leagueid"`
	Positive        int             `json:"positive_votes"`
	Negative        int             `json:"negative_votes"`
	GameMode        int             `json:"game_mode"`
	//Drafts          Draft           `json:"picks_ban"`
}

type PlayerDetails struct {
	Id           int       `json:"account_id"`
	Position     int       `json:"player_slot"`
	Hero         int       `json:"hero_id"`
	Item0        int       `json:"item_0"`
	Item1        int       `json:"item_1"`
	Item2        int       `json:"item_2"`
	Item3        int       `json:"item_3"`
	Item4        int       `json:"item_4"`
	Item5        int       `json:"item_5"`
	Kills        int       `json:"kills"`
	Deaths       int       `json:"deaths"`
	Assists      int       `json:"assists"`
	LeaverStatus int       `json:"leaver_status"`
	Gold         int       `json:"gold"`
	LH           int       `json:"last_hits"`
	DH           int       `json:"denies"`
	GPM          int       `json:"gold_per_min"`
	XPM          int       `json:"xp_per_min"`
	GS           int       `json:"gold_spent"`
	HD           int       `json:"hero_damage"`
	TD           int       `json:"tower_damage"`
	HH           int       `json:"hero_healing"`
	Level        int       `json:"level"`
	Abilities    []Ability `json:"ability_upgrades"`
}

type Ability struct {
	Id    int `json:"ability"`
	Time  int `json:"time"`
	Level int `json:"level"`
}

type Leagues struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	Name        string `json:"name"`
	Id          int    `json:"leagueid"`
	Description string `json:"description"`
	URL         string `json:"tournament_url"`
}

/*
 * Returns a list of matches, filterable by various parameters.
 * See https://wiki.teamfortress.com/wiki/WebAPI/GetMatchHistory for more information.
 */
func (s *DOTA2MatchesServices) GetMatchHistory(accountID int, gameMode int, skill int, heroID int, minPlayers int, leagueID int, startAtMatchID int, limit int, tournamentOnly bool) *MatchHistory {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("game_mode", strconv.Itoa(gameMode))
	params.Set("skill", strconv.Itoa(skill))
	params.Set("hero_id", strconv.Itoa(heroID))
	params.Set("min_players", strconv.Itoa(minPlayers))
	params.Set("league_id", strconv.Itoa(leagueID))
	params.Set("start_at_match_id", strconv.Itoa(startAtMatchID))

	if limit > 0 {
		params.Set("matches_requested", strconv.Itoa(limit))
	} else {
		params.Set("matches_requested", "5")
	}

	if tournamentOnly {
		params.Set("tournament_games_only", "1")
	}

	history := new(MatchHistory)
	_, err := s.client.Get(baseDOTA2MatchEndpoint+"/GetMatchHistory/v1", params, history)
	failOnError(err)

	return history
}

/*
 * Returns information about a particular match.
 * See https://wiki.teamfortress.com/wiki/WebAPI/GetMatchDetails for more information.
 */
func (s *DOTA2MatchesServices) GetMatchDetails(matchID int) *MatchDetails {
	params := url.Values{}
	params.Set("match_id", strconv.Itoa(matchID))

	match := new(MatchDetails)
	_, err := s.client.Get(baseDOTA2MatchEndpoint+"/GetMatchDetails/v1", params, match)
	failOnError(err)

	return match
}

/*
 * Returns a player's team and position.
 * The player's slot is stored as an 8-bit uint: 0 0 0 0 0 0 0 0.
 * The first bit (LtR) represents the player's team (i.e. 1 / True = Dire).
 * The final 3 bits represents the player's position within a team.
 */
func (p PlayerDetails) GetPosition() (bool, int) {
	isDire := false
	if (p.Position >> 7) == 1 {
		isDire = true
	}
	position := p.Position & 111
	return isDire, position
}

func (m MatchDetails) GetPosition(accountID int) (bool, int) {
	isDire, position := false, 0
	for _, player := range m.Players {
		if accountID == player.Id {
			return player.GetPosition()
		}
	}
	return isDire, position
}

/*
 * Returns information about DotaTV-supported leagues.
 * See https://wiki.teamfortress.com/wiki/WebAPI/GetLeagueListing for more information.
 */
func (s *DOTA2MatchesServices) GetLeagueListing() *Leagues {
	leagues := new(Leagues)
	_, err := s.client.Get(baseDOTA2MatchEndpoint+"/GetLeagueListing/v1", nil, leagues)
	failOnError(err)
	return leagues
}

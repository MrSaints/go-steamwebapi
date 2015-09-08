package steamwebapi

import (
	//"log"
	"net/url"
	"sort"
	"strconv"
	"sync"
)

// DOTA2MatchesServices handles communication with the
// IDOTA2Match related methods of the Steam Web API.
//
// Steam Web API docs: https://wiki.teamfortress.com/wiki/WebAPI#Dota_2
type DOTA2MatchesServices struct {
	client *Client
}

// MatchHistory represents a filtered list of Dota 2 matches.
// It is often used for a player's match history (recent matches).
type MatchHistory struct {
	Status    int     `json:"status"`
	Limit     int     `json:"num_results"`
	Total     int     `json:"total_results"`
	Remaining int     `json:"results_remaining"`
	Matches   []Match `json:"matches"`
}

// Match represents a single Dota 2 match from
// a filtered list of matches in a MatchHistory.
// It is not as detailed as MatchDetails, but it is all that is
// offered from the MatchHistory API.
type Match struct {
	ID       int      `json:"match_id"`
	Sequence int      `json:"match_seq_num"`
	Start    int      `json:"start_time"`
	Type     int      `json:"lobby_type"`
	Players  []Player `json:"players"`
}

// Player represents a single Dota 2 player from
// a list of players in a Match.
type Player struct {
	ID       int `json:"account_id"`
	Position int `json:"player_slot"`
	Hero     int `json:"hero_id"`
}

// Matches represents a list of detailed matches.
type Matches []MatchDetails

// MatchDetails represents information about a particular match.
type MatchDetails struct {
	Players         []PlayerDetails `json:"players"`
	RadiantWin      bool            `json:"radiant_win"`
	Duration        int             `json:"duration"`
	Start           int             `json:"start_time"`
	ID              int             `json:"match_id"`
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

// PlayerDetails represents information about a player in a particular match.
type PlayerDetails struct {
	ID           int       `json:"account_id"`
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

// Ability represents information about a player's ability upgrades
// (e.g. what, and when an ability was learnt).
type Ability struct {
	ID    int `json:"ability"`
	Time  int `json:"time"`
	Level int `json:"level"`
}

// Leagues represents a list of DotaTV-supported leagues.
type Leagues struct {
	Leagues []League `json:"leagues"`
}

// League represents information about a DotaTV-supported league.
type League struct {
	Name        string `json:"name"`
	ID          int    `json:"leagueid"`
	Description string `json:"description"`
	URL         string `json:"tournament_url"`
}

// GetMatchHistory returns a list of matches, filterable by various parameters.
// See https://wiki.teamfortress.com/wiki/WebAPI/GetMatchHistory for more information.
func (s *DOTA2MatchesServices) GetMatchHistory(accountID int, gameMode int, skill int, heroID int, minPlayers int, leagueID int, startAtMatchID int, limit int, tournamentOnly bool) (MatchHistory, error) {
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

	return *history, err
}

// GetMatchDetails returns information about a particular match.
// See https://wiki.teamfortress.com/wiki/WebAPI/GetMatchDetails for more information.
func (s *DOTA2MatchesServices) GetMatchDetails(matchID int) (MatchDetails, error) {
	params := url.Values{}
	params.Set("match_id", strconv.Itoa(matchID))

	match := new(MatchDetails)
	_, err := s.client.Get(baseDOTA2MatchEndpoint+"/GetMatchDetails/v1", params, match)

	return *match, err
}

// GetDetails returns detailed information about a Match from MatchHistory
// using GetMatchDetails.
// It requires a DOTA2MatchesServices client.
func (m Match) GetDetails(s *DOTA2MatchesServices) (MatchDetails, error) {
	return s.GetMatchDetails(m.ID)
}

// GetDetails returns detailed information about Matches in a MatchHistory
// using using GetDetails.
// It requires a DOTA2MatchesServices client.
func (h MatchHistory) GetDetails(s *DOTA2MatchesServices) Matches {
	history := h.Matches
	total := len(history)

	out := make(chan MatchDetails, total)
	var wg sync.WaitGroup
	wg.Add(total)

	for _, m := range history {
		go func(m Match, s *DOTA2MatchesServices) {
			// Supress errors (TODO: terminate on error)
			md, _ := m.GetDetails(s)
			out <- md
			wg.Done()
		}(m, s)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var matches Matches
	for m := range out {
		matches = append(matches, m)
	}
	sort.Sort(ByMatchID(matches))
	return matches
}

// GetPosition returns a player's team, and position from PlayerDetails.
// The player's slot is stored as an 8-bit uint: 00000000.
// The first bit (LtR) represents the player's team (i.e. 1 / True = Dire).
// The final 3 bits represents the player's position within a team.
func (p PlayerDetails) GetPosition() (bool, int) {
	isDire := false
	if (p.Position >> 7) == 1 {
		isDire = true
	}
	position := p.Position & 111
	return isDire, position
}

// GetPositionByAccount returns a player's team, and position
// using their account ID from MatchDetails.
func (m MatchDetails) GetPositionByAccount(aid int32) (bool, int) {
	isDire, pos := false, 0
	for _, p := range m.Players {
		if aid == int32(p.ID) {
			return p.GetPosition()
		}
	}
	return isDire, pos
}

// CommunityID returns a 32-bit Steam community ID from a 64-bit account ID.
func CommunityID(aid int64) int32 {
	return int32(aid)
}

// GetLeagueListing returns information about DotaTV-supported leagues.
// See https://wiki.teamfortress.com/wiki/WebAPI/GetLeagueListing for more information.
func (s *DOTA2MatchesServices) GetLeagueListing() (Leagues, error) {
	leagues := new(Leagues)
	_, err := s.client.Get(baseDOTA2MatchEndpoint+"/GetLeagueListing/v1", nil, leagues)

	return *leagues, err
}

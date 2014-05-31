package godoto

import (
    "encoding/json"
    "net/url"
    "strconv"
)

type MatchHistory struct {
    Status      int         `json:"status"`
    Limit       int         `json:"num_results"`
    Total       int         `json:"total_results"`
    Remaining   int         `json:"results_remaining"`
    Matches     []Match     `json:"matches"`
}

type Match struct {
    Id          int         `json:"match_id"`
    Sequence    int         `json:"match_seq_num"`
    Start       int         `json:"start_time"`
    Type        int         `json:"lobby_type"`
    Players     []Player    `json:"players"`
}

type Player struct {
    Id          int         `json:"account_id"`
    Position    int         `json:"player_slot"`
    Hero        int         `json:"hero_id"`
}

func GetMatchHistory(accountID int, gameMode int, skill int, heroID int, minPlayers int, leagueID int, startAtMatchID int, limit int, tournamentOnly bool) (history MatchHistory) {
    api := DotaAPI("GetMatchHistory", true)

    api.Params = url.Values{}
    api.Params.Set("account_id", strconv.Itoa(accountID))
    api.Params.Set("game_mode", strconv.Itoa(gameMode))
    api.Params.Set("skill", strconv.Itoa(skill))
    api.Params.Set("hero_id", strconv.Itoa(heroID))
    api.Params.Set("min_players", strconv.Itoa(minPlayers))
    api.Params.Set("league_id", strconv.Itoa(leagueID))
    api.Params.Set("start_at_match_id", strconv.Itoa(startAtMatchID))

    if limit > 0 {
        api.Params.Set("matches_requested", strconv.Itoa(limit))
    }

    if tournamentOnly {
        api.Params.Set("tournament_games_only", "1")
    }

    result := api.GetResult()

    history = MatchHistory{}
    err := json.Unmarshal(result.Data, &history)
    pError(err)
    return
}
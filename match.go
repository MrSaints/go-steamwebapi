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
    Id              int         `json:"account_id"`
    Position        int         `json:"player_slot"`
    Hero            int         `json:"hero_id"`
    Item0           int         `json:"item_0"`
    Item1           int         `json:"item_1"`
    Item2           int         `json:"item_2"`
    Item3           int         `json:"item_3"`
    Item4           int         `json:"item_4"`
    Item5           int         `json:"item_5"`
    Kills           int         `json:"kills"`
    Deaths          int         `json:"deaths"`
    Assists         int         `json:"assists"`
    LeaverStatus    int         `json:"leaver_status"`
    Gold            int         `json:"gold"`
    LH              int         `json:"last_hits"`
    DH              int         `json:"denies"`
    GPM             int         `json:"gold_per_min"`
    XPM             int         `json:"xp_per_min"`
    GS              int         `json:"gold_spent"`
    HD              int         `json:"hero_damage"`
    TD              int         `json:"tower_damage"`
    HH              int         `json:"hero_healing"`
    Level           int         `json:"level"`
    Abilities       []Ability   `json:"ability_upgrades"`
}

type Ability struct {
    Id      int `json:"ability"`
    Time    int `json:"time"`
    Level   int `json:"level"`
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
    } else {
        api.Params.Set("matches_requested", "5")
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

func GetMatchDetails(matchID int) (match MatchDetails) {
    api := DotaAPI("GetMatchDetails", true)
    api.Params = url.Values{}
    api.Params.Set("match_id", strconv.Itoa(matchID))

    result := api.GetResult()

    match = MatchDetails{}
    err := json.Unmarshal(result.Data, &match)
    pError(err)
    return
}

func (this Match) GetDetails() MatchDetails {
    return GetMatchDetails(this.Id)
}

func (this MatchHistory) GetDetails() (matches Matches) {
    history := this.Matches
    for _, element := range history {
        matches = append(matches, element.GetDetails())
    }
    return
}

func (this PlayerDetails) GetPosition() (isDire bool, position int) {
    isDire = false
    if (this.Position & (1 << 7)) >> 7 == 1 {
        isDire = true
    }
    position = this.Position & 111
    return
}

func (this MatchDetails) GetPosition(accountID int) (isDire bool, position int) {
    isDire, position = false, 0
    // IMPL QUICK BOOL SEARCH
    for _, player := range this.Players {
        if accountID == player.Id {
            isDire, position = player.GetPosition() 
        }
    }
    return
}
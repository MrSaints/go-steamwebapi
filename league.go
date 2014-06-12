package godoto

import (
    "encoding/json"
)

type Leagues struct {
    Leagues []League    `json:"leagues"`
}

type League struct {
    Name        string  `json:"name"`
    Id          int     `json:"leagueid"`
    Description string  `json:"description"`
    URL         string  `json:"tournament_url"`
}

func GetLeagueListing() (leagues Leagues) {
    result := DotaAPI("GetLeagueListing", true).GetResult()

    err := json.Unmarshal(result.Data, &leagues)
    pError(err)
    return
}

func GetTournamentPrizePool(leagueID int) interface{} {
    result := DotaAPI("GetTournamentPrizePool", false).GetResult()
    var data map[string]interface{}
    err := json.Unmarshal(result.Data, &data)
    pError(err)
    return data["prize_pool"]
}
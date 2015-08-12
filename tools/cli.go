package main

import (
    "os"
    "github.com/codegangsta/cli"
)

func main() {
    commands := []cli.Command{
        {
            Name: "GetHeroes",
            ShortName: "h",
            Usage: "Returns a list of heroes within Dota 2",
            Flags: []cli.Flag{
                cli.BoolFlag{Name: "msgpack", Usage: "Encodes and dumps data in MsgPack"},
            },
            Action: GetHeroes,
        },
        {
            Name: "GetLeagueListing",
            ShortName: "l",
            Usage: "Returns information about DotaTV-supported leagues",
            Flags: []cli.Flag{
                cli.BoolFlag{Name: "msgpack", Usage: "Encodes and dumps data in MsgPack"},
            },
            Action: GetLeagueListing,
        },
        {
            Name: "GetTournamentPrizePool",
            ShortName: "pp",
            Usage: "Returns the current prizepool for specific tournaments",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "id", Value: 600, Usage: "A list of league IDs can be found via the GetLeagueListing method"},
            },
            Action: GetTournamentPrizePool,
        },
        {
            Name: "GetMatchHistory",
            ShortName: "mh",
            Usage: "Returns a list of matches, filterable by various parameters",
            Flags: []cli.Flag{
                cli.IntFlag{Name: "accountId", Value: 0, Usage: "32-bit account ID"},
                cli.IntFlag{Name: "gameMode", Value: 0, Usage: "Game mode"},
                cli.IntFlag{Name: "skill", Value: 0, Usage: "Skill bracket for matches"},
                cli.IntFlag{Name: "heroID", Value: 0, Usage: "A list of hero IDs can be found via the GetHeroes method"},
                cli.IntFlag{Name: "minPlayers", Value: 0, Usage: "Minimum amount of players in a match for the match to be returned"},
                cli.IntFlag{Name: "leagueID", Value: 0, Usage: "Only return matches from this league. A list of league IDs can be found via the GetLeagueListing method"},
                cli.IntFlag{Name: "startAtMatchID", Value: 0, Usage: "Start searching for matches equal to or older than this match ID"},
                cli.IntFlag{Name: "limit", Value: 0, Usage: "Amount of matches to include in results (default: 5)"},
                cli.BoolFlag{Name: "tournamentOnly", Usage: "Whether to limit results to tournament matches"},
                cli.BoolFlag{Name: "summary", Usage: "Include an overview of matches (e.g. W/L)"},
            },
            Action: GetMatchHistory,
        },
    }

    app := cli.NewApp()
    app.Name = "godoto cli"
    app.Usage = "Fetch Dota 2 data via Steam's Web API"
    app.Version = "2.0.0"
    app.Commands = commands
    app.Run(os.Args)
}
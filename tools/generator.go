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
                cli.BoolFlag{"msgpack", "Encodes and dumps data in MsgPack"},
            },
            Action: GetHeroes,
        },
        {
            Name: "GetLeagueListing",
            ShortName: "l",
            Usage: "Returns information about DotaTV-supported leagues",
            Flags: []cli.Flag{
                cli.BoolFlag{"msgpack", "Encodes and dumps data in MsgPack"},
            },
            Action: GetLeagueListing,
        },
        {
            Name: "GetTournamentPrizePool",
            ShortName: "pp",
            Usage: "Returns the current prizepool for specific tournaments",
            Flags: []cli.Flag{
                cli.IntFlag{"id", 600, "A list of league IDs can be found via the GetLeagueListing method"},
            },
            Action: GetTournamentPrizePool,
        },
        {
            Name: "GetMatchHistory",
            ShortName: "mh",
            Usage: "Returns a list of matches, filterable by various parameters",
            Flags: []cli.Flag{
                cli.IntFlag{"accountId", 0, "32-bit account ID"},
                cli.IntFlag{"gameMode", 0, "Game mode"},
                cli.IntFlag{"skill", 0, "Skill bracket for matches"},
                cli.IntFlag{"heroID", 0, "A list of hero IDs can be found via the GetHeroes method"},
                cli.IntFlag{"minPlayers", 0, "Minimum amount of players in a match for the match to be returned"},
                cli.IntFlag{"leagueID", 0, "Only return matches from this league. A list of league IDs can be found via the GetLeagueListing method"},
                cli.IntFlag{"startAtMatchID", 0, "Start searching for matches equal to or older than this match ID"},
                cli.IntFlag{"limit", 0, "Amount of matches to include in results (default: 5)"},
                cli.BoolFlag{"tournamentOnly", "Whether to limit results to tournament matches"},
                cli.BoolFlag{"summary", "Include an overview of matches (e.g. W/L)"},
            },
            Action: GetMatchHistory,
        },
    }

    app := cli.NewApp()
    app.Name = "GoDoto"
    app.Version = "1.0.0"
    app.Commands = commands
    app.Run(os.Args)
    os.Exit(0)
}
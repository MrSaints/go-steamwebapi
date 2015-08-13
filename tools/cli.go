package main

import (
    "io/ioutil"
    "log"
    "os"
    "github.com/codegangsta/cli"
    "github.com/mrsaints/godoto"
    "github.com/ugorji/go/codec"
)

var (
    client *godoto.Client
)

func Dump(fileName string, data interface{}) (err error) {
    var (
        packed []byte
        handle codec.MsgpackHandle
    )
    encoder := codec.NewEncoderBytes(&packed, &handle)
    encoder.Encode(data)

    err = ioutil.WriteFile(fileName, packed, 0644)
    return
}

func GetHeroes(c *cli.Context) {
    heroes := client.DOTA2.GetHeroes()

    if c.Bool("msgpack") {
        file_error := Dump("heroes.bin", heroes)
        if file_error != nil {
            panic(file_error)
        }

        log.Println("Encoded and dumped list of heroes in MsgPack.")
    } else {
        log.Println(heroes)
    }

    log.Printf("Total heroes: %v", heroes.Count)
}

func GetLeagueListing(c *cli.Context) {
    listing := client.DOTA2Matches.GetLeagueListing()

    if c.Bool("msgpack") {
        file_error := Dump("leagues.bin", listing.Leagues)
        if file_error != nil {
            panic(file_error)
        }

        log.Println("Encoded and dumped list of leagues in MsgPack.")
    } else {
        log.Println(listing)
    }

    log.Printf("Total leagues: %v", len(listing.Leagues))
}

func GetTournamentPrizePool(c *cli.Context) {
    prize := client.DOTA2.GetTournamentPrizePool(c.Int("leagueID"))
    log.Printf("Prize pool: %v", prize)
}

func GetMatchHistory(c *cli.Context) {
    accountID := c.Int("accountId")
    history := client.DOTA2Matches.GetMatchHistory(accountID, c.Int("gameMode"), c.Int("skill"), c.Int("heroID"), c.Int("minPlayers"), c.Int("leagueID"), c.Int("startAtMatchID"), c.Int("limit"), c.Bool("tournamentOnly"))

    log.Println(history)
}

func main() {
    client = godoto.NewClient("")

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
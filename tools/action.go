package main

import (
    "io/ioutil"
    "log"
    "github.com/codegangsta/cli"
    "github.com/mrsaints/godoto"
    "github.com/ugorji/go/codec"
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
    heroes := godoto.GetHeroes()

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
    listing := godoto.GetLeagueListing()

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
    prize := godoto.GetTournamentPrizePool(c.Int("leagueID"))
    log.Printf("Prize pool: %v", prize)
}

func GetMatchHistory(c *cli.Context) {
    accountID := c.Int("accountId")
    history := godoto.GetMatchHistory(accountID, c.Int("gameMode"), c.Int("skill"), c.Int("heroID"), c.Int("minPlayers"), c.Int("leagueID"), c.Int("startAtMatchID"), c.Int("limit"), c.Bool("tournamentOnly"))

    if c.Bool("summary") {
        result := ""
        matches := history.GetDetails()

        for _, match := range matches {
            if accountID == 0 {
                if match.RadiantWin {
                    result += "R"
                } else {
                    result += "D"
                }
                continue
            }

            isDire, _ := match.GetPosition(accountID)
            if match.RadiantWin && !isDire {
                result += "W"
            } else if !match.RadiantWin && isDire {
                result += "W"
            } else {
                result += "L"
            }
        }

        log.Println(result)
    } else {
        log.Println(history)
    }
}
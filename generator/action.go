package main

import (
    "io/ioutil"
    "log"
    "github.com/codegangsta/cli"
    "github.com/MrSaints/godoto"
    "github.com/ugorji/go/codec"
)

func GetHeroes(c *cli.Context) {
    heroes := godoto.GetHeroes()

    if c.Bool("msgpack") {
        var (
            packed []byte
            handle codec.MsgpackHandle
        )
        encoder := codec.NewEncoderBytes(&packed, &handle)
        encoder.Encode(heroes)

        file_error := ioutil.WriteFile("heroes.bin", packed, 0644)
        if file_error != nil {
            panic(file_error)
        }

        log.Print("Encoded and dumped list of heroes in MsgPack.")
    } else {
        log.Print(heroes)
    }

    log.Printf("Total heroes: %v", heroes.Count)
}

func GetMatchHistory(c *cli.Context) {
    matches := godoto.GetMatchHistory(c.Int("accountId"), c.Int("gameMode"), c.Int("skill"), c.Int("heroID"), c.Int("minPlayers"), c.Int("leagueID"), c.Int("startAtMatchID"), c.Int("limit"), c.Bool("tournamentOnly"))
    log.Print(matches)
}
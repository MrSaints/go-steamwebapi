package main

import (
    "io/ioutil"
	"log"
	"github.com/MrSaints/godoto"
	"github.com/ugorji/go/codec"
)

func main() {
	heroes := godoto.GetHeroes()

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

    log.Print("Complete!")
}
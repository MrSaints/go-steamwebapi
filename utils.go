package godoto

import (
    "log"
)

func failOnError(err error) {
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
}
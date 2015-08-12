package godoto

import (
    "log"
)

func failOnError(err error) {
    if err != nil {
        log.Fatalf(err)
        panic(err)
    }
}
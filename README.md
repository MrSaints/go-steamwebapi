# go-steamwebapi

A simple and idiomatic Go package for programmatically working with [Steam's Web API][] (particularly its [Dota 2 Web API][]).

## Installation

```shell
go get github.com/mrsaints/go-steamwebapi
```

## Usage

Import the `steamwebapi` package:

```go
import "github.com/mrsaints/go-steamwebapi/steamwebapi"
```

Construct a Steam Web API client:

```go
// Accepts an API key as a parameter
// Alternatively, you can set the API key via the STEAM_API_KEY environment variable
client = steamwebapi.NewClient("")
```

Using the newly constructed client, call methods via its respective service (`DOTA2` and `DOTA2Matches`):

```go
// Returns information about a particular match.
client.DOTA2Matches.GetMatchDetails(matchId)

// Returns a list of heroes within Dota 2.
client.DOTA2.GetHeroes()

// Returns the current prize pool for specific tournaments.
client.DOTA2.GetTournamentPrizePool(leagueId)
```

View the [GoDocs](https://godoc.org/github.com/MrSaints/go-steamwebapi/steamwebapi), [examples](https://github.com/MrSaints/go-steamwebapi/tree/master/examples) or [code](https://github.com/MrSaints/go-steamwebapi/tree/master/steamwebapi) for more information (i.e. available services / methods).


[Steam's Web API]: https://developer.valvesoftware.com/wiki/Steam_Web_API
[Dota 2 Web API]: https://wiki.teamfortress.com/wiki/WebAPI#Dota_2

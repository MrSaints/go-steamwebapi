package godoto

import (
    "testing"
)

func BenchmarkGetLeagueListing(b *testing.B) {
    for i := 0; i < b.N; i++ {
        GetLeagueListing()
    }
}

func BenchmarkGetTournamentPrizePool(b *testing.B) {
    for i := 0; i < b.N; i++ {
        GetTournamentPrizePool(600)
    }
}
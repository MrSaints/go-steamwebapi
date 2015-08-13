package steamwebapi

import (
	"testing"
)

func BenchmarkGetHeroes(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.DOTA2.GetHeroes()
	}
}

func BenchmarkGetTournamentPrizePool(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.DOTA2.GetTournamentPrizePool(600)
	}
}

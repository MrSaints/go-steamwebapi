package steamwebapi

import (
	"testing"
)

func BenchmarkGetMatchHistory(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.DOTA2Matches.GetMatchHistory(47724064, 0, 0, 0, 0, 0, 0, 0, false)
	}
}

func BenchmarkGetMatchDetails(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.DOTA2Matches.GetMatchDetails(714452368)
	}
}

func BenchmarkMatchHistory_GetDetails(b *testing.B) {
	setup()
	history := client.DOTA2Matches.GetMatchHistory(47724064, 0, 0, 0, 0, 0, 0, 0, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		history.GetDetails(client.DOTA2Matches)
	}
}

func BenchmarkMatchDetails_GetPositionByAccount(b *testing.B) {
	setup()
	match := client.DOTA2Matches.GetMatchDetails(714452368)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		match.GetPositionByAccount(7653193)
	}
}

func BenchmarkGetLeagueListing(b *testing.B) {
	setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.DOTA2Matches.GetLeagueListing()
	}
}

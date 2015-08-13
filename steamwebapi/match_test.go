package steamwebapi

import (
	"testing"
)

func BenchmarkGetMatchHistory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMatchHistory(47724064, 0, 0, 0, 0, 0, 0, 0, false)
	}
}

func BenchmarkGetMatchDetails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMatchDetails(714452368)
	}
}

func BenchmarkMatchHistoryGetDetails(b *testing.B) {
	history := GetMatchHistory(47724064, 0, 0, 0, 0, 0, 0, 0, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		history.GetDetails()
	}
}

func BenchmarkMatchDetailsGetPosition(b *testing.B) {
	match := GetMatchDetails(714452368)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		match.GetPosition(7653193)
	}
}

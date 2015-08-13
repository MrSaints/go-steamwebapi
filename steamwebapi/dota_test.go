package steamwebapi

import (
	"testing"
)

func BenchmarkGetHeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetHeroes()
	}
}

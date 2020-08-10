package wild_test

import (
	"fmt"
	"testing"

	"github.com/gertd/wild"
)

func TestMatchCI(t *testing.T) {
	for tid, tcase := range CaseInsensitiveTestCases() {
		tcase := tcase
		testID := fmt.Sprintf("%d", tid)
		t.Run(testID, func(t *testing.T) {
			if wild.Match(tcase.pattern, tcase.search, tcase.caseinsensitive) != tcase.expected {
				t.Logf("failed tid: %s tcase: %+v", testID, tcase)
				t.Fail()
			}
		})
	}
}

func BenchmarkMatchCI1000(b *testing.B)   { benchmarkMatchCI(b) }
func BenchmarkMatchCI10000(b *testing.B)  { benchmarkMatchCI(b) }
func BenchmarkMatchCI100000(b *testing.B) { benchmarkMatchCI(b) }

func benchmarkMatchCI(b *testing.B) {
	tcases := CaseInsensitiveTestCases()
	for n := 0; n < b.N; n++ {
		for tid, tcase := range tcases {
			if wild.Match(tcase.pattern, tcase.search, tcase.caseinsensitive) != tcase.expected {
				b.Logf("failed tid: %d tcase: %+v", tid, tcase)
			}
		}
	}
}

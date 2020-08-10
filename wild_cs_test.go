package wild_test

import (
	"fmt"
	"testing"

	"github.com/gertd/wild"
)

func TestMatchCS(t *testing.T) {
	for tid, tcase := range CaseSensitiveTestCases() {
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

func BenchmarkMatchCS1000(b *testing.B)   { benchmarkMatchCS(b) }
func BenchmarkMatchCS10000(b *testing.B)  { benchmarkMatchCS(b) }
func BenchmarkMatchCS100000(b *testing.B) { benchmarkMatchCS(b) }

func benchmarkMatchCS(b *testing.B) {
	tcases := CaseSensitiveTestCases()
	for n := 0; n < b.N; n++ {
		for tid, tcase := range tcases {
			if wild.Match(tcase.pattern, tcase.search, tcase.caseinsensitive) != tcase.expected {
				b.Logf("failed tid: %d tcase: %+v", tid, tcase)
			}
		}
	}
}

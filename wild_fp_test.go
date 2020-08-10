package wild_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

// TestFilepathMatchCS - calibration test comparing with filepath.Match() function (case-sensitive).
func TestFilepathMatchCS(t *testing.T) {
	for tid, tcase := range CaseSensitiveTestCases() {
		tcase := tcase
		testID := fmt.Sprintf("%d", tid)
		t.Run(testID, func(t *testing.T) {
			matched, err := filepath.Match(tcase.pattern, tcase.search)
			if err != nil {
				t.Logf("error tid: %s tcase: %+v error: %+v", testID, tcase, err)
				t.Fail()
			}
			if tcase.expected != matched {
				t.Logf("failed tid: %s tcase: %+v", testID, tcase)
				t.Fail()
			}
		})
	}
}

func BenchmarkFilepathMatchCS1000(b *testing.B)   { benchmarkFilepathMatchCS(b) }
func BenchmarkFilepathMatchCS10000(b *testing.B)  { benchmarkFilepathMatchCS(b) }
func BenchmarkFilepathMatchCS100000(b *testing.B) { benchmarkFilepathMatchCS(b) }

func benchmarkFilepathMatchCS(b *testing.B) {
	tcases := CaseSensitiveTestCases()
	for n := 0; n < b.N; n++ {
		for tid, tcase := range tcases {
			matched, err := filepath.Match(tcase.pattern, tcase.search)
			if err != nil {
				b.Logf("error tid: %d tcase: %+v error: %+v", tid, tcase, err)
			}
			if tcase.expected != matched {
				b.Logf("failed tid: %d tcase: %+v", tid, tcase)
			}
		}
	}
}

// matchCI - case-insensitive filepath.Match, by lower casing both inputs
func matchCI(pattern, name string) (matched bool, err error) {
	p := strings.ToLower(pattern)
	s := strings.ToLower(name)
	return filepath.Match(p, s)
}

// TestFilepathMatchCI - calibration test comparing with filepath.Match() function (case-insensitive).
func TestFilepathMatchCI(t *testing.T) {
	for tid, tcase := range CaseInsensitiveTestCases() {
		tcase := tcase
		testID := fmt.Sprintf("%d", tid)
		t.Run(testID, func(t *testing.T) {
			matched, err := matchCI(tcase.pattern, tcase.search)
			if err != nil {
				t.Logf("error tid: %s tcase: %+v error: %+v", testID, tcase, err)
				t.Fail()
			}
			if tcase.expected != matched {
				t.Logf("failed tid: %s tcase: %+v", testID, tcase)
				t.Fail()
			}
		})
	}
}

func BenchmarkFilepathMatchCI1000(b *testing.B)   { benchmarkFilepathMatchCI(b) }
func BenchmarkFilepathMatchCI10000(b *testing.B)  { benchmarkFilepathMatchCI(b) }
func BenchmarkFilepathMatchCI100000(b *testing.B) { benchmarkFilepathMatchCI(b) }

func benchmarkFilepathMatchCI(b *testing.B) {
	tcases := CaseInsensitiveTestCases()
	for n := 0; n < b.N; n++ {
		for tid, tcase := range tcases {

			matched, err := matchCI(tcase.pattern, tcase.search)
			if err != nil {
				b.Logf("error tid: %d tcase: %+v error: %+v", tid, tcase, err)
			}
			if tcase.expected != matched {
				b.Logf("failed tid: %d tcase: %+v", tid, tcase)
			}
		}
	}
}

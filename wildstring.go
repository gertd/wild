package wild

import "unicode"

const (
	null         = '\x00'
	asterisk     = '*'
	questionmark = '?'
)

// wildString - mimics null terminated string behavior
type wildString string

// Char - get character on index
func (s wildString) Char(i int) rune {
	if i < len(s) {
		// return unicode.ToLower(rune(s[i]))
		return rune(s[i])
	}
	return null
}

// IsEOS - returns true index position is at or past end of string.
func (s wildString) IsEOS(i int) bool {
	return !(i < len(s))
}

// IsAsterisk - returns true if index character is an asterisk (*).
func (s wildString) IsAsterisk(i int) bool {
	return i < len(s) && s[i] == asterisk
}

// IsQuestionMark - returns true if index character (i) is an questionmark (?).
func (s wildString) IsQuestionMark(i int) bool {
	return i < len(s) && s[i] == questionmark
}

// EqualCS - .
func EqualCS(p, s rune) bool {
	return p == s
}

// NotEqualCS - .
func NotEqualCS(p, s rune) bool {
	return p != s
}

// EqualCI - .
func EqualCI(p, s rune) bool {
	return toLower(p) == toLower(s)
}

// NotEqualCI - .
func NotEqualCI(p, s rune) bool {
	return toLower(p) != toLower(s)
}

// CompFn - rune comparison function.
type CompFn func(rune, rune) bool

func toLower(r rune) rune {
	if r > 64 && r < 91 {
		return r ^ 32
	}
	if r < 128 {
		return r
	}
	return unicode.ToLower(r)
}

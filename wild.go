package wild

const (
// asterisk     = '*'
// questionmark = '?'
)

// Match - wildcard match function (case-sensitive).
//
// adoptation of the wildcard match implementation of Kirk J. Krauss documented in Dr Dobbs
// https://www.drdobbs.com/architecture-and-design/matching-wildcards-an-empirical-way-to-t/240169123
//
// Input parameters:
//  pattern         - A (potentially) corresponding string with wildcards
//  search          - A string without wildcards
//  ci              - Case insensitive comparison
//
// Return value(s):
//  bool - pattern matched search string
func Match(pattern, search string, ci bool) bool {

	pStr := wildString(pattern)
	sStr := wildString(search)

	var (
		pIdx     int    = 0 // Index for both tame and wild strings in upper loop
		sIdx     int        // Index for tame string, set going into lower loop
		pSeq     int        // Index for prospective match after '*' (wild string)
		sSeq     int        // Index for prospective match (tame string)
		NotEqual CompFn     // Not equal rune comparison function, abstracting case-sensitivity
	)

	if ci {
		NotEqual = NotEqualCI
	} else {
		NotEqual = NotEqualCS
	}

	// Find a first wildcard, if one exists, and the beginning of any
	// prospectively matching sequence after it.
	for {
		// Check for the end from the start.  Get out fast, if possible.
		if sStr.IsEOS(pIdx) {
			if !pStr.IsEOS(pIdx) {
				for pStr.IsAsterisk(pIdx) {
					pIdx++

					if pStr.IsEOS(pIdx) {
						return true // "ab" matches "ab*".
					}
				}

				return false // "abcd" doesn't match "abc".
			}

			return true // "abc" matches "abc".
		} else if pStr.IsAsterisk(pIdx) {
			// Got wild: set up for the second loop and skip on down there.
			sIdx = pIdx

			for pStr.IsAsterisk(pIdx) {
				pIdx++
			}

			if pStr.IsEOS(pIdx) {
				return true // "abc*" matches "abcd".
			}

			// Search for the next prospective match.
			if !pStr.IsQuestionMark(pIdx) {
				for NotEqual(pStr.Char(pIdx), sStr.Char(sIdx)) {
					sIdx++
					if sStr.IsEOS(sIdx) {
						return false // "a*bc" doesn't match "ab".
					}
				}
			}

			// Keep fallback positions for retry in case of incomplete match.
			pSeq = pIdx
			sSeq = sIdx
			break
		} else if NotEqual(pStr.Char(pIdx), sStr.Char(pIdx)) && !pStr.IsQuestionMark(pIdx) {
			return false // "abc" doesn't match "abd".
		}

		pIdx++ // Everything's a match, so far.
	}

	// Find any further wildcards and any further matching sequences.
	for {
		if pStr.IsAsterisk(pIdx) {
			// Got wild again.
			for pStr.IsAsterisk(pIdx) {
				pIdx++
			}

			if pStr.IsEOS(pIdx) {
				return true // "ab*c*" matches "abcd".
			}

			if sStr.IsEOS(sIdx) {
				return false // "*bcd*" doesn't match "abc".
			}

			// Search for the next prospective match.
			if !pStr.IsQuestionMark(pIdx) {
				for NotEqual(pStr.Char(pIdx), sStr.Char(sIdx)) {
					sIdx++
					if sStr.IsEOS(sIdx) {
						return false // "a*b*c" doesn't match "ab".
					}
				}
			}

			// Keep the new fallback positions.
			pSeq = pIdx
			sSeq = sIdx
		} else if NotEqual(pStr.Char(pIdx), sStr.Char(sIdx)) && !pStr.IsQuestionMark(pIdx) {
			// The equivalent portion of the upper loop is really simple.
			if sStr.IsEOS(sIdx) {
				return false // "*bcd" doesn't match "abc".
			}

			// A fine time for questions.
			for pStr.IsQuestionMark(pSeq) {
				pSeq++
				sSeq++
			}

			pIdx = pSeq

			// Fall back, but never so far again.
			sSeq++
			for NotEqual(pStr.Char(pIdx), sStr.Char(sSeq)) {
				if sStr.IsEOS(sSeq) {
					return false // "*a*b" doesn't match "ac".
				}
				sSeq++
			}

			sIdx = sSeq
		}

		// Another check for the end, at the end.
		if sStr.IsEOS(sIdx) {
			// return true:  "*bc" matches "abc".
			// return false: "*bc" doesn't match "abcd".
			return pStr.IsEOS(pIdx)
		}

		pIdx++ // Everything's still a match.
		sIdx++
	}
}

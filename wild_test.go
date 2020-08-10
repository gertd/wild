//nolint
package wild_test

type TestCase struct {
	search          string
	pattern         string
	caseinsensitive bool
	expected        bool
}

func CaseSensitiveTestCases() []TestCase {
	return []TestCase{
		// Case with first wildcard after total match.
		{"Hi", "Hi*", false, true},

		// Case with mismatch after '*'
		{"abc", "ab*d", false, false},

		// Cases with repeating character sequences.
		{"abcccd", "*ccd", false, true},
		{"mississipissippi", "*issip*ss*", false, true},
		{"xxxx*zzzzzzzzy*f", "xxxx*zzy*fffff", false, false},
		{"xxxx*zzzzzzzzy*f", "xxx*zzy*f", false, true},
		{"xxxxzzzzzzzzyf", "xxxx*zzy*fffff", false, false},
		{"xxxxzzzzzzzzyf", "xxxx*zzy*f", false, true},
		{"xyxyxyzyxyz", "xy*z*xyz", false, true},
		{"mississippi", "*sip*", false, true},
		{"xyxyxyxyz", "xy*xyz", false, true},
		{"mississippi", "mi*sip*", false, true},
		{"ababac", "*abac*", false, true},
		{"ababac", "*abac*", false, true},
		{"aaazz", "a*zz*", false, true},
		{"a12b12", "*12*23", false, false},
		{"a12b12", "a12b", false, false},
		{"a12b12", "*12*12*", false, true},

		{"caaab", "*a?b", false, true}, // From DDJ reader Andy Belf

		// Additional cases where the '*' char appears in the tame string.
		{"*", "*", false, true},
		{"a*abab", "a*b", false, true},
		{"a*r", "a*", false, true},
		{"a*ar", "a*aar", false, false},

		// More double wildcard scenarios.
		{"XYXYXYZYXYz", "XY*Z*XYz", false, true},
		{"missisSIPpi", "*SIP*", false, true},
		{"mississipPI", "*issip*PI", false, true},
		{"xyxyxyxyz", "xy*xyz", false, true},
		{"miSsissippi", "mi*sip*", false, true},
		{"miSsissippi", "mi*Sip*", false, false},
		{"abAbac", "*Abac*", false, true},
		{"abAbac", "*Abac*", false, true},
		{"aAazz", "a*zz*", false, true},
		{"A12b12", "*12*23", false, false},
		{"a12B12", "*12*12*", false, true},
		{"oWn", "*oWn*", false, true},

		// Completely tame (no wildcards) cases.
		{"bLah", "bLah", false, true},
		{"bLah", "bLaH", false, false},

		// Simple mixed wildcard tests suggested by IBMer Marlin Deckert.
		{"a", "*?", false, true},
		{"ab", "*?", false, true},
		{"abc", "*?", false, true},

		// More mixed wildcard tests including coverage for false positives.
		{"a", "??", false, false},
		{"ab", "?*?", false, true},
		{"ab", "*?*?*", false, true},
		{"abc", "?**?*?", false, true},
		{"abc", "?**?*&?", false, false},
		{"abcd", "?b*??", false, true},
		{"abcd", "?a*??", false, false},
		{"abcd", "?**?c?", false, true},
		{"abcd", "?**?d?", false, false},
		{"abcde", "?*b*?*d*?", false, true},

		// Single-character-match cases.
		{"bLah", "bL?h", false, true},
		{"bLaaa", "bLa?", false, false},
		{"bLah", "bLa?", false, true},
		{"bLaH", "?Lah", false, false},
		{"bLaH", "?LaH", false, true},

		// Many-wildcard scenarios.
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*aa*aaa*a*a*b", false, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*a*aa*aaa*fa*ga*b*", false, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*a*x*aaa*fa*ga*b*", false, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*aaaa*fa*ga*gggg*b*", false, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*aaaa*fa*ga*ggg*b*", false, true},
		{"aaabbaabbaab", "*aabbaa*a*", false, true},
		{"a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", "a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", false, true},
		{"aaaaaaaaaaaaaaaaa", "*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", false, true},
		{"aaaaaaaaaaaaaaaa", "*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", false, false},
		{"abc*abcd*abcde*abcdef*abcdefg*abcdefgh*abcdefghi*abcdefghij*abcdefghijk*abcdefghijkl*abcdefghijklm*abcdefghijklmn", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*", false, false},
		{"abc*abcd*abcde*abcdef*abcdefg*abcdefgh*abcdefghi*abcdefghij*abcdefghijk*abcdefghijkl*abcdefghijklm*abcdefghijklmn", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*", false, true},
		{"abc*abcd*abcd*abc*abcd", "abc*abc*abc*abc*abc", false, false},
		{"abc*abcd*abcd*abc*abcd*abcd*abc*abcd*abc*abc*abcd", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abcd", false, true},
		{"abc", "********a********b********c********", false, true},
		{"********a********b********c********", "abc", false, false},
		{"abc", "********a********b********b********", false, false},
		{"*abc*", "***a*b*c***", false, true},

		// A case-insensitive algorithm test.
		{"mississippi", "*issip*PI", false, false},

		// Tests suggested by other DDJ readers
		{"", "?", false, false},
		{"", "*?", false, false},
		{"", "", false, true},
		{"a", "", false, false},

		// Case with last character mismatch.
		{"abc", "abd", false, false},

		// Cases with repeating character sequences.
		{"abcccd", "abcccd", false, true},
		{"mississipissippi", "mississipissippi", false, true},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyfffff", false, false},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyf", false, true},
		{"xxxxzzzzzzzzyf", "xxxxzzy.fffff", false, false},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyf", false, true},
		{"xyxyxyzyxyz", "xyxyxyzyxyz", false, true},
		{"mississippi", "mississippi", false, true},
		{"xyxyxyxyz", "xyxyxyxyz", false, true},
		{"m ississippi", "m ississippi", false, true},
		{"ababac", "ababac?", false, false},
		{"dababac", "ababac", false, false},
		{"aaazz", "aaazz", false, true},
		{"a12b12", "1212", false, false},
		{"a12b12", "a12b", false, false},
		{"a12b12", "a12b12", false, true},

		// A mix of cases
		{"n", "n", false, true},
		{"aabab", "aabab", false, true},
		{"ar", "ar", false, true},
		{"aar", "aaar", false, false},
		{"XYXYXYZYXYz", "XYXYXYZYXYz", false, true},
		{"missisSIPpi", "missisSIPpi", false, true},
		{"mississipPI", "mississipPI", false, true},
		{"xyxyxyxyz", "xyxyxyxyz", false, true},
		{"miSsissippi", "miSsissippi", false, true},
		{"miSsissippi", "miSsisSippi", false, false},
		{"abAbac", "abAbac", false, true},
		{"abAbac", "abAbac", false, true},
		{"aAazz", "aAazz", false, true},
		{"A12b12", "A12b123", false, false},
		{"a12B12", "a12B12", false, true},
		{"oWn", "oWn", false, true},
		{"bLah", "bLah", false, true},
		{"bLah", "bLaH", false, false},

		// Single '?' cases.
		{"a", "a", false, true},
		{"ab", "a?", false, true},
		{"abc", "ab?", false, true},

		// Mixed '?' cases.
		{"a", "??", false, false},
		{"ab", "??", false, true},
		{"abc", "???", false, true},
		{"abcd", "????", false, true},
		{"abc", "????", false, false},
		{"abcd", "?b??", false, true},
		{"abcd", "?a??", false, false},
		{"abcd", "??c?", false, true},
		{"abcd", "??d?", false, false},
		{"abcde", "?b?d*?", false, true},

		// Longer string scenarios.
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", false, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", false, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajaxalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", false, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaggggagaaaaaaaab", false, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", false, true},
		{"aaabbaabbaab", "aaabbaabbaab", false, true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false, true},
		{"aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", false, true},
		{"aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", false, false},
		{"abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc", false, false},
		{"abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", "abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", false, true},
		{"abcabcdabcdabcabcd", "abcabc?abcabcabc", false, false},
		{"abcabcdabcdabcabcdabcdabcabcdabcabcabcd", "abcabc?abc?abcabc?abc?abc?bc?abc?bc?bcd", false, true},
		{"?abc?", "?abc?", false, true},

		// A simple case
		{"", "abd", false, false},

		// Cases with repeating character sequences
		{"", "abcccd", false, false},
		{"", "mississipissippi", false, false},
		{"", "xxxxzzzzzzzzyfffff", false, false},
		{"", "xxxxzzzzzzzzyf", false, false},
		{"", "xxxxzzy.fffff", false, false},
		{"", "xxxxzzzzzzzzyf", false, false},
		{"", "xyxyxyzyxyz", false, false},
		{"", "mississippi", false, false},
		{"", "xyxyxyxyz", false, false},
		{"", "m ississippi", false, false},
		{"", "ababac*", false, false},
		{"", "ababac", false, false},
		{"", "aaazz", false, false},
		{"", "1212", false, false},
		{"", "a12b", false, false},
		{"", "a12b12", false, false},

		// A mix of cases
		{"", "n", false, false},
		{"", "aabab", false, false},
		{"", "ar", false, false},
		{"", "aaar", false, false},
		{"", "XYXYXYZYXYz", false, false},
		{"", "missisSIPpi", false, false},
		{"", "mississipPI", false, false},
		{"", "xyxyxyxyz", false, false},
		{"", "miSsissippi", false, false},
		{"", "miSsisSippi", false, false},
		{"", "abAbac", false, false},
		{"", "abAbac", false, false},
		{"", "aAazz", false, false},
		{"", "A12b123", false, false},
		{"", "a12B12", false, false},
		{"", "oWn", false, false},
		{"", "bLah", false, false},
		{"", "bLaH", false, false},

		// Both strings empty
		{"", "", false, true},

		// Another simple case
		{"abc", "", false, false},

		// Cases with repeating character sequences.
		{"abcccd", "", false, false},
		{"mississipissippi", "", false, false},
		{"xxxxzzzzzzzzyf", "", false, false},
		{"xxxxzzzzzzzzyf", "", false, false},
		{"xxxxzzzzzzzzyf", "", false, false},
		{"xxxxzzzzzzzzyf", "", false, false},
		{"xyxyxyzyxyz", "", false, false},
		{"mississippi", "", false, false},
		{"xyxyxyxyz", "", false, false},
		{"m ississippi", "", false, false},
		{"ababac", "", false, false},
		{"dababac", "", false, false},
		{"aaazz", "", false, false},
		{"a12b12", "", false, false},
		{"a12b12", "", false, false},
		{"a12b12", "", false, false},

		// A mix of cases
		{"n", "", false, false},
		{"aabab", "", false, false},
		{"ar", "", false, false},
		{"aar", "", false, false},
		{"XYXYXYZYXYz", "", false, false},
		{"missisSIPpi", "", false, false},
		{"mississipPI", "", false, false},
		{"xyxyxyxyz", "", false, false},
		{"miSsissippi", "", false, false},
		{"miSsissippi", "", false, false},
		{"abAbac", "", false, false},
		{"abAbac", "", false, false},
		{"aAazz", "", false, false},
		{"A12b12", "", false, false},
		{"a12B12", "", false, false},
		{"oWn", "", false, false},
		{"bLah", "", false, false},
		{"bLah", "", false, false},
	}
}

func CaseInsensitiveTestCases() []TestCase {
	return []TestCase{
		// Case with first wildcard after total match.
		{"Hi", "Hi*", true, true},

		// Case with mismatch after '*'
		{"abc", "ab*d", true, false},

		// Cases with repeating character sequences.
		{"abcccd", "*ccd", true, true},
		{"mississipissippi", "*issip*ss*", true, true},
		{"xxxx*zzzzzzzzy*f", "xxxx*zzy*fffff", true, false},
		{"xxxx*zzzzzzzzy*f", "xxx*zzy*f", true, true},
		{"xxxxzzzzzzzzyf", "xxxx*zzy*fffff", true, false},
		{"xxxxzzzzzzzzyf", "xxxx*zzy*f", true, true},
		{"xyxyxyzyxyz", "xy*z*xyz", true, true},
		{"mississippi", "*sip*", true, true},
		{"xyxyxyxyz", "xy*xyz", true, true},
		{"mississippi", "mi*sip*", true, true},
		{"ababac", "*abac*", true, true},
		{"ababac", "*abac*", true, true},
		{"aaazz", "a*zz*", true, true},
		{"a12b12", "*12*23", true, false},
		{"a12b12", "a12b", true, false},
		{"a12b12", "*12*12*", true, true},

		{"caaab", "*a?b", true, true}, // From DDJ reader Andy Belf

		// Additional cases where the '*' char appears in the tame string.
		{"*", "*", true, true},
		{"a*abab", "a*b", true, true},
		{"a*r", "a*", true, true},
		{"a*ar", "a*aar", true, false},

		// More double wildcard scenarios.
		{"XYXYXYZYXYz", "XY*Z*XYz", true, true},
		{"missisSIPpi", "*SIP*", true, true},
		{"mississipPI", "*issip*PI", true, true},
		{"xyxyxyxyz", "xy*xyz", true, true},
		{"miSsissippi", "mi*sip*", true, true},
		{"miSsissippi", "mi*Sip*", true, true},
		{"abAbac", "*Abac*", true, true},
		{"abAbac", "*Abac*", true, true},
		{"aAazz", "a*zz*", true, true},
		{"A12b12", "*12*23", true, false},
		{"a12B12", "*12*12*", true, true},
		{"oWn", "*oWn*", true, true},

		// Completely tame (no wildcards) cases.
		{"bLah", "bLah", true, true},
		{"bLah", "bLaH", true, true},

		// Simple mixed wildcard tests suggested by IBMer Marlin Deckert.
		{"a", "*?", true, true},
		{"ab", "*?", true, true},
		{"abc", "*?", true, true},

		// More mixed wildcard tests including coverage for false positives.
		{"a", "??", true, false},
		{"ab", "?*?", true, true},
		{"ab", "*?*?*", true, true},
		{"abc", "?**?*?", true, true},
		{"abc", "?**?*&?", true, false},
		{"abcd", "?b*??", true, true},
		{"abcd", "?a*??", true, false},
		{"abcd", "?**?c?", true, true},
		{"abcd", "?**?d?", true, false},
		{"abcde", "?*b*?*d*?", true, true},

		// Single-character-match cases.
		{"bLah", "bL?h", true, true},
		{"bLaaa", "bLa?", true, false},
		{"bLah", "bLa?", true, true},
		{"bLaH", "?Lah", true, true},
		{"bLaH", "?LaH", true, true},

		// Many-wildcard scenarios.
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*aa*aaa*a*a*b", true, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*a*aa*aaa*fa*ga*b*", true, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*a*x*aaa*fa*ga*b*", true, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*aaaa*fa*ga*gggg*b*", true, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "*a*b*ba*ca*aaaa*fa*ga*ggg*b*", true, true},
		{"aaabbaabbaab", "*aabbaa*a*", true, true},
		{"a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", "a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", true, true},
		{"aaaaaaaaaaaaaaaaa", "*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", true, true},
		{"aaaaaaaaaaaaaaaa", "*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*", true, false},
		{"abc*abcd*abcde*abcdef*abcdefg*abcdefgh*abcdefghi*abcdefghij*abcdefghijk*abcdefghijkl*abcdefghijklm*abcdefghijklmn", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*", true, false},
		{"abc*abcd*abcde*abcdef*abcdefg*abcdefgh*abcdefghi*abcdefghij*abcdefghijk*abcdefghijkl*abcdefghijklm*abcdefghijklmn", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*", true, true},
		{"abc*abcd*abcd*abc*abcd", "abc*abc*abc*abc*abc", true, false},
		{"abc*abcd*abcd*abc*abcd*abcd*abc*abcd*abc*abc*abcd", "abc*abc*abc*abc*abc*abc*abc*abc*abc*abc*abcd", true, true},
		{"abc", "********a********b********c********", true, true},
		{"********a********b********c********", "abc", true, false},
		{"abc", "********a********b********b********", true, false},
		{"*abc*", "***a*b*c***", true, true},

		// A case-insensitive algorithm test.
		{"mississippi", "*issip*PI", true, true},

		// Tests suggested by other DDJ readers
		{"", "?", true, false},
		{"", "*?", true, false},
		{"", "", true, true},
		{"a", "", true, false},

		// Case with last character mismatch.
		{"abc", "abd", true, false},

		// Cases with repeating character sequences.
		{"abcccd", "abcccd", true, true},
		{"mississipissippi", "mississipissippi", true, true},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyfffff", true, false},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyf", true, true},
		{"xxxxzzzzzzzzyf", "xxxxzzy.fffff", true, false},
		{"xxxxzzzzzzzzyf", "xxxxzzzzzzzzyf", true, true},
		{"xyxyxyzyxyz", "xyxyxyzyxyz", true, true},
		{"mississippi", "mississippi", true, true},
		{"xyxyxyxyz", "xyxyxyxyz", true, true},
		{"m ississippi", "m ississippi", true, true},
		{"ababac", "ababac?", true, false},
		{"dababac", "ababac", true, false},
		{"aaazz", "aaazz", true, true},
		{"a12b12", "1212", true, false},
		{"a12b12", "a12b", true, false},
		{"a12b12", "a12b12", true, true},

		// A mix of cases
		{"n", "n", true, true},
		{"aabab", "aabab", true, true},
		{"ar", "ar", true, true},
		{"aar", "aaar", true, false},
		{"XYXYXYZYXYz", "XYXYXYZYXYz", true, true},
		{"missisSIPpi", "missisSIPpi", true, true},
		{"mississipPI", "mississipPI", true, true},
		{"xyxyxyxyz", "xyxyxyxyz", true, true},
		{"miSsissippi", "miSsissippi", true, true},
		{"miSsissippi", "miSsisSippi", true, true},
		{"abAbac", "abAbac", true, true},
		{"abAbac", "abAbac", true, true},
		{"aAazz", "aAazz", true, true},
		{"A12b12", "A12b123", true, false},
		{"a12B12", "a12B12", true, true},
		{"oWn", "oWn", true, true},
		{"bLah", "bLah", true, true},
		{"bLah", "bLaH", true, true},

		// Single '?' cases.
		{"a", "a", true, true},
		{"ab", "a?", true, true},
		{"abc", "ab?", true, true},

		// Mixed '?' cases.
		{"a", "??", true, false},
		{"ab", "??", true, true},
		{"abc", "???", true, true},
		{"abcd", "????", true, true},
		{"abc", "????", true, false},
		{"abcd", "?b??", true, true},
		{"abcd", "?a??", true, false},
		{"abcd", "??c?", true, true},
		{"abcd", "??d?", true, false},
		{"abcde", "?b?d*?", true, true},

		// Longer string scenarios.
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", true, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", true, true},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajaxalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", true, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaggggagaaaaaaaab", true, false},
		{"abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", "abababababababababababababababababababaacacacacacacacadaeafagahaiajakalaaaaaaaaaaaaaaaaaffafagaagggagaaaaaaaab", true, true},
		{"aaabbaabbaab", "aaabbaabbaab", true, true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true, true},
		{"aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", true, true},
		{"aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", true, false},
		{"abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc", true, false},
		{"abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", "abcabcdabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijklabcdefghijklmabcdefghijklmn", true, true},
		{"abcabcdabcdabcabcd", "abcabc?abcabcabc", true, false},
		{"abcabcdabcdabcabcdabcdabcabcdabcabcabcd", "abcabc?abc?abcabc?abc?abc?bc?abc?bc?bcd", true, true},
		{"?abc?", "?abc?", true, true},

		// A simple case
		{"", "abd", true, false},

		// Cases with repeating character sequences
		{"", "abcccd", true, false},
		{"", "mississipissippi", true, false},
		{"", "xxxxzzzzzzzzyfffff", true, false},
		{"", "xxxxzzzzzzzzyf", true, false},
		{"", "xxxxzzy.fffff", true, false},
		{"", "xxxxzzzzzzzzyf", true, false},
		{"", "xyxyxyzyxyz", true, false},
		{"", "mississippi", true, false},
		{"", "xyxyxyxyz", true, false},
		{"", "m ississippi", true, false},
		{"", "ababac*", true, false},
		{"", "ababac", true, false},
		{"", "aaazz", true, false},
		{"", "1212", true, false},
		{"", "a12b", true, false},
		{"", "a12b12", true, false},

		// A mix of cases
		{"", "n", true, false},
		{"", "aabab", true, false},
		{"", "ar", true, false},
		{"", "aaar", true, false},
		{"", "XYXYXYZYXYz", true, false},
		{"", "missisSIPpi", true, false},
		{"", "mississipPI", true, false},
		{"", "xyxyxyxyz", true, false},
		{"", "miSsissippi", true, false},
		{"", "miSsisSippi", true, false},
		{"", "abAbac", true, false},
		{"", "abAbac", true, false},
		{"", "aAazz", true, false},
		{"", "A12b123", true, false},
		{"", "a12B12", true, false},
		{"", "oWn", true, false},
		{"", "bLah", true, false},
		{"", "bLaH", true, false},

		// Both strings empty
		{"", "", true, true},

		// Another simple case
		{"abc", "", true, false},

		// Cases with repeating character sequences.
		{"abcccd", "", true, false},
		{"mississipissippi", "", true, false},
		{"xxxxzzzzzzzzyf", "", true, false},
		{"xxxxzzzzzzzzyf", "", true, false},
		{"xxxxzzzzzzzzyf", "", true, false},
		{"xxxxzzzzzzzzyf", "", true, false},
		{"xyxyxyzyxyz", "", true, false},
		{"mississippi", "", true, false},
		{"xyxyxyxyz", "", true, false},
		{"m ississippi", "", true, false},
		{"ababac", "", true, false},
		{"dababac", "", true, false},
		{"aaazz", "", true, false},
		{"a12b12", "", true, false},
		{"a12b12", "", true, false},
		{"a12b12", "", true, false},

		// A mix of cases
		{"n", "", true, false},
		{"aabab", "", true, false},
		{"ar", "", true, false},
		{"aar", "", true, false},
		{"XYXYXYZYXYz", "", true, false},
		{"missisSIPpi", "", true, false},
		{"mississipPI", "", true, false},
		{"xyxyxyxyz", "", true, false},
		{"miSsissippi", "", true, false},
		{"miSsissippi", "", true, false},
		{"abAbac", "", true, false},
		{"abAbac", "", true, false},
		{"aAazz", "", true, false},
		{"A12b12", "", true, false},
		{"a12B12", "", true, false},
		{"oWn", "", true, false},
		{"bLah", "", true, false},
		{"bLah", "", true, false},
	}
}

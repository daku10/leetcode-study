package main

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	// super simple algorhythm.
	// KMP is better.
	for i := 0; i <= lenHaystack - lenNeedle; i++ {
		if haystack[i:i + lenNeedle] == needle {
			return i
		}
	}

	return -1    
}

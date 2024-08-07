package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Helper function to count character frequencies
	countFreq := func(s string) [26]int {
		count := [26]int{}
		for _, char := range s {
			count[char-'a']++
		}
		return count
	}

	// Count frequencies of s1
	s1Count := countFreq(s1)

	// Count frequencies of the first window of s2
	windowCount := countFreq(s2[:len(s1)])

	// Check if the first window matches
	if s1Count == windowCount {
		return true
	}

	// Slide the window over s2
	for i := len(s1); i < len(s2); i++ {
		// Include the next character in the window
		windowCount[s2[i]-'a']++
		// Exclude the character that is left out of the window
		windowCount[s2[i-len(s1)]-'a']--

		// Check if the current window matches
		if s1Count == windowCount {
			return true
		}
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaoo"
	fmt.Println(checkInclusion(s1, s2)) // Output: false

	s1 = "abc"
	s2 = "ccccbbbbaaaa"
	fmt.Println(checkInclusion(s1, s2)) // Output: false

	s1 = "ab"
	s2 = "eidbaooo"
	fmt.Println(checkInclusion(s1, s2)) // Output: true

	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2)) // Output: false
}

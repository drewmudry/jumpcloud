package main

import (
	"fmt"
	"strings"
)

func Solution(message string, K int) string {
	// Early exit: if message fits, return as is
	if len(message) <= K {
		return message
	}
	
	// Handle edge cases
	if K < 3 {
		return message[:K]
	}
	
	// Split into words
	words := strings.Fields(message)
	if len(words) == 0 {
		return message[:K]
	}
	
	// Pre-calculate word lengths to avoid repeated len() calls
	wordLengths := make([]int, len(words))
	for i, word := range words {
		wordLengths[i] = len(word)
	}
	
	// Try to fit words with " ..." (4 chars)
	if K >= 4 {
		maxLength := K - 4
		
		// Start with all words and remove from the end
		for i := len(words); i > 0; i-- {
			// Calculate total length for first i words
			totalLength := 0
			for j := 0; j < i; j++ {
				if j > 0 {
					totalLength++ // space
				}
				totalLength += wordLengths[j]
			}
			
			// If this fits, we found our solution
			if totalLength <= maxLength {
				result := strings.Join(words[:i], " ")
				return result + " ..."
			}
		}
	}
	
	// If we couldn't fit any words with " ...", return "..."
	if K >= 3 {
		return "..."
	}
	
	// Fallback for very small K
	return message[:K]
}

func main() {
	// Test cases
	fmt.Println(Solution("And now here is my secret", 15)) // "And now ..."
	fmt.Println(Solution("super dog", 4))                  // "..."
	fmt.Println(Solution("hello world", 20))               // "hello world"
	fmt.Println(Solution("a", 5))                         // "a"
	fmt.Println(Solution("this is a very long message", 10)) // "this ..."
	
	// Verify the length
	result := Solution("And now here is my secret", 15)
	fmt.Printf("Result: '%s' (length: %d)\n", result, len(result))
	
	result2 := Solution("super dog", 4)
	fmt.Printf("Result: '%s' (length: %d)\n", result2, len(result2))
}
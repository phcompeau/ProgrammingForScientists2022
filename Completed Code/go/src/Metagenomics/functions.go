package main

//Richness determines the richness of a frequency map representing a sample or ecosystem.
//Input: a map of strings to integers representing a sample.
//Output: the number of strings detected in this sample (i.e., its richness).
func Richness(freq map[string]int) int {
	//richness is just the number of keys in our map
	count := 0
	//range over all the values of the map and increment a count
	//whenever the value is nonzero.
	for _, val := range freq {
		if val > 0 {
			count++
		} else if val < 0 { // any negative keys?
			panic("Error: negative integer in frequency map given to Richness.")
		}
	}

	return count
}

//SimpsonsIndex computes the Simpson's index (evenness metric)
//of a sample.
//Input: A frequency map of strings to integers.
//Output: The Simpson's index of this frequency map.
func SimpsonsIndex(freq map[string]int) float64 {
	return 0.0
}

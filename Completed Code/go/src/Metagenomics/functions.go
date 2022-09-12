package main

import (
	"sort"
)

//BetaDiversityMatrix
//Input: A map of frequency maps allMaps and a string distanceMetric representing "Bray-Curtis" or "Jaccard".
//Output: A sorted list samples of the sample names, as well as a 2-D array D such that D[i][j] is the distance from samples[i] to samples[j] using the distance metric indicated by distanceMetric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {
	//first, let's sort the keys (sample IDs) of allMaps

	samples := make([]string, 0)

	//range over the maps and append each sample ID
	for sampleID := range allMaps {
		samples = append(samples, sampleID)
	}

	//let's sort these!
	sort.Strings(samples)

	//now build the distance matrix
	numSamples := len(allMaps)

	//make a distance matrix with zero values
	distanceMatrix := InitializeSquareMatrix(numSamples)

	//range over distance matrix and set all values
	for row := range distanceMatrix {
		for col := range distanceMatrix[row] {
			//want distance between two things. What are they?

			freqMap1 := allMaps[samples[row]]
			freqMap2 := allMaps[samples[col]]

			//now compute distance. Which one?
			if distMetric == "Jaccard" {
				distanceMatrix[row][col] = JaccardDistance(freqMap1, freqMap2)
			} else if distMetric == "Bray-Curtis" {
				distanceMatrix[row][col] = BrayCurtisDistance(freqMap1, freqMap2)
			} else {
				panic("no")
			}
		}
	}

	return samples, distanceMatrix

}

func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)
	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

//RichnessMap computes the richness of an arbitrary number of samples.
//Input: A map of frequency maps allMaps whose keys are strings.
//Output: A map R such that R[sample] is the richness of allMaps[sample].
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	r := make(map[string]int)

	//range through all the maps and for each one, calculate the richness
	for sampleID := range allMaps {
		//calculate the richness of current sample
		currentSample := allMaps[sampleID]
		currentRichness := Richness(currentSample)

		//assign current value of r this richness
		r[sampleID] = currentRichness
	}

	return r
}

//EvennessMap computes the Simpson's index of an arbitrary number of samples.
//Input: A map of frequency maps allMaps whose keys are strings.
//Output: A map E such that E[sample] is the Simpson’s index of allMaps[sample].
func EvennessMap(allMaps map[string](map[string]int)) map[string]float64 {
	e := make(map[string]float64)

	//range through all the maps and for each one, calculate the Simpson's index
	for sampleID := range allMaps {
		currentSample := allMaps[sampleID]
		currentEvenness := SimpsonsIndex(currentSample)

		//set the simpson's index
		e[sampleID] = currentEvenness
	}

	return e
}

//BrayCurtisDistance computes the Bray-Curtis distance between two samples.
//Input: two frequency maps (of strings to ints)
//Output: A float64 representing the Bray-Curtis distance between these maps.
func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	total1 := SumOfValues(sample1)
	total2 := SumOfValues(sample2)

	av := Average(float64(total1), float64(total2))

	sum := SumOfMinima(sample1, sample2)

	return 1.0 - float64(sum)/av
}

func Average(x, y float64) float64 {
	return (x + y) / 2.0
}

//JaccardDistance computes the Jaccard distance between two samples.
//Input: two frequency maps (of strings to ints)
//Output: A float64 representing the Jaccard distance between these maps.
func JaccardDistance(sample1, sample2 map[string]int) float64 {
	sumMin := SumOfMinima(sample1, sample2)
	sumMax := SumOfMaxima(sample1, sample2)

	return 1.0 - float64(sumMin)/float64(sumMax)
}

//SumOfMinima
//Input: Two frequency maps of strings to integers
//Output: Sum of minimum values over the two maps for each shared key in the frequency maps
func SumOfMinima(sample1, sample2 map[string]int) int {
	sum := 0

	// range through the keys of one of the maps.
	for key := range sample1 {
		//is this key present in the other map?

		_, exists := sample2[key] // exists = true if sample2[key] exists

		if exists { // yes
			//add the minimum to current sum
			sum += Min2(sample1[key], sample2[key])
		}
		//if no, take no action (or add zero)\
	}
	return sum
}

func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//SumOfMaxima
//Input: Two frequency maps of strings to integers
//Output: Sum of maximum values over two maps for each key present in either frequency map; if a key is present in one map but not the other, add its value to the sum.
func SumOfMaxima(sample1, sample2 map[string]int) int {
	sum := 0
	//range through all keys of sample 1
	for pattern := range sample1 {
		_, exists := sample2[pattern] //does this key occur in sample 2?
		//if yes, add max of the two values to sum
		if exists {
			sum += Max2(sample1[pattern], sample2[pattern])
		} else {
			//if no, add value of sample1[key] to sum
			sum += sample1[pattern]
		}
	}

	//range through the keys of sample 2.
	for pattern := range sample2 {
		//does this key occur in sample 1?
		_, exists := sample1[pattern]

		//if yes, take no action.
		if !exists {
			//if no, add value of sample2[key] to sum
			sum += sample2[pattern]
		}
	}

	return sum
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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
	simpson := 0.0

	//need to know the sum of all values in the frequency map
	total := SumOfValues(freq)

	//iterate over map, and square the probability of "choosing" the current element twice with replacement
	for _, val := range freq {
		probability := float64(val) / float64(total)
		simpson += probability * probability
	}

	return simpson
}

//SumOfValues sums all values in a frequency map.
//Input: a map of strings to integers.
//Output: the sum of all values. Panics if there is a negative value.
func SumOfValues(freq map[string]int) int {
	total := 0

	for _, val := range freq {
		if val < 0 {
			panic("Error: negative value given to SumOfValues.")
		}
		total += val
	}

	return total
}

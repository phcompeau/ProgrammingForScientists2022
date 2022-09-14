package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRichness(t *testing.T) {
	//first, declare our test type
	type test struct {
		frequencyMap map[string]int
		answer       int
	}

	inputDirectory := "tests/Richness/input/"
	outputDirectory := "tests/Richness/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].frequencyMap = ReadFrequencyMapFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadIntegerFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := Richness(test.frequencyMap)

		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %d, and the correct richness is %d", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the frequency map is", test.frequencyMap, "the richness is", test.answer)
		}
	}

}

func TestSimpsonsIndex(t *testing.T) {
	//first, declare our test type
	type test struct {
		frequencyMap map[string]int
		answer       float64
	}

	inputDirectory := "tests/SimpsonsIndex/input/"
	outputDirectory := "tests/SimpsonsIndex/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].frequencyMap = ReadFrequencyMapFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadFloatFromFile(outputDirectory, outputFiles[i])
	}

	//are the tests correct?
	for i, test := range tests {
		outcome := SimpsonsIndex(test.frequencyMap)
		var numDigits uint = 4

		if roundFloat(outcome, numDigits) != roundFloat(test.answer, numDigits) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the frequency map is", test.frequencyMap, "the Simpson's index is", test.answer)
		}
	}

}

func TestBrayCurtis(t *testing.T) {
	//first, declare our test type
	type test struct {
		frequencyMap1, frequencyMap2 map[string]int
		answer                       float64
	}

	inputDirectory := "tests/BrayCurtis/input/"
	outputDirectory := "tests/BrayCurtis/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].frequencyMap1, tests[i].frequencyMap2 = ReadTwoFrequencyMapsFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadFloatFromFile(outputDirectory, outputFiles[i])
	}

	//are the tests correct?
	for i, test := range tests {
		outcome := BrayCurtisDistance(test.frequencyMap1, test.frequencyMap2)

		var numDigits uint = 4

		if roundFloat(outcome, numDigits) != roundFloat(test.answer, numDigits) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the frequency maps are", test.frequencyMap1, "and", test.frequencyMap2, "the Bray-Curtis distance is", test.answer)
		}
	}

}

func TestJaccard(t *testing.T) {
	//first, declare our test type
	type test struct {
		frequencyMap1, frequencyMap2 map[string]int
		answer                       float64
	}

	inputDirectory := "tests/Jaccard/input/"
	outputDirectory := "tests/Jaccard/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].frequencyMap1, tests[i].frequencyMap2 = ReadTwoFrequencyMapsFromFile(inputDirectory, inputFiles[i])
		tests[i].answer = ReadFloatFromFile(outputDirectory, outputFiles[i])
	}

	//are the tests correct?
	for i, test := range tests {
		outcome := JaccardDistance(test.frequencyMap1, test.frequencyMap2)

		var numDigits uint = 4

		if roundFloat(outcome, numDigits) != roundFloat(test.answer, numDigits) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the frequency maps are", test.frequencyMap1, "and", test.frequencyMap2, "the Jaccard distance is", test.answer)
		}
	}

}

func ReadFloatFromFile(directory string, file os.FileInfo) float64 {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.ParseFloat(outputLines[0], 64)

	if err != nil {
		panic(err)
	}

	return answer
}

func ReadIntegerFromFile(directory string, file os.FileInfo) int {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.Atoi(outputLines[0])

	if err != nil {
		panic(err)
	}

	return answer
}

func ReadFrequencyMapFromFile(directory string, inputFile os.FileInfo) map[string]int {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//make the map that will store our frequency map
	frequencyMap := make(map[string]int)

	//each line of the file corresponds to a single line of the frequency map
	for _, inputLine := range inputLines {

		//read out the current line
		currentLine := strings.Split(inputLine, " ")
		//currentLine has two strings corresponding to the key and value
		currentString := currentLine[0]
		currentVal, err := strconv.Atoi(currentLine[1])
		if err != nil {
			panic(err)
		}

		//if we make it here, everything is OK, so append to the input map
		frequencyMap[currentString] = currentVal
	}
	return frequencyMap
}

func ReadTwoFrequencyMapsFromFile(directory string, inputFile os.FileInfo) (map[string]int, map[string]int) {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//make the maps that will store our frequency maps
	frequencyMap1 := make(map[string]int)
	frequencyMap2 := make(map[string]int)

	mapIndex := 1

	//each line of the file corresponds to a single line of the frequency map
	for _, inputLine := range inputLines {
		if inputLine == "-" {
			mapIndex = 2
			continue
		}

		//read out the current line
		currentLine := strings.Split(inputLine, " ")
		//currentLine has two strings corresponding to the key and value
		currentString := currentLine[0]
		currentVal, err := strconv.Atoi(currentLine[1])

		if err != nil {
			panic(err)
		}

		//if we make it here, everything is OK, so append to the appropriate map
		if mapIndex == 1 {
			frequencyMap1[currentString] = currentVal
		} else {
			frequencyMap2[currentString] = currentVal
		}
	}
	return frequencyMap1, frequencyMap2
}

func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in input directory.")
	}
	if length1 == 0 {
		panic("No files present in output directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

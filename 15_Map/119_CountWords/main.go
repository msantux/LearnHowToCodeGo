package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Open a file
	f, err := os.Open("book.txt")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// The frequency of words in the file
	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %v\n", err)
	}

	// Sort the word frequencies
	pairs := sortWordFrequency(words)

	// Print the sorted results
	for _, v := range pairs {
		fmt.Printf("%-30v %v\n", v.Key, v.Value)
	}

	// Word with greatest frequency, and it's frequency
	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord: %v\n", err)
	}
	fmt.Printf("%v has a frequency %v", w, n)
}

func freq(r io.Reader) (map[string]int, error) {
	// create a map to store word frequencies
	wordFreq := map[string]int{}

	// create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	// read each word and update frequencies
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}

// For sorting frequency of words
type Pair struct {
	Key   string
	Value int
}

// implement the Len, Less and Swap methods in PairList
// to satisfy the sort.Interface interface.
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value //Sorting in descent order
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	// Convert the map to a PairList
	pairs := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	// Sort the PairsList
	sort.Sort(pairs)

	return pairs
}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}

	maxW := "" // word with max frequency
	maxF := 0  // frequency of that word

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}

	return maxW, maxF, nil
}

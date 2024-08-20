package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Definition struct {
	Word        string `json:"word"`
	Form        string `json:"type"`
	Description string `json:"description"`
}

var (
	dictionary []Definition
)

func unmarshal() error {
	contents, err := os.ReadFile("data/EDMTDictionary.json")
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	if err := json.Unmarshal(contents, &dictionary); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func SetupDictionary() {
	if err := unmarshal(); err != nil {
		fmt.Println("Error setting up dictionary:", err)
	}
}

func preprocess(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	return word
}

func BinarySearch(search string) [][]string {
	search = preprocess(search)
	searchWord := wordToA(search)
	left := 0
	right := len(dictionary) - 1
	for left <= right {
		middle := (left + right) / 2
		current := preprocess(dictionary[middle].Word)
		currentWord := wordToA(current)
		if search == current {
			return checkMultiple(middle)
		}
		for i := 0; i < len(searchWord); i++ {
			if i == len(currentWord) { // e.g. search: management, current: manage
				left = middle + 1
				break
			}
			if currentWord[i] < searchWord[i] {
				left = middle + 1
				break
			}
			if currentWord[i] > searchWord[i] {
				right = middle - 1
				break
			}
			if i == len(searchWord)-1 && i != len(currentWord)-1 { // search: manage, current: management
				right = middle - 1
			}
		}
	}
	return [][]string{}
}

const windowSize = 25

func checkMultiple(loc int) [][]string {
	var wordList [][]string
	before := loc - windowSize
	after := loc + windowSize
	if before < 0 {
		before = 0
	}
	if after > len(dictionary)-1 {
		after = len(dictionary) - 1
	}
	word := preprocess(dictionary[loc].Word)
	for i := before; i <= after; i++ {
		if preprocess(dictionary[i].Word) == word {
			wordList = append(wordList, []string{dictionary[i].Word, dictionary[i].Form, dictionary[i].Description})
		}
	}
	return wordList
}

func wordToA(word string) []string {
	var wordArray []string
	for _, i := range word {
		wordArray = append(wordArray, string(i))
	}
	return wordArray
}

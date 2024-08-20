package services

import (
	"encoding/json"
	"fmt"
	"math"
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
		panic(err)
	}
	err = json.Unmarshal(contents, &dictionary)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return err
}

func SetupDictionary() {
	unmarshal()
}

func preprocess(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	return word
}

func BinarySearch(search string) []string {
	search = preprocess(search)
	left := 0
	right := len(dictionary) - 1
	for left <= right {
		middle := int(math.Floor(float64(left+right) / 2))
		searchWord := wordToA(search)
		current := preprocess(dictionary[middle].Word)
		currentWord := wordToA(current)
		if search == current {
			result := []string{dictionary[middle].Word, dictionary[middle].Form, dictionary[middle].Description}
			return result
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
	return []string{}
}

func wordToA(word string) []string {
	var wordArray []string
	for _, i := range word {
		wordArray = append(wordArray, string(i))
	}
	return wordArray
}

//{
//    "word": "Abandon",
//    "type": "(v. t.)",
//    "description": "To cast or drive out; to banish; to expel; to reject."
//  },

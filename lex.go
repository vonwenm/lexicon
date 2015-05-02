package lex

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// DefaultDataPath gets the lexicon word list file location from the LEXICON_DATA environment variable.
func DefaultDataPath() string {
	dir := os.Getenv("LEXICON_DATA")
	if dir == "" {
		log.Fatal("Set LEXICON_DATA variable to directory of lexicon data files")
	}
	return dir
}

// CountChars counts all the characters in a language lexicon.
func CountChars(words map[string]bool) map[rune]int {
	chars := make(map[rune]int)
	for word := range words {
		for _, ch := range word {
			chars[ch]++
		}
	}
	return chars
}

// CountSubstr counts all the substrings of given length in a language lexicon.
func CountSubstr(words map[string]bool, length int) map[string]int {
	substrs := make(map[string]int)
	for word := range words {
		list := ListSubstr(word, length)
		for _, substr := range list {
			substrs[substr]++
		}
	}
	return substrs
}

// ListSubstr lists all the substrings of a word of given length.
func ListSubstr(word string, length int) []string {
	// Cast word as a rune slice.
	runes := []rune(word)
	n := len(runes)
	subcount := n - length + 1
	if subcount < 0 {
		subcount = 0
	}
	substrs := make([]string, subcount)
	if n >= length {
		max := n - length
		for i := 0; i <= max; i++ {
			// Cast portion of rune slice back to string.
			substr := string(runes[i : i+length])
			substrs = append(substrs, substr)
		}
	}
	return substrs
}

// LoadAllLangs loads all the language files.
func LoadAllLangs(langDir string) map[string]map[string]bool {
	// Load the languages.
	langFiles, err := ioutil.ReadDir(langDir)
	if err != nil {
		log.Fatal(err)
	}
	langWords := make(map[string]map[string]bool)
	for _, langFile := range langFiles {
		name := langFile.Name()
		path := path.Join(langDir, name)
		langWords[name] = LoadLang(path)
	}
	return langWords
}

// LoadLang loads a language file.
func LoadLang(langFile string) map[string]bool {
	// Open file.
	handle, err := os.Open(langFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Scan file line by line.
	words := make(map[string]bool)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		word := strings.TrimSpace(strings.ToLower(line))
		words[word] = true
	}
	return words
}
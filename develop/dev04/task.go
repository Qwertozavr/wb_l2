package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagramm(input []string) map[string][]string {
	if len(input) < 2 {
		return nil
	}

	result := make(map[string][]string)

	for i := 0; i < len(input)-1; i++ {
		input[i] = strings.ToLower(input[i])
		if _, ok := result[input[i]]; !ok {
			result[input[i]] = append(result[input[i]], input[i])
			for j := i + 1; j < len(input); j++ {
				if len(input[i]) == len(input[j]) {
					var equal bool
					for _, rune_i := range input[i] {
						equal = false
						for _, rune_j := range input[j] {
							if rune_i == rune_j {
								equal = true
								break
							}
						}
						if !equal {
							break
						}
					}
					if equal {
						result[input[i]] = append(result[input[i]], input[j])
						result[input[j]] = []string{}
					}
				}
			}
		}
	}
	for key, anagram := range result {
		if len(anagram) < 2 {
			delete(result, key)
			continue
		}
		sort.Slice(anagram, func(i, j int) bool { return anagram[i] < anagram[j] })
	}

	return result
}

func main() {
	fmt.Println(Anagramm([]string{"слиток", "пятак", "столик", "пятка", "тяпка", "листок", "арбуз"}))
}

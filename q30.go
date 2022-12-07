package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordMap := make(map[string][]int)
	arr := [][]int{}
	answer := []int{}
	for i, w := range words {
		_, exists := wordMap[w]
		if !exists {
			wordMap[w] = make([]int, 0)
		}
		wordMap[w] = append(wordMap[w], i+1)
	}
	wordLen := len(words[0])
	numWords := len(words)

	for i := 0; i < len(s)-wordLen+1; i++ {
		currWord := s[i : i+wordLen]
		val, exist := wordMap[currWord]
		if exist {
			arr = append(arr, val)
		} else {
			arr = append(arr, nil)
		}
	}
	//filling in the last values
	for i := 1; i < wordLen; i++ {
		arr = append(arr, nil)
	}
	lenCompleteString := wordLen * numWords
	// fmt.Println(lenCompleteString)
	// fmt.Println(arr)
	for i := 0; i < len(arr)-lenCompleteString+1; i++ {
		indexMap := make(map[int]bool)
		for j := i; j < i+lenCompleteString; {
			//required to ensure that the matching word is still within our bounds
			stillInBound := j+wordLen <= i+lenCompleteString
			foundWords := arr[j] //this is a list
			if len(foundWords) > 0 && stillInBound {
				//now we need to iterate and add these words
				for _, k := range foundWords {
					_, exists := indexMap[k]
					if !exists {
						indexMap[k] = true
						break
					}
				}
				j += wordLen
			} else {
				j += 1
			}
		}
		//if all are found (this is a hack but with unique words should work)
		if len(indexMap) == numWords {
			answer = append(answer, i)
		}
	}
	return answer
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

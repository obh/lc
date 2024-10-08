package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordMap := make(map[string][]int)
	// arr := [][]int{}
	arr2 := []string{}
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
	fmt.Println(len(s) - wordLen + 1)
	for i := 0; i < len(s)-wordLen+1; i++ {
		currWord := s[i : i+wordLen]
		_, exist := wordMap[currWord]
		if exist {
			// arr = append(arr, val)
			arr2 = append(arr2, currWord)
		} else {
			// arr = append(arr, nil)
			arr2 = append(arr2, "")
		}
	}
	//filling in the last values
	for i := 1; i < wordLen; i++ {
		// arr = append(arr, nil)
		arr2 = append(arr2, "")
	}
	lenCompleteString := wordLen * numWords
	fmt.Println(lenCompleteString)
	// fmt.Println(len(arr))
	for i := 0; i < len(arr2)-lenCompleteString+1; i++ {
		indexMap := make(map[string]int)

		for j := i; j < i+lenCompleteString; {
			//required to ensure that the matching word is still within our bounds
			stillInBound := j+wordLen <= i+lenCompleteString
			foundWords := wordMap[arr2[j]] //this is a list
			if len(foundWords) > 0 && stillInBound {
				//arr2[j] is the word
				_, exists := indexMap[arr2[j]]
				if !exists {
					indexMap[arr2[j]] = 0
				}
				// we can only have upto max words
				indexMap[arr2[j]] = min(indexMap[arr2[j]]+1, len(foundWords))
				//now we need to iterate and add these words
				// fmt.Println("found words --> ", foundWords)
				// for _, k := range foundWords {
				// 	_, exists := indexMap[k]
				// 	if !exists {
				// 		indexMap[k] = true
				// 		break
				// 	}
				// }
				j += wordLen
			} else {
				j += 1
			}
		}

		//if all are found (this is a hack but with unique words should work)
		// fmt.Println(i, indexMap)
		count := 0
		for _, v := range indexMap {
			count += v
		}
		if count == numWords {
			answer = append(answer, i)
		}
	}
	return answer
}

func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	// fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}

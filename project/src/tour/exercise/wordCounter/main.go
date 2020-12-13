package main

import (
	"strings"
	"sort"
	"fmt"
)

func main() {
	s := "By the time Doug Raysby’s wife was allowed to enter his hospital room, " +
		"it was too late to be sure whether he even knew she was there. " +
		"After a feverish fight with the coronavirus, he lay unconscious on the bed. " +
		"His wife cried through an N95 mask, while a computer tablet flashed a video stream of his children saying goodbye." +
		"But now, signs are shifting: More than 1,000 Americans are dying of the coronavirus every day on average, a 50 percent increase in the last month. " +
		"Iowa, Minnesota, New Mexico, Tennessee and Wisconsin have recorded more deaths over the last seven days than in any other week of the pandemic. " +
		"Twice this past week, there have been more than 1,400 deaths reported in a single day.It’s getting bad and it’s potentially going to get a lot worse,” " +
		"said Jennifer Nuzzo, an epidemiologist and senior scholar at the Johns Hopkins Center for Health Security. “The months ahead are looking quite horrifying.”" +
		"For families like Mr. Raysby’s, the pain of personal loss has combined with a sense of anger that the nation, exhausted after nine months of the pandemic, " +
		"has grown inured to the death toll, even as its pace is now quickening once more." +
		"“Do you see this human being? Do you realize?” said Kathy James, the mother-in-law of Mr. Raysby, " +
		"a 57-year-old factory supervisor in Sioux Falls, S.D., who liked to hunt pheasants on the weekends. " +
		"In the weeks since Mr. Raysby died of the virus, Ms. James said she had wanted to wave a photo of him — a quiet, " +
		"bespectacled man who once wooed her daughter with purple roses — at the world."

	wordsCount := WordCount(s)

	// sort ref.
	// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range wordsCount {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for _, kv := range ss {
		fmt.Printf("%s, %d \n", kv.Key, kv.Value)
	}
}

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, value := range strings.Fields(s) {
		elem, ok := wordMap[value]
		if ok {
			wordMap[value] = elem + 1
		} else {
			wordMap[value] = 1
		}
	}
	return wordMap
}

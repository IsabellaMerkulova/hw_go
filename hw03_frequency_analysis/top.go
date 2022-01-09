package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordStat struct {
	Word string
	N    int
}

type wordsStats []WordStat

func sortWordsStats(wordsQty map[string]int) wordsStats {
	sortedStats := make(wordsStats, len(wordsQty))
	i := 0
	for word, qty := range wordsQty {
		sortedStats[i] = WordStat{Word: word, N: qty}
		i++
	}
	sort.Slice(sortedStats, func(i, j int) bool {
		if sortedStats[i].N == sortedStats[j].N {
			return sortedStats[i].Word < sortedStats[j].Word
		}
		return sortedStats[i].N > sortedStats[j].N
	})
	return sortedStats
}

func Top10(input string) []string {
	words := strings.Fields(input)
	wordsQtyMap := make(map[string]int)

	for _, word := range words {
		wordsQtyMap[word]++
	}

	result := make([]string, 0, 10)
	for i, stat := range sortWordsStats(wordsQtyMap) {
		if i == 10 {
			break
		}
		result = append(result, stat.Word)
	}
	return result
}

package ranking

import (
	"math"
	"sort"

	metrics "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/metrics"
)

type Priority struct {
	Package  string
	Name     string
	Rank     int
	Priority int
}

func Prioritize(rankabels []metrics.Rankable, top uint) []Priority {
	sort.Slice(rankabels, func(i, j int) bool {
		return rankabels[i].GetScore() > rankabels[j].GetScore()
	})

	priorities := make([]Priority, min(int(top), len(rankabels)))

	if len(priorities) == 0 {
		return priorities
	}

	currentScore := rankabels[0].GetScore()
	currentRank := 1
	for i := 0; i < len(priorities); i++ {
		score := rankabels[i].GetScore()
		if currentScore != score {
			currentScore = score
			currentRank++
		}

		priorities[i] = Priority{
			Package: rankabels[i].GetPackage(),
			Name:    rankabels[i].GetName(),
			Rank:    currentRank,
		}
	}

	calculatePriority(priorities, len(rankabels))

	return priorities
}

func calculatePriority(priorities []Priority, totalCount int) {
	countWithRankMoreThanCurrent := 0
	currentRank := 0

	for i, j := 0, 0; i < len(priorities); i++ {
		for ; j < len(priorities) && priorities[j].Rank <= currentRank; j++ {
			countWithRankMoreThanCurrent++
		}

		if j < len(priorities) {
			currentRank = priorities[j].Rank
		}

		priorities[i].Priority = int(math.Round(100 - float64(countWithRankMoreThanCurrent*100.0)/float64(totalCount)))
	}
}

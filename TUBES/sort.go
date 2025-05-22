package main

func SelectionSortByWins(ascending bool) {
	for i := 0; i < TeamCount-1; i++ {
		idx := i
		for j := i + 1; j < TeamCount; j++ {
			if (ascending && Teams[j].Wins < Teams[idx].Wins) || (!ascending && Teams[j].Wins > Teams[idx].Wins) {
				idx = j
			}
		}
		Teams[i], Teams[idx] = Teams[idx], Teams[i]
	}
}

func InsertionSortByScoreDiff(ascending bool) {
	for i := 1; i < TeamCount; i++ {
		key := Teams[i]
		j := i - 1
		for j >= 0 && ((ascending && Teams[j].ScoreDiff > key.ScoreDiff) || (!ascending && Teams[j].ScoreDiff < key.ScoreDiff)) {
			Teams[j+1] = Teams[j]
			j--
		}
		Teams[j+1] = key
	}
}

func BubbleSortByPoints(ascending bool) {
	for i := 0; i < TeamCount-1; i++ {
		for j := 0; j < TeamCount-i-1; j++ {
			if (ascending && Teams[j].Points > Teams[j+1].Points) || (!ascending && Teams[j].Points < Teams[j+1].Points) {
				Teams[j], Teams[j+1] = Teams[j+1], Teams[j]
			}
		}
	}
}

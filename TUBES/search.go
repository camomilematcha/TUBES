package main

import "strings"

func SequentialSearchByName(name string) int {
	for i := 0; i < TeamCount; i++ {
		if strings.ToLower(Teams[i].Name) == strings.ToLower(name) {
			return i
		}
	}
	return -1
}

func BinarySearchByID(target int) int {
	low, high := 0, TeamCount-1
	for low <= high {
		mid := (low + high) / 2
		if Teams[mid].ID == target {
			return mid
		} else if Teams[mid].ID < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

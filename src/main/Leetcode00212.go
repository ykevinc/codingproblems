package main

import (
	"sort"
)

func findWords(board [][]byte, words []string) []string {
	s := make([]string, 0, len(words))
	mapWord := make(map[string]bool, len(words))
	for _, word := range words {
		mapWord[word] = true
	}
	for word := range mapWord {
		found := false
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if findWord(board, i, j, word) {
					s = append(s, word)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	sort.Strings(s)
	return s
}

func findWord(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	s := board[i][j]
	board[i][j] = '0'
	found := false
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		if findWord(board, i+dir[0], j+dir[1], word[1:]) {
			found = true
			break
		}
	}
	board[i][j] = s
	return found
}

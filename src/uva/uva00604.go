/**
 * Created with IntelliJ IDEA.
 * User: Kevin
 * Date: 17/02/13
 * Time: 10:21 AM
 */
package uva00604

import (
	_ "fmt"
	"fmt"
	"os"
	"bufio"
	"sort"
	_ "text/scanner"
	_ "strings"
	_ "strconv"
)

type Direction struct {
	x, y int
}

type Board struct{
	m [][]byte
	b [][]bool
	w int
	h int
}

func NewBoard(rank int) *Board {
	board := Board{make([][]byte, rank), make([][]bool, rank), rank, rank}
	for i := range board.m {
		board.m[i] = make([]byte, rank)
		board.b[i] = make([]bool, rank)
	}
	return &board
}

func isPigEwu(buffer []byte) bool {
	vowelsCount := 0;
	for _, value := range buffer {
		switch value {
			case 'A','E','I','O','U','Y':
			vowelsCount++;
		}
	}
	return vowelsCount == 2
}

func isValid(x, y, w, h int) bool {
	if x < 0 || x >= w || y < 0 || y >= h {return false}
	return true
}

func fillBoggles(board *Board, directions []Direction, requiredPigEwuSize int, dictionary map[string] bool, x, y int, buffer []byte, ifExistThenMark bool) {
	if !isValid(x, y, board.w, board.h) || board.b[x][y] == true {return}
	board.b[x][y] = true
	buffer = append(buffer, board.m[x][y])
	if len(buffer) == requiredPigEwuSize {
		if isPigEwu(buffer) {
			_, exist := dictionary[string(buffer[:])]
			//  Set value to false to indicate the word is not yet common in both boards
			if !ifExistThenMark {
				dictionary[string(buffer[:])] = false
			} else if exist {
				dictionary[string(buffer[:])] = true
			}
		}
	}else {
		// Walk the boggle
		for i := range directions {
			fillBoggles(board, directions, requiredPigEwuSize, dictionary, directions[i].x + x, directions[i].y + y, buffer, ifExistThenMark)
		}
	}
	buffer = buffer[0:len(buffer) - 1]
	board.b[x][y] = false
}

func boggle(in *bufio.Reader, boardA, boardB *Board, rank int, directions []Direction) bool {
	for i := 0; i < rank; i++ {
		for j := 0; j < rank; j++ {
			fmt.Fscanf(in, "%c ", &boardA.m[i][j]);
		}
		for j := 0; j < rank; j++ {
			fmt.Fscanf(in, "%c ", &boardB.m[i][j]);
		}
	}
	if boardA.m[0][0] == '#' {return false}
	//for i := 0; i < rank;i++ {fmt.Println(string(boardA.m[i][:]))}
	//for i := 0; i < rank;i++ {fmt.Println(string(boardB.m[i][:]))}

	// Build the dictionary to hold  PigEwu words
	dictionary := make(map[string]bool)

	// Find all PigEwu words on board A and fill into dictionary
	for x := 0 ; x < boardA.w ; x++ {
		for y := 0 ; y < boardA.h ; y++ {
			fillBoggles(boardA, directions, rank, dictionary, x, y, make([]byte,0, rank), false)
		}
	}

	// Find all PigEwu words on board B and check if exists in dictionary; possible enhancement: pre-check if the letter in board B start with dictionary's first letters
	for x := 0 ; x < boardB.w ; x++ {
		for y := 0 ; y < boardB.h ; y++ {
			fillBoggles(boardB, directions, rank, dictionary, x, y, make([]byte,0, rank), true)
		}
	}

	// Pick the common words in both boards and output an alphabetically-sorted list
	words := make([]string,0, len(dictionary))
	for key, value := range dictionary {
		if value {words = append(words, key)}
	}
	sort.Strings(words)
	if len(words) == 0 {
		fmt.Println("There are no common words for this pair of boggle boards.")
	} else {
		for _, word := range words {
			fmt.Println(word)
		}
	}
	fmt.Println();
	return true
}

func Main() {
	in := bufio.NewReader(os.Stdin)
	rank := 4
	directions := make([]Direction,0, 8);
	for x := -1 ; x < 2 ; x++ {
		for y := -1 ; y < 2 ; y++ {
			if x == 0 && y == 0 {continue}
			directions = append(directions, Direction{x, y})
		}
	}
	boardA := NewBoard(rank)
	boardB := NewBoard(rank)
	for boggle(in, boardA, boardB, rank, directions) {}
}

/*

D F F B    W A S U
T U G I    B R E T
O K J M    Y A P Q
K M B E    L O Y R

Z W A V    G S F U
U N C O    A H F T
Y T G I    G N A L
H G P M    B O O B

#

There are no common words for this pair of boggle boards.

ANGO
AOGN
GNAO
GOAN
NAOG
NGOA
OANG
OGNA

*/

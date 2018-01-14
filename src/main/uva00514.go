/**
 * Created with IntelliJ IDEA.
 * User: kevinchen
 * Date: 16/02/13
 * Time: 1:26 PM
 */
package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
	"text/scanner"
)

func reverseEquals(normal []int, reverse []int, startOffset, reverseOffset int, size int) bool {
	if reverseOffset >= len(reverse) {
		return false
	}
	i := 0
	for startOffset < len(normal) && reverseOffset >= 0 && i < size {
		if normal[startOffset] != reverse[reverseOffset] {
			return false
		}
		startOffset++
		reverseOffset--
		i++
	}
	return i == size
}

func rails(in *bufio.Reader, trainA []int, trainB []int) int {
	line, _ := in.ReadString('\n')
	var s scanner.Scanner
	s.Init(strings.NewReader(line)).Scan()
	coachNum, _ := strconv.Atoi(s.TokenText())

	if coachNum == 0 {
		return 0
	}

	// Initialize train A with ordered coaches
	trainA = trainA[:0]
	for i := 0; i < coachNum; i++ {
		trainA = append(trainA, i+1)
	}

	// Read train B blocks
	for ; ; trainB = trainB[:0] {
		line, _ := in.ReadString('\n')
		s.Init(strings.NewReader(line))
		tok := s.Scan()
		for tok != scanner.EOF {
			coach, _ := strconv.Atoi(s.TokenText())
			trainB = append(trainB, coach)
			tok = s.Scan()
		}
		if len(trainB) == 1 && trainB[0] == 0 || len(trainB) <= 0 {
			fmt.Println()
			return coachNum
		}

		// Search a identical coach for disconnecting train A, and reverse train A to see if order is possible
		possible := true
		for i := 0; i < len(trainB); i++ {
			matching := false
			for j := i; j < len(trainA); j++ {
				if trainB[i] == trainA[j] {
					matching = true
					if !reverseEquals(trainB, trainA, i, j, j-i) {
						possible = false
						break
					}
					i = j
				}
			}
			if !possible || !matching {
				possible = false
				break
			}
		}
		if possible {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
	return 0
}

/*

5
1 2 3 4 5
5 4 1 2 3
0
6
6 5 4 3 2 1
0
0

Yes
No

Yes


5
1 4 3 2 5
0
0

YES

*/

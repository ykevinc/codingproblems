/**
 * Created with IntelliJ IDEA.
 * User: kevinchen
 * Date: 16/02/13
 * Time: 1:26 PM
 */
package uva00514

import (
	_ "fmt"
	"fmt"
	"os"
	"bufio"
	"text/scanner"
	"strings"
	"strconv"
)

func reverseEquals(normal []int, reverse []int, startOffset, reverseOffset int, size int) bool {
	if reverseOffset >= len(reverse) { return false }
	i := 0
	for ; startOffset < len(normal) && reverseOffset >= 0 && i < size; {
		if normal[startOffset] != reverse[reverseOffset] { return false; }
		startOffset++;
		reverseOffset--;
		i++;
	}
	return i == size;
}

func rails(in *bufio.Reader, trainA []int, trainB []int) int {

	line, _ := in.ReadString('\n');
	var s scanner.Scanner
	s.Init(strings.NewReader(line)).Scan()
	coachNum, _ := strconv.Atoi(s.TokenText());

	if coachNum == 0 {return 0}

	// Initialize a order train from A
	trainA = trainA[:0]
	for i := 0; i < coachNum; i++ {
		trainA = append(trainA, i + 1)
	}

	// Read train B blocks
	for ; ; trainB = trainB[:0] {
		line, _ := in.ReadString('\n')
		s.Init(strings.NewReader(line))
		tok := s.Scan()
		for tok != scanner.EOF {
			coach, _ := strconv.Atoi(s.TokenText());
			trainB = append(trainB, coach)
			tok = s.Scan()
		}
		if len(trainB) == 1 && trainB[0] == 0 || len(trainB) <= 0 {
			fmt.Println();
			return coachNum;
		}

		// Search a identical coach as the disconnect point on train B, and reverse train B to see if order is possible
		possible := true
		for i := 0; i < len(trainA); i++ {
			matching := false;
			for j := i; j < len(trainB); j++ {
				if trainA[i] == trainB[j] {
					matching = true;
					if !reverseEquals(trainA, trainB, i, j, j - i) { possible = false; break}
					i = j + 1;
				}
			}
			if !possible || !matching { possible = false; break}
		}
		if possible { fmt.Println("Yes"); } else { fmt.Println("No"); }
	}
	return 0
}

func Main() {
	in := bufio.NewReader(os.Stdin)
	trainA := make([]int,0, 1000)
	trainB := make([]int,0, 1000)
	for rails(in, trainA, trainB) > 0 {}
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

5
1 4 3 2 5
0
0

*/

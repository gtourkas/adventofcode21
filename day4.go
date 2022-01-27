package adventofcode21

import (
	"strconv"
	"strings"
)

func calcBingoScore(lines []string, winnerIsTheLast bool) int {

	// parse the drawn numbers
	drawnNumbersStr := strings.Split(lines[0],",")
	drawnNumbers := make([]int, len(drawnNumbersStr))
	for i:=0; i<len(drawnNumbersStr); i++ {
		n := drawnNumbersStr[i]
		drawnNumbers[i], _ = strconv.Atoi(n)
	}

	// parse the player boards
	type boardItem struct {
		row int
		col int
		drawn bool
	}
	var boards [][]*boardItem

	player := 0
	row := 0
	boards = append(boards, make([]*boardItem, 100))
	for i:=2; i<len(lines); i++ {
		l := lines[i]
		if l == "" {
			player += 1
			row = 0
			boards = append(boards, make([]*boardItem, 100))
			continue
		}
		numbersStr := strings.Split(strings.Replace(strings.TrimSpace(l), "  ", " ", -1), " ")
		for col:=0; col<len(numbersStr); col++ {
			number, _ := strconv.Atoi(numbersStr[col])
			boards[player][number] = &boardItem{
				row: row,
				col: col,
			}
		}
		row += 1
	}

	players := player+1

	// go through the drawn numbers counting the marked numbers per row and col for each player
	// the first row or col that reaches 5 marked numbers wins
	markedRows := make([][]int, players)
	markedCols := make([][]int, players)
	for player:=0; player<players; player++ {
		markedRows[player] = make([]int, 5)
		markedCols[player] = make([]int, 5)
	}

	winners := make([]bool, players)
	lastWinnerScore := 0
	for i:=0; i<len(drawnNumbers); i++ {
		number := drawnNumbers[i]
		for player:=0; player<players; player++ {
			boardItem := boards[player][number]
			if boardItem == nil {
				continue
			}
			boardItem.drawn = true
			markedRows[player][boardItem.row] += 1
			markedCols[player][boardItem.col] += 1
			winner := false
			if markedRows[player][boardItem.row] == 5 {
				winner = true
			}
			if markedCols[player][boardItem.col] == 5 {
				winner = true
			}
			if winner && !winners[player] {
				sum := 0
				for k:=0; k<len(boards[player]); k++ {
					boardItem = boards[player][k]
					if boardItem != nil && !boardItem.drawn {
						sum += k
					}
				}
				winnerScore := sum * number
				if !winnerIsTheLast {
					return winnerScore
				}
				winners[player] = true
				lastWinnerScore = winnerScore
			}
		}
	}
	return lastWinnerScore
}

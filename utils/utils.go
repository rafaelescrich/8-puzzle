package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CopySlice copies a slice
func CopySlice(dst [][]int, src [][]int) {
	dst[0][0] = src[0][0]
	dst[0][1] = src[0][1]
	dst[0][2] = src[0][2]
	dst[1][0] = src[1][0]
	dst[1][1] = src[1][1]
	dst[1][2] = src[1][2]
	dst[2][0] = src[2][0]
	dst[2][1] = src[2][1]
	dst[2][2] = src[2][2]

}

// BoardStringer converts a board to a string
func BoardStringer(board [][]int) string {
	str := ""
	for _, i := range board[0] {
		str += strconv.Itoa(i)
	}
	for _, i := range board[1] {
		str += strconv.Itoa(i)
	}
	for _, i := range board[2] {
		str += strconv.Itoa(i)
	}
	return str
}

// StatePrinter prints a state
func StatePrinter(state [][]int) {
	fmt.Println("_______")
	fmt.Println(state[0])
	fmt.Println(state[1])
	fmt.Println(state[2])
	fmt.Println("-------")
}

// InputParser parses a txt input to a board
func InputParser(fileName string) [][]int {
	board := make([][]int, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	rawInput := make([]byte, 17)
	_, err = file.Read(rawInput)

	var inputArray []int
	var inpStr []string
	inputLines := strings.Split(string(rawInput), "\n")
	for i := 0; i < len(inputLines); i++ {
		inpStr = strings.Split(inputLines[i], " ")
		for j := 0; j < len(inpStr); j++ {
			ij, err := strconv.ParseInt(inpStr[j], 0, 32)
			if err != nil {
				log.Fatal(err)
			}
			inputArray = append(inputArray, int(ij))
		}
	}
	board[0] = inputArray[0:3]
	board[1] = inputArray[3:6]
	board[2] = inputArray[6:9]
	return board

}

// Solvable receives two boards and returns if the game is solvable
func Solvable(board [][]int, goal [][]int) bool {
	var invGoal int
	var boardArray, goalArray [9]int

	counter := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardArray[counter] = board[i][j]
			goalArray[counter] = goal[i][j]
			counter++
		}
	}

	fmt.Println(boardArray)

	for i := range goalArray {
		for j := i + 1; j < len(goal); j++ {
			if goalArray[i] > goalArray[j] && goalArray[j] != 0 {
				invGoal++
			}
		}
	}
	invBoard := 0
	for i := range boardArray {
		for j := i + 1; j < len(boardArray); j++ {
			if boardArray[i] > boardArray[j] && boardArray[j] != 0 {
				invBoard++
			}
		}
	}
	return invGoal%2 == invBoard%2
}

package main

import (
	"fmt"
	"time"

	"github.com/rafaelescrich/8-puzzle/astar"
	"github.com/rafaelescrich/8-puzzle/search"
	"github.com/rafaelescrich/8-puzzle/utils"
)

func main() {

	fmt.Println("Iniciando...", time.Now())
	startTime := time.Now()
	board := utils.InputParser("input.txt")
	goal := utils.InputParser("goal.txt")
	if board == nil {
		return
	}

	if !utils.Solvable(board, goal) {
		fmt.Println("Sem solução")
		return
	}
	utils.StatePrinter(board)
	startState := search.NewState(board, goal)
	var solution *search.State

	solution, expanded := astar.Solve(startState, goal)

	if solution == nil {
		fmt.Println("Sem solução")
		return
	}
	s := solution
	steps := make([]*search.State, solution.NumMoves)
	for i := 0; i < solution.NumMoves; i++ {
		steps[i] = s
		s = s.Parent
	}
	for i := len(steps)/2 - 1; i >= 0; i-- {
		opp := len(steps) - 1 - i
		steps[i], steps[opp] = steps[opp], steps[i]
	}
	for _, next := range steps {
		utils.StatePrinter(next.Board)
	}

	fmt.Println("Custo da solução:", solution.NumMoves)
	fmt.Println("Nós expandidos: ", expanded)
	elapsed := time.Since(startTime)
	fmt.Println("Processo levou cerca de:", elapsed)
	return
}

package astar

import (
	"container/heap"
	_ "fmt"

	"github.com/rafaelescrich/8-puzzle/priorityQueue"
	"github.com/rafaelescrich/8-puzzle/search"
	"github.com/rafaelescrich/8-puzzle/utils"
)

// Solve is the function to solve the game
func Solve(start search.State, goal [][]int) *search.State {
	var frontier, expanded int
	states := make(map[string]search.State)
	pq := make(priorityQueue.PriorityQueue, 0)
	key := utils.BoardStringer(start.Board)
	states[key] = start
	heap.Push(&pq, &priorityQueue.Item{Value: key, Priority: 0, Index: 0})

	for pq.Len() != 0 {

		currentItem := heap.Pop(&pq).(*priorityQueue.Item)
		current := states[currentItem.Value]
		expanded++

		if current.IsGoal(goal) {
			return &current
		}
		for _, next := range current.PossibleMoves(goal) {
			key := utils.BoardStringer(next.Board)
			if old, exists := states[key]; !exists || next.Distance < old.Distance {
				states[key] = next
				heap.Push(&pq, &priorityQueue.Item{Value: key, Priority: next.Distance + next.NumMoves, Index: 0})
				frontier++
			}
		}
	}
	return nil
}

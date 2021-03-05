package main

import "fmt"

const (
	PLAYER     = "X"
	CLEAR_PATH = "."
	OBSTACLE   = "#"
	TREASURE   = "$"
)

// Initiate MAP
var MAP = [][]string{
	{OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE},
	{OBSTACLE, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, TREASURE, OBSTACLE},
	{OBSTACLE, CLEAR_PATH, OBSTACLE, OBSTACLE, OBSTACLE, CLEAR_PATH, CLEAR_PATH, OBSTACLE},
	{OBSTACLE, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, OBSTACLE, CLEAR_PATH, OBSTACLE, OBSTACLE},
	{OBSTACLE, CLEAR_PATH, OBSTACLE, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, CLEAR_PATH, OBSTACLE},
	{OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE, OBSTACLE},
}

// Position
type Position struct {
	x int
	y int
}

// isVisited
// Check position has been visited or not
func isVisited(visited []Position, pos Position) bool {
	for _, v := range visited {
		if v == pos {
			return true
		}
	}
	return false
}

// bfs
// Find treasure location using Breadth-first search algorithm (bfs)
// ref: https://en.wikipedia.org/wiki/Breadth-first_search
func bfs(grid [][]string, start Position) []Position {
	queue := [][]Position{{start}}
	visited := []Position{start}

	height := len(grid)
	width := len(grid[0])

	for len(queue) != 0 {
		path := queue[0]
		if len(queue) > 2 {
			queue = queue[1:]
		}
		pos := path[len(path)-1]
		if grid[pos.y][pos.x] == TREASURE {
			return path
		}
		right := Position{pos.x + 1, pos.y}
		left := Position{pos.x - 1, pos.y}
		bottom := Position{pos.x, pos.y + 1}
		top := Position{pos.x, pos.y - 1}
		directions := []Position{top, bottom, right, left}
		for _, d := range directions {
			if (d.x >= 0 && d.x < width) && (d.y >= 0 && d.y < height) && (grid[d.y][d.x] != OBSTACLE) && !isVisited(visited, d) {
				z := append(path, d)
				queue = append([][]Position{z}, queue...)
				visited = append(visited, d)
			}
		}
	}

	return nil
}

// render
// Render player position
func render(grid [][]string, playerPosition Position)  {
	fmt.Printf("Player find treasure at position: x = %d, y = %d\n", playerPosition.x, playerPosition.y)
	for y, _ := range grid {
		for x, _ := range grid[y] {
			node := grid[y][x]
			if y == playerPosition.y && x == playerPosition.x {
				node = PLAYER
			}
			fmt.Print(node)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	start := Position{1, 4}
	positions := bfs(MAP, start)
	for _, position := range positions {
		render(MAP, position)
	}
	treasurePosition := positions[len(positions) - 1]
	fmt.Printf("Tresure found at position: x = %d, y = %d\n", treasurePosition.x, treasurePosition.y)
}

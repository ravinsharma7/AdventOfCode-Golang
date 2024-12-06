package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Position struct {
	X, Y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directions = []Position{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

// Parse the input map
func parseMap(input string) ([][]rune, Position, Direction) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var guardPos Position
	var guardDir Direction

	for i, line := range lines {
		grid[i] = []rune(line)
		for j, char := range line {
			switch char {
			case '^':
				guardPos = Position{i, j}
				guardDir = Up
			case '>':
				guardPos = Position{i, j}
				guardDir = Right
			case 'v':
				guardPos = Position{i, j}
				guardDir = Down
			case '<':
				guardPos = Position{i, j}
				guardDir = Left
			}
		}
	}
	return grid, guardPos, guardDir
}

// Simulate the guard's patrol with an optional obstruction
func simulateWithObstruction(grid [][]rune, guardPos Position, guardDir Direction, obstruction Position) bool {
	visitedStates := map[struct {
		Position
		Direction
	}]bool{}

	rows := len(grid)
	cols := len(grid[0])

	// Add the obstruction if specified
	if obstruction.X >= 0 && obstruction.Y >= 0 {
		grid[obstruction.X][obstruction.Y] = '#'
	}

	for {
		state := struct {
			Position
			Direction
		}{guardPos, guardDir}

		if visitedStates[state] {
			return true // Loop detected
		}
		visitedStates[state] = true

		nextPos := Position{
			X: guardPos.X + directions[guardDir].X,
			Y: guardPos.Y + directions[guardDir].Y,
		}

		if nextPos.X < 0 || nextPos.X >= rows || nextPos.Y < 0 || nextPos.Y >= cols {
			break
		}

		if grid[nextPos.X][nextPos.Y] == '#' {
			guardDir = (guardDir + 1) % 4
		} else {
			guardPos = nextPos
		}
	}

	return false
}

// Find valid obstruction positions in parallel
func findValidObstructions(grid [][]rune, guardPos Position, guardDir Direction) int {
	rows := len(grid)
	cols := len(grid[0])
	var validCount int
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}

	// Divide work into chunks for parallel processing
	chunkSize := rows / 4 // Number of rows per chunk
	for startRow := 0; startRow < rows; startRow += chunkSize {
		endRow := startRow + chunkSize
		if endRow > rows {
			endRow = rows
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localCount := 0

			for x := start; x < end; x++ {
				for y := 0; y < cols; y++ {
					if grid[x][y] == '.' && !(x == guardPos.X && y == guardPos.Y) {
						if simulateWithObstruction(copyGrid(grid), guardPos, guardDir, Position{x, y}) {
							localCount++
						}
					}
				}
			}

			// Update global count with thread-safe access
			mutex.Lock()
			validCount += localCount
			mutex.Unlock()
		}(startRow, endRow)
	}

	wg.Wait()
	return validCount
}

// Utility function to copy a grid
func copyGrid(grid [][]rune) [][]rune {
	copied := make([][]rune, len(grid))
	for i := range grid {
		copied[i] = append([]rune{}, grid[i]...)
	}
	return copied
}

func main() {
	input := `
#.............#.#...........................................................#..........#...........................#..........#...
............#......................#.......#..#.........#...................##....................................................
..............#...........#..#.......#.........#.#...........................#................................#......#............
.....#......................................#......#..#...................#....................#......................#...........
...#............................#................................#................................................................
..........#..................#.....#......................................................#.......................................
..#........#.......#.#...............................................................#......#.................................#...
.#....#................#................#.....#...................#...............................................#..#............
.......................................................................................#...#.#....#........##......#........#.#.##
............................................#.................#.....................#..........................#..................
....#..........##......................#.................................................................................#........
......#.............##.........#.......#.........#.......#.............#...#.......#..............................................
.#............................................................................#..........#....#...................................
...#..#....#..............................#...........................................................................#...........
.............#....................................................................................................................
................................#..................##...................................#...............................#......#..
..........#.........#.................#..........###...........................................#............#...#.................
.............#....................................#.......#.#.#............#............................................#.........
......#.#..............................................................................................#.........................#
...#...#.........#.........#..#......#................##.................#........................................................
.................................#......................#........#................#........................#......................
.......................##..............#................#................#..............................#.........................
.......#...#.....#.....##........................#..................##.....#......................................................
.............................................#...........................#..........................................#.....#.......
.........................#...............#....#..............#.......#...............#.#.................#.................#......
...........................................................#....#..................................#..............................
..........................#.................................................#............................##....#..................
.................#...#.........................................................#........................#.....#..#................
.........................#....#................................................#...#.....................#....#...........#.......
................................................##.......#.#....................................#.......##........................
..........##...#.......................#....#..#....................#........................#....................................
#......#.......................#............#.....................................................................#...............
.......#.......#...#..#..........#.......................................#.............#......#...................................
...............#............#..................#.................................................................................#
.......#...........#.......................##.....#.........#...........#...................#......#.........#.............##....#
...................#.........................#......#..........................................#..........#.........#......#......
.....#............#..............#...#......#....................#.............#..................................................
................................................#........#.....#...................#.....................................#.....#..
......................#.#...##......................##...............................................#.................#..........
#...................#.........................#...........#.....#................#.....................#...................#......
....#................#.....#.............#.....................................#....#...#.................................#.......
.........#.....................................................................................................#.#.......#........
.........#..........##..............................#............................#.#.............................#................
#....#.#..............................................##...#...#...............#.............#....................................
........#..#.#...................#.......................................................................#.............#..#.#...#.
.................................................................................................................#.......#........
............................................#..#...............##........#..........#.......................................#....#
........#...........................#............................#.....#........................#..........#...#..........#.......
.#...........#...................................................#...........................................#................#.#.
............#...............................#........#...........#..........#...........#.#..............#........................
...............................................................................##.#................#..............................
.#.......................#.#...........#..................................#.......................................................
..#..............................................##................#.......#....#..............#.....................#...#........
...................#.....#................................................#..........#.........#.........#.....................#..
........#................#................................................................................#......#................
...................#...#....#.........................#.........................#....................#..........................#.
.................#.......................................................#................................................#.......
..#...........................##.#.............#..#............................................#.......................#.#........
..#......#.........................................................................#.......................#......................
...........................................................................................................#............#......#..
............................................................^..........................................#..........#...............
..#.........#....................#..........................................#........#.........................................#..
........#...............#....#.............#......#.........................................................................#.....
.........#......#.....#..............#...........................................................#..............#.................
.........#...................................................#...###...........#.......#.##.......................................
.....................##.................#.........##.....................#................#................#.........#............
.......................................................................#..............#..............#.................#..........
.##.........#..........................................................................#....................................#.....
.........................#.............................#........................................................................#.
........#.............................................#......#....#.....................................#....................#....
........#...#..............................................................................#.........#............................
..............................#................................#..........#..............#..................................#.....
........##..#...................................#................................#.................#.........#......#..#........#.
..................#......#...................................................................#....................................
..........#.............#..............#........#............#...............#..#..........#..............#.#...#............#....
............................##....#..............#......................................................................#.........
...............................................#.....#...........#.....##...............#..#......................................
.........................#.......#...............................#...........#.........#...........#....#.........................
.#............#...#...............#..............................................................#..........................#....#
#...........#......................................................#.............................#......#.........................
..................................................#....#...............#.........................#................................
................................................#..........................#.......#..............................#....##.......#.
........#......#..#............#................#............................#......#.##.........#....................#..#........
.......................................................#..#.#....................#..#......................................#......
.....................#.......................................................................#............#.......................
.....................##........................#...................#.....................#........................................
..........................#...........................................................................................#...........
..#..........#...................#..............##..#.....................#...................#..............#...........#........
#................................#....##......#.............#..............#.......................#.........#...........#........
..................#................#...............................#.....#..........#................#..............#.............
................#.#.........#.....##...............................................................#..............#...............
..#...............................................................................................#...............................
..........#.................#...........#.......#.........................#...#.......#...........................................
..##....#......................................#.#.......................................#.................#......................
.......................#..#...#.##......................................................................#.........#...#...........
...............#.........#..............#.....................................................#...................#...............
......#.................#............#......#...................#..#.........................#........................#........#..
......................................................................#...........................#........#......................
..........#..........#........................#.............##....................................................................
....#................................................#.......................#................#...............#...................
.................#....................................................#......#....................................................
.#........#.......................................................................................................................
......##..................#....#......#.................#....................#.....................#..#...........#...............
#...#...#...#..............#........#.......................#.#..........#...............................#......#.................
#...#...................#.............#........#..........#...............................................................#...#...
..#................................#...#.........................#...................#......#.....................................
...#...###..........#............#...............................................................#..........#.....................
............#..............#.......#.....#..........#.....#.........................................................#............#
..#......................#..#...#................................#......#.......#.....#..................#.........#..............
....##......#.......................................................#..............................................#..............
...............................#.....#....#.................................................#.......#............#................
..........#....#......................#........#..........................#......#................................................
......#.............................#...........#..................#...........................#................................#.
...........................##..........................................#.....#.#.....................#................#...........
..#..............#.........................................................#........................#.........#.................#.
#.....#.....#......................#....................#........................................#...........#...........#......#.
.......................................#..#................................#.........#.......#..#.....#..........................#
..........#..............##...............#..........................................................#......#......#......#....#..
.............#..............#..#......#........#..#......##.......................#...........#.......#...........................
......................#.......................................#.......#..#.....................................#.........#...##...
...........................#...........................#......#....#.............................#.............................#..
..................#............................#....#........##.........................#..............#.............#............
.......#....................#.....#........#.......#.#...#...............#.....#.#.........................#.......#.#............
.....#.........#.......#..........#.................#..#...#..................................................##..................
................#........................................................#........................................................
....##......#........#..................#.........................#...................................#...#...................#...
..............#....................#.......#.......................#.............#......#..............#..........................
...........................#..............#.....#............#.....................#........#....#.......................#........
........#..##................#.....#.#..............#..............................#........................#.........#...........
.........................................................#...........................##..........#........#.##....................
`

	// Start the timer
	start := time.Now()

	// Parse the map and calculate valid obstructions
	grid, startPos, startDir := parseMap(input)
	result := findValidObstructions(grid, startPos, startDir)

	// Stop the timer
	elapsed := time.Since(start)

	// Display the results
	fmt.Printf("Number of valid obstruction positions: %d\n", result)
	fmt.Printf("Execution time: %s\n", elapsed)
	fmt.Printf("Execution time (nanoseconds): %d ns\n", elapsed.Nanoseconds())
	fmt.Printf("Execution time (microseconds): %.3f µs\n", float64(elapsed.Nanoseconds())/1000)
	fmt.Printf("Execution time (milliseconds): %.3f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
}

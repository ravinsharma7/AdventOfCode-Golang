package main

import (
	"fmt"
	"strings"
)

// Check if three points are collinear using cross product:
// (x2 - x1)*(y3 - y1) - (y2 - y1)*(x3 - x1) == 0
func collinear(x1, y1, x2, y2, x3, y3 int) bool {
	return (x2 - x1)*(y3 - y1) == (y2 - y1)*(x3 - x1)
}

// Parse the input map and identify antennas
func parseMap(input string) ([][]rune, map[rune][][2]int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	// Map from frequency (character) to list of antenna coordinates
	antennas := make(map[rune][][2]int)

	for i, line := range lines {
		grid[i] = []rune(line)
		for j, ch := range line {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], [2]int{i, j})
			}
		}
	}
	return grid, antennas
}

func main() {
	// Example input from the problem statement
	input := `
...O.....0...............................p..k.....
O.........o....w..T.........................p.....
..................w..........oM...................
.............................................Y....
o.............T...........................z.....pk
.....................................z..Y....t.F..
...........T..........................F.......Y...
...................A............z...k..M..........
....O.........j....w.....................M........
..........w....T..................M..k............
.............5.............................F.....t
......................A.............F....E........
.....................S.........A..................
.P................................................
........................A...E.............x...t...
............j.........................t.........x.
.......................j.........................x
....................................E........c....
.............P.......E............................
...............j..5...............q...............
..............P..............................Qc...
..........C..........s................c........x..
.............C...r................................
.....C......V..........q...................Q......
...........yX.........q...................Q.......
.....X....................e.............m.........
.2.................e..7....m.......c..............
......i..........e...K..............f....U...WQ...
...X....................e....................V...Y
...............5..X.....0.........................
..C..i......5..3...sK......J.........f..B.........
2............3.............0I...a.........BNb.....
.........................K............m...........
.r........3...............s....7...m.v...f.......b
........3........7.....Iy..........q...b.N........
.....R.1.......................n.....a.B.......VN.
......R.........9...................a...W.........
..........7.6................S....................
..............r.......s...0........nb....W..f..B..
...2...........I......K...........................
..............................u...n............U..
............r......y.............O............W...
.......R..........v..u................N...V.......
..........R.8..4.9..y........u....................
...8...............v................J.............
.....8..............4.........Z.........n.....J.U.
...........4i....................Z..S.............
..............9...........1.u.S................J..
8...6....9..4......a........Z.1...................
....................v..i.............Z............`
	grid, antennas := parseMap(input)
	rows := len(grid)
	cols := len(grid[0])

	antinodeSet := make(map[[2]int]bool)

	// For each frequency, consider all pairs of antennas (A,B)
	// Every point on the line formed by (A,B) that lies within the map is an antinode
	// Also, the antennas themselves (if that frequency has >=2 antennas) are antinodes
	for _, antList := range antennas {
		if len(antList) < 2 {
			// If only one antenna of this frequency, no line can form => no antinodes from it
			continue
		}
		// Mark all antennas of this freq as antinodes since they line up with at least one other
		for _, a := range antList {
			antinodeSet[a] = true
		}

		// Generate all pairs
		for i := 0; i < len(antList); i++ {
			for j := i + 1; j < len(antList); j++ {
				A := antList[i]
				B := antList[j]

				// For each point in the grid, check if it is collinear with A and B
				for x := 0; x < rows; x++ {
					for y := 0; y < cols; y++ {
						if collinear(A[0], A[1], B[0], B[1], x, y) {
							antinodeSet[[2]int{x, y}] = true
						}
					}
				}
			}
		}
	}

	// Count how many unique antinode positions
	result := len(antinodeSet)

	fmt.Println(result)
}

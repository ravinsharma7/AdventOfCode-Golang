package main

import (
	"fmt"
	"strings"
)

// Position represents a coordinate on the map
type Position struct {
	x, y int
}

// Directions for movement: up, down, left, right
var directions = []Position{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func main() {
	// Input topographic map as a multi-line string
	input := `
210121012321234586543237650120901076540121001
101234985490147897890198701221812389430430012
567105676787458721209877654336765410321569123
498014785010969630118960165445898923210678234
309323494323874549823454278321080854100123945
211234515690123678989923329876571765987637876
110987601781010501677815410989432894543540345
021010532782323432596506901098126545652901276
012123445693437010487417892387087036761892987
323098530696598923300306103456592123890743210
014567621787127654211215296503431014125650309
329643452650034018987894387112321265034321478
498762167541005329876012379012450996543406567
567433098932116456705430498763567887456715458
016526187654327643212321567654654321019823309
327817656783298984103435433216789670321094215
438908943090125675678976124105894589452385434
567617812181034106789780045674583076581276123
106787801654323295610191230983012109690187012
245096998745016784327898101092103298781298432
332165432101565410678765432387654567654382321
456976541001274328769867211498565650103401450
697889034012389219654358300567656743212519567
787765121965476101011249401456549865435658498
657854320876210894320134562340134678321067019
548943013456367885012021076543221589801354323
038765412387478976523256789601100499832231234
129914302198565985430143898702341300740140145
878801243089478876943232169219654211655657656
965760651078769217854321078348778902598748912
054450782369854308765010127651267143456732103
123321895451043219678098763410356012349801654
103456986032112008769123452345441078959798723
212387872149800123452101011096732163478760810
301095433458945654343000125687823454307451923
454346721567238769289210234776916543210312367
965239810123189378176321349885407892321208498
870125676034003210065443456790398901765089567
345676787165612349898652105381210340894176506
210985696278973656776789965450678256783215410
101974320161087298785649874367589109891006321
007876013232196105694234701298432098782107890
012565434983435234523105650165601105673456921
323476325876524321012023498074320234504305430
434789016789012010123110567789210343212214321
`

	// Parse the input into a 2D grid of integers
	grid := parseInput(input)
	rows := len(grid)
	if rows == 0 {
		fmt.Println(0)
		return
	}
	cols := len(grid[0])

	// Find all trailheads (positions with height 0)
	trailheads := []Position{}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 0 {
				trailheads = append(trailheads, Position{x, y})
			}
		}
	}

	// For each trailhead, compute the number of reachable height 9 positions
	sumScores := 0
	for _, th := range trailheads {
		score := computeScore(grid, th, rows, cols)
		sumScores += score
	}

	// Output the sum of all trailhead scores
	fmt.Println(sumScores)
}

// parseInput converts the input string into a 2D slice of integers
func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			grid[y][x] = int(char - '0')
		}
	}
	return grid
}

// computeScore performs BFS from the trailhead and counts reachable height 9 positions
func computeScore(grid [][]int, start Position, rows, cols int) int {
	// Queue for BFS
	queue := []Position{start}
	// Visited map to prevent revisiting positions
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[start.y][start.x] = true

	// Set to store unique reachable height 9 positions
	reachableNines := make(map[Position]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Current height
		currentHeight := grid[current.y][current.x]

		// If current height is 9, add to reachableNines
		if currentHeight == 9 {
			reachableNines[current] = true
			// Do not continue from a height of 9
			continue
		}

		// Next expected height
		nextHeight := currentHeight + 1

		// Explore all four directions
		for _, dir := range directions {
			newX := current.x + dir.x
			newY := current.y + dir.y

			// Check boundaries
			if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
				continue
			}

			// Check if already visited
			if visited[newY][newX] {
				continue
			}

			// Check if the next position has the exact next height
			if grid[newY][newX] == nextHeight {
				visited[newY][newX] = true
				queue = append(queue, Position{newX, newY})
			}
		}
	}

	// The score is the number of unique reachable height 9 positions
	return len(reachableNines)
}

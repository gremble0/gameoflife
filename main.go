package main

import (
    "math/rand"
    "fmt"
    "time"
)

const (
    DEAD = " \033[91mx\033[0m"
    ALIVE = " \033[92mx\033[0m"

    ROWS = 20
    COLUMNS = 20
)

func initializeBoard() [ROWS][COLUMNS]string {
    var cells [ROWS][COLUMNS]string
    for i := 0; i < ROWS; i++ {
        for j := 0; j < COLUMNS; j++ {
	    random := rand.Intn(10)
	    if random == 1 { 
		cells[j][i] = ALIVE 
	    } else { 
		cells[j][i] = DEAD 
	    }
        }
    }

    return cells
}

func drawBoard(cells [ROWS][COLUMNS]string) {
    for i := 0; i < ROWS; i++ {
        for j := 0; j < COLUMNS; j++ {
	    fmt.Printf(cells[j][i])
        }
	fmt.Printf("\n")
    }
}

func updateBoard(cells [ROWS][COLUMNS]string) [ROWS][COLUMNS]string {
    for row := 0; row < ROWS; row++ {
	for col := 0; col < COLUMNS; col++ {
	    aliveNeighbors := 0
	    // Row above
	    if row > 0 {
		if col > 0 && cells[row - 1][col - 1] == ALIVE { aliveNeighbors++ }
		if cells[row - 1][col] == ALIVE { aliveNeighbors++ }
		if col + 1 < COLUMNS && cells[row - 1][col + 1] == ALIVE { aliveNeighbors++ }
	    }

	    // Current row
	    if col > 0 && cells[row][col - 1] == ALIVE { aliveNeighbors++ }
	    if col + 1 < COLUMNS && cells[row][col + 1] == ALIVE { aliveNeighbors++ }

	    // Row below
	    if row + 1 < ROWS {
		if col > 0 && cells[row + 1][col - 1] == ALIVE { aliveNeighbors++ }
		if cells[row + 1][col] == ALIVE { aliveNeighbors++ }
		if col + 1 < COLUMNS && cells[row + 1][col + 1] == ALIVE { aliveNeighbors++ }
	    }

	    if aliveNeighbors == 3 && cells[row][col] == DEAD { cells[row][col] = ALIVE }
	    if !(2 <= aliveNeighbors && aliveNeighbors <= 3) && cells[row][col] == ALIVE { cells[row][col] = DEAD }
	}
    }

    return cells
}

func main() {
    var cells = initializeBoard()

    drawBoard(cells)

    for true {
	fmt.Println()
	time.Sleep(time.Second)

	// Optionally clear terminal after each iteration
	// fmt.Print("\033[H\033[2J")

	cells = updateBoard(cells)
	drawBoard(cells)
    }
}

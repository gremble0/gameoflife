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
    for i := 0; i < ROWS; i++ {
	for j := 0; j < COLUMNS; j++ {

	}
    }
    return cells
}

func main() {
    var cells = initializeBoard()

    drawBoard(cells)

    for true {
	time.Sleep(time.Second)
	// Clear terminal
	fmt.Print("\033[H\033[2J")

	cells = updateBoard(cells)
	drawBoard(cells)
    }
}

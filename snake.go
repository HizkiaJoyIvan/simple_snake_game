package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const (
	width  = 20
	height = 20
)

var (
	gameOver bool
	x, y     int
	fruitX   int
	fruitY   int
	score    int
)

type eDirection int

const (
	STOP eDirection = iota
	LEFT
	RIGHT
	UP
	DOWN
)

func setup() {
	gameOver = false
	x = width / 2
	y = height / 2
	fruitX = rand.Intn(width)
	fruitY = rand.Intn(height)
	score = 0
}

func draw() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println('#')

}

func input() {

}

func logic() {

}

func main() {
	draw()
	for !gameOver {
		draw()
		input()
		logic()
	}
}

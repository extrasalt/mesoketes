package main

import "fmt"

//A battle consists of several days
type Battle struct {
	Days []Day
}

//Each day has several attacks
type Day struct {
	Attacks []Attack
	N       int
	E       int
	S       int
	W       int
}

//Alias int to direction to create Direction constants
type Direction int
type Attack struct {
	Dir      Direction
	Strength int
}

const (
	North = iota
	East
	South
	West
)

//City has a wall
type City struct {
	Wall
}

type Wall struct {
	BreachCount int //Total Number of breaches.
	N           int //Northern Wall's strength
	E           int //Eastern
	S           int //Southern
	W           int //Western
}

func main() {

	fmt.Println(West)
}

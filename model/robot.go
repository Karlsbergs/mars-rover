package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Robot struct {
	X           int64
	Y           int64
	Orientation string
	Commands    string

	Grid Grid
	Lost bool
}

func (r *Robot) Init(build string, grid Grid) error {
	var err error

	build = strings.ReplaceAll(build, " ", "")

	if !strings.HasPrefix(build, "(") {
		return fmt.Errorf("'%s' does not start with an open bracket (", build)
	}

	build = build[1:]

	closedBracketIndex := strings.Index(build, ")")

	if closedBracketIndex == -1 {
		return fmt.Errorf("'%s' no closed bracket )", build)
	}

	coordinateVals := strings.Split(build[0:closedBracketIndex], ",")

	if len(coordinateVals) != 3 {
		return fmt.Errorf("'%s' missing coordinates", build)
	}

	// Get X coordinate
	r.X, err = strconv.ParseInt(coordinateVals[0], 10, 32)
	if err != nil {
		return fmt.Errorf("'%s' x coordinates is not numeric", build)
	}

	// Get Y coordinate
	r.Y, err = strconv.ParseInt(coordinateVals[1], 10, 32)
	if err != nil {
		return fmt.Errorf("'%s' y coordinates is not numeric", build)
	}

	r.Orientation = coordinateVals[2]

	if r.Orientation != "N" && r.Orientation != "E" && r.Orientation != "S" && r.Orientation != "W" {
		return fmt.Errorf("'%s' orientation not one N,E,S or W", build)
	}

	// Get commands
	r.Commands = build[closedBracketIndex+1:]

	if len(r.Commands) == 0 {
		return fmt.Errorf("'%s' no commands", build)
	}

	// Check all commands are legit characters
	for _, c := range r.Commands {
		cmd := string(c)

		if cmd != "L" && cmd != "F" && cmd != "R" {
			return fmt.Errorf("'%s' commands must be one of L, F or R", build)
		}
	}

	// Check grid size value
	if grid.M < 1 || grid.N < 1 {
		return fmt.Errorf("'%s' invalid grid size", build)
	}

	r.Grid = grid

	return nil
}

// RunCommands - Run all commands, directing robot around grid
func (r *Robot) RunCommands() string {
	for _, c := range r.Commands {
		if r.Lost {
			// If robot is lost break from the loop
			break
		}

		switch string(c) {
		case "L":
			r.TurnLeft()
		case "R":
			r.TurnRight()
		case "F":
			r.MoveForwardForward()
		}
	}

	return r.ToString()
}

func (r *Robot) TurnLeft() {
	switch r.Orientation {
	case "N":
		r.Orientation = "W"
	case "E":
		r.Orientation = "N"
	case "S":
		r.Orientation = "E"
	case "W":
		r.Orientation = "S"
	}
}

func (r *Robot) TurnRight() {
	switch r.Orientation {
	case "N":
		r.Orientation = "E"
	case "E":
		r.Orientation = "S"
	case "S":
		r.Orientation = "W"
	case "W":
		r.Orientation = "N"
	}
}

// MoveForwardForward - Moves the robot forward, direction depends on orientation
// If CheckInRange fails, the X\Y position is reverted to get last know good position
func (r *Robot) MoveForwardForward() {	
	switch r.Orientation {
	case "N":
		r.Y++
		if !r.CheckInRange() {
			r.Y--
		}
	case "E":
		r.X++
		if !r.CheckInRange() {
			r.X--
		}
	case "S":
		r.Y--
		if !r.CheckInRange() {
			r.Y++
		}
	case "W":
		r.X--
		if !r.CheckInRange() {
			r.X++
		}
	}
}

// CheckInRange - Check the robot coordinates are still within the grid boundaries
func (r *Robot) CheckInRange() bool {
	if r.X < 0 || r.X > r.Grid.N || r.Y < 0 || r.Y > r.Grid.M {
		r.Lost = true
		return false
	}

	return true
}

// ToString - Show robot final coordinates and Orientation
func (r *Robot) ToString() string {
	lost := ""

	if r.Lost {
		lost = " LOST"
	}

	return fmt.Sprintf("(%d, %d, %s)%s", r.X, r.Y, r.Orientation, lost)
}
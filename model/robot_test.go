package model

import (
	"testing"
)

// Test creating a new robot with build string
func TestInit(t *testing.T) {
	var err error
	robot := Robot{}
	
	err = robot.Init("   (0   ,    0  ,  E )       LF   R    ", Grid{M: 4, N: 8})
	if err != nil {
		t.Errorf("Create should not have created an error. Error: %s", err)
	}

	err = robot.Init("0,0,N) LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed missing ( check")
	}		

	err = robot.Init("(0,0,N LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed missing ) check")
	}		

	err = robot.Init("(0,0) LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed missing coordinates check")
	}	

	err = robot.Init("(x,0,E) LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed X numeric check")
	}

	err = robot.Init("(0,x,E) LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed Y numeric check")
	}	

	err = robot.Init("(0,0,C) LFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed orientation value check")
	}	
	
	err = robot.Init("(0,0,N) ", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed missing command check")
	}		

	err = robot.Init("(0,0,N) OLFR", Grid{M: 4, N: 8})
	if err == nil {
		t.Errorf("Should have failed command value check")
	}	

	err = robot.Init("(0,0,N) LFR", Grid{M: 0, N: 0})
	if err == nil {
		t.Errorf("Should have failed grid value check")
	}	
}

// Test turing left Orientation
func TestLeft(t *testing.T) {
	grid := Grid{M: 4, N: 8}

	robot := Robot{X: 0, Y: 0, Orientation: "N", Commands: "", Grid: grid}

	if robot.Orientation != "N" {
		t.Errorf("Result '%s' expected 'N'", robot.Orientation)
	}

	robot.TurnLeft()

	if robot.Orientation != "W" {
		t.Errorf("Result '%s' expected 'W'", robot.Orientation)
	}

	robot.TurnLeft()

	if robot.Orientation != "S" {
		t.Errorf("Result '%s' expected 'S'", robot.Orientation)
	}

	robot.TurnLeft()

	if robot.Orientation != "E" {
		t.Errorf("Result '%s' expected 'E'", robot.Orientation)
	}

	robot.TurnLeft()

	if robot.Orientation != "N" {
		t.Errorf("Result '%s' expected 'N'", robot.Orientation)
	}
}

// Test turing right orientation
func TestRight(t *testing.T) {
	grid := Grid{M: 4, N: 8}

	robot := Robot{X: 0, Y: 0, Orientation: "N", Commands: "", Grid: grid}

	if robot.Orientation != "N" {
		t.Errorf("Result '%s' expected 'N'", robot.Orientation)
	}

	robot.TurnRight()

	if robot.Orientation != "E" {
		t.Errorf("Result '%s' expected 'E'", robot.Orientation)
	}

	robot.TurnRight()

	if robot.Orientation != "S" {
		t.Errorf("Result '%s' expected 'S'", robot.Orientation)
	}

	robot.TurnRight()

	if robot.Orientation != "W" {
		t.Errorf("Result '%s' expected 'W'", robot.Orientation)
	}

	robot.TurnRight()

	if robot.Orientation != "N" {
		t.Errorf("Result '%s' expected 'N'", robot.Orientation)
	}
}

func TestMoveForwardNorth(t *testing.T) {
	grid := Grid{M: 4, N: 4}

	robot := Robot{X: 0, Y: 3, Orientation: "N", Commands: "", Grid: grid}

	// Move robot inside the grid
	robot.MoveForwardForward()

	// A move forward should move the robot by 1 row
	if robot.X != 0 || robot.Y != 4 || robot.Orientation != "N" || robot.Lost != false {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:3 Orientation:N Lost:False'", robot.X, robot.Y, robot.Orientation)
	}

	// Move robot outside the grid
	robot.MoveForwardForward()

	// A move forward from these Coordinates will result in a lost robot. Previous coordinates should be intact.
	if robot.X != 0 || robot.Y != 4 || robot.Orientation != "N"  || robot.Lost != true {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:5 Orientation:N Lost:True'", robot.X, robot.Y, robot.Orientation)
	}
}

func TestMoveForwardWEST(t *testing.T) {
	grid := Grid{M: 4, N: 4}

	robot := Robot{X: 1, Y: 0, Orientation: "W", Commands: "", Grid: grid}

	// Move robot inside the grid
	robot.MoveForwardForward()

	// A move forward should move the robot by 1 row
	if robot.X != 0 || robot.Y != 0 || robot.Orientation != "W" || robot.Lost != false {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:0 Orientation:W Lost:False'", robot.X, robot.Y, robot.Orientation)
	}

	// Move robot outside the grid
	robot.MoveForwardForward()

	// A move forward from these Coordinates will result in a lost robot. Previous coordinates should be intact.
	if robot.X != 0 || robot.Y != 0 || robot.Orientation != "W"  || robot.Lost != true {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:0 Orientation:N Lost:True'", robot.X, robot.Y, robot.Orientation)
	}
}

func TestMoveForwardSOUTH(t *testing.T) {
	grid := Grid{M: 4, N: 4}

	robot := Robot{X: 0, Y: 1, Orientation: "S", Commands: "", Grid: grid}

	// Move robot inside the grid
	robot.MoveForwardForward()

	// A move forward should move the robot by 1 row
	if robot.X != 0 || robot.Y != 0 || robot.Orientation != "S" || robot.Lost != false {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:0 Orientation:S Lost:False'", robot.X, robot.Y, robot.Orientation)
	}

	// Move robot outside the grid
	robot.MoveForwardForward()

	// A move forward from these Coordinates will result in a lost robot. Previous coordinates should be intact.
	if robot.X != 0 || robot.Y != 0 || robot.Orientation != "S"  || robot.Lost != true {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:0 Y:0 Orientation:S Lost:True'", robot.X, robot.Y, robot.Orientation)
	}
}

func TestMoveForwardEAST(t *testing.T) {
	grid := Grid{M: 4, N: 4}

	robot := Robot{X: 3, Y: 0, Orientation: "E", Commands: "", Grid: grid}

	// Move robot inside the grid
	robot.MoveForwardForward()

	// A move forward should move the robot by 1 row
	if robot.X != 4 || robot.Y != 0 || robot.Orientation != "E" || robot.Lost != false {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:4 Y:0 Orientation:E Lost:False'", robot.X, robot.Y, robot.Orientation)
	}

	// Move robot outside the grid
	robot.MoveForwardForward()

	// A move forward from these Coordinates will result in a lost robot. Previous coordinates should be intact.
	if robot.X != 4 || robot.Y != 0 || robot.Orientation != "E"  || robot.Lost != true {
		t.Errorf("Result 'X:%d Y:%d Orientation:%s' expected 'X:4 Y:0 Orientation:E Lost:True'", robot.X, robot.Y, robot.Orientation)
	}
}

// Test e2e scenarios via RunCommand
func TestAll(t *testing.T) {
	grid := Grid{M: 4, N: 8}

	testRobot(t, "(2, 3, E) LFRFF", grid, "(4, 4, E)")
	testRobot(t, "(0, 2, N) FFLRFF", grid, "(0, 4, N) LOST")
	testRobot(t, "(2, 3, N) FLLFR", grid, "(2, 3, W)")
	testRobot(t, "(1, 0, S) FFRLF", grid, "(1, 0, S) LOST")
}

func testRobot(t *testing.T, build string, grid Grid, expectedResult string) {
	robot := Robot{}
	robot.Init(build, grid)

	result := robot.RunCommands()

	if result != expectedResult {
		t.Errorf("Result '%s', expected '%s'", result, expectedResult)
	}
}
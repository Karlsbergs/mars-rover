package main

import (
	"bufio"
	"fmt"
	"marsrover/model"
	"os"
	"strings"
)

var Robots []model.Robot

// Example robot build configs:
// (2, 3, E) LFRFF
// (0, 2, N) FFLRFF
// (2, 3, N) FLLFR
// (1, 0, S) FFRLF

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get grid size
	gridValues, _ := reader.ReadString('\n')

	grid := model.Grid{}
	err := grid.Init(gridValues)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// Get each robot build config
		robotBuild, _ := reader.ReadString('\n')
		robotBuild = strings.Replace(robotBuild, "\n", "", -1)

		if robotBuild == "" {
			// End receiving new robots
			break
		}

		RobotDeploy(grid, robotBuild)
	}

	RobotRunCommands()
}

func RobotDeploy(grid model.Grid, build string) {
	robot := model.Robot{}
	err := robot.Init(build, grid)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save robot
	Robots = append(Robots, robot)
}

func RobotRunCommands() {
	for _, r := range Robots {
		result := r.RunCommands()
		fmt.Println(result)
	}
}

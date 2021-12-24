package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Grid struct {
	M int64
	N int64
}

func (g *Grid) Init(build string) error {
	var err error

	build = strings.TrimSpace(build)

	// Remove any duplicate spaces between values
	space := regexp.MustCompile(`\s+`)
	build = space.ReplaceAllString(build, " ")	

	// Split build string to get m & n values
	vals := strings.Split(build, " ")

	if len(vals) != 2 {
		return fmt.Errorf("'%s' 2 values required", build)
	}

	// Get M value
	g.M, err = strconv.ParseInt(vals[0], 10, 32)
	if err != nil {
		return fmt.Errorf("'%s' m grid value is not numeric", build)
	}

	// Get N value
	g.N, err = strconv.ParseInt(vals[1], 10, 32)
	if err != nil {
		return fmt.Errorf("'%s' n grid value is not numeric", build)
	}	

	return nil
}
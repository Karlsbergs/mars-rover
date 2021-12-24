package model

import "testing"

func TestCreate(t *testing.T) {
	var err error
	grid := Grid{}

	err = grid.Init("   4    8   ")
	if err != nil {
		t.Errorf("Should not have failed. Error: %s", err)
	}

	err = grid.Init("")
	if err == nil {
		t.Errorf("Should have failed missing values check")
	}

	err = grid.Init("x 2")
	if err == nil {
		t.Errorf("Should have failed non-numeric values check")
	}

	err = grid.Init("1 x")
	if err == nil {
		t.Errorf("Should have failed non-numeric values check")
	}
}
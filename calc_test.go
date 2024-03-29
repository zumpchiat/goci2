package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}

}

func TestSub(t *testing.T) {

	result := sub(2, 3)

	if result != -1 {
		t.Error("The result must be 1")
	}
}

func TestMult(t *testing.T) {

	result := mult(2, 3)

	if result != 6 {
		t.Error("The result must be 6")
	}
}

func TestDiv(t *testing.T) {

	result := div(10, 2)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

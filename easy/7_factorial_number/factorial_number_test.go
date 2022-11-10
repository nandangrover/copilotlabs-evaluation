package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	if factorial(1) != 1 {
		t.Errorf("factorial(1) = %d; want 1", factorial(1))
	}

	if factorial(2) != 2 {
		t.Errorf("factorial(2) = %d; want 2", factorial(2))
	}

	if factorial(3) != 6 {
		t.Errorf("factorial(3) = %d; want 6", factorial(3))
	}

	if factorial(4) != 24 {
		t.Errorf("factorial(4) = %d; want 24", factorial(4))
	}

	if factorial(5) != 120 {
		t.Errorf("factorial(5) = %d; want 120", factorial(5))
	}

	if factorial(6) != 720 {
		t.Errorf("factorial(6) = %d; want 720", factorial(6))
	}
}

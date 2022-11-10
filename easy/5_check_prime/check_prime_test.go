package main

import "testing"

func TestPrime(t *testing.T) {
	if isPrime(1) {
		t.Errorf("isPrime(1) = true, false expected")
	}
	if !isPrime(17) {
		t.Errorf("isPrime(17) = false, true expected")
	}
	if !isPrime(2) {
		t.Errorf("isPrime(2) = false, true expected")
	}
	if !isPrime(3) {
		t.Errorf("isPrime(3) = false, true expected")
	}
}

package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	if !isLeapYear(2000) {
		t.Error("expected true, got false")
	}
	if isLeapYear(2001) {
		t.Error("expected false, got true")
	}
	if isLeapYear(2002) {
		t.Error("expected false, got true")
	}
	if isLeapYear(2003) {
		t.Error("expected false, got true")
	}
	if !isLeapYear(2004) {
		t.Error("expected true, got false")
	}
	if isLeapYear(2100) {
		t.Error("expected false, got true")
	}
}

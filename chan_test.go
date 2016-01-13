package main

import "testing"

func TestAge(t *testing.T) {
	bob := Person{16, "Bob"}

	if bob.is_underage() != true {
		t.Errorf("Underage error")
	}
}

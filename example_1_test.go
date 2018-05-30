package main // this test belongs to the main package

import (
	"testing" // a test must import the testing package
)

func TestAlwaysPassing(t *testing.T) {
	t.Log("this test always passes")
}

func TestAlwaysFailing(t *testing.T) {
	// t.Error function reports failure, the string it accepts represents
	// the error that causes this test to fail
	t.Error("this test always fails")
}

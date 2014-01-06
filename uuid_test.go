package main

import (
	"testing"
)

func TestMake(t *testing.T) {
	u, err := MakeUUID()
	if err != nil {
		t.Errorf("Expected to generate UUID without problems, error thrown: %s", err.Error())
		return
	}
	if len(u) != 16 {
		t.Errorf("Should have been a 16 byte guid: %s", (*u).String())
		return
	}
}

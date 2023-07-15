package main

import (
	"fmt"
	"testing"
)

// "https://steamcommunity.com/profiles/76561197993684328"
func TestParseSteamProfile(t *testing.T) {
	err, result := profileToId64("https://steamcommunity.com/id/ximf")
	fmt.Println("test ")
	t.Log(result)
	if err != nil {
		t.Error(err)
	}

	if result == 76561197993684328 {
		t.Logf("Result is correct: %v", result)
	} else {
		t.Errorf("Result is incorrect: %v", result)
	}

}

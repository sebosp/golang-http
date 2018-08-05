package main

import "testing"
import "os"
import "reflect"

func TestGetEnvVar(t *testing.T) {
	// Arrange
	os.Unsetenv("NONEXISTENT")
	os.Setenv("EMPTY", "")
	os.Setenv("EXISTENT", "hello")
	cases := []struct {
		input       string
		value       string
		shouldError bool
	}{
		{"NONEXISTENT", "", true},
		{"EXISTENT", "hello", false},
		{"EMPTY", "", false},
	}
	for _, c := range cases {
		// Act:
		response, err := GetEnvVar(c.input)
		// Assert
		if err == nil {
			if c.shouldError {
				t.Errorf("On input '%s', expected error, got success.", c.input)
				continue
			}
			if response != c.value {
				t.Errorf("On input '%s', expected '%s', got '%s'.", c.input, c.value, response)
			}
		} else if !c.shouldError {
			t.Errorf("On input '%s', expected success, got error: %s.", c.input, err)
		}
	}

}
func TestSetDataFromEnv(t *testing.T) {
	// Arrange Full data available
	os.Setenv("PLAYER_NAME", "GoTester")
	os.Setenv("ENV_NAME", "TESTING_SUITE")
	os.Setenv("COLOR", "blue")
	playerTest1 := &Player{
		name:        "x",
		environment: "x",
		color:       "gray",
	}
	playerExpect := &Player{
		name:        "GoTester",
		environment: "TESTING_SUITE",
		color:       "blue",
	}
	// Act:
	err := playerTest1.setDataFromEnv()
	// Assert
	if err != nil {
		t.Errorf("For properly set env, expected success, got error: %s", err)
	} else {
		if !reflect.DeepEqual(playerExpect, playerTest1) {
			t.Errorf("For properly set env, expected data '%+v', got '%+v'.", playerExpect, playerTest1)
		}
	}
	// Arrange Optional data not available, the color will not be overwritten.
	os.Setenv("PLAYER_NAME", "GoTester")
	os.Setenv("ENV_NAME", "TESTING_SUITE")
	os.Unsetenv("COLOR")
	playerTest2 := &Player{
		name:        "x",
		environment: "x",
		color:       "red",
	}
	playerExpect.color = "red"
	// Act
	err = playerTest2.setDataFromEnv()
	// Assert
	if err != nil {
		t.Errorf("For partially set env, expected success, got error: %s", err)
	} else {
		if !reflect.DeepEqual(playerExpect, playerTest2) {
			t.Errorf("For partially set env, expected data '%+v', got '%+v'.", playerExpect, playerTest2)
		}
	}
	// Arrange Required ENV_NAME is unset
	os.Setenv("PLAYER_NAME", "GoTester")
	os.Unsetenv("ENV_NAME")
	playerTest3 := &Player{
		name:        "x",
		environment: "x",
		color:       "red",
	}
	// Act
	err = playerTest3.setDataFromEnv()
	// Assert
	if err == nil {
		t.Errorf("For unset required vars, expected error, got success: %+v", playerTest3)
	}
	// Arrange Required PLAYER_NAME is unset
	os.Setenv("ENV_NAME", "TESTING_SUITE")
	os.Unsetenv("PLAYER_NAME")
	playerTest4 := &Player{
		name:        "x",
		environment: "x",
		color:       "red",
	}
	// Act
	err = playerTest4.setDataFromEnv()
	// Assert
	if err == nil {
		t.Errorf("For unset required vars, expected error, got success: %+v", playerTest4)
	}
}
func TestGatherData(t *testing.T) {
	// Arrange Full data available
	os.Setenv("PLAYER_NAME", "GoTester")
	os.Setenv("ENV_NAME", "TESTING_SUITE")
	os.Setenv("COLOR", "blue")
	playerTest1 := &Player{
		name:        "x",
		environment: "x",
		color:       "gray",
	}
	playerExpect := &Player{
		name:        "GoTester",
		environment: "TESTING_SUITE",
		color:       "blue",
	}
	// Act
	err := playerTest1.GatherData()
	// Assert
	if err != nil {
		t.Errorf("For properly set env, expected success, got error: %s", err)
	} else {
		if !reflect.DeepEqual(playerExpect, playerTest1) {
			t.Errorf("For properly set env, expected data '%+v', got '%+v'.", playerExpect, playerTest1)
		}
	}
}

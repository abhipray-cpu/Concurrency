package main

/*
Use model checking tools to explore all possible states or configurations of a concurrent system to verify properties like safety (nothing bad happens) and liveness (something good eventually happens).
Tools like SPIN or TLA+ can be used for formal verification of concurrent algorithms
*/

import (
	"testing"
)

func TestTrafficLightStateTransitions(t *testing.T) {
	light := NewTrafficLight()

	// Initial state should be Red
	if light.State() != Red {
		t.Fatalf("Expected initial state to be Red, got %s", light.State())
	}

	// Test state transitions
	expectedTransitions := []TrafficLightState{Green, Yellow, Red}
	for _, expectedState := range expectedTransitions {
		light.ChangeState()
		if light.State() != expectedState {
			t.Fatalf("Expected state to be %s, got %s", expectedState, light.State())
		}
	}
}

func TestTrafficLightFullCycle(t *testing.T) {
	light := NewTrafficLight()

	// Perform a full cycle of state transitions
	for i := 0; i < 3; i++ {
		light.ChangeState()
	}

	// The state should be back to Red after a full cycle
	if light.State() != Red {
		t.Fatalf("Expected state to be Red after a full cycle, got %s", light.State())
	}
}

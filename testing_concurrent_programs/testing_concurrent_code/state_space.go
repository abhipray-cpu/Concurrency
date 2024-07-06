package main

type TrafficLightState string

const (
	Green  TrafficLightState = "Green"
	Yellow TrafficLightState = "Yellow"
	Red    TrafficLightState = "Red"
)

type TrafficLight struct {
	state TrafficLightState
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{state: Red}
}

func (t *TrafficLight) ChangeState() {
	switch t.state {
	case Red:
		t.state = Green
	case Green:
		t.state = Yellow
	case Yellow:
		t.state = Red
	}
}

func (t *TrafficLight) State() TrafficLightState {
	return t.state
}

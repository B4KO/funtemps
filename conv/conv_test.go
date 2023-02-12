package conv

import "testing"
import "reflect"

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 500, want: 260},
		{input: 500.45, want: 260.25},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
		{input: 500, want: 533.15},
		{input: 500.45, want: 533.40},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelciusToFarenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56, want: 132.8},
		{input: 500, want: 932},
		{input: 500.45, want: 932.81},
	}

	for _, tc := range tests {
		got := CelciusToFarenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelciusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56, want: 329.15},
		{input: 500, want: 773.15},
		{input: 500.45, want: 773.6},
	}

	for _, tc := range tests {
		got := CelciusToKelvin(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelcius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56, want: -217.15},
		{input: 500, want: 226.85},
		{input: 500.45, want: 227.3},
	}

	for _, tc := range tests {
		got := KelvinToCelcius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56, want: -358.87},
		{input: 500, want: 440.33},
		{input: 500.45, want: 441.14},
	}

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

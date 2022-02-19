package main

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "print stuff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hello()
		})
	}
}

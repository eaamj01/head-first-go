package main

import "testing"

func Test_exercise1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "pass",
			want: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exercise1(); got != tt.want {
				t.Errorf("exercise1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Add 1 + 2",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		}, {
			name: "Add 2.5 + -2",
			args: args{
				a: 2.5,
				b: -2,
			},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

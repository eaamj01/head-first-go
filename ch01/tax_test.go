package main

import "testing"

// func Test_tax(t *testing.T) {
// 	type args struct {
// 		price   int
// 		taxRate float64
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		{
// 			name: "pass",
// 			args: args{
// 				price:   100,
// 				taxRate: 0.08,
// 			},
// 			want: 108,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tax(tt.args.price, tt.args.taxRate); got != tt.want {
// 				t.Errorf("tax() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_tax(t *testing.T) {
	type args struct {
		price   int
		taxRate float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 bool
	}{
		{
			name: "In budget",
			args: args{
				price:   100,
				taxRate: 0.08,
			},
			want:  108,
			want1: true,
		},
		{
			name: "Not in budget",
			args: args{
				price:   200,
				taxRate: 0.08,
			},
			want:  216,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tax(tt.args.price, tt.args.taxRate)
			if got != tt.want {
				t.Errorf("tax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

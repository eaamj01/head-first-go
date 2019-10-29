package compare

import "testing"

func TestLarger(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test first larger",
			args: args{
				a: 2,
				b: 1,
			},
			want: 2,
		},
		{
			name: "Test second larger",
			args: args{
				a: 4,
				b: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Larger(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Larger(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

package prose

import (
	"fmt"
	"testing"
)

func TestJoinWithCommas(t *testing.T) {
	type args struct {
		phrases []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 2 elements",
			args: args{
				phrases: []string{"apple", "orange"},
			},
			want: "apple and orange",
		}, {
			name: "Join 3 elements",
			args: args{
				phrases: []string{"apple", "orange", "pear"},
			},
			want: "apple, orange, and pear",
		}, {
			name: "Just 1 element",
			args: args{
				phrases: []string{"apple"},
			},
			want: "apple",
		}, {
			name: "No element",
			args: args{
				phrases: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinWithCommas(tt.args.phrases); got != tt.want {
				t.Errorf("JoinWithCommas(%#v) = %v, want %v", tt.args.phrases, got, tt.want)
			}
		})
	}
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(erroString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(erroString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(erroString(list, got, want))
	}
}

func erroString(list []string, got string, want string) string {
	// Display the slice passed to JoinWithCommas in debug format
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

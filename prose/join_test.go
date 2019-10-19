package prose

import "testing"

func TestJoinWithCommas(t *testing.T) {
	type args struct {
		phrases []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "JoinWithCommas",
			args: args{
				phrases: []string{"apple", "orange"},
			},
			want: "apple and orange",
		}, {
			name: "JoinWithCommas",
			args: args{
				phrases: []string{"apple", "orange", "pear"},
			},
			want: "apple, orange, and pear",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinWithCommas(tt.args.phrases); got != tt.want {
				t.Errorf("JoinWithCommas() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestTwoElements(t *testing.T) {
// 	list := []string{"apple", "orange"}
// 	if JoinWithCommas(list) != "apple and orange" {
// 		t.Error("didn't match expected value")
// 	}
// }

// func TestThreeElements(t *testing.T) {
// 	list := []string{"apple", "orange", "pear"}
// 	if JoinWithCommas(list) != "apple, orange, and pear" {
// 		t.Error("didn't match expected value")
// 	}
// }

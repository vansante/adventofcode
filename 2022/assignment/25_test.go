package assignment

import "testing"

func Test_snafu2dec(t *testing.T) {
	tests := []struct {
		snafu string
		want  int64
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"21", 11},
		{"1=-1=", 353},
		{"12111", 906},
		{"20012", 1257},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	}
	for _, tt := range tests {
		t.Run(tt.snafu, func(t *testing.T) {
			got := snafu2dec(tt.snafu)
			if got != tt.want {
				t.Errorf("snafu2dec() = %v, want %v", got, tt.want)
			}
			snaf := dec2snafu(got)
			if dec2snafu(got) != tt.snafu {
				t.Errorf("dec2snafu() = %v, want %v", snaf, tt.snafu)
			}
		})
	}
}

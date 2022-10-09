package main

import "testing"

func TestHexToInt(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Quick Real data Test",
			args: args{"0x0000000000000000000000000000000000000000000000000000000000011bb4"},
			want: 72628,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToInt(tt.args.hex); got != tt.want {
				t.Errorf("HexToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

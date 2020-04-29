package timeout

import "testing"

func TestExec(t *testing.T) {
	type args struct {
		waitSecond int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec(tt.args.waitSecond)
		})
	}
}

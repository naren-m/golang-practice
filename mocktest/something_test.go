package main

import "testing"

func TestDoThings(t *testing.T) {
	type args struct {
		in0 string
		in1 uint64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoThings(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoThings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoThings() = %v, want %v", got, tt.want)
			}
		})
	}
}

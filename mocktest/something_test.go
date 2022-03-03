package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

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
		name: "test name",
		args: ['stirng', 24]
		want: 24
		wantErr: False
	},
	{
		// TODO: Add test cases.
		name: "Negative name",
		args: ['stirng', 24]
		want: 23
		wantErr: True
	}
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

func TestDoNothing(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoNothing()
		})
	}
}

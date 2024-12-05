package main

import "testing"

func Test_sendingTickets(t *testing.T) {
	type args struct {
		userTickets uint
		firstName   string
		lastName    string
		email       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendingTickets(tt.args.userTickets, tt.args.firstName, tt.args.lastName, tt.args.email)
		})
	}
}

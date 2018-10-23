package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var tests = []struct {
	name        string
	password    string
	wantErr     bool
	cost        int
	expectedErr interface{}
}{
	{
		name:     "Hash sample password",
		password: "foo",
	},
	{
		name:     "Hash very short password",
		password: "x",
	},
	{
		name:     "Hash very short password",
		password: "x",
	},
	{
		name:     "Hash long password",
		password: "012345678901234567890123456789012345678901234567890123456",
	},
	{
		name:        "Invalid cost (too low)",
		cost:        1,
		expectedErr: "cost 1 is outside allowed range",
		wantErr:     true,
	},
	{
		name:        "Invalid cost (too high)",
		cost:        32,
		expectedErr: "cost 32 is outside allowed range",
		wantErr:     true,
	},
	{
		name:        "No password provided",
		expectedErr: "you must provide the password through stdin",
		wantErr:     true,
	},
}

func TestBcryptGenerateCmd_Execute(t *testing.T) {

	for _, tt := range tests {
		cmd := NewBcryptGenerateCmd()

		t.Run(tt.name, func(t *testing.T) {
			var err error

			expectedCost := DefaultCost
			if tt.cost != 0 {
				expectedCost = tt.cost
			}
			cmd.Cost = expectedCost
			out := &bytes.Buffer{}
			cmd.OutWriter = out
			if tt.password != "" {
				in := bytes.NewBufferString(tt.password)
				cmd.InReader = in
			}
			err = cmd.Execute([]string{})

			stdout := out.String()

			if (err != nil) != tt.wantErr {
				t.Errorf("BcryptGenerateCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if tt.expectedErr != nil {
					assert.Regexp(t, tt.expectedErr, err, "Expected error %v to match %v", err, tt.expectedErr)
				}
			} else {
				require.NoError(t, bcrypt.CompareHashAndPassword([]byte(stdout), []byte(tt.password)))
			}
		})
	}
}

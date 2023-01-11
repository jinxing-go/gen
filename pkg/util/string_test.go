package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudly(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "user to User",
			args: "user",
			want: "User",
		},
		{
			name: "userName to UserName",
			args: "userName",
			want: "UserName",
		},
		{
			name: "user_id to UserId",
			args: "user_id",
			want: "UserId",
		},
		{
			name: "user name to UserName",
			args: "user name",
			want: "UserName",
		},
		{
			name: "UserName to UserName",
			args: "UserName",
			want: "UserName",
		},
		{
			name: "User Name to UserName",
			args: "User Name",
			want: "UserName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Studly(tt.args); got != tt.want {
				t.Errorf("Studly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncludes(t *testing.T) {
	assert.True(t, Includes(1, []int{1, 2, 3}))
	assert.False(t, Includes(4, []int{1, 2, 3}))

	assert.True(t, Includes("1", []string{"1", "2", "3"}))
	assert.False(t, Includes("10", []string{"1", "2", "3"}))
}

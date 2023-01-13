package mysql

import "testing"

func Test_toType(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "varchar(16)",
			args: "varchar(16)",
			want: "varchar",
		},
		{
			name: "decimal(10,2)",
			args: "decimal(10,2)",
			want: "decimal",
		},
		{
			name: "int(11)",
			args: "int(11)",
			want: "int",
		},
		{
			name: "bigint(20)",
			args: "bigint(20)",
			want: "bigint",
		},
		{
			name: "int unsigned",
			args: "int unsigned",
			want: "int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toType(tt.args); got != tt.want {
				t.Errorf("toType() = %v, want %v", got, tt.want)
			}
		})
	}
}

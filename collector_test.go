package main

import "testing"

func TestCallGroup(t *testing.T) {
	tests := []struct {
		name string
		row  map[string]interface{}
		want string
	}{
		{
			name: "top-level data",
			row:  map[string]interface{}{"data": "sales"},
			want: "sales",
		},
		{
			name: "top-level presence_data",
			row:  map[string]interface{}{"presence_data": "support"},
			want: "support",
		},
		{
			name: "top-level variable_data",
			row:  map[string]interface{}{"variable_data": "support"},
			want: "support",
		},
		{
			name: "nested variables data",
			row:  map[string]interface{}{"variables": map[string]interface{}{"data": "billing"}},
			want: "billing",
		},
		{
			name: "missing data",
			row:  map[string]interface{}{"uuid": "abc"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callGroup(tt.row); got != tt.want {
				t.Fatalf("callGroup() = %q, want %q", got, tt.want)
			}
		})
	}
}

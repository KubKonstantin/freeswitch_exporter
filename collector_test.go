package main

import "testing"

func TestCallGroup(t *testing.T) {
	tests := []struct {
		name string
		row  map[string]interface{}
		want string
	}{
		{
			name: "top-level group",
			row:  map[string]interface{}{"group": "telemed"},
			want: "telemed",
		},
		{
			name: "top-level variable_group",
			row:  map[string]interface{}{"variable_group": "telemed"},
			want: "telemed",
		},
		{
			name: "top-level data fallback",
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
			name: "nested variables data fallback",
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

func TestParseShowRows(t *testing.T) {
	response := []byte("uuid,direction,presence_data,variable_data\nabc,inbound,sales,\ndef,outbound,,support\n\n2 total.\n")

	rows, err := parseShowRows(response)
	if err != nil {
		t.Fatalf("parseShowRows() error = %v", err)
	}
	if got, want := len(rows), 2; got != want {
		t.Fatalf("len(rows) = %d, want %d", got, want)
	}
	if got, want := callGroup(rows[0]), "sales"; got != want {
		t.Fatalf("callGroup(rows[0]) = %q, want %q", got, want)
	}
	if got, want := callGroup(rows[1]), "support"; got != want {
		t.Fatalf("callGroup(rows[1]) = %q, want %q", got, want)
	}
}

func TestCleanVariableValue(t *testing.T) {
	tests := []struct {
		name  string
		value []byte
		want  string
	}{
		{name: "plain value", value: []byte("sales\n"), want: "sales"},
		{name: "undefined", value: []byte("_undef_\n"), want: ""},
		{name: "error", value: []byte("-ERR No such channel!\n"), want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanVariableValue(tt.value); got != tt.want {
				t.Fatalf("cleanVariableValue() = %q, want %q", got, tt.want)
			}
		})
	}
}

package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func Test_readOutputFile(t *testing.T) {
	type test struct {
		arg  string
		want string
	}

	tests := []test{
		{"2020-02-25", shroveTuesday},
		{"2021-02-16", shroveTuesday},
		{"2022-03-01", shroveTuesday},
		{"2023-02-21", shroveTuesday},
		{"2024-02-13", shroveTuesday},
		{"2025-03-04", shroveTuesday},
		{"2026-02-17", shroveTuesday},
		{"2027-02-09", shroveTuesday},
		{"2028-02-29", shroveTuesday},
		{"2029-02-13", shroveTuesday},
	}

	dates := make(map[string]Date)
	b, err := os.ReadFile("../dates.json")
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(b, &dates)

	for _, tt := range tests {
		t.Run("shroveTuesdayTest", func(t *testing.T) {
			d := dates[tt.arg]
			seen := false
			for _, k := range d.Specialday {
				if k == tt.want {
					seen = true
				}
			}
			if !seen {
				t.Errorf("%s not found under special days for %s\n", tt.want, tt.arg)
			}
		})
	}
}

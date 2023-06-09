package cai2hcl

import (
	"testing"
)

func TestConvertProject(t *testing.T) {
	cases := []struct {
		name string
	}{
		{name: "project_create"},
		{name: "project_iam"},
	}
	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			err := assertTestData(c.name)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

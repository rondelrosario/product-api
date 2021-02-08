package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "ron",
		Price: 1.00,
		SKU:   "ron-ron-rooon",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

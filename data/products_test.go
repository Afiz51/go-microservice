package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Gboyega",
		Price: 1,
		SKU:   "abcd-ef-e",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

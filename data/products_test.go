package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 1.99,
		SKU:   "abc-abd-bcd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Fatal("Success")
}

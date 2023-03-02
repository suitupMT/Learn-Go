package main

// in go there are no classes
//but this is a custom created object
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
//constructor, takes a name inside and returns a bill object
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

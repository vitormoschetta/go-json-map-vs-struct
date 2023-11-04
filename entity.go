package main

type Supplier struct {
	Name    string
	Address string
}

type Product struct {
	Name  string
	Price float64
}

type Order struct {
	OrderID  int
	Supplier Supplier
	Items    []Product
}

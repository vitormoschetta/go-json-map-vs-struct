package main

type Address struct {
	Street string
	City   string
}

type Supplier struct {
	Name    string
	Address Address
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

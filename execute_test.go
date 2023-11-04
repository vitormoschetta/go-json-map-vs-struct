package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalOk(t *testing.T) {
	// Arrange
	inputJSON := `{
		"OrderID": 1,
		"Supplier": {
			"Name": "Acme",
			"Address": "123 Main St"
		},
		"Items": [
			{
				"Name": "Widget",
				"Price": 1.99
			},
			{
				"Name": "Gadget",
				"Price": 2.99
			}
		]
	}`

	// Act
	var order Order
	err := json.Unmarshal([]byte(inputJSON), &order)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, order.OrderID)
	assert.Equal(t, "Acme", order.Supplier.Name)
	assert.Equal(t, "123 Main St", order.Supplier.Address)
	assert.Equal(t, 2, len(order.Items))
	assert.Equal(t, "Widget", order.Items[0].Name)
	assert.Equal(t, 1.99, order.Items[0].Price)
}

func TestUnmarshalToMapStringInterfaceOk(t *testing.T) {
	// Arrange
	inputJSON := `{
		"OrderID": 1,
		"Supplier": {
			"Name": "Acme",
			"Address": "123 Main St"
		},
		"Items": [
			{
				"Name": "Widget",
				"Price": 1.99
			},
			{
				"Name": "Gadget",
				"Price": 2.99
			}
		]
	}`

	// Act
	var order map[string]interface{}
	err := json.Unmarshal([]byte(inputJSON), &order)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1.0, order["OrderID"]) // Note: JSON numbers are always float64 | pt-br: Números JSON são sempre float64.
	assert.Equal(t, "Acme", order["Supplier"].(map[string]interface{})["Name"])
	assert.Equal(t, "123 Main St", order["Supplier"].(map[string]interface{})["Address"])
	assert.Equal(t, 1.99, order["Items"].([]interface{})[0].(map[string]interface{})["Price"])
}

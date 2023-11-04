package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalToStructOk(t *testing.T) {
	// Arrange
	inputJSON := buildJSON()
	var order Order

	// Act
	err := json.Unmarshal([]byte(inputJSON), &order)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, order.OrderID)
	assert.Equal(t, "Acme", order.Supplier.Name)
	assert.Equal(t, "123 Main St", order.Supplier.Address.Street)
	assert.Equal(t, 2, len(order.Items))
	assert.Equal(t, "Widget", order.Items[0].Name)
	assert.Equal(t, 1.99, order.Items[0].Price)
}

func TestUnmarshalToMapOk(t *testing.T) {
	// Arrange
	inputJSON := buildJSON()
	var order map[string]interface{}

	// Act
	err := json.Unmarshal([]byte(inputJSON), &order)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1.0, order["OrderID"]) // Note: JSON numbers are always float64 | pt-br: Números JSON são sempre float64.
	assert.Equal(t, "Acme", order["Supplier"].(map[string]interface{})["Name"])
	assert.Equal(t, "123 Main St", order["Supplier"].(map[string]interface{})["Address"].(map[string]interface{})["Street"])
	assert.Equal(t, 1.99, order["Items"].([]interface{})[0].(map[string]interface{})["Price"])
}

func TestConvertMapToStructOk(t *testing.T) {
	// Arrange
	inputJSON := buildJSON()
	var order map[string]interface{}
	err := json.Unmarshal([]byte(inputJSON), &order)
	assert.Nil(t, err)

	// Act
	var orderStruct Order
	err = MapToStruct(order, &orderStruct)
	assert.Nil(t, err)

	// Assert
	assert.Equal(t, 1, orderStruct.OrderID)
	assert.Equal(t, "Acme", orderStruct.Supplier.Name)
	assert.Equal(t, "123 Main St", orderStruct.Supplier.Address.Street)
	assert.Equal(t, 2, len(orderStruct.Items))
	assert.Equal(t, "Widget", orderStruct.Items[0].Name)
	assert.Equal(t, 1.99, orderStruct.Items[0].Price)
}

func TestConvertStructToMapOk(t *testing.T) {
	// Arrange
	inputJSON := buildJSON()
	var order Order
	err := json.Unmarshal([]byte(inputJSON), &order)
	assert.Nil(t, err)

	// Act
	orderMap, err := StructToMap(order)

	// Assert
	assert.Equal(t, 1.0, orderMap["OrderID"]) // Note: JSON numbers are always float64 | pt-br: Números JSON são sempre float64.
	assert.Equal(t, "Acme", orderMap["Supplier"].(map[string]interface{})["Name"])
	assert.Equal(t, "123 Main St", orderMap["Supplier"].(map[string]interface{})["Address"].(map[string]interface{})["Street"])
	assert.Equal(t, 1.99, orderMap["Items"].([]interface{})[0].(map[string]interface{})["Price"])
}

func buildJSON() string {
	return `{
		"OrderID": 1,
		"Supplier": {
			"Name": "Acme",
			"Address": {
				"Street": "123 Main St",
				"City": "Anytown"
			}
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
}

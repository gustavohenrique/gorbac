package gorbac_test

import (
	"testing"

	rbac "github.com/gustavohenrique/gorbac"
	"github.com/gustavohenrique/gorbac/test/assert"
)

func TestConvertSimpleStructToMap(t *testing.T) {
	type Struct struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}
	instance := Struct{ID: 1, Name: "Gustavo"}
	converted := rbac.Struct(instance).ToMap()
	assert.Equal(t, converted["id"], instance.ID)
	assert.Equal(t, converted["name"], instance.Name)
}

func TestConvertPointerStructToMap(t *testing.T) {
	type Struct struct {
		ID int32 `json:"id"`
	}
	instance := Struct{ID: 1}
	converted := rbac.Struct(&instance).ToMap()
	assert.Equal(t, converted["id"], instance.ID)
}

func TestConvertNestedStructToMap(t *testing.T) {
	type Product struct {
		Name string `json:"name"`
	}
	product := Product{Name: "Notebook"}
	type Order struct {
		Product Product `json:"product"`
	}
	order := Order{product}
	converted := rbac.Struct(order).ToMap()
	productMap := converted["product"].(map[string]interface{})
	assert.Equal(t, productMap["name"], product.Name)
}

func TestConvertStructWithSliceToMap(t *testing.T) {
	type Struct struct {
		Items []string `json:"items"`
	}
	items := []string{"item1", "item2"}
	instance := Struct{items}
	converted := rbac.Struct(instance).ToMap()
	list := converted["items"].([]string)
	assert.Equal(t, len(list), len(items))
	assert.Equal(t, list[0], items[0])
}

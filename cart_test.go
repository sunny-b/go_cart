package cart

import (
	"testing"
)

func TestPut(t *testing.T) {
	mockItem := &Item{
		itemBase: &itemBase{
			name:        "mock",
			description: "mock item",
			price:       9.99,
			productID:   "product-id",
		},
		id: "item-id",
	}
	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	q, p, err := cart.Put(mockItem)
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if q != 1 {
		t.Log("should equal 1 quantity")
		t.Fail()
	}
	if p != 9.99 {
		t.Log("price should equal 9.99")
		t.Fail()
	}
}

func TestPutTwoItems(t *testing.T) {
	mockItem1 := &Item{
		itemBase: &itemBase{
			name:        "mock",
			description: "mock item",
			price:       9.99,
			productID:   "product-id",
		},
	}
	mockItem2 := &Item{
		itemBase: &itemBase{
			name:        "mock",
			description: "mock item",
			price:       9.99,
			productID:   "product-id",
		},
		id: "item-id2",
	}

	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	cart.Put(mockItem1)

	q, p, err := cart.Put(mockItem2)
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if q != 2 {
		t.Log("should equal 2 quantity")
		t.Fail()
	}
	if p != 19.98 {
		t.Logf("price should equal 19.98 but equal %v", p)
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	mockItem := &Item{
		itemBase: &itemBase{
			name:        "mock",
			description: "mock item",
			price:       9.99,
			productID:   "product-id",
		},
		id: "item-id",
	}

	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	q, p, err := cart.Put(mockItem)
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if q != 1 {
		t.Log("should equal 1 quantity")
		t.Fail()
	}
	if p != 9.99 {
		t.Logf("price should equal 9.99 but equal %v", p)
		t.Fail()
	}

	item, err := cart.Get("item-id")
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if item.id != mockItem.id {
		t.Logf("name should equal %v but equals %v", mockItem.id, item.id)
		t.Fail()
	}
}

func TestGetFail(t *testing.T) {
	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	item, err := cart.Get("item-id")
	if err == nil {
		t.Log("should return with error")
		t.Fail()
	}
	if item != nil {
		t.Logf("name should equal nil but equals %v", item)
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	mockItem := &Item{
		itemBase: &itemBase{
			name:        "mock",
			description: "mock item",
			price:       9.99,
			productID:   "product-id",
		},
		id: "item-id",
	}

	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	q, p, err := cart.Put(mockItem)
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if q != 1 {
		t.Log("should equal 1 quantity")
		t.Fail()
	}
	if p != 9.99 {
		t.Logf("price should equal 9.99 but equal %v", p)
		t.Fail()
	}

	item, err := cart.Delete("item-id")
	if err != nil {
		t.Log("should not return with error")
		t.Fail()
	}
	if item.id != mockItem.id {
		t.Logf("name should equal %v but equals %v", mockItem.id, item.id)
		t.Fail()
	}
	if cart.quantity != 0 {
		t.Logf("cart should be empty but has %v items", cart.quantity)
		t.Fail()
	}
}

func TestDeleteFail(t *testing.T) {
	cart := &cart{
		items:      make(map[string]ItemWrapper),
		quantity:   0,
		totalPrice: 0,
	}

	item, err := cart.Delete("item-id")
	if err == nil {
		t.Log("should return with error")
		t.Fail()
	}
	if item != nil {
		t.Logf("name should equal nil but equals %v", item)
		t.Fail()
	}
}

package cart

import "errors"

// Cart describes putting, removing, and getting items into the cart
type Cart interface {
	Put(ci *Item) (int, float32, error)
	Get(itemID string) (*Item, error)
	Delete(itemID string) (*Item, error)
}

type cart struct {
	quantity   int
	totalPrice float32
	items      map[string]ItemWrapper
}

// Item is a struct describing a product item
type Item struct {
	*itemBase
	id string
}

func (c *cart) Put(i *Item) (int, float32, error) {
	if ci, ok := c.items[i.productID]; ok {
		ci.Add(i)
	} else {
		c.addNewItem(i)
	}

	c.putUpdate(i.price)

	return c.quantity, c.totalPrice, nil
}

func (c *cart) Get(itemID string) (*Item, error) {
	ci, present := c.findCartItem(itemID)
	if !present {
		return nil, errors.New("item was not found")
	}

	return &Item{
		itemBase: &itemBase{
			productID:   ci.GetProductID(),
			name:        ci.GetName(),
			description: ci.GetDescription(),
			price:       ci.GetPrice(),
		},
		id: itemID,
	}, nil
}

func (c *cart) Delete(itemID string) (*Item, error) {
	ci, present := c.findCartItem(itemID)
	if !present {
		return nil, errors.New("item was not found")
	}

	ci.Remove(itemID)

	if ci.GetQuantity() == 0 {
		delete(c.items, ci.GetProductID())
	}

	c.deleteUpdate(ci.GetPrice())

	return &Item{
		id: itemID,
		itemBase: &itemBase{
			productID:   ci.GetProductID(),
			name:        ci.GetName(),
			description: ci.GetDescription(),
			price:       ci.GetPrice(),
		},
	}, nil
}

func (c *cart) findCartItem(itemID string) (ItemWrapper, bool) {
	for _, ci := range c.items {
		if ci.Exists(itemID) {
			return ci, true
		}
	}

	return nil, false
}

func (c *cart) putUpdate(amount float32) {
	c.totalPrice += amount
	c.quantity++
}

func (c *cart) deleteUpdate(amount float32) {
	c.totalPrice -= amount
	c.quantity--
}

func (c *cart) addNewItem(i *Item) {
	nci := newItemWrapper(i)

	c.items[nci.GetProductID()] = nci
}

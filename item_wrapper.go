package cart

func newItemWrapper(i *Item) ItemWrapper {
	return &cartItem{
		itemBase: &itemBase{
			price:       i.price,
			name:        i.name,
			description: i.description,
			productID:   i.productID,
		},
		itemMap:  itemMap{i.id: i.version},
		quantity: 1,
	}
}

// ItemWrapper describes an individual cart item in the cart
type ItemWrapper interface {
	Add(i *Item)
	Remove(itemID string)
	Exists(itemID string) bool
	GetProductID() string
	GetName() string
	GetDescription() string
	GetPrice() float32
	GetQuantity() int
}

type cartItem struct {
	*itemBase
	quantity int
	itemMap  itemMap
}

type itemMap map[string]*Version

func (ci *cartItem) Add(i *Item) {
	if _, present := ci.itemMap[i.id]; !present {
		ci.itemMap[i.id] = i.version
		ci.quantity++
	}
}

func (ci *cartItem) Remove(itemID string) {
	if _, present := ci.itemMap[itemID]; present {
		delete(ci.itemMap, itemID)
		ci.quantity--
	}
}

func (ci *cartItem) Exists(itemID string) bool {
	_, exists := ci.itemMap[itemID]
	return exists
}

func (ci *cartItem) GetProductID() string {
	return ci.productID
}

func (ci *cartItem) GetName() string {
	return ci.name
}

func (ci *cartItem) GetDescription() string {
	return ci.description
}

func (ci *cartItem) GetPrice() float32 {
	return ci.price
}

func (ci *cartItem) GetQuantity() int {
	return ci.quantity
}

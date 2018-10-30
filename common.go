package cart

type itemBase struct {
	productID   string
	price       float32
	name        string
	description string
	version     *Version
}

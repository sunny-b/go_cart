package cart

// CRDT describes the interface of the CRDT data structure
type CRDT interface {
	LocalSync(so SyncOperation) error
	RemoteSync(so SyncOperation) error
}

type crdt struct {
	id         string
	cart       Cart
	versionMap VersionMap
}

// Version is the version of a crdt. Contains crdt id and version number
type Version struct {
	id     string
	number int
}

func (c *crdt) LocalSync(so SyncOperation) error {
	so.AddVersion(c.localVersion())

	return c.sync(so)
}

func (c *crdt) RemoteSync(so SyncOperation) error {
	return c.sync(so)
}

func (c *crdt) sync(so SyncOperation) error {
	if so.IsPut() {
		c.cart.Put(so.GetItem())
	} else if so.IsDelete() {
		c.cart.Delete(so.GetItem().id)
	}

	c.updateVersion(so.GetVersion())

	return nil
}

func (c *crdt) updateVersion(v *Version) {
	c.versionMap.Set(v)
}

func (c *crdt) localVersion() *Version {
	return c.versionMap.Get(c.id)
}

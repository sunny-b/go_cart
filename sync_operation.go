package cart

// SyncOperation describes the interface for operations that sync crdts together
type SyncOperation interface {
	AddVersion(v *Version)
	GetVersion() *Version
	GetItem() *Item
	IsPut() bool
	IsDelete() bool
}

type syncOperation struct {
	opType  string
	version *Version
	item    *Item
}

// AddVersion adds a version to the sync operation struct
func (so *syncOperation) AddVersion(v *Version) {
	so.version = v
	so.item.version = v
}

// GetVersion returns the version property of the SyncOperation
func (so *syncOperation) GetVersion() *Version {
	return so.version
}

// GetItem returns the item property of the SyncOperation
func (so *syncOperation) GetItem() *Item {
	return so.item
}

// IsPut denotes if operation is a put operation
func (so *syncOperation) IsPut() bool {
	return so.opType == "put"
}

// IsDelete denotes if operation is a delete operation
func (so *syncOperation) IsDelete() bool {
	return so.opType == "delete"
}

package cart

// VersionMap is the interface that describes the map of version for crdt
type VersionMap interface {
	Get(id string) *Version
	Set(v *Version)
}

type versionMap struct {
	versions map[string]*Version
}

// Set sets the version for corresponding crdt in the map
func (vm *versionMap) Set(v *Version) {
	vm.versions[v.id] = v
}

func (vm *versionMap) Get(id string) *Version {
	return vm.versions[id]
}

package proto

type SystemTreeSync struct {
	Added   []string // UUID of added instances
	Deleted []string // UUID of deleted instances
	Updated []string // UUID of updated instances
	Version uint     // Version of instance tree
	Tree    Instance // Full instance tree
}

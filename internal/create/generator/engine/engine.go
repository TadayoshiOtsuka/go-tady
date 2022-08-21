package engine

type IEngine interface {
	Start(src, rootName, packageName string) error
}

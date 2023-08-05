package mapx

type MultiMap[K any, V any] struct {
	m mapi[K, []V]
}

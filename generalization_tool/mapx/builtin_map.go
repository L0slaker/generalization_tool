package mapx

// builtinMap 是对 map 的二次封装
// 主要用于各种装饰器模式中被装饰的那个
type builtinMap[K comparable, V any] struct {
	data map[K]V
}

func (b *builtinMap[K, V]) Put(key K, val V) error {
	b.data[key] = val
	return nil
}

func (b *builtinMap[K, V]) Get(key K) (V, bool) {
	val, ok := b.data[key]
	return val, ok
}

func (b *builtinMap[K, V]) Delete(k K) (V, bool) {
	v, ok := b.data[k]
	delete(b.data, k)
	return v, ok
}

// Keys 返回的 key 是随机的。即便对于同一个实例，调用两次，得到的结果都可能不同。
func (b *builtinMap[K, V]) Keys() []K {
	return Keys[K, V](b.data)
}

func (b *builtinMap[K, V]) Values() []V {
	return Values[K, V](b.data)
}

func newBuiltinMap[K comparable, V any](cap int) *builtinMap[K, V] {
	return &builtinMap[K, V]{
		data: make(map[K]V, cap),
	}
}

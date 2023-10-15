package jsmap

type JSMap[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any]() *JSMap[K, V] {
	return &JSMap[K, V]{
		m: make(map[K]V),
	}
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/clear
func (mp *JSMap[K, V]) Clear() {
	for k := range mp.m {
		delete(mp.m, k)
	}
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/delete
func (mp *JSMap[K, V]) Delete(k K) {
	delete(mp.m, k)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/get
func (mp *JSMap[K, V]) Get(k K) (V, bool) {
	v, ok := mp.m[k]
	return v, ok
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/set
func (mp *JSMap[K, V]) Set(k K, v V) {
	mp.m[k] = v
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/forEach
func (mp *JSMap[K, V]) Foreach(f func(k K, v V)) {
	for k, v := range mp.m {
		f(k, v)
	}
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/has
func (mp *JSMap[K, V]) Has(k K) bool {
	_, ok := mp.m[k]
	return ok
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/keys
func (mp *JSMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(mp.m))
	for k := range mp.m {
		keys = append(keys, k)
	}
	return keys
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Map/values
func (mp *JSMap[K, V]) Values() []V {
	values := make([]V, 0, len(mp.m))
	for _, v := range mp.m {
		values = append(values, v)
	}
	return values
}

func (mp *JSMap[K, V]) Size() int {
	return len(mp.m)
}

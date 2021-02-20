package cache

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/sets"
)

// Indexer 使用多个索引扩展了 Store，并限制了每个累加器只能容纳当前对象
//
// 这里有3种字符串需要说明：
// 1. 一个存储键，在 Store 接口中定义（其实就是对象键）
// 2. 一个索引的名称（相当于索引分类名称）
// 3. 索引键，由 IndexFunc 生成，可以是一个字段值或从对象中计算出来的任何字符串
type Indexer interface {
	Store // 继承了 Store 存储接口，所以说 Indexer 也是存储
	// indexName 是索引类名称，obj 是对象，计算 obj 在 indexName 索引类中的索引键，然后通过索引键把所有的对象取出来
	// 获取 obj 对象在索引类中的索引键相匹配的对象
	Index(indexName string, obj interface{}) ([]interface{}, error)
	// indexKey 是 indexName 索引分类中的一个索引键
	// 函数返回 indexKey 指定的所有对象键 IndexKeys
	IndexKeys(indexName, indexedValue string) ([]string, error)
	// ListIndexFuncValues 返回所有的索引值给定的索引
	ListIndexFuncValues(indexName string) []string
	// ByIndex returns the stored objects whose set of indexed values
	// for the named index includes the given indexed value
	ByIndex(indexName, indexedValue string) ([]interface{}, error)
	// GetIndexer return the indexers
	GetIndexers() Indexers

	// AddIndexers adds more indexers to this store.  If you call this after you already have data
	// in the store, the results are undefined.
	AddIndexers(newIndexers Indexers) error
}

// IndexFunc 知道怎么计算一个对象的索引键集合
type IndexFunc func(obj interface{}) ([]string, error)

// IndexFuncToKeyFuncAdapter adapts an indexFunc to a keyFunc.  This is only useful if your index function returns
// unique values for every object.  This conversion can create errors when more than one key is found.  You
// should prefer to make proper key and index functions.
func IndexFuncToKeyFuncAdapter(indexFunc IndexFunc) KeyFunc {
	return func(obj interface{}) (string, error) {
		indexKeys, err := indexFunc(obj)
		if err != nil {
			return "", err
		}
		if len(indexKeys) > 1 {
			return "", fmt.Errorf("too many keys: %v", indexKeys)
		}
		if len(indexKeys) == 0 {
			return "", fmt.Errorf("unexpected empty indexKeys")
		}
		return indexKeys[0], nil
	}
}

const (
	// NamespaceIndex is the lookup name for the most comment index function, which is to index by the namespace field.
	NamespaceIndex string = "namespace"
)

// MetaNamespaceIndexFunc is a default index function that indexes based on an object's namespace
func MetaNamespaceIndexFunc(obj interface{}) ([]string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return []string{""}, fmt.Errorf("object has no meta: %v", err)
	}
	return []string{meta.GetNamespace()}, nil
}

// Index maps the indexed value to a set of keys in the store that match on that value
type Index map[string]sets.String

// Indexers maps a name to a IndexFunc
type Indexers map[string]IndexFunc

// Indices maps a name to an Index
type Indices map[string]Index

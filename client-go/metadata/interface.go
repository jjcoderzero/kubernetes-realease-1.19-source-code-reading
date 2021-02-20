package metadata

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

// Interface 允许调用者从任何Kubernetes兼容的资源API中获得元数据(以部分对象的形式)。
type Interface interface {
	Resource(resource schema.GroupVersionResource) Getter
}

// ResourceInterface包含一组可以通过对象的元数据在对象上调用的方法。服务器不支持更新，但是补丁可以用于更新将处理的动作。
type ResourceInterface interface {
	Delete(ctx context.Context, name string, options metav1.DeleteOptions, subresources ...string) error
	DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(ctx context.Context, name string, options metav1.GetOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
	List(ctx context.Context, opts metav1.ListOptions) (*metav1.PartialObjectMetadataList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
}

// Getter handles both namespaced and non-namespaced resource types consistently.
type Getter interface {
	Namespace(string) ResourceInterface
	ResourceInterface
}

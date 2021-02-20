package metadatalister

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// Lister helps list resources.
type Lister interface {
	// List lists all resources in the indexer.
	List(selector labels.Selector) (ret []*metav1.PartialObjectMetadata, err error)
	// Get retrieves a resource from the indexer with the given name
	Get(name string) (*metav1.PartialObjectMetadata, error)
	// Namespace returns an object that can list and get resources in a given namespace.
	Namespace(namespace string) NamespaceLister
}

// NamespaceLister helps list and get resources.
type NamespaceLister interface {
	// List lists all resources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*metav1.PartialObjectMetadata, err error)
	// Get retrieves a resource from the indexer for a given namespace and name.
	Get(name string) (*metav1.PartialObjectMetadata, error)
}

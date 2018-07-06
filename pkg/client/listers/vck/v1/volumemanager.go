/*
<insert-license-here>
*/
// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/ppkube/vck/pkg/apis/vck/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeManagerLister helps list VolumeManagers.
type VolumeManagerLister interface {
	// List lists all VolumeManagers in the indexer.
	List(selector labels.Selector) (ret []*v1.VolumeManager, err error)
	// VolumeManagers returns an object that can list and get VolumeManagers.
	VolumeManagers(namespace string) VolumeManagerNamespaceLister
	VolumeManagerListerExpansion
}

// volumeManagerLister implements the VolumeManagerLister interface.
type volumeManagerLister struct {
	indexer cache.Indexer
}

// NewVolumeManagerLister returns a new VolumeManagerLister.
func NewVolumeManagerLister(indexer cache.Indexer) VolumeManagerLister {
	return &volumeManagerLister{indexer: indexer}
}

// List lists all VolumeManagers in the indexer.
func (s *volumeManagerLister) List(selector labels.Selector) (ret []*v1.VolumeManager, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VolumeManager))
	})
	return ret, err
}

// VolumeManagers returns an object that can list and get VolumeManagers.
func (s *volumeManagerLister) VolumeManagers(namespace string) VolumeManagerNamespaceLister {
	return volumeManagerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeManagerNamespaceLister helps list and get VolumeManagers.
type VolumeManagerNamespaceLister interface {
	// List lists all VolumeManagers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.VolumeManager, err error)
	// Get retrieves the VolumeManager from the indexer for a given namespace and name.
	Get(name string) (*v1.VolumeManager, error)
	VolumeManagerNamespaceListerExpansion
}

// volumeManagerNamespaceLister implements the VolumeManagerNamespaceLister
// interface.
type volumeManagerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeManagers in the indexer for a given namespace.
func (s volumeManagerNamespaceLister) List(selector labels.Selector) (ret []*v1.VolumeManager, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VolumeManager))
	})
	return ret, err
}

// Get retrieves the VolumeManager from the indexer for a given namespace and name.
func (s volumeManagerNamespaceLister) Get(name string) (*v1.VolumeManager, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("volumemanager"), name)
	}
	return obj.(*v1.VolumeManager), nil
}

/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumegroupsnapshot/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeGroupSnapshotContentLister helps list VolumeGroupSnapshotContents.
// All objects returned here must be treated as read-only.
type VolumeGroupSnapshotContentLister interface {
	// List lists all VolumeGroupSnapshotContents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeGroupSnapshotContent, err error)
	// Get retrieves the VolumeGroupSnapshotContent from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.VolumeGroupSnapshotContent, error)
	VolumeGroupSnapshotContentListerExpansion
}

// volumeGroupSnapshotContentLister implements the VolumeGroupSnapshotContentLister interface.
type volumeGroupSnapshotContentLister struct {
	indexer cache.Indexer
}

// NewVolumeGroupSnapshotContentLister returns a new VolumeGroupSnapshotContentLister.
func NewVolumeGroupSnapshotContentLister(indexer cache.Indexer) VolumeGroupSnapshotContentLister {
	return &volumeGroupSnapshotContentLister{indexer: indexer}
}

// List lists all VolumeGroupSnapshotContents in the indexer.
func (s *volumeGroupSnapshotContentLister) List(selector labels.Selector) (ret []*v1beta1.VolumeGroupSnapshotContent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeGroupSnapshotContent))
	})
	return ret, err
}

// Get retrieves the VolumeGroupSnapshotContent from the index for a given name.
func (s *volumeGroupSnapshotContentLister) Get(name string) (*v1beta1.VolumeGroupSnapshotContent, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumegroupsnapshotcontent"), name)
	}
	return obj.(*v1beta1.VolumeGroupSnapshotContent), nil
}
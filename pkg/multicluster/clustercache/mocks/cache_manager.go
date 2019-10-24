// Code generated by MockGen. DO NOT EDIT.
// Source: cache_manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clustercache "github.com/solo-io/solo-kit/pkg/multicluster/clustercache"
	rest "k8s.io/client-go/rest"
)

// MockClusterCache is a mock of ClusterCache interface
type MockClusterCache struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCacheMockRecorder
}

// MockClusterCacheMockRecorder is the mock recorder for MockClusterCache
type MockClusterCacheMockRecorder struct {
	mock *MockClusterCache
}

// NewMockClusterCache creates a new mock instance
func NewMockClusterCache(ctrl *gomock.Controller) *MockClusterCache {
	mock := &MockClusterCache{ctrl: ctrl}
	mock.recorder = &MockClusterCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterCache) EXPECT() *MockClusterCacheMockRecorder {
	return m.recorder
}

// IsClusterCache mocks base method
func (m *MockClusterCache) IsClusterCache() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsClusterCache")
}

// IsClusterCache indicates an expected call of IsClusterCache
func (mr *MockClusterCacheMockRecorder) IsClusterCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterCache", reflect.TypeOf((*MockClusterCache)(nil).IsClusterCache))
}

// MockCacheGetter is a mock of CacheGetter interface
type MockCacheGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCacheGetterMockRecorder
}

// MockCacheGetterMockRecorder is the mock recorder for MockCacheGetter
type MockCacheGetterMockRecorder struct {
	mock *MockCacheGetter
}

// NewMockCacheGetter creates a new mock instance
func NewMockCacheGetter(ctrl *gomock.Controller) *MockCacheGetter {
	mock := &MockCacheGetter{ctrl: ctrl}
	mock.recorder = &MockCacheGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheGetter) EXPECT() *MockCacheGetterMockRecorder {
	return m.recorder
}

// GetCache mocks base method
func (m *MockCacheGetter) GetCache(cluster string, restConfig *rest.Config) clustercache.ClusterCache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", cluster, restConfig)
	ret0, _ := ret[0].(clustercache.ClusterCache)
	return ret0
}

// GetCache indicates an expected call of GetCache
func (mr *MockCacheGetterMockRecorder) GetCache(cluster, restConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockCacheGetter)(nil).GetCache), cluster, restConfig)
}

// MockCacheManager is a mock of CacheManager interface
type MockCacheManager struct {
	ctrl     *gomock.Controller
	recorder *MockCacheManagerMockRecorder
}

// MockCacheManagerMockRecorder is the mock recorder for MockCacheManager
type MockCacheManagerMockRecorder struct {
	mock *MockCacheManager
}

// NewMockCacheManager creates a new mock instance
func NewMockCacheManager(ctrl *gomock.Controller) *MockCacheManager {
	mock := &MockCacheManager{ctrl: ctrl}
	mock.recorder = &MockCacheManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheManager) EXPECT() *MockCacheManagerMockRecorder {
	return m.recorder
}

// ClusterAdded mocks base method
func (m *MockCacheManager) ClusterAdded(cluster string, restConfig *rest.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterAdded", cluster, restConfig)
}

// ClusterAdded indicates an expected call of ClusterAdded
func (mr *MockCacheManagerMockRecorder) ClusterAdded(cluster, restConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAdded", reflect.TypeOf((*MockCacheManager)(nil).ClusterAdded), cluster, restConfig)
}

// ClusterRemoved mocks base method
func (m *MockCacheManager) ClusterRemoved(cluster string, restConfig *rest.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterRemoved", cluster, restConfig)
}

// ClusterRemoved indicates an expected call of ClusterRemoved
func (mr *MockCacheManagerMockRecorder) ClusterRemoved(cluster, restConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRemoved", reflect.TypeOf((*MockCacheManager)(nil).ClusterRemoved), cluster, restConfig)
}

// GetCache mocks base method
func (m *MockCacheManager) GetCache(cluster string, restConfig *rest.Config) clustercache.ClusterCache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", cluster, restConfig)
	ret0, _ := ret[0].(clustercache.ClusterCache)
	return ret0
}

// GetCache indicates an expected call of GetCache
func (mr *MockCacheManagerMockRecorder) GetCache(cluster, restConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockCacheManager)(nil).GetCache), cluster, restConfig)
}
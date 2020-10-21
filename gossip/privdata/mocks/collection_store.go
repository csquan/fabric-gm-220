// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ledger "github.com/VoneChain-CS/fabric-gm/core/ledger"
import mock "github.com/stretchr/testify/mock"
import peer "github.com/hyperledger/fabric-protos-go/peer"
import privdata "github.com/VoneChain-CS/fabric-gm/core/common/privdata"

// CollectionStore is an autogenerated mock type for the CollectionStore type
type CollectionStore struct {
	mock.Mock
}

// AccessFilter provides a mock function with given fields: channelName, collectionPolicyConfig
func (_m *CollectionStore) AccessFilter(channelName string, collectionPolicyConfig *peer.CollectionPolicyConfig) (privdata.Filter, error) {
	ret := _m.Called(channelName, collectionPolicyConfig)

	var r0 privdata.Filter
	if rf, ok := ret.Get(0).(func(string, *peer.CollectionPolicyConfig) privdata.Filter); ok {
		r0 = rf(channelName, collectionPolicyConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *peer.CollectionPolicyConfig) error); ok {
		r1 = rf(channelName, collectionPolicyConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollection provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollection(_a0 privdata.CollectionCriteria) (privdata.Collection, error) {
	ret := _m.Called(_a0)

	var r0 privdata.Collection
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria) privdata.Collection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionAccessPolicy provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionAccessPolicy(_a0 privdata.CollectionCriteria) (privdata.CollectionAccessPolicy, error) {
	ret := _m.Called(_a0)

	var r0 privdata.CollectionAccessPolicy
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria) privdata.CollectionAccessPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.CollectionAccessPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionConfig provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionConfig(_a0 privdata.CollectionCriteria) (*peer.StaticCollectionConfig, error) {
	ret := _m.Called(_a0)

	var r0 *peer.StaticCollectionConfig
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria) *peer.StaticCollectionConfig); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*peer.StaticCollectionConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionConfigPackage provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionConfigPackage(_a0 privdata.CollectionCriteria) (*peer.CollectionConfigPackage, error) {
	ret := _m.Called(_a0)

	var r0 *peer.CollectionConfigPackage
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria) *peer.CollectionConfigPackage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*peer.CollectionConfigPackage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveCollectionPersistenceConfigs provides a mock function with given fields: _a0
func (_m *CollectionStore) RetrieveCollectionPersistenceConfigs(_a0 privdata.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error) {
	ret := _m.Called(_a0)

	var r0 privdata.CollectionPersistenceConfigs
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria) privdata.CollectionPersistenceConfigs); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(privdata.CollectionPersistenceConfigs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveReadWritePermission provides a mock function with given fields: _a0, _a1, _a2
func (_m *CollectionStore) RetrieveReadWritePermission(_a0 privdata.CollectionCriteria, _a1 *peer.SignedProposal, _a2 ledger.QueryExecutor) (bool, bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) bool); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

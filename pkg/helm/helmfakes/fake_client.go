// Code generated by counterfeiter. DO NOT EDIT.
package helmfakes

import (
	"sync"

	"github.com/billiford/go-clouddriver/pkg/helm"
)

type FakeClient struct {
	GetIndexStub        func() (helm.Index, error)
	getIndexMutex       sync.RWMutex
	getIndexArgsForCall []struct{}
	getIndexReturns     struct {
		result1 helm.Index
		result2 error
	}
	getIndexReturnsOnCall map[int]struct {
		result1 helm.Index
		result2 error
	}
	GetChartStub        func(string, string) ([]byte, error)
	getChartMutex       sync.RWMutex
	getChartArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getChartReturns struct {
		result1 []byte
		result2 error
	}
	getChartReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) GetIndex() (helm.Index, error) {
	fake.getIndexMutex.Lock()
	ret, specificReturn := fake.getIndexReturnsOnCall[len(fake.getIndexArgsForCall)]
	fake.getIndexArgsForCall = append(fake.getIndexArgsForCall, struct{}{})
	fake.recordInvocation("GetIndex", []interface{}{})
	fake.getIndexMutex.Unlock()
	if fake.GetIndexStub != nil {
		return fake.GetIndexStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getIndexReturns.result1, fake.getIndexReturns.result2
}

func (fake *FakeClient) GetIndexCallCount() int {
	fake.getIndexMutex.RLock()
	defer fake.getIndexMutex.RUnlock()
	return len(fake.getIndexArgsForCall)
}

func (fake *FakeClient) GetIndexReturns(result1 helm.Index, result2 error) {
	fake.GetIndexStub = nil
	fake.getIndexReturns = struct {
		result1 helm.Index
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetIndexReturnsOnCall(i int, result1 helm.Index, result2 error) {
	fake.GetIndexStub = nil
	if fake.getIndexReturnsOnCall == nil {
		fake.getIndexReturnsOnCall = make(map[int]struct {
			result1 helm.Index
			result2 error
		})
	}
	fake.getIndexReturnsOnCall[i] = struct {
		result1 helm.Index
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetChart(arg1 string, arg2 string) ([]byte, error) {
	fake.getChartMutex.Lock()
	ret, specificReturn := fake.getChartReturnsOnCall[len(fake.getChartArgsForCall)]
	fake.getChartArgsForCall = append(fake.getChartArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChart", []interface{}{arg1, arg2})
	fake.getChartMutex.Unlock()
	if fake.GetChartStub != nil {
		return fake.GetChartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getChartReturns.result1, fake.getChartReturns.result2
}

func (fake *FakeClient) GetChartCallCount() int {
	fake.getChartMutex.RLock()
	defer fake.getChartMutex.RUnlock()
	return len(fake.getChartArgsForCall)
}

func (fake *FakeClient) GetChartArgsForCall(i int) (string, string) {
	fake.getChartMutex.RLock()
	defer fake.getChartMutex.RUnlock()
	return fake.getChartArgsForCall[i].arg1, fake.getChartArgsForCall[i].arg2
}

func (fake *FakeClient) GetChartReturns(result1 []byte, result2 error) {
	fake.GetChartStub = nil
	fake.getChartReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetChartReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetChartStub = nil
	if fake.getChartReturnsOnCall == nil {
		fake.getChartReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getChartReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIndexMutex.RLock()
	defer fake.getIndexMutex.RUnlock()
	fake.getChartMutex.RLock()
	defer fake.getChartMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ helm.Client = new(FakeClient)
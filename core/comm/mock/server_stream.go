
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	context "context"
	sync "sync"

	metadata "google.golang.org/grpc/metadata"
)

type ServerStream struct {
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct {
	}
	contextReturns struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	RecvMsgStub        func(interface{}) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		arg1 interface{}
	}
	recvMsgReturns struct {
		result1 error
	}
	recvMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SendHeaderStub        func(metadata.MD) error
	sendHeaderMutex       sync.RWMutex
	sendHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	sendHeaderReturns struct {
		result1 error
	}
	sendHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	SendMsgStub        func(interface{}) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		arg1 interface{}
	}
	sendMsgReturns struct {
		result1 error
	}
	sendMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SetHeaderStub        func(metadata.MD) error
	setHeaderMutex       sync.RWMutex
	setHeaderArgsForCall []struct {
		arg1 metadata.MD
	}
	setHeaderReturns struct {
		result1 error
	}
	setHeaderReturnsOnCall map[int]struct {
		result1 error
	}
	SetTrailerStub        func(metadata.MD)
	setTrailerMutex       sync.RWMutex
	setTrailerArgsForCall []struct {
		arg1 metadata.MD
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ServerStream) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct {
	}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contextReturns
	return fakeReturns.result1
}

func (fake *ServerStream) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *ServerStream) ContextCalls(stub func() context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = stub
}

func (fake *ServerStream) ContextReturns(result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *ServerStream) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *ServerStream) RecvMsg(arg1 interface{}) error {
	fake.recvMsgMutex.Lock()
	ret, specificReturn := fake.recvMsgReturnsOnCall[len(fake.recvMsgArgsForCall)]
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("RecvMsg", []interface{}{arg1})
	fake.recvMsgMutex.Unlock()
	if fake.RecvMsgStub != nil {
		return fake.RecvMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.recvMsgReturns
	return fakeReturns.result1
}

func (fake *ServerStream) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *ServerStream) RecvMsgCalls(stub func(interface{}) error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = stub
}

func (fake *ServerStream) RecvMsgArgsForCall(i int) interface{} {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	argsForCall := fake.recvMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServerStream) RecvMsgReturns(result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) RecvMsgReturnsOnCall(i int, result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	if fake.recvMsgReturnsOnCall == nil {
		fake.recvMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recvMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SendHeader(arg1 metadata.MD) error {
	fake.sendHeaderMutex.Lock()
	ret, specificReturn := fake.sendHeaderReturnsOnCall[len(fake.sendHeaderArgsForCall)]
	fake.sendHeaderArgsForCall = append(fake.sendHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SendHeader", []interface{}{arg1})
	fake.sendHeaderMutex.Unlock()
	if fake.SendHeaderStub != nil {
		return fake.SendHeaderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendHeaderReturns
	return fakeReturns.result1
}

func (fake *ServerStream) SendHeaderCallCount() int {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	return len(fake.sendHeaderArgsForCall)
}

func (fake *ServerStream) SendHeaderCalls(stub func(metadata.MD) error) {
	fake.sendHeaderMutex.Lock()
	defer fake.sendHeaderMutex.Unlock()
	fake.SendHeaderStub = stub
}

func (fake *ServerStream) SendHeaderArgsForCall(i int) metadata.MD {
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	argsForCall := fake.sendHeaderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServerStream) SendHeaderReturns(result1 error) {
	fake.sendHeaderMutex.Lock()
	defer fake.sendHeaderMutex.Unlock()
	fake.SendHeaderStub = nil
	fake.sendHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SendHeaderReturnsOnCall(i int, result1 error) {
	fake.sendHeaderMutex.Lock()
	defer fake.sendHeaderMutex.Unlock()
	fake.SendHeaderStub = nil
	if fake.sendHeaderReturnsOnCall == nil {
		fake.sendHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SendMsg(arg1 interface{}) error {
	fake.sendMsgMutex.Lock()
	ret, specificReturn := fake.sendMsgReturnsOnCall[len(fake.sendMsgArgsForCall)]
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("SendMsg", []interface{}{arg1})
	fake.sendMsgMutex.Unlock()
	if fake.SendMsgStub != nil {
		return fake.SendMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendMsgReturns
	return fakeReturns.result1
}

func (fake *ServerStream) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *ServerStream) SendMsgCalls(stub func(interface{}) error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = stub
}

func (fake *ServerStream) SendMsgArgsForCall(i int) interface{} {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	argsForCall := fake.sendMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServerStream) SendMsgReturns(result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SendMsgReturnsOnCall(i int, result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	if fake.sendMsgReturnsOnCall == nil {
		fake.sendMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SetHeader(arg1 metadata.MD) error {
	fake.setHeaderMutex.Lock()
	ret, specificReturn := fake.setHeaderReturnsOnCall[len(fake.setHeaderArgsForCall)]
	fake.setHeaderArgsForCall = append(fake.setHeaderArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetHeader", []interface{}{arg1})
	fake.setHeaderMutex.Unlock()
	if fake.SetHeaderStub != nil {
		return fake.SetHeaderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setHeaderReturns
	return fakeReturns.result1
}

func (fake *ServerStream) SetHeaderCallCount() int {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	return len(fake.setHeaderArgsForCall)
}

func (fake *ServerStream) SetHeaderCalls(stub func(metadata.MD) error) {
	fake.setHeaderMutex.Lock()
	defer fake.setHeaderMutex.Unlock()
	fake.SetHeaderStub = stub
}

func (fake *ServerStream) SetHeaderArgsForCall(i int) metadata.MD {
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	argsForCall := fake.setHeaderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServerStream) SetHeaderReturns(result1 error) {
	fake.setHeaderMutex.Lock()
	defer fake.setHeaderMutex.Unlock()
	fake.SetHeaderStub = nil
	fake.setHeaderReturns = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SetHeaderReturnsOnCall(i int, result1 error) {
	fake.setHeaderMutex.Lock()
	defer fake.setHeaderMutex.Unlock()
	fake.SetHeaderStub = nil
	if fake.setHeaderReturnsOnCall == nil {
		fake.setHeaderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setHeaderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ServerStream) SetTrailer(arg1 metadata.MD) {
	fake.setTrailerMutex.Lock()
	fake.setTrailerArgsForCall = append(fake.setTrailerArgsForCall, struct {
		arg1 metadata.MD
	}{arg1})
	fake.recordInvocation("SetTrailer", []interface{}{arg1})
	fake.setTrailerMutex.Unlock()
	if fake.SetTrailerStub != nil {
		fake.SetTrailerStub(arg1)
	}
}

func (fake *ServerStream) SetTrailerCallCount() int {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	return len(fake.setTrailerArgsForCall)
}

func (fake *ServerStream) SetTrailerCalls(stub func(metadata.MD)) {
	fake.setTrailerMutex.Lock()
	defer fake.setTrailerMutex.Unlock()
	fake.SetTrailerStub = stub
}

func (fake *ServerStream) SetTrailerArgsForCall(i int) metadata.MD {
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	argsForCall := fake.setTrailerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServerStream) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	fake.sendHeaderMutex.RLock()
	defer fake.sendHeaderMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.setHeaderMutex.RLock()
	defer fake.setHeaderMutex.RUnlock()
	fake.setTrailerMutex.RLock()
	defer fake.setTrailerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ServerStream) recordInvocation(key string, args []interface{}) {
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

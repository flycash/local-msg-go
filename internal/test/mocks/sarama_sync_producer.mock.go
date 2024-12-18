// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/IBM/sarama (interfaces: SyncProducer)
//
// Generated by this command:
//
//	mockgen -destination=internal/test/mocks/sarama_sync_producer.mock.go -typed -package=mocks github.com/IBM/sarama SyncProducer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	sarama "github.com/IBM/sarama"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncProducer is a mock of SyncProducer interface.
type MockSyncProducer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncProducerMockRecorder
}

// MockSyncProducerMockRecorder is the mock recorder for MockSyncProducer.
type MockSyncProducerMockRecorder struct {
	mock *MockSyncProducer
}

// NewMockSyncProducer creates a new mock instance.
func NewMockSyncProducer(ctrl *gomock.Controller) *MockSyncProducer {
	mock := &MockSyncProducer{ctrl: ctrl}
	mock.recorder = &MockSyncProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncProducer) EXPECT() *MockSyncProducerMockRecorder {
	return m.recorder
}

// AbortTxn mocks base method.
func (m *MockSyncProducer) AbortTxn() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortTxn")
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortTxn indicates an expected call of AbortTxn.
func (mr *MockSyncProducerMockRecorder) AbortTxn() *SyncProducerAbortTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTxn", reflect.TypeOf((*MockSyncProducer)(nil).AbortTxn))
	return &SyncProducerAbortTxnCall{Call: call}
}

// SyncProducerAbortTxnCall wrap *gomock.Call
type SyncProducerAbortTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerAbortTxnCall) Return(arg0 error) *SyncProducerAbortTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerAbortTxnCall) Do(f func() error) *SyncProducerAbortTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerAbortTxnCall) DoAndReturn(f func() error) *SyncProducerAbortTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddMessageToTxn mocks base method.
func (m *MockSyncProducer) AddMessageToTxn(arg0 *sarama.ConsumerMessage, arg1 string, arg2 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessageToTxn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMessageToTxn indicates an expected call of AddMessageToTxn.
func (mr *MockSyncProducerMockRecorder) AddMessageToTxn(arg0, arg1, arg2 any) *SyncProducerAddMessageToTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessageToTxn", reflect.TypeOf((*MockSyncProducer)(nil).AddMessageToTxn), arg0, arg1, arg2)
	return &SyncProducerAddMessageToTxnCall{Call: call}
}

// SyncProducerAddMessageToTxnCall wrap *gomock.Call
type SyncProducerAddMessageToTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerAddMessageToTxnCall) Return(arg0 error) *SyncProducerAddMessageToTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerAddMessageToTxnCall) Do(f func(*sarama.ConsumerMessage, string, *string) error) *SyncProducerAddMessageToTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerAddMessageToTxnCall) DoAndReturn(f func(*sarama.ConsumerMessage, string, *string) error) *SyncProducerAddMessageToTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddOffsetsToTxn mocks base method.
func (m *MockSyncProducer) AddOffsetsToTxn(arg0 map[string][]*sarama.PartitionOffsetMetadata, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOffsetsToTxn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOffsetsToTxn indicates an expected call of AddOffsetsToTxn.
func (mr *MockSyncProducerMockRecorder) AddOffsetsToTxn(arg0, arg1 any) *SyncProducerAddOffsetsToTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOffsetsToTxn", reflect.TypeOf((*MockSyncProducer)(nil).AddOffsetsToTxn), arg0, arg1)
	return &SyncProducerAddOffsetsToTxnCall{Call: call}
}

// SyncProducerAddOffsetsToTxnCall wrap *gomock.Call
type SyncProducerAddOffsetsToTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerAddOffsetsToTxnCall) Return(arg0 error) *SyncProducerAddOffsetsToTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerAddOffsetsToTxnCall) Do(f func(map[string][]*sarama.PartitionOffsetMetadata, string) error) *SyncProducerAddOffsetsToTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerAddOffsetsToTxnCall) DoAndReturn(f func(map[string][]*sarama.PartitionOffsetMetadata, string) error) *SyncProducerAddOffsetsToTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BeginTxn mocks base method.
func (m *MockSyncProducer) BeginTxn() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxn")
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTxn indicates an expected call of BeginTxn.
func (mr *MockSyncProducerMockRecorder) BeginTxn() *SyncProducerBeginTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxn", reflect.TypeOf((*MockSyncProducer)(nil).BeginTxn))
	return &SyncProducerBeginTxnCall{Call: call}
}

// SyncProducerBeginTxnCall wrap *gomock.Call
type SyncProducerBeginTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerBeginTxnCall) Return(arg0 error) *SyncProducerBeginTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerBeginTxnCall) Do(f func() error) *SyncProducerBeginTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerBeginTxnCall) DoAndReturn(f func() error) *SyncProducerBeginTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockSyncProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncProducerMockRecorder) Close() *SyncProducerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncProducer)(nil).Close))
	return &SyncProducerCloseCall{Call: call}
}

// SyncProducerCloseCall wrap *gomock.Call
type SyncProducerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerCloseCall) Return(arg0 error) *SyncProducerCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerCloseCall) Do(f func() error) *SyncProducerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerCloseCall) DoAndReturn(f func() error) *SyncProducerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CommitTxn mocks base method.
func (m *MockSyncProducer) CommitTxn() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTxn")
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTxn indicates an expected call of CommitTxn.
func (mr *MockSyncProducerMockRecorder) CommitTxn() *SyncProducerCommitTxnCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTxn", reflect.TypeOf((*MockSyncProducer)(nil).CommitTxn))
	return &SyncProducerCommitTxnCall{Call: call}
}

// SyncProducerCommitTxnCall wrap *gomock.Call
type SyncProducerCommitTxnCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerCommitTxnCall) Return(arg0 error) *SyncProducerCommitTxnCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerCommitTxnCall) Do(f func() error) *SyncProducerCommitTxnCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerCommitTxnCall) DoAndReturn(f func() error) *SyncProducerCommitTxnCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsTransactional mocks base method.
func (m *MockSyncProducer) IsTransactional() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTransactional")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTransactional indicates an expected call of IsTransactional.
func (mr *MockSyncProducerMockRecorder) IsTransactional() *SyncProducerIsTransactionalCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTransactional", reflect.TypeOf((*MockSyncProducer)(nil).IsTransactional))
	return &SyncProducerIsTransactionalCall{Call: call}
}

// SyncProducerIsTransactionalCall wrap *gomock.Call
type SyncProducerIsTransactionalCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerIsTransactionalCall) Return(arg0 bool) *SyncProducerIsTransactionalCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerIsTransactionalCall) Do(f func() bool) *SyncProducerIsTransactionalCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerIsTransactionalCall) DoAndReturn(f func() bool) *SyncProducerIsTransactionalCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessage mocks base method.
func (m *MockSyncProducer) SendMessage(arg0 *sarama.ProducerMessage) (int32, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockSyncProducerMockRecorder) SendMessage(arg0 any) *SyncProducerSendMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSyncProducer)(nil).SendMessage), arg0)
	return &SyncProducerSendMessageCall{Call: call}
}

// SyncProducerSendMessageCall wrap *gomock.Call
type SyncProducerSendMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerSendMessageCall) Return(arg0 int32, arg1 int64, arg2 error) *SyncProducerSendMessageCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerSendMessageCall) Do(f func(*sarama.ProducerMessage) (int32, int64, error)) *SyncProducerSendMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerSendMessageCall) DoAndReturn(f func(*sarama.ProducerMessage) (int32, int64, error)) *SyncProducerSendMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessages mocks base method.
func (m *MockSyncProducer) SendMessages(arg0 []*sarama.ProducerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessages", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages.
func (mr *MockSyncProducerMockRecorder) SendMessages(arg0 any) *SyncProducerSendMessagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockSyncProducer)(nil).SendMessages), arg0)
	return &SyncProducerSendMessagesCall{Call: call}
}

// SyncProducerSendMessagesCall wrap *gomock.Call
type SyncProducerSendMessagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerSendMessagesCall) Return(arg0 error) *SyncProducerSendMessagesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerSendMessagesCall) Do(f func([]*sarama.ProducerMessage) error) *SyncProducerSendMessagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerSendMessagesCall) DoAndReturn(f func([]*sarama.ProducerMessage) error) *SyncProducerSendMessagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TxnStatus mocks base method.
func (m *MockSyncProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxnStatus")
	ret0, _ := ret[0].(sarama.ProducerTxnStatusFlag)
	return ret0
}

// TxnStatus indicates an expected call of TxnStatus.
func (mr *MockSyncProducerMockRecorder) TxnStatus() *SyncProducerTxnStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxnStatus", reflect.TypeOf((*MockSyncProducer)(nil).TxnStatus))
	return &SyncProducerTxnStatusCall{Call: call}
}

// SyncProducerTxnStatusCall wrap *gomock.Call
type SyncProducerTxnStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SyncProducerTxnStatusCall) Return(arg0 sarama.ProducerTxnStatusFlag) *SyncProducerTxnStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SyncProducerTxnStatusCall) Do(f func() sarama.ProducerTxnStatusFlag) *SyncProducerTxnStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SyncProducerTxnStatusCall) DoAndReturn(f func() sarama.ProducerTxnStatusFlag) *SyncProducerTxnStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

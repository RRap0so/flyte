// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// WorkflowExecutionEventWriter is an autogenerated mock type for the WorkflowExecutionEventWriter type
type WorkflowExecutionEventWriter struct {
	mock.Mock
}

// Run provides a mock function with given fields:
func (_m *WorkflowExecutionEventWriter) Run() {
	_m.Called()
}

// Write provides a mock function with given fields: workflowExecutionEvent
func (_m *WorkflowExecutionEventWriter) Write(workflowExecutionEvent admin.WorkflowExecutionEventRequest) {
	_m.Called(workflowExecutionEvent)
}

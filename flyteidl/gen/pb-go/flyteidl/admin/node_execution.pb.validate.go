// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/node_execution.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = core.NodeExecution_Phase(0)

	_ = core.CatalogCacheStatus(0)
)

// define the regex for a UUID once up-front
var _node_execution_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on NodeExecutionGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionGetRequestValidationError is the validation error returned by
// NodeExecutionGetRequest.Validate if the designated constraints aren't met.
type NodeExecutionGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionGetRequestValidationError) ErrorName() string {
	return "NodeExecutionGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionGetRequestValidationError{}

// Validate checks the field values on NodeExecutionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetWorkflowExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionListRequestValidationError{
				field:  "WorkflowExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Limit

	// no validation rules for Token

	// no validation rules for Filters

	if v, ok := interface{}(m.GetSortBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionListRequestValidationError{
				field:  "SortBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UniqueParentId

	return nil
}

// NodeExecutionListRequestValidationError is the validation error returned by
// NodeExecutionListRequest.Validate if the designated constraints aren't met.
type NodeExecutionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionListRequestValidationError) ErrorName() string {
	return "NodeExecutionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionListRequestValidationError{}

// Validate checks the field values on NodeExecutionForTaskListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionForTaskListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTaskExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionForTaskListRequestValidationError{
				field:  "TaskExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Limit

	// no validation rules for Token

	// no validation rules for Filters

	if v, ok := interface{}(m.GetSortBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionForTaskListRequestValidationError{
				field:  "SortBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionForTaskListRequestValidationError is the validation error
// returned by NodeExecutionForTaskListRequest.Validate if the designated
// constraints aren't met.
type NodeExecutionForTaskListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionForTaskListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionForTaskListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionForTaskListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionForTaskListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionForTaskListRequestValidationError) ErrorName() string {
	return "NodeExecutionForTaskListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionForTaskListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionForTaskListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionForTaskListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionForTaskListRequestValidationError{}

// Validate checks the field values on NodeExecution with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NodeExecution) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InputUri

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionValidationError is the validation error returned by
// NodeExecution.Validate if the designated constraints aren't met.
type NodeExecutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionValidationError) ErrorName() string { return "NodeExecutionValidationError" }

// Error satisfies the builtin error interface
func (e NodeExecutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionValidationError{}

// Validate checks the field values on NodeExecutionMetaData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionMetaData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RetryGroup

	// no validation rules for IsParentNode

	// no validation rules for SpecNodeId

	// no validation rules for IsDynamic

	return nil
}

// NodeExecutionMetaDataValidationError is the validation error returned by
// NodeExecutionMetaData.Validate if the designated constraints aren't met.
type NodeExecutionMetaDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionMetaDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionMetaDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionMetaDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionMetaDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionMetaDataValidationError) ErrorName() string {
	return "NodeExecutionMetaDataValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionMetaDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionMetaData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionMetaDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionMetaDataValidationError{}

// Validate checks the field values on NodeExecutionList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *NodeExecutionList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNodeExecutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionListValidationError{
					field:  fmt.Sprintf("NodeExecutions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// NodeExecutionListValidationError is the validation error returned by
// NodeExecutionList.Validate if the designated constraints aren't met.
type NodeExecutionListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionListValidationError) ErrorName() string {
	return "NodeExecutionListValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionListValidationError{}

// Validate checks the field values on NodeExecutionClosure with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionClosure) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Phase

	if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionClosureValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionClosureValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionClosureValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeckUri

	// no validation rules for DynamicJobSpecUri

	switch m.OutputResult.(type) {

	case *NodeExecutionClosure_OutputUri:
		// no validation rules for OutputUri

	case *NodeExecutionClosure_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionClosureValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *NodeExecutionClosure_OutputData:

		if v, ok := interface{}(m.GetOutputData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionClosureValidationError{
					field:  "OutputData",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.TargetMetadata.(type) {

	case *NodeExecutionClosure_WorkflowNodeMetadata:

		if v, ok := interface{}(m.GetWorkflowNodeMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionClosureValidationError{
					field:  "WorkflowNodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *NodeExecutionClosure_TaskNodeMetadata:

		if v, ok := interface{}(m.GetTaskNodeMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionClosureValidationError{
					field:  "TaskNodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeExecutionClosureValidationError is the validation error returned by
// NodeExecutionClosure.Validate if the designated constraints aren't met.
type NodeExecutionClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionClosureValidationError) ErrorName() string {
	return "NodeExecutionClosureValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionClosureValidationError{}

// Validate checks the field values on WorkflowNodeMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowNodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowNodeMetadataValidationError{
				field:  "ExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowNodeMetadataValidationError is the validation error returned by
// WorkflowNodeMetadata.Validate if the designated constraints aren't met.
type WorkflowNodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowNodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowNodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowNodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowNodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowNodeMetadataValidationError) ErrorName() string {
	return "WorkflowNodeMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowNodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowNodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowNodeMetadataValidationError{}

// Validate checks the field values on TaskNodeMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TaskNodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CacheStatus

	if v, ok := interface{}(m.GetCatalogKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskNodeMetadataValidationError{
				field:  "CatalogKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CheckpointUri

	return nil
}

// TaskNodeMetadataValidationError is the validation error returned by
// TaskNodeMetadata.Validate if the designated constraints aren't met.
type TaskNodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskNodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskNodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskNodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskNodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskNodeMetadataValidationError) ErrorName() string { return "TaskNodeMetadataValidationError" }

// Error satisfies the builtin error interface
func (e TaskNodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskNodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskNodeMetadataValidationError{}

// Validate checks the field values on DynamicWorkflowNodeMetadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DynamicWorkflowNodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamicWorkflowNodeMetadataValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCompiledWorkflow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamicWorkflowNodeMetadataValidationError{
				field:  "CompiledWorkflow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DynamicJobSpecUri

	return nil
}

// DynamicWorkflowNodeMetadataValidationError is the validation error returned
// by DynamicWorkflowNodeMetadata.Validate if the designated constraints
// aren't met.
type DynamicWorkflowNodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamicWorkflowNodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamicWorkflowNodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamicWorkflowNodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamicWorkflowNodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamicWorkflowNodeMetadataValidationError) ErrorName() string {
	return "DynamicWorkflowNodeMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e DynamicWorkflowNodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicWorkflowNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamicWorkflowNodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamicWorkflowNodeMetadataValidationError{}

// Validate checks the field values on NodeExecutionGetDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionGetDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionGetDataRequestValidationError is the validation error returned
// by NodeExecutionGetDataRequest.Validate if the designated constraints
// aren't met.
type NodeExecutionGetDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionGetDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionGetDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionGetDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionGetDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionGetDataRequestValidationError) ErrorName() string {
	return "NodeExecutionGetDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionGetDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionGetDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionGetDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionGetDataRequestValidationError{}

// Validate checks the field values on NodeExecutionGetDataResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionGetDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "Outputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFullInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "FullInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFullOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "FullOutputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDynamicWorkflow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "DynamicWorkflow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFlyteUrls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionGetDataResponseValidationError{
				field:  "FlyteUrls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NodeExecutionGetDataResponseValidationError is the validation error returned
// by NodeExecutionGetDataResponse.Validate if the designated constraints
// aren't met.
type NodeExecutionGetDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionGetDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionGetDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionGetDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionGetDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionGetDataResponseValidationError) ErrorName() string {
	return "NodeExecutionGetDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionGetDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionGetDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionGetDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionGetDataResponseValidationError{}

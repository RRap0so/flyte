// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/core/artifact_id.proto

package core

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
)

// define the regex for a UUID once up-front
var _artifact_id_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ArtifactKey with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ArtifactKey) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	// no validation rules for Name

	return nil
}

// ArtifactKeyValidationError is the validation error returned by
// ArtifactKey.Validate if the designated constraints aren't met.
type ArtifactKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactKeyValidationError) ErrorName() string { return "ArtifactKeyValidationError" }

// Error satisfies the builtin error interface
func (e ArtifactKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifactKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactKeyValidationError{}

// Validate checks the field values on ArtifactBindingData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ArtifactBindingData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Index

	// no validation rules for PartitionKey

	// no validation rules for Transform

	return nil
}

// ArtifactBindingDataValidationError is the validation error returned by
// ArtifactBindingData.Validate if the designated constraints aren't met.
type ArtifactBindingDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactBindingDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactBindingDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactBindingDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactBindingDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactBindingDataValidationError) ErrorName() string {
	return "ArtifactBindingDataValidationError"
}

// Error satisfies the builtin error interface
func (e ArtifactBindingDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifactBindingData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactBindingDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactBindingDataValidationError{}

// Validate checks the field values on InputBindingData with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *InputBindingData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Var

	return nil
}

// InputBindingDataValidationError is the validation error returned by
// InputBindingData.Validate if the designated constraints aren't met.
type InputBindingDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InputBindingDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InputBindingDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InputBindingDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InputBindingDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InputBindingDataValidationError) ErrorName() string { return "InputBindingDataValidationError" }

// Error satisfies the builtin error interface
func (e InputBindingDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInputBindingData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InputBindingDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InputBindingDataValidationError{}

// Validate checks the field values on LabelValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LabelValue) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Value.(type) {

	case *LabelValue_StaticValue:
		// no validation rules for StaticValue

	case *LabelValue_TriggeredBinding:

		if v, ok := interface{}(m.GetTriggeredBinding()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LabelValueValidationError{
					field:  "TriggeredBinding",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LabelValue_InputBinding:

		if v, ok := interface{}(m.GetInputBinding()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LabelValueValidationError{
					field:  "InputBinding",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LabelValueValidationError is the validation error returned by
// LabelValue.Validate if the designated constraints aren't met.
type LabelValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LabelValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LabelValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LabelValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LabelValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LabelValueValidationError) ErrorName() string { return "LabelValueValidationError" }

// Error satisfies the builtin error interface
func (e LabelValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLabelValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LabelValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LabelValueValidationError{}

// Validate checks the field values on Partitions with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Partitions) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetValue() {
		_ = val

		// no validation rules for Value[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PartitionsValidationError{
					field:  fmt.Sprintf("Value[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PartitionsValidationError is the validation error returned by
// Partitions.Validate if the designated constraints aren't met.
type PartitionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PartitionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PartitionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PartitionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PartitionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PartitionsValidationError) ErrorName() string { return "PartitionsValidationError" }

// Error satisfies the builtin error interface
func (e PartitionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPartitions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PartitionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PartitionsValidationError{}

// Validate checks the field values on ArtifactID with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ArtifactID) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetArtifactKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArtifactIDValidationError{
				field:  "ArtifactKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Version

	switch m.Dimensions.(type) {

	case *ArtifactID_Partitions:

		if v, ok := interface{}(m.GetPartitions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArtifactIDValidationError{
					field:  "Partitions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ArtifactIDValidationError is the validation error returned by
// ArtifactID.Validate if the designated constraints aren't met.
type ArtifactIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactIDValidationError) ErrorName() string { return "ArtifactIDValidationError" }

// Error satisfies the builtin error interface
func (e ArtifactIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifactID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactIDValidationError{}

// Validate checks the field values on ArtifactTag with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ArtifactTag) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetArtifactKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArtifactTagValidationError{
				field:  "ArtifactKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArtifactTagValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ArtifactTagValidationError is the validation error returned by
// ArtifactTag.Validate if the designated constraints aren't met.
type ArtifactTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactTagValidationError) ErrorName() string { return "ArtifactTagValidationError" }

// Error satisfies the builtin error interface
func (e ArtifactTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifactTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactTagValidationError{}

// Validate checks the field values on ArtifactQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ArtifactQuery) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Identifier.(type) {

	case *ArtifactQuery_ArtifactId:

		if v, ok := interface{}(m.GetArtifactId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArtifactQueryValidationError{
					field:  "ArtifactId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ArtifactQuery_ArtifactTag:

		if v, ok := interface{}(m.GetArtifactTag()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArtifactQueryValidationError{
					field:  "ArtifactTag",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ArtifactQuery_Uri:
		// no validation rules for Uri

	case *ArtifactQuery_Binding:

		if v, ok := interface{}(m.GetBinding()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArtifactQueryValidationError{
					field:  "Binding",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ArtifactQueryValidationError is the validation error returned by
// ArtifactQuery.Validate if the designated constraints aren't met.
type ArtifactQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArtifactQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArtifactQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArtifactQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArtifactQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArtifactQueryValidationError) ErrorName() string { return "ArtifactQueryValidationError" }

// Error satisfies the builtin error interface
func (e ArtifactQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArtifactQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArtifactQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArtifactQueryValidationError{}

// Validate checks the field values on Trigger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trigger) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTriggerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TriggerValidationError{
				field:  "TriggerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTriggers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TriggerValidationError{
					field:  fmt.Sprintf("Triggers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TriggerValidationError is the validation error returned by Trigger.Validate
// if the designated constraints aren't met.
type TriggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TriggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TriggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TriggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TriggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TriggerValidationError) ErrorName() string { return "TriggerValidationError" }

// Error satisfies the builtin error interface
func (e TriggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrigger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TriggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TriggerValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/notify/template/notify_template.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTemplateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTemplateRequestMultiError, or nil if none found.
func (m *ListTemplateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTemplateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNotifyId() < 1 {
		err := ListTemplateRequestValidationError{
			field:  "NotifyId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListTemplateRequestMultiError(errors)
	}

	return nil
}

// ListTemplateRequestMultiError is an error wrapping multiple validation
// errors returned by ListTemplateRequest.ValidateAll() if the designated
// constraints aren't met.
type ListTemplateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTemplateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTemplateRequestMultiError) AllErrors() []error { return m }

// ListTemplateRequestValidationError is the validation error returned by
// ListTemplateRequest.Validate if the designated constraints aren't met.
type ListTemplateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTemplateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTemplateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTemplateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTemplateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTemplateRequestValidationError) ErrorName() string {
	return "ListTemplateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTemplateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTemplateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTemplateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTemplateRequestValidationError{}

// Validate checks the field values on ListTemplateReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTemplateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTemplateReplyMultiError, or nil if none found.
func (m *ListTemplateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTemplateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTemplateReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTemplateReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTemplateReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListTemplateReplyMultiError(errors)
	}

	return nil
}

// ListTemplateReplyMultiError is an error wrapping multiple validation errors
// returned by ListTemplateReply.ValidateAll() if the designated constraints
// aren't met.
type ListTemplateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTemplateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTemplateReplyMultiError) AllErrors() []error { return m }

// ListTemplateReplyValidationError is the validation error returned by
// ListTemplateReply.Validate if the designated constraints aren't met.
type ListTemplateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTemplateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTemplateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTemplateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTemplateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTemplateReplyValidationError) ErrorName() string {
	return "ListTemplateReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListTemplateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTemplateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTemplateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTemplateReplyValidationError{}

// Validate checks the field values on CreateTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTemplateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTemplateRequestMultiError, or nil if none found.
func (m *CreateTemplateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTemplateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NotifyId

	// no validation rules for ChannelId

	// no validation rules for Content

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if len(errors) > 0 {
		return CreateTemplateRequestMultiError(errors)
	}

	return nil
}

// CreateTemplateRequestMultiError is an error wrapping multiple validation
// errors returned by CreateTemplateRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateTemplateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTemplateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTemplateRequestMultiError) AllErrors() []error { return m }

// CreateTemplateRequestValidationError is the validation error returned by
// CreateTemplateRequest.Validate if the designated constraints aren't met.
type CreateTemplateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTemplateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTemplateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTemplateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTemplateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTemplateRequestValidationError) ErrorName() string {
	return "CreateTemplateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTemplateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTemplateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTemplateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTemplateRequestValidationError{}

// Validate checks the field values on CreateTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTemplateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTemplateReplyMultiError, or nil if none found.
func (m *CreateTemplateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTemplateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateTemplateReplyMultiError(errors)
	}

	return nil
}

// CreateTemplateReplyMultiError is an error wrapping multiple validation
// errors returned by CreateTemplateReply.ValidateAll() if the designated
// constraints aren't met.
type CreateTemplateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTemplateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTemplateReplyMultiError) AllErrors() []error { return m }

// CreateTemplateReplyValidationError is the validation error returned by
// CreateTemplateReply.Validate if the designated constraints aren't met.
type CreateTemplateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTemplateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTemplateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTemplateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTemplateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTemplateReplyValidationError) ErrorName() string {
	return "CreateTemplateReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTemplateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTemplateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTemplateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTemplateReplyValidationError{}

// Validate checks the field values on UpdateTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTemplateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTemplateRequestMultiError, or nil if none found.
func (m *UpdateTemplateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTemplateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := UpdateTemplateRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NotifyId

	// no validation rules for ChannelId

	// no validation rules for Content

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if len(errors) > 0 {
		return UpdateTemplateRequestMultiError(errors)
	}

	return nil
}

// UpdateTemplateRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateTemplateRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateTemplateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTemplateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTemplateRequestMultiError) AllErrors() []error { return m }

// UpdateTemplateRequestValidationError is the validation error returned by
// UpdateTemplateRequest.Validate if the designated constraints aren't met.
type UpdateTemplateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTemplateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTemplateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTemplateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTemplateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTemplateRequestValidationError) ErrorName() string {
	return "UpdateTemplateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTemplateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTemplateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTemplateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTemplateRequestValidationError{}

// Validate checks the field values on UpdateTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTemplateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTemplateReplyMultiError, or nil if none found.
func (m *UpdateTemplateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTemplateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateTemplateReplyMultiError(errors)
	}

	return nil
}

// UpdateTemplateReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateTemplateReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateTemplateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTemplateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTemplateReplyMultiError) AllErrors() []error { return m }

// UpdateTemplateReplyValidationError is the validation error returned by
// UpdateTemplateReply.Validate if the designated constraints aren't met.
type UpdateTemplateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTemplateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTemplateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTemplateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTemplateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTemplateReplyValidationError) ErrorName() string {
	return "UpdateTemplateReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTemplateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTemplateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTemplateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTemplateReplyValidationError{}

// Validate checks the field values on DeleteTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTemplateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTemplateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTemplateRequestMultiError, or nil if none found.
func (m *DeleteTemplateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTemplateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := DeleteTemplateRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTemplateRequestMultiError(errors)
	}

	return nil
}

// DeleteTemplateRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteTemplateRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteTemplateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTemplateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTemplateRequestMultiError) AllErrors() []error { return m }

// DeleteTemplateRequestValidationError is the validation error returned by
// DeleteTemplateRequest.Validate if the designated constraints aren't met.
type DeleteTemplateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTemplateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTemplateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTemplateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTemplateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTemplateRequestValidationError) ErrorName() string {
	return "DeleteTemplateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTemplateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTemplateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTemplateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTemplateRequestValidationError{}

// Validate checks the field values on DeleteTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTemplateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTemplateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTemplateReplyMultiError, or nil if none found.
func (m *DeleteTemplateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTemplateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTemplateReplyMultiError(errors)
	}

	return nil
}

// DeleteTemplateReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteTemplateReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteTemplateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTemplateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTemplateReplyMultiError) AllErrors() []error { return m }

// DeleteTemplateReplyValidationError is the validation error returned by
// DeleteTemplateReply.Validate if the designated constraints aren't met.
type DeleteTemplateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTemplateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTemplateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTemplateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTemplateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTemplateReplyValidationError) ErrorName() string {
	return "DeleteTemplateReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTemplateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTemplateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTemplateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTemplateReplyValidationError{}

// Validate checks the field values on ListTemplateReply_Channel with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTemplateReply_Channel) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTemplateReply_Channel with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTemplateReply_ChannelMultiError, or nil if none found.
func (m *ListTemplateReply_Channel) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTemplateReply_Channel) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	if len(errors) > 0 {
		return ListTemplateReply_ChannelMultiError(errors)
	}

	return nil
}

// ListTemplateReply_ChannelMultiError is an error wrapping multiple validation
// errors returned by ListTemplateReply_Channel.ValidateAll() if the
// designated constraints aren't met.
type ListTemplateReply_ChannelMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTemplateReply_ChannelMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTemplateReply_ChannelMultiError) AllErrors() []error { return m }

// ListTemplateReply_ChannelValidationError is the validation error returned by
// ListTemplateReply_Channel.Validate if the designated constraints aren't met.
type ListTemplateReply_ChannelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTemplateReply_ChannelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTemplateReply_ChannelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTemplateReply_ChannelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTemplateReply_ChannelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTemplateReply_ChannelValidationError) ErrorName() string {
	return "ListTemplateReply_ChannelValidationError"
}

// Error satisfies the builtin error interface
func (e ListTemplateReply_ChannelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTemplateReply_Channel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTemplateReply_ChannelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTemplateReply_ChannelValidationError{}

// Validate checks the field values on ListTemplateReply_Template with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTemplateReply_Template) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTemplateReply_Template with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTemplateReply_TemplateMultiError, or nil if none found.
func (m *ListTemplateReply_Template) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTemplateReply_Template) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NotifyId

	// no validation rules for ChannelId

	// no validation rules for Content

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if all {
		switch v := interface{}(m.GetChannel()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTemplateReply_TemplateValidationError{
					field:  "Channel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTemplateReply_TemplateValidationError{
					field:  "Channel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChannel()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTemplateReply_TemplateValidationError{
				field:  "Channel",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if len(errors) > 0 {
		return ListTemplateReply_TemplateMultiError(errors)
	}

	return nil
}

// ListTemplateReply_TemplateMultiError is an error wrapping multiple
// validation errors returned by ListTemplateReply_Template.ValidateAll() if
// the designated constraints aren't met.
type ListTemplateReply_TemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTemplateReply_TemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTemplateReply_TemplateMultiError) AllErrors() []error { return m }

// ListTemplateReply_TemplateValidationError is the validation error returned
// by ListTemplateReply_Template.Validate if the designated constraints aren't met.
type ListTemplateReply_TemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTemplateReply_TemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTemplateReply_TemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTemplateReply_TemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTemplateReply_TemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTemplateReply_TemplateValidationError) ErrorName() string {
	return "ListTemplateReply_TemplateValidationError"
}

// Error satisfies the builtin error interface
func (e ListTemplateReply_TemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTemplateReply_Template.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTemplateReply_TemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTemplateReply_TemplateValidationError{}

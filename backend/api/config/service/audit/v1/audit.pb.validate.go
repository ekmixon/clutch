// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/audit/v1/audit.proto

package auditv1

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
)

// Validate checks the field values on EventFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventFilter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Field

	switch m.Value.(type) {

	case *EventFilter_Text:
		// no validation rules for Text

	}

	return nil
}

// EventFilterValidationError is the validation error returned by
// EventFilter.Validate if the designated constraints aren't met.
type EventFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventFilterValidationError) ErrorName() string { return "EventFilterValidationError" }

// Error satisfies the builtin error interface
func (e EventFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventFilterValidationError{}

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Filter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Denylist

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on SinkConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SinkConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SinkConfigValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SinkConfigValidationError is the validation error returned by
// SinkConfig.Validate if the designated constraints aren't met.
type SinkConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SinkConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SinkConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SinkConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SinkConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SinkConfigValidationError) ErrorName() string { return "SinkConfigValidationError" }

// Error satisfies the builtin error interface
func (e SinkConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSinkConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SinkConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SinkConfigValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.StorageProvider.(type) {

	case *Config_DbProvider:

		if len(m.GetDbProvider()) < 1 {
			return ConfigValidationError{
				field:  "DbProvider",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *Config_InMemory:
		// no validation rules for InMemory

	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

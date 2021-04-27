// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/authn/v1/storage.proto

package authnv1

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

// Validate checks the field values on StorageConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StorageConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetEncryptionPassphrase()) < 1 {
		return StorageConfigValidationError{
			field:  "EncryptionPassphrase",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// StorageConfigValidationError is the validation error returned by
// StorageConfig.Validate if the designated constraints aren't met.
type StorageConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageConfigValidationError) ErrorName() string { return "StorageConfigValidationError" }

// Error satisfies the builtin error interface
func (e StorageConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageConfigValidationError{}

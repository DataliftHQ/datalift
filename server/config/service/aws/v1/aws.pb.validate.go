// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/aws/v1/aws.proto

package awsv1

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

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetRegions()) < 1 {
		err := ConfigValidationError{
			field:  "Regions",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetClientConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigValidationError{
					field:  "ClientConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigValidationError{
					field:  "ClientConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClientConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "ClientConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PrimaryAccountAliasDisplayName

	// no validation rules for AwsConfigProfileName

	for idx, item := range m.GetAdditionalAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("AdditionalAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("AdditionalAccounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("AdditionalAccounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}

	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

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

// Validate checks the field values on ClientConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientConfigMultiError, or
// nil if none found.
func (m *ClientConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRetries() < 0 {
		err := ClientConfigValidationError{
			field:  "Retries",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ClientConfigMultiError(errors)
	}

	return nil
}

// ClientConfigMultiError is an error wrapping multiple validation errors
// returned by ClientConfig.ValidateAll() if the designated constraints aren't met.
type ClientConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientConfigMultiError) AllErrors() []error { return m }

// ClientConfigValidationError is the validation error returned by
// ClientConfig.Validate if the designated constraints aren't met.
type ClientConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientConfigValidationError) ErrorName() string { return "ClientConfigValidationError" }

// Error satisfies the builtin error interface
func (e ClientConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientConfigValidationError{}

// Validate checks the field values on AWSAccount with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AWSAccount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSAccount with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AWSAccountMultiError, or
// nil if none found.
func (m *AWSAccount) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSAccount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAlias()) < 1 {
		err := AWSAccountValidationError{
			field:  "Alias",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetAccountNumber()) < 1 {
		err := AWSAccountValidationError{
			field:  "AccountNumber",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetIamRole()) < 1 {
		err := AWSAccountValidationError{
			field:  "IamRole",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRegions()) < 1 {
		err := AWSAccountValidationError{
			field:  "Regions",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AWSAccountMultiError(errors)
	}

	return nil
}

// AWSAccountMultiError is an error wrapping multiple validation errors
// returned by AWSAccount.ValidateAll() if the designated constraints aren't met.
type AWSAccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSAccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSAccountMultiError) AllErrors() []error { return m }

// AWSAccountValidationError is the validation error returned by
// AWSAccount.Validate if the designated constraints aren't met.
type AWSAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSAccountValidationError) ErrorName() string { return "AWSAccountValidationError" }

// Error satisfies the builtin error interface
func (e AWSAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSAccountValidationError{}

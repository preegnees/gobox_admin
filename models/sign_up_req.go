// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignUpReq sign up req
//
// swagger:model SignUpReq
type SignUpReq struct {

	// Ваша электоронная почта
	// Example: email@gmail.com
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// Код который пришел вам на почту
	// Example: 123456
	// Required: true
	EmailCode *int64 `json:"email_code"`

	// Ваш пароль
	// Required: true
	// Min Length: 6
	Password *string `json:"password"`

	// Ваше никнейм
	// Required: true
	// Min Length: 6
	Username *string `json:"username"`
}

// Validate validates this sign up req
func (m *SignUpReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignUpReq) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignUpReq) validateEmailCode(formats strfmt.Registry) error {

	if err := validate.Required("email_code", "body", m.EmailCode); err != nil {
		return err
	}

	return nil
}

func (m *SignUpReq) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 6); err != nil {
		return err
	}

	return nil
}

func (m *SignUpReq) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 6); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sign up req based on context it is used
func (m *SignUpReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignUpReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignUpReq) UnmarshalBinary(b []byte) error {
	var res SignUpReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

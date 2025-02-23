// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TableInfoResp Table information response
//
// swagger:model TableInfoResp
type TableInfoResp struct {

	// Keyspace
	KeySpace string `json:"keySpace,omitempty"`

	// Relation type
	// Enum: [SYSTEM_TABLE_RELATION USER_TABLE_RELATION INDEX_TABLE_RELATION]
	RelationType string `json:"relationType,omitempty"`

	// Size in bytes
	// Read Only: true
	SizeBytes float64 `json:"sizeBytes,omitempty"`

	// Table name
	TableName string `json:"tableName,omitempty"`

	// Table type
	// Enum: [YQL_TABLE_TYPE REDIS_TABLE_TYPE PGSQL_TABLE_TYPE TRANSACTION_STATUS_TABLE_TYPE]
	TableType string `json:"tableType,omitempty"`

	// Table UUID
	// Read Only: true
	// Format: uuid
	TableUUID strfmt.UUID `json:"tableUUID,omitempty"`
}

// Validate validates this table info resp
func (m *TableInfoResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tableInfoRespTypeRelationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYSTEM_TABLE_RELATION","USER_TABLE_RELATION","INDEX_TABLE_RELATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tableInfoRespTypeRelationTypePropEnum = append(tableInfoRespTypeRelationTypePropEnum, v)
	}
}

const (

	// TableInfoRespRelationTypeSYSTEMTABLERELATION captures enum value "SYSTEM_TABLE_RELATION"
	TableInfoRespRelationTypeSYSTEMTABLERELATION string = "SYSTEM_TABLE_RELATION"

	// TableInfoRespRelationTypeUSERTABLERELATION captures enum value "USER_TABLE_RELATION"
	TableInfoRespRelationTypeUSERTABLERELATION string = "USER_TABLE_RELATION"

	// TableInfoRespRelationTypeINDEXTABLERELATION captures enum value "INDEX_TABLE_RELATION"
	TableInfoRespRelationTypeINDEXTABLERELATION string = "INDEX_TABLE_RELATION"
)

// prop value enum
func (m *TableInfoResp) validateRelationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tableInfoRespTypeRelationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TableInfoResp) validateRelationType(formats strfmt.Registry) error {
	if swag.IsZero(m.RelationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRelationTypeEnum("relationType", "body", m.RelationType); err != nil {
		return err
	}

	return nil
}

var tableInfoRespTypeTableTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["YQL_TABLE_TYPE","REDIS_TABLE_TYPE","PGSQL_TABLE_TYPE","TRANSACTION_STATUS_TABLE_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tableInfoRespTypeTableTypePropEnum = append(tableInfoRespTypeTableTypePropEnum, v)
	}
}

const (

	// TableInfoRespTableTypeYQLTABLETYPE captures enum value "YQL_TABLE_TYPE"
	TableInfoRespTableTypeYQLTABLETYPE string = "YQL_TABLE_TYPE"

	// TableInfoRespTableTypeREDISTABLETYPE captures enum value "REDIS_TABLE_TYPE"
	TableInfoRespTableTypeREDISTABLETYPE string = "REDIS_TABLE_TYPE"

	// TableInfoRespTableTypePGSQLTABLETYPE captures enum value "PGSQL_TABLE_TYPE"
	TableInfoRespTableTypePGSQLTABLETYPE string = "PGSQL_TABLE_TYPE"

	// TableInfoRespTableTypeTRANSACTIONSTATUSTABLETYPE captures enum value "TRANSACTION_STATUS_TABLE_TYPE"
	TableInfoRespTableTypeTRANSACTIONSTATUSTABLETYPE string = "TRANSACTION_STATUS_TABLE_TYPE"
)

// prop value enum
func (m *TableInfoResp) validateTableTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tableInfoRespTypeTableTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TableInfoResp) validateTableType(formats strfmt.Registry) error {
	if swag.IsZero(m.TableType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTableTypeEnum("tableType", "body", m.TableType); err != nil {
		return err
	}

	return nil
}

func (m *TableInfoResp) validateTableUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TableUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("tableUUID", "body", "uuid", m.TableUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this table info resp based on the context it is used
func (m *TableInfoResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSizeBytes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTableUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableInfoResp) contextValidateSizeBytes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sizeBytes", "body", float64(m.SizeBytes)); err != nil {
		return err
	}

	return nil
}

func (m *TableInfoResp) contextValidateTableUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tableUUID", "body", strfmt.UUID(m.TableUUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableInfoResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableInfoResp) UnmarshalBinary(b []byte) error {
	var res TableInfoResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

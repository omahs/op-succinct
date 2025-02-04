// Code generated by ent, DO NOT EDIT.

package proofrequest

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the proofrequest type in the database.
	Label = "proof_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStartBlock holds the string denoting the start_block field in the database.
	FieldStartBlock = "start_block"
	// FieldEndBlock holds the string denoting the end_block field in the database.
	FieldEndBlock = "end_block"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRequestAddedTime holds the string denoting the request_added_time field in the database.
	FieldRequestAddedTime = "request_added_time"
	// FieldProverRequestID holds the string denoting the prover_request_id field in the database.
	FieldProverRequestID = "prover_request_id"
	// FieldProofRequestTime holds the string denoting the proof_request_time field in the database.
	FieldProofRequestTime = "proof_request_time"
	// FieldLastUpdatedTime holds the string denoting the last_updated_time field in the database.
	FieldLastUpdatedTime = "last_updated_time"
	// FieldL1BlockNumber holds the string denoting the l1_block_number field in the database.
	FieldL1BlockNumber = "l1_block_number"
	// FieldL1BlockHash holds the string denoting the l1_block_hash field in the database.
	FieldL1BlockHash = "l1_block_hash"
	// FieldProof holds the string denoting the proof field in the database.
	FieldProof = "proof"
	// Table holds the table name of the proofrequest in the database.
	Table = "proof_requests"
)

// Columns holds all SQL columns for proofrequest fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldStartBlock,
	FieldEndBlock,
	FieldStatus,
	FieldRequestAddedTime,
	FieldProverRequestID,
	FieldProofRequestTime,
	FieldLastUpdatedTime,
	FieldL1BlockNumber,
	FieldL1BlockHash,
	FieldProof,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeSPAN Type = "SPAN"
	TypeAGG  Type = "AGG"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeSPAN, TypeAGG:
		return nil
	default:
		return fmt.Errorf("proofrequest: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusUNREQ    Status = "UNREQ"
	StatusREQ      Status = "REQ"
	StatusFAILED   Status = "FAILED"
	StatusCOMPLETE Status = "COMPLETE"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUNREQ, StatusREQ, StatusFAILED, StatusCOMPLETE:
		return nil
	default:
		return fmt.Errorf("proofrequest: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ProofRequest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStartBlock orders the results by the start_block field.
func ByStartBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartBlock, opts...).ToFunc()
}

// ByEndBlock orders the results by the end_block field.
func ByEndBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndBlock, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRequestAddedTime orders the results by the request_added_time field.
func ByRequestAddedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestAddedTime, opts...).ToFunc()
}

// ByProverRequestID orders the results by the prover_request_id field.
func ByProverRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProverRequestID, opts...).ToFunc()
}

// ByProofRequestTime orders the results by the proof_request_time field.
func ByProofRequestTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProofRequestTime, opts...).ToFunc()
}

// ByLastUpdatedTime orders the results by the last_updated_time field.
func ByLastUpdatedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdatedTime, opts...).ToFunc()
}

// ByL1BlockNumber orders the results by the l1_block_number field.
func ByL1BlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldL1BlockNumber, opts...).ToFunc()
}

// ByL1BlockHash orders the results by the l1_block_hash field.
func ByL1BlockHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldL1BlockHash, opts...).ToFunc()
}

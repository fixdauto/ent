// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package tenant

import (
	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the tenant type in the database.
	Label = "tenant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// Table holds the table name of the tenant in the database.
	Table = "tenants"
)

// Columns holds all SQL columns for tenant fields.
var Columns = []string{
	FieldID,
	FieldName,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/facebook/ent/examples/privacytenant/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

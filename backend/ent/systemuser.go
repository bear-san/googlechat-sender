// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bear-san/googlechat-sender/backend/ent/systemuser"
)

// SystemUser is the model entity for the SystemUser schema.
type SystemUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SystemUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case systemuser.FieldID, systemuser.FieldEmail, systemuser.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SystemUser fields.
func (su *SystemUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case systemuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				su.ID = value.String
			}
		case systemuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				su.Email = value.String
			}
		case systemuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				su.Name = value.String
			}
		default:
			su.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SystemUser.
// This includes values selected through modifiers, order, etc.
func (su *SystemUser) Value(name string) (ent.Value, error) {
	return su.selectValues.Get(name)
}

// Update returns a builder for updating this SystemUser.
// Note that you need to call SystemUser.Unwrap() before calling this method if this SystemUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (su *SystemUser) Update() *SystemUserUpdateOne {
	return NewSystemUserClient(su.config).UpdateOne(su)
}

// Unwrap unwraps the SystemUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (su *SystemUser) Unwrap() *SystemUser {
	_tx, ok := su.config.driver.(*txDriver)
	if !ok {
		panic("ent: SystemUser is not a transactional entity")
	}
	su.config.driver = _tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *SystemUser) String() string {
	var builder strings.Builder
	builder.WriteString("SystemUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", su.ID))
	builder.WriteString("email=")
	builder.WriteString(su.Email)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(su.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SystemUsers is a parsable slice of SystemUser.
type SystemUsers []*SystemUser

// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// LinGroup is the model entity for the LinGroup schema.
type LinGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// 分组名称，例如：搬砖者
	Name string `json:"name,omitempty"`
	// Info holds the value of the "info" field.
	// 分组信息：例如：搬砖的人
	Info string `json:"info,omitempty"`
	// Level holds the value of the "level" field.
	// 分组级别 1：root 2：guest 3：user（root、guest分组只能存在一个)
	Level int8 `json:"level,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LinGroupQuery when eager-loading is set.
	Edges LinGroupEdges `json:"edges"`
}

// LinGroupEdges holds the relations/edges for other nodes in the graph.
type LinGroupEdges struct {
	// LinUser holds the value of the lin_user edge.
	LinUser []*LinUser `json:"lin_user,omitempty"`
	// LinPermission holds the value of the lin_permission edge.
	LinPermission []*LinPermission `json:"lin_permission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LinUserOrErr returns the LinUser value or an error if the edge
// was not loaded in eager-loading.
func (e LinGroupEdges) LinUserOrErr() ([]*LinUser, error) {
	if e.loadedTypes[0] {
		return e.LinUser, nil
	}
	return nil, &NotLoadedError{edge: "lin_user"}
}

// LinPermissionOrErr returns the LinPermission value or an error if the edge
// was not loaded in eager-loading.
func (e LinGroupEdges) LinPermissionOrErr() ([]*LinPermission, error) {
	if e.loadedTypes[1] {
		return e.LinPermission, nil
	}
	return nil, &NotLoadedError{edge: "lin_permission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LinGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case lingroup.FieldID, lingroup.FieldLevel:
			values[i] = new(sql.NullInt64)
		case lingroup.FieldName, lingroup.FieldInfo:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LinGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LinGroup fields.
func (lg *LinGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lingroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lg.ID = int(value.Int64)
		case lingroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				lg.Name = value.String
			}
		case lingroup.FieldInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value.Valid {
				lg.Info = value.String
			}
		case lingroup.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				lg.Level = int8(value.Int64)
			}
		}
	}
	return nil
}

// QueryLinUser queries the "lin_user" edge of the LinGroup entity.
func (lg *LinGroup) QueryLinUser() *LinUserQuery {
	return (&LinGroupClient{config: lg.config}).QueryLinUser(lg)
}

// QueryLinPermission queries the "lin_permission" edge of the LinGroup entity.
func (lg *LinGroup) QueryLinPermission() *LinPermissionQuery {
	return (&LinGroupClient{config: lg.config}).QueryLinPermission(lg)
}

// Update returns a builder for updating this LinGroup.
// Note that you need to call LinGroup.Unwrap() before calling this method if this LinGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (lg *LinGroup) Update() *LinGroupUpdateOne {
	return (&LinGroupClient{config: lg.config}).UpdateOne(lg)
}

// Unwrap unwraps the LinGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lg *LinGroup) Unwrap() *LinGroup {
	tx, ok := lg.config.driver.(*txDriver)
	if !ok {
		panic("model: LinGroup is not a transactional entity")
	}
	lg.config.driver = tx.drv
	return lg
}

// String implements the fmt.Stringer.
func (lg *LinGroup) String() string {
	var builder strings.Builder
	builder.WriteString("LinGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", lg.ID))
	builder.WriteString(", name=")
	builder.WriteString(lg.Name)
	builder.WriteString(", info=")
	builder.WriteString(lg.Info)
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", lg.Level))
	builder.WriteByte(')')
	return builder.String()
}

// LinGroups is a parsable slice of LinGroup.
type LinGroups []*LinGroup

func (lg LinGroups) config(cfg config) {
	for _i := range lg {
		lg[_i].config = cfg
	}
}

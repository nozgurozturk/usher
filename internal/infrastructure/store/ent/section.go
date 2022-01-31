// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/section"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SectionQuery when eager-loading is set.
	Edges           SectionEdges `json:"edges"`
	layout_sections *uuid.UUID
}

// SectionEdges holds the relations/edges for other nodes in the graph.
type SectionEdges struct {
	// Layout holds the value of the layout edge.
	Layout *Layout `json:"layout,omitempty"`
	// Rows holds the value of the rows edge.
	Rows []*Row `json:"rows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LayoutOrErr returns the Layout value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SectionEdges) LayoutOrErr() (*Layout, error) {
	if e.loadedTypes[0] {
		if e.Layout == nil {
			// The edge layout was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: layout.Label}
		}
		return e.Layout, nil
	}
	return nil, &NotLoadedError{edge: "layout"}
}

// RowsOrErr returns the Rows value or an error if the edge
// was not loaded in eager-loading.
func (e SectionEdges) RowsOrErr() ([]*Row, error) {
	if e.loadedTypes[1] {
		return e.Rows, nil
	}
	return nil, &NotLoadedError{edge: "rows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case section.FieldName:
			values[i] = new(sql.NullString)
		case section.FieldID:
			values[i] = new(uuid.UUID)
		case section.ForeignKeys[0]: // layout_sections
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Section", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case section.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case section.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case section.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field layout_sections", values[i])
			} else if value.Valid {
				s.layout_sections = new(uuid.UUID)
				*s.layout_sections = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryLayout queries the "layout" edge of the Section entity.
func (s *Section) QueryLayout() *LayoutQuery {
	return (&SectionClient{config: s.config}).QueryLayout(s)
}

// QueryRows queries the "rows" edge of the Section entity.
func (s *Section) QueryRows() *RowQuery {
	return (&SectionClient{config: s.config}).QueryRows(s)
}

// Update returns a builder for updating this Section.
// Note that you need to call Section.Unwrap() before calling this method if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return (&SectionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Section entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section

func (s Sections) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
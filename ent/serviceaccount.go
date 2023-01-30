// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/BradHacker/compsole/ent/serviceaccount"
	"github.com/google/uuid"
)

// ServiceAccount is the model entity for the ServiceAccount schema.
type ServiceAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	// [REQUIRED] The display/common name for the service account.
	DisplayName string `json:"display_name,omitempty"`
	// APIKey holds the value of the "api_key" field.
	// [REQUIRED] The API key for the service account. Equivalent to a username.
	APIKey uuid.UUID `json:"api_key,omitempty"`
	// APISecret holds the value of the "api_secret" field.
	// [REQUIRED] The API secret for the service account. This value MUST be protected.
	APISecret uuid.UUID `json:"api_secret,omitempty"`
	// Active holds the value of the "active" field.
	// [REQUIRED] Determines whether or not the service account is active or not
	Active bool `json:"active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceAccountQuery when eager-loading is set.
	Edges ServiceAccountEdges `json:"edges"`
}

// ServiceAccountEdges holds the relations/edges for other nodes in the graph.
type ServiceAccountEdges struct {
	// ServiceAccountToToken holds the value of the ServiceAccountToToken edge.
	ServiceAccountToToken []*ServiceToken `json:"ServiceAccountToToken,omitempty"`
	// ServiceAccountToActions holds the value of the ServiceAccountToActions edge.
	ServiceAccountToActions []*Action `json:"ServiceAccountToActions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ServiceAccountToTokenOrErr returns the ServiceAccountToToken value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceAccountEdges) ServiceAccountToTokenOrErr() ([]*ServiceToken, error) {
	if e.loadedTypes[0] {
		return e.ServiceAccountToToken, nil
	}
	return nil, &NotLoadedError{edge: "ServiceAccountToToken"}
}

// ServiceAccountToActionsOrErr returns the ServiceAccountToActions value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceAccountEdges) ServiceAccountToActionsOrErr() ([]*Action, error) {
	if e.loadedTypes[1] {
		return e.ServiceAccountToActions, nil
	}
	return nil, &NotLoadedError{edge: "ServiceAccountToActions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceAccount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case serviceaccount.FieldActive:
			values[i] = new(sql.NullBool)
		case serviceaccount.FieldDisplayName:
			values[i] = new(sql.NullString)
		case serviceaccount.FieldID, serviceaccount.FieldAPIKey, serviceaccount.FieldAPISecret:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ServiceAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceAccount fields.
func (sa *ServiceAccount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case serviceaccount.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sa.ID = *value
			}
		case serviceaccount.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				sa.DisplayName = value.String
			}
		case serviceaccount.FieldAPIKey:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value != nil {
				sa.APIKey = *value
			}
		case serviceaccount.FieldAPISecret:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field api_secret", values[i])
			} else if value != nil {
				sa.APISecret = *value
			}
		case serviceaccount.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				sa.Active = value.Bool
			}
		}
	}
	return nil
}

// QueryServiceAccountToToken queries the "ServiceAccountToToken" edge of the ServiceAccount entity.
func (sa *ServiceAccount) QueryServiceAccountToToken() *ServiceTokenQuery {
	return (&ServiceAccountClient{config: sa.config}).QueryServiceAccountToToken(sa)
}

// QueryServiceAccountToActions queries the "ServiceAccountToActions" edge of the ServiceAccount entity.
func (sa *ServiceAccount) QueryServiceAccountToActions() *ActionQuery {
	return (&ServiceAccountClient{config: sa.config}).QueryServiceAccountToActions(sa)
}

// Update returns a builder for updating this ServiceAccount.
// Note that you need to call ServiceAccount.Unwrap() before calling this method if this ServiceAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *ServiceAccount) Update() *ServiceAccountUpdateOne {
	return (&ServiceAccountClient{config: sa.config}).UpdateOne(sa)
}

// Unwrap unwraps the ServiceAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *ServiceAccount) Unwrap() *ServiceAccount {
	tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ServiceAccount is not a transactional entity")
	}
	sa.config.driver = tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *ServiceAccount) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceAccount(")
	builder.WriteString(fmt.Sprintf("id=%v", sa.ID))
	builder.WriteString(", display_name=")
	builder.WriteString(sa.DisplayName)
	builder.WriteString(", api_key=")
	builder.WriteString(fmt.Sprintf("%v", sa.APIKey))
	builder.WriteString(", api_secret=")
	builder.WriteString(fmt.Sprintf("%v", sa.APISecret))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", sa.Active))
	builder.WriteByte(')')
	return builder.String()
}

// ServiceAccounts is a parsable slice of ServiceAccount.
type ServiceAccounts []*ServiceAccount

func (sa ServiceAccounts) config(cfg config) {
	for _i := range sa {
		sa[_i].config = cfg
	}
}

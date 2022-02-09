// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/validation"
	"github.com/google/uuid"
)

// Script is the model entity for the Script schema.
type Script struct {
	config ` json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HclID holds the value of the "hcl_id" field.
	HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" hcl:"name,attr"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty" hcl:"language,attr"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" hcl:"description,optional"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty" hcl:"source,attr"`
	// SourceType holds the value of the "source_type" field.
	SourceType string `json:"source_type,omitempty" hcl:"source_type,attr"`
	// Cooldown holds the value of the "cooldown" field.
	Cooldown int `json:"cooldown,omitempty" hcl:"cooldown,optional"`
	// Timeout holds the value of the "timeout" field.
	Timeout int `json:"timeout,omitempty" hcl:"timeout,optional"`
	// IgnoreErrors holds the value of the "ignore_errors" field.
	IgnoreErrors bool `json:"ignore_errors,omitempty" hcl:"ignore_errors,optional"`
	// Args holds the value of the "args" field.
	Args []string `json:"args,omitempty" hcl:"args,optional"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty" hcl:"disabled,optional" `
	// Vars holds the value of the "vars" field.
	Vars map[string]string `json:"vars,omitempty" hcl:"vars,optional"`
	// AbsPath holds the value of the "abs_path" field.
	AbsPath string `json:"abs_path,omitempty" hcl:"abs_path,optional"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,optional"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScriptQuery when eager-loading is set.
	Edges ScriptEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// ScriptToUser holds the value of the ScriptToUser edge.
	HCLScriptToUser []*User `json:"ScriptToUser,omitempty" hcl:"maintainer,block"`
	// ScriptToFinding holds the value of the ScriptToFinding edge.
	HCLScriptToFinding []*Finding `json:"ScriptToFinding,omitempty" hcl:"finding,block"`
	// ScriptToEnvironment holds the value of the ScriptToEnvironment edge.
	HCLScriptToEnvironment *Environment `json:"ScriptToEnvironment,omitempty"`
	// ScriptToValidation holds the value of the ScriptToValidation edge.
	HCLScriptToValidation *Validation `json:"ScriptToValidation,omitempty" hcl:"validation,block"`
	//
	environment_environment_to_script *uuid.UUID
	script_script_to_validation       *uuid.UUID
}

// ScriptEdges holds the relations/edges for other nodes in the graph.
type ScriptEdges struct {
	// ScriptToUser holds the value of the ScriptToUser edge.
	ScriptToUser []*User `json:"ScriptToUser,omitempty" hcl:"maintainer,block"`
	// ScriptToFinding holds the value of the ScriptToFinding edge.
	ScriptToFinding []*Finding `json:"ScriptToFinding,omitempty" hcl:"finding,block"`
	// ScriptToEnvironment holds the value of the ScriptToEnvironment edge.
	ScriptToEnvironment *Environment `json:"ScriptToEnvironment,omitempty"`
	// ScriptToValidation holds the value of the ScriptToValidation edge.
	ScriptToValidation *Validation `json:"ScriptToValidation,omitempty" hcl:"validation,block"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ScriptToUserOrErr returns the ScriptToUser value or an error if the edge
// was not loaded in eager-loading.
func (e ScriptEdges) ScriptToUserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.ScriptToUser, nil
	}
	return nil, &NotLoadedError{edge: "ScriptToUser"}
}

// ScriptToFindingOrErr returns the ScriptToFinding value or an error if the edge
// was not loaded in eager-loading.
func (e ScriptEdges) ScriptToFindingOrErr() ([]*Finding, error) {
	if e.loadedTypes[1] {
		return e.ScriptToFinding, nil
	}
	return nil, &NotLoadedError{edge: "ScriptToFinding"}
}

// ScriptToEnvironmentOrErr returns the ScriptToEnvironment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScriptEdges) ScriptToEnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[2] {
		if e.ScriptToEnvironment == nil {
			// The edge ScriptToEnvironment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.ScriptToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "ScriptToEnvironment"}
}

// ScriptToValidationOrErr returns the ScriptToValidation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScriptEdges) ScriptToValidationOrErr() (*Validation, error) {
	if e.loadedTypes[3] {
		if e.ScriptToValidation == nil {
			// The edge ScriptToValidation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: validation.Label}
		}
		return e.ScriptToValidation, nil
	}
	return nil, &NotLoadedError{edge: "ScriptToValidation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Script) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case script.FieldArgs, script.FieldVars, script.FieldTags:
			values[i] = new([]byte)
		case script.FieldIgnoreErrors, script.FieldDisabled:
			values[i] = new(sql.NullBool)
		case script.FieldCooldown, script.FieldTimeout:
			values[i] = new(sql.NullInt64)
		case script.FieldHclID, script.FieldName, script.FieldLanguage, script.FieldDescription, script.FieldSource, script.FieldSourceType, script.FieldAbsPath:
			values[i] = new(sql.NullString)
		case script.FieldID:
			values[i] = new(uuid.UUID)
		case script.ForeignKeys[0]: // environment_environment_to_script
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case script.ForeignKeys[1]: // script_script_to_validation
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Script", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Script fields.
func (s *Script) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case script.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case script.FieldHclID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hcl_id", values[i])
			} else if value.Valid {
				s.HclID = value.String
			}
		case script.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case script.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				s.Language = value.String
			}
		case script.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case script.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				s.Source = value.String
			}
		case script.FieldSourceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_type", values[i])
			} else if value.Valid {
				s.SourceType = value.String
			}
		case script.FieldCooldown:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cooldown", values[i])
			} else if value.Valid {
				s.Cooldown = int(value.Int64)
			}
		case script.FieldTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				s.Timeout = int(value.Int64)
			}
		case script.FieldIgnoreErrors:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ignore_errors", values[i])
			} else if value.Valid {
				s.IgnoreErrors = value.Bool
			}
		case script.FieldArgs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field args", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Args); err != nil {
					return fmt.Errorf("unmarshal field args: %w", err)
				}
			}
		case script.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				s.Disabled = value.Bool
			}
		case script.FieldVars:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vars", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Vars); err != nil {
					return fmt.Errorf("unmarshal field vars: %w", err)
				}
			}
		case script.FieldAbsPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abs_path", values[i])
			} else if value.Valid {
				s.AbsPath = value.String
			}
		case script.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case script.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field environment_environment_to_script", values[i])
			} else if value.Valid {
				s.environment_environment_to_script = new(uuid.UUID)
				*s.environment_environment_to_script = *value.S.(*uuid.UUID)
			}
		case script.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field script_script_to_validation", values[i])
			} else if value.Valid {
				s.script_script_to_validation = new(uuid.UUID)
				*s.script_script_to_validation = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryScriptToUser queries the "ScriptToUser" edge of the Script entity.
func (s *Script) QueryScriptToUser() *UserQuery {
	return (&ScriptClient{config: s.config}).QueryScriptToUser(s)
}

// QueryScriptToFinding queries the "ScriptToFinding" edge of the Script entity.
func (s *Script) QueryScriptToFinding() *FindingQuery {
	return (&ScriptClient{config: s.config}).QueryScriptToFinding(s)
}

// QueryScriptToEnvironment queries the "ScriptToEnvironment" edge of the Script entity.
func (s *Script) QueryScriptToEnvironment() *EnvironmentQuery {
	return (&ScriptClient{config: s.config}).QueryScriptToEnvironment(s)
}

// QueryScriptToValidation queries the "ScriptToValidation" edge of the Script entity.
func (s *Script) QueryScriptToValidation() *ValidationQuery {
	return (&ScriptClient{config: s.config}).QueryScriptToValidation(s)
}

// Update returns a builder for updating this Script.
// Note that you need to call Script.Unwrap() before calling this method if this Script
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Script) Update() *ScriptUpdateOne {
	return (&ScriptClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Script entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Script) Unwrap() *Script {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Script is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Script) String() string {
	var builder strings.Builder
	builder.WriteString("Script(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", hcl_id=")
	builder.WriteString(s.HclID)
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", language=")
	builder.WriteString(s.Language)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteString(", source=")
	builder.WriteString(s.Source)
	builder.WriteString(", source_type=")
	builder.WriteString(s.SourceType)
	builder.WriteString(", cooldown=")
	builder.WriteString(fmt.Sprintf("%v", s.Cooldown))
	builder.WriteString(", timeout=")
	builder.WriteString(fmt.Sprintf("%v", s.Timeout))
	builder.WriteString(", ignore_errors=")
	builder.WriteString(fmt.Sprintf("%v", s.IgnoreErrors))
	builder.WriteString(", args=")
	builder.WriteString(fmt.Sprintf("%v", s.Args))
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", s.Disabled))
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", s.Vars))
	builder.WriteString(", abs_path=")
	builder.WriteString(s.AbsPath)
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", s.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// Scripts is a parsable slice of Script.
type Scripts []*Script

func (s Scripts) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}

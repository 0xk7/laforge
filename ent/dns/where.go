// Code generated by entc, DO NOT EDIT.

package dns

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HclID applies equality check predicate on the "hcl_id" field. It's identical to HclIDEQ.
func HclID(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHclID), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// RootDomain applies equality check predicate on the "root_domain" field. It's identical to RootDomainEQ.
func RootDomain(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDomain), v))
	})
}

// HclIDEQ applies the EQ predicate on the "hcl_id" field.
func HclIDEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHclID), v))
	})
}

// HclIDNEQ applies the NEQ predicate on the "hcl_id" field.
func HclIDNEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHclID), v))
	})
}

// HclIDIn applies the In predicate on the "hcl_id" field.
func HclIDIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHclID), v...))
	})
}

// HclIDNotIn applies the NotIn predicate on the "hcl_id" field.
func HclIDNotIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHclID), v...))
	})
}

// HclIDGT applies the GT predicate on the "hcl_id" field.
func HclIDGT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHclID), v))
	})
}

// HclIDGTE applies the GTE predicate on the "hcl_id" field.
func HclIDGTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHclID), v))
	})
}

// HclIDLT applies the LT predicate on the "hcl_id" field.
func HclIDLT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHclID), v))
	})
}

// HclIDLTE applies the LTE predicate on the "hcl_id" field.
func HclIDLTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHclID), v))
	})
}

// HclIDContains applies the Contains predicate on the "hcl_id" field.
func HclIDContains(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHclID), v))
	})
}

// HclIDHasPrefix applies the HasPrefix predicate on the "hcl_id" field.
func HclIDHasPrefix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHclID), v))
	})
}

// HclIDHasSuffix applies the HasSuffix predicate on the "hcl_id" field.
func HclIDHasSuffix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHclID), v))
	})
}

// HclIDEqualFold applies the EqualFold predicate on the "hcl_id" field.
func HclIDEqualFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHclID), v))
	})
}

// HclIDContainsFold applies the ContainsFold predicate on the "hcl_id" field.
func HclIDContainsFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHclID), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// RootDomainEQ applies the EQ predicate on the "root_domain" field.
func RootDomainEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDomain), v))
	})
}

// RootDomainNEQ applies the NEQ predicate on the "root_domain" field.
func RootDomainNEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRootDomain), v))
	})
}

// RootDomainIn applies the In predicate on the "root_domain" field.
func RootDomainIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRootDomain), v...))
	})
}

// RootDomainNotIn applies the NotIn predicate on the "root_domain" field.
func RootDomainNotIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRootDomain), v...))
	})
}

// RootDomainGT applies the GT predicate on the "root_domain" field.
func RootDomainGT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRootDomain), v))
	})
}

// RootDomainGTE applies the GTE predicate on the "root_domain" field.
func RootDomainGTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRootDomain), v))
	})
}

// RootDomainLT applies the LT predicate on the "root_domain" field.
func RootDomainLT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRootDomain), v))
	})
}

// RootDomainLTE applies the LTE predicate on the "root_domain" field.
func RootDomainLTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRootDomain), v))
	})
}

// RootDomainContains applies the Contains predicate on the "root_domain" field.
func RootDomainContains(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRootDomain), v))
	})
}

// RootDomainHasPrefix applies the HasPrefix predicate on the "root_domain" field.
func RootDomainHasPrefix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRootDomain), v))
	})
}

// RootDomainHasSuffix applies the HasSuffix predicate on the "root_domain" field.
func RootDomainHasSuffix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRootDomain), v))
	})
}

// RootDomainEqualFold applies the EqualFold predicate on the "root_domain" field.
func RootDomainEqualFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRootDomain), v))
	})
}

// RootDomainContainsFold applies the ContainsFold predicate on the "root_domain" field.
func RootDomainContainsFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRootDomain), v))
	})
}

// HasDNSToEnvironment applies the HasEdge predicate on the "DNSToEnvironment" edge.
func HasDNSToEnvironment() predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DNSToEnvironmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DNSToEnvironmentTable, DNSToEnvironmentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDNSToEnvironmentWith applies the HasEdge predicate on the "DNSToEnvironment" edge with a given conditions (other predicates).
func HasDNSToEnvironmentWith(preds ...predicate.Environment) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DNSToEnvironmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DNSToEnvironmentTable, DNSToEnvironmentPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDNSToCompetition applies the HasEdge predicate on the "DNSToCompetition" edge.
func HasDNSToCompetition() predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DNSToCompetitionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DNSToCompetitionTable, DNSToCompetitionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDNSToCompetitionWith applies the HasEdge predicate on the "DNSToCompetition" edge with a given conditions (other predicates).
func HasDNSToCompetitionWith(preds ...predicate.Competition) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DNSToCompetitionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DNSToCompetitionTable, DNSToCompetitionPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		p(s.Not())
	})
}
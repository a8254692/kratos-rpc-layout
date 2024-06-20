// Code generated by ent, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// Pass applies equality check predicate on the "pass" field. It's identical to PassEQ.
func Pass(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPass, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPhone, v))
}

// PassEQ applies the EQ predicate on the "pass" field.
func PassEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPass, v))
}

// PassNEQ applies the NEQ predicate on the "pass" field.
func PassNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPass, v))
}

// PassIn applies the In predicate on the "pass" field.
func PassIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldPass, vs...))
}

// PassNotIn applies the NotIn predicate on the "pass" field.
func PassNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldPass, vs...))
}

// PassGT applies the GT predicate on the "pass" field.
func PassGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldPass, v))
}

// PassGTE applies the GTE predicate on the "pass" field.
func PassGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldPass, v))
}

// PassLT applies the LT predicate on the "pass" field.
func PassLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldPass, v))
}

// PassLTE applies the LTE predicate on the "pass" field.
func PassLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldPass, v))
}

// PassContains applies the Contains predicate on the "pass" field.
func PassContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldPass, v))
}

// PassHasPrefix applies the HasPrefix predicate on the "pass" field.
func PassHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldPass, v))
}

// PassHasSuffix applies the HasSuffix predicate on the "pass" field.
func PassHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldPass, v))
}

// PassIsNil applies the IsNil predicate on the "pass" field.
func PassIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldPass))
}

// PassNotNil applies the NotNil predicate on the "pass" field.
func PassNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldPass))
}

// PassEqualFold applies the EqualFold predicate on the "pass" field.
func PassEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldPass, v))
}

// PassContainsFold applies the ContainsFold predicate on the "pass" field.
func PassContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldPass, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldPhone, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(sql.NotPredicates(p))
}

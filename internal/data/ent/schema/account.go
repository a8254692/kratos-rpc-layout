package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Comment("自增ID").Unique(),

		field.String("pass").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Optional().Comment("密码"),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)", // Override MySQL.
		}).Optional().Default("").Comment("email"),

		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Optional().Default("").Comment("手机号"),
	}

}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

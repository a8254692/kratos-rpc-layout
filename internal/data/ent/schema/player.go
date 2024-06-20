package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

func (Player) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "player"},
	}
}

// Fields of the Player.
func (Player) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.Int64("account_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Comment("账户ID"),

		field.Int64("player_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Comment("玩家ID"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Optional(),

		field.Int32("gender").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(1).Comment("性别"),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Optional().Comment("头像地址"),
	}

}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return nil
}

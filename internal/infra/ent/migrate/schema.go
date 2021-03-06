// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TTeatokUserColumns holds the columns for the "t_teatok_user" table.
	TTeatokUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "mobile", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "gender", Type: field.TypeInt},
		{Name: "job", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TTeatokUserTable holds the schema information for the "t_teatok_user" table.
	TTeatokUserTable = &schema.Table{
		Name:        "t_teatok_user",
		Columns:     TTeatokUserColumns,
		PrimaryKey:  []*schema.Column{TTeatokUserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TTeatokUserTable,
	}
)

func init() {
	TTeatokUserTable.Annotation = &entsql.Annotation{
		Table: "t_teatok_user",
	}
}

// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.
//
// Code has been copied from the ent examples folder to demonstrate common schema patterns.
// You can find the original code here: https://github.com/ent/ent/tree/master/entc/integration/edgeschema

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("groups").
			Through("joined_users", UserGroup.Type),
		edge.From("tags", Tag.Type).
			Ref("groups").
			Through("group_tags", GroupTag.Type),
	}
}
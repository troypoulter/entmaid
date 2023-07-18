// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.
//
// Code has been copied from the ent examples folder to demonstrate common schema patterns.
// The Policy has been removed as this isn't too related to the ER generated diagram.
// You can find the original code here: https://github.com/ent/ent/tree/master/entc/integration/edgeschema

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TweetLike holds the schema definition for the TweetLike entity.
type TweetLike struct {
	ent.Schema
}

func (TweetLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
	}
}

// Fields of the TweetLike.
func (TweetLike) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the TweetLike.
func (TweetLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}

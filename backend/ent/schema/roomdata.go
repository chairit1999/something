package schema

import "github.com/facebookincubator/ent"

// RoomData holds the schema definition for the RoomData entity.
type RoomData struct {
	ent.Schema
}

// Fields of the RoomData.
func (RoomData) Fields() []ent.Field {
	return nil
}

// Edges of the RoomData.
func (RoomData) Edges() []ent.Edge {
	return nil
}

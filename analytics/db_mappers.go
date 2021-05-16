package analytics

import "time"

type DBExampleStruct struct {
	Id         string    `db:"id" json:"id" bson:"_id"`
	Label      string    `db:"label" json:"label" bson:"label"`
	EntityType string    `db:"entity_type" json:"entity_type" bson:"entity_type"`
	EntityId   string    `db:"entity_id" json:"entity_id" bson:"entity_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at" bson:"updated_at"`
}

package model

type Netflix struct {
	// ID      primitive.ObjectID `json:"_id,ommitempty" bson:"_id,ommitempty" `
	Movie   string `json:"movie,ommitempty"`
	Watched bool   `json:"watched,ommitempty" `
}

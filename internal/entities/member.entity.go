package entities

type Member struct {
	ID    string  `json:"_id" bson:"_id" validate:"required,uuid4"`
	Name  string  `json:"name" bson:"name" validate:"required,uuid4"`
	Email float64 `json:"email" bson:"email" validate:"required,float64"`
}

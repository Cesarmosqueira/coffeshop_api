package product

type Product struct {
	ID		string		`mongo:"-" json:"id,omitempty" bson:"_id"`
	Name	string		`mongo:"name" json:"name"`
	Cost	float32		`mongo:"cost" json:"cost"`
}

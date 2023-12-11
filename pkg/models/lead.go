package models


type Lead struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Capital float64 `json:"capital"`
	Age     uint8   `json:"age"`
}

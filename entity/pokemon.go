package entity

type Pokemon struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Type1      string  `json:"type1"`
	Type2      string  `json:"type2"`
	Total      float32 `json:"total"`
	HP         float32 `json:"hp"`
	Attack     float32 `json:"attack"`
	Defence    float32 `json:"defense"`
	SPAttack   float32 `json:"spattack"`
	SPDefence  float32 `json:"spdefence"`
	Speed      float32 `json:"speed"`
	Generation float32 `json:"generation"`
	Lengendary bool    `json:"legendary"`
}

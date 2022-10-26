package entity

type PokemonCSV struct {
	Name       string  `csv:"Name"`
	Type1      string  `csv:"Type 1"`
	Type2      string  `csv:"Type 2"`
	Total      float32 `csv:"Total"`
	HP         float32 `csv:"HP"`
	Attack     float32 `csv:"Attack"`
	Defence    float32 `csv:"Defense"`
	SPAttack   float32 `csv:"Sp. Atk"`
	SPDefence  float32 `csv:"Sp. Def"`
	Speed      float32 `csv:"Speed"`
	Generation float32 `csv:"Generation"`
	Lengendary bool    `csv:"Legendary"`
}

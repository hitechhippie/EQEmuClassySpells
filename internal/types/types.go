package types

type ClassySpellMappings struct {
	Spells []SpellMap `yaml:"Spells"`
}

type SpellMap struct {
	SpellName   string `yaml:"SpellName"`
	SpellEffect int    `yaml:"SpellEffect"`
}

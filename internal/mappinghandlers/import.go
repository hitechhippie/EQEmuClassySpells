package mappinghandlers

import (
	"EQEmuClassySpells/internal/types"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

func ExportMappingsFromTxt(f string, o string) error {
	var data types.ClassySpellMappings
	var dataEntries []types.SpellMap
	var parseReader *csv.Reader

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	parseReader = csv.NewReader(file)
	parseReader.Comma = '^'

	for {
		row, err := parseReader.Read()
		if err == io.EOF {
			break
		}

		spellName := ""
		effectID := ""

		for fieldID, field := range row {
			switch fieldID {
			case 1:
				spellName = field
			case 145:
				effectID = field
			}

			if spellName != "" && effectID != "" {
				effectIDint, _ := strconv.Atoi(effectID)

				var newEntry types.SpellMap

				newEntry.SpellName = spellName
				newEntry.SpellEffect = effectIDint

				// Append the new entry to the main data set
				dataEntries = append(dataEntries, newEntry)

				// Zero the iterative placeholders
				spellName = ""
				effectID = ""
			}
		}
	}

	data.Spells = dataEntries

	out, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	outFile, err := os.Create(o)
	if err != nil {
		return err
	}

	_, err = outFile.WriteString("---\n")
	if err != nil {
		return err
	}
	_, err = outFile.WriteString(string(out))
	if err != nil {
		return err
	}

	return nil
}

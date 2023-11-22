package mappinghandlers

import (
	"EQEmuClassySpells/internal/types"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadMappingsFromYAML(f string) (*types.ClassySpellMappings, error) {
	file, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var data types.ClassySpellMappings

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func UpdateSpellsFileFromMappings(mappings *types.ClassySpellMappings, f string) error {
	var parseReader *csv.Reader
	var newWriter io.Writer

	fileIn, err := os.Open(f)
	if err != nil {
		return err
	}

	parseReader = csv.NewReader(fileIn)
	parseReader.Comma = '^'
	parseReader.LazyQuotes = true

	fileOut, err := os.Create(f + "_classyspells")
	if err != nil {
		return err
	}
	defer fileOut.Close()

	newWriter = bufio.NewWriter(fileOut)

	for {
		row, err := parseReader.Read()
		if err == io.EOF {
			break
		}

		spellEffect := ""
		newRow := []string{}

		for fieldID, field := range row {
			if fieldID == 0 {
				newRow = append(newRow, field)
				continue
			}

			if fieldID == 1 {
				for _, mapping := range mappings.Spells {
					if mapping.SpellName == field {
						spellEffect = fmt.Sprint(mapping.SpellEffect)
					}
				}
			}

			if fieldID == 145 {
				if spellEffect != "" {
					newRow = append(newRow, "^"+spellEffect)
					spellEffect = ""
					continue
				} else {
					newRow = append(newRow, "^"+field)
					continue
				}

			} else {
				newRow = append(newRow, "^"+field)
			}

			if fieldID == len(row)-1 {
				for _, data := range newRow {
					_, _ = newWriter.Write([]byte(data))
				}

				newWriter.Write([]byte("\n"))

				break
			}
		}
	}

	return nil
}

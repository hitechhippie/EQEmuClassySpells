package main

import (
	"EQEmuClassySpells/internal/arguments"
	"EQEmuClassySpells/internal/mappinghandlers"
	"EQEmuClassySpells/internal/types"
	"fmt"
	"os"
)

//var classySpellMappings *types.ClassySpellMappings

func main() {
	var err error
	var args []string

	var sourceSpellsFilePath string
	mappingYamlFilePath := "./spellmap.yaml"

	var spellMappings *types.ClassySpellMappings

	args = os.Args[1:]
	if len(args) == 0 {
		arguments.Help()
		os.Exit(1)
	}

	for argIndex, arg := range args {
		if arg == "--help" {
			arguments.Help()
			os.Exit(0)
		} else if arg == "--import" {
			sourceSpellsFilePath = args[argIndex+1]
			err = mappinghandlers.ExportMappingsFromTxt(sourceSpellsFilePath, mappingYamlFilePath)
			if err != nil {
				fmt.Println("ERROR: processing import:", err)
				os.Exit(1)
			}

			fmt.Println("Updated YAML mappings file:", mappingYamlFilePath)
			os.Exit(0)

		} else if arg == "--update" {
			sourceSpellsFilePath = args[argIndex+1]
			spellMappings, err = mappinghandlers.LoadMappingsFromYAML(mappingYamlFilePath)
			if err != nil {
				fmt.Println("ERROR: reading spell mappings:", err)
				os.Exit(1)
			}

			mappinghandlers.UpdateSpellsFileFromMappings(spellMappings, sourceSpellsFilePath)
			if err != nil {
				fmt.Println("ERROR: writing new spells file:", err)
				os.Exit(1)
			}

			fmt.Println("Updated new spell file,(_classyspells)", "from:", sourceSpellsFilePath, "using mappings:", mappingYamlFilePath)
			os.Exit(0)
		} else {
			arguments.Help()
			os.Exit(0)
		}
	}

	//classySpellMappings, err = mappinghandlers.LoadMappingsFromYAML(destMappingYamlFilePath)

	/*for _, spell := range classySpellMappings.Spells {
		fmt.Println("Spell:", spell.SpellName, " // ID:", spell.SpellEffect)
	}*/
	os.Exit(0)
}

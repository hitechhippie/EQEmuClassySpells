package arguments

import "fmt"

func Help() {
	fmt.Println("Usage: classyspells [OPTION]... [FILE]")
	fmt.Println("Import spell mappings from a known spells_us.txt file into yaml format or update an existing spells_us.txt file's spell effects entries for classic spell effects.")
	fmt.Println("--import --update --help")
}

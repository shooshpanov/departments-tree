package main

import "fmt"

const delimiter = "***"

type Department struct {
	departmentID    string
	departmentName  string
	parentID string
}

func write(dep []Department, parent string, depth int) {
	for _, d := range dep {
		if d.parentID == parent {
			for i := 0; i < depth; i++ {
				fmt.Print(delimiter)
			}
			fmt.Print(d.departmentName, "\n\n")
			write(dep, d.departmentID, depth + 1)
		}
	}
}

func main() {
	data := []Department{
		{"01", "Dept A", "-1"},
		{"02", "subA01", "01"},
		{"03", "subA02", "01"},
		{"04", "subsubA02", "03"},
		{"06", "Dept B", "-1"},
		{"07", "subB01", "06"},
		{"08", "subB02", "06"},
		{"09", "subsubB02", "08"},
	}

	write(data, "-1", 0)
}

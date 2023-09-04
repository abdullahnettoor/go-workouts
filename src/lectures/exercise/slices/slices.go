//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func showAssemblyLine(title string, line []Part) {
	fmt.Println(title)
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
	fmt.Println()
}

func main() {
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Screws"
	assemblyLine[1] = "Nuts"
	assemblyLine[2] = "Bolts"

	showAssemblyLine("Parts to assemble", assemblyLine)

	assemblyLine = append(assemblyLine, "Small Screws", "Lid")
	showAssemblyLine("Two more parts added", assemblyLine)

	assemblyLine = assemblyLine[3:]
	showAssemblyLine("After fitting first 3 parts", assemblyLine)
}

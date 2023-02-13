package main

import (
	"fmt"
	"os"

	"github.com/johanhenriksson/spirv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: reflect <shader.spv>")
		os.Exit(1)
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	shader, err := spirv.FromSource(bytes)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// entry points
	fmt.Println("default entry point: ", shader.EntryPoint.Name)
	fmt.Println("default shader stage:", shader.EntryPoint.Stage)
	for _, entry := range shader.EntryPoints {
		fmt.Println("  entry point: ", entry.Name)
		fmt.Println("  shader stage:", shader.EntryPoint.Stage)
	}

	// inputs
	fmt.Println()
	fmt.Println("inputs:")
	inputs, err := shader.EnumerateInputVariables()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	for i, input := range inputs {
		fmt.Println("  name:    ", input.Name)
		fmt.Println("  location:", input.Location)
		fmt.Println("  type:    ", input.Type)
		fmt.Println("  storage: ", input.StorageClass)
		if i < len(inputs)-1 {
			fmt.Println()
		}
	}

	// descriptors
	fmt.Println()
	fmt.Println("descriptors:")
	descriptorSets, err := shader.EnumerateDescriptorSets()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	for i, desc := range descriptorSets {
		fmt.Println("  set:", desc.Set)
		for j, binding := range desc.Bindings {
			fmt.Println("    binding:", binding.Binding)
			fmt.Println("    name:   ", binding.Name)
			fmt.Println("    type:   ", binding.Type)
			fmt.Println("    storage:", binding.StorageClass)
			if j < len(desc.Bindings)-1 {
				fmt.Println()

			}
		}
		if i < len(descriptorSets)-1 {
			fmt.Println()
		}
	}
}

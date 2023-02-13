package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johanhenriksson/spirv"
)

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	mod, err := spirv.FromSource(bytes)
	if err != nil {
		log.Fatal(err)
	}
	defer mod.Destroy()

	log.Println("inputs:")
	inputs, err := mod.EnumerateInputVariables()
	if err != nil {
		log.Fatal(err)
	}
	for _, input := range inputs {
		log.Println("  name:         ", input.Name)
		log.Println("  location:     ", input.Location)
		log.Println("  type:         ", fmt.Sprintf("0x%x", input.TypeFlags))
		log.Println("  storage class:", input.StorageClass)
	}

	log.Println("descriptors:")
	descriptorSets, err := mod.EnumerateDescriptorSets()
	if err != nil {
		log.Fatal(err)
	}
	for _, desc := range descriptorSets {
		log.Println("  set:", desc.Set)
		for _, binding := range desc.Bindings {
			log.Println("    binding:", binding.Binding)
			log.Println("    name:   ", binding.Name)
		}
	}
}

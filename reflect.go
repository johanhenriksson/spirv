package spirv

type Input struct {
	Name         string
	Location     int
	StorageClass int
	TypeFlags    int
	TypeName     string
}

type DescriptorSet struct {
	Set      int
	Bindings []DescriptorBinding
}

type DescriptorBinding struct {
	Binding int
	Name    string
}

package spirv

type Input struct {
	Name         string
	Location     int
	StorageClass StorageClass
	Type         TypeDescription
}

type DescriptorSet struct {
	Set      int
	Bindings []DescriptorBinding
}

type DescriptorBinding struct {
	Binding         int
	Name            string
	InputAttachment int
	StorageClass    StorageClass
	Type            TypeDescription
}

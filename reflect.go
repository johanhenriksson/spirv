package spirv

import (
	"fmt"
	"strings"
)

type Input struct {
	Name         string
	Location     int
	StorageClass int
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
	Type            TypeDescription
}

type TypeDescription struct {
	Name             string
	Flags            int
	StructMemberName string
	StorageClass     int
	Members          []TypeDescription
}

func (t TypeDescription) String() string {
	sb := strings.Builder{}
	if len(t.Name) > 0 {
		sb.WriteString(t.Name)
		sb.WriteRune(' ')
	}
	sb.WriteString(fmt.Sprintf("0x%x", t.Flags))
	if len(t.Members) > 0 {
		sb.WriteString(" { ")
		for i, m := range t.Members {
			sb.WriteString(m.StructMemberName)
			sb.WriteString(": ")
			sb.WriteString(m.String())
			if i < len(t.Members)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(" }")
	}
	return sb.String()
}

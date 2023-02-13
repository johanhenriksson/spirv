package spirv

// #include "./spirv-reflect/spirv_reflect.h"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type TypeDescription struct {
	ptr              *C.SpvReflectTypeDescription
	Name             string
	Flags            int
	StructMemberName string
	StorageClass     StorageClass
	Members          []TypeDescription
}

func parseTypeDescription(t *C.SpvReflectTypeDescription) TypeDescription {
	members := make([]TypeDescription, 0, t.member_count)
	memberSlice := ptrToSlice[C.SpvReflectTypeDescription](
		unsafe.Pointer(t.members), int(t.member_count), int(t.member_count))
	for _, member := range memberSlice {
		m := parseTypeDescription(&member)
		members = append(members, m)
	}

	return TypeDescription{
		ptr: t,

		Name:             C.GoString(t.type_name),
		Flags:            int(t.type_flags),
		StructMemberName: C.GoString(t.struct_member_name),
		StorageClass:     StorageClass(t.storage_class),
		Members:          members,
	}
}

func (t TypeDescription) String() string {
	sb := strings.Builder{}
	if len(t.Name) > 0 {
		sb.WriteString(t.Name)
		sb.WriteRune(' ')
	}
	// sb.WriteString(fmt.Sprintf("0x%x", t.Flags))
	sb.WriteString(typeFlagsToString(t))
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

func typeFlagsToString(t TypeDescription) string {
	if t.Flags&C.SPV_REFLECT_TYPE_FLAG_MATRIX > 0 {
		return fmt.Sprintf("mat%dx%d",
			t.ptr.traits.numeric.matrix.row_count,
			t.ptr.traits.numeric.matrix.column_count)
	} else if t.Flags&C.SPV_REFLECT_TYPE_FLAG_VECTOR > 0 {
		return fmt.Sprintf("vec%d", t.ptr.traits.numeric.vector.component_count)
	}

	switch t.Flags & 0xF {
	case C.SPV_REFLECT_TYPE_FLAG_VOID:
		return "void"
	case C.SPV_REFLECT_TYPE_FLAG_BOOL:
		return "bool"
	case C.SPV_REFLECT_TYPE_FLAG_INT:
		if t.ptr.traits.numeric.scalar.signedness > 0 {
			return "int"
		} else {
			return "uint"
		}
	case C.SPV_REFLECT_TYPE_FLAG_FLOAT:
		switch t.ptr.traits.numeric.scalar.width {
		case 32:
			return "float"
		case 64:
			return "double"
		}

	case C.SPV_REFLECT_TYPE_FLAG_STRUCT:
		return "struct"
	}
	return fmt.Sprintf("0x%x", t.Flags)
}

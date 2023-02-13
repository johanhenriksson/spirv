package spirv

// #include "./spirv-reflect/spirv_reflect.h"
import "C"

type StorageClass C.SpvStorageClass

const (
	StorageClassInput          = StorageClass(C.SpvStorageClassInput)
	StorageClassUniform        = StorageClass(C.SpvStorageClassUniform)
	StorageClassOutput         = StorageClass(C.SpvStorageClassOutput)
	StorageClassImage          = StorageClass(C.SpvStorageClassImage)
	StorageClassPushConstant   = StorageClass(C.SpvStorageClassPushConstant)
	StorageClassStorageBuffer  = StorageClass(C.SpvStorageClassStorageBuffer)
	StorageClassWorkgroup      = StorageClass(C.SpvStorageClassWorkgroup)
	StorageClassCrossWorkgroup = StorageClass(C.SpvStorageClassCrossWorkgroup)
	StorageClassFunction       = StorageClass(C.SpvStorageClassFunction)
	StorageClassGeneric        = StorageClass(C.SpvStorageClassGeneric)
	StorageClassPrivate        = StorageClass(C.SpvStorageClassPrivate)
)

func (s StorageClass) String() string {
	switch C.SpvStorageClass(s) {
	case C.SpvStorageClassUniformConstant:
		return "UniformConstant"
	case C.SpvStorageClassInput:
		return "Input"
	case C.SpvStorageClassUniform:
		return "Uniform"
	case C.SpvStorageClassOutput:
		return "Output"
	case C.SpvStorageClassWorkgroup:
		return "Workgroup"
	case C.SpvStorageClassCrossWorkgroup:
		return "CrossWorkgroup"
	case C.SpvStorageClassPrivate:
		return "Private"
	case C.SpvStorageClassFunction:
		return "Function"
	case C.SpvStorageClassGeneric:
		return "Generic"
	case C.SpvStorageClassPushConstant:
		return "PushConstant"
	case C.SpvStorageClassAtomicCounter:
		return "AtomicCounter"
	case C.SpvStorageClassImage:
		return "Image"
	case C.SpvStorageClassStorageBuffer:
		return "StorageBuffer"
	case C.SpvStorageClassCallableDataKHR:

	}
	return "Unknown"
}

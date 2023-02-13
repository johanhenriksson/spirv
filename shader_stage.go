package spirv

// #include "./spirv-reflect/spirv_reflect.h"
import "C"

type Stage C.SpvReflectShaderStageFlagBits

const (
	StageVertex   = Stage(C.SPV_REFLECT_SHADER_STAGE_VERTEX_BIT)
	StageFragment = Stage(C.SPV_REFLECT_SHADER_STAGE_FRAGMENT_BIT)
)

func (s Stage) String() string {
	switch s {
	case StageVertex:
		return "Vertex"
	case StageFragment:
		return "Fragment"
	}
	panic("unmapped stage")
}

package spirv

/*
#include "./spirv-reflect/spirv_reflect.h"
#include "./spirv-reflect/spirv_reflect.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type ShaderModule struct {
	ptr *C.SpvReflectShaderModule
}

func checkResult(result C.SpvReflectResult) error {
	if result != C.SPV_REFLECT_RESULT_SUCCESS {
		return fmt.Errorf("failed to create shader module")
	}
	return nil
}

func FromSource(source []byte) (*ShaderModule, error) {
	spirv_bytes := C.ulong(len(source))
	spirv_ptr := unsafe.Pointer(&source[0])
	var ptr C.SpvReflectShaderModule
	res := C.spvReflectCreateShaderModule(spirv_bytes, spirv_ptr, &ptr)
	if err := checkResult(res); err != nil {
		return nil, err
	}
	return &ShaderModule{
		ptr: &ptr,
	}, nil
}

func (sm *ShaderModule) EnumerateInputVariables() ([]Input, error) {
	var count C.uint32_t
	res := C.spvReflectEnumerateInputVariables(sm.ptr, &count, nil)
	if err := checkResult(res); err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	inputs := make([]*C.SpvReflectInterfaceVariable, count)
	res = C.spvReflectEnumerateInputVariables(sm.ptr, &count, &inputs[0])
	if err := checkResult(res); err != nil {
		return nil, err
	}

	parsed := make([]Input, 0, len(inputs))
	for _, input := range inputs {
		if input.location == 0xFFFFFFFF {
			continue
		}
		name := C.GoString(input.name)
		parsed = append(parsed, Input{
			Name:         name,
			Location:     int(input.location),
			StorageClass: int(input.storage_class),
			TypeName:     C.GoString(input.type_description.type_name),
			TypeFlags:    int(input.type_description.type_flags),
		})
	}
	return parsed, nil
}

func (sm *ShaderModule) EnumerateDescriptorSets() ([]DescriptorSet, error) {
	var count C.uint32_t
	res := C.spvReflectEnumerateDescriptorSets(sm.ptr, &count, nil)
	if err := checkResult(res); err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	descriptors := make([]*C.SpvReflectDescriptorSet, count)
	res = C.spvReflectEnumerateDescriptorSets(sm.ptr, &count, &descriptors[0])
	if err := checkResult(res); err != nil {
		return nil, err
	}

	parsed := make([]DescriptorSet, 0, len(descriptors))
	for _, desc := range descriptors {
		bindingSlice := PtrToSlice[*C.SpvReflectDescriptorBinding](
			unsafe.Pointer(desc.bindings), int(desc.binding_count), int(desc.binding_count))

		bindings := make([]DescriptorBinding, 0, desc.binding_count)
		for _, binding := range bindingSlice {
			bindings = append(bindings, DescriptorBinding{
				Name:    C.GoString(binding.name),
				Binding: int(binding.binding),
			})
		}

		parsed = append(parsed, DescriptorSet{
			Set:      int(desc.set),
			Bindings: bindings,
		})
	}

	return parsed, nil
}

func (sm *ShaderModule) Destroy() {
	C.spvReflectDestroyShaderModule(sm.ptr)
	sm.ptr = nil
}

func PtrToSlice[T any](ptr unsafe.Pointer, len int, cap int) []T {
	var sl = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr, len, cap}
	return *(*[]T)(unsafe.Pointer(&sl))
}

package spirv

/*
#include "./spirv-reflect/spirv_reflect.h"
#include "./spirv-reflect/spirv_reflect.c"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type ShaderModule struct {
	ptr *C.SpvReflectShaderModule

	EntryPoint  EntryPoint
	EntryPoints []EntryPoint
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

	shader := &ShaderModule{
		ptr: &ptr,
		EntryPoint: EntryPoint{
			ID:    int(ptr.entry_point_id),
			Name:  C.GoString(ptr.entry_point_name),
			Stage: Stage(ptr.shader_stage),
		},
	}
	runtime.SetFinalizer(shader, func(shader *ShaderModule) {
		C.spvReflectDestroyShaderModule(shader.ptr)
	})

	if ptr.entry_point_count > 0 {
		// enumerate entry points
		shader.EntryPoints = make([]EntryPoint, 0, ptr.entry_point_count)
		entrypointSlice := ptrToSlice[C.SpvReflectEntryPoint](
			unsafe.Pointer(ptr.entry_points), int(ptr.entry_point_count), int(ptr.entry_point_count))
		for _, entry := range entrypointSlice {
			shader.EntryPoints = append(shader.EntryPoints, EntryPoint{
				ID:    int(entry.id),
				Name:  C.GoString(entry.name),
				Stage: Stage(entry.shader_stage),
			})
		}
	}

	return shader, nil
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
			StorageClass: StorageClass(input.storage_class),
			Type:         parseTypeDescription(input.type_description),
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
		bindingSlice := ptrToSlice[*C.SpvReflectDescriptorBinding](
			unsafe.Pointer(desc.bindings), int(desc.binding_count), int(desc.binding_count))

		bindings := make([]DescriptorBinding, 0, desc.binding_count)
		for _, binding := range bindingSlice {
			bindings = append(bindings, DescriptorBinding{
				Name:    C.GoString(binding.name),
				Binding: int(binding.binding),
				Type:    parseTypeDescription(binding.type_description),
			})
		}

		parsed = append(parsed, DescriptorSet{
			Set:      int(desc.set),
			Bindings: bindings,
		})
	}

	return parsed, nil
}

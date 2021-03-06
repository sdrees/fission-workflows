package builtin

import (
	"testing"

	"github.com/fission/fission-workflows/pkg/types"
)

func TestFunctionNoopPassInput(t *testing.T) {
	expected := "noopnoop"
	internalFunctionTest(t,
		&FunctionNoop{},
		&types.TaskInvocationSpec{
			Inputs: map[string]*types.TypedValue{
				NOOP_INPUT: parseUnsafe(expected),
			},
		},
		expected)
}

func TestFunctionNoopEmpty(t *testing.T) {
	internalFunctionTest(t,
		&FunctionNoop{},
		&types.TaskInvocationSpec{
			Inputs: map[string]*types.TypedValue{},
		},
		nil)
}

func TestFunctionNoopObject(t *testing.T) {
	internalFunctionTest(t,
		&FunctionNoop{},
		&types.TaskInvocationSpec{
			Inputs: map[string]*types.TypedValue{
				"foo":     parseUnsafe(true),
				"bar":     parseUnsafe(false),
				"default": parseUnsafe("hello"),
			},
		},
		"hello")
}

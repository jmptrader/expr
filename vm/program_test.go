package vm_test

import (
	"strings"
	"testing"

	"github.com/antonmedv/expr/vm"
)

func TestProgram_Disassemble(t *testing.T) {
	for op := vm.OpPush; op < vm.OpEnd; op++ {
		program := vm.Program{
			Constants: []interface{}{true},
			Bytecode:  []vm.Opcode{op},
			Arguments: []int{0},
		}
		d := program.Disassemble()
		if strings.Contains(d, "\t0x") {
			t.Errorf("cannot disassemble all opcodes")
		}
	}
}

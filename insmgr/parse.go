package insmgr

import (
	"bufio"
	"io"

	"github.com/v2fly/VSign/instructions"
)

func ReadAllIns(reader io.Reader) []instructions.Instruction {
	ins := make([]instructions.Instruction, 0, 1024)
	bufreader := bufio.NewReader(reader)
	for {
		s, err := bufreader.ReadString('\n')
		if err == io.EOF {
			return ins
		}
		s = s[:len(s)-len("\n")]
		insw := instructions.UnpackInstruction(s)
		ins = append(ins, insw)
	}

}

func YieldAll(mgr InstructionMgr, ins []instructions.Instruction) {
	for _, w := range ins {
		mgr.SubmitIns(w)
	}
}

func SortAll(r io.Reader, w io.Writer) {
	ins := ReadAllIns(r)
	SortIns(ins)
	YieldAll(NewOutputInsMgr(w), ins)
}

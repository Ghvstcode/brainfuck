package vm

import "io"

//Machine is the core of the interpreter
//it defines various components of the brainfuck(bf) language
type Machine struct {
	//code is the bf code which could be written by a bf programmer
	code string
	// instruction pointer specifies which code should be executed next
	ip int
	//In BF you have an array of 30k 1byte memory blocks
	memory [30000]int
	//Data pointer points to the memory cell from above. It holds the index, that index is usually zero based
	dp     int
	//YKTV, bf has "." & "," for reading and writing to STD
	input  io.Reader
	output io.Writer
}

//Get a new instance of the machine
func NewMachine(c string, in io.Reader, out io.Writer) *Machine{
	return &Machine{
		code:    c,
		input:   in,
		output:  out,
	}
}

func (m *Machine) Run() {
	//Loop for as long as the instruction pointer is less than the length of the code written
	for m.ip < len(m.code) {

		switch m.code[m.ip] {
		case '+':
			m.memory[m.dp]++
		case '-':
			m.memory[m.dp]--
		case '>':
			m.dp++
		case '<':
			m.dp--
		}

		m.ip++
	}
}

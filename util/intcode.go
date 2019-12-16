package util

import (
	"strconv"
	"strings"
)

var instructions = map[int]instruction{
	1:  add,
	2:  multiply,
	3:  saveInput,
	4:  output,
	5:  jumpIfTrue,
	6:  jumpIfFalse,
	7:  lessThan,
	8:  equals,
	9:  relativeBaseOffset,
	99: halt,
}

var numberOfArguments = map[int]int{
	1:  3,
	2:  3,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	9:  1,
	99: 0,
}

type Computer struct {
	program        [] int
	inputs         [] int
	pc             int
	output         int
	halted         bool
	outputProduced bool
	relativeBase   int
	memory         map[int]int
}

func NewComputer(program []int, inputs []int) *Computer {
	return &Computer{
		program:        program,
		inputs:         inputs,
		pc:             0,
		output:         0,
		halted:         false,
		outputProduced: false,
		relativeBase:   0,
		memory:         make(map[int]int),
	}
}

type arg struct {
	address int
	value   int
}

type instruction func(comp *Computer, args []arg) int

func add(c *Computer, args []arg) int {
	c.set(args[2].address, args[0].value+args[1].value)
	return c.pc + 4
}

func multiply(c *Computer, args []arg) int {
	c.set(args[2].address, args[0].value*args[1].value)
	return c.pc + 4
}

func saveInput(c *Computer, args []arg) int {
	input := c.getInput()
	c.set(args[0].address, input)
	return c.pc + 2
}

func output(c *Computer, args []arg) int {
	c.output = args[0].value
	c.outputProduced = true
	return c.pc + 2
}

func jumpIfTrue(c *Computer, args []arg) int {
	if args[0].value != 0 {
		return args[1].value
	}
	return c.pc + 3
}

func jumpIfFalse(c *Computer, args []arg) int {
	if args[0].value == 0 {
		return args[1].value
	}
	return c.pc + 3
}

func lessThan(c *Computer, args []arg) int {
	if args[0].value < args[1].value {
		c.set(args[2].address, 1)
	} else {
		c.set(args[2].address, 0)
	}
	return c.pc + 4
}

func equals(c *Computer, args []arg) int {
	if args[0].value == args[1].value {
		c.set(args[2].address, 1)
	} else {
		c.set(args[2].address, 0)
	}
	return c.pc + 4
}

func relativeBaseOffset(c *Computer, args []arg) int {
	c.relativeBase += args[0].value
	return c.pc + 2
}

func halt(c *Computer, args []arg) int {
	c.halted = true
	return c.pc
}

func (c *Computer) Run() int {
	c.outputProduced = false
	for !c.halted && !c.outputProduced {
		opcode := c.program[c.pc] % 100
		instruction := instructions[opcode]
		args := c.parseInstruction(numberOfArguments[opcode])
		c.pc = instruction(c, args)
	}
	return c.output
}

func (c *Computer) Get(address int) int {
	if address < len(c.program) {
		return c.program[address]
	}
	if val, ok := c.memory[c.relativeBase]; ok {
		return val
	}
	return 0
}

func (c *Computer) Output() int {
	return c.output
}

func (c *Computer) Input(in int) {
	//prepend input
	c.inputs = append([]int{in}, c.inputs...)
}

func (c *Computer) IsHalted() bool {
	return c.halted
}

func (c *Computer) set(address, value int) {
	if address < len(c.program) {
		c.program[address] = value
	} else {
		c.memory[address] = value
	}
}

func (c *Computer) getInput() int {
	//pops input
	var val int
	val, c.inputs = c.inputs[len(c.inputs)-1], c.inputs[:len(c.inputs)-1]
	return val
}

func (c *Computer) parseInstruction(numArgs int) []arg {
	s := strconv.Itoa(c.program[c.pc])

	// 2->00002, 102->00102
	instructionStr := strings.Repeat("0", 5-len(s)) + s

	inst := Atoi(instructionStr[3:])
	modes := Reverse(instructionStr[0:3])

	var args [] arg
	for i := 0; i < numberOfArguments[inst]; i++ {
		val := c.program[c.pc+i+1]
		switch modes[i : i+1] {
		case "0":
			args = append(args, arg{address: val, value: c.readAddress(val)})
		case "1":
			args = append(args, arg{address: val, value: val})
		case "2":
			val += c.relativeBase
			args = append(args, arg{address: val, value: c.readAddress(val)})
		}
	}
	return args
}

func (c *Computer) readAddress(address int) int {
	if address < len(c.program) {
		return c.program[address]
	}
	if val, ok := c.memory[address]; ok {
		return val
	}
	return 0
}

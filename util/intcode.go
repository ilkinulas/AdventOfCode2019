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
	99: 0,
}

type Computer struct {
	program        [] int
	inputs         [] int
	pc             int
	output         int
	halted         bool
	outputProduced bool
}

func NewComputer(program []int, inputs []int) *Computer {
	return &Computer{
		program:        program,
		inputs:         inputs,
		pc:             0,
		output:         0,
		halted:         false,
		outputProduced: false,
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
	return c.program[address]
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
	c.program[address] = value
}

func (c *Computer) getInput() int {
	//pops input
	var val int
	val, c.inputs = c.inputs[len(c.inputs)-1], c.inputs[:len(c.inputs)-1]
	return val
}

func (c *Computer) SILBUNU() []int {
	return c.inputs
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
		if modes[i:i+1] == "1" {
			//immediate mode
			args = append(args, arg{address: val, value: val})
		} else {
			args = append(args, arg{address: val, value: c.program[val]})
		}
	}
	return args
}

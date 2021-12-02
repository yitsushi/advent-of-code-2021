package day02

// Command is a command.
type Command string

const (
	// ForwardCommand moves forward.
	ForwardCommand Command = "forward"
	// DownCommand moves down.
	DownCommand Command = "down"
	// UpCommand moves up.
	UpCommand Command = "up"
	// NoCommand does nothing.
	NoCommand Command = ""
)

// Instruction represents one instruction in the planned course.
type Instruction struct {
	Command Command
	Value   int64
}

// NewInstruction creates a new NoCommand instruction.
func NewInstruction() Instruction {
	return Instruction{
		Command: NoCommand,
		Value:   0,
	}
}

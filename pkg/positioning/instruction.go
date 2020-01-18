package positioning

const (
	left    = "L"
	right   = "R"
	forward = "M"
)

type Instruction string

// AllowedInstructions can be left, right and forward.
var AllowedInstructions = []Instruction{
	left,
	right,
	forward,
}

// IsInstructionAllowed checks if an instruction specified is present in the list of available instructions.
// If the instruction is not there in the allowed list, return false . If the supplied instruction is an allowed one
// then true is returned.
func IsInstructionAllowed(instruction Instruction) bool {
	for _, allowedInstruction := range AllowedInstructions {
		if allowedInstruction == instruction {
			return true
		}
	}
	return false
}

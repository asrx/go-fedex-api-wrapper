package SimpleType

type SpecialInstructionsStatusCode string

const (
	SpecialInstructionsStatusCodeACCEPTED SpecialInstructionsStatusCode = "ACCEPTED"

	SpecialInstructionsStatusCodeCANCELLED SpecialInstructionsStatusCode = "CANCELLED"

	SpecialInstructionsStatusCodeDENIED SpecialInstructionsStatusCode = "DENIED"

	SpecialInstructionsStatusCodeHELD SpecialInstructionsStatusCode = "HELD"

	SpecialInstructionsStatusCodeMODIFIED SpecialInstructionsStatusCode = "MODIFIED"

	SpecialInstructionsStatusCodeRELINQUISHED SpecialInstructionsStatusCode = "RELINQUISHED"

	SpecialInstructionsStatusCodeREQUESTED SpecialInstructionsStatusCode = "REQUESTED"

	SpecialInstructionsStatusCodeSET SpecialInstructionsStatusCode = "SET"
)

package sdkcommon

type Opcode string

const (
	OpcodeNone           Opcode = "None"
	OpcodeStep           Opcode = "Step"
	OpcodeStepRun        Opcode = "StepRun"
	OpcodeStepError      Opcode = "StepError"
	OpcodeStepPlanned    Opcode = "StepPlanned"
	OpcodeSleep          Opcode = "Sleep"
	OpcodeWaitForEvent   Opcode = "WaitForEvent"
	OpcodeInvokeFunction Opcode = "InvokeFunction"
)

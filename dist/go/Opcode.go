package sdkcommon

type Opcode string

const (
	OpcodeInvokefunction Opcode = "InvokeFunction"
	OpcodeSleep          Opcode = "Sleep"
	OpcodeSteperror      Opcode = "StepError"
	OpcodeStepplanned    Opcode = "StepPlanned"
	OpcodeSteprun        Opcode = "StepRun"
	OpcodeWaitforevent   Opcode = "WaitForEvent"
)

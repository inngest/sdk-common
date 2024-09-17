import enum


class Opcode(enum.Enum):
    INVOKE_FUNCTION = "InvokeFunction"
    SLEEP = "Sleep"
    STEP_ERROR = "StepError"
    STEP_PLANNED = "StepPlanned"
    STEP_RUN = "StepRun"
    WAIT_FOR_EVENT = "WaitForEvent"

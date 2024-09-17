import enum


class Opcode(enum.Enum):
    NONE = "None"
    STEP = "Step"
    STEP_RUN = "StepRun"
    STEP_ERROR = "StepError"
    STEP_PLANNED = "StepPlanned"
    SLEEP = "Sleep"
    WAIT_FOR_EVENT = "WaitForEvent"
    INVOKE_FUNCTION = "InvokeFunction"

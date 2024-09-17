export const Opcode = {
  None: "None",
  Step: "Step",
  StepRun: "StepRun",
  StepError: "StepError",
  StepPlanned: "StepPlanned",
  Sleep: "Sleep",
  WaitForEvent: "WaitForEvent",
  InvokeFunction: "InvokeFunction",
} as const;

export type Opcode = (typeof Opcode)[keyof typeof Opcode];

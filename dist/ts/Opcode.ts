export const Opcode = {
  Invokefunction: "InvokeFunction",
  Sleep: "Sleep",
  Steperror: "StepError",
  Stepplanned: "StepPlanned",
  Steprun: "StepRun",
  Waitforevent: "WaitForEvent",
} as const;

export type Opcode = (typeof Opcode)[keyof typeof Opcode];

export const InternalEvent = {
  FunctionFailed: "function.failed",
} as const;

export type InternalEvent = (typeof InternalEvent)[keyof typeof InternalEvent];

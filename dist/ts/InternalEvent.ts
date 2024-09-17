export const InternalEvent = {
  None: "none",
  FunctionFailed: "inngest/function.failed",
} as const;

export type InternalEvent = (typeof InternalEvent)[keyof typeof InternalEvent];

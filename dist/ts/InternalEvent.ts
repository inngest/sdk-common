export const InternalEvent = {
  FunctionFailed: "inngest/function.failed",
} as const;

export type InternalEvent = (typeof InternalEvent)[keyof typeof InternalEvent];

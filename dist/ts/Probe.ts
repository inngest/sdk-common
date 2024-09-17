export const Probe = {
  None: "none",
  Trust: "trust",
} as const;

export type Probe = (typeof Probe)[keyof typeof Probe];

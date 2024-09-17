export const Probe = {
  Trust: "trust",
} as const;

export type Probe = (typeof Probe)[keyof typeof Probe];

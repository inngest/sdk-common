export const ServerKind = {
  None: "none",
  Cloud: "cloud",
  Dev: "dev",
} as const;

export type ServerKind = (typeof ServerKind)[keyof typeof ServerKind];

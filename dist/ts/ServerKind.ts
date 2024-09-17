export const ServerKind = {
  Cloud: "cloud",
  Dev: "dev",
} as const;

export type ServerKind = (typeof ServerKind)[keyof typeof ServerKind];

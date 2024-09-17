export const Mode = {
  None: "none",
  Cloud: "cloud",
  Dev: "dev",
} as const;

export type Mode = (typeof Mode)[keyof typeof Mode];

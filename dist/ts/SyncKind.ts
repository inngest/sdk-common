export const SyncKind = {
  None: "none",
  OutOfBand: "out_of_band",
  InBand: "in_band",
} as const;

export type SyncKind = (typeof SyncKind)[keyof typeof SyncKind];

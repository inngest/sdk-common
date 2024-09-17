export const SyncKind = {
  InBand: "in_band",
  OutOfBand: "out_of_band",
} as const;

export type SyncKind = (typeof SyncKind)[keyof typeof SyncKind];

export const QueryParamKey = {
  Deployid: "deployId",
  Fnid: "fnId",
  Probe: "probe",
  Stepid: "stepId",
} as const;

export type QueryParamKey = (typeof QueryParamKey)[keyof typeof QueryParamKey];

export const QueryParamKey = {
  None: "none",
  DeployId: "deployId",
  FnId: "fnId",
  Probe: "probe",
  StepId: "stepId",
} as const;

export type QueryParamKey = (typeof QueryParamKey)[keyof typeof QueryParamKey];

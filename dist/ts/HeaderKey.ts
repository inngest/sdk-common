export const HeaderKey = {
  None: "none",
  Authorization: "authorization",
  ContentType: "content-type",
  RetryAfter: "retry-after",
  ServerTiming: "server-timing",
  UserAgent: "user-agent",
  InngestEnv: "x-inngest-env",
  InngestExpectedServerKind: "x-inngest-expected-server-kind",
  InngestFramework: "x-inngest-framework",
  InngestNoRetry: "x-inngest-no-retry",
  InngestReqVersion: "x-inngest-req-version",
  InngestSdk: "x-inngest-sdk",
  InngestServerKind: "x-inngest-server-kind",
  InngestSignature: "x-inngest-signature",
  InngestSyncKind: "x-inngest-sync-kind",
} as const;

export type HeaderKey = (typeof HeaderKey)[keyof typeof HeaderKey];

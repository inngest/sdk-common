export const HeaderKey = {
  Authorization: "authorization",
  ContentType: "content-type",
  RetryAfter: "retry-after",
  ServerTiming: "server-timing",
  UserAgent: "user-agent",
  InngestEnv: "inngest-env",
  InngestExpectedServerKind: "inngest-expected-server-kind",
  InngestFramework: "inngest-framework",
  InngestNoRetry: "inngest-no-retry",
  InngestReqVersion: "inngest-req-version",
  InngestSdk: "inngest-sdk",
  InngestServerKind: "inngest-server-kind",
  InngestSignature: "inngest-signature",
  InngestSyncKind: "inngest-sync-kind",
} as const;

export type HeaderKey = (typeof HeaderKey)[keyof typeof HeaderKey];

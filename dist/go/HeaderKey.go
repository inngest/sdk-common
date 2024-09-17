package sdkcommon

type HeaderKey string

const (
	HeaderKeyNone                      HeaderKey = "none"
	HeaderKeyAuthorization             HeaderKey = "authorization"
	HeaderKeyContentType               HeaderKey = "content-type"
	HeaderKeyRetryAfter                HeaderKey = "retry-after"
	HeaderKeyServerTiming              HeaderKey = "server-timing"
	HeaderKeyUserAgent                 HeaderKey = "user-agent"
	HeaderKeyInngestEnv                HeaderKey = "x-inngest-env"
	HeaderKeyInngestExpectedServerKind HeaderKey = "x-inngest-expected-server-kind"
	HeaderKeyInngestFramework          HeaderKey = "x-inngest-framework"
	HeaderKeyInngestNoRetry            HeaderKey = "x-inngest-no-retry"
	HeaderKeyInngestReqVersion         HeaderKey = "x-inngest-req-version"
	HeaderKeyInngestSdk                HeaderKey = "x-inngest-sdk"
	HeaderKeyInngestServerKind         HeaderKey = "x-inngest-server-kind"
	HeaderKeyInngestSignature          HeaderKey = "x-inngest-signature"
	HeaderKeyInngestSyncKind           HeaderKey = "x-inngest-sync-kind"
)

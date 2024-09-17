package sdkcommon

type HeaderKey string

const (
	HeaderKeyAuthorization             HeaderKey = "authorization"
	HeaderKeyContentType               HeaderKey = "content-type"
	HeaderKeyRetryAfter                HeaderKey = "retry-after"
	HeaderKeyServerTiming              HeaderKey = "server-timing"
	HeaderKeyUserAgent                 HeaderKey = "user-agent"
	HeaderKeyInngestEnv                HeaderKey = "inngest-env"
	HeaderKeyInngestExpectedServerKind HeaderKey = "inngest-expected-server-kind"
	HeaderKeyInngestFramework          HeaderKey = "inngest-framework"
	HeaderKeyInngestNoRetry            HeaderKey = "inngest-no-retry"
	HeaderKeyInngestReqVersion         HeaderKey = "inngest-req-version"
	HeaderKeyInngestSdk                HeaderKey = "inngest-sdk"
	HeaderKeyInngestServerKind         HeaderKey = "inngest-server-kind"
	HeaderKeyInngestSignature          HeaderKey = "inngest-signature"
	HeaderKeyInngestSyncKind           HeaderKey = "inngest-sync-kind"
)

import enum


class HeaderKey(enum.Enum):
    NONE = "none"
    AUTHORIZATION = "authorization"
    CONTENT_TYPE = "content-type"
    RETRY_AFTER = "retry-after"
    SERVER_TIMING = "server-timing"
    USER_AGENT = "user-agent"
    INNGEST_ENV = "x-inngest-env"
    INNGEST_EXPECTED_SERVER_KIND = "x-inngest-expected-server-kind"
    INNGEST_FRAMEWORK = "x-inngest-framework"
    INNGEST_NO_RETRY = "x-inngest-no-retry"
    INNGEST_REQ_VERSION = "x-inngest-req-version"
    INNGEST_SDK = "x-inngest-sdk"
    INNGEST_SERVER_KIND = "x-inngest-server-kind"
    INNGEST_SIGNATURE = "x-inngest-signature"
    INNGEST_SYNC_KIND = "x-inngest-sync-kind"

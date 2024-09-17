import enum


class HeaderKey(enum.Enum):
    AUTHORIZATION = "authorization"
    CONTENT_TYPE = "content-type"
    RETRY_AFTER = "retry-after"
    SERVER_TIMING = "server-timing"
    USER_AGENT = "user-agent"
    INNGEST_ENV = "inngest-env"
    INNGEST_EXPECTED_SERVER_KIND = "inngest-expected-server-kind"
    INNGEST_FRAMEWORK = "inngest-framework"
    INNGEST_NO_RETRY = "inngest-no-retry"
    INNGEST_REQ_VERSION = "inngest-req-version"
    INNGEST_SDK = "inngest-sdk"
    INNGEST_SERVER_KIND = "inngest-server-kind"
    INNGEST_SIGNATURE = "inngest-signature"
    INNGEST_SYNC_KIND = "inngest-sync-kind"

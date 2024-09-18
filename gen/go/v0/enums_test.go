package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderValues(t *testing.T) {
	tests := []struct {
		enum  Header
		value string
	}{
		{
			enum:  Header_HEADER_NONE,
			value: "none",
		},
		{
			enum:  Header_HEADER_AUTHORIZATION,
			value: "authorization",
		},
		{
			enum:  Header_HEADER_CONTENT_TYPE,
			value: "content-type",
		},
		{
			enum:  Header_HEADER_RETRY_AFTER,
			value: "retry-after",
		},
		{
			enum:  Header_HEADER_SERVER_TIMING,
			value: "server-timing",
		},
		{
			enum:  Header_HEADER_USER_AGENT,
			value: "user-agent",
		},
		{
			enum:  Header_HEADER_X_INNGEST_ENV,
			value: "x-inngest-env",
		},
		{
			enum:  Header_HEADER_X_INNGEST_EXPECTED_SERVER_KIND,
			value: "x-inngest-expected-server-kind",
		},
		{
			enum:  Header_HEADER_X_INNGEST_FRAMEWORK,
			value: "x-inngest-framework",
		},
		{
			enum:  Header_HEADER_X_INNGEST_NO_RETRY,
			value: "x-inngest-no-retry",
		},
		{
			enum:  Header_HEADER_X_INNGEST_REQ_VERSION,
			value: "x-inngest-req-version",
		},
		{
			enum:  Header_HEADER_X_INNGEST_SDK,
			value: "x-inngest-sdk",
		},
		{
			enum:  Header_HEADER_X_INNGEST_SERVER_KIND,
			value: "x-inngest-server-kind",
		},
		{
			enum:  Header_HEADER_X_INNGEST_SIGNATURE,
			value: "x-inngest-signature",
		},
		{
			enum:  Header_HEADER_X_INNGEST_SYNC_KIND,
			value: "x-inngest-sync-kind",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestModeValues(t *testing.T) {
	tests := []struct {
		enum  Mode
		value string
	}{
		{
			enum:  Mode_MODE_NONE,
			value: "none",
		},
		{
			enum:  Mode_MODE_CLOUD,
			value: "cloud",
		},
		{
			enum:  Mode_MODE_DEV,
			value: "dev",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

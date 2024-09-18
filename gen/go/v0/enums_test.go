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

func TestOpcodeValues(t *testing.T) {
	tests := []struct {
		enum  Opcode
		value string
	}{
		{
			enum:  Opcode_OPCODE_NONE,
			value: "None",
		},
		{
			enum:  Opcode_OPCODE_STEP,
			value: "Step",
		},
		{
			enum:  Opcode_OPCODE_STEP_RUN,
			value: "StepRun",
		},
		{
			enum:  Opcode_OPCODE_STEP_ERROR,
			value: "StepError",
		},
		{
			enum:  Opcode_OPCODE_STEP_PLANNED,
			value: "StepPlanned",
		},
		{
			enum:  Opcode_OPCODE_SLEEP,
			value: "Sleep",
		},
		{
			enum:  Opcode_OPCODE_WAIT_FOR_EVENT,
			value: "WaitForEvent",
		},
		{
			enum:  Opcode_OPCODE_INVOKE_FUNCTION,
			value: "InvokeFunction",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestInternalEventValues(t *testing.T) {
	tests := []struct {
		enum  InternalEvent
		value string
	}{
		{
			enum:  InternalEvent_INTERNAL_EVENT_NONE,
			value: "none",
		},
		{
			enum:  InternalEvent_INTERNAL_EVENT_FUNCTION_FAILED,
			value: "inngest/function.failed",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestProbeValues(t *testing.T) {
	tests := []struct {
		enum  Probe
		value string
	}{
		{
			enum:  Probe_PROBE_NONE,
			value: "none",
		},
		{
			enum:  Probe_PROBE_TRUST,
			value: "trust",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestQueryParamValues(t *testing.T) {
	tests := []struct {
		enum  QueryParam
		value string
	}{
		{
			enum:  QueryParam_QUERY_PARAM_NONE,
			value: "none",
		},
		{
			enum:  QueryParam_QUERY_PARAM_DEPLOY_ID,
			value: "deployId",
		},
		{
			enum:  QueryParam_QUERY_PARAM_FN_ID,
			value: "fnId",
		},
		{
			enum:  QueryParam_QUERY_PARAM_PROBE,
			value: "probe",
		},
		{
			enum:  QueryParam_QUERY_PARAM_STEP_ID,
			value: "stepId",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestServerKindValues(t *testing.T) {
	tests := []struct {
		enum  ServerKind
		value string
	}{
		{
			enum:  ServerKind_SERVER_KIND_NONE,
			value: "none",
		},
		{
			enum:  ServerKind_SERVER_KIND_CLOUD,
			value: "cloud",
		},
		{
			enum:  ServerKind_SERVER_KIND_DEV,
			value: "dev",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

func TestSyncKindValues(t *testing.T) {
	tests := []struct {
		enum  SyncKind
		value string
	}{
		{
			enum:  SyncKind_SYNC_KIND_NONE,
			value: "none",
		},
		{
			enum:  SyncKind_SYNC_KIND_OUT_OF_BAND,
			value: "out_of_band",
		},
		{
			enum:  SyncKind_SYNC_KIND_IN_BAND,
			value: "in_band",
		},
	}

	for _, test := range tests {
		t.Run(test.enum.String(), func(t *testing.T) {
			assert.Equal(t, test.value, test.enum.Str())
		})
	}
}

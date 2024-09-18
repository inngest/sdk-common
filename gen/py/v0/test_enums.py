#!/usr/bin/env python3

from gen.py.v0 import enums_augmented
from gen.py.v0 import enums_pb2 as enums


def test_header():
    assert enums.Header.str(enums.HEADER_NONE) == "none"
    assert enums.Header.str(enums.HEADER_AUTHORIZATION) == "authorization"
    assert enums.Header.str(enums.HEADER_CONTENT_TYPE) == "content-type"
    assert enums.Header.str(enums.HEADER_RETRY_AFTER) == "retry-after"
    assert enums.Header.str(enums.HEADER_SERVER_TIMING) == "server-timing"
    assert enums.Header.str(enums.HEADER_USER_AGENT) == "user-agent"
    assert (
        enums.Header.str(enums.HEADER_X_INNGEST_EXPECTED_SERVER_KIND)
        == "x-inngest-expected-server-kind"
    )
    assert enums.Header.str(enums.HEADER_X_INNGEST_FRAMEWORK) == "x-inngest-framework"
    assert enums.Header.str(enums.HEADER_X_INNGEST_NO_RETRY) == "x-inngest-no-retry"
    assert (
        enums.Header.str(enums.HEADER_X_INNGEST_REQ_VERSION) == "x-inngest-req-version"
    )
    assert enums.Header.str(enums.HEADER_X_INNGEST_SDK) == "x-inngest-sdk"
    assert (
        enums.Header.str(enums.HEADER_X_INNGEST_SERVER_KIND) == "x-inngest-server-kind"
    )
    assert enums.Header.str(enums.HEADER_X_INNGEST_SIGNATURE) == "x-inngest-signature"
    assert enums.Header.str(enums.HEADER_X_INNGEST_SYNC_KIND) == "x-inngest-sync-kind"


def test_internal_event():
    assert enums.InternalEvent.str(enums.INTERNAL_EVENT_NONE) == "none"
    assert (
        enums.InternalEvent.str(enums.INTERNAL_EVENT_FUNCTION_FAILED)
        == "inngest/function.failed"
    )


def test_mode():
    assert enums.Mode.str(enums.MODE_NONE) == "none"
    assert enums.Mode.str(enums.MODE_CLOUD) == "cloud"
    assert enums.Mode.str(enums.MODE_DEV) == "dev"


def test_opcode():
    assert enums.Opcode.str(enums.OPCODE_NONE) == "None"
    assert enums.Opcode.str(enums.OPCODE_STEP) == "Step"
    assert enums.Opcode.str(enums.OPCODE_STEP_RUN) == "StepRun"
    assert enums.Opcode.str(enums.OPCODE_STEP_ERROR) == "StepError"
    assert enums.Opcode.str(enums.OPCODE_STEP_PLANNED) == "StepPlanned"
    assert enums.Opcode.str(enums.OPCODE_SLEEP) == "Sleep"
    assert enums.Opcode.str(enums.OPCODE_WAIT_FOR_EVENT) == "WaitForEvent"
    assert enums.Opcode.str(enums.OPCODE_INVOKE_FUNCTION) == "InvokeFunction"


def test_probe():
    assert enums.Probe.str(enums.PROBE_NONE) == "none"
    assert enums.Probe.str(enums.PROBE_TRUST) == "trust"


def test_query_param():
    assert enums.QueryParam.str(enums.QUERY_PARAM_NONE) == "none"
    assert enums.QueryParam.str(enums.QUERY_PARAM_DEPLOY_ID) == "deployId"
    assert enums.QueryParam.str(enums.QUERY_PARAM_FN_ID) == "fnId"
    assert enums.QueryParam.str(enums.QUERY_PARAM_PROBE) == "probe"
    assert enums.QueryParam.str(enums.QUERY_PARAM_STEP_ID) == "stepId"


def test_server_kind():
    assert enums.ServerKind.str(enums.SERVER_KIND_NONE) == "none"
    assert enums.ServerKind.str(enums.SERVER_KIND_CLOUD) == "cloud"
    assert enums.ServerKind.str(enums.SERVER_KIND_DEV) == "dev"


def test_sync_kind():
    assert enums.SyncKind.str(enums.SYNC_KIND_NONE) == "none"
    assert enums.SyncKind.str(enums.SYNC_KIND_OUT_OF_BAND) == "out_of_band"
    assert enums.SyncKind.str(enums.SYNC_KIND_IN_BAND) == "in_band"

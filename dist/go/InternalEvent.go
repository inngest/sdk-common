package sdkcommon

type InternalEvent string

const (
	InternalEventNone           InternalEvent = "none"
	InternalEventFunctionFailed InternalEvent = "inngest/function.failed"
)

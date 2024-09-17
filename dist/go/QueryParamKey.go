package sdkcommon

type QueryParamKey string

const (
	QueryParamKeyNone     QueryParamKey = "none"
	QueryParamKeyDeployId QueryParamKey = "deployId"
	QueryParamKeyFnId     QueryParamKey = "fnId"
	QueryParamKeyProbe    QueryParamKey = "probe"
	QueryParamKeyStepId   QueryParamKey = "stepId"
)

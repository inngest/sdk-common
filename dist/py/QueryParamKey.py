import enum


class QueryParamKey(enum.Enum):
    NONE = "none"
    DEPLOY_ID = "deployId"
    FN_ID = "fnId"
    PROBE = "probe"
    STEP_ID = "stepId"

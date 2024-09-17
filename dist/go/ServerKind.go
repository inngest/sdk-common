package sdkcommon

type ServerKind string

const (
	ServerKindNone  ServerKind = "none"
	ServerKindCloud ServerKind = "cloud"
	ServerKindDev   ServerKind = "dev"
)

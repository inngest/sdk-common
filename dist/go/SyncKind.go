package sdkcommon

type SyncKind string

const (
	SyncKindNone      SyncKind = "none"
	SyncKindOutOfBand SyncKind = "out_of_band"
	SyncKindInBand    SyncKind = "in_band"
)

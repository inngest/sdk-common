package sdkcommon

type SyncKind string

const (
	SyncKindInBand    SyncKind = "in_band"
	SyncKindOutOfBand SyncKind = "out_of_band"
)

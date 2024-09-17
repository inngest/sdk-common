import enum


class SyncKind(enum.Enum):
    NONE = "none"
    OUT_OF_BAND = "out_of_band"
    IN_BAND = "in_band"

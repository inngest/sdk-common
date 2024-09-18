from gen.py.v0 import enums_pb2 as enums


# Monkey patch and add a `str` method to enums classes for easier access of string value


# Header enum
def __header(v):
    return enums._HEADER.values_by_number[v].GetOptions().Extensions[enums.str]


enums.Header.str = __header


# InternalEvent enum
def __internal_event(v):
    return enums._INTERNALEVENT.values_by_number[v].GetOptions().Extensions[enums.str]


enums.InternalEvent.str = __internal_event


# Mode Enum
def __mode(v):
    return enums._MODE.values_by_number[v].GetOptions().Extensions[enums.str]


enums.Mode.str = __mode


# Opcode enum
def __opcode(v):
    return enums._OPCODE.values_by_number[v].GetOptions().Extensions[enums.str]


enums.Opcode.str = __opcode


# Probe enum
def __probe(v):
    return enums._PROBE.values_by_number[v].GetOptions().Extensions[enums.str]


enums.Probe.str = __probe


# QueryParam enum
def __query_param(v):
    return enums._QUERYPARAM.values_by_number[v].GetOptions().Extensions[enums.str]


enums.QueryParam.str = __query_param


# ServerKind enum
def __server_kind(v):
    return enums._SERVERKIND.values_by_number[v].GetOptions().Extensions[enums.str]


enums.ServerKind.str = __server_kind


# SyncKind enum
def __sync_kind(v):
    return enums._SYNCKIND.values_by_number[v].GetOptions().Extensions[enums.str]


enums.SyncKind.str = __sync_kind

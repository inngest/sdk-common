package sdk

import "google.golang.org/protobuf/proto"

func (h Header) Str() string {
	vals := h.Descriptor().Values()
	desc := vals.ByNumber(h.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

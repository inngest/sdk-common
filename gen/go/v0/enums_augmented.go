package sdk

import "google.golang.org/protobuf/proto"

func (h Header) Str() string {
	vals := h.Descriptor().Values()
	desc := vals.ByNumber(h.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (i InternalEvent) Str() string {
	vals := i.Descriptor().Values()
	desc := vals.ByNumber(i.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (m Mode) Str() string {
	vals := m.Descriptor().Values()
	desc := vals.ByNumber(m.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (op Opcode) Str() string {
	vals := op.Descriptor().Values()
	desc := vals.ByNumber(op.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (p Probe) Str() string {
	vals := p.Descriptor().Values()
	desc := vals.ByNumber(p.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (qp QueryParam) Str() string {
	vals := qp.Descriptor().Values()
	desc := vals.ByNumber(qp.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (sk ServerKind) Str() string {
	vals := sk.Descriptor().Values()
	desc := vals.ByNumber(sk.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

func (sk SyncKind) Str() string {
	vals := sk.Descriptor().Values()
	desc := vals.ByNumber(sk.Number())
	ext := proto.GetExtension(desc.Options(), E_Str)

	return ext.(string)
}

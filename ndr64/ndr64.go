package ndr64

type Encoder interface {
	Encode() []byte
}

type NDRState struct {
	refIDCtr      uint64
	deferredBytes []byte
	bytes         []byte
}

func NewNDRState() NDRState {
	return NDRState{refIDCtr: 1}
}

func (state *NDRState) Marshal(in Encoder) {
	state.bytes = append(state.bytes, in.Encode()...)
}

func (state *NDRState) MarshalDeferred(in Encoder) {
	// TODO: Actual refids
	state.bytes = append(state.bytes, EncodeUint(state.refIDCtr)...)
	state.deferredBytes = append(state.deferredBytes, in.Encode()...)
}

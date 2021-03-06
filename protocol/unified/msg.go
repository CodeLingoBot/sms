package unified

type MsgHead struct {
	Length uint32
	CmdID  uint32
	Seq    uint32
	Status uint32

	RecvConnID uint32
	SendConnID uint32
}

type MsgBindReq struct {
	MsgHead

	Account  string
	Password string
}

type MsgBindAck struct {
	MsgHead
	Succ bool
}

type MsgHeartBeat struct {
	MsgHead
}

type MsgSubmitReq struct {
	MsgHead

	SrcAddr string
	DstAddr string
	FeeAddr string

	LinkID string
}

type MsgSubmitAck struct {
	MsgID [8]byte
}

type MsgMultiSubmitReq struct {
}

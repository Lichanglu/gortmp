package rtmp

const (
	TIMESTAMP_MAX      = uint32(2000000000)
	TIMESTAMP_AUTO     = uint32(0)
	TIMESTAMP_EXTENDED = 0xFFFFFF
)

const (
	CHUNK_STREAM_ID_PROTOCOL     = uint32(2)
	CHUNK_STREAM_ID_COMMAND      = uint32(3)
	CHUNK_STREAM_ID_USER_CONTROL = uint32(4)
)

const (
	HEADER_FORMAT_FULL                   = 0x00
	HEADER_FORMAT_SAME_STREAM            = 0x01
	HEADER_FORMAT_SAME_LENGTH_AND_STREAM = 0x02
	HEADER_FORMAT_CONTINUATION           = 0x03
)

const (
	MESSAGE_TYPE_NONE               = 0x00
	MESSAGE_TYPE_CHUNK_SIZE         = 0x01
	MESSAGE_TYPE_ABORT              = 0x02
	MESSAGE_TYPE_ACK                = 0x03
	MESSAGE_TYPE_PING               = 0x04
	MESSAGE_TYPE_ACK_SIZE           = 0x05
	MESSAGE_TYPE_BANDWIDTH          = 0x06
	MESSAGE_TYPE_AUDIO              = 0x08
	MESSAGE_TYPE_VIDEO              = 0x09
	MESSAGE_TYPE_FLEX               = 0x0F
	MESSAGE_TYPE_AMF3_SHARED_OBJECT = 0x10
	MESSAGE_TYPE_AMF3               = 0x11
	MESSAGE_TYPE_INVOKE             = 0x12
	MESSAGE_TYPE_AMF0_SHARED_OBJECT = 0x13
	MESSAGE_TYPE_AMF0               = 0x14
	MESSAGE_TYPE_FLV                = 0x16
)

const (
	MESSAGE_DISPATCH_QUEUE_LENGTH = 100
)

const (
	DEFAULT_CHUNK_SIZE  = uint32(128)
	DEFAULT_WINDOW_SIZE = uint32(2500000)
)
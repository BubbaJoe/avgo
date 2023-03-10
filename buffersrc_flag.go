package avgo

//#cgo pkg-config: libavfilter
//#include <libavfilter/buffersrc.h>
import "C"

type BuffersrcFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavfilter/buffersrc.h#L36
const (
	BuffersrcFlagNoCheckFormat = BuffersrcFlag(C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT)
	BuffersrcFlagPush          = BuffersrcFlag(C.AV_BUFFERSRC_FLAG_PUSH)
	BuffersrcFlagKeepRef       = BuffersrcFlag(C.AV_BUFFERSRC_FLAG_KEEP_REF)
)

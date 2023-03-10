package avgo

//#cgo pkg-config: libavfilter
//#include <libavfilter/buffersink.h>
import "C"

type BuffersinkFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavfilter/buffersink.h#L89
const (
	BuffersinkFlagPeek      = BuffersinkFlag(C.AV_BUFFERSINK_FLAG_PEEK)
	BuffersinkFlagNoRequest = BuffersinkFlag(C.AV_BUFFERSINK_FLAG_NO_REQUEST)
)

package avgo

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavcodec/codec.h#L202
type Codec struct {
	c *C.struct_AVCodec
}

func newCodecFromC(c *C.struct_AVCodec) *Codec {
	if c == nil {
		return nil
	}
	return &Codec{c: c}
}

func (c *Codec) Name() string {
	return C.GoString(c.c.name)
}

func (c *Codec) String() string {
	return c.Name()
}

func (c *Codec) ChannelLayouts() (o []ChannelLayout) {
	if c.c.channel_layouts == nil {
		return nil
	}
	size := unsafe.Sizeof(*c.c.channel_layouts)
	for i := 0; ; i++ {
		p := *(*C.int64_t)(unsafe.Pointer(uintptr(unsafe.Pointer(c.c.channel_layouts)) + uintptr(i)*size))
		if p == 0 {
			break
		}
		o = append(o, ChannelLayout(p))
	}
	return
}

func (c *Codec) IsDecoder() bool {
	return int(C.av_codec_is_decoder(c.c)) != 0
}

func (c *Codec) IsEncoder() bool {
	return int(C.av_codec_is_encoder(c.c)) != 0
}

func (c *Codec) PixelFormats() (o []PixelFormat) {
	if c.c.pix_fmts == nil {
		return nil
	}
	size := unsafe.Sizeof(*c.c.pix_fmts)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.c.pix_fmts)) + uintptr(i)*size))
		if p == C.AV_PIX_FMT_NONE {
			break
		}
		o = append(o, PixelFormat(p))
	}
	return
}

func (c *Codec) SampleFormats() (o []SampleFormat) {
	if c.c.sample_fmts == nil {
		return nil
	}
	size := unsafe.Sizeof(*c.c.sample_fmts)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.c.sample_fmts)) + uintptr(i)*size))
		if p == C.AV_SAMPLE_FMT_NONE {
			break
		}
		o = append(o, SampleFormat(p))
	}
	return
}

func FindDecoder(id CodecID) *Codec {
	return newCodecFromC(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

func FindDecoderByName(n string) *Codec {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return newCodecFromC(C.avcodec_find_decoder_by_name(cn))
}

func FindEncoder(id CodecID) *Codec {
	return newCodecFromC(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

func FindEncoderByName(n string) *Codec {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return newCodecFromC(C.avcodec_find_encoder_by_name(cn))
}

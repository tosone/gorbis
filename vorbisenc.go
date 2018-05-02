package gorbis

/*
#cgo LDFLAGS: -lvorbisenc
#include <vorbis/vorbisenc.h>
*/
import "C"
import (
	"unsafe"
)

func encodeInit(vi *Info, channels int32, rate int32,
	max_bitrate int32, nominal_bitrate int32, min_bitrate int32) int {
	return int(C.vorbis_encode_init((*C.vorbis_info)(unsafe.Pointer(vi)),
		C.long(channels), C.long(rate), C.long(max_bitrate),
		C.long(nominal_bitrate), C.long(min_bitrate)))
}

func encodeSetupManaged(vi *Info, channels int32, rate int32,
	max_bitrate int32, nominal_bitrate int32, min_bitrate int32) int {
	return int(C.vorbis_encode_setup_managed((*C.vorbis_info)(unsafe.Pointer(vi)),
		C.long(channels), C.long(rate), C.long(max_bitrate),
		C.long(nominal_bitrate), C.long(min_bitrate)))
}

func encodeSetupVbr(vi *Info, channels int32, rate int32, quality float32) int {
	return int(C.vorbis_encode_setup_vbr((*C.vorbis_info)(unsafe.Pointer(vi)),
		C.long(channels), C.long(rate), C.float(quality)))
}

func encodeInitVbr(vi *Info, channels int32, rate int32, base_quality float32) int {
	return int(C.vorbis_encode_init_vbr((*C.vorbis_info)(unsafe.Pointer(vi)),
		C.long(channels), C.long(rate), C.float(base_quality)))
}

func encodeSetupInit(vi *Info) int {
	return int(C.vorbis_encode_setup_init((*C.vorbis_info)(unsafe.Pointer(vi))))
}

func encodeCtl(vi *Info, number int, arg *uintptr) int {
	return int(C.vorbis_encode_ctl((*C.vorbis_info)(unsafe.Pointer(vi)),
		C.int(number), unsafe.Pointer(arg)))
}

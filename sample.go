package gst

/*
#include <gst/gst.h>
*/
import "C"
import "github.com/ziutek/glib"

type Sample C.GstSample

func (s *Sample) g() *C.GstSample {
	return (*C.GstSample)(s)
}

func (s *Sample) Ref() *Sample {
	return (*Sample)(C.gst_sample_ref(s.g()))
}

func (s *Sample) Unref() {
	C.gst_sample_unref(s.g())
}

func (s *Sample) GetBuffer() *Buffer {
	return (*Buffer)(C.gst_sample_get_buffer(s.g()))
}

func SampleFromPointer(p interface{}) *Sample {
	return (*Sample)(p.(glib.Pointer))
}

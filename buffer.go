package gst

/*
#include <gst/gst.h>
*/
import "C"
import "unsafe"

type MapFlags int

const (
	MAP_READ  = MapFlags(C.GST_MAP_READ)
	MAP_WRITE = MapFlags(C.GST_MAP_WRITE)
)

type Buffer C.GstBuffer

func (b *Buffer) g() *C.GstBuffer {
	return (*C.GstBuffer)(b)
}

func (b *Buffer) Ref() *Buffer {
	return (*Buffer)(C.gst_buffer_ref(b.g()))
}

func (b *Buffer) Unref() {
	C.gst_buffer_unref(b.g())
}

func (b *Buffer) Map(flags MapFlags) *MapInfo {
	info := new(MapInfo)
	if int(C.gst_buffer_map(b.g(), info.g(), C.GstMapFlags(flags))) == 0 {
		return nil
	}
	return info
}

func (b *Buffer) Unmap(m *MapInfo) {
	C.gst_buffer_unmap(b.g(), m.g())
}

func NewBufferWrapped(data []byte) *Buffer {
	str := C.CString(string(data))
	// borrows the pointer
	return (*Buffer)(C.gst_buffer_new_wrapped(str, C.gsize(len(data))))
}

type MapInfo C.GstMapInfo

func (m *MapInfo) g() *C.GstMapInfo {
	return (*C.GstMapInfo)(m)
}

func (m *MapInfo) Data() []byte {
	var b []byte
	b = (*[1 << 30]byte)(unsafe.Pointer(m.data))[0:uint(m.size)]
	return b
}

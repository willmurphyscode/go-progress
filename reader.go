package progress

import (
	"io"

	"github.com/hashicorp/go-multierror"
)

// Reader should wrap another reader (acts as a bytes pass through)
type Reader struct {
	reader io.Reader
	monitor *Manual
}

func NewSizedReader(reader io.Reader, size int64) *Reader {
	return &Reader{
		reader: reader,
		monitor: &Manual{
			Total: size,
		},
	}
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: reader,
		monitor: &Manual{
			Total: -1,
		},
	}
}

func NewProxyReader(reader io.Reader, monitor *Manual) *Reader {
	return &Reader{
		reader: reader,
		monitor: monitor,
	}
}

func (r *Reader) SetReader(reader io.Reader) {
	r.reader = reader
}

func (r *Reader) SetCompleted() {
	r.monitor.Err = multierror.Append(r.monitor.Err, ErrCompleted)
}

func (r *Reader) Read(p []byte) (n int, err error) {
	bytes, err := r.reader.Read(p)
	r.monitor.N += int64(bytes)
	if err != nil {
		r.monitor.Err = multierror.Append(r.monitor.Err, err)
	}
	return bytes, err
}

func (r *Reader) Current() int64 {
	return r.monitor.N
}

func (r *Reader) Size() int64 {
	return r.monitor.Total
}

func (r *Reader) Error() error {
	return r.monitor.Err
}

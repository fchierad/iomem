package iomem

import (
	"errors"
)

//Mem is the io.Writer in-memory storage.
// Created to store a slice of Log.Print* output in memory for a multitude of purposes.
type Mem struct {
	data []byte
	size int
}

// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("bytes.Buffer: too large")

//New returns an initialized Mem object that
// implements the io.Writer interface
func New(n int) *Mem {
	if n == 0 {
		n = 1
	}
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return &Mem{
		data: make([]byte, 0, 2*n),
		size: n,
	}
}

//String fulfills fmt.Stringer interface
// it will return the internal bytes array as a string
func (m *Mem) String() string {
	return string(m.data)
}

//Reset empties the in memory data
func (m *Mem) Reset() {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	m.data = make([]byte, 0, 2*m.size)
}

// Write appends the contents of p to the buffer, truncating the original
// data to keep it from passing the maximum size.
// If p is greater than the maximum size, the last (size) bytes
// from p will be stored as the data. err is always nil.
func (m *Mem) Write(p []byte) (n int, err error) {
	n = len(p)
	w := p
	switch pl, dl := len(p), len(m.data); {
	case pl > m.size:
		n = m.size
		m.data = m.data[dl:]
		w = p[len(p)-n:]
	case pl == m.size:
		m.data = m.data[dl:]
	case pl < m.size && pl+dl > m.size:
		m.data = m.data[dl+pl-m.size:]
	}
	m.data = append(m.data, w...)
	return
}

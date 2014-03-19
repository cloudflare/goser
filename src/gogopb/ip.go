package gogopb

import (
	"bytes"
	"io"
	"net"
)

type IP net.IP

func (this IP) Marshal() ([]byte, error) {
	return []byte(this), nil
}

func (this IP) MarshalTo(data []byte) (int, error) {
	n := copy(data, this)
	return n, nil
}

func (this *IP) Unmarshal(data []byte) error {
	if data == nil {
		*this = nil
		return nil
	}
	if len(*this) != len(data) {
		*this = IP(make([]byte, len(data)))
	}
	copy(*this, data)
	return nil
}

func (this IP) Size() int {
	return len(this)
}

type rander interface {
	Int63() int64
}

func NewPopulatedIP(r rander) *IP {
	ip := IP(net.IPv4(1, 2, 3, 4))
	return &ip
}

func (this IP) Equal(other IP) bool {
	return bytes.Equal(this, other)
}

func (this IP) MarshalJSON() ([]byte, error) {
	return []byte("\"" + net.IP(this).String() + "\""), nil
}

func (this *IP) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return io.ErrShortBuffer
	}
	*this = IP(net.ParseIP(string(data[1 : len(data)-1])))
	return nil
}

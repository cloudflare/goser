package goser

import (
	"bytes"
	"capnp"
	"encoding/json"
	"github.com/glycerine/go-capnproto"
	"github.com/golang/protobuf/proto"
	"gogopb_both"
	"pb"
	"testing"
)

func BenchmarkPopulatePb(b *testing.B) {
	var record pb.Log
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newPbLog(&record)
	}
}

func BenchmarkPopulateGogopb(b *testing.B) {
	var record gogopb.Log
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newGogopbLog(&record)
	}
}

func BenchmarkPopulateCapnp(b *testing.B) {
	buf := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		segment := capn.NewBuffer(buf[:0])
		record := capnp.NewRootLog(segment)
		newCapnpLog(&record)
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	var record gogopb.Log
	newGogopbLog(&record)

	buf, err := json.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&record)
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkMarshalPb(b *testing.B) {
	var record pb.Log
	newPbLog(&record)

	buf, err := proto.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&record)
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkMarshalGogopb(b *testing.B) {
	var record gogopb.Log
	newGogopbLog(&record)

	buf, err := proto.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&record)
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkMarshalCapnp(b *testing.B) {
	segment := capn.NewBuffer(make([]byte, 0, 1<<20))
	record := capnp.NewRootLog(segment)
	newCapnpLog(&record)

	var buf bytes.Buffer
	_, err := segment.WriteTo(&buf)
	if err != nil {
		b.Fatalf("WriteTo: %v", err)
	}
	b.SetBytes(int64(len(buf.Bytes())))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := segment.WriteTo(&buf)
		if err != nil {
			b.Fatalf("WriteTo: %v", err)
		}
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	// generate the same marshalled buffer
	var record gogopb.Log
	newGogopbLog(&record)
	buf, err := json.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	var tmp gogopb.Log
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(buf, &tmp)
		if err != nil {
			b.Fatalf("Unmarshal: %v", err)
		}
	}
}

func BenchmarkUnmarshalPb(b *testing.B) {
	var record pb.Log
	newPbLog(&record)
	buf, err := proto.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	var tmp pb.Log
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(buf, &tmp)
		if err != nil {
			b.Fatalf("Unmarshal: %v", err)
		}
	}
}

func BenchmarkUnmarshalGogopb(b *testing.B) {
	// generate the same marshalled buffer
	var record pb.Log
	newPbLog(&record)
	buf, err := proto.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	var tmp gogopb.Log
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(buf, &tmp)
		if err != nil {
			b.Fatalf("Unmarshal: %v", err)
		}
	}
}

func BenchmarkUnmarshalCapnp(b *testing.B) {
	segment := capn.NewBuffer(make([]byte, 0, 1<<20))
	record := capnp.NewRootLog(segment)
	newCapnpLog(&record)

	var buf bytes.Buffer
	_, err := segment.WriteTo(&buf)
	if err != nil {
		b.Fatalf("WriteTo: %v", err)
	}
	b.SetBytes(int64(len(buf.Bytes())))

	segmentBuf := bytes.NewBuffer(make([]byte, 0, 1<<20))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := bytes.NewBuffer(buf.Bytes())
		seg, err := capn.ReadFromStream(r, segmentBuf)
		if err != nil {
			b.Fatalf("ReadFromStream: %v", err)
		}
		record := capnp.ReadRootLog(seg)
		_ = record
	}
}

func BenchmarkUnmarshalCapnpZeroCopy(b *testing.B) {
	segment := capn.NewBuffer(make([]byte, 0, 1<<20))
	record := capnp.NewRootLog(segment)
	newCapnpLog(&record)

	var buf bytes.Buffer
	_, err := segment.WriteTo(&buf)
	if err != nil {
		b.Fatalf("WriteTo: %v", err)
	}
	b.SetBytes(int64(len(buf.Bytes())))

	data := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seg, _, err := capn.ReadFromMemoryZeroCopy(data)
		if err != nil {
			b.Fatalf("ReadFromStream: %v", err)
		}
		record := capnp.ReadRootLog(seg)
		_ = record
	}
}

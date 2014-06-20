package capnp

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
	"unsafe"
)

type HTTP C.Struct

func NewHTTP(s *C.Segment) HTTP           { return HTTP(s.NewStruct(16, 4)) }
func NewRootHTTP(s *C.Segment) HTTP       { return HTTP(s.NewRootStruct(16, 4)) }
func ReadRootHTTP(s *C.Segment) HTTP      { return HTTP(s.Root(0).ToStruct()) }
func (s HTTP) Protocol() HTTPProtocol     { return HTTPProtocol(C.Struct(s).Get16(0)) }
func (s HTTP) SetProtocol(v HTTPProtocol) { C.Struct(s).Set16(0, uint16(v)) }
func (s HTTP) Status() uint16             { return C.Struct(s).Get16(2) }
func (s HTTP) SetStatus(v uint16)         { C.Struct(s).Set16(2, v) }
func (s HTTP) HostStatus() uint16         { return C.Struct(s).Get16(4) }
func (s HTTP) SetHostStatus(v uint16)     { C.Struct(s).Set16(4, v) }
func (s HTTP) UpStatus() uint16           { return C.Struct(s).Get16(6) }
func (s HTTP) SetUpStatus(v uint16)       { C.Struct(s).Set16(6, v) }
func (s HTTP) Method() HTTPMethod         { return HTTPMethod(C.Struct(s).Get16(8)) }
func (s HTTP) SetMethod(v HTTPMethod)     { C.Struct(s).Set16(8, uint16(v)) }
func (s HTTP) ContentType() string        { return C.Struct(s).GetObject(0).ToText() }
func (s HTTP) SetContentType(v string)    { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s HTTP) UserAgent() string          { return C.Struct(s).GetObject(1).ToText() }
func (s HTTP) SetUserAgent(v string)      { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s HTTP) Referer() string            { return C.Struct(s).GetObject(2).ToText() }
func (s HTTP) SetReferer(v string)        { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }
func (s HTTP) RequestURI() string         { return C.Struct(s).GetObject(3).ToText() }
func (s HTTP) SetRequestURI(v string)     { C.Struct(s).SetObject(3, s.Segment.NewText(v)) }
func (s HTTP) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"protocol\":")
	if err != nil {
		return err
	}
	{
		s := s.Protocol()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"status\":")
	if err != nil {
		return err
	}
	{
		s := s.Status()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"hostStatus\":")
	if err != nil {
		return err
	}
	{
		s := s.HostStatus()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"upStatus\":")
	if err != nil {
		return err
	}
	{
		s := s.UpStatus()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"method\":")
	if err != nil {
		return err
	}
	{
		s := s.Method()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"contentType\":")
	if err != nil {
		return err
	}
	{
		s := s.ContentType()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"userAgent\":")
	if err != nil {
		return err
	}
	{
		s := s.UserAgent()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"referer\":")
	if err != nil {
		return err
	}
	{
		s := s.Referer()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"requestURI\":")
	if err != nil {
		return err
	}
	{
		s := s.RequestURI()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HTTP) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type HTTP_List C.PointerList

func NewHTTPList(s *C.Segment, sz int) HTTP_List { return HTTP_List(s.NewCompositeList(16, 4, sz)) }
func (s HTTP_List) Len() int                     { return C.PointerList(s).Len() }
func (s HTTP_List) At(i int) HTTP                { return HTTP(C.PointerList(s).At(i).ToStruct()) }
func (s HTTP_List) ToArray() []HTTP              { return *(*[]HTTP)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type HTTPProtocol uint16

const (
	HTTPPROTOCOL_UNKNOWN HTTPProtocol = 0
	HTTPPROTOCOL_HTTP10               = 1
	HTTPPROTOCOL_HTTP11               = 2
	HTTPPROTOCOL_MAX                  = 3
)

func (c HTTPProtocol) String() string {
	switch c {
	case HTTPPROTOCOL_UNKNOWN:
		return "unknown"
	case HTTPPROTOCOL_HTTP10:
		return "http10"
	case HTTPPROTOCOL_HTTP11:
		return "http11"
	case HTTPPROTOCOL_MAX:
		return "max"
	default:
		return ""
	}
}

type HTTPProtocol_List C.PointerList

func NewHTTPProtocolList(s *C.Segment, sz int) HTTPProtocol_List {
	return HTTPProtocol_List(s.NewUInt16List(sz))
}
func (s HTTPProtocol_List) Len() int              { return C.UInt16List(s).Len() }
func (s HTTPProtocol_List) At(i int) HTTPProtocol { return HTTPProtocol(C.UInt16List(s).At(i)) }
func (s HTTPProtocol_List) ToArray() []HTTPProtocol {
	return *(*[]HTTPProtocol)(unsafe.Pointer(C.UInt16List(s).ToEnumArray()))
}
func (s HTTPProtocol) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HTTPProtocol) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type HTTPMethod uint16

const (
	HTTPMETHOD_UNKNOWN  HTTPMethod = 0
	HTTPMETHOD_GET                 = 1
	HTTPMETHOD_POST                = 2
	HTTPMETHOD_DELETE              = 3
	HTTPMETHOD_PUT                 = 4
	HTTPMETHOD_HEAD                = 5
	HTTPMETHOD_PURGE               = 6
	HTTPMETHOD_OPTIONS             = 7
	HTTPMETHOD_PROPFIND            = 8
	HTTPMETHOD_MKCOL               = 9
	HTTPMETHOD_PATCH               = 10
	HTTPMETHOD_MAX                 = 11
)

func (c HTTPMethod) String() string {
	switch c {
	case HTTPMETHOD_UNKNOWN:
		return "unknown"
	case HTTPMETHOD_GET:
		return "get"
	case HTTPMETHOD_POST:
		return "post"
	case HTTPMETHOD_DELETE:
		return "delete"
	case HTTPMETHOD_PUT:
		return "put"
	case HTTPMETHOD_HEAD:
		return "head"
	case HTTPMETHOD_PURGE:
		return "purge"
	case HTTPMETHOD_OPTIONS:
		return "options"
	case HTTPMETHOD_PROPFIND:
		return "propfind"
	case HTTPMETHOD_MKCOL:
		return "mkcol"
	case HTTPMETHOD_PATCH:
		return "patch"
	case HTTPMETHOD_MAX:
		return "max"
	default:
		return ""
	}
}

type HTTPMethod_List C.PointerList

func NewHTTPMethodList(s *C.Segment, sz int) HTTPMethod_List {
	return HTTPMethod_List(s.NewUInt16List(sz))
}
func (s HTTPMethod_List) Len() int            { return C.UInt16List(s).Len() }
func (s HTTPMethod_List) At(i int) HTTPMethod { return HTTPMethod(C.UInt16List(s).At(i)) }
func (s HTTPMethod_List) ToArray() []HTTPMethod {
	return *(*[]HTTPMethod)(unsafe.Pointer(C.UInt16List(s).ToEnumArray()))
}
func (s HTTPMethod) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s HTTPMethod) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type CacheStatus uint16

const (
	CACHESTATUS_UNKNOWN CacheStatus = 0
	CACHESTATUS_MISS                = 1
	CACHESTATUS_EXPIRED             = 2
	CACHESTATUS_HIT                 = 3
	CACHESTATUS_MAX                 = 4
)

func (c CacheStatus) String() string {
	switch c {
	case CACHESTATUS_UNKNOWN:
		return "unknown"
	case CACHESTATUS_MISS:
		return "miss"
	case CACHESTATUS_EXPIRED:
		return "expired"
	case CACHESTATUS_HIT:
		return "hit"
	case CACHESTATUS_MAX:
		return "max"
	default:
		return ""
	}
}

type CacheStatus_List C.PointerList

func NewCacheStatusList(s *C.Segment, sz int) CacheStatus_List {
	return CacheStatus_List(s.NewUInt16List(sz))
}
func (s CacheStatus_List) Len() int             { return C.UInt16List(s).Len() }
func (s CacheStatus_List) At(i int) CacheStatus { return CacheStatus(C.UInt16List(s).At(i)) }
func (s CacheStatus_List) ToArray() []CacheStatus {
	return *(*[]CacheStatus)(unsafe.Pointer(C.UInt16List(s).ToEnumArray()))
}
func (s CacheStatus) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s CacheStatus) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type Origin C.Struct

func NewOrigin(s *C.Segment) Origin           { return Origin(s.NewStruct(8, 2)) }
func NewRootOrigin(s *C.Segment) Origin       { return Origin(s.NewRootStruct(8, 2)) }
func ReadRootOrigin(s *C.Segment) Origin      { return Origin(s.Root(0).ToStruct()) }
func (s Origin) Ip() []byte                   { return C.Struct(s).GetObject(0).ToData() }
func (s Origin) SetIp(v []byte)               { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s Origin) Port() uint16                 { return C.Struct(s).Get16(0) }
func (s Origin) SetPort(v uint16)             { C.Struct(s).Set16(0, v) }
func (s Origin) Hostname() string             { return C.Struct(s).GetObject(1).ToText() }
func (s Origin) SetHostname(v string)         { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s Origin) Protocol() OriginProtocol     { return OriginProtocol(C.Struct(s).Get16(2)) }
func (s Origin) SetProtocol(v OriginProtocol) { C.Struct(s).Set16(2, uint16(v)) }
func (s Origin) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ip\":")
	if err != nil {
		return err
	}
	{
		s := s.Ip()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"port\":")
	if err != nil {
		return err
	}
	{
		s := s.Port()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"hostname\":")
	if err != nil {
		return err
	}
	{
		s := s.Hostname()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"protocol\":")
	if err != nil {
		return err
	}
	{
		s := s.Protocol()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Origin) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type Origin_List C.PointerList

func NewOriginList(s *C.Segment, sz int) Origin_List { return Origin_List(s.NewCompositeList(8, 2, sz)) }
func (s Origin_List) Len() int                       { return C.PointerList(s).Len() }
func (s Origin_List) At(i int) Origin                { return Origin(C.PointerList(s).At(i).ToStruct()) }
func (s Origin_List) ToArray() []Origin {
	return *(*[]Origin)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type OriginProtocol uint16

const (
	ORIGINPROTOCOL_UNKNOWN OriginProtocol = 0
	ORIGINPROTOCOL_HTTP                   = 1
	ORIGINPROTOCOL_HTTPS                  = 2
	ORIGINPROTOCOL_MAX                    = 3
)

func (c OriginProtocol) String() string {
	switch c {
	case ORIGINPROTOCOL_UNKNOWN:
		return "unknown"
	case ORIGINPROTOCOL_HTTP:
		return "http"
	case ORIGINPROTOCOL_HTTPS:
		return "https"
	case ORIGINPROTOCOL_MAX:
		return "max"
	default:
		return ""
	}
}

type OriginProtocol_List C.PointerList

func NewOriginProtocolList(s *C.Segment, sz int) OriginProtocol_List {
	return OriginProtocol_List(s.NewUInt16List(sz))
}
func (s OriginProtocol_List) Len() int                { return C.UInt16List(s).Len() }
func (s OriginProtocol_List) At(i int) OriginProtocol { return OriginProtocol(C.UInt16List(s).At(i)) }
func (s OriginProtocol_List) ToArray() []OriginProtocol {
	return *(*[]OriginProtocol)(unsafe.Pointer(C.UInt16List(s).ToEnumArray()))
}
func (s OriginProtocol) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s OriginProtocol) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type ZonePlan uint16

const (
	ZONEPLAN_UNKNOWN ZonePlan = 0
	ZONEPLAN_FREE             = 1
	ZONEPLAN_PRO              = 2
	ZONEPLAN_BIZ              = 3
	ZONEPLAN_ENT              = 4
	ZONEPLAN_MAX              = 5
)

func (c ZonePlan) String() string {
	switch c {
	case ZONEPLAN_FREE:
		return "Free"
	case ZONEPLAN_PRO:
		return "Pro"
	case ZONEPLAN_BIZ:
		return "Business"
	case ZONEPLAN_ENT:
		return "Enterprise"
	case ZONEPLAN_MAX:
		return "max"
	default:
		return ""
	}
}

type ZonePlan_List C.PointerList

func NewZonePlanList(s *C.Segment, sz int) ZonePlan_List { return ZonePlan_List(s.NewUInt16List(sz)) }
func (s ZonePlan_List) Len() int                         { return C.UInt16List(s).Len() }
func (s ZonePlan_List) At(i int) ZonePlan                { return ZonePlan(C.UInt16List(s).At(i)) }
func (s ZonePlan_List) ToArray() []ZonePlan {
	return *(*[]ZonePlan)(unsafe.Pointer(C.UInt16List(s).ToEnumArray()))
}
func (s ZonePlan) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ZonePlan) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type Log C.Struct

func NewLog(s *C.Segment) Log              { return Log(s.NewStruct(32, 6)) }
func NewRootLog(s *C.Segment) Log          { return Log(s.NewRootStruct(32, 6)) }
func ReadRootLog(s *C.Segment) Log         { return Log(s.Root(0).ToStruct()) }
func (s Log) Timestamp() int64             { return int64(C.Struct(s).Get64(0)) }
func (s Log) SetTimestamp(v int64)         { C.Struct(s).Set64(0, uint64(v)) }
func (s Log) ZoneId() uint32               { return C.Struct(s).Get32(8) }
func (s Log) SetZoneId(v uint32)           { C.Struct(s).Set32(8, v) }
func (s Log) ZonePlan() ZonePlan           { return ZonePlan(C.Struct(s).Get16(12)) }
func (s Log) SetZonePlan(v ZonePlan)       { C.Struct(s).Set16(12, uint16(v)) }
func (s Log) Http() HTTP                   { return HTTP(C.Struct(s).GetObject(0).ToStruct()) }
func (s Log) SetHttp(v HTTP)               { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Log) Origin() Origin               { return Origin(C.Struct(s).GetObject(1).ToStruct()) }
func (s Log) SetOrigin(v Origin)           { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Log) Country() Country             { return Country(C.Struct(s).Get16(14)) }
func (s Log) SetCountry(v Country)         { C.Struct(s).Set16(14, uint16(v)) }
func (s Log) CacheStatus() CacheStatus     { return CacheStatus(C.Struct(s).Get16(16)) }
func (s Log) SetCacheStatus(v CacheStatus) { C.Struct(s).Set16(16, uint16(v)) }
func (s Log) ServerIp() []byte             { return C.Struct(s).GetObject(2).ToData() }
func (s Log) SetServerIp(v []byte)         { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }
func (s Log) ServerName() string           { return C.Struct(s).GetObject(3).ToText() }
func (s Log) SetServerName(v string)       { C.Struct(s).SetObject(3, s.Segment.NewText(v)) }
func (s Log) RemoteIp() []byte             { return C.Struct(s).GetObject(4).ToData() }
func (s Log) SetRemoteIp(v []byte)         { C.Struct(s).SetObject(4, s.Segment.NewData(v)) }
func (s Log) BytesDlv() uint64             { return C.Struct(s).Get64(24) }
func (s Log) SetBytesDlv(v uint64)         { C.Struct(s).Set64(24, v) }
func (s Log) RayId() string                { return C.Struct(s).GetObject(5).ToText() }
func (s Log) SetRayId(v string)            { C.Struct(s).SetObject(5, s.Segment.NewText(v)) }
func (s Log) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"timestamp\":")
	if err != nil {
		return err
	}
	{
		s := s.Timestamp()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"zoneId\":")
	if err != nil {
		return err
	}
	{
		s := s.ZoneId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"zonePlan\":")
	if err != nil {
		return err
	}
	{
		s := s.ZonePlan()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"http\":")
	if err != nil {
		return err
	}
	{
		s := s.Http()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"origin\":")
	if err != nil {
		return err
	}
	{
		s := s.Origin()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"country\":")
	if err != nil {
		return err
	}
	{
		s := s.Country()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"cacheStatus\":")
	if err != nil {
		return err
	}
	{
		s := s.CacheStatus()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"serverIp\":")
	if err != nil {
		return err
	}
	{
		s := s.ServerIp()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"serverName\":")
	if err != nil {
		return err
	}
	{
		s := s.ServerName()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"remoteIp\":")
	if err != nil {
		return err
	}
	{
		s := s.RemoteIp()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"bytesDlv\":")
	if err != nil {
		return err
	}
	{
		s := s.BytesDlv()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"rayId\":")
	if err != nil {
		return err
	}
	{
		s := s.RayId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Log) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}

type Log_List C.PointerList

func NewLogList(s *C.Segment, sz int) Log_List { return Log_List(s.NewCompositeList(32, 6, sz)) }
func (s Log_List) Len() int                    { return C.PointerList(s).Len() }
func (s Log_List) At(i int) Log                { return Log(C.PointerList(s).At(i).ToStruct()) }
func (s Log_List) ToArray() []Log              { return *(*[]Log)(unsafe.Pointer(C.PointerList(s).ToArray())) }

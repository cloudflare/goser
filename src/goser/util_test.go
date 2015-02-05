package goser

import (
	"capnp"
	"github.com/golang/protobuf/proto"
	"gogopb_both"
	"net"
	"pb"
	"time"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.146 Safari/537.36"

func newPbLog(record *pb.Log) {
	record.Timestamp = proto.Int64(time.Now().UnixNano())
	record.ZoneId = proto.Uint32(123456)
	record.ZonePlan = pb.ZonePlan_FREE.Enum()

	record.Http = &pb.HTTP{
		Protocol:    pb.HTTP_HTTP11.Enum(),
		Status:      proto.Uint32(200),
		HostStatus:  proto.Uint32(503),
		UpStatus:    proto.Uint32(520),
		Method:      pb.HTTP_GET.Enum(),
		ContentType: proto.String("text/html"),
		UserAgent:   proto.String(userAgent),
		Referer:     proto.String("https://www.cloudflare.com/"),
		RequestURI:  proto.String("/cdn-cgi/trace"),
	}

	record.Origin = &pb.Origin{
		Ip:       []byte(net.IPv4(1, 2, 3, 4).To4()),
		Port:     proto.Uint32(8080),
		Hostname: proto.String("www.example.com"),
		Protocol: pb.Origin_HTTPS.Enum(),
	}

	record.Country = pb.Country_US.Enum()
	record.CacheStatus = pb.CacheStatus_HIT.Enum()
	record.ServerIp = []byte(net.IPv4(192, 168, 1, 1).To4())
	record.ServerName = proto.String("metal.cloudflare.com")
	record.RemoteIp = []byte(net.IPv4(10, 1, 2, 3).To4())
	record.BytesDlv = proto.Uint64(123456)
	record.RayId = proto.String("10c73629cce30078-LAX")
}

func newGogopbLog(record *gogopb.Log) {
	record.Timestamp = time.Now().UnixNano()
	record.ZoneId = 123456
	record.ZonePlan = gogopb.ZonePlan_FREE

	record.Http = gogopb.HTTP{
		Protocol:    gogopb.HTTP_HTTP11,
		Status:      200,
		HostStatus:  503,
		UpStatus:    520,
		Method:      gogopb.HTTP_GET,
		ContentType: "text/html",
		UserAgent:   userAgent,
		Referer:     "https://www.cloudflare.com/",
		RequestURI:  "/cdn-cgi/trace",
	}

	record.Origin = gogopb.Origin{
		Ip:       gogopb.IP(net.IPv4(1, 2, 3, 4).To4()),
		Port:     8080,
		Hostname: "www.example.com",
		Protocol: gogopb.Origin_HTTPS,
	}

	record.Country = gogopb.Country_US
	record.CacheStatus = gogopb.CacheStatus_HIT
	record.ServerIp = gogopb.IP(net.IPv4(192, 168, 1, 1).To4())
	record.ServerName = "metal.cloudflare.com"
	record.RemoteIp = gogopb.IP(net.IPv4(10, 1, 2, 3).To4())
	record.BytesDlv = 123456
	record.RayId = "10c73629cce30078-LAX"
}

func newCapnpLog(record *capnp.Log) {
	record.SetTimestamp(time.Now().UnixNano())
	record.SetZoneId(123456)
	record.SetZonePlan(capnp.ZONEPLAN_FREE)

	http := capnp.NewHTTP(record.Segment)
	record.SetHttp(http)
	http.SetProtocol(capnp.HTTPPROTOCOL_HTTP11)
	http.SetStatus(200)
	http.SetHostStatus(503)
	http.SetUpStatus(520)
	http.SetMethod(capnp.HTTPMETHOD_GET)
	http.SetContentType("text/html")
	http.SetUserAgent(userAgent)
	http.SetReferer("http://www.w3.org/hypertext/DataSources/Overview.html")
	http.SetRequestURI("/a9pPJR1.jpg")

	origin := capnp.NewOrigin(record.Segment)
	record.SetOrigin(origin)
	origin.SetIp([]byte(net.IPv4(1, 2, 3, 4).To4()))
	origin.SetPort(8080)
	origin.SetHostname("www.cloudflare.com")
	origin.SetProtocol(capnp.ORIGINPROTOCOL_HTTPS)

	record.SetCacheStatus(capnp.CACHESTATUS_HIT)
	record.SetServerIp([]byte(net.IPv4(192, 168, 1, 1).To4()))
	record.SetServerName("metal.cloudflare.com")
	record.SetRemoteIp([]byte(net.IPv4(10, 1, 2, 3).To4()))
	record.SetBytesDlv(123456)
	record.SetRayId("10c73629cce30078-LAX")
}

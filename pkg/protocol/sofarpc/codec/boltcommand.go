package codec

import (
	"gitlab.alipay-inc.com/afe/mosn/pkg/protocol/sofarpc"
	"net"
)

//command defination
type boltCommand struct {
	protocol byte  // 1 boltv1, 2 boltv2, 13 tr request, 14 tr response
	cmdCode  int16 // 0 心跳,1 request, 2 response
	version  byte  // ver2
	cmdType  byte  // 0 response,1 是request,2是 req oneway
	codec    byte
	//protoSwitchStatus byte

	id            uint32
	classLength   int16
	headerLength  int16
	contentLength int

	class   []byte
	header  []byte
	content []byte

	invokeContext interface{}
}

// bolt request command
type boltRequestCommand struct {
	//rpc command
	boltCommand

	//request command
	timeout int

	//rpc request command
	requestObject interface{}
	requestClass  string

	//customSerializer CustomSerializer
	requestHeader map[string]string
	arriveTime    int64

	//add ver1 and switchcode for boltv2 request
	ver1       byte
	switchCode byte
}

// bolt response command
type boltResponseCommand struct {
	//rpc command
	boltCommand

	//response command
	responseStatus     int16
	responseTimeMillis int64
	responseHost       net.Addr
	cause              error

	//rpc response command
	responseObject interface{}
	responseClass  string

	//customSerializer CustomSerializer
	responseHeader map[string]string

	errorMsg string

	//add ver1 and switchcode for boltv2 response
	ver1       byte
	switchCode byte
}

func (b *boltCommand) GetProtocolCode() byte {
	//return sofarpc.PROTOCOL_CODE_V1
	return b.protocol
}

func (b *boltCommand) GetCmdCode() int16 {
	return b.cmdCode
}

func (b *boltCommand) GetId() uint32 {
	return b.id
}

func (b *boltCommand) GetClass() []byte {
	return b.class
}

func (b *boltCommand) GetHeader() []byte {
	return b.header
}

func (b *boltCommand) GetContent() []byte {
	return b.content
}

// FOR BOLRV2 Request
func (b *boltRequestCommand) SetVer1(ver1 byte) {
	b.ver1 = ver1
}
func (b *boltRequestCommand) GetVer1() byte {
	return b.ver1
}
func (b *boltRequestCommand) SetSwitch(switchCode byte) {
	b.switchCode = switchCode
}
func (b *boltRequestCommand) GetSwitch() byte {
	return b.switchCode
}

func (b *boltRequestCommand) GetProtocolCode() byte {
	return sofarpc.PROTOCOL_CODE_V1
}

func (b *boltRequestCommand) GetCmdCode() int16 {
	return b.cmdCode
}

func (b *boltRequestCommand) GetId() uint32 {
	return b.id
}

func (b *boltRequestCommand) SetTimeout(timeout int) {
	b.timeout = timeout
}

func (b *boltRequestCommand) SetArriveTime(arriveTime int64) {
	b.arriveTime = arriveTime
}
func (b *boltRequestCommand) GetArriveTime() int64 {
	return b.arriveTime
}

func (b *boltRequestCommand) SetRequestHeader(headerMap map[string]string) {
	b.requestHeader = headerMap
}

func (b *boltRequestCommand) GetRequestHeader() map[string]string {
	return b.requestHeader
}

func (b *boltResponseCommand) GetProtocolCode() byte {
	return sofarpc.PROTOCOL_CODE_V1
}

func (b *boltResponseCommand) GetCmdCode() int16 {
	return b.cmdCode
}

func (b *boltResponseCommand) GetId() uint32 {
	return b.id
}

func (b *boltResponseCommand) SetResponseStatus(status int16) {
	b.responseStatus = status
}

func (b *boltResponseCommand) SetResponseTimeMillis(responseTime int64) {
	b.responseTimeMillis = responseTime
}

func (b *boltResponseCommand) SetResponseHost(host net.Addr) {
	b.responseHost = host
}

//FOR BOLTV2 RESPONSE
func (b *boltResponseCommand) SetVer1(ver1 byte) {
	b.ver1 = ver1
}
func (b *boltResponseCommand) GetVer1() byte {
	return b.ver1
}

func (b *boltResponseCommand) SetSwitch(switchCode byte) {
	b.switchCode = switchCode
}
func (b *boltResponseCommand) GetSwitch() byte {
	return b.switchCode
}

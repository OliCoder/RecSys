package engine

import (
	"github.com/OliCoder/RecSys/settings"
	"github.com/apache/thrift/lib/go/thrift"
	log "github.com/sirupsen/logrus"
	"net"
)
var (
	tFactory *thrift.TBufferedTransportFactory
	pFactory *thrift.TBinaryProtocolFactory
	transport thrift.TTransport

	ClientTransport *thrift.TSocket
)

func init() {
	host := settings.GetConfigInstance().RpcHost
	port := settings.GetConfigInstance().RpcPort

	tFactory = thrift.NewTBufferedTransportFactory(8192)
	pFactory = thrift.NewTBinaryProtocolFactory(false, true)
	ClientTransport, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Errorf("Error opening socket to %s:%s", host, port)
	}

	transport, _ = tFactory.GetTransport(ClientTransport)
	if err = ClientTransport.Open(); err != nil {
		log.Errorf("Error opening socket to %s:%s", host, port)
	}
}

// TODO very ugly!!! just for demo!! need to write a client pool
func NewClient() *EngineServiceClient {
	return NewEngineServiceClientFactory(transport, pFactory)
}
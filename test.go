package main

func test() {
	//tFactory := thrift.NewTBufferedTransportFactory(8192)
	//pFactory := thrift.NewTBinaryProtocolFactory(false, true)
	//
	//host := GetConfigInstance().GetRpcHost()
	//port := GetConfigInstance().GetRpcPort()
	//transport, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	//if err != nil {
	//	log.Panicf("Fail to resolve address %s:%s, err:%v", host, port, err)
	//}
	//
	//useTransport, err := tFactory.GetTransport(transport)
	//if err != nil {
	//	log.Panic("Fail to get transport. err:%v", err)
	//}
	//client := engine.NewEngineServiceClientFactory(useTransport, pFactory)
	//if err := transport.Open(); err != nil {
	//	log.Errorf("Fail to open socket on %s:%s, err:%v", host, port, err)
	//}
	//defer transport.Close()

	//req := "hello"
	//res, err := client.Echo(req)
	//if err != nil {
	//	log.Errorf("call funtion failed, err:%v", err)
	//}
	//log.Info(res)
}
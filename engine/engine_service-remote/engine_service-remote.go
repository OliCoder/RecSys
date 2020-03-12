// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"engine"
)

var _ = engine.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  bool UpdateEngineGroup(string groupConf)")
  fmt.Fprintln(os.Stderr, "  i64 Predict(UserProfile userProfile, i64 movieId)")
  fmt.Fprintln(os.Stderr, "   Recommend(UserProfile userProfile, i32 topk)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := engine.NewEngineServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "UpdateEngineGroup":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UpdateEngineGroup requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.UpdateEngineGroup(context.Background(), value0))
    fmt.Print("\n")
    break
  case "Predict":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "Predict requires 2 args")
      flag.Usage()
    }
    arg10 := flag.Arg(1)
    mbTrans11 := thrift.NewTMemoryBufferLen(len(arg10))
    defer mbTrans11.Close()
    _, err12 := mbTrans11.WriteString(arg10)
    if err12 != nil {
      Usage()
      return
    }
    factory13 := thrift.NewTJSONProtocolFactory()
    jsProt14 := factory13.GetProtocol(mbTrans11)
    argvalue0 := engine.NewUserProfile()
    err15 := argvalue0.Read(jsProt14)
    if err15 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err16 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err16 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.Predict(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "Recommend":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "Recommend requires 2 args")
      flag.Usage()
    }
    arg17 := flag.Arg(1)
    mbTrans18 := thrift.NewTMemoryBufferLen(len(arg17))
    defer mbTrans18.Close()
    _, err19 := mbTrans18.WriteString(arg17)
    if err19 != nil {
      Usage()
      return
    }
    factory20 := thrift.NewTJSONProtocolFactory()
    jsProt21 := factory20.GetProtocol(mbTrans18)
    argvalue0 := engine.NewUserProfile()
    err22 := argvalue0.Read(jsProt21)
    if err22 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err23 := (strconv.Atoi(flag.Arg(2)))
    if err23 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.Recommend(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}

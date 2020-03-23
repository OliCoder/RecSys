// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package engine

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - UserId
//  - MovieWatchedNumRecently
type UserProfile struct {
  UserId int64 `thrift:"userId,1" db:"userId" json:"userId"`
  MovieWatchedNumRecently int32 `thrift:"movieWatchedNumRecently,2" db:"movieWatchedNumRecently" json:"movieWatchedNumRecently"`
}

func NewUserProfile() *UserProfile {
  return &UserProfile{}
}


func (p *UserProfile) GetUserId() int64 {
  return p.UserId
}

func (p *UserProfile) GetMovieWatchedNumRecently() int32 {
  return p.MovieWatchedNumRecently
}
func (p *UserProfile) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *UserProfile)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.UserId = v
}
  return nil
}

func (p *UserProfile)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MovieWatchedNumRecently = v
}
  return nil
}

func (p *UserProfile) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("UserProfile"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *UserProfile) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("userId", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userId: ", p), err) }
  if err := oprot.WriteI64(int64(p.UserId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.userId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userId: ", p), err) }
  return err
}

func (p *UserProfile) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("movieWatchedNumRecently", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:movieWatchedNumRecently: ", p), err) }
  if err := oprot.WriteI32(int32(p.MovieWatchedNumRecently)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.movieWatchedNumRecently (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:movieWatchedNumRecently: ", p), err) }
  return err
}

func (p *UserProfile) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("UserProfile(%+v)", *p)
}

type EngineService interface {
  // Parameters:
  //  - GroupConf
  UpdateEngineGroup(ctx context.Context, groupConf string) (r bool, err error)
  // Parameters:
  //  - UserProfile
  //  - MovieId
  Predict(ctx context.Context, userProfile *UserProfile, movieId int64) (r int64, err error)
  // Parameters:
  //  - UserProfile
  //  - Topk
  Recommend(ctx context.Context, userProfile *UserProfile, topk int32) (r []int64, err error)
}

type EngineServiceClient struct {
  c thrift.TClient
}

func NewEngineServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *EngineServiceClient {
  return &EngineServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewEngineServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *EngineServiceClient {
  return &EngineServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewEngineServiceClient(c thrift.TClient) *EngineServiceClient {
  return &EngineServiceClient{
    c: c,
  }
}

func (p *EngineServiceClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - GroupConf
func (p *EngineServiceClient) UpdateEngineGroup(ctx context.Context, groupConf string) (r bool, err error) {
  var _args0 EngineServiceUpdateEngineGroupArgs
  _args0.GroupConf = groupConf
  var _result1 EngineServiceUpdateEngineGroupResult
  if err = p.Client_().Call(ctx, "UpdateEngineGroup", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - UserProfile
//  - MovieId
func (p *EngineServiceClient) Predict(ctx context.Context, userProfile *UserProfile, movieId int64) (r int64, err error) {
  var _args2 EngineServicePredictArgs
  _args2.UserProfile = userProfile
  _args2.MovieId = movieId
  var _result3 EngineServicePredictResult
  if err = p.Client_().Call(ctx, "Predict", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

// Parameters:
//  - UserProfile
//  - Topk
func (p *EngineServiceClient) Recommend(ctx context.Context, userProfile *UserProfile, topk int32) (r []int64, err error) {
  var _args4 EngineServiceRecommendArgs
  _args4.UserProfile = userProfile
  _args4.Topk = topk
  var _result5 EngineServiceRecommendResult
  if err = p.Client_().Call(ctx, "Recommend", &_args4, &_result5); err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

type EngineServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler EngineService
}

func (p *EngineServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *EngineServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *EngineServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewEngineServiceProcessor(handler EngineService) *EngineServiceProcessor {

  self6 := &EngineServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["UpdateEngineGroup"] = &engineServiceProcessorUpdateEngineGroup{handler:handler}
  self6.processorMap["Predict"] = &engineServiceProcessorPredict{handler:handler}
  self6.processorMap["Recommend"] = &engineServiceProcessorRecommend{handler:handler}
return self6
}

func (p *EngineServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x7

}

type engineServiceProcessorUpdateEngineGroup struct {
  handler EngineService
}

func (p *engineServiceProcessorUpdateEngineGroup) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := EngineServiceUpdateEngineGroupArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("UpdateEngineGroup", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := EngineServiceUpdateEngineGroupResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.UpdateEngineGroup(ctx, args.GroupConf); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing UpdateEngineGroup: " + err2.Error())
    oprot.WriteMessageBegin("UpdateEngineGroup", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("UpdateEngineGroup", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type engineServiceProcessorPredict struct {
  handler EngineService
}

func (p *engineServiceProcessorPredict) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := EngineServicePredictArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Predict", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := EngineServicePredictResult{}
var retval int64
  var err2 error
  if retval, err2 = p.handler.Predict(ctx, args.UserProfile, args.MovieId); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Predict: " + err2.Error())
    oprot.WriteMessageBegin("Predict", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Predict", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type engineServiceProcessorRecommend struct {
  handler EngineService
}

func (p *engineServiceProcessorRecommend) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := EngineServiceRecommendArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Recommend", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := EngineServiceRecommendResult{}
var retval []int64
  var err2 error
  if retval, err2 = p.handler.Recommend(ctx, args.UserProfile, args.Topk); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Recommend: " + err2.Error())
    oprot.WriteMessageBegin("Recommend", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("Recommend", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - GroupConf
type EngineServiceUpdateEngineGroupArgs struct {
  GroupConf string `thrift:"groupConf,1" db:"groupConf" json:"groupConf"`
}

func NewEngineServiceUpdateEngineGroupArgs() *EngineServiceUpdateEngineGroupArgs {
  return &EngineServiceUpdateEngineGroupArgs{}
}


func (p *EngineServiceUpdateEngineGroupArgs) GetGroupConf() string {
  return p.GroupConf
}
func (p *EngineServiceUpdateEngineGroupArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServiceUpdateEngineGroupArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.GroupConf = v
}
  return nil
}

func (p *EngineServiceUpdateEngineGroupArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("UpdateEngineGroup_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServiceUpdateEngineGroupArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("groupConf", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:groupConf: ", p), err) }
  if err := oprot.WriteString(string(p.GroupConf)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.groupConf (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:groupConf: ", p), err) }
  return err
}

func (p *EngineServiceUpdateEngineGroupArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServiceUpdateEngineGroupArgs(%+v)", *p)
}

// Attributes:
//  - Success
type EngineServiceUpdateEngineGroupResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEngineServiceUpdateEngineGroupResult() *EngineServiceUpdateEngineGroupResult {
  return &EngineServiceUpdateEngineGroupResult{}
}

var EngineServiceUpdateEngineGroupResult_Success_DEFAULT bool
func (p *EngineServiceUpdateEngineGroupResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return EngineServiceUpdateEngineGroupResult_Success_DEFAULT
  }
return *p.Success
}
func (p *EngineServiceUpdateEngineGroupResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *EngineServiceUpdateEngineGroupResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServiceUpdateEngineGroupResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *EngineServiceUpdateEngineGroupResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("UpdateEngineGroup_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServiceUpdateEngineGroupResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *EngineServiceUpdateEngineGroupResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServiceUpdateEngineGroupResult(%+v)", *p)
}

// Attributes:
//  - UserProfile
//  - MovieId
type EngineServicePredictArgs struct {
  UserProfile *UserProfile `thrift:"userProfile,1" db:"userProfile" json:"userProfile"`
  MovieId int64 `thrift:"movieId,2" db:"movieId" json:"movieId"`
}

func NewEngineServicePredictArgs() *EngineServicePredictArgs {
  return &EngineServicePredictArgs{}
}

var EngineServicePredictArgs_UserProfile_DEFAULT *UserProfile
func (p *EngineServicePredictArgs) GetUserProfile() *UserProfile {
  if !p.IsSetUserProfile() {
    return EngineServicePredictArgs_UserProfile_DEFAULT
  }
return p.UserProfile
}

func (p *EngineServicePredictArgs) GetMovieId() int64 {
  return p.MovieId
}
func (p *EngineServicePredictArgs) IsSetUserProfile() bool {
  return p.UserProfile != nil
}

func (p *EngineServicePredictArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServicePredictArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.UserProfile = &UserProfile{}
  if err := p.UserProfile.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
  }
  return nil
}

func (p *EngineServicePredictArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MovieId = v
}
  return nil
}

func (p *EngineServicePredictArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Predict_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServicePredictArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("userProfile", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userProfile: ", p), err) }
  if err := p.UserProfile.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userProfile: ", p), err) }
  return err
}

func (p *EngineServicePredictArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("movieId", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:movieId: ", p), err) }
  if err := oprot.WriteI64(int64(p.MovieId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.movieId (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:movieId: ", p), err) }
  return err
}

func (p *EngineServicePredictArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServicePredictArgs(%+v)", *p)
}

// Attributes:
//  - Success
type EngineServicePredictResult struct {
  Success *int64 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEngineServicePredictResult() *EngineServicePredictResult {
  return &EngineServicePredictResult{}
}

var EngineServicePredictResult_Success_DEFAULT int64
func (p *EngineServicePredictResult) GetSuccess() int64 {
  if !p.IsSetSuccess() {
    return EngineServicePredictResult_Success_DEFAULT
  }
return *p.Success
}
func (p *EngineServicePredictResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *EngineServicePredictResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServicePredictResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *EngineServicePredictResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Predict_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServicePredictResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI64(int64(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *EngineServicePredictResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServicePredictResult(%+v)", *p)
}

// Attributes:
//  - UserProfile
//  - Topk
type EngineServiceRecommendArgs struct {
  UserProfile *UserProfile `thrift:"userProfile,1" db:"userProfile" json:"userProfile"`
  Topk int32 `thrift:"topk,2" db:"topk" json:"topk"`
}

func NewEngineServiceRecommendArgs() *EngineServiceRecommendArgs {
  return &EngineServiceRecommendArgs{}
}

var EngineServiceRecommendArgs_UserProfile_DEFAULT *UserProfile
func (p *EngineServiceRecommendArgs) GetUserProfile() *UserProfile {
  if !p.IsSetUserProfile() {
    return EngineServiceRecommendArgs_UserProfile_DEFAULT
  }
return p.UserProfile
}

func (p *EngineServiceRecommendArgs) GetTopk() int32 {
  return p.Topk
}
func (p *EngineServiceRecommendArgs) IsSetUserProfile() bool {
  return p.UserProfile != nil
}

func (p *EngineServiceRecommendArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServiceRecommendArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.UserProfile = &UserProfile{}
  if err := p.UserProfile.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
  }
  return nil
}

func (p *EngineServiceRecommendArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Topk = v
}
  return nil
}

func (p *EngineServiceRecommendArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Recommend_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServiceRecommendArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("userProfile", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userProfile: ", p), err) }
  if err := p.UserProfile.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userProfile: ", p), err) }
  return err
}

func (p *EngineServiceRecommendArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("topk", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:topk: ", p), err) }
  if err := oprot.WriteI32(int32(p.Topk)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.topk (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:topk: ", p), err) }
  return err
}

func (p *EngineServiceRecommendArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServiceRecommendArgs(%+v)", *p)
}

// Attributes:
//  - Success
type EngineServiceRecommendResult struct {
  Success []int64 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEngineServiceRecommendResult() *EngineServiceRecommendResult {
  return &EngineServiceRecommendResult{}
}

var EngineServiceRecommendResult_Success_DEFAULT []int64

func (p *EngineServiceRecommendResult) GetSuccess() []int64 {
  return p.Success
}
func (p *EngineServiceRecommendResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *EngineServiceRecommendResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *EngineServiceRecommendResult)  ReadField0(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]int64, 0, size)
  p.Success =  tSlice
  for i := 0; i < size; i ++ {
var _elem8 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem8 = v
}
    p.Success = append(p.Success, _elem8)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *EngineServiceRecommendResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Recommend_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EngineServiceRecommendResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteListBegin(thrift.I64, len(p.Success)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Success {
      if err := oprot.WriteI64(int64(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *EngineServiceRecommendResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EngineServiceRecommendResult(%+v)", *p)
}


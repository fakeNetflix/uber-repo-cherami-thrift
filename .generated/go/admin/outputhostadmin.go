// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package admin

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type OutputHostAdmin interface {
	// Parameters:
	//  - Request
	ConsumerGroupsUpdated(request *ConsumerGroupsUpdatedRequest) (err error)
	// Parameters:
	//  - Request
	UnloadConsumerGroups(request *UnloadConsumerGroupsRequest) (err error)
}

type OutputHostAdminClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewOutputHostAdminClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *OutputHostAdminClient {
	return &OutputHostAdminClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewOutputHostAdminClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *OutputHostAdminClient {
	return &OutputHostAdminClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Request
func (p *OutputHostAdminClient) ConsumerGroupsUpdated(request *ConsumerGroupsUpdatedRequest) (err error) {
	if err = p.sendConsumerGroupsUpdated(request); err != nil {
		return
	}
	return p.recvConsumerGroupsUpdated()
}

func (p *OutputHostAdminClient) sendConsumerGroupsUpdated(request *ConsumerGroupsUpdatedRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("consumerGroupsUpdated", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := OutputHostAdminConsumerGroupsUpdatedArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *OutputHostAdminClient) recvConsumerGroupsUpdated() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "consumerGroupsUpdated" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "consumerGroupsUpdated failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "consumerGroupsUpdated failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error15 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error16 error
		error16, err = error15.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error16
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "consumerGroupsUpdated failed: invalid message type")
		return
	}
	result := OutputHostAdminConsumerGroupsUpdatedResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

// Parameters:
//  - Request
func (p *OutputHostAdminClient) UnloadConsumerGroups(request *UnloadConsumerGroupsRequest) (err error) {
	if err = p.sendUnloadConsumerGroups(request); err != nil {
		return
	}
	return p.recvUnloadConsumerGroups()
}

func (p *OutputHostAdminClient) sendUnloadConsumerGroups(request *UnloadConsumerGroupsRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("unloadConsumerGroups", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := OutputHostAdminUnloadConsumerGroupsArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *OutputHostAdminClient) recvUnloadConsumerGroups() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "unloadConsumerGroups" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "unloadConsumerGroups failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "unloadConsumerGroups failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error17 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error18 error
		error18, err = error17.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error18
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "unloadConsumerGroups failed: invalid message type")
		return
	}
	result := OutputHostAdminUnloadConsumerGroupsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type OutputHostAdminProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      OutputHostAdmin
}

func (p *OutputHostAdminProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *OutputHostAdminProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *OutputHostAdminProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewOutputHostAdminProcessor(handler OutputHostAdmin) *OutputHostAdminProcessor {

	self19 := &OutputHostAdminProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self19.processorMap["consumerGroupsUpdated"] = &outputHostAdminProcessorConsumerGroupsUpdated{handler: handler}
	self19.processorMap["unloadConsumerGroups"] = &outputHostAdminProcessorUnloadConsumerGroups{handler: handler}
	return self19
}

func (p *OutputHostAdminProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x20 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x20.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x20

}

type outputHostAdminProcessorConsumerGroupsUpdated struct {
	handler OutputHostAdmin
}

func (p *outputHostAdminProcessorConsumerGroupsUpdated) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := OutputHostAdminConsumerGroupsUpdatedArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("consumerGroupsUpdated", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := OutputHostAdminConsumerGroupsUpdatedResult{}
	var err2 error
	if err2 = p.handler.ConsumerGroupsUpdated(args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing consumerGroupsUpdated: "+err2.Error())
		oprot.WriteMessageBegin("consumerGroupsUpdated", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("consumerGroupsUpdated", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type outputHostAdminProcessorUnloadConsumerGroups struct {
	handler OutputHostAdmin
}

func (p *outputHostAdminProcessorUnloadConsumerGroups) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := OutputHostAdminUnloadConsumerGroupsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("unloadConsumerGroups", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := OutputHostAdminUnloadConsumerGroupsResult{}
	var err2 error
	if err2 = p.handler.UnloadConsumerGroups(args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing unloadConsumerGroups: "+err2.Error())
		oprot.WriteMessageBegin("unloadConsumerGroups", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("unloadConsumerGroups", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type OutputHostAdminConsumerGroupsUpdatedArgs struct {
	Request *ConsumerGroupsUpdatedRequest `thrift:"request,1" json:"request"`
}

func NewOutputHostAdminConsumerGroupsUpdatedArgs() *OutputHostAdminConsumerGroupsUpdatedArgs {
	return &OutputHostAdminConsumerGroupsUpdatedArgs{}
}

var OutputHostAdminConsumerGroupsUpdatedArgs_Request_DEFAULT *ConsumerGroupsUpdatedRequest

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) GetRequest() *ConsumerGroupsUpdatedRequest {
	if !p.IsSetRequest() {
		return OutputHostAdminConsumerGroupsUpdatedArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *OutputHostAdminConsumerGroupsUpdatedArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
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

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &ConsumerGroupsUpdatedRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("consumerGroupsUpdated_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *OutputHostAdminConsumerGroupsUpdatedArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OutputHostAdminConsumerGroupsUpdatedArgs(%+v)", *p)
}

type OutputHostAdminConsumerGroupsUpdatedResult struct {
}

func NewOutputHostAdminConsumerGroupsUpdatedResult() *OutputHostAdminConsumerGroupsUpdatedResult {
	return &OutputHostAdminConsumerGroupsUpdatedResult{}
}

func (p *OutputHostAdminConsumerGroupsUpdatedResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *OutputHostAdminConsumerGroupsUpdatedResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("consumerGroupsUpdated_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OutputHostAdminConsumerGroupsUpdatedResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OutputHostAdminConsumerGroupsUpdatedResult(%+v)", *p)
}

// Attributes:
//  - Request
type OutputHostAdminUnloadConsumerGroupsArgs struct {
	Request *UnloadConsumerGroupsRequest `thrift:"request,1" json:"request"`
}

func NewOutputHostAdminUnloadConsumerGroupsArgs() *OutputHostAdminUnloadConsumerGroupsArgs {
	return &OutputHostAdminUnloadConsumerGroupsArgs{}
}

var OutputHostAdminUnloadConsumerGroupsArgs_Request_DEFAULT *UnloadConsumerGroupsRequest

func (p *OutputHostAdminUnloadConsumerGroupsArgs) GetRequest() *UnloadConsumerGroupsRequest {
	if !p.IsSetRequest() {
		return OutputHostAdminUnloadConsumerGroupsArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *OutputHostAdminUnloadConsumerGroupsArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *OutputHostAdminUnloadConsumerGroupsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
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

func (p *OutputHostAdminUnloadConsumerGroupsArgs) readField1(iprot thrift.TProtocol) error {
	p.Request = &UnloadConsumerGroupsRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *OutputHostAdminUnloadConsumerGroupsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("unloadConsumerGroups_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OutputHostAdminUnloadConsumerGroupsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *OutputHostAdminUnloadConsumerGroupsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OutputHostAdminUnloadConsumerGroupsArgs(%+v)", *p)
}

type OutputHostAdminUnloadConsumerGroupsResult struct {
}

func NewOutputHostAdminUnloadConsumerGroupsResult() *OutputHostAdminUnloadConsumerGroupsResult {
	return &OutputHostAdminUnloadConsumerGroupsResult{}
}

func (p *OutputHostAdminUnloadConsumerGroupsResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *OutputHostAdminUnloadConsumerGroupsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("unloadConsumerGroups_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OutputHostAdminUnloadConsumerGroupsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OutputHostAdminUnloadConsumerGroupsResult(%+v)", *p)
}
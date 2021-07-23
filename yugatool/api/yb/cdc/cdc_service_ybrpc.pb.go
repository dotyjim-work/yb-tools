// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// The following only applies to changes made to this file as part of YugaByte development.
//
// Portions Copyright (c) YugaByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.  See the License for the specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-ybrpc. DO NOT EDIT.

package cdc

import (
	"github.com/go-logr/logr"
	"github.com/yugabyte/yb-tools/protoc-gen-ybrpc/pkg/message"
)

// service: yb.cdc.CDCService
// service: CDCService
type CDCService interface {
	CreateCDCStream(request *CreateCDCStreamRequestPB) (*CreateCDCStreamResponsePB, error)
	DeleteCDCStream(request *DeleteCDCStreamRequestPB) (*DeleteCDCStreamResponsePB, error)
	ListTablets(request *ListTabletsRequestPB) (*ListTabletsResponsePB, error)
	GetChanges(request *GetChangesRequestPB) (*GetChangesResponsePB, error)
	GetCheckpoint(request *GetCheckpointRequestPB) (*GetCheckpointResponsePB, error)
	UpdateCdcReplicatedIndex(request *UpdateCdcReplicatedIndexRequestPB) (*UpdateCdcReplicatedIndexResponsePB, error)
	BootstrapProducer(request *BootstrapProducerRequestPB) (*BootstrapProducerResponsePB, error)
	GetLatestEntryOpId(request *GetLatestEntryOpIdRequestPB) (*GetLatestEntryOpIdResponsePB, error)
}

type CDCServiceImpl struct {
	Log       logr.Logger
	Messenger message.Messenger
}

func (s *CDCServiceImpl) CreateCDCStream(request *CreateCDCStreamRequestPB) (*CreateCDCStreamResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "CreateCDCStream", "message", request)
	response := &CreateCDCStreamResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "CreateCDCStream", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "CreateCDCStream", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) DeleteCDCStream(request *DeleteCDCStreamRequestPB) (*DeleteCDCStreamResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "DeleteCDCStream", "message", request)
	response := &DeleteCDCStreamResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "DeleteCDCStream", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "DeleteCDCStream", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) ListTablets(request *ListTabletsRequestPB) (*ListTabletsResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "ListTablets", "message", request)
	response := &ListTabletsResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "ListTablets", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "ListTablets", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) GetChanges(request *GetChangesRequestPB) (*GetChangesResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "GetChanges", "message", request)
	response := &GetChangesResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "GetChanges", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "GetChanges", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) GetCheckpoint(request *GetCheckpointRequestPB) (*GetCheckpointResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "GetCheckpoint", "message", request)
	response := &GetCheckpointResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "GetCheckpoint", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "GetCheckpoint", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) UpdateCdcReplicatedIndex(request *UpdateCdcReplicatedIndexRequestPB) (*UpdateCdcReplicatedIndexResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "UpdateCdcReplicatedIndex", "message", request)
	response := &UpdateCdcReplicatedIndexResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "UpdateCdcReplicatedIndex", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "UpdateCdcReplicatedIndex", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) BootstrapProducer(request *BootstrapProducerRequestPB) (*BootstrapProducerResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "BootstrapProducer", "message", request)
	response := &BootstrapProducerResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "BootstrapProducer", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "BootstrapProducer", "message", response)

	return response, nil
}

func (s *CDCServiceImpl) GetLatestEntryOpId(request *GetLatestEntryOpIdRequestPB) (*GetLatestEntryOpIdResponsePB, error) {
	s.Log.V(1).Info("sending RPC message", "service", "yb.cdc.CDCService", "method", "GetLatestEntryOpId", "message", request)
	response := &GetLatestEntryOpIdResponsePB{}

	err := s.Messenger.SendMessage("yb.cdc.CDCService", "GetLatestEntryOpId", request.ProtoReflect().Interface(), response.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	s.Log.V(1).Info("received RPC response", "service", "yb.cdc.CDCService", "method", "GetLatestEntryOpId", "message", response)

	return response, nil
}

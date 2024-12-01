// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var api_jack_pb = require('../api/jack_pb.js');

function serialize_server_GetStudentRequest(arg) {
  if (!(arg instanceof api_jack_pb.GetStudentRequest)) {
    throw new Error('Expected argument of type server.GetStudentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_server_GetStudentRequest(buffer_arg) {
  return api_jack_pb.GetStudentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_server_GetStudentResponse(arg) {
  if (!(arg instanceof api_jack_pb.GetStudentResponse)) {
    throw new Error('Expected argument of type server.GetStudentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_server_GetStudentResponse(buffer_arg) {
  return api_jack_pb.GetStudentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var JackService = exports.JackService = {
  getStudent: {
    path: '/server.Jack/GetStudent',
    requestStream: false,
    responseStream: false,
    requestType: api_jack_pb.GetStudentRequest,
    responseType: api_jack_pb.GetStudentResponse,
    requestSerialize: serialize_server_GetStudentRequest,
    requestDeserialize: deserialize_server_GetStudentRequest,
    responseSerialize: serialize_server_GetStudentResponse,
    responseDeserialize: deserialize_server_GetStudentResponse,
  },
};

exports.JackClient = grpc.makeGenericClientConstructor(JackService);

# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import stringdb_pb2 as stringdb__pb2


class StringDbStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetValue = channel.unary_unary(
                '/stringdb.StringDb/GetValue',
                request_serializer=stringdb__pb2.GetValueRequest.SerializeToString,
                response_deserializer=stringdb__pb2.GetValueReply.FromString,
                )
        self.SetValue = channel.unary_unary(
                '/stringdb.StringDb/SetValue',
                request_serializer=stringdb__pb2.SetValueRequest.SerializeToString,
                response_deserializer=stringdb__pb2.SetValueReply.FromString,
                )
        self.CountValue = channel.unary_unary(
                '/stringdb.StringDb/CountValue',
                request_serializer=stringdb__pb2.CountValueRequest.SerializeToString,
                response_deserializer=stringdb__pb2.CountValueReply.FromString,
                )


class StringDbServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetValue(self, request, context):
        """Get the value stored on the server for a given key
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetValue(self, request, context):
        """Set the server's value for a given key
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CountValue(self, request, context):
        """Count the size of the server's value for a given key
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_StringDbServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetValue': grpc.unary_unary_rpc_method_handler(
                    servicer.GetValue,
                    request_deserializer=stringdb__pb2.GetValueRequest.FromString,
                    response_serializer=stringdb__pb2.GetValueReply.SerializeToString,
            ),
            'SetValue': grpc.unary_unary_rpc_method_handler(
                    servicer.SetValue,
                    request_deserializer=stringdb__pb2.SetValueRequest.FromString,
                    response_serializer=stringdb__pb2.SetValueReply.SerializeToString,
            ),
            'CountValue': grpc.unary_unary_rpc_method_handler(
                    servicer.CountValue,
                    request_deserializer=stringdb__pb2.CountValueRequest.FromString,
                    response_serializer=stringdb__pb2.CountValueReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'stringdb.StringDb', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class StringDb(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetValue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/stringdb.StringDb/GetValue',
            stringdb__pb2.GetValueRequest.SerializeToString,
            stringdb__pb2.GetValueReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetValue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/stringdb.StringDb/SetValue',
            stringdb__pb2.SetValueRequest.SerializeToString,
            stringdb__pb2.SetValueReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CountValue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/stringdb.StringDb/CountValue',
            stringdb__pb2.CountValueRequest.SerializeToString,
            stringdb__pb2.CountValueReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

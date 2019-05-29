# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from datapipe.metrics import services_pb2 as datapipe_dot_metrics_dot_services__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


class MetricsServiceStub(object):
  """*
  Service for providing data stored in the backend
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreatePodMetrics = channel.unary_unary(
        '/containersai.datapipe.metrics.MetricsService/CreatePodMetrics',
        request_serializer=datapipe_dot_metrics_dot_services__pb2.CreatePodMetricsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreateNodeMetrics = channel.unary_unary(
        '/containersai.datapipe.metrics.MetricsService/CreateNodeMetrics',
        request_serializer=datapipe_dot_metrics_dot_services__pb2.CreateNodeMetricsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.ListPodMetrics = channel.unary_unary(
        '/containersai.datapipe.metrics.MetricsService/ListPodMetrics',
        request_serializer=datapipe_dot_metrics_dot_services__pb2.ListPodMetricsRequest.SerializeToString,
        response_deserializer=datapipe_dot_metrics_dot_services__pb2.ListPodMetricsResponse.FromString,
        )
    self.ListNodeMetrics = channel.unary_unary(
        '/containersai.datapipe.metrics.MetricsService/ListNodeMetrics',
        request_serializer=datapipe_dot_metrics_dot_services__pb2.ListNodeMetricsRequest.SerializeToString,
        response_deserializer=datapipe_dot_metrics_dot_services__pb2.ListNodeMetricsResponse.FromString,
        )


class MetricsServiceServicer(object):
  """*
  Service for providing data stored in the backend
  """

  def CreatePodMetrics(self, request, context):
    """Used to create metrics data of pods
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNodeMetrics(self, request, context):
    """Used to create metrics data of nodes
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPodMetrics(self, request, context):
    """Used to list pod metric data
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListNodeMetrics(self, request, context):
    """Used to list node metric data
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_MetricsServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreatePodMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePodMetrics,
          request_deserializer=datapipe_dot_metrics_dot_services__pb2.CreatePodMetricsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreateNodeMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNodeMetrics,
          request_deserializer=datapipe_dot_metrics_dot_services__pb2.CreateNodeMetricsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'ListPodMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodMetrics,
          request_deserializer=datapipe_dot_metrics_dot_services__pb2.ListPodMetricsRequest.FromString,
          response_serializer=datapipe_dot_metrics_dot_services__pb2.ListPodMetricsResponse.SerializeToString,
      ),
      'ListNodeMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.ListNodeMetrics,
          request_deserializer=datapipe_dot_metrics_dot_services__pb2.ListNodeMetricsRequest.FromString,
          response_serializer=datapipe_dot_metrics_dot_services__pb2.ListNodeMetricsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'containersai.datapipe.metrics.MetricsService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))

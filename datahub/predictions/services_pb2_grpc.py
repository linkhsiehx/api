# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from datahub.predictions import services_pb2 as datahub_dot_predictions_dot_services__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


class PredictionsServiceStub(object):
  """*
  Service for providing data stored in the backend
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreatePodPredictions = channel.unary_unary(
        '/containersai.datahub.predictions.PredictionsService/CreatePodPredictions',
        request_serializer=datahub_dot_predictions_dot_services__pb2.CreatePodPredictionsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreateNodePredictions = channel.unary_unary(
        '/containersai.datahub.predictions.PredictionsService/CreateNodePredictions',
        request_serializer=datahub_dot_predictions_dot_services__pb2.CreateNodePredictionsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.ListPodPredictions = channel.unary_unary(
        '/containersai.datahub.predictions.PredictionsService/ListPodPredictions',
        request_serializer=datahub_dot_predictions_dot_services__pb2.ListPodPredictionsRequest.SerializeToString,
        response_deserializer=datahub_dot_predictions_dot_services__pb2.ListPodPredictionsResponse.FromString,
        )
    self.ListNodePredictions = channel.unary_unary(
        '/containersai.datahub.predictions.PredictionsService/ListNodePredictions',
        request_serializer=datahub_dot_predictions_dot_services__pb2.ListNodePredictionsRequest.SerializeToString,
        response_deserializer=datahub_dot_predictions_dot_services__pb2.ListNodePredictionsResponse.FromString,
        )


class PredictionsServiceServicer(object):
  """*
  Service for providing data stored in the backend
  """

  def CreatePodPredictions(self, request, context):
    """Used to create predictions of pods
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNodePredictions(self, request, context):
    """Used to create predictions of nodes
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPodPredictions(self, request, context):
    """Used to list pod predictions
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListNodePredictions(self, request, context):
    """Used to list node predictions
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PredictionsServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreatePodPredictions': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePodPredictions,
          request_deserializer=datahub_dot_predictions_dot_services__pb2.CreatePodPredictionsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreateNodePredictions': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNodePredictions,
          request_deserializer=datahub_dot_predictions_dot_services__pb2.CreateNodePredictionsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'ListPodPredictions': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodPredictions,
          request_deserializer=datahub_dot_predictions_dot_services__pb2.ListPodPredictionsRequest.FromString,
          response_serializer=datahub_dot_predictions_dot_services__pb2.ListPodPredictionsResponse.SerializeToString,
      ),
      'ListNodePredictions': grpc.unary_unary_rpc_method_handler(
          servicer.ListNodePredictions,
          request_deserializer=datahub_dot_predictions_dot_services__pb2.ListNodePredictionsRequest.FromString,
          response_serializer=datahub_dot_predictions_dot_services__pb2.ListNodePredictionsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'containersai.datahub.predictions.PredictionsService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))

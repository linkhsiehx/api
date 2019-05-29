# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/recommendations/recommendations.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.recommendations import types_pb2 as datahub_dot_recommendations_dot_types__pb2
from datahub.resources import resources_pb2 as datahub_dot_resources_dot_resources__pb2
from datahub.resources import types_pb2 as datahub_dot_resources_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/recommendations/recommendations.proto',
  package='containersai.datahub.recommendations',
  syntax='proto3',
  serialized_options=_b('Z4github.com/containers-ai/api/datahub/recommendations'),
  serialized_pb=_b('\n-datahub/recommendations/recommendations.proto\x12$containersai.datahub.recommendations\x1a#datahub/recommendations/types.proto\x1a!datahub/resources/resources.proto\x1a\x1d\x64\x61tahub/resources/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb4\x02\n\x17\x43ontainerRecommendation\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bgranularity\x18\x02 \x01(\x03\x12.\n\nstart_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12G\n\tresources\x18\x05 \x01(\x0b\x32\x34.containersai.datahub.resources.ResourceRequirements\x12O\n\x11initial_resources\x18\x06 \x01(\x0b\x32\x34.containersai.datahub.resources.ResourceRequirements\"\x91\x03\n\x11PodRecommendation\x12G\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32..containersai.datahub.resources.NamespacedName\x12\x42\n\x0etop_controller\x18\x02 \x01(\x0b\x32*.containersai.datahub.resources.Controller\x12\x19\n\x11recommendation_id\x18\x03 \x01(\t\x12 \n\x18\x61pply_recommendation_now\x18\x04 \x01(\x08\x12P\n\x11\x61ssign_pod_policy\x18\x05 \x01(\x0b\x32\x35.containersai.datahub.recommendations.AssignPodPolicy\x12`\n\x19\x63ontainer_recommendations\x18\x06 \x03(\x0b\x32=.containersai.datahub.recommendations.ContainerRecommendation\"\xdc\x01\n\x18\x43ontrollerRecommendation\x12_\n\x13recommendation_type\x18\x01 \x01(\x0e\x32\x42.containersai.datahub.recommendations.ControllerRecommendationType\x12_\n\x13recommendation_spec\x18\x02 \x01(\x0b\x32\x42.containersai.datahub.recommendations.ControllerRecommendationSpecB6Z4github.com/containers-ai/api/datahub/recommendationsb\x06proto3')
  ,
  dependencies=[datahub_dot_recommendations_dot_types__pb2.DESCRIPTOR,datahub_dot_resources_dot_resources__pb2.DESCRIPTOR,datahub_dot_resources_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_CONTAINERRECOMMENDATION = _descriptor.Descriptor(
  name='ContainerRecommendation',
  full_name='containersai.datahub.recommendations.ContainerRecommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.recommendations.ContainerRecommendation.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='granularity', full_name='containersai.datahub.recommendations.ContainerRecommendation.granularity', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.datahub.recommendations.ContainerRecommendation.start_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.datahub.recommendations.ContainerRecommendation.end_time', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='containersai.datahub.recommendations.ContainerRecommendation.resources', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initial_resources', full_name='containersai.datahub.recommendations.ContainerRecommendation.initial_resources', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=224,
  serialized_end=532,
)


_PODRECOMMENDATION = _descriptor.Descriptor(
  name='PodRecommendation',
  full_name='containersai.datahub.recommendations.PodRecommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containersai.datahub.recommendations.PodRecommendation.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='top_controller', full_name='containersai.datahub.recommendations.PodRecommendation.top_controller', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='recommendation_id', full_name='containersai.datahub.recommendations.PodRecommendation.recommendation_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='apply_recommendation_now', full_name='containersai.datahub.recommendations.PodRecommendation.apply_recommendation_now', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='assign_pod_policy', full_name='containersai.datahub.recommendations.PodRecommendation.assign_pod_policy', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container_recommendations', full_name='containersai.datahub.recommendations.PodRecommendation.container_recommendations', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=535,
  serialized_end=936,
)


_CONTROLLERRECOMMENDATION = _descriptor.Descriptor(
  name='ControllerRecommendation',
  full_name='containersai.datahub.recommendations.ControllerRecommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='recommendation_type', full_name='containersai.datahub.recommendations.ControllerRecommendation.recommendation_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='recommendation_spec', full_name='containersai.datahub.recommendations.ControllerRecommendation.recommendation_spec', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=939,
  serialized_end=1159,
)

_CONTAINERRECOMMENDATION.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTAINERRECOMMENDATION.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CONTAINERRECOMMENDATION.fields_by_name['resources'].message_type = datahub_dot_resources_dot_types__pb2._RESOURCEREQUIREMENTS
_CONTAINERRECOMMENDATION.fields_by_name['initial_resources'].message_type = datahub_dot_resources_dot_types__pb2._RESOURCEREQUIREMENTS
_PODRECOMMENDATION.fields_by_name['namespaced_name'].message_type = datahub_dot_resources_dot_types__pb2._NAMESPACEDNAME
_PODRECOMMENDATION.fields_by_name['top_controller'].message_type = datahub_dot_resources_dot_resources__pb2._CONTROLLER
_PODRECOMMENDATION.fields_by_name['assign_pod_policy'].message_type = datahub_dot_recommendations_dot_types__pb2._ASSIGNPODPOLICY
_PODRECOMMENDATION.fields_by_name['container_recommendations'].message_type = _CONTAINERRECOMMENDATION
_CONTROLLERRECOMMENDATION.fields_by_name['recommendation_type'].enum_type = datahub_dot_recommendations_dot_types__pb2._CONTROLLERRECOMMENDATIONTYPE
_CONTROLLERRECOMMENDATION.fields_by_name['recommendation_spec'].message_type = datahub_dot_recommendations_dot_types__pb2._CONTROLLERRECOMMENDATIONSPEC
DESCRIPTOR.message_types_by_name['ContainerRecommendation'] = _CONTAINERRECOMMENDATION
DESCRIPTOR.message_types_by_name['PodRecommendation'] = _PODRECOMMENDATION
DESCRIPTOR.message_types_by_name['ControllerRecommendation'] = _CONTROLLERRECOMMENDATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerRecommendation = _reflection.GeneratedProtocolMessageType('ContainerRecommendation', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINERRECOMMENDATION,
  __module__ = 'datahub.recommendations.recommendations_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.recommendations.ContainerRecommendation)
  ))
_sym_db.RegisterMessage(ContainerRecommendation)

PodRecommendation = _reflection.GeneratedProtocolMessageType('PodRecommendation', (_message.Message,), dict(
  DESCRIPTOR = _PODRECOMMENDATION,
  __module__ = 'datahub.recommendations.recommendations_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.recommendations.PodRecommendation)
  ))
_sym_db.RegisterMessage(PodRecommendation)

ControllerRecommendation = _reflection.GeneratedProtocolMessageType('ControllerRecommendation', (_message.Message,), dict(
  DESCRIPTOR = _CONTROLLERRECOMMENDATION,
  __module__ = 'datahub.recommendations.recommendations_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.recommendations.ControllerRecommendation)
  ))
_sym_db.RegisterMessage(ControllerRecommendation)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: voltha_protos/health.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from voltha_protos import meta_pb2 as voltha__protos_dot_meta__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='voltha_protos/health.proto',
  package='voltha',
  syntax='proto3',
  serialized_options=_b('Z+github.com/opencord/voltha-protos/go/voltha'),
  serialized_pb=_b('\n\x1avoltha_protos/health.proto\x12\x06voltha\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x18voltha_protos/meta.proto\"}\n\x0cHealthStatus\x12\x36\n\x05state\x18\x01 \x01(\x0e\x32 .voltha.HealthStatus.HealthStateB\x05\xe8\xf6\xcd\x1d\x01\"5\n\x0bHealthState\x12\x0b\n\x07HEALTHY\x10\x00\x12\x0e\n\nOVERLOADED\x10\x01\x12\t\n\x05\x44YING\x10\x02\x32\x61\n\rHealthService\x12P\n\x0fGetHealthStatus\x12\x16.google.protobuf.Empty\x1a\x14.voltha.HealthStatus\"\x0f\x82\xd3\xe4\x93\x02\t\x12\x07/healthB-Z+github.com/opencord/voltha-protos/go/volthab\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,voltha__protos_dot_meta__pb2.DESCRIPTOR,])



_HEALTHSTATUS_HEALTHSTATE = _descriptor.EnumDescriptor(
  name='HealthState',
  full_name='voltha.HealthStatus.HealthState',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HEALTHY', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OVERLOADED', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DYING', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=195,
  serialized_end=248,
)
_sym_db.RegisterEnumDescriptor(_HEALTHSTATUS_HEALTHSTATE)


_HEALTHSTATUS = _descriptor.Descriptor(
  name='HealthStatus',
  full_name='voltha.HealthStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='state', full_name='voltha.HealthStatus.state', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\350\366\315\035\001'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _HEALTHSTATUS_HEALTHSTATE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=123,
  serialized_end=248,
)

_HEALTHSTATUS.fields_by_name['state'].enum_type = _HEALTHSTATUS_HEALTHSTATE
_HEALTHSTATUS_HEALTHSTATE.containing_type = _HEALTHSTATUS
DESCRIPTOR.message_types_by_name['HealthStatus'] = _HEALTHSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HealthStatus = _reflection.GeneratedProtocolMessageType('HealthStatus', (_message.Message,), dict(
  DESCRIPTOR = _HEALTHSTATUS,
  __module__ = 'voltha_protos.health_pb2'
  # @@protoc_insertion_point(class_scope:voltha.HealthStatus)
  ))
_sym_db.RegisterMessage(HealthStatus)


DESCRIPTOR._options = None
_HEALTHSTATUS.fields_by_name['state']._options = None

_HEALTHSERVICE = _descriptor.ServiceDescriptor(
  name='HealthService',
  full_name='voltha.HealthService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=250,
  serialized_end=347,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetHealthStatus',
    full_name='voltha.HealthService.GetHealthStatus',
    index=0,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_HEALTHSTATUS,
    serialized_options=_b('\202\323\344\223\002\t\022\007/health'),
  ),
])
_sym_db.RegisterServiceDescriptor(_HEALTHSERVICE)

DESCRIPTOR.services_by_name['HealthService'] = _HEALTHSERVICE

# @@protoc_insertion_point(module_scope)

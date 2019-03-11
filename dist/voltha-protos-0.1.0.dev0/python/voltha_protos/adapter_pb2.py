# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: voltha_protos/adapter.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from voltha_protos import common_pb2 as voltha__protos_dot_common__pb2
from voltha_protos import meta_pb2 as voltha__protos_dot_meta__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='voltha_protos/adapter.proto',
  package='voltha',
  syntax='proto3',
  serialized_options=_b('Z+github.com/opencord/voltha-go/protos/voltha'),
  serialized_pb=_b('\n\x1bvoltha_protos/adapter.proto\x12\x06voltha\x1a\x19google/protobuf/any.proto\x1a\x1avoltha_protos/common.proto\x1a\x18voltha_protos/meta.proto\"n\n\rAdapterConfig\x12,\n\tlog_level\x18\x01 \x01(\x0e\x32\x19.voltha.LogLevel.LogLevel\x12/\n\x11\x61\x64\x64itional_config\x18@ \x01(\x0b\x32\x14.google.protobuf.Any\"\xcb\x01\n\x07\x41\x64\x61pter\x12\x11\n\x02id\x18\x01 \x01(\tB\x05\xe8\xf6\xcd\x1d\x01\x12\x15\n\x06vendor\x18\x02 \x01(\tB\x05\xe8\xf6\xcd\x1d\x01\x12\x16\n\x07version\x18\x03 \x01(\tB\x05\xe8\xf6\xcd\x1d\x01\x12%\n\x06\x63onfig\x18\x10 \x01(\x0b\x32\x15.voltha.AdapterConfig\x12;\n\x16\x61\x64\x64itional_description\x18@ \x01(\x0b\x32\x14.google.protobuf.AnyB\x05\xe8\xf6\xcd\x1d\x01\x12\x1a\n\x12logical_device_ids\x18\x04 \x03(\t\"*\n\x08\x41\x64\x61pters\x12\x1e\n\x05items\x18\x01 \x03(\x0b\x32\x0f.voltha.AdapterB-Z+github.com/opencord/voltha-go/protos/volthab\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_any__pb2.DESCRIPTOR,voltha__protos_dot_common__pb2.DESCRIPTOR,voltha__protos_dot_meta__pb2.DESCRIPTOR,])




_ADAPTERCONFIG = _descriptor.Descriptor(
  name='AdapterConfig',
  full_name='voltha.AdapterConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='log_level', full_name='voltha.AdapterConfig.log_level', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='additional_config', full_name='voltha.AdapterConfig.additional_config', index=1,
      number=64, type=11, cpp_type=10, label=1,
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
  serialized_start=120,
  serialized_end=230,
)


_ADAPTER = _descriptor.Descriptor(
  name='Adapter',
  full_name='voltha.Adapter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='voltha.Adapter.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\350\366\315\035\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vendor', full_name='voltha.Adapter.vendor', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\350\366\315\035\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='voltha.Adapter.version', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\350\366\315\035\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config', full_name='voltha.Adapter.config', index=3,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='additional_description', full_name='voltha.Adapter.additional_description', index=4,
      number=64, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\350\366\315\035\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='logical_device_ids', full_name='voltha.Adapter.logical_device_ids', index=5,
      number=4, type=9, cpp_type=9, label=3,
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
  serialized_start=233,
  serialized_end=436,
)


_ADAPTERS = _descriptor.Descriptor(
  name='Adapters',
  full_name='voltha.Adapters',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='voltha.Adapters.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=438,
  serialized_end=480,
)

_ADAPTERCONFIG.fields_by_name['log_level'].enum_type = voltha__protos_dot_common__pb2._LOGLEVEL_LOGLEVEL
_ADAPTERCONFIG.fields_by_name['additional_config'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_ADAPTER.fields_by_name['config'].message_type = _ADAPTERCONFIG
_ADAPTER.fields_by_name['additional_description'].message_type = google_dot_protobuf_dot_any__pb2._ANY
_ADAPTERS.fields_by_name['items'].message_type = _ADAPTER
DESCRIPTOR.message_types_by_name['AdapterConfig'] = _ADAPTERCONFIG
DESCRIPTOR.message_types_by_name['Adapter'] = _ADAPTER
DESCRIPTOR.message_types_by_name['Adapters'] = _ADAPTERS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AdapterConfig = _reflection.GeneratedProtocolMessageType('AdapterConfig', (_message.Message,), dict(
  DESCRIPTOR = _ADAPTERCONFIG,
  __module__ = 'voltha_protos.adapter_pb2'
  # @@protoc_insertion_point(class_scope:voltha.AdapterConfig)
  ))
_sym_db.RegisterMessage(AdapterConfig)

Adapter = _reflection.GeneratedProtocolMessageType('Adapter', (_message.Message,), dict(
  DESCRIPTOR = _ADAPTER,
  __module__ = 'voltha_protos.adapter_pb2'
  # @@protoc_insertion_point(class_scope:voltha.Adapter)
  ))
_sym_db.RegisterMessage(Adapter)

Adapters = _reflection.GeneratedProtocolMessageType('Adapters', (_message.Message,), dict(
  DESCRIPTOR = _ADAPTERS,
  __module__ = 'voltha_protos.adapter_pb2'
  # @@protoc_insertion_point(class_scope:voltha.Adapters)
  ))
_sym_db.RegisterMessage(Adapters)


DESCRIPTOR._options = None
_ADAPTER.fields_by_name['id']._options = None
_ADAPTER.fields_by_name['vendor']._options = None
_ADAPTER.fields_by_name['version']._options = None
_ADAPTER.fields_by_name['additional_description']._options = None
# @@protoc_insertion_point(module_scope)

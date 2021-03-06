# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: voltha_protos/schema.proto

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


DESCRIPTOR = _descriptor.FileDescriptor(
  name='voltha_protos/schema.proto',
  package='schema',
  syntax='proto3',
  serialized_options=_b('Z+github.com/opencord/voltha-protos/go/schema'),
  serialized_pb=_b('\n\x1avoltha_protos/schema.proto\x12\x06schema\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"A\n\tProtoFile\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\r\n\x05proto\x18\x02 \x01(\t\x12\x12\n\ndescriptor\x18\x03 \x01(\x0c\"U\n\x07Schemas\x12!\n\x06protos\x18\x01 \x03(\x0b\x32\x11.schema.ProtoFile\x12\x14\n\x0cswagger_from\x18\x02 \x01(\t\x12\x11\n\tyang_from\x18\x03 \x01(\t2V\n\rSchemaService\x12\x45\n\tGetSchema\x12\x16.google.protobuf.Empty\x1a\x0f.schema.Schemas\"\x0f\x82\xd3\xe4\x93\x02\t\x12\x07/schemaB-Z+github.com/opencord/voltha-protos/go/schemab\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_PROTOFILE = _descriptor.Descriptor(
  name='ProtoFile',
  full_name='schema.ProtoFile',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='file_name', full_name='schema.ProtoFile.file_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proto', full_name='schema.ProtoFile.proto', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='descriptor', full_name='schema.ProtoFile.descriptor', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=97,
  serialized_end=162,
)


_SCHEMAS = _descriptor.Descriptor(
  name='Schemas',
  full_name='schema.Schemas',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='protos', full_name='schema.Schemas.protos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='swagger_from', full_name='schema.Schemas.swagger_from', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='yang_from', full_name='schema.Schemas.yang_from', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=164,
  serialized_end=249,
)

_SCHEMAS.fields_by_name['protos'].message_type = _PROTOFILE
DESCRIPTOR.message_types_by_name['ProtoFile'] = _PROTOFILE
DESCRIPTOR.message_types_by_name['Schemas'] = _SCHEMAS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProtoFile = _reflection.GeneratedProtocolMessageType('ProtoFile', (_message.Message,), dict(
  DESCRIPTOR = _PROTOFILE,
  __module__ = 'voltha_protos.schema_pb2'
  # @@protoc_insertion_point(class_scope:schema.ProtoFile)
  ))
_sym_db.RegisterMessage(ProtoFile)

Schemas = _reflection.GeneratedProtocolMessageType('Schemas', (_message.Message,), dict(
  DESCRIPTOR = _SCHEMAS,
  __module__ = 'voltha_protos.schema_pb2'
  # @@protoc_insertion_point(class_scope:schema.Schemas)
  ))
_sym_db.RegisterMessage(Schemas)


DESCRIPTOR._options = None

_SCHEMASERVICE = _descriptor.ServiceDescriptor(
  name='SchemaService',
  full_name='schema.SchemaService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=251,
  serialized_end=337,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetSchema',
    full_name='schema.SchemaService.GetSchema',
    index=0,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_SCHEMAS,
    serialized_options=_b('\202\323\344\223\002\t\022\007/schema'),
  ),
])
_sym_db.RegisterServiceDescriptor(_SCHEMASERVICE)

DESCRIPTOR.services_by_name['SchemaService'] = _SCHEMASERVICE

# @@protoc_insertion_point(module_scope)

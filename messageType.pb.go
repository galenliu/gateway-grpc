// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: messageType.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_PluginRegisterRequest              MessageType = 0
	MessageType_PluginRegisterResponse             MessageType = 1
	MessageType_PluginUnloadRequest                MessageType = 2
	MessageType_PluginUnloadResponse               MessageType = 3
	MessageType_PluginErrorNotification            MessageType = 4
	MessageType_AdapterAddedNotification           MessageType = 4096
	MessageType_AdapterCancelPairingCommand        MessageType = 4100
	MessageType_AdapterPairingPromptNotification   MessageType = 4101
	MessageType_AdapterUnpairingPromptNotification MessageType = 4102
	MessageType_AdapterRemoveDeviceRequest         MessageType = 4103
	MessageType_AdapterRemoveDeviceResponse        MessageType = 4104
	MessageType_AdapterCancelRemoveDeviceCommand   MessageType = 4105
	MessageType_AdapterUnloadRequest               MessageType = 4097
	MessageType_AdapterUnloadResponse              MessageType = 4098
	MessageType_AdapterStartPairingCommand         MessageType = 4099
	MessageType_ApiHandlerAddedNotification        MessageType = 20480
	MessageType_ApiHandlerApiRequest               MessageType = 20483
	MessageType_ApiHandlerApiResponse              MessageType = 20484
	MessageType_ApiHandlerUnloadRequest            MessageType = 20481
	MessageType_ApiHandlerUnloadResponse           MessageType = 20482
	MessageType_DeviceActionStatusNotification     MessageType = 8201
	MessageType_DeviceAddedNotification            MessageType = 8192
	MessageType_DeviceConnectedStateNotification   MessageType = 8197
	MessageType_DeviceDebugCommand                 MessageType = 8206
	MessageType_DeviceEventNotification            MessageType = 8200
	MessageType_DevicePropertyChangedNotification  MessageType = 8199
	MessageType_DeviceRemoveActionRequest          MessageType = 8202
	MessageType_DeviceRemoveActionResponse         MessageType = 8203
	MessageType_DeviceRequestActionRequest         MessageType = 8204
	MessageType_DeviceRequestActionResponse        MessageType = 8205
	MessageType_DeviceSavedNotification            MessageType = 8207
	MessageType_DeviceSetCredentialsRequest        MessageType = 8195
	MessageType_DeviceSetCredentialsResponse       MessageType = 8196
	MessageType_DeviceSetPinRequest                MessageType = 8193
	MessageType_DeviceSetPinResponse               MessageType = 8194
	MessageType_DeviceSetPropertyCommand           MessageType = 8198
	MessageType_MockAdapterAddDeviceRequest        MessageType = 61440
	MessageType_MockAdapterAddDeviceResponse       MessageType = 61441
	MessageType_MockAdapterClearStateRequest       MessageType = 61446
	MessageType_MockAdapterClearStateResponse      MessageType = 61447
	MessageType_MockAdapterPairDeviceCommand       MessageType = 61444
	MessageType_MockAdapterRemoveDeviceRequest     MessageType = 61442
	MessageType_MockAdapterRemoveDeviceResponse    MessageType = 61443
	MessageType_MockAdapterUnpairDeviceCommand     MessageType = 61445
	MessageType_NotifierAddedNotification          MessageType = 12288
	MessageType_NotifierUnloadRequest              MessageType = 12289
	MessageType_NotifierUnloadResponse             MessageType = 12290
	MessageType_OutletAddedNotification            MessageType = 16384
	MessageType_OutletNotifyRequest                MessageType = 16386
	MessageType_OutletNotifyResponse               MessageType = 16387
	MessageType_OutletRemovedNotification          MessageType = 16385
	MessageType_ServiceAddedNotification           MessageType = 81000
	MessageType_ServiceSetPropertyValueRequest     MessageType = 81100
	MessageType_ServicePropertyChangedNotification MessageType = 81101
	MessageType_ServiceActionsStatusNotification   MessageType = 81109
	MessageType_ServiceGetThingsRequest            MessageType = 81001
	MessageType_ServiceGetThingsResponse           MessageType = 81002
	MessageType_ServiceGetThingRequest             MessageType = 81003
	MessageType_ServiceGetThingResponse            MessageType = 81004
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:     "PluginRegisterRequest",
		1:     "PluginRegisterResponse",
		2:     "PluginUnloadRequest",
		3:     "PluginUnloadResponse",
		4:     "PluginErrorNotification",
		4096:  "AdapterAddedNotification",
		4100:  "AdapterCancelPairingCommand",
		4101:  "AdapterPairingPromptNotification",
		4102:  "AdapterUnpairingPromptNotification",
		4103:  "AdapterRemoveDeviceRequest",
		4104:  "AdapterRemoveDeviceResponse",
		4105:  "AdapterCancelRemoveDeviceCommand",
		4097:  "AdapterUnloadRequest",
		4098:  "AdapterUnloadResponse",
		4099:  "AdapterStartPairingCommand",
		20480: "ApiHandlerAddedNotification",
		20483: "ApiHandlerApiRequest",
		20484: "ApiHandlerApiResponse",
		20481: "ApiHandlerUnloadRequest",
		20482: "ApiHandlerUnloadResponse",
		8201:  "DeviceActionStatusNotification",
		8192:  "DeviceAddedNotification",
		8197:  "DeviceConnectedStateNotification",
		8206:  "DeviceDebugCommand",
		8200:  "DeviceEventNotification",
		8199:  "DevicePropertyChangedNotification",
		8202:  "DeviceRemoveActionRequest",
		8203:  "DeviceRemoveActionResponse",
		8204:  "DeviceRequestActionRequest",
		8205:  "DeviceRequestActionResponse",
		8207:  "DeviceSavedNotification",
		8195:  "DeviceSetCredentialsRequest",
		8196:  "DeviceSetCredentialsResponse",
		8193:  "DeviceSetPinRequest",
		8194:  "DeviceSetPinResponse",
		8198:  "DeviceSetPropertyCommand",
		61440: "MockAdapterAddDeviceRequest",
		61441: "MockAdapterAddDeviceResponse",
		61446: "MockAdapterClearStateRequest",
		61447: "MockAdapterClearStateResponse",
		61444: "MockAdapterPairDeviceCommand",
		61442: "MockAdapterRemoveDeviceRequest",
		61443: "MockAdapterRemoveDeviceResponse",
		61445: "MockAdapterUnpairDeviceCommand",
		12288: "NotifierAddedNotification",
		12289: "NotifierUnloadRequest",
		12290: "NotifierUnloadResponse",
		16384: "OutletAddedNotification",
		16386: "OutletNotifyRequest",
		16387: "OutletNotifyResponse",
		16385: "OutletRemovedNotification",
		81000: "ServiceAddedNotification",
		81100: "ServiceSetPropertyValueRequest",
		81101: "ServicePropertyChangedNotification",
		81109: "ServiceActionsStatusNotification",
		81001: "ServiceGetThingsRequest",
		81002: "ServiceGetThingsResponse",
		81003: "ServiceGetThingRequest",
		81004: "ServiceGetThingResponse",
	}
	MessageType_value = map[string]int32{
		"PluginRegisterRequest":              0,
		"PluginRegisterResponse":             1,
		"PluginUnloadRequest":                2,
		"PluginUnloadResponse":               3,
		"PluginErrorNotification":            4,
		"AdapterAddedNotification":           4096,
		"AdapterCancelPairingCommand":        4100,
		"AdapterPairingPromptNotification":   4101,
		"AdapterUnpairingPromptNotification": 4102,
		"AdapterRemoveDeviceRequest":         4103,
		"AdapterRemoveDeviceResponse":        4104,
		"AdapterCancelRemoveDeviceCommand":   4105,
		"AdapterUnloadRequest":               4097,
		"AdapterUnloadResponse":              4098,
		"AdapterStartPairingCommand":         4099,
		"ApiHandlerAddedNotification":        20480,
		"ApiHandlerApiRequest":               20483,
		"ApiHandlerApiResponse":              20484,
		"ApiHandlerUnloadRequest":            20481,
		"ApiHandlerUnloadResponse":           20482,
		"DeviceActionStatusNotification":     8201,
		"DeviceAddedNotification":            8192,
		"DeviceConnectedStateNotification":   8197,
		"DeviceDebugCommand":                 8206,
		"DeviceEventNotification":            8200,
		"DevicePropertyChangedNotification":  8199,
		"DeviceRemoveActionRequest":          8202,
		"DeviceRemoveActionResponse":         8203,
		"DeviceRequestActionRequest":         8204,
		"DeviceRequestActionResponse":        8205,
		"DeviceSavedNotification":            8207,
		"DeviceSetCredentialsRequest":        8195,
		"DeviceSetCredentialsResponse":       8196,
		"DeviceSetPinRequest":                8193,
		"DeviceSetPinResponse":               8194,
		"DeviceSetPropertyCommand":           8198,
		"MockAdapterAddDeviceRequest":        61440,
		"MockAdapterAddDeviceResponse":       61441,
		"MockAdapterClearStateRequest":       61446,
		"MockAdapterClearStateResponse":      61447,
		"MockAdapterPairDeviceCommand":       61444,
		"MockAdapterRemoveDeviceRequest":     61442,
		"MockAdapterRemoveDeviceResponse":    61443,
		"MockAdapterUnpairDeviceCommand":     61445,
		"NotifierAddedNotification":          12288,
		"NotifierUnloadRequest":              12289,
		"NotifierUnloadResponse":             12290,
		"OutletAddedNotification":            16384,
		"OutletNotifyRequest":                16386,
		"OutletNotifyResponse":               16387,
		"OutletRemovedNotification":          16385,
		"ServiceAddedNotification":           81000,
		"ServiceSetPropertyValueRequest":     81100,
		"ServicePropertyChangedNotification": 81101,
		"ServiceActionsStatusNotification":   81109,
		"ServiceGetThingsRequest":            81001,
		"ServiceGetThingsResponse":           81002,
		"ServiceGetThingRequest":             81003,
		"ServiceGetThingResponse":            81004,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_null    Type = 0
	Type_boolean Type = 1
	Type_object  Type = 2
	Type_array   Type = 3
	Type_number  Type = 4
	Type_integer Type = 5
	Type_string  Type = 6
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "null",
		1: "boolean",
		2: "object",
		3: "array",
		4: "number",
		5: "integer",
		6: "string",
	}
	Type_value = map[string]int32{
		"null":    0,
		"boolean": 1,
		"object":  2,
		"array":   3,
		"number":  4,
		"integer": 5,
		"string":  6,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{1}
}

type Unit int32

const (
	Unit_degreeCentigrade Unit = 0
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "degreeCentigrade",
	}
	Unit_value = map[string]int32{
		"degreeCentigrade": 0,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[2].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[2]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{2}
}

type Language int32

const (
	Language_zh_cn Language = 0
	Language_zh_hk Language = 1
	Language_zh_tw Language = 2
	Language_en_us Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "zh_cn",
		1: "zh_hk",
		2: "zh_tw",
		3: "en_us",
	}
	Language_value = map[string]int32{
		"zh_cn": 0,
		"zh_hk": 1,
		"zh_tw": 2,
		"en_us": 3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_messageType_proto_enumTypes[3].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_messageType_proto_enumTypes[3]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_messageType_proto_rawDescGZIP(), []int{3}
}

var File_messageType_proto protoreflect.FileDescriptor

var file_messageType_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x91, 0x0f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x55, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x03, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x18, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x80, 0x20, 0x12, 0x20, 0x0a, 0x1b, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x84, 0x20, 0x12, 0x25, 0x0a, 0x20, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x85, 0x20, 0x12, 0x27, 0x0a, 0x22, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x55, 0x6e, 0x70,
	0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x86, 0x20, 0x12, 0x1f, 0x0a, 0x1a, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x87, 0x20, 0x12, 0x20, 0x0a, 0x1b,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x88, 0x20, 0x12, 0x25,
	0x0a, 0x20, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x10, 0x89, 0x20, 0x12, 0x19, 0x0a, 0x14, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x81, 0x20,
	0x12, 0x1a, 0x0a, 0x15, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x82, 0x20, 0x12, 0x1f, 0x0a, 0x1a,
	0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x83, 0x20, 0x12, 0x21, 0x0a,
	0x1b, 0x41, 0x70, 0x69, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x80, 0xa0, 0x01,
	0x12, 0x1a, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x83, 0xa0, 0x01, 0x12, 0x1b, 0x0a, 0x15,
	0x41, 0x70, 0x69, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x84, 0xa0, 0x01, 0x12, 0x1d, 0x0a, 0x17, 0x41, 0x70, 0x69,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x10, 0x81, 0xa0, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x41, 0x70, 0x69, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x10, 0x82, 0xa0, 0x01, 0x12, 0x23, 0x0a, 0x1e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x89, 0x40, 0x12, 0x1c, 0x0a,
	0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x80, 0x40, 0x12, 0x25, 0x0a, 0x20, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x85, 0x40, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x8e, 0x40, 0x12, 0x1c, 0x0a, 0x17, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x88, 0x40, 0x12, 0x26, 0x0a, 0x21, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x87,
	0x40, 0x12, 0x1e, 0x0a, 0x19, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x8a,
	0x40, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10,
	0x8b, 0x40, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x10, 0x8c, 0x40, 0x12, 0x20, 0x0a, 0x1b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x10, 0x8d, 0x40, 0x12, 0x1c, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x61, 0x76, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x8f, 0x40, 0x12, 0x20, 0x0a, 0x1b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x83, 0x40, 0x12, 0x21, 0x0a, 0x1c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x84, 0x40, 0x12, 0x18, 0x0a, 0x13, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10,
	0x81, 0x40, 0x12, 0x19, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x50,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x82, 0x40, 0x12, 0x1d, 0x0a,
	0x18, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x86, 0x40, 0x12, 0x21, 0x0a, 0x1b,
	0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x80, 0xe0, 0x03, 0x12,
	0x22, 0x0a, 0x1c, 0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10,
	0x81, 0xe0, 0x03, 0x12, 0x22, 0x0a, 0x1c, 0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0x86, 0xe0, 0x03, 0x12, 0x23, 0x0a, 0x1d, 0x4d, 0x6f, 0x63, 0x6b, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x87, 0xe0, 0x03, 0x12, 0x22, 0x0a, 0x1c,
	0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x84, 0xe0, 0x03,
	0x12, 0x24, 0x0a, 0x1e, 0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x82, 0xe0, 0x03, 0x12, 0x25, 0x0a, 0x1f, 0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x83, 0xe0, 0x03, 0x12, 0x24, 0x0a,
	0x1e, 0x4d, 0x6f, 0x63, 0x6b, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x55, 0x6e, 0x70, 0x61,
	0x69, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10,
	0x85, 0xe0, 0x03, 0x12, 0x1e, 0x0a, 0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x80, 0x60, 0x12, 0x1a, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x55,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x81, 0x60, 0x12,
	0x1b, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x82, 0x60, 0x12, 0x1d, 0x0a, 0x17,
	0x4f, 0x75, 0x74, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x80, 0x80, 0x01, 0x12, 0x19, 0x0a, 0x13, 0x4f,
	0x75, 0x74, 0x6c, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x82, 0x80, 0x01, 0x12, 0x1a, 0x0a, 0x14, 0x4f, 0x75, 0x74, 0x6c, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x83,
	0x80, 0x01, 0x12, 0x1f, 0x0a, 0x19, 0x4f, 0x75, 0x74, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x81, 0x80, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0xe8, 0xf8, 0x04, 0x12, 0x24, 0x0a, 0x1e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xcc, 0xf9, 0x04, 0x12, 0x28, 0x0a, 0x22, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0xcd, 0xf9, 0x04, 0x12, 0x26, 0x0a, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xd5, 0xf9, 0x04, 0x12, 0x1d, 0x0a, 0x17, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xe9, 0xf8, 0x04, 0x12, 0x1e, 0x0a, 0x18, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0xea, 0xf8, 0x04, 0x12, 0x1c, 0x0a, 0x16, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x10, 0xeb, 0xf8, 0x04, 0x12, 0x1d, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x10, 0xec, 0xf8, 0x04, 0x2a, 0x59, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x10, 0x06, 0x2a, 0x1c, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x64, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x69, 0x67, 0x72, 0x61, 0x64, 0x65, 0x10, 0x00,
	0x2a, 0x36, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x7a, 0x68, 0x5f, 0x63, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x7a, 0x68, 0x5f, 0x68, 0x6b,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x7a, 0x68, 0x5f, 0x74, 0x77, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x65, 0x6e, 0x5f, 0x75, 0x73, 0x10, 0x03, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2e, 0x2f, 0x2e,
	0x2e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messageType_proto_rawDescOnce sync.Once
	file_messageType_proto_rawDescData = file_messageType_proto_rawDesc
)

func file_messageType_proto_rawDescGZIP() []byte {
	file_messageType_proto_rawDescOnce.Do(func() {
		file_messageType_proto_rawDescData = protoimpl.X.CompressGZIP(file_messageType_proto_rawDescData)
	})
	return file_messageType_proto_rawDescData
}

var file_messageType_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_messageType_proto_goTypes = []interface{}{
	(MessageType)(0), // 0: messageType
	(Type)(0),        // 1: type
	(Unit)(0),        // 2: unit
	(Language)(0),    // 3: language
}
var file_messageType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messageType_proto_init() }
func file_messageType_proto_init() {
	if File_messageType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messageType_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messageType_proto_goTypes,
		DependencyIndexes: file_messageType_proto_depIdxs,
		EnumInfos:         file_messageType_proto_enumTypes,
	}.Build()
	File_messageType_proto = out.File
	file_messageType_proto_rawDesc = nil
	file_messageType_proto_goTypes = nil
	file_messageType_proto_depIdxs = nil
}

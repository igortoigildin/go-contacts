// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/proto/service.proto

package chatpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto_service_proto protoreflect.FileDescriptor

const file_api_proto_service_proto_rawDesc = "" +
	"\n" +
	"\x17api/proto/service.proto\x12\x04chat\x1a\x18api/proto/messages.proto2\xa5\x02\n" +
	"\vChatService\x12A\n" +
	"\n" +
	"CreateChat\x12\x17.chat.CreateChatRequest\x1a\x18.chat.CreateChatResponse\"\x00\x12G\n" +
	"\fGetUserChats\x12\x19.chat.GetUserChatsRequest\x1a\x1a.chat.GetUserChatsResponse\"\x00\x12D\n" +
	"\vSendMessage\x12\x18.chat.SendMessageRequest\x1a\x19.chat.SendMessageResponse\"\x00\x12D\n" +
	"\vGetMessages\x12\x18.chat.GetMessagesRequest\x1a\x19.chat.GetMessagesResponse\"\x00B<Z:github.com/igortoigildin/go-contacts/chat/api/proto;chatpbb\x06proto3"

var file_api_proto_service_proto_goTypes = []any{
	(*CreateChatRequest)(nil),    // 0: chat.CreateChatRequest
	(*GetUserChatsRequest)(nil),  // 1: chat.GetUserChatsRequest
	(*SendMessageRequest)(nil),   // 2: chat.SendMessageRequest
	(*GetMessagesRequest)(nil),   // 3: chat.GetMessagesRequest
	(*CreateChatResponse)(nil),   // 4: chat.CreateChatResponse
	(*GetUserChatsResponse)(nil), // 5: chat.GetUserChatsResponse
	(*SendMessageResponse)(nil),  // 6: chat.SendMessageResponse
	(*GetMessagesResponse)(nil),  // 7: chat.GetMessagesResponse
}
var file_api_proto_service_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.CreateChat:input_type -> chat.CreateChatRequest
	1, // 1: chat.ChatService.GetUserChats:input_type -> chat.GetUserChatsRequest
	2, // 2: chat.ChatService.SendMessage:input_type -> chat.SendMessageRequest
	3, // 3: chat.ChatService.GetMessages:input_type -> chat.GetMessagesRequest
	4, // 4: chat.ChatService.CreateChat:output_type -> chat.CreateChatResponse
	5, // 5: chat.ChatService.GetUserChats:output_type -> chat.GetUserChatsResponse
	6, // 6: chat.ChatService.SendMessage:output_type -> chat.SendMessageResponse
	7, // 7: chat.ChatService.GetMessages:output_type -> chat.GetMessagesResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_service_proto_init() }
func file_api_proto_service_proto_init() {
	if File_api_proto_service_proto != nil {
		return
	}
	file_api_proto_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_service_proto_rawDesc), len(file_api_proto_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_service_proto_goTypes,
		DependencyIndexes: file_api_proto_service_proto_depIdxs,
	}.Build()
	File_api_proto_service_proto = out.File
	file_api_proto_service_proto_goTypes = nil
	file_api_proto_service_proto_depIdxs = nil
}

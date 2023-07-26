// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "github.com/kubernetes-csi/csi-proxy/client/api/smb/v1beta2"
	impl "github.com/kubernetes-csi/csi-proxy/pkg/server/smb/impl"
)

func autoConvert_v1beta2_NewSmbGlobalMappingRequest_To_impl_NewSmbGlobalMappingRequest(in *v1beta2.NewSmbGlobalMappingRequest, out *impl.NewSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	out.LocalPath = in.LocalPath
	out.Username = in.Username
	out.Password = in.Password
	return nil
}

// Convert_v1beta2_NewSmbGlobalMappingRequest_To_impl_NewSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_v1beta2_NewSmbGlobalMappingRequest_To_impl_NewSmbGlobalMappingRequest(in *v1beta2.NewSmbGlobalMappingRequest, out *impl.NewSmbGlobalMappingRequest) error {
	return autoConvert_v1beta2_NewSmbGlobalMappingRequest_To_impl_NewSmbGlobalMappingRequest(in, out)
}

func autoConvert_impl_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in *impl.NewSmbGlobalMappingRequest, out *v1beta2.NewSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	out.LocalPath = in.LocalPath
	out.Username = in.Username
	out.Password = in.Password
	return nil
}

// Convert_impl_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_impl_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in *impl.NewSmbGlobalMappingRequest, out *v1beta2.NewSmbGlobalMappingRequest) error {
	return autoConvert_impl_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in, out)
}

func autoConvert_v1beta2_NewSmbGlobalMappingResponse_To_impl_NewSmbGlobalMappingResponse(in *v1beta2.NewSmbGlobalMappingResponse, out *impl.NewSmbGlobalMappingResponse) error {
	return nil
}

// Convert_v1beta2_NewSmbGlobalMappingResponse_To_impl_NewSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_v1beta2_NewSmbGlobalMappingResponse_To_impl_NewSmbGlobalMappingResponse(in *v1beta2.NewSmbGlobalMappingResponse, out *impl.NewSmbGlobalMappingResponse) error {
	return autoConvert_v1beta2_NewSmbGlobalMappingResponse_To_impl_NewSmbGlobalMappingResponse(in, out)
}

func autoConvert_impl_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in *impl.NewSmbGlobalMappingResponse, out *v1beta2.NewSmbGlobalMappingResponse) error {
	return nil
}

// Convert_impl_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_impl_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in *impl.NewSmbGlobalMappingResponse, out *v1beta2.NewSmbGlobalMappingResponse) error {
	return autoConvert_impl_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in, out)
}

func autoConvert_v1beta2_RemoveSmbGlobalMappingRequest_To_impl_RemoveSmbGlobalMappingRequest(in *v1beta2.RemoveSmbGlobalMappingRequest, out *impl.RemoveSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	return nil
}

// Convert_v1beta2_RemoveSmbGlobalMappingRequest_To_impl_RemoveSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_v1beta2_RemoveSmbGlobalMappingRequest_To_impl_RemoveSmbGlobalMappingRequest(in *v1beta2.RemoveSmbGlobalMappingRequest, out *impl.RemoveSmbGlobalMappingRequest) error {
	return autoConvert_v1beta2_RemoveSmbGlobalMappingRequest_To_impl_RemoveSmbGlobalMappingRequest(in, out)
}

func autoConvert_impl_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in *impl.RemoveSmbGlobalMappingRequest, out *v1beta2.RemoveSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	return nil
}

// Convert_impl_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_impl_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in *impl.RemoveSmbGlobalMappingRequest, out *v1beta2.RemoveSmbGlobalMappingRequest) error {
	return autoConvert_impl_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in, out)
}

func autoConvert_v1beta2_RemoveSmbGlobalMappingResponse_To_impl_RemoveSmbGlobalMappingResponse(in *v1beta2.RemoveSmbGlobalMappingResponse, out *impl.RemoveSmbGlobalMappingResponse) error {
	return nil
}

// Convert_v1beta2_RemoveSmbGlobalMappingResponse_To_impl_RemoveSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_v1beta2_RemoveSmbGlobalMappingResponse_To_impl_RemoveSmbGlobalMappingResponse(in *v1beta2.RemoveSmbGlobalMappingResponse, out *impl.RemoveSmbGlobalMappingResponse) error {
	return autoConvert_v1beta2_RemoveSmbGlobalMappingResponse_To_impl_RemoveSmbGlobalMappingResponse(in, out)
}

func autoConvert_impl_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in *impl.RemoveSmbGlobalMappingResponse, out *v1beta2.RemoveSmbGlobalMappingResponse) error {
	return nil
}

// Convert_impl_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_impl_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in *impl.RemoveSmbGlobalMappingResponse, out *v1beta2.RemoveSmbGlobalMappingResponse) error {
	return autoConvert_impl_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in, out)
}
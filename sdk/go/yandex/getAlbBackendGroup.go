// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlbBackendGroup(ctx *pulumi.Context, args *LookupAlbBackendGroupArgs, opts ...pulumi.InvokeOption) (*LookupAlbBackendGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAlbBackendGroupResult
	err := ctx.Invoke("yandex:index/getAlbBackendGroup:getAlbBackendGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlbBackendGroup.
type LookupAlbBackendGroupArgs struct {
	BackendGroupId  *string                            `pulumi:"backendGroupId"`
	Description     *string                            `pulumi:"description"`
	FolderId        *string                            `pulumi:"folderId"`
	GrpcBackends    []GetAlbBackendGroupGrpcBackend    `pulumi:"grpcBackends"`
	HttpBackends    []GetAlbBackendGroupHttpBackend    `pulumi:"httpBackends"`
	Labels          map[string]string                  `pulumi:"labels"`
	Name            *string                            `pulumi:"name"`
	SessionAffinity *GetAlbBackendGroupSessionAffinity `pulumi:"sessionAffinity"`
	StreamBackends  []GetAlbBackendGroupStreamBackend  `pulumi:"streamBackends"`
}

// A collection of values returned by getAlbBackendGroup.
type LookupAlbBackendGroupResult struct {
	BackendGroupId string                          `pulumi:"backendGroupId"`
	CreatedAt      string                          `pulumi:"createdAt"`
	Description    string                          `pulumi:"description"`
	FolderId       string                          `pulumi:"folderId"`
	GrpcBackends   []GetAlbBackendGroupGrpcBackend `pulumi:"grpcBackends"`
	HttpBackends   []GetAlbBackendGroupHttpBackend `pulumi:"httpBackends"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                            `pulumi:"id"`
	Labels          map[string]string                 `pulumi:"labels"`
	Name            string                            `pulumi:"name"`
	SessionAffinity GetAlbBackendGroupSessionAffinity `pulumi:"sessionAffinity"`
	StreamBackends  []GetAlbBackendGroupStreamBackend `pulumi:"streamBackends"`
}

func LookupAlbBackendGroupOutput(ctx *pulumi.Context, args LookupAlbBackendGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAlbBackendGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAlbBackendGroupResultOutput, error) {
			args := v.(LookupAlbBackendGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getAlbBackendGroup:getAlbBackendGroup", args, LookupAlbBackendGroupResultOutput{}, options).(LookupAlbBackendGroupResultOutput), nil
		}).(LookupAlbBackendGroupResultOutput)
}

// A collection of arguments for invoking getAlbBackendGroup.
type LookupAlbBackendGroupOutputArgs struct {
	BackendGroupId  pulumi.StringPtrInput                     `pulumi:"backendGroupId"`
	Description     pulumi.StringPtrInput                     `pulumi:"description"`
	FolderId        pulumi.StringPtrInput                     `pulumi:"folderId"`
	GrpcBackends    GetAlbBackendGroupGrpcBackendArrayInput   `pulumi:"grpcBackends"`
	HttpBackends    GetAlbBackendGroupHttpBackendArrayInput   `pulumi:"httpBackends"`
	Labels          pulumi.StringMapInput                     `pulumi:"labels"`
	Name            pulumi.StringPtrInput                     `pulumi:"name"`
	SessionAffinity GetAlbBackendGroupSessionAffinityPtrInput `pulumi:"sessionAffinity"`
	StreamBackends  GetAlbBackendGroupStreamBackendArrayInput `pulumi:"streamBackends"`
}

func (LookupAlbBackendGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbBackendGroupArgs)(nil)).Elem()
}

// A collection of values returned by getAlbBackendGroup.
type LookupAlbBackendGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAlbBackendGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbBackendGroupResult)(nil)).Elem()
}

func (o LookupAlbBackendGroupResultOutput) ToLookupAlbBackendGroupResultOutput() LookupAlbBackendGroupResultOutput {
	return o
}

func (o LookupAlbBackendGroupResultOutput) ToLookupAlbBackendGroupResultOutputWithContext(ctx context.Context) LookupAlbBackendGroupResultOutput {
	return o
}

func (o LookupAlbBackendGroupResultOutput) BackendGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.BackendGroupId }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) GrpcBackends() GetAlbBackendGroupGrpcBackendArrayOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) []GetAlbBackendGroupGrpcBackend { return v.GrpcBackends }).(GetAlbBackendGroupGrpcBackendArrayOutput)
}

func (o LookupAlbBackendGroupResultOutput) HttpBackends() GetAlbBackendGroupHttpBackendArrayOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) []GetAlbBackendGroupHttpBackend { return v.HttpBackends }).(GetAlbBackendGroupHttpBackendArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlbBackendGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupAlbBackendGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlbBackendGroupResultOutput) SessionAffinity() GetAlbBackendGroupSessionAffinityOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) GetAlbBackendGroupSessionAffinity { return v.SessionAffinity }).(GetAlbBackendGroupSessionAffinityOutput)
}

func (o LookupAlbBackendGroupResultOutput) StreamBackends() GetAlbBackendGroupStreamBackendArrayOutput {
	return o.ApplyT(func(v LookupAlbBackendGroupResult) []GetAlbBackendGroupStreamBackend { return v.StreamBackends }).(GetAlbBackendGroupStreamBackendArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlbBackendGroupResultOutput{})
}

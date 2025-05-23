// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpcGateway(ctx *pulumi.Context, args *LookupVpcGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpcGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcGatewayResult
	err := ctx.Invoke("yandex:index/getVpcGateway:getVpcGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcGateway.
type LookupVpcGatewayArgs struct {
	FolderId            *string                           `pulumi:"folderId"`
	GatewayId           *string                           `pulumi:"gatewayId"`
	Name                *string                           `pulumi:"name"`
	SharedEgressGateway *GetVpcGatewaySharedEgressGateway `pulumi:"sharedEgressGateway"`
}

// A collection of values returned by getVpcGateway.
type LookupVpcGatewayResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	GatewayId   string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                            `pulumi:"id"`
	Labels              map[string]string                 `pulumi:"labels"`
	Name                string                            `pulumi:"name"`
	SharedEgressGateway *GetVpcGatewaySharedEgressGateway `pulumi:"sharedEgressGateway"`
}

func LookupVpcGatewayOutput(ctx *pulumi.Context, args LookupVpcGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpcGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcGatewayResultOutput, error) {
			args := v.(LookupVpcGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getVpcGateway:getVpcGateway", args, LookupVpcGatewayResultOutput{}, options).(LookupVpcGatewayResultOutput), nil
		}).(LookupVpcGatewayResultOutput)
}

// A collection of arguments for invoking getVpcGateway.
type LookupVpcGatewayOutputArgs struct {
	FolderId            pulumi.StringPtrInput                    `pulumi:"folderId"`
	GatewayId           pulumi.StringPtrInput                    `pulumi:"gatewayId"`
	Name                pulumi.StringPtrInput                    `pulumi:"name"`
	SharedEgressGateway GetVpcGatewaySharedEgressGatewayPtrInput `pulumi:"sharedEgressGateway"`
}

func (LookupVpcGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getVpcGateway.
type LookupVpcGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpcGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcGatewayResult)(nil)).Elem()
}

func (o LookupVpcGatewayResultOutput) ToLookupVpcGatewayResultOutput() LookupVpcGatewayResultOutput {
	return o
}

func (o LookupVpcGatewayResultOutput) ToLookupVpcGatewayResultOutputWithContext(ctx context.Context) LookupVpcGatewayResultOutput {
	return o
}

func (o LookupVpcGatewayResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupVpcGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayResultOutput) SharedEgressGateway() GetVpcGatewaySharedEgressGatewayPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayResult) *GetVpcGatewaySharedEgressGateway { return v.SharedEgressGateway }).(GetVpcGatewaySharedEgressGatewayPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcGatewayResultOutput{})
}

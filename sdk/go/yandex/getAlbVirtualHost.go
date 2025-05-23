// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlbVirtualHost(ctx *pulumi.Context, args *LookupAlbVirtualHostArgs, opts ...pulumi.InvokeOption) (*LookupAlbVirtualHostResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAlbVirtualHostResult
	err := ctx.Invoke("yandex:index/getAlbVirtualHost:getAlbVirtualHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlbVirtualHost.
type LookupAlbVirtualHostArgs struct {
	HttpRouterId  *string `pulumi:"httpRouterId"`
	Name          *string `pulumi:"name"`
	VirtualHostId *string `pulumi:"virtualHostId"`
}

// A collection of values returned by getAlbVirtualHost.
type LookupAlbVirtualHostResult struct {
	Authorities  []string `pulumi:"authorities"`
	HttpRouterId string   `pulumi:"httpRouterId"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                                  `pulumi:"id"`
	ModifyRequestHeaders  []GetAlbVirtualHostModifyRequestHeader  `pulumi:"modifyRequestHeaders"`
	ModifyResponseHeaders []GetAlbVirtualHostModifyResponseHeader `pulumi:"modifyResponseHeaders"`
	Name                  string                                  `pulumi:"name"`
	RateLimits            []GetAlbVirtualHostRateLimit            `pulumi:"rateLimits"`
	RouteOptions          []GetAlbVirtualHostRouteOption          `pulumi:"routeOptions"`
	Routes                []GetAlbVirtualHostRoute                `pulumi:"routes"`
	VirtualHostId         string                                  `pulumi:"virtualHostId"`
}

func LookupAlbVirtualHostOutput(ctx *pulumi.Context, args LookupAlbVirtualHostOutputArgs, opts ...pulumi.InvokeOption) LookupAlbVirtualHostResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAlbVirtualHostResultOutput, error) {
			args := v.(LookupAlbVirtualHostArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getAlbVirtualHost:getAlbVirtualHost", args, LookupAlbVirtualHostResultOutput{}, options).(LookupAlbVirtualHostResultOutput), nil
		}).(LookupAlbVirtualHostResultOutput)
}

// A collection of arguments for invoking getAlbVirtualHost.
type LookupAlbVirtualHostOutputArgs struct {
	HttpRouterId  pulumi.StringPtrInput `pulumi:"httpRouterId"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	VirtualHostId pulumi.StringPtrInput `pulumi:"virtualHostId"`
}

func (LookupAlbVirtualHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbVirtualHostArgs)(nil)).Elem()
}

// A collection of values returned by getAlbVirtualHost.
type LookupAlbVirtualHostResultOutput struct{ *pulumi.OutputState }

func (LookupAlbVirtualHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlbVirtualHostResult)(nil)).Elem()
}

func (o LookupAlbVirtualHostResultOutput) ToLookupAlbVirtualHostResultOutput() LookupAlbVirtualHostResultOutput {
	return o
}

func (o LookupAlbVirtualHostResultOutput) ToLookupAlbVirtualHostResultOutputWithContext(ctx context.Context) LookupAlbVirtualHostResultOutput {
	return o
}

func (o LookupAlbVirtualHostResultOutput) Authorities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []string { return v.Authorities }).(pulumi.StringArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) HttpRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) string { return v.HttpRouterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlbVirtualHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlbVirtualHostResultOutput) ModifyRequestHeaders() GetAlbVirtualHostModifyRequestHeaderArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []GetAlbVirtualHostModifyRequestHeader {
		return v.ModifyRequestHeaders
	}).(GetAlbVirtualHostModifyRequestHeaderArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) ModifyResponseHeaders() GetAlbVirtualHostModifyResponseHeaderArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []GetAlbVirtualHostModifyResponseHeader {
		return v.ModifyResponseHeaders
	}).(GetAlbVirtualHostModifyResponseHeaderArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlbVirtualHostResultOutput) RateLimits() GetAlbVirtualHostRateLimitArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []GetAlbVirtualHostRateLimit { return v.RateLimits }).(GetAlbVirtualHostRateLimitArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) RouteOptions() GetAlbVirtualHostRouteOptionArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []GetAlbVirtualHostRouteOption { return v.RouteOptions }).(GetAlbVirtualHostRouteOptionArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) Routes() GetAlbVirtualHostRouteArrayOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) []GetAlbVirtualHostRoute { return v.Routes }).(GetAlbVirtualHostRouteArrayOutput)
}

func (o LookupAlbVirtualHostResultOutput) VirtualHostId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlbVirtualHostResult) string { return v.VirtualHostId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlbVirtualHostResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiGateway(ctx *pulumi.Context, args *LookupApiGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApiGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiGatewayResult
	err := ctx.Invoke("yandex:index/getApiGateway:getApiGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiGateway.
type LookupApiGatewayArgs struct {
	ApiGatewayId     *string                     `pulumi:"apiGatewayId"`
	Canary           *GetApiGatewayCanary        `pulumi:"canary"`
	Connectivity     *GetApiGatewayConnectivity  `pulumi:"connectivity"`
	CustomDomains    []GetApiGatewayCustomDomain `pulumi:"customDomains"`
	ExecutionTimeout *string                     `pulumi:"executionTimeout"`
	FolderId         *string                     `pulumi:"folderId"`
	Name             *string                     `pulumi:"name"`
	Variables        map[string]string           `pulumi:"variables"`
}

// A collection of values returned by getApiGateway.
type LookupApiGatewayResult struct {
	ApiGatewayId     *string                     `pulumi:"apiGatewayId"`
	Canary           *GetApiGatewayCanary        `pulumi:"canary"`
	Connectivity     *GetApiGatewayConnectivity  `pulumi:"connectivity"`
	CreatedAt        string                      `pulumi:"createdAt"`
	CustomDomains    []GetApiGatewayCustomDomain `pulumi:"customDomains"`
	Description      string                      `pulumi:"description"`
	Domain           string                      `pulumi:"domain"`
	ExecutionTimeout string                      `pulumi:"executionTimeout"`
	FolderId         *string                     `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                   `pulumi:"id"`
	Labels     map[string]string        `pulumi:"labels"`
	LogGroupId string                   `pulumi:"logGroupId"`
	LogOptions []GetApiGatewayLogOption `pulumi:"logOptions"`
	Name       *string                  `pulumi:"name"`
	Status     string                   `pulumi:"status"`
	// Deprecated: The 'user_domains' field has been deprecated. Please use 'custom_domains' instead.
	UserDomains []string          `pulumi:"userDomains"`
	Variables   map[string]string `pulumi:"variables"`
}

func LookupApiGatewayOutput(ctx *pulumi.Context, args LookupApiGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupApiGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApiGatewayResultOutput, error) {
			args := v.(LookupApiGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getApiGateway:getApiGateway", args, LookupApiGatewayResultOutput{}, options).(LookupApiGatewayResultOutput), nil
		}).(LookupApiGatewayResultOutput)
}

// A collection of arguments for invoking getApiGateway.
type LookupApiGatewayOutputArgs struct {
	ApiGatewayId     pulumi.StringPtrInput               `pulumi:"apiGatewayId"`
	Canary           GetApiGatewayCanaryPtrInput         `pulumi:"canary"`
	Connectivity     GetApiGatewayConnectivityPtrInput   `pulumi:"connectivity"`
	CustomDomains    GetApiGatewayCustomDomainArrayInput `pulumi:"customDomains"`
	ExecutionTimeout pulumi.StringPtrInput               `pulumi:"executionTimeout"`
	FolderId         pulumi.StringPtrInput               `pulumi:"folderId"`
	Name             pulumi.StringPtrInput               `pulumi:"name"`
	Variables        pulumi.StringMapInput               `pulumi:"variables"`
}

func (LookupApiGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getApiGateway.
type LookupApiGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupApiGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayResult)(nil)).Elem()
}

func (o LookupApiGatewayResultOutput) ToLookupApiGatewayResultOutput() LookupApiGatewayResultOutput {
	return o
}

func (o LookupApiGatewayResultOutput) ToLookupApiGatewayResultOutputWithContext(ctx context.Context) LookupApiGatewayResultOutput {
	return o
}

func (o LookupApiGatewayResultOutput) ApiGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *string { return v.ApiGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupApiGatewayResultOutput) Canary() GetApiGatewayCanaryPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *GetApiGatewayCanary { return v.Canary }).(GetApiGatewayCanaryPtrOutput)
}

func (o LookupApiGatewayResultOutput) Connectivity() GetApiGatewayConnectivityPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *GetApiGatewayConnectivity { return v.Connectivity }).(GetApiGatewayConnectivityPtrOutput)
}

func (o LookupApiGatewayResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) CustomDomains() GetApiGatewayCustomDomainArrayOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) []GetApiGatewayCustomDomain { return v.CustomDomains }).(GetApiGatewayCustomDomainArrayOutput)
}

func (o LookupApiGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) ExecutionTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.ExecutionTimeout }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) FolderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *string { return v.FolderId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApiGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupApiGatewayResultOutput) LogGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.LogGroupId }).(pulumi.StringOutput)
}

func (o LookupApiGatewayResultOutput) LogOptions() GetApiGatewayLogOptionArrayOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) []GetApiGatewayLogOption { return v.LogOptions }).(GetApiGatewayLogOptionArrayOutput)
}

func (o LookupApiGatewayResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApiGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

// Deprecated: The 'user_domains' field has been deprecated. Please use 'custom_domains' instead.
func (o LookupApiGatewayResultOutput) UserDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) []string { return v.UserDomains }).(pulumi.StringArrayOutput)
}

func (o LookupApiGatewayResultOutput) Variables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiGatewayResult) map[string]string { return v.Variables }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiGatewayResultOutput{})
}

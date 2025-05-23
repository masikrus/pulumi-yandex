// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSwsAdvancedRateLimiterProfile(ctx *pulumi.Context, args *LookupSwsAdvancedRateLimiterProfileArgs, opts ...pulumi.InvokeOption) (*LookupSwsAdvancedRateLimiterProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSwsAdvancedRateLimiterProfileResult
	err := ctx.Invoke("yandex:index/getSwsAdvancedRateLimiterProfile:getSwsAdvancedRateLimiterProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSwsAdvancedRateLimiterProfile.
type LookupSwsAdvancedRateLimiterProfileArgs struct {
	AdvancedRateLimiterProfileId *string `pulumi:"advancedRateLimiterProfileId"`
	CloudId                      *string `pulumi:"cloudId"`
	FolderId                     *string `pulumi:"folderId"`
	Name                         *string `pulumi:"name"`
}

// A collection of values returned by getSwsAdvancedRateLimiterProfile.
type LookupSwsAdvancedRateLimiterProfileResult struct {
	AdvancedRateLimiterProfileId *string                                                   `pulumi:"advancedRateLimiterProfileId"`
	AdvancedRateLimiterRules     []GetSwsAdvancedRateLimiterProfileAdvancedRateLimiterRule `pulumi:"advancedRateLimiterRules"`
	CloudId                      string                                                    `pulumi:"cloudId"`
	CreatedAt                    string                                                    `pulumi:"createdAt"`
	Description                  string                                                    `pulumi:"description"`
	FolderId                     string                                                    `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string            `pulumi:"id"`
	Labels map[string]string `pulumi:"labels"`
	Name   string            `pulumi:"name"`
}

func LookupSwsAdvancedRateLimiterProfileOutput(ctx *pulumi.Context, args LookupSwsAdvancedRateLimiterProfileOutputArgs, opts ...pulumi.InvokeOption) LookupSwsAdvancedRateLimiterProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSwsAdvancedRateLimiterProfileResultOutput, error) {
			args := v.(LookupSwsAdvancedRateLimiterProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getSwsAdvancedRateLimiterProfile:getSwsAdvancedRateLimiterProfile", args, LookupSwsAdvancedRateLimiterProfileResultOutput{}, options).(LookupSwsAdvancedRateLimiterProfileResultOutput), nil
		}).(LookupSwsAdvancedRateLimiterProfileResultOutput)
}

// A collection of arguments for invoking getSwsAdvancedRateLimiterProfile.
type LookupSwsAdvancedRateLimiterProfileOutputArgs struct {
	AdvancedRateLimiterProfileId pulumi.StringPtrInput `pulumi:"advancedRateLimiterProfileId"`
	CloudId                      pulumi.StringPtrInput `pulumi:"cloudId"`
	FolderId                     pulumi.StringPtrInput `pulumi:"folderId"`
	Name                         pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSwsAdvancedRateLimiterProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSwsAdvancedRateLimiterProfileArgs)(nil)).Elem()
}

// A collection of values returned by getSwsAdvancedRateLimiterProfile.
type LookupSwsAdvancedRateLimiterProfileResultOutput struct{ *pulumi.OutputState }

func (LookupSwsAdvancedRateLimiterProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSwsAdvancedRateLimiterProfileResult)(nil)).Elem()
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) ToLookupSwsAdvancedRateLimiterProfileResultOutput() LookupSwsAdvancedRateLimiterProfileResultOutput {
	return o
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) ToLookupSwsAdvancedRateLimiterProfileResultOutputWithContext(ctx context.Context) LookupSwsAdvancedRateLimiterProfileResultOutput {
	return o
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) AdvancedRateLimiterProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) *string { return v.AdvancedRateLimiterProfileId }).(pulumi.StringPtrOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) AdvancedRateLimiterRules() GetSwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleArrayOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) []GetSwsAdvancedRateLimiterProfileAdvancedRateLimiterRule {
		return v.AdvancedRateLimiterRules
	}).(GetSwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleArrayOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.CloudId }).(pulumi.StringOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSwsAdvancedRateLimiterProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupSwsAdvancedRateLimiterProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSwsAdvancedRateLimiterProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSwsAdvancedRateLimiterProfileResultOutput{})
}

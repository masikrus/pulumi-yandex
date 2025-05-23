// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIamServiceAccount(ctx *pulumi.Context, args *LookupIamServiceAccountArgs, opts ...pulumi.InvokeOption) (*LookupIamServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIamServiceAccountResult
	err := ctx.Invoke("yandex:index/getIamServiceAccount:getIamServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamServiceAccount.
type LookupIamServiceAccountArgs struct {
	FolderId         *string `pulumi:"folderId"`
	Name             *string `pulumi:"name"`
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

// A collection of values returned by getIamServiceAccount.
type LookupIamServiceAccountResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	Name             string `pulumi:"name"`
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

func LookupIamServiceAccountOutput(ctx *pulumi.Context, args LookupIamServiceAccountOutputArgs, opts ...pulumi.InvokeOption) LookupIamServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIamServiceAccountResultOutput, error) {
			args := v.(LookupIamServiceAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getIamServiceAccount:getIamServiceAccount", args, LookupIamServiceAccountResultOutput{}, options).(LookupIamServiceAccountResultOutput), nil
		}).(LookupIamServiceAccountResultOutput)
}

// A collection of arguments for invoking getIamServiceAccount.
type LookupIamServiceAccountOutputArgs struct {
	FolderId         pulumi.StringPtrInput `pulumi:"folderId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ServiceAccountId pulumi.StringPtrInput `pulumi:"serviceAccountId"`
}

func (LookupIamServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getIamServiceAccount.
type LookupIamServiceAccountResultOutput struct{ *pulumi.OutputState }

func (LookupIamServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamServiceAccountResult)(nil)).Elem()
}

func (o LookupIamServiceAccountResultOutput) ToLookupIamServiceAccountResultOutput() LookupIamServiceAccountResultOutput {
	return o
}

func (o LookupIamServiceAccountResultOutput) ToLookupIamServiceAccountResultOutputWithContext(ctx context.Context) LookupIamServiceAccountResultOutput {
	return o
}

func (o LookupIamServiceAccountResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupIamServiceAccountResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupIamServiceAccountResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIamServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIamServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIamServiceAccountResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamServiceAccountResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamServiceAccountResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoggingGroup(ctx *pulumi.Context, args *LookupLoggingGroupArgs, opts ...pulumi.InvokeOption) (*LookupLoggingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoggingGroupResult
	err := ctx.Invoke("yandex:index/getLoggingGroup:getLoggingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoggingGroup.
type LookupLoggingGroupArgs struct {
	FolderId *string `pulumi:"folderId"`
	GroupId  *string `pulumi:"groupId"`
	Name     *string `pulumi:"name"`
}

// A collection of values returned by getLoggingGroup.
type LookupLoggingGroupResult struct {
	CloudId     string `pulumi:"cloudId"`
	CreatedAt   string `pulumi:"createdAt"`
	DataStream  string `pulumi:"dataStream"`
	Description string `pulumi:"description"`
	FolderId    string `pulumi:"folderId"`
	GroupId     string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string            `pulumi:"id"`
	Labels          map[string]string `pulumi:"labels"`
	Name            string            `pulumi:"name"`
	RetentionPeriod string            `pulumi:"retentionPeriod"`
	Status          string            `pulumi:"status"`
}

func LookupLoggingGroupOutput(ctx *pulumi.Context, args LookupLoggingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupLoggingGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoggingGroupResultOutput, error) {
			args := v.(LookupLoggingGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getLoggingGroup:getLoggingGroup", args, LookupLoggingGroupResultOutput{}, options).(LookupLoggingGroupResultOutput), nil
		}).(LookupLoggingGroupResultOutput)
}

// A collection of arguments for invoking getLoggingGroup.
type LookupLoggingGroupOutputArgs struct {
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	GroupId  pulumi.StringPtrInput `pulumi:"groupId"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupLoggingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingGroupArgs)(nil)).Elem()
}

// A collection of values returned by getLoggingGroup.
type LookupLoggingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupLoggingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggingGroupResult)(nil)).Elem()
}

func (o LookupLoggingGroupResultOutput) ToLookupLoggingGroupResultOutput() LookupLoggingGroupResultOutput {
	return o
}

func (o LookupLoggingGroupResultOutput) ToLookupLoggingGroupResultOutputWithContext(ctx context.Context) LookupLoggingGroupResultOutput {
	return o
}

func (o LookupLoggingGroupResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.CloudId }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) DataStream() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.DataStream }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLoggingGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupLoggingGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) RetentionPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.RetentionPeriod }).(pulumi.StringOutput)
}

func (o LookupLoggingGroupResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoggingGroupResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggingGroupResultOutput{})
}

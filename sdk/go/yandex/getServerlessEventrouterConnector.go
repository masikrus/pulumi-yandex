// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerlessEventrouterConnector(ctx *pulumi.Context, args *LookupServerlessEventrouterConnectorArgs, opts ...pulumi.InvokeOption) (*LookupServerlessEventrouterConnectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerlessEventrouterConnectorResult
	err := ctx.Invoke("yandex:index/getServerlessEventrouterConnector:getServerlessEventrouterConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerlessEventrouterConnector.
type LookupServerlessEventrouterConnectorArgs struct {
	ConnectorId *string `pulumi:"connectorId"`
	Name        *string `pulumi:"name"`
}

// A collection of values returned by getServerlessEventrouterConnector.
type LookupServerlessEventrouterConnectorResult struct {
	BusId              string  `pulumi:"busId"`
	CloudId            string  `pulumi:"cloudId"`
	ConnectorId        *string `pulumi:"connectorId"`
	CreatedAt          string  `pulumi:"createdAt"`
	DeletionProtection bool    `pulumi:"deletionProtection"`
	Description        string  `pulumi:"description"`
	FolderId           string  `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string                                 `pulumi:"id"`
	Labels map[string]string                      `pulumi:"labels"`
	Name   *string                                `pulumi:"name"`
	Yds    []GetServerlessEventrouterConnectorYd  `pulumi:"yds"`
	Ymqs   []GetServerlessEventrouterConnectorYmq `pulumi:"ymqs"`
}

func LookupServerlessEventrouterConnectorOutput(ctx *pulumi.Context, args LookupServerlessEventrouterConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupServerlessEventrouterConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServerlessEventrouterConnectorResultOutput, error) {
			args := v.(LookupServerlessEventrouterConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getServerlessEventrouterConnector:getServerlessEventrouterConnector", args, LookupServerlessEventrouterConnectorResultOutput{}, options).(LookupServerlessEventrouterConnectorResultOutput), nil
		}).(LookupServerlessEventrouterConnectorResultOutput)
}

// A collection of arguments for invoking getServerlessEventrouterConnector.
type LookupServerlessEventrouterConnectorOutputArgs struct {
	ConnectorId pulumi.StringPtrInput `pulumi:"connectorId"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupServerlessEventrouterConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessEventrouterConnectorArgs)(nil)).Elem()
}

// A collection of values returned by getServerlessEventrouterConnector.
type LookupServerlessEventrouterConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupServerlessEventrouterConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessEventrouterConnectorResult)(nil)).Elem()
}

func (o LookupServerlessEventrouterConnectorResultOutput) ToLookupServerlessEventrouterConnectorResultOutput() LookupServerlessEventrouterConnectorResultOutput {
	return o
}

func (o LookupServerlessEventrouterConnectorResultOutput) ToLookupServerlessEventrouterConnectorResultOutputWithContext(ctx context.Context) LookupServerlessEventrouterConnectorResultOutput {
	return o
}

func (o LookupServerlessEventrouterConnectorResultOutput) BusId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.BusId }).(pulumi.StringOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.CloudId }).(pulumi.StringOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) ConnectorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) *string { return v.ConnectorId }).(pulumi.StringPtrOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServerlessEventrouterConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) Yds() GetServerlessEventrouterConnectorYdArrayOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) []GetServerlessEventrouterConnectorYd { return v.Yds }).(GetServerlessEventrouterConnectorYdArrayOutput)
}

func (o LookupServerlessEventrouterConnectorResultOutput) Ymqs() GetServerlessEventrouterConnectorYmqArrayOutput {
	return o.ApplyT(func(v LookupServerlessEventrouterConnectorResult) []GetServerlessEventrouterConnectorYmq {
		return v.Ymqs
	}).(GetServerlessEventrouterConnectorYmqArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerlessEventrouterConnectorResultOutput{})
}

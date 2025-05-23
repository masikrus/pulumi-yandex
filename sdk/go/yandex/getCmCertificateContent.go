// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCmCertificateContent(ctx *pulumi.Context, args *GetCmCertificateContentArgs, opts ...pulumi.InvokeOption) (*GetCmCertificateContentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCmCertificateContentResult
	err := ctx.Invoke("yandex:index/getCmCertificateContent:getCmCertificateContent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCmCertificateContent.
type GetCmCertificateContentArgs struct {
	CertificateId    *string `pulumi:"certificateId"`
	FolderId         *string `pulumi:"folderId"`
	Name             *string `pulumi:"name"`
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	WaitValidation   *bool   `pulumi:"waitValidation"`
}

// A collection of values returned by getCmCertificateContent.
type GetCmCertificateContentResult struct {
	CertificateId *string  `pulumi:"certificateId"`
	Certificates  []string `pulumi:"certificates"`
	FolderId      *string  `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	PrivateKey       string  `pulumi:"privateKey"`
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	WaitValidation   *bool   `pulumi:"waitValidation"`
}

func GetCmCertificateContentOutput(ctx *pulumi.Context, args GetCmCertificateContentOutputArgs, opts ...pulumi.InvokeOption) GetCmCertificateContentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCmCertificateContentResultOutput, error) {
			args := v.(GetCmCertificateContentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getCmCertificateContent:getCmCertificateContent", args, GetCmCertificateContentResultOutput{}, options).(GetCmCertificateContentResultOutput), nil
		}).(GetCmCertificateContentResultOutput)
}

// A collection of arguments for invoking getCmCertificateContent.
type GetCmCertificateContentOutputArgs struct {
	CertificateId    pulumi.StringPtrInput `pulumi:"certificateId"`
	FolderId         pulumi.StringPtrInput `pulumi:"folderId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PrivateKeyFormat pulumi.StringPtrInput `pulumi:"privateKeyFormat"`
	WaitValidation   pulumi.BoolPtrInput   `pulumi:"waitValidation"`
}

func (GetCmCertificateContentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmCertificateContentArgs)(nil)).Elem()
}

// A collection of values returned by getCmCertificateContent.
type GetCmCertificateContentResultOutput struct{ *pulumi.OutputState }

func (GetCmCertificateContentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCmCertificateContentResult)(nil)).Elem()
}

func (o GetCmCertificateContentResultOutput) ToGetCmCertificateContentResultOutput() GetCmCertificateContentResultOutput {
	return o
}

func (o GetCmCertificateContentResultOutput) ToGetCmCertificateContentResultOutputWithContext(ctx context.Context) GetCmCertificateContentResultOutput {
	return o
}

func (o GetCmCertificateContentResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

func (o GetCmCertificateContentResultOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o GetCmCertificateContentResultOutput) FolderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.FolderId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCmCertificateContentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCmCertificateContentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCmCertificateContentResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o GetCmCertificateContentResultOutput) PrivateKeyFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *string { return v.PrivateKeyFormat }).(pulumi.StringPtrOutput)
}

func (o GetCmCertificateContentResultOutput) WaitValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCmCertificateContentResult) *bool { return v.WaitValidation }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCmCertificateContentResultOutput{})
}

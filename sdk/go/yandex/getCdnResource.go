// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCdnResource(ctx *pulumi.Context, args *LookupCdnResourceArgs, opts ...pulumi.InvokeOption) (*LookupCdnResourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCdnResourceResult
	err := ctx.Invoke("yandex:index/getCdnResource:getCdnResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCdnResource.
type LookupCdnResourceArgs struct {
	Active             *bool                         `pulumi:"active"`
	Cname              *string                       `pulumi:"cname"`
	FolderId           *string                       `pulumi:"folderId"`
	Labels             map[string]string             `pulumi:"labels"`
	Options            *GetCdnResourceOptions        `pulumi:"options"`
	OriginGroupId      *int                          `pulumi:"originGroupId"`
	OriginGroupName    *string                       `pulumi:"originGroupName"`
	OriginProtocol     *string                       `pulumi:"originProtocol"`
	ResourceId         *string                       `pulumi:"resourceId"`
	SecondaryHostnames []string                      `pulumi:"secondaryHostnames"`
	SslCertificate     *GetCdnResourceSslCertificate `pulumi:"sslCertificate"`
	UpdatedAt          *string                       `pulumi:"updatedAt"`
}

// A collection of values returned by getCdnResource.
type LookupCdnResourceResult struct {
	Active    *bool  `pulumi:"active"`
	Cname     string `pulumi:"cname"`
	CreatedAt string `pulumi:"createdAt"`
	FolderId  string `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                       `pulumi:"id"`
	Labels             map[string]string            `pulumi:"labels"`
	Options            GetCdnResourceOptions        `pulumi:"options"`
	OriginGroupId      *int                         `pulumi:"originGroupId"`
	OriginGroupName    *string                      `pulumi:"originGroupName"`
	OriginProtocol     *string                      `pulumi:"originProtocol"`
	ProviderCname      string                       `pulumi:"providerCname"`
	ResourceId         string                       `pulumi:"resourceId"`
	SecondaryHostnames []string                     `pulumi:"secondaryHostnames"`
	SslCertificate     GetCdnResourceSslCertificate `pulumi:"sslCertificate"`
	UpdatedAt          string                       `pulumi:"updatedAt"`
}

func LookupCdnResourceOutput(ctx *pulumi.Context, args LookupCdnResourceOutputArgs, opts ...pulumi.InvokeOption) LookupCdnResourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCdnResourceResultOutput, error) {
			args := v.(LookupCdnResourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getCdnResource:getCdnResource", args, LookupCdnResourceResultOutput{}, options).(LookupCdnResourceResultOutput), nil
		}).(LookupCdnResourceResultOutput)
}

// A collection of arguments for invoking getCdnResource.
type LookupCdnResourceOutputArgs struct {
	Active             pulumi.BoolPtrInput                  `pulumi:"active"`
	Cname              pulumi.StringPtrInput                `pulumi:"cname"`
	FolderId           pulumi.StringPtrInput                `pulumi:"folderId"`
	Labels             pulumi.StringMapInput                `pulumi:"labels"`
	Options            GetCdnResourceOptionsPtrInput        `pulumi:"options"`
	OriginGroupId      pulumi.IntPtrInput                   `pulumi:"originGroupId"`
	OriginGroupName    pulumi.StringPtrInput                `pulumi:"originGroupName"`
	OriginProtocol     pulumi.StringPtrInput                `pulumi:"originProtocol"`
	ResourceId         pulumi.StringPtrInput                `pulumi:"resourceId"`
	SecondaryHostnames pulumi.StringArrayInput              `pulumi:"secondaryHostnames"`
	SslCertificate     GetCdnResourceSslCertificatePtrInput `pulumi:"sslCertificate"`
	UpdatedAt          pulumi.StringPtrInput                `pulumi:"updatedAt"`
}

func (LookupCdnResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnResourceArgs)(nil)).Elem()
}

// A collection of values returned by getCdnResource.
type LookupCdnResourceResultOutput struct{ *pulumi.OutputState }

func (LookupCdnResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCdnResourceResult)(nil)).Elem()
}

func (o LookupCdnResourceResultOutput) ToLookupCdnResourceResultOutput() LookupCdnResourceResultOutput {
	return o
}

func (o LookupCdnResourceResultOutput) ToLookupCdnResourceResultOutputWithContext(ctx context.Context) LookupCdnResourceResultOutput {
	return o
}

func (o LookupCdnResourceResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupCdnResourceResultOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.Cname }).(pulumi.StringOutput)
}

func (o LookupCdnResourceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupCdnResourceResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCdnResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCdnResourceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupCdnResourceResultOutput) Options() GetCdnResourceOptionsOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) GetCdnResourceOptions { return v.Options }).(GetCdnResourceOptionsOutput)
}

func (o LookupCdnResourceResultOutput) OriginGroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) *int { return v.OriginGroupId }).(pulumi.IntPtrOutput)
}

func (o LookupCdnResourceResultOutput) OriginGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) *string { return v.OriginGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupCdnResourceResultOutput) OriginProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) *string { return v.OriginProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupCdnResourceResultOutput) ProviderCname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.ProviderCname }).(pulumi.StringOutput)
}

func (o LookupCdnResourceResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupCdnResourceResultOutput) SecondaryHostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) []string { return v.SecondaryHostnames }).(pulumi.StringArrayOutput)
}

func (o LookupCdnResourceResultOutput) SslCertificate() GetCdnResourceSslCertificateOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) GetCdnResourceSslCertificate { return v.SslCertificate }).(GetCdnResourceSslCertificateOutput)
}

func (o LookupCdnResourceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCdnResourceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCdnResourceResultOutput{})
}

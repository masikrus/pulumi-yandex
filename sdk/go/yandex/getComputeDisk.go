// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeDisk(ctx *pulumi.Context, args *LookupComputeDiskArgs, opts ...pulumi.InvokeOption) (*LookupComputeDiskResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComputeDiskResult
	err := ctx.Invoke("yandex:index/getComputeDisk:getComputeDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeDisk.
type LookupComputeDiskArgs struct {
	DiskId              *string                            `pulumi:"diskId"`
	DiskPlacementPolicy *GetComputeDiskDiskPlacementPolicy `pulumi:"diskPlacementPolicy"`
	FolderId            *string                            `pulumi:"folderId"`
	Name                *string                            `pulumi:"name"`
}

// A collection of values returned by getComputeDisk.
type LookupComputeDiskResult struct {
	BlockSize           int                                `pulumi:"blockSize"`
	CreatedAt           string                             `pulumi:"createdAt"`
	Description         string                             `pulumi:"description"`
	DiskId              string                             `pulumi:"diskId"`
	DiskPlacementPolicy *GetComputeDiskDiskPlacementPolicy `pulumi:"diskPlacementPolicy"`
	FolderId            string                             `pulumi:"folderId"`
	HardwareGenerations []GetComputeDiskHardwareGeneration `pulumi:"hardwareGenerations"`
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	ImageId     string            `pulumi:"imageId"`
	InstanceIds []string          `pulumi:"instanceIds"`
	KmsKeyId    string            `pulumi:"kmsKeyId"`
	Labels      map[string]string `pulumi:"labels"`
	Name        string            `pulumi:"name"`
	ProductIds  []string          `pulumi:"productIds"`
	Size        int               `pulumi:"size"`
	SnapshotId  string            `pulumi:"snapshotId"`
	Status      string            `pulumi:"status"`
	Type        string            `pulumi:"type"`
	Zone        string            `pulumi:"zone"`
}

func LookupComputeDiskOutput(ctx *pulumi.Context, args LookupComputeDiskOutputArgs, opts ...pulumi.InvokeOption) LookupComputeDiskResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeDiskResultOutput, error) {
			args := v.(LookupComputeDiskArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getComputeDisk:getComputeDisk", args, LookupComputeDiskResultOutput{}, options).(LookupComputeDiskResultOutput), nil
		}).(LookupComputeDiskResultOutput)
}

// A collection of arguments for invoking getComputeDisk.
type LookupComputeDiskOutputArgs struct {
	DiskId              pulumi.StringPtrInput                     `pulumi:"diskId"`
	DiskPlacementPolicy GetComputeDiskDiskPlacementPolicyPtrInput `pulumi:"diskPlacementPolicy"`
	FolderId            pulumi.StringPtrInput                     `pulumi:"folderId"`
	Name                pulumi.StringPtrInput                     `pulumi:"name"`
}

func (LookupComputeDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeDiskArgs)(nil)).Elem()
}

// A collection of values returned by getComputeDisk.
type LookupComputeDiskResultOutput struct{ *pulumi.OutputState }

func (LookupComputeDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeDiskResult)(nil)).Elem()
}

func (o LookupComputeDiskResultOutput) ToLookupComputeDiskResultOutput() LookupComputeDiskResultOutput {
	return o
}

func (o LookupComputeDiskResultOutput) ToLookupComputeDiskResultOutputWithContext(ctx context.Context) LookupComputeDiskResultOutput {
	return o
}

func (o LookupComputeDiskResultOutput) BlockSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) int { return v.BlockSize }).(pulumi.IntOutput)
}

func (o LookupComputeDiskResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.DiskId }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) DiskPlacementPolicy() GetComputeDiskDiskPlacementPolicyPtrOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) *GetComputeDiskDiskPlacementPolicy { return v.DiskPlacementPolicy }).(GetComputeDiskDiskPlacementPolicyPtrOutput)
}

func (o LookupComputeDiskResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) HardwareGenerations() GetComputeDiskHardwareGenerationArrayOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) []GetComputeDiskHardwareGeneration { return v.HardwareGenerations }).(GetComputeDiskHardwareGenerationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupComputeDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o LookupComputeDiskResultOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupComputeDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) ProductIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) []string { return v.ProductIds }).(pulumi.StringArrayOutput)
}

func (o LookupComputeDiskResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o LookupComputeDiskResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupComputeDiskResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeDiskResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeDiskResultOutput{})
}

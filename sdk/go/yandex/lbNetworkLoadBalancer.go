// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LbNetworkLoadBalancer struct {
	pulumi.CustomResourceState

	// Flag that marks the network load balancer as available to zonal shift.
	AllowZonalShift pulumi.BoolOutput `pulumi:"allowZonalShift"`
	// An AttachedTargetGroup resource.
	AttachedTargetGroups LbNetworkLoadBalancerAttachedTargetGroupArrayOutput `pulumi:"attachedTargetGroups"`
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
	// `internalAddressSpec` should be specified.
	Listeners LbNetworkLoadBalancerListenerArrayOutput `pulumi:"listeners"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewLbNetworkLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLbNetworkLoadBalancer(ctx *pulumi.Context,
	name string, args *LbNetworkLoadBalancerArgs, opts ...pulumi.ResourceOption) (*LbNetworkLoadBalancer, error) {
	if args == nil {
		args = &LbNetworkLoadBalancerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbNetworkLoadBalancer
	err := ctx.RegisterResource("yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbNetworkLoadBalancer gets an existing LbNetworkLoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbNetworkLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbNetworkLoadBalancerState, opts ...pulumi.ResourceOption) (*LbNetworkLoadBalancer, error) {
	var resource LbNetworkLoadBalancer
	err := ctx.ReadResource("yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbNetworkLoadBalancer resources.
type lbNetworkLoadBalancerState struct {
	// Flag that marks the network load balancer as available to zonal shift.
	AllowZonalShift *bool `pulumi:"allowZonalShift"`
	// An AttachedTargetGroup resource.
	AttachedTargetGroups []LbNetworkLoadBalancerAttachedTargetGroup `pulumi:"attachedTargetGroups"`
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
	// `internalAddressSpec` should be specified.
	Listeners []LbNetworkLoadBalancerListener `pulumi:"listeners"`
	// The resource name.
	Name *string `pulumi:"name"`
	// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
	RegionId *string `pulumi:"regionId"`
	// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
	Type *string `pulumi:"type"`
}

type LbNetworkLoadBalancerState struct {
	// Flag that marks the network load balancer as available to zonal shift.
	AllowZonalShift pulumi.BoolPtrInput
	// An AttachedTargetGroup resource.
	AttachedTargetGroups LbNetworkLoadBalancerAttachedTargetGroupArrayInput
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
	// `internalAddressSpec` should be specified.
	Listeners LbNetworkLoadBalancerListenerArrayInput
	// The resource name.
	Name pulumi.StringPtrInput
	// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
	RegionId pulumi.StringPtrInput
	// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
	Type pulumi.StringPtrInput
}

func (LbNetworkLoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbNetworkLoadBalancerState)(nil)).Elem()
}

type lbNetworkLoadBalancerArgs struct {
	// Flag that marks the network load balancer as available to zonal shift.
	AllowZonalShift *bool `pulumi:"allowZonalShift"`
	// An AttachedTargetGroup resource.
	AttachedTargetGroups []LbNetworkLoadBalancerAttachedTargetGroup `pulumi:"attachedTargetGroups"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
	// `internalAddressSpec` should be specified.
	Listeners []LbNetworkLoadBalancerListener `pulumi:"listeners"`
	// The resource name.
	Name *string `pulumi:"name"`
	// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
	RegionId *string `pulumi:"regionId"`
	// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a LbNetworkLoadBalancer resource.
type LbNetworkLoadBalancerArgs struct {
	// Flag that marks the network load balancer as available to zonal shift.
	AllowZonalShift pulumi.BoolPtrInput
	// An AttachedTargetGroup resource.
	AttachedTargetGroups LbNetworkLoadBalancerAttachedTargetGroupArrayInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
	// `internalAddressSpec` should be specified.
	Listeners LbNetworkLoadBalancerListenerArrayInput
	// The resource name.
	Name pulumi.StringPtrInput
	// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
	RegionId pulumi.StringPtrInput
	// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
	Type pulumi.StringPtrInput
}

func (LbNetworkLoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbNetworkLoadBalancerArgs)(nil)).Elem()
}

type LbNetworkLoadBalancerInput interface {
	pulumi.Input

	ToLbNetworkLoadBalancerOutput() LbNetworkLoadBalancerOutput
	ToLbNetworkLoadBalancerOutputWithContext(ctx context.Context) LbNetworkLoadBalancerOutput
}

func (*LbNetworkLoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LbNetworkLoadBalancer)(nil)).Elem()
}

func (i *LbNetworkLoadBalancer) ToLbNetworkLoadBalancerOutput() LbNetworkLoadBalancerOutput {
	return i.ToLbNetworkLoadBalancerOutputWithContext(context.Background())
}

func (i *LbNetworkLoadBalancer) ToLbNetworkLoadBalancerOutputWithContext(ctx context.Context) LbNetworkLoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbNetworkLoadBalancerOutput)
}

// LbNetworkLoadBalancerArrayInput is an input type that accepts LbNetworkLoadBalancerArray and LbNetworkLoadBalancerArrayOutput values.
// You can construct a concrete instance of `LbNetworkLoadBalancerArrayInput` via:
//
//	LbNetworkLoadBalancerArray{ LbNetworkLoadBalancerArgs{...} }
type LbNetworkLoadBalancerArrayInput interface {
	pulumi.Input

	ToLbNetworkLoadBalancerArrayOutput() LbNetworkLoadBalancerArrayOutput
	ToLbNetworkLoadBalancerArrayOutputWithContext(context.Context) LbNetworkLoadBalancerArrayOutput
}

type LbNetworkLoadBalancerArray []LbNetworkLoadBalancerInput

func (LbNetworkLoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbNetworkLoadBalancer)(nil)).Elem()
}

func (i LbNetworkLoadBalancerArray) ToLbNetworkLoadBalancerArrayOutput() LbNetworkLoadBalancerArrayOutput {
	return i.ToLbNetworkLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LbNetworkLoadBalancerArray) ToLbNetworkLoadBalancerArrayOutputWithContext(ctx context.Context) LbNetworkLoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbNetworkLoadBalancerArrayOutput)
}

// LbNetworkLoadBalancerMapInput is an input type that accepts LbNetworkLoadBalancerMap and LbNetworkLoadBalancerMapOutput values.
// You can construct a concrete instance of `LbNetworkLoadBalancerMapInput` via:
//
//	LbNetworkLoadBalancerMap{ "key": LbNetworkLoadBalancerArgs{...} }
type LbNetworkLoadBalancerMapInput interface {
	pulumi.Input

	ToLbNetworkLoadBalancerMapOutput() LbNetworkLoadBalancerMapOutput
	ToLbNetworkLoadBalancerMapOutputWithContext(context.Context) LbNetworkLoadBalancerMapOutput
}

type LbNetworkLoadBalancerMap map[string]LbNetworkLoadBalancerInput

func (LbNetworkLoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbNetworkLoadBalancer)(nil)).Elem()
}

func (i LbNetworkLoadBalancerMap) ToLbNetworkLoadBalancerMapOutput() LbNetworkLoadBalancerMapOutput {
	return i.ToLbNetworkLoadBalancerMapOutputWithContext(context.Background())
}

func (i LbNetworkLoadBalancerMap) ToLbNetworkLoadBalancerMapOutputWithContext(ctx context.Context) LbNetworkLoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbNetworkLoadBalancerMapOutput)
}

type LbNetworkLoadBalancerOutput struct{ *pulumi.OutputState }

func (LbNetworkLoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbNetworkLoadBalancer)(nil)).Elem()
}

func (o LbNetworkLoadBalancerOutput) ToLbNetworkLoadBalancerOutput() LbNetworkLoadBalancerOutput {
	return o
}

func (o LbNetworkLoadBalancerOutput) ToLbNetworkLoadBalancerOutputWithContext(ctx context.Context) LbNetworkLoadBalancerOutput {
	return o
}

// Flag that marks the network load balancer as available to zonal shift.
func (o LbNetworkLoadBalancerOutput) AllowZonalShift() pulumi.BoolOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.BoolOutput { return v.AllowZonalShift }).(pulumi.BoolOutput)
}

// An AttachedTargetGroup resource.
func (o LbNetworkLoadBalancerOutput) AttachedTargetGroups() LbNetworkLoadBalancerAttachedTargetGroupArrayOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) LbNetworkLoadBalancerAttachedTargetGroupArrayOutput {
		return v.AttachedTargetGroups
	}).(LbNetworkLoadBalancerAttachedTargetGroupArrayOutput)
}

// The creation timestamp of the resource.
func (o LbNetworkLoadBalancerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The `true` value means that resource is protected from accidental deletion.
func (o LbNetworkLoadBalancerOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// The resource description.
func (o LbNetworkLoadBalancerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o LbNetworkLoadBalancerOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o LbNetworkLoadBalancerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Listener specification that will be used by a network load balancer. > One of `externalAddressSpec` or
// `internalAddressSpec` should be specified.
func (o LbNetworkLoadBalancerOutput) Listeners() LbNetworkLoadBalancerListenerArrayOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) LbNetworkLoadBalancerListenerArrayOutput { return v.Listeners }).(LbNetworkLoadBalancerListenerArrayOutput)
}

// The resource name.
func (o LbNetworkLoadBalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
func (o LbNetworkLoadBalancerOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
func (o LbNetworkLoadBalancerOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbNetworkLoadBalancer) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type LbNetworkLoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LbNetworkLoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbNetworkLoadBalancer)(nil)).Elem()
}

func (o LbNetworkLoadBalancerArrayOutput) ToLbNetworkLoadBalancerArrayOutput() LbNetworkLoadBalancerArrayOutput {
	return o
}

func (o LbNetworkLoadBalancerArrayOutput) ToLbNetworkLoadBalancerArrayOutputWithContext(ctx context.Context) LbNetworkLoadBalancerArrayOutput {
	return o
}

func (o LbNetworkLoadBalancerArrayOutput) Index(i pulumi.IntInput) LbNetworkLoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbNetworkLoadBalancer {
		return vs[0].([]*LbNetworkLoadBalancer)[vs[1].(int)]
	}).(LbNetworkLoadBalancerOutput)
}

type LbNetworkLoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LbNetworkLoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbNetworkLoadBalancer)(nil)).Elem()
}

func (o LbNetworkLoadBalancerMapOutput) ToLbNetworkLoadBalancerMapOutput() LbNetworkLoadBalancerMapOutput {
	return o
}

func (o LbNetworkLoadBalancerMapOutput) ToLbNetworkLoadBalancerMapOutputWithContext(ctx context.Context) LbNetworkLoadBalancerMapOutput {
	return o
}

func (o LbNetworkLoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LbNetworkLoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbNetworkLoadBalancer {
		return vs[0].(map[string]*LbNetworkLoadBalancer)[vs[1].(string)]
	}).(LbNetworkLoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbNetworkLoadBalancerInput)(nil)).Elem(), &LbNetworkLoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbNetworkLoadBalancerArrayInput)(nil)).Elem(), LbNetworkLoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbNetworkLoadBalancerMapInput)(nil)).Elem(), LbNetworkLoadBalancerMap{})
	pulumi.RegisterOutputType(LbNetworkLoadBalancerOutput{})
	pulumi.RegisterOutputType(LbNetworkLoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LbNetworkLoadBalancerMapOutput{})
}

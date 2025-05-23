// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"errors"
	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsZone struct {
	pulumi.CustomResourceState

	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks pulumi.StringArrayOutput `pulumi:"privateNetworks"`
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	Public pulumi.BoolOutput `pulumi:"public"`
	// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDnsZone registers a new resource with the given unique name, arguments, and options.
func NewDnsZone(ctx *pulumi.Context,
	name string, args *DnsZoneArgs, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsZone
	err := ctx.RegisterResource("yandex:index/dnsZone:DnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsZone gets an existing DnsZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsZoneState, opts ...pulumi.ResourceOption) (*DnsZone, error) {
	var resource DnsZone
	err := ctx.ReadResource("yandex:index/dnsZone:DnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsZone resources.
type dnsZoneState struct {
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
	// The resource name.
	Name *string `pulumi:"name"`
	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks []string `pulumi:"privateNetworks"`
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	Public *bool `pulumi:"public"`
	// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
	Zone *string `pulumi:"zone"`
}

type DnsZoneState struct {
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
	// The resource name.
	Name pulumi.StringPtrInput
	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks pulumi.StringArrayInput
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	Public pulumi.BoolPtrInput
	// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
	Zone pulumi.StringPtrInput
}

func (DnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneState)(nil)).Elem()
}

type dnsZoneArgs struct {
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name.
	Name *string `pulumi:"name"`
	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks []string `pulumi:"privateNetworks"`
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	Public *bool `pulumi:"public"`
	// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a DnsZone resource.
type DnsZoneArgs struct {
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// The resource name.
	Name pulumi.StringPtrInput
	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks pulumi.StringArrayInput
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
	// Cloud resources.
	Public pulumi.BoolPtrInput
	// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
	Zone pulumi.StringInput
}

func (DnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsZoneArgs)(nil)).Elem()
}

type DnsZoneInput interface {
	pulumi.Input

	ToDnsZoneOutput() DnsZoneOutput
	ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput
}

func (*DnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (i *DnsZone) ToDnsZoneOutput() DnsZoneOutput {
	return i.ToDnsZoneOutputWithContext(context.Background())
}

func (i *DnsZone) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneOutput)
}

// DnsZoneArrayInput is an input type that accepts DnsZoneArray and DnsZoneArrayOutput values.
// You can construct a concrete instance of `DnsZoneArrayInput` via:
//
//	DnsZoneArray{ DnsZoneArgs{...} }
type DnsZoneArrayInput interface {
	pulumi.Input

	ToDnsZoneArrayOutput() DnsZoneArrayOutput
	ToDnsZoneArrayOutputWithContext(context.Context) DnsZoneArrayOutput
}

type DnsZoneArray []DnsZoneInput

func (DnsZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
}

func (i DnsZoneArray) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return i.ToDnsZoneArrayOutputWithContext(context.Background())
}

func (i DnsZoneArray) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneArrayOutput)
}

// DnsZoneMapInput is an input type that accepts DnsZoneMap and DnsZoneMapOutput values.
// You can construct a concrete instance of `DnsZoneMapInput` via:
//
//	DnsZoneMap{ "key": DnsZoneArgs{...} }
type DnsZoneMapInput interface {
	pulumi.Input

	ToDnsZoneMapOutput() DnsZoneMapOutput
	ToDnsZoneMapOutputWithContext(context.Context) DnsZoneMapOutput
}

type DnsZoneMap map[string]DnsZoneInput

func (DnsZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (i DnsZoneMap) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return i.ToDnsZoneMapOutputWithContext(context.Background())
}

func (i DnsZoneMap) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsZoneMapOutput)
}

type DnsZoneOutput struct{ *pulumi.OutputState }

func (DnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsZone)(nil)).Elem()
}

func (o DnsZoneOutput) ToDnsZoneOutput() DnsZoneOutput {
	return o
}

func (o DnsZoneOutput) ToDnsZoneOutputWithContext(ctx context.Context) DnsZoneOutput {
	return o
}

// The creation timestamp of the resource.
func (o DnsZoneOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The `true` value means that resource is protected from accidental deletion.
func (o DnsZoneOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// The resource description.
func (o DnsZoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o DnsZoneOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o DnsZoneOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name.
func (o DnsZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
func (o DnsZoneOutput) PrivateNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringArrayOutput { return v.PrivateNetworks }).(pulumi.StringArrayOutput)
}

// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
// Cloud resources.
func (o DnsZoneOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.BoolOutput { return v.Public }).(pulumi.BoolOutput)
}

// The DNS name of this zone, e.g. `example.com.`. Must ends with dot.
func (o DnsZoneOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsZone) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DnsZoneArrayOutput struct{ *pulumi.OutputState }

func (DnsZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsZone)(nil)).Elem()
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutput() DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) ToDnsZoneArrayOutputWithContext(ctx context.Context) DnsZoneArrayOutput {
	return o
}

func (o DnsZoneArrayOutput) Index(i pulumi.IntInput) DnsZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].([]*DnsZone)[vs[1].(int)]
	}).(DnsZoneOutput)
}

type DnsZoneMapOutput struct{ *pulumi.OutputState }

func (DnsZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsZone)(nil)).Elem()
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutput() DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) ToDnsZoneMapOutputWithContext(ctx context.Context) DnsZoneMapOutput {
	return o
}

func (o DnsZoneMapOutput) MapIndex(k pulumi.StringInput) DnsZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsZone {
		return vs[0].(map[string]*DnsZone)[vs[1].(string)]
	}).(DnsZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneInput)(nil)).Elem(), &DnsZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneArrayInput)(nil)).Elem(), DnsZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsZoneMapInput)(nil)).Elem(), DnsZoneMap{})
	pulumi.RegisterOutputType(DnsZoneOutput{})
	pulumi.RegisterOutputType(DnsZoneArrayOutput{})
	pulumi.RegisterOutputType(DnsZoneMapOutput{})
}

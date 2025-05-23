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

type VpcSubnet struct {
	pulumi.CustomResourceState

	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Options for DHCP client.
	DhcpOptions VpcSubnetDhcpOptionsPtrOutput `pulumi:"dhcpOptions"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
	// subnet.
	RouteTableId pulumi.StringPtrOutput `pulumi:"routeTableId"`
	// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
	// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
	// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
	V4CidrBlocks pulumi.StringArrayOutput `pulumi:"v4CidrBlocks"`
	// An optional list of blocks of IPv6 addresses that are owned by this subnet.
	V6CidrBlocks pulumi.StringArrayOutput `pulumi:"v6CidrBlocks"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcSubnet registers a new resource with the given unique name, arguments, and options.
func NewVpcSubnet(ctx *pulumi.Context,
	name string, args *VpcSubnetArgs, opts ...pulumi.ResourceOption) (*VpcSubnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.V4CidrBlocks == nil {
		return nil, errors.New("invalid value for required argument 'V4CidrBlocks'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcSubnet
	err := ctx.RegisterResource("yandex:index/vpcSubnet:VpcSubnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcSubnet gets an existing VpcSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcSubnetState, opts ...pulumi.ResourceOption) (*VpcSubnet, error) {
	var resource VpcSubnet
	err := ctx.ReadResource("yandex:index/vpcSubnet:VpcSubnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcSubnet resources.
type vpcSubnetState struct {
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Options for DHCP client.
	DhcpOptions *VpcSubnetDhcpOptions `pulumi:"dhcpOptions"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name.
	Name *string `pulumi:"name"`
	// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
	NetworkId *string `pulumi:"networkId"`
	// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
	// subnet.
	RouteTableId *string `pulumi:"routeTableId"`
	// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
	// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
	// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
	V4CidrBlocks []string `pulumi:"v4CidrBlocks"`
	// An optional list of blocks of IPv6 addresses that are owned by this subnet.
	V6CidrBlocks []string `pulumi:"v6CidrBlocks"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone *string `pulumi:"zone"`
}

type VpcSubnetState struct {
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Options for DHCP client.
	DhcpOptions VpcSubnetDhcpOptionsPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// The resource name.
	Name pulumi.StringPtrInput
	// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
	NetworkId pulumi.StringPtrInput
	// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
	// subnet.
	RouteTableId pulumi.StringPtrInput
	// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
	// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
	// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
	V4CidrBlocks pulumi.StringArrayInput
	// An optional list of blocks of IPv6 addresses that are owned by this subnet.
	V6CidrBlocks pulumi.StringArrayInput
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringPtrInput
}

func (VpcSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSubnetState)(nil)).Elem()
}

type vpcSubnetArgs struct {
	// The resource description.
	Description *string `pulumi:"description"`
	// Options for DHCP client.
	DhcpOptions *VpcSubnetDhcpOptions `pulumi:"dhcpOptions"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name.
	Name *string `pulumi:"name"`
	// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
	NetworkId string `pulumi:"networkId"`
	// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
	// subnet.
	RouteTableId *string `pulumi:"routeTableId"`
	// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
	// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
	// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
	V4CidrBlocks []string `pulumi:"v4CidrBlocks"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcSubnet resource.
type VpcSubnetArgs struct {
	// The resource description.
	Description pulumi.StringPtrInput
	// Options for DHCP client.
	DhcpOptions VpcSubnetDhcpOptionsPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// The resource name.
	Name pulumi.StringPtrInput
	// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
	NetworkId pulumi.StringInput
	// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
	// subnet.
	RouteTableId pulumi.StringPtrInput
	// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
	// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
	// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
	V4CidrBlocks pulumi.StringArrayInput
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringPtrInput
}

func (VpcSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcSubnetArgs)(nil)).Elem()
}

type VpcSubnetInput interface {
	pulumi.Input

	ToVpcSubnetOutput() VpcSubnetOutput
	ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput
}

func (*VpcSubnet) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSubnet)(nil)).Elem()
}

func (i *VpcSubnet) ToVpcSubnetOutput() VpcSubnetOutput {
	return i.ToVpcSubnetOutputWithContext(context.Background())
}

func (i *VpcSubnet) ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetOutput)
}

// VpcSubnetArrayInput is an input type that accepts VpcSubnetArray and VpcSubnetArrayOutput values.
// You can construct a concrete instance of `VpcSubnetArrayInput` via:
//
//	VpcSubnetArray{ VpcSubnetArgs{...} }
type VpcSubnetArrayInput interface {
	pulumi.Input

	ToVpcSubnetArrayOutput() VpcSubnetArrayOutput
	ToVpcSubnetArrayOutputWithContext(context.Context) VpcSubnetArrayOutput
}

type VpcSubnetArray []VpcSubnetInput

func (VpcSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcSubnet)(nil)).Elem()
}

func (i VpcSubnetArray) ToVpcSubnetArrayOutput() VpcSubnetArrayOutput {
	return i.ToVpcSubnetArrayOutputWithContext(context.Background())
}

func (i VpcSubnetArray) ToVpcSubnetArrayOutputWithContext(ctx context.Context) VpcSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetArrayOutput)
}

// VpcSubnetMapInput is an input type that accepts VpcSubnetMap and VpcSubnetMapOutput values.
// You can construct a concrete instance of `VpcSubnetMapInput` via:
//
//	VpcSubnetMap{ "key": VpcSubnetArgs{...} }
type VpcSubnetMapInput interface {
	pulumi.Input

	ToVpcSubnetMapOutput() VpcSubnetMapOutput
	ToVpcSubnetMapOutputWithContext(context.Context) VpcSubnetMapOutput
}

type VpcSubnetMap map[string]VpcSubnetInput

func (VpcSubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcSubnet)(nil)).Elem()
}

func (i VpcSubnetMap) ToVpcSubnetMapOutput() VpcSubnetMapOutput {
	return i.ToVpcSubnetMapOutputWithContext(context.Background())
}

func (i VpcSubnetMap) ToVpcSubnetMapOutputWithContext(ctx context.Context) VpcSubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcSubnetMapOutput)
}

type VpcSubnetOutput struct{ *pulumi.OutputState }

func (VpcSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetOutput) ToVpcSubnetOutput() VpcSubnetOutput {
	return o
}

func (o VpcSubnetOutput) ToVpcSubnetOutputWithContext(ctx context.Context) VpcSubnetOutput {
	return o
}

// The creation timestamp of the resource.
func (o VpcSubnetOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The resource description.
func (o VpcSubnetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Options for DHCP client.
func (o VpcSubnetOutput) DhcpOptions() VpcSubnetDhcpOptionsPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) VpcSubnetDhcpOptionsPtrOutput { return v.DhcpOptions }).(VpcSubnetDhcpOptionsPtrOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o VpcSubnetOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o VpcSubnetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name.
func (o VpcSubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
func (o VpcSubnetOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
// subnet.
func (o VpcSubnetOutput) RouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringPtrOutput { return v.RouteTableId }).(pulumi.StringPtrOutput)
}

// A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
// subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
// network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
func (o VpcSubnetOutput) V4CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringArrayOutput { return v.V4CidrBlocks }).(pulumi.StringArrayOutput)
}

// An optional list of blocks of IPv6 addresses that are owned by this subnet.
func (o VpcSubnetOutput) V6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringArrayOutput { return v.V6CidrBlocks }).(pulumi.StringArrayOutput)
}

// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
// provided, the default provider zone will be used.
func (o VpcSubnetOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcSubnet) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcSubnetArrayOutput struct{ *pulumi.OutputState }

func (VpcSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetArrayOutput) ToVpcSubnetArrayOutput() VpcSubnetArrayOutput {
	return o
}

func (o VpcSubnetArrayOutput) ToVpcSubnetArrayOutputWithContext(ctx context.Context) VpcSubnetArrayOutput {
	return o
}

func (o VpcSubnetArrayOutput) Index(i pulumi.IntInput) VpcSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcSubnet {
		return vs[0].([]*VpcSubnet)[vs[1].(int)]
	}).(VpcSubnetOutput)
}

type VpcSubnetMapOutput struct{ *pulumi.OutputState }

func (VpcSubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcSubnet)(nil)).Elem()
}

func (o VpcSubnetMapOutput) ToVpcSubnetMapOutput() VpcSubnetMapOutput {
	return o
}

func (o VpcSubnetMapOutput) ToVpcSubnetMapOutputWithContext(ctx context.Context) VpcSubnetMapOutput {
	return o
}

func (o VpcSubnetMapOutput) MapIndex(k pulumi.StringInput) VpcSubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcSubnet {
		return vs[0].(map[string]*VpcSubnet)[vs[1].(string)]
	}).(VpcSubnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetInput)(nil)).Elem(), &VpcSubnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetArrayInput)(nil)).Elem(), VpcSubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcSubnetMapInput)(nil)).Elem(), VpcSubnetMap{})
	pulumi.RegisterOutputType(VpcSubnetOutput{})
	pulumi.RegisterOutputType(VpcSubnetArrayOutput{})
	pulumi.RegisterOutputType(VpcSubnetMapOutput{})
}

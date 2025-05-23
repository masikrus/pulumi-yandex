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

type ContainerRegistryIpPermission struct {
	pulumi.CustomResourceState

	// List of configured CIDRs, from which `pull` is allowed.
	Pulls pulumi.StringArrayOutput `pulumi:"pulls"`
	// List of configured CIDRs, from which `push` is allowed.
	Pushes pulumi.StringArrayOutput `pulumi:"pushes"`
	// The ID of the registry that ip restrictions applied to.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
}

// NewContainerRegistryIpPermission registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryIpPermission(ctx *pulumi.Context,
	name string, args *ContainerRegistryIpPermissionArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryIpPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistryIpPermission
	err := ctx.RegisterResource("yandex:index/containerRegistryIpPermission:ContainerRegistryIpPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryIpPermission gets an existing ContainerRegistryIpPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryIpPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryIpPermissionState, opts ...pulumi.ResourceOption) (*ContainerRegistryIpPermission, error) {
	var resource ContainerRegistryIpPermission
	err := ctx.ReadResource("yandex:index/containerRegistryIpPermission:ContainerRegistryIpPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryIpPermission resources.
type containerRegistryIpPermissionState struct {
	// List of configured CIDRs, from which `pull` is allowed.
	Pulls []string `pulumi:"pulls"`
	// List of configured CIDRs, from which `push` is allowed.
	Pushes []string `pulumi:"pushes"`
	// The ID of the registry that ip restrictions applied to.
	RegistryId *string `pulumi:"registryId"`
}

type ContainerRegistryIpPermissionState struct {
	// List of configured CIDRs, from which `pull` is allowed.
	Pulls pulumi.StringArrayInput
	// List of configured CIDRs, from which `push` is allowed.
	Pushes pulumi.StringArrayInput
	// The ID of the registry that ip restrictions applied to.
	RegistryId pulumi.StringPtrInput
}

func (ContainerRegistryIpPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIpPermissionState)(nil)).Elem()
}

type containerRegistryIpPermissionArgs struct {
	// List of configured CIDRs, from which `pull` is allowed.
	Pulls []string `pulumi:"pulls"`
	// List of configured CIDRs, from which `push` is allowed.
	Pushes []string `pulumi:"pushes"`
	// The ID of the registry that ip restrictions applied to.
	RegistryId string `pulumi:"registryId"`
}

// The set of arguments for constructing a ContainerRegistryIpPermission resource.
type ContainerRegistryIpPermissionArgs struct {
	// List of configured CIDRs, from which `pull` is allowed.
	Pulls pulumi.StringArrayInput
	// List of configured CIDRs, from which `push` is allowed.
	Pushes pulumi.StringArrayInput
	// The ID of the registry that ip restrictions applied to.
	RegistryId pulumi.StringInput
}

func (ContainerRegistryIpPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIpPermissionArgs)(nil)).Elem()
}

type ContainerRegistryIpPermissionInput interface {
	pulumi.Input

	ToContainerRegistryIpPermissionOutput() ContainerRegistryIpPermissionOutput
	ToContainerRegistryIpPermissionOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionOutput
}

func (*ContainerRegistryIpPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIpPermission)(nil)).Elem()
}

func (i *ContainerRegistryIpPermission) ToContainerRegistryIpPermissionOutput() ContainerRegistryIpPermissionOutput {
	return i.ToContainerRegistryIpPermissionOutputWithContext(context.Background())
}

func (i *ContainerRegistryIpPermission) ToContainerRegistryIpPermissionOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIpPermissionOutput)
}

// ContainerRegistryIpPermissionArrayInput is an input type that accepts ContainerRegistryIpPermissionArray and ContainerRegistryIpPermissionArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryIpPermissionArrayInput` via:
//
//	ContainerRegistryIpPermissionArray{ ContainerRegistryIpPermissionArgs{...} }
type ContainerRegistryIpPermissionArrayInput interface {
	pulumi.Input

	ToContainerRegistryIpPermissionArrayOutput() ContainerRegistryIpPermissionArrayOutput
	ToContainerRegistryIpPermissionArrayOutputWithContext(context.Context) ContainerRegistryIpPermissionArrayOutput
}

type ContainerRegistryIpPermissionArray []ContainerRegistryIpPermissionInput

func (ContainerRegistryIpPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIpPermission)(nil)).Elem()
}

func (i ContainerRegistryIpPermissionArray) ToContainerRegistryIpPermissionArrayOutput() ContainerRegistryIpPermissionArrayOutput {
	return i.ToContainerRegistryIpPermissionArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryIpPermissionArray) ToContainerRegistryIpPermissionArrayOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIpPermissionArrayOutput)
}

// ContainerRegistryIpPermissionMapInput is an input type that accepts ContainerRegistryIpPermissionMap and ContainerRegistryIpPermissionMapOutput values.
// You can construct a concrete instance of `ContainerRegistryIpPermissionMapInput` via:
//
//	ContainerRegistryIpPermissionMap{ "key": ContainerRegistryIpPermissionArgs{...} }
type ContainerRegistryIpPermissionMapInput interface {
	pulumi.Input

	ToContainerRegistryIpPermissionMapOutput() ContainerRegistryIpPermissionMapOutput
	ToContainerRegistryIpPermissionMapOutputWithContext(context.Context) ContainerRegistryIpPermissionMapOutput
}

type ContainerRegistryIpPermissionMap map[string]ContainerRegistryIpPermissionInput

func (ContainerRegistryIpPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIpPermission)(nil)).Elem()
}

func (i ContainerRegistryIpPermissionMap) ToContainerRegistryIpPermissionMapOutput() ContainerRegistryIpPermissionMapOutput {
	return i.ToContainerRegistryIpPermissionMapOutputWithContext(context.Background())
}

func (i ContainerRegistryIpPermissionMap) ToContainerRegistryIpPermissionMapOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIpPermissionMapOutput)
}

type ContainerRegistryIpPermissionOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIpPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIpPermission)(nil)).Elem()
}

func (o ContainerRegistryIpPermissionOutput) ToContainerRegistryIpPermissionOutput() ContainerRegistryIpPermissionOutput {
	return o
}

func (o ContainerRegistryIpPermissionOutput) ToContainerRegistryIpPermissionOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionOutput {
	return o
}

// List of configured CIDRs, from which `pull` is allowed.
func (o ContainerRegistryIpPermissionOutput) Pulls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerRegistryIpPermission) pulumi.StringArrayOutput { return v.Pulls }).(pulumi.StringArrayOutput)
}

// List of configured CIDRs, from which `push` is allowed.
func (o ContainerRegistryIpPermissionOutput) Pushes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerRegistryIpPermission) pulumi.StringArrayOutput { return v.Pushes }).(pulumi.StringArrayOutput)
}

// The ID of the registry that ip restrictions applied to.
func (o ContainerRegistryIpPermissionOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryIpPermission) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

type ContainerRegistryIpPermissionArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIpPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIpPermission)(nil)).Elem()
}

func (o ContainerRegistryIpPermissionArrayOutput) ToContainerRegistryIpPermissionArrayOutput() ContainerRegistryIpPermissionArrayOutput {
	return o
}

func (o ContainerRegistryIpPermissionArrayOutput) ToContainerRegistryIpPermissionArrayOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionArrayOutput {
	return o
}

func (o ContainerRegistryIpPermissionArrayOutput) Index(i pulumi.IntInput) ContainerRegistryIpPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistryIpPermission {
		return vs[0].([]*ContainerRegistryIpPermission)[vs[1].(int)]
	}).(ContainerRegistryIpPermissionOutput)
}

type ContainerRegistryIpPermissionMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIpPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIpPermission)(nil)).Elem()
}

func (o ContainerRegistryIpPermissionMapOutput) ToContainerRegistryIpPermissionMapOutput() ContainerRegistryIpPermissionMapOutput {
	return o
}

func (o ContainerRegistryIpPermissionMapOutput) ToContainerRegistryIpPermissionMapOutputWithContext(ctx context.Context) ContainerRegistryIpPermissionMapOutput {
	return o
}

func (o ContainerRegistryIpPermissionMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryIpPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistryIpPermission {
		return vs[0].(map[string]*ContainerRegistryIpPermission)[vs[1].(string)]
	}).(ContainerRegistryIpPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIpPermissionInput)(nil)).Elem(), &ContainerRegistryIpPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIpPermissionArrayInput)(nil)).Elem(), ContainerRegistryIpPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIpPermissionMapInput)(nil)).Elem(), ContainerRegistryIpPermissionMap{})
	pulumi.RegisterOutputType(ContainerRegistryIpPermissionOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIpPermissionArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIpPermissionMapOutput{})
}

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

type ResourcemanagerFolderIamBinding struct {
	pulumi.CustomResourceState

	// The ID of the folder to attach a policy to.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewResourcemanagerFolderIamBinding registers a new resource with the given unique name, arguments, and options.
func NewResourcemanagerFolderIamBinding(ctx *pulumi.Context,
	name string, args *ResourcemanagerFolderIamBindingArgs, opts ...pulumi.ResourceOption) (*ResourcemanagerFolderIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourcemanagerFolderIamBinding
	err := ctx.RegisterResource("yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcemanagerFolderIamBinding gets an existing ResourcemanagerFolderIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcemanagerFolderIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcemanagerFolderIamBindingState, opts ...pulumi.ResourceOption) (*ResourcemanagerFolderIamBinding, error) {
	var resource ResourcemanagerFolderIamBinding
	err := ctx.ReadResource("yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcemanagerFolderIamBinding resources.
type resourcemanagerFolderIamBindingState struct {
	// The ID of the folder to attach a policy to.
	FolderId *string `pulumi:"folderId"`
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Members []string `pulumi:"members"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type ResourcemanagerFolderIamBindingState struct {
	// The ID of the folder to attach a policy to.
	FolderId pulumi.StringPtrInput
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerFolderIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderIamBindingState)(nil)).Elem()
}

type resourcemanagerFolderIamBindingArgs struct {
	// The ID of the folder to attach a policy to.
	FolderId string `pulumi:"folderId"`
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Members []string `pulumi:"members"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a ResourcemanagerFolderIamBinding resource.
type ResourcemanagerFolderIamBindingArgs struct {
	// The ID of the folder to attach a policy to.
	FolderId pulumi.StringInput
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (ResourcemanagerFolderIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcemanagerFolderIamBindingArgs)(nil)).Elem()
}

type ResourcemanagerFolderIamBindingInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamBindingOutput() ResourcemanagerFolderIamBindingOutput
	ToResourcemanagerFolderIamBindingOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingOutput
}

func (*ResourcemanagerFolderIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (i *ResourcemanagerFolderIamBinding) ToResourcemanagerFolderIamBindingOutput() ResourcemanagerFolderIamBindingOutput {
	return i.ToResourcemanagerFolderIamBindingOutputWithContext(context.Background())
}

func (i *ResourcemanagerFolderIamBinding) ToResourcemanagerFolderIamBindingOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamBindingOutput)
}

// ResourcemanagerFolderIamBindingArrayInput is an input type that accepts ResourcemanagerFolderIamBindingArray and ResourcemanagerFolderIamBindingArrayOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderIamBindingArrayInput` via:
//
//	ResourcemanagerFolderIamBindingArray{ ResourcemanagerFolderIamBindingArgs{...} }
type ResourcemanagerFolderIamBindingArrayInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamBindingArrayOutput() ResourcemanagerFolderIamBindingArrayOutput
	ToResourcemanagerFolderIamBindingArrayOutputWithContext(context.Context) ResourcemanagerFolderIamBindingArrayOutput
}

type ResourcemanagerFolderIamBindingArray []ResourcemanagerFolderIamBindingInput

func (ResourcemanagerFolderIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (i ResourcemanagerFolderIamBindingArray) ToResourcemanagerFolderIamBindingArrayOutput() ResourcemanagerFolderIamBindingArrayOutput {
	return i.ToResourcemanagerFolderIamBindingArrayOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderIamBindingArray) ToResourcemanagerFolderIamBindingArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamBindingArrayOutput)
}

// ResourcemanagerFolderIamBindingMapInput is an input type that accepts ResourcemanagerFolderIamBindingMap and ResourcemanagerFolderIamBindingMapOutput values.
// You can construct a concrete instance of `ResourcemanagerFolderIamBindingMapInput` via:
//
//	ResourcemanagerFolderIamBindingMap{ "key": ResourcemanagerFolderIamBindingArgs{...} }
type ResourcemanagerFolderIamBindingMapInput interface {
	pulumi.Input

	ToResourcemanagerFolderIamBindingMapOutput() ResourcemanagerFolderIamBindingMapOutput
	ToResourcemanagerFolderIamBindingMapOutputWithContext(context.Context) ResourcemanagerFolderIamBindingMapOutput
}

type ResourcemanagerFolderIamBindingMap map[string]ResourcemanagerFolderIamBindingInput

func (ResourcemanagerFolderIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (i ResourcemanagerFolderIamBindingMap) ToResourcemanagerFolderIamBindingMapOutput() ResourcemanagerFolderIamBindingMapOutput {
	return i.ToResourcemanagerFolderIamBindingMapOutputWithContext(context.Background())
}

func (i ResourcemanagerFolderIamBindingMap) ToResourcemanagerFolderIamBindingMapOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcemanagerFolderIamBindingMapOutput)
}

type ResourcemanagerFolderIamBindingOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (o ResourcemanagerFolderIamBindingOutput) ToResourcemanagerFolderIamBindingOutput() ResourcemanagerFolderIamBindingOutput {
	return o
}

func (o ResourcemanagerFolderIamBindingOutput) ToResourcemanagerFolderIamBindingOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingOutput {
	return o
}

// The ID of the folder to attach a policy to.
func (o ResourcemanagerFolderIamBindingOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcemanagerFolderIamBinding) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
func (o ResourcemanagerFolderIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourcemanagerFolderIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
func (o ResourcemanagerFolderIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcemanagerFolderIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ResourcemanagerFolderIamBindingOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourcemanagerFolderIamBinding) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

type ResourcemanagerFolderIamBindingArrayOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (o ResourcemanagerFolderIamBindingArrayOutput) ToResourcemanagerFolderIamBindingArrayOutput() ResourcemanagerFolderIamBindingArrayOutput {
	return o
}

func (o ResourcemanagerFolderIamBindingArrayOutput) ToResourcemanagerFolderIamBindingArrayOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingArrayOutput {
	return o
}

func (o ResourcemanagerFolderIamBindingArrayOutput) Index(i pulumi.IntInput) ResourcemanagerFolderIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcemanagerFolderIamBinding {
		return vs[0].([]*ResourcemanagerFolderIamBinding)[vs[1].(int)]
	}).(ResourcemanagerFolderIamBindingOutput)
}

type ResourcemanagerFolderIamBindingMapOutput struct{ *pulumi.OutputState }

func (ResourcemanagerFolderIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcemanagerFolderIamBinding)(nil)).Elem()
}

func (o ResourcemanagerFolderIamBindingMapOutput) ToResourcemanagerFolderIamBindingMapOutput() ResourcemanagerFolderIamBindingMapOutput {
	return o
}

func (o ResourcemanagerFolderIamBindingMapOutput) ToResourcemanagerFolderIamBindingMapOutputWithContext(ctx context.Context) ResourcemanagerFolderIamBindingMapOutput {
	return o
}

func (o ResourcemanagerFolderIamBindingMapOutput) MapIndex(k pulumi.StringInput) ResourcemanagerFolderIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcemanagerFolderIamBinding {
		return vs[0].(map[string]*ResourcemanagerFolderIamBinding)[vs[1].(string)]
	}).(ResourcemanagerFolderIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderIamBindingInput)(nil)).Elem(), &ResourcemanagerFolderIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderIamBindingArrayInput)(nil)).Elem(), ResourcemanagerFolderIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcemanagerFolderIamBindingMapInput)(nil)).Elem(), ResourcemanagerFolderIamBindingMap{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamBindingOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamBindingArrayOutput{})
	pulumi.RegisterOutputType(ResourcemanagerFolderIamBindingMapOutput{})
}

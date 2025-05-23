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

type KmsSymmetricKeyIamBinding struct {
	pulumi.CustomResourceState

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
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringOutput `pulumi:"symmetricKeyId"`
}

// NewKmsSymmetricKeyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewKmsSymmetricKeyIamBinding(ctx *pulumi.Context,
	name string, args *KmsSymmetricKeyIamBindingArgs, opts ...pulumi.ResourceOption) (*KmsSymmetricKeyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.SymmetricKeyId == nil {
		return nil, errors.New("invalid value for required argument 'SymmetricKeyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KmsSymmetricKeyIamBinding
	err := ctx.RegisterResource("yandex:index/kmsSymmetricKeyIamBinding:KmsSymmetricKeyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsSymmetricKeyIamBinding gets an existing KmsSymmetricKeyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsSymmetricKeyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsSymmetricKeyIamBindingState, opts ...pulumi.ResourceOption) (*KmsSymmetricKeyIamBinding, error) {
	var resource KmsSymmetricKeyIamBinding
	err := ctx.ReadResource("yandex:index/kmsSymmetricKeyIamBinding:KmsSymmetricKeyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsSymmetricKeyIamBinding resources.
type kmsSymmetricKeyIamBindingState struct {
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
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId *string `pulumi:"symmetricKeyId"`
}

type KmsSymmetricKeyIamBindingState struct {
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
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringPtrInput
}

func (KmsSymmetricKeyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSymmetricKeyIamBindingState)(nil)).Elem()
}

type kmsSymmetricKeyIamBindingArgs struct {
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
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId string `pulumi:"symmetricKeyId"`
}

// The set of arguments for constructing a KmsSymmetricKeyIamBinding resource.
type KmsSymmetricKeyIamBindingArgs struct {
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
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringInput
}

func (KmsSymmetricKeyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSymmetricKeyIamBindingArgs)(nil)).Elem()
}

type KmsSymmetricKeyIamBindingInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamBindingOutput() KmsSymmetricKeyIamBindingOutput
	ToKmsSymmetricKeyIamBindingOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingOutput
}

func (*KmsSymmetricKeyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (i *KmsSymmetricKeyIamBinding) ToKmsSymmetricKeyIamBindingOutput() KmsSymmetricKeyIamBindingOutput {
	return i.ToKmsSymmetricKeyIamBindingOutputWithContext(context.Background())
}

func (i *KmsSymmetricKeyIamBinding) ToKmsSymmetricKeyIamBindingOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamBindingOutput)
}

// KmsSymmetricKeyIamBindingArrayInput is an input type that accepts KmsSymmetricKeyIamBindingArray and KmsSymmetricKeyIamBindingArrayOutput values.
// You can construct a concrete instance of `KmsSymmetricKeyIamBindingArrayInput` via:
//
//	KmsSymmetricKeyIamBindingArray{ KmsSymmetricKeyIamBindingArgs{...} }
type KmsSymmetricKeyIamBindingArrayInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamBindingArrayOutput() KmsSymmetricKeyIamBindingArrayOutput
	ToKmsSymmetricKeyIamBindingArrayOutputWithContext(context.Context) KmsSymmetricKeyIamBindingArrayOutput
}

type KmsSymmetricKeyIamBindingArray []KmsSymmetricKeyIamBindingInput

func (KmsSymmetricKeyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (i KmsSymmetricKeyIamBindingArray) ToKmsSymmetricKeyIamBindingArrayOutput() KmsSymmetricKeyIamBindingArrayOutput {
	return i.ToKmsSymmetricKeyIamBindingArrayOutputWithContext(context.Background())
}

func (i KmsSymmetricKeyIamBindingArray) ToKmsSymmetricKeyIamBindingArrayOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamBindingArrayOutput)
}

// KmsSymmetricKeyIamBindingMapInput is an input type that accepts KmsSymmetricKeyIamBindingMap and KmsSymmetricKeyIamBindingMapOutput values.
// You can construct a concrete instance of `KmsSymmetricKeyIamBindingMapInput` via:
//
//	KmsSymmetricKeyIamBindingMap{ "key": KmsSymmetricKeyIamBindingArgs{...} }
type KmsSymmetricKeyIamBindingMapInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamBindingMapOutput() KmsSymmetricKeyIamBindingMapOutput
	ToKmsSymmetricKeyIamBindingMapOutputWithContext(context.Context) KmsSymmetricKeyIamBindingMapOutput
}

type KmsSymmetricKeyIamBindingMap map[string]KmsSymmetricKeyIamBindingInput

func (KmsSymmetricKeyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (i KmsSymmetricKeyIamBindingMap) ToKmsSymmetricKeyIamBindingMapOutput() KmsSymmetricKeyIamBindingMapOutput {
	return i.ToKmsSymmetricKeyIamBindingMapOutputWithContext(context.Background())
}

func (i KmsSymmetricKeyIamBindingMap) ToKmsSymmetricKeyIamBindingMapOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamBindingMapOutput)
}

type KmsSymmetricKeyIamBindingOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (o KmsSymmetricKeyIamBindingOutput) ToKmsSymmetricKeyIamBindingOutput() KmsSymmetricKeyIamBindingOutput {
	return o
}

func (o KmsSymmetricKeyIamBindingOutput) ToKmsSymmetricKeyIamBindingOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingOutput {
	return o
}

// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
func (o KmsSymmetricKeyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
func (o KmsSymmetricKeyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o KmsSymmetricKeyIamBindingOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamBinding) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
func (o KmsSymmetricKeyIamBindingOutput) SymmetricKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamBinding) pulumi.StringOutput { return v.SymmetricKeyId }).(pulumi.StringOutput)
}

type KmsSymmetricKeyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (o KmsSymmetricKeyIamBindingArrayOutput) ToKmsSymmetricKeyIamBindingArrayOutput() KmsSymmetricKeyIamBindingArrayOutput {
	return o
}

func (o KmsSymmetricKeyIamBindingArrayOutput) ToKmsSymmetricKeyIamBindingArrayOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingArrayOutput {
	return o
}

func (o KmsSymmetricKeyIamBindingArrayOutput) Index(i pulumi.IntInput) KmsSymmetricKeyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsSymmetricKeyIamBinding {
		return vs[0].([]*KmsSymmetricKeyIamBinding)[vs[1].(int)]
	}).(KmsSymmetricKeyIamBindingOutput)
}

type KmsSymmetricKeyIamBindingMapOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSymmetricKeyIamBinding)(nil)).Elem()
}

func (o KmsSymmetricKeyIamBindingMapOutput) ToKmsSymmetricKeyIamBindingMapOutput() KmsSymmetricKeyIamBindingMapOutput {
	return o
}

func (o KmsSymmetricKeyIamBindingMapOutput) ToKmsSymmetricKeyIamBindingMapOutputWithContext(ctx context.Context) KmsSymmetricKeyIamBindingMapOutput {
	return o
}

func (o KmsSymmetricKeyIamBindingMapOutput) MapIndex(k pulumi.StringInput) KmsSymmetricKeyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsSymmetricKeyIamBinding {
		return vs[0].(map[string]*KmsSymmetricKeyIamBinding)[vs[1].(string)]
	}).(KmsSymmetricKeyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamBindingInput)(nil)).Elem(), &KmsSymmetricKeyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamBindingArrayInput)(nil)).Elem(), KmsSymmetricKeyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamBindingMapInput)(nil)).Elem(), KmsSymmetricKeyIamBindingMap{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamBindingOutput{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamBindingMapOutput{})
}

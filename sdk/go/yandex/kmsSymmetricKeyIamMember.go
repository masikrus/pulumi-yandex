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

type KmsSymmetricKeyIamMember struct {
	pulumi.CustomResourceState

	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringOutput `pulumi:"symmetricKeyId"`
}

// NewKmsSymmetricKeyIamMember registers a new resource with the given unique name, arguments, and options.
func NewKmsSymmetricKeyIamMember(ctx *pulumi.Context,
	name string, args *KmsSymmetricKeyIamMemberArgs, opts ...pulumi.ResourceOption) (*KmsSymmetricKeyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.SymmetricKeyId == nil {
		return nil, errors.New("invalid value for required argument 'SymmetricKeyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KmsSymmetricKeyIamMember
	err := ctx.RegisterResource("yandex:index/kmsSymmetricKeyIamMember:KmsSymmetricKeyIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsSymmetricKeyIamMember gets an existing KmsSymmetricKeyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsSymmetricKeyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsSymmetricKeyIamMemberState, opts ...pulumi.ResourceOption) (*KmsSymmetricKeyIamMember, error) {
	var resource KmsSymmetricKeyIamMember
	err := ctx.ReadResource("yandex:index/kmsSymmetricKeyIamMember:KmsSymmetricKeyIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsSymmetricKeyIamMember resources.
type kmsSymmetricKeyIamMemberState struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member *string `pulumi:"member"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId *string `pulumi:"symmetricKeyId"`
}

type KmsSymmetricKeyIamMemberState struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member pulumi.StringPtrInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringPtrInput
}

func (KmsSymmetricKeyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSymmetricKeyIamMemberState)(nil)).Elem()
}

type kmsSymmetricKeyIamMemberArgs struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member string `pulumi:"member"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId string `pulumi:"symmetricKeyId"`
}

// The set of arguments for constructing a KmsSymmetricKeyIamMember resource.
type KmsSymmetricKeyIamMemberArgs struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member pulumi.StringInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
	SymmetricKeyId pulumi.StringInput
}

func (KmsSymmetricKeyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsSymmetricKeyIamMemberArgs)(nil)).Elem()
}

type KmsSymmetricKeyIamMemberInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamMemberOutput() KmsSymmetricKeyIamMemberOutput
	ToKmsSymmetricKeyIamMemberOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberOutput
}

func (*KmsSymmetricKeyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (i *KmsSymmetricKeyIamMember) ToKmsSymmetricKeyIamMemberOutput() KmsSymmetricKeyIamMemberOutput {
	return i.ToKmsSymmetricKeyIamMemberOutputWithContext(context.Background())
}

func (i *KmsSymmetricKeyIamMember) ToKmsSymmetricKeyIamMemberOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamMemberOutput)
}

// KmsSymmetricKeyIamMemberArrayInput is an input type that accepts KmsSymmetricKeyIamMemberArray and KmsSymmetricKeyIamMemberArrayOutput values.
// You can construct a concrete instance of `KmsSymmetricKeyIamMemberArrayInput` via:
//
//	KmsSymmetricKeyIamMemberArray{ KmsSymmetricKeyIamMemberArgs{...} }
type KmsSymmetricKeyIamMemberArrayInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamMemberArrayOutput() KmsSymmetricKeyIamMemberArrayOutput
	ToKmsSymmetricKeyIamMemberArrayOutputWithContext(context.Context) KmsSymmetricKeyIamMemberArrayOutput
}

type KmsSymmetricKeyIamMemberArray []KmsSymmetricKeyIamMemberInput

func (KmsSymmetricKeyIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (i KmsSymmetricKeyIamMemberArray) ToKmsSymmetricKeyIamMemberArrayOutput() KmsSymmetricKeyIamMemberArrayOutput {
	return i.ToKmsSymmetricKeyIamMemberArrayOutputWithContext(context.Background())
}

func (i KmsSymmetricKeyIamMemberArray) ToKmsSymmetricKeyIamMemberArrayOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamMemberArrayOutput)
}

// KmsSymmetricKeyIamMemberMapInput is an input type that accepts KmsSymmetricKeyIamMemberMap and KmsSymmetricKeyIamMemberMapOutput values.
// You can construct a concrete instance of `KmsSymmetricKeyIamMemberMapInput` via:
//
//	KmsSymmetricKeyIamMemberMap{ "key": KmsSymmetricKeyIamMemberArgs{...} }
type KmsSymmetricKeyIamMemberMapInput interface {
	pulumi.Input

	ToKmsSymmetricKeyIamMemberMapOutput() KmsSymmetricKeyIamMemberMapOutput
	ToKmsSymmetricKeyIamMemberMapOutputWithContext(context.Context) KmsSymmetricKeyIamMemberMapOutput
}

type KmsSymmetricKeyIamMemberMap map[string]KmsSymmetricKeyIamMemberInput

func (KmsSymmetricKeyIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (i KmsSymmetricKeyIamMemberMap) ToKmsSymmetricKeyIamMemberMapOutput() KmsSymmetricKeyIamMemberMapOutput {
	return i.ToKmsSymmetricKeyIamMemberMapOutputWithContext(context.Background())
}

func (i KmsSymmetricKeyIamMemberMap) ToKmsSymmetricKeyIamMemberMapOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsSymmetricKeyIamMemberMapOutput)
}

type KmsSymmetricKeyIamMemberOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (o KmsSymmetricKeyIamMemberOutput) ToKmsSymmetricKeyIamMemberOutput() KmsSymmetricKeyIamMemberOutput {
	return o
}

func (o KmsSymmetricKeyIamMemberOutput) ToKmsSymmetricKeyIamMemberOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberOutput {
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
func (o KmsSymmetricKeyIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
func (o KmsSymmetricKeyIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o KmsSymmetricKeyIamMemberOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamMember) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Symmetric Key ID to apply a binding to.
func (o KmsSymmetricKeyIamMemberOutput) SymmetricKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsSymmetricKeyIamMember) pulumi.StringOutput { return v.SymmetricKeyId }).(pulumi.StringOutput)
}

type KmsSymmetricKeyIamMemberArrayOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (o KmsSymmetricKeyIamMemberArrayOutput) ToKmsSymmetricKeyIamMemberArrayOutput() KmsSymmetricKeyIamMemberArrayOutput {
	return o
}

func (o KmsSymmetricKeyIamMemberArrayOutput) ToKmsSymmetricKeyIamMemberArrayOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberArrayOutput {
	return o
}

func (o KmsSymmetricKeyIamMemberArrayOutput) Index(i pulumi.IntInput) KmsSymmetricKeyIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsSymmetricKeyIamMember {
		return vs[0].([]*KmsSymmetricKeyIamMember)[vs[1].(int)]
	}).(KmsSymmetricKeyIamMemberOutput)
}

type KmsSymmetricKeyIamMemberMapOutput struct{ *pulumi.OutputState }

func (KmsSymmetricKeyIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsSymmetricKeyIamMember)(nil)).Elem()
}

func (o KmsSymmetricKeyIamMemberMapOutput) ToKmsSymmetricKeyIamMemberMapOutput() KmsSymmetricKeyIamMemberMapOutput {
	return o
}

func (o KmsSymmetricKeyIamMemberMapOutput) ToKmsSymmetricKeyIamMemberMapOutputWithContext(ctx context.Context) KmsSymmetricKeyIamMemberMapOutput {
	return o
}

func (o KmsSymmetricKeyIamMemberMapOutput) MapIndex(k pulumi.StringInput) KmsSymmetricKeyIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsSymmetricKeyIamMember {
		return vs[0].(map[string]*KmsSymmetricKeyIamMember)[vs[1].(string)]
	}).(KmsSymmetricKeyIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamMemberInput)(nil)).Elem(), &KmsSymmetricKeyIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamMemberArrayInput)(nil)).Elem(), KmsSymmetricKeyIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsSymmetricKeyIamMemberMapInput)(nil)).Elem(), KmsSymmetricKeyIamMemberMap{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamMemberOutput{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamMemberArrayOutput{})
	pulumi.RegisterOutputType(KmsSymmetricKeyIamMemberMapOutput{})
}

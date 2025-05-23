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

type KmsAsymmetricEncryptionKeyIamMember struct {
	pulumi.CustomResourceState

	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringOutput `pulumi:"asymmetricEncryptionKeyId"`
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
}

// NewKmsAsymmetricEncryptionKeyIamMember registers a new resource with the given unique name, arguments, and options.
func NewKmsAsymmetricEncryptionKeyIamMember(ctx *pulumi.Context,
	name string, args *KmsAsymmetricEncryptionKeyIamMemberArgs, opts ...pulumi.ResourceOption) (*KmsAsymmetricEncryptionKeyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AsymmetricEncryptionKeyId == nil {
		return nil, errors.New("invalid value for required argument 'AsymmetricEncryptionKeyId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KmsAsymmetricEncryptionKeyIamMember
	err := ctx.RegisterResource("yandex:index/kmsAsymmetricEncryptionKeyIamMember:KmsAsymmetricEncryptionKeyIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsAsymmetricEncryptionKeyIamMember gets an existing KmsAsymmetricEncryptionKeyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsAsymmetricEncryptionKeyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsAsymmetricEncryptionKeyIamMemberState, opts ...pulumi.ResourceOption) (*KmsAsymmetricEncryptionKeyIamMember, error) {
	var resource KmsAsymmetricEncryptionKeyIamMember
	err := ctx.ReadResource("yandex:index/kmsAsymmetricEncryptionKeyIamMember:KmsAsymmetricEncryptionKeyIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsAsymmetricEncryptionKeyIamMember resources.
type kmsAsymmetricEncryptionKeyIamMemberState struct {
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId *string `pulumi:"asymmetricEncryptionKeyId"`
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
}

type KmsAsymmetricEncryptionKeyIamMemberState struct {
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringPtrInput
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
}

func (KmsAsymmetricEncryptionKeyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymmetricEncryptionKeyIamMemberState)(nil)).Elem()
}

type kmsAsymmetricEncryptionKeyIamMemberArgs struct {
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId string `pulumi:"asymmetricEncryptionKeyId"`
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
}

// The set of arguments for constructing a KmsAsymmetricEncryptionKeyIamMember resource.
type KmsAsymmetricEncryptionKeyIamMemberArgs struct {
	// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringInput
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
}

func (KmsAsymmetricEncryptionKeyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymmetricEncryptionKeyIamMemberArgs)(nil)).Elem()
}

type KmsAsymmetricEncryptionKeyIamMemberInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamMemberOutput() KmsAsymmetricEncryptionKeyIamMemberOutput
	ToKmsAsymmetricEncryptionKeyIamMemberOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberOutput
}

func (*KmsAsymmetricEncryptionKeyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (i *KmsAsymmetricEncryptionKeyIamMember) ToKmsAsymmetricEncryptionKeyIamMemberOutput() KmsAsymmetricEncryptionKeyIamMemberOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamMemberOutputWithContext(context.Background())
}

func (i *KmsAsymmetricEncryptionKeyIamMember) ToKmsAsymmetricEncryptionKeyIamMemberOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamMemberOutput)
}

// KmsAsymmetricEncryptionKeyIamMemberArrayInput is an input type that accepts KmsAsymmetricEncryptionKeyIamMemberArray and KmsAsymmetricEncryptionKeyIamMemberArrayOutput values.
// You can construct a concrete instance of `KmsAsymmetricEncryptionKeyIamMemberArrayInput` via:
//
//	KmsAsymmetricEncryptionKeyIamMemberArray{ KmsAsymmetricEncryptionKeyIamMemberArgs{...} }
type KmsAsymmetricEncryptionKeyIamMemberArrayInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamMemberArrayOutput() KmsAsymmetricEncryptionKeyIamMemberArrayOutput
	ToKmsAsymmetricEncryptionKeyIamMemberArrayOutputWithContext(context.Context) KmsAsymmetricEncryptionKeyIamMemberArrayOutput
}

type KmsAsymmetricEncryptionKeyIamMemberArray []KmsAsymmetricEncryptionKeyIamMemberInput

func (KmsAsymmetricEncryptionKeyIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (i KmsAsymmetricEncryptionKeyIamMemberArray) ToKmsAsymmetricEncryptionKeyIamMemberArrayOutput() KmsAsymmetricEncryptionKeyIamMemberArrayOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamMemberArrayOutputWithContext(context.Background())
}

func (i KmsAsymmetricEncryptionKeyIamMemberArray) ToKmsAsymmetricEncryptionKeyIamMemberArrayOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamMemberArrayOutput)
}

// KmsAsymmetricEncryptionKeyIamMemberMapInput is an input type that accepts KmsAsymmetricEncryptionKeyIamMemberMap and KmsAsymmetricEncryptionKeyIamMemberMapOutput values.
// You can construct a concrete instance of `KmsAsymmetricEncryptionKeyIamMemberMapInput` via:
//
//	KmsAsymmetricEncryptionKeyIamMemberMap{ "key": KmsAsymmetricEncryptionKeyIamMemberArgs{...} }
type KmsAsymmetricEncryptionKeyIamMemberMapInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamMemberMapOutput() KmsAsymmetricEncryptionKeyIamMemberMapOutput
	ToKmsAsymmetricEncryptionKeyIamMemberMapOutputWithContext(context.Context) KmsAsymmetricEncryptionKeyIamMemberMapOutput
}

type KmsAsymmetricEncryptionKeyIamMemberMap map[string]KmsAsymmetricEncryptionKeyIamMemberInput

func (KmsAsymmetricEncryptionKeyIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (i KmsAsymmetricEncryptionKeyIamMemberMap) ToKmsAsymmetricEncryptionKeyIamMemberMapOutput() KmsAsymmetricEncryptionKeyIamMemberMapOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamMemberMapOutputWithContext(context.Background())
}

func (i KmsAsymmetricEncryptionKeyIamMemberMap) ToKmsAsymmetricEncryptionKeyIamMemberMapOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamMemberMapOutput)
}

type KmsAsymmetricEncryptionKeyIamMemberOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamMemberOutput) ToKmsAsymmetricEncryptionKeyIamMemberOutput() KmsAsymmetricEncryptionKeyIamMemberOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamMemberOutput) ToKmsAsymmetricEncryptionKeyIamMemberOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberOutput {
	return o
}

// The [Yandex Key Management Service](https://yandex.cloud/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
func (o KmsAsymmetricEncryptionKeyIamMemberOutput) AsymmetricEncryptionKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamMember) pulumi.StringOutput { return v.AsymmetricEncryptionKeyId }).(pulumi.StringOutput)
}

// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
func (o KmsAsymmetricEncryptionKeyIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
func (o KmsAsymmetricEncryptionKeyIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o KmsAsymmetricEncryptionKeyIamMemberOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamMember) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

type KmsAsymmetricEncryptionKeyIamMemberArrayOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamMemberArrayOutput) ToKmsAsymmetricEncryptionKeyIamMemberArrayOutput() KmsAsymmetricEncryptionKeyIamMemberArrayOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamMemberArrayOutput) ToKmsAsymmetricEncryptionKeyIamMemberArrayOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberArrayOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamMemberArrayOutput) Index(i pulumi.IntInput) KmsAsymmetricEncryptionKeyIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsAsymmetricEncryptionKeyIamMember {
		return vs[0].([]*KmsAsymmetricEncryptionKeyIamMember)[vs[1].(int)]
	}).(KmsAsymmetricEncryptionKeyIamMemberOutput)
}

type KmsAsymmetricEncryptionKeyIamMemberMapOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymmetricEncryptionKeyIamMember)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamMemberMapOutput) ToKmsAsymmetricEncryptionKeyIamMemberMapOutput() KmsAsymmetricEncryptionKeyIamMemberMapOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamMemberMapOutput) ToKmsAsymmetricEncryptionKeyIamMemberMapOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamMemberMapOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamMemberMapOutput) MapIndex(k pulumi.StringInput) KmsAsymmetricEncryptionKeyIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsAsymmetricEncryptionKeyIamMember {
		return vs[0].(map[string]*KmsAsymmetricEncryptionKeyIamMember)[vs[1].(string)]
	}).(KmsAsymmetricEncryptionKeyIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamMemberInput)(nil)).Elem(), &KmsAsymmetricEncryptionKeyIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamMemberArrayInput)(nil)).Elem(), KmsAsymmetricEncryptionKeyIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamMemberMapInput)(nil)).Elem(), KmsAsymmetricEncryptionKeyIamMemberMap{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamMemberOutput{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamMemberArrayOutput{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamMemberMapOutput{})
}

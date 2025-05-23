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

type OrganizationmanagerOrganizationIamMember struct {
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
	// The ID of the organization to attach the policy to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewOrganizationmanagerOrganizationIamMember registers a new resource with the given unique name, arguments, and options.
func NewOrganizationmanagerOrganizationIamMember(ctx *pulumi.Context,
	name string, args *OrganizationmanagerOrganizationIamMemberArgs, opts ...pulumi.ResourceOption) (*OrganizationmanagerOrganizationIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationmanagerOrganizationIamMember
	err := ctx.RegisterResource("yandex:index/organizationmanagerOrganizationIamMember:OrganizationmanagerOrganizationIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationmanagerOrganizationIamMember gets an existing OrganizationmanagerOrganizationIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationmanagerOrganizationIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationmanagerOrganizationIamMemberState, opts ...pulumi.ResourceOption) (*OrganizationmanagerOrganizationIamMember, error) {
	var resource OrganizationmanagerOrganizationIamMember
	err := ctx.ReadResource("yandex:index/organizationmanagerOrganizationIamMember:OrganizationmanagerOrganizationIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationmanagerOrganizationIamMember resources.
type organizationmanagerOrganizationIamMemberState struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member *string `pulumi:"member"`
	// The ID of the organization to attach the policy to.
	OrganizationId *string `pulumi:"organizationId"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type OrganizationmanagerOrganizationIamMemberState struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member pulumi.StringPtrInput
	// The ID of the organization to attach the policy to.
	OrganizationId pulumi.StringPtrInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationmanagerOrganizationIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationmanagerOrganizationIamMemberState)(nil)).Elem()
}

type organizationmanagerOrganizationIamMemberArgs struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member string `pulumi:"member"`
	// The ID of the organization to attach the policy to.
	OrganizationId string `pulumi:"organizationId"`
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a OrganizationmanagerOrganizationIamMember resource.
type OrganizationmanagerOrganizationIamMemberArgs struct {
	// An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
	// values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
	// **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
	// federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
	// **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
	// **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
	// All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
	// system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
	Member pulumi.StringInput
	// The ID of the organization to attach the policy to.
	OrganizationId pulumi.StringInput
	// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (OrganizationmanagerOrganizationIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationmanagerOrganizationIamMemberArgs)(nil)).Elem()
}

type OrganizationmanagerOrganizationIamMemberInput interface {
	pulumi.Input

	ToOrganizationmanagerOrganizationIamMemberOutput() OrganizationmanagerOrganizationIamMemberOutput
	ToOrganizationmanagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberOutput
}

func (*OrganizationmanagerOrganizationIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (i *OrganizationmanagerOrganizationIamMember) ToOrganizationmanagerOrganizationIamMemberOutput() OrganizationmanagerOrganizationIamMemberOutput {
	return i.ToOrganizationmanagerOrganizationIamMemberOutputWithContext(context.Background())
}

func (i *OrganizationmanagerOrganizationIamMember) ToOrganizationmanagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOrganizationIamMemberOutput)
}

// OrganizationmanagerOrganizationIamMemberArrayInput is an input type that accepts OrganizationmanagerOrganizationIamMemberArray and OrganizationmanagerOrganizationIamMemberArrayOutput values.
// You can construct a concrete instance of `OrganizationmanagerOrganizationIamMemberArrayInput` via:
//
//	OrganizationmanagerOrganizationIamMemberArray{ OrganizationmanagerOrganizationIamMemberArgs{...} }
type OrganizationmanagerOrganizationIamMemberArrayInput interface {
	pulumi.Input

	ToOrganizationmanagerOrganizationIamMemberArrayOutput() OrganizationmanagerOrganizationIamMemberArrayOutput
	ToOrganizationmanagerOrganizationIamMemberArrayOutputWithContext(context.Context) OrganizationmanagerOrganizationIamMemberArrayOutput
}

type OrganizationmanagerOrganizationIamMemberArray []OrganizationmanagerOrganizationIamMemberInput

func (OrganizationmanagerOrganizationIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (i OrganizationmanagerOrganizationIamMemberArray) ToOrganizationmanagerOrganizationIamMemberArrayOutput() OrganizationmanagerOrganizationIamMemberArrayOutput {
	return i.ToOrganizationmanagerOrganizationIamMemberArrayOutputWithContext(context.Background())
}

func (i OrganizationmanagerOrganizationIamMemberArray) ToOrganizationmanagerOrganizationIamMemberArrayOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOrganizationIamMemberArrayOutput)
}

// OrganizationmanagerOrganizationIamMemberMapInput is an input type that accepts OrganizationmanagerOrganizationIamMemberMap and OrganizationmanagerOrganizationIamMemberMapOutput values.
// You can construct a concrete instance of `OrganizationmanagerOrganizationIamMemberMapInput` via:
//
//	OrganizationmanagerOrganizationIamMemberMap{ "key": OrganizationmanagerOrganizationIamMemberArgs{...} }
type OrganizationmanagerOrganizationIamMemberMapInput interface {
	pulumi.Input

	ToOrganizationmanagerOrganizationIamMemberMapOutput() OrganizationmanagerOrganizationIamMemberMapOutput
	ToOrganizationmanagerOrganizationIamMemberMapOutputWithContext(context.Context) OrganizationmanagerOrganizationIamMemberMapOutput
}

type OrganizationmanagerOrganizationIamMemberMap map[string]OrganizationmanagerOrganizationIamMemberInput

func (OrganizationmanagerOrganizationIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (i OrganizationmanagerOrganizationIamMemberMap) ToOrganizationmanagerOrganizationIamMemberMapOutput() OrganizationmanagerOrganizationIamMemberMapOutput {
	return i.ToOrganizationmanagerOrganizationIamMemberMapOutputWithContext(context.Background())
}

func (i OrganizationmanagerOrganizationIamMemberMap) ToOrganizationmanagerOrganizationIamMemberMapOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOrganizationIamMemberMapOutput)
}

type OrganizationmanagerOrganizationIamMemberOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOrganizationIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationmanagerOrganizationIamMemberOutput) ToOrganizationmanagerOrganizationIamMemberOutput() OrganizationmanagerOrganizationIamMemberOutput {
	return o
}

func (o OrganizationmanagerOrganizationIamMemberOutput) ToOrganizationmanagerOrganizationIamMemberOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberOutput {
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
func (o OrganizationmanagerOrganizationIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationmanagerOrganizationIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the organization to attach the policy to.
func (o OrganizationmanagerOrganizationIamMemberOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationmanagerOrganizationIamMember) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
func (o OrganizationmanagerOrganizationIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationmanagerOrganizationIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o OrganizationmanagerOrganizationIamMemberOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OrganizationmanagerOrganizationIamMember) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

type OrganizationmanagerOrganizationIamMemberArrayOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOrganizationIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationmanagerOrganizationIamMemberArrayOutput) ToOrganizationmanagerOrganizationIamMemberArrayOutput() OrganizationmanagerOrganizationIamMemberArrayOutput {
	return o
}

func (o OrganizationmanagerOrganizationIamMemberArrayOutput) ToOrganizationmanagerOrganizationIamMemberArrayOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberArrayOutput {
	return o
}

func (o OrganizationmanagerOrganizationIamMemberArrayOutput) Index(i pulumi.IntInput) OrganizationmanagerOrganizationIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationmanagerOrganizationIamMember {
		return vs[0].([]*OrganizationmanagerOrganizationIamMember)[vs[1].(int)]
	}).(OrganizationmanagerOrganizationIamMemberOutput)
}

type OrganizationmanagerOrganizationIamMemberMapOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOrganizationIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationmanagerOrganizationIamMember)(nil)).Elem()
}

func (o OrganizationmanagerOrganizationIamMemberMapOutput) ToOrganizationmanagerOrganizationIamMemberMapOutput() OrganizationmanagerOrganizationIamMemberMapOutput {
	return o
}

func (o OrganizationmanagerOrganizationIamMemberMapOutput) ToOrganizationmanagerOrganizationIamMemberMapOutputWithContext(ctx context.Context) OrganizationmanagerOrganizationIamMemberMapOutput {
	return o
}

func (o OrganizationmanagerOrganizationIamMemberMapOutput) MapIndex(k pulumi.StringInput) OrganizationmanagerOrganizationIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationmanagerOrganizationIamMember {
		return vs[0].(map[string]*OrganizationmanagerOrganizationIamMember)[vs[1].(string)]
	}).(OrganizationmanagerOrganizationIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOrganizationIamMemberInput)(nil)).Elem(), &OrganizationmanagerOrganizationIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOrganizationIamMemberArrayInput)(nil)).Elem(), OrganizationmanagerOrganizationIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOrganizationIamMemberMapInput)(nil)).Elem(), OrganizationmanagerOrganizationIamMemberMap{})
	pulumi.RegisterOutputType(OrganizationmanagerOrganizationIamMemberOutput{})
	pulumi.RegisterOutputType(OrganizationmanagerOrganizationIamMemberArrayOutput{})
	pulumi.RegisterOutputType(OrganizationmanagerOrganizationIamMemberMapOutput{})
}

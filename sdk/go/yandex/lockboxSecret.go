// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LockboxSecret struct {
	pulumi.CustomResourceState

	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Payload specification for password generation.
	PasswordPayloadSpecification LockboxSecretPasswordPayloadSpecificationPtrOutput `pulumi:"passwordPayloadSpecification"`
	// The Yandex Cloud Lockbox secret status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewLockboxSecret registers a new resource with the given unique name, arguments, and options.
func NewLockboxSecret(ctx *pulumi.Context,
	name string, args *LockboxSecretArgs, opts ...pulumi.ResourceOption) (*LockboxSecret, error) {
	if args == nil {
		args = &LockboxSecretArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LockboxSecret
	err := ctx.RegisterResource("yandex:index/lockboxSecret:LockboxSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLockboxSecret gets an existing LockboxSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLockboxSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LockboxSecretState, opts ...pulumi.ResourceOption) (*LockboxSecret, error) {
	var resource LockboxSecret
	err := ctx.ReadResource("yandex:index/lockboxSecret:LockboxSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LockboxSecret resources.
type lockboxSecretState struct {
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Payload specification for password generation.
	PasswordPayloadSpecification *LockboxSecretPasswordPayloadSpecification `pulumi:"passwordPayloadSpecification"`
	// The Yandex Cloud Lockbox secret status.
	Status *string `pulumi:"status"`
}

type LockboxSecretState struct {
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KmsKeyId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Payload specification for password generation.
	PasswordPayloadSpecification LockboxSecretPasswordPayloadSpecificationPtrInput
	// The Yandex Cloud Lockbox secret status.
	Status pulumi.StringPtrInput
}

func (LockboxSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*lockboxSecretState)(nil)).Elem()
}

type lockboxSecretArgs struct {
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Payload specification for password generation.
	PasswordPayloadSpecification *LockboxSecretPasswordPayloadSpecification `pulumi:"passwordPayloadSpecification"`
}

// The set of arguments for constructing a LockboxSecret resource.
type LockboxSecretArgs struct {
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KmsKeyId pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Payload specification for password generation.
	PasswordPayloadSpecification LockboxSecretPasswordPayloadSpecificationPtrInput
}

func (LockboxSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lockboxSecretArgs)(nil)).Elem()
}

type LockboxSecretInput interface {
	pulumi.Input

	ToLockboxSecretOutput() LockboxSecretOutput
	ToLockboxSecretOutputWithContext(ctx context.Context) LockboxSecretOutput
}

func (*LockboxSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**LockboxSecret)(nil)).Elem()
}

func (i *LockboxSecret) ToLockboxSecretOutput() LockboxSecretOutput {
	return i.ToLockboxSecretOutputWithContext(context.Background())
}

func (i *LockboxSecret) ToLockboxSecretOutputWithContext(ctx context.Context) LockboxSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockboxSecretOutput)
}

// LockboxSecretArrayInput is an input type that accepts LockboxSecretArray and LockboxSecretArrayOutput values.
// You can construct a concrete instance of `LockboxSecretArrayInput` via:
//
//	LockboxSecretArray{ LockboxSecretArgs{...} }
type LockboxSecretArrayInput interface {
	pulumi.Input

	ToLockboxSecretArrayOutput() LockboxSecretArrayOutput
	ToLockboxSecretArrayOutputWithContext(context.Context) LockboxSecretArrayOutput
}

type LockboxSecretArray []LockboxSecretInput

func (LockboxSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LockboxSecret)(nil)).Elem()
}

func (i LockboxSecretArray) ToLockboxSecretArrayOutput() LockboxSecretArrayOutput {
	return i.ToLockboxSecretArrayOutputWithContext(context.Background())
}

func (i LockboxSecretArray) ToLockboxSecretArrayOutputWithContext(ctx context.Context) LockboxSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockboxSecretArrayOutput)
}

// LockboxSecretMapInput is an input type that accepts LockboxSecretMap and LockboxSecretMapOutput values.
// You can construct a concrete instance of `LockboxSecretMapInput` via:
//
//	LockboxSecretMap{ "key": LockboxSecretArgs{...} }
type LockboxSecretMapInput interface {
	pulumi.Input

	ToLockboxSecretMapOutput() LockboxSecretMapOutput
	ToLockboxSecretMapOutputWithContext(context.Context) LockboxSecretMapOutput
}

type LockboxSecretMap map[string]LockboxSecretInput

func (LockboxSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LockboxSecret)(nil)).Elem()
}

func (i LockboxSecretMap) ToLockboxSecretMapOutput() LockboxSecretMapOutput {
	return i.ToLockboxSecretMapOutputWithContext(context.Background())
}

func (i LockboxSecretMap) ToLockboxSecretMapOutputWithContext(ctx context.Context) LockboxSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LockboxSecretMapOutput)
}

type LockboxSecretOutput struct{ *pulumi.OutputState }

func (LockboxSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LockboxSecret)(nil)).Elem()
}

func (o LockboxSecretOutput) ToLockboxSecretOutput() LockboxSecretOutput {
	return o
}

func (o LockboxSecretOutput) ToLockboxSecretOutputWithContext(ctx context.Context) LockboxSecretOutput {
	return o
}

// The creation timestamp of the resource.
func (o LockboxSecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The `true` value means that resource is protected from accidental deletion.
func (o LockboxSecretOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// The resource description.
func (o LockboxSecretOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o LockboxSecretOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
func (o LockboxSecretOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o LockboxSecretOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name.
func (o LockboxSecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Payload specification for password generation.
func (o LockboxSecretOutput) PasswordPayloadSpecification() LockboxSecretPasswordPayloadSpecificationPtrOutput {
	return o.ApplyT(func(v *LockboxSecret) LockboxSecretPasswordPayloadSpecificationPtrOutput {
		return v.PasswordPayloadSpecification
	}).(LockboxSecretPasswordPayloadSpecificationPtrOutput)
}

// The Yandex Cloud Lockbox secret status.
func (o LockboxSecretOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LockboxSecret) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type LockboxSecretArrayOutput struct{ *pulumi.OutputState }

func (LockboxSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LockboxSecret)(nil)).Elem()
}

func (o LockboxSecretArrayOutput) ToLockboxSecretArrayOutput() LockboxSecretArrayOutput {
	return o
}

func (o LockboxSecretArrayOutput) ToLockboxSecretArrayOutputWithContext(ctx context.Context) LockboxSecretArrayOutput {
	return o
}

func (o LockboxSecretArrayOutput) Index(i pulumi.IntInput) LockboxSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LockboxSecret {
		return vs[0].([]*LockboxSecret)[vs[1].(int)]
	}).(LockboxSecretOutput)
}

type LockboxSecretMapOutput struct{ *pulumi.OutputState }

func (LockboxSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LockboxSecret)(nil)).Elem()
}

func (o LockboxSecretMapOutput) ToLockboxSecretMapOutput() LockboxSecretMapOutput {
	return o
}

func (o LockboxSecretMapOutput) ToLockboxSecretMapOutputWithContext(ctx context.Context) LockboxSecretMapOutput {
	return o
}

func (o LockboxSecretMapOutput) MapIndex(k pulumi.StringInput) LockboxSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LockboxSecret {
		return vs[0].(map[string]*LockboxSecret)[vs[1].(string)]
	}).(LockboxSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LockboxSecretInput)(nil)).Elem(), &LockboxSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*LockboxSecretArrayInput)(nil)).Elem(), LockboxSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LockboxSecretMapInput)(nil)).Elem(), LockboxSecretMap{})
	pulumi.RegisterOutputType(LockboxSecretOutput{})
	pulumi.RegisterOutputType(LockboxSecretArrayOutput{})
	pulumi.RegisterOutputType(LockboxSecretMapOutput{})
}

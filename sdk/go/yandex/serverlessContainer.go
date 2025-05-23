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

type ServerlessContainer struct {
	pulumi.CustomResourceState

	// Concurrency of Yandex Cloud Serverless Container.
	Concurrency pulumi.IntOutput `pulumi:"concurrency"`
	// Network access. If specified the revision will be attached to specified network.
	Connectivity ServerlessContainerConnectivityPtrOutput `pulumi:"connectivity"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
	CoreFraction pulumi.IntOutput `pulumi:"coreFraction"`
	// Cores (**1+**) of the Yandex Cloud Serverless Container.
	Cores pulumi.IntPtrOutput `pulumi:"cores"`
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
	ExecutionTimeout pulumi.StringOutput `pulumi:"executionTimeout"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container.
	Image ServerlessContainerImageOutput `pulumi:"image"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Options for logging from Yandex Cloud Serverless Container.
	LogOptions ServerlessContainerLogOptionsPtrOutput `pulumi:"logOptions"`
	// Memory in megabytes (**aligned to 128 MB**).
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Options set the access mode to revision's metadata endpoints.
	MetadataOptions ServerlessContainerMetadataOptionsOutput `pulumi:"metadataOptions"`
	// Mounts for Yandex Cloud Serverless Container.
	Mounts ServerlessContainerMountArrayOutput `pulumi:"mounts"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provision policy. If specified the revision will have prepared instances.
	ProvisionPolicy ServerlessContainerProvisionPolicyPtrOutput `pulumi:"provisionPolicy"`
	// Last revision ID of the Yandex Cloud Serverless Container.
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Runtime for Yandex Cloud Serverless Container.
	Runtime ServerlessContainerRuntimeOutput `pulumi:"runtime"`
	// Secrets for Yandex Cloud Serverless Container.
	Secrets ServerlessContainerSecretArrayOutput `pulumi:"secrets"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrOutput `pulumi:"serviceAccountId"`
	// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
	//
	// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
	StorageMounts ServerlessContainerStorageMountArrayOutput `pulumi:"storageMounts"`
	// Invoke URL for the Yandex Cloud Serverless Container.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewServerlessContainer registers a new resource with the given unique name, arguments, and options.
func NewServerlessContainer(ctx *pulumi.Context,
	name string, args *ServerlessContainerArgs, opts ...pulumi.ResourceOption) (*ServerlessContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerlessContainer
	err := ctx.RegisterResource("yandex:index/serverlessContainer:ServerlessContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessContainer gets an existing ServerlessContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessContainerState, opts ...pulumi.ResourceOption) (*ServerlessContainer, error) {
	var resource ServerlessContainer
	err := ctx.ReadResource("yandex:index/serverlessContainer:ServerlessContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessContainer resources.
type serverlessContainerState struct {
	// Concurrency of Yandex Cloud Serverless Container.
	Concurrency *int `pulumi:"concurrency"`
	// Network access. If specified the revision will be attached to specified network.
	Connectivity *ServerlessContainerConnectivity `pulumi:"connectivity"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
	CoreFraction *int `pulumi:"coreFraction"`
	// Cores (**1+**) of the Yandex Cloud Serverless Container.
	Cores *int `pulumi:"cores"`
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
	ExecutionTimeout *string `pulumi:"executionTimeout"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container.
	Image *ServerlessContainerImage `pulumi:"image"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Options for logging from Yandex Cloud Serverless Container.
	LogOptions *ServerlessContainerLogOptions `pulumi:"logOptions"`
	// Memory in megabytes (**aligned to 128 MB**).
	Memory *int `pulumi:"memory"`
	// Options set the access mode to revision's metadata endpoints.
	MetadataOptions *ServerlessContainerMetadataOptions `pulumi:"metadataOptions"`
	// Mounts for Yandex Cloud Serverless Container.
	Mounts []ServerlessContainerMount `pulumi:"mounts"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Provision policy. If specified the revision will have prepared instances.
	ProvisionPolicy *ServerlessContainerProvisionPolicy `pulumi:"provisionPolicy"`
	// Last revision ID of the Yandex Cloud Serverless Container.
	RevisionId *string `pulumi:"revisionId"`
	// Runtime for Yandex Cloud Serverless Container.
	Runtime *ServerlessContainerRuntime `pulumi:"runtime"`
	// Secrets for Yandex Cloud Serverless Container.
	Secrets []ServerlessContainerSecret `pulumi:"secrets"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
	//
	// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
	StorageMounts []ServerlessContainerStorageMount `pulumi:"storageMounts"`
	// Invoke URL for the Yandex Cloud Serverless Container.
	Url *string `pulumi:"url"`
}

type ServerlessContainerState struct {
	// Concurrency of Yandex Cloud Serverless Container.
	Concurrency pulumi.IntPtrInput
	// Network access. If specified the revision will be attached to specified network.
	Connectivity ServerlessContainerConnectivityPtrInput
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
	CoreFraction pulumi.IntPtrInput
	// Cores (**1+**) of the Yandex Cloud Serverless Container.
	Cores pulumi.IntPtrInput
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
	ExecutionTimeout pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// Revision deployment image for Yandex Cloud Serverless Container.
	Image ServerlessContainerImagePtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Options for logging from Yandex Cloud Serverless Container.
	LogOptions ServerlessContainerLogOptionsPtrInput
	// Memory in megabytes (**aligned to 128 MB**).
	Memory pulumi.IntPtrInput
	// Options set the access mode to revision's metadata endpoints.
	MetadataOptions ServerlessContainerMetadataOptionsPtrInput
	// Mounts for Yandex Cloud Serverless Container.
	Mounts ServerlessContainerMountArrayInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Provision policy. If specified the revision will have prepared instances.
	ProvisionPolicy ServerlessContainerProvisionPolicyPtrInput
	// Last revision ID of the Yandex Cloud Serverless Container.
	RevisionId pulumi.StringPtrInput
	// Runtime for Yandex Cloud Serverless Container.
	Runtime ServerlessContainerRuntimePtrInput
	// Secrets for Yandex Cloud Serverless Container.
	Secrets ServerlessContainerSecretArrayInput
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrInput
	// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
	//
	// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
	StorageMounts ServerlessContainerStorageMountArrayInput
	// Invoke URL for the Yandex Cloud Serverless Container.
	Url pulumi.StringPtrInput
}

func (ServerlessContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessContainerState)(nil)).Elem()
}

type serverlessContainerArgs struct {
	// Concurrency of Yandex Cloud Serverless Container.
	Concurrency *int `pulumi:"concurrency"`
	// Network access. If specified the revision will be attached to specified network.
	Connectivity *ServerlessContainerConnectivity `pulumi:"connectivity"`
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
	CoreFraction *int `pulumi:"coreFraction"`
	// Cores (**1+**) of the Yandex Cloud Serverless Container.
	Cores *int `pulumi:"cores"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
	ExecutionTimeout *string `pulumi:"executionTimeout"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// Revision deployment image for Yandex Cloud Serverless Container.
	Image ServerlessContainerImage `pulumi:"image"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Options for logging from Yandex Cloud Serverless Container.
	LogOptions *ServerlessContainerLogOptions `pulumi:"logOptions"`
	// Memory in megabytes (**aligned to 128 MB**).
	Memory int `pulumi:"memory"`
	// Options set the access mode to revision's metadata endpoints.
	MetadataOptions *ServerlessContainerMetadataOptions `pulumi:"metadataOptions"`
	// Mounts for Yandex Cloud Serverless Container.
	Mounts []ServerlessContainerMount `pulumi:"mounts"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Provision policy. If specified the revision will have prepared instances.
	ProvisionPolicy *ServerlessContainerProvisionPolicy `pulumi:"provisionPolicy"`
	// Runtime for Yandex Cloud Serverless Container.
	Runtime *ServerlessContainerRuntime `pulumi:"runtime"`
	// Secrets for Yandex Cloud Serverless Container.
	Secrets []ServerlessContainerSecret `pulumi:"secrets"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
	//
	// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
	StorageMounts []ServerlessContainerStorageMount `pulumi:"storageMounts"`
}

// The set of arguments for constructing a ServerlessContainer resource.
type ServerlessContainerArgs struct {
	// Concurrency of Yandex Cloud Serverless Container.
	Concurrency pulumi.IntPtrInput
	// Network access. If specified the revision will be attached to specified network.
	Connectivity ServerlessContainerConnectivityPtrInput
	// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
	CoreFraction pulumi.IntPtrInput
	// Cores (**1+**) of the Yandex Cloud Serverless Container.
	Cores pulumi.IntPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
	ExecutionTimeout pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// Revision deployment image for Yandex Cloud Serverless Container.
	Image ServerlessContainerImageInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Options for logging from Yandex Cloud Serverless Container.
	LogOptions ServerlessContainerLogOptionsPtrInput
	// Memory in megabytes (**aligned to 128 MB**).
	Memory pulumi.IntInput
	// Options set the access mode to revision's metadata endpoints.
	MetadataOptions ServerlessContainerMetadataOptionsPtrInput
	// Mounts for Yandex Cloud Serverless Container.
	Mounts ServerlessContainerMountArrayInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Provision policy. If specified the revision will have prepared instances.
	ProvisionPolicy ServerlessContainerProvisionPolicyPtrInput
	// Runtime for Yandex Cloud Serverless Container.
	Runtime ServerlessContainerRuntimePtrInput
	// Secrets for Yandex Cloud Serverless Container.
	Secrets ServerlessContainerSecretArrayInput
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrInput
	// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
	//
	// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
	StorageMounts ServerlessContainerStorageMountArrayInput
}

func (ServerlessContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessContainerArgs)(nil)).Elem()
}

type ServerlessContainerInput interface {
	pulumi.Input

	ToServerlessContainerOutput() ServerlessContainerOutput
	ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput
}

func (*ServerlessContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessContainer)(nil)).Elem()
}

func (i *ServerlessContainer) ToServerlessContainerOutput() ServerlessContainerOutput {
	return i.ToServerlessContainerOutputWithContext(context.Background())
}

func (i *ServerlessContainer) ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerOutput)
}

// ServerlessContainerArrayInput is an input type that accepts ServerlessContainerArray and ServerlessContainerArrayOutput values.
// You can construct a concrete instance of `ServerlessContainerArrayInput` via:
//
//	ServerlessContainerArray{ ServerlessContainerArgs{...} }
type ServerlessContainerArrayInput interface {
	pulumi.Input

	ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput
	ToServerlessContainerArrayOutputWithContext(context.Context) ServerlessContainerArrayOutput
}

type ServerlessContainerArray []ServerlessContainerInput

func (ServerlessContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessContainer)(nil)).Elem()
}

func (i ServerlessContainerArray) ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput {
	return i.ToServerlessContainerArrayOutputWithContext(context.Background())
}

func (i ServerlessContainerArray) ToServerlessContainerArrayOutputWithContext(ctx context.Context) ServerlessContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerArrayOutput)
}

// ServerlessContainerMapInput is an input type that accepts ServerlessContainerMap and ServerlessContainerMapOutput values.
// You can construct a concrete instance of `ServerlessContainerMapInput` via:
//
//	ServerlessContainerMap{ "key": ServerlessContainerArgs{...} }
type ServerlessContainerMapInput interface {
	pulumi.Input

	ToServerlessContainerMapOutput() ServerlessContainerMapOutput
	ToServerlessContainerMapOutputWithContext(context.Context) ServerlessContainerMapOutput
}

type ServerlessContainerMap map[string]ServerlessContainerInput

func (ServerlessContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessContainer)(nil)).Elem()
}

func (i ServerlessContainerMap) ToServerlessContainerMapOutput() ServerlessContainerMapOutput {
	return i.ToServerlessContainerMapOutputWithContext(context.Background())
}

func (i ServerlessContainerMap) ToServerlessContainerMapOutputWithContext(ctx context.Context) ServerlessContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessContainerMapOutput)
}

type ServerlessContainerOutput struct{ *pulumi.OutputState }

func (ServerlessContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerOutput) ToServerlessContainerOutput() ServerlessContainerOutput {
	return o
}

func (o ServerlessContainerOutput) ToServerlessContainerOutputWithContext(ctx context.Context) ServerlessContainerOutput {
	return o
}

// Concurrency of Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Concurrency() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.IntOutput { return v.Concurrency }).(pulumi.IntOutput)
}

// Network access. If specified the revision will be attached to specified network.
func (o ServerlessContainerOutput) Connectivity() ServerlessContainerConnectivityPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerConnectivityPtrOutput { return v.Connectivity }).(ServerlessContainerConnectivityPtrOutput)
}

// Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) CoreFraction() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.IntOutput { return v.CoreFraction }).(pulumi.IntOutput)
}

// Cores (**1+**) of the Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.IntPtrOutput { return v.Cores }).(pulumi.IntPtrOutput)
}

// The creation timestamp of the resource.
func (o ServerlessContainerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The resource description.
func (o ServerlessContainerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) ExecutionTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.ExecutionTimeout }).(pulumi.StringOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o ServerlessContainerOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Revision deployment image for Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Image() ServerlessContainerImageOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerImageOutput { return v.Image }).(ServerlessContainerImageOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o ServerlessContainerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Options for logging from Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) LogOptions() ServerlessContainerLogOptionsPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerLogOptionsPtrOutput { return v.LogOptions }).(ServerlessContainerLogOptionsPtrOutput)
}

// Memory in megabytes (**aligned to 128 MB**).
func (o ServerlessContainerOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// Options set the access mode to revision's metadata endpoints.
func (o ServerlessContainerOutput) MetadataOptions() ServerlessContainerMetadataOptionsOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerMetadataOptionsOutput { return v.MetadataOptions }).(ServerlessContainerMetadataOptionsOutput)
}

// Mounts for Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Mounts() ServerlessContainerMountArrayOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerMountArrayOutput { return v.Mounts }).(ServerlessContainerMountArrayOutput)
}

// The resource name.
func (o ServerlessContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provision policy. If specified the revision will have prepared instances.
func (o ServerlessContainerOutput) ProvisionPolicy() ServerlessContainerProvisionPolicyPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerProvisionPolicyPtrOutput { return v.ProvisionPolicy }).(ServerlessContainerProvisionPolicyPtrOutput)
}

// Last revision ID of the Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Runtime for Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Runtime() ServerlessContainerRuntimeOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerRuntimeOutput { return v.Runtime }).(ServerlessContainerRuntimeOutput)
}

// Secrets for Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Secrets() ServerlessContainerSecretArrayOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerSecretArrayOutput { return v.Secrets }).(ServerlessContainerSecretArrayOutput)
}

// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
func (o ServerlessContainerOutput) ServiceAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringPtrOutput { return v.ServiceAccountId }).(pulumi.StringPtrOutput)
}

// (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
//
// Deprecated: The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
func (o ServerlessContainerOutput) StorageMounts() ServerlessContainerStorageMountArrayOutput {
	return o.ApplyT(func(v *ServerlessContainer) ServerlessContainerStorageMountArrayOutput { return v.StorageMounts }).(ServerlessContainerStorageMountArrayOutput)
}

// Invoke URL for the Yandex Cloud Serverless Container.
func (o ServerlessContainerOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessContainer) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ServerlessContainerArrayOutput struct{ *pulumi.OutputState }

func (ServerlessContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerArrayOutput) ToServerlessContainerArrayOutput() ServerlessContainerArrayOutput {
	return o
}

func (o ServerlessContainerArrayOutput) ToServerlessContainerArrayOutputWithContext(ctx context.Context) ServerlessContainerArrayOutput {
	return o
}

func (o ServerlessContainerArrayOutput) Index(i pulumi.IntInput) ServerlessContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerlessContainer {
		return vs[0].([]*ServerlessContainer)[vs[1].(int)]
	}).(ServerlessContainerOutput)
}

type ServerlessContainerMapOutput struct{ *pulumi.OutputState }

func (ServerlessContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessContainer)(nil)).Elem()
}

func (o ServerlessContainerMapOutput) ToServerlessContainerMapOutput() ServerlessContainerMapOutput {
	return o
}

func (o ServerlessContainerMapOutput) ToServerlessContainerMapOutputWithContext(ctx context.Context) ServerlessContainerMapOutput {
	return o
}

func (o ServerlessContainerMapOutput) MapIndex(k pulumi.StringInput) ServerlessContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerlessContainer {
		return vs[0].(map[string]*ServerlessContainer)[vs[1].(string)]
	}).(ServerlessContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerInput)(nil)).Elem(), &ServerlessContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerArrayInput)(nil)).Elem(), ServerlessContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessContainerMapInput)(nil)).Elem(), ServerlessContainerMap{})
	pulumi.RegisterOutputType(ServerlessContainerOutput{})
	pulumi.RegisterOutputType(ServerlessContainerArrayOutput{})
	pulumi.RegisterOutputType(ServerlessContainerMapOutput{})
}

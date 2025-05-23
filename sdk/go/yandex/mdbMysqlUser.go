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

type MdbMysqlUser struct {
	pulumi.CustomResourceState

	// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
	// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin pulumi.StringOutput `pulumi:"authenticationPlugin"`
	// The ID of the MySQL cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
	// parameters are set to `-1`, backend default values will be actually used.
	ConnectionLimits MdbMysqlUserConnectionLimitsOutput `pulumi:"connectionLimits"`
	// Connection Manager connection configuration. Filled in by the server automatically.
	ConnectionManager pulumi.StringMapOutput `pulumi:"connectionManager"`
	// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
	// is ignored during updating. > **Must specify either password or generate_password**.
	GeneratePassword pulumi.BoolPtrOutput `pulumi:"generatePassword"`
	// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
	// use empty list. If the attribute is not specified there will be no changes.
	GlobalPermissions pulumi.StringArrayOutput `pulumi:"globalPermissions"`
	// The name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of the user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Set of permissions granted to the user.
	Permissions MdbMysqlUserPermissionArrayOutput `pulumi:"permissions"`
}

// NewMdbMysqlUser registers a new resource with the given unique name, arguments, and options.
func NewMdbMysqlUser(ctx *pulumi.Context,
	name string, args *MdbMysqlUserArgs, opts ...pulumi.ResourceOption) (*MdbMysqlUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MdbMysqlUser
	err := ctx.RegisterResource("yandex:index/mdbMysqlUser:MdbMysqlUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbMysqlUser gets an existing MdbMysqlUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbMysqlUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbMysqlUserState, opts ...pulumi.ResourceOption) (*MdbMysqlUser, error) {
	var resource MdbMysqlUser
	err := ctx.ReadResource("yandex:index/mdbMysqlUser:MdbMysqlUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbMysqlUser resources.
type mdbMysqlUserState struct {
	// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
	// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin *string `pulumi:"authenticationPlugin"`
	// The ID of the MySQL cluster.
	ClusterId *string `pulumi:"clusterId"`
	// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
	// parameters are set to `-1`, backend default values will be actually used.
	ConnectionLimits *MdbMysqlUserConnectionLimits `pulumi:"connectionLimits"`
	// Connection Manager connection configuration. Filled in by the server automatically.
	ConnectionManager map[string]string `pulumi:"connectionManager"`
	// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
	// is ignored during updating. > **Must specify either password or generate_password**.
	GeneratePassword *bool `pulumi:"generatePassword"`
	// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
	// use empty list. If the attribute is not specified there will be no changes.
	GlobalPermissions []string `pulumi:"globalPermissions"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user.
	Password *string `pulumi:"password"`
	// Set of permissions granted to the user.
	Permissions []MdbMysqlUserPermission `pulumi:"permissions"`
}

type MdbMysqlUserState struct {
	// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
	// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin pulumi.StringPtrInput
	// The ID of the MySQL cluster.
	ClusterId pulumi.StringPtrInput
	// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
	// parameters are set to `-1`, backend default values will be actually used.
	ConnectionLimits MdbMysqlUserConnectionLimitsPtrInput
	// Connection Manager connection configuration. Filled in by the server automatically.
	ConnectionManager pulumi.StringMapInput
	// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
	// is ignored during updating. > **Must specify either password or generate_password**.
	GeneratePassword pulumi.BoolPtrInput
	// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
	// use empty list. If the attribute is not specified there will be no changes.
	GlobalPermissions pulumi.StringArrayInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user.
	Password pulumi.StringPtrInput
	// Set of permissions granted to the user.
	Permissions MdbMysqlUserPermissionArrayInput
}

func (MdbMysqlUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMysqlUserState)(nil)).Elem()
}

type mdbMysqlUserArgs struct {
	// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
	// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin *string `pulumi:"authenticationPlugin"`
	// The ID of the MySQL cluster.
	ClusterId string `pulumi:"clusterId"`
	// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
	// parameters are set to `-1`, backend default values will be actually used.
	ConnectionLimits *MdbMysqlUserConnectionLimits `pulumi:"connectionLimits"`
	// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
	// is ignored during updating. > **Must specify either password or generate_password**.
	GeneratePassword *bool `pulumi:"generatePassword"`
	// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
	// use empty list. If the attribute is not specified there will be no changes.
	GlobalPermissions []string `pulumi:"globalPermissions"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user.
	Password *string `pulumi:"password"`
	// Set of permissions granted to the user.
	Permissions []MdbMysqlUserPermission `pulumi:"permissions"`
}

// The set of arguments for constructing a MdbMysqlUser resource.
type MdbMysqlUserArgs struct {
	// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
	// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin pulumi.StringPtrInput
	// The ID of the MySQL cluster.
	ClusterId pulumi.StringInput
	// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
	// parameters are set to `-1`, backend default values will be actually used.
	ConnectionLimits MdbMysqlUserConnectionLimitsPtrInput
	// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
	// is ignored during updating. > **Must specify either password or generate_password**.
	GeneratePassword pulumi.BoolPtrInput
	// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
	// use empty list. If the attribute is not specified there will be no changes.
	GlobalPermissions pulumi.StringArrayInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user.
	Password pulumi.StringPtrInput
	// Set of permissions granted to the user.
	Permissions MdbMysqlUserPermissionArrayInput
}

func (MdbMysqlUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMysqlUserArgs)(nil)).Elem()
}

type MdbMysqlUserInput interface {
	pulumi.Input

	ToMdbMysqlUserOutput() MdbMysqlUserOutput
	ToMdbMysqlUserOutputWithContext(ctx context.Context) MdbMysqlUserOutput
}

func (*MdbMysqlUser) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMysqlUser)(nil)).Elem()
}

func (i *MdbMysqlUser) ToMdbMysqlUserOutput() MdbMysqlUserOutput {
	return i.ToMdbMysqlUserOutputWithContext(context.Background())
}

func (i *MdbMysqlUser) ToMdbMysqlUserOutputWithContext(ctx context.Context) MdbMysqlUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlUserOutput)
}

// MdbMysqlUserArrayInput is an input type that accepts MdbMysqlUserArray and MdbMysqlUserArrayOutput values.
// You can construct a concrete instance of `MdbMysqlUserArrayInput` via:
//
//	MdbMysqlUserArray{ MdbMysqlUserArgs{...} }
type MdbMysqlUserArrayInput interface {
	pulumi.Input

	ToMdbMysqlUserArrayOutput() MdbMysqlUserArrayOutput
	ToMdbMysqlUserArrayOutputWithContext(context.Context) MdbMysqlUserArrayOutput
}

type MdbMysqlUserArray []MdbMysqlUserInput

func (MdbMysqlUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMysqlUser)(nil)).Elem()
}

func (i MdbMysqlUserArray) ToMdbMysqlUserArrayOutput() MdbMysqlUserArrayOutput {
	return i.ToMdbMysqlUserArrayOutputWithContext(context.Background())
}

func (i MdbMysqlUserArray) ToMdbMysqlUserArrayOutputWithContext(ctx context.Context) MdbMysqlUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlUserArrayOutput)
}

// MdbMysqlUserMapInput is an input type that accepts MdbMysqlUserMap and MdbMysqlUserMapOutput values.
// You can construct a concrete instance of `MdbMysqlUserMapInput` via:
//
//	MdbMysqlUserMap{ "key": MdbMysqlUserArgs{...} }
type MdbMysqlUserMapInput interface {
	pulumi.Input

	ToMdbMysqlUserMapOutput() MdbMysqlUserMapOutput
	ToMdbMysqlUserMapOutputWithContext(context.Context) MdbMysqlUserMapOutput
}

type MdbMysqlUserMap map[string]MdbMysqlUserInput

func (MdbMysqlUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMysqlUser)(nil)).Elem()
}

func (i MdbMysqlUserMap) ToMdbMysqlUserMapOutput() MdbMysqlUserMapOutput {
	return i.ToMdbMysqlUserMapOutputWithContext(context.Background())
}

func (i MdbMysqlUserMap) ToMdbMysqlUserMapOutputWithContext(ctx context.Context) MdbMysqlUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMysqlUserMapOutput)
}

type MdbMysqlUserOutput struct{ *pulumi.OutputState }

func (MdbMysqlUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMysqlUser)(nil)).Elem()
}

func (o MdbMysqlUserOutput) ToMdbMysqlUserOutput() MdbMysqlUserOutput {
	return o
}

func (o MdbMysqlUserOutput) ToMdbMysqlUserOutputWithContext(ctx context.Context) MdbMysqlUserOutput {
	return o
}

// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version
// 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
func (o MdbMysqlUserOutput) AuthenticationPlugin() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringOutput { return v.AuthenticationPlugin }).(pulumi.StringOutput)
}

// The ID of the MySQL cluster.
func (o MdbMysqlUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// User's connection limits. If the attribute is not specified there will be no changes. Default value is `-1`. When these
// parameters are set to `-1`, backend default values will be actually used.
func (o MdbMysqlUserOutput) ConnectionLimits() MdbMysqlUserConnectionLimitsOutput {
	return o.ApplyT(func(v *MdbMysqlUser) MdbMysqlUserConnectionLimitsOutput { return v.ConnectionLimits }).(MdbMysqlUserConnectionLimitsOutput)
}

// Connection Manager connection configuration. Filled in by the server automatically.
func (o MdbMysqlUserOutput) ConnectionManager() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringMapOutput { return v.ConnectionManager }).(pulumi.StringMapOutput)
}

// Generate password using Connection Manager. Allowed values: `true` or `false`. It's used only during user creation and
// is ignored during updating. > **Must specify either password or generate_password**.
func (o MdbMysqlUserOutput) GeneratePassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.BoolPtrOutput { return v.GeneratePassword }).(pulumi.BoolPtrOutput)
}

// List user's global permissions. Allowed permissions: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` for clear list
// use empty list. If the attribute is not specified there will be no changes.
func (o MdbMysqlUserOutput) GlobalPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringArrayOutput { return v.GlobalPermissions }).(pulumi.StringArrayOutput)
}

// The name of the user.
func (o MdbMysqlUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password of the user.
func (o MdbMysqlUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbMysqlUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Set of permissions granted to the user.
func (o MdbMysqlUserOutput) Permissions() MdbMysqlUserPermissionArrayOutput {
	return o.ApplyT(func(v *MdbMysqlUser) MdbMysqlUserPermissionArrayOutput { return v.Permissions }).(MdbMysqlUserPermissionArrayOutput)
}

type MdbMysqlUserArrayOutput struct{ *pulumi.OutputState }

func (MdbMysqlUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMysqlUser)(nil)).Elem()
}

func (o MdbMysqlUserArrayOutput) ToMdbMysqlUserArrayOutput() MdbMysqlUserArrayOutput {
	return o
}

func (o MdbMysqlUserArrayOutput) ToMdbMysqlUserArrayOutputWithContext(ctx context.Context) MdbMysqlUserArrayOutput {
	return o
}

func (o MdbMysqlUserArrayOutput) Index(i pulumi.IntInput) MdbMysqlUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbMysqlUser {
		return vs[0].([]*MdbMysqlUser)[vs[1].(int)]
	}).(MdbMysqlUserOutput)
}

type MdbMysqlUserMapOutput struct{ *pulumi.OutputState }

func (MdbMysqlUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMysqlUser)(nil)).Elem()
}

func (o MdbMysqlUserMapOutput) ToMdbMysqlUserMapOutput() MdbMysqlUserMapOutput {
	return o
}

func (o MdbMysqlUserMapOutput) ToMdbMysqlUserMapOutputWithContext(ctx context.Context) MdbMysqlUserMapOutput {
	return o
}

func (o MdbMysqlUserMapOutput) MapIndex(k pulumi.StringInput) MdbMysqlUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbMysqlUser {
		return vs[0].(map[string]*MdbMysqlUser)[vs[1].(string)]
	}).(MdbMysqlUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlUserInput)(nil)).Elem(), &MdbMysqlUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlUserArrayInput)(nil)).Elem(), MdbMysqlUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMysqlUserMapInput)(nil)).Elem(), MdbMysqlUserMap{})
	pulumi.RegisterOutputType(MdbMysqlUserOutput{})
	pulumi.RegisterOutputType(MdbMysqlUserArrayOutput{})
	pulumi.RegisterOutputType(MdbMysqlUserMapOutput{})
}

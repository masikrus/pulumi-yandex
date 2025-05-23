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

type OrganizationmanagerOsLoginSettings struct {
	pulumi.CustomResourceState

	// The organization to manage it's OsLogin Settings.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// SSH Certificate settings.
	SshCertificateSettings OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrOutput `pulumi:"sshCertificateSettings"`
	// Users SSH key settings.
	UserSshKeySettings OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrOutput `pulumi:"userSshKeySettings"`
}

// NewOrganizationmanagerOsLoginSettings registers a new resource with the given unique name, arguments, and options.
func NewOrganizationmanagerOsLoginSettings(ctx *pulumi.Context,
	name string, args *OrganizationmanagerOsLoginSettingsArgs, opts ...pulumi.ResourceOption) (*OrganizationmanagerOsLoginSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationmanagerOsLoginSettings
	err := ctx.RegisterResource("yandex:index/organizationmanagerOsLoginSettings:OrganizationmanagerOsLoginSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationmanagerOsLoginSettings gets an existing OrganizationmanagerOsLoginSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationmanagerOsLoginSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationmanagerOsLoginSettingsState, opts ...pulumi.ResourceOption) (*OrganizationmanagerOsLoginSettings, error) {
	var resource OrganizationmanagerOsLoginSettings
	err := ctx.ReadResource("yandex:index/organizationmanagerOsLoginSettings:OrganizationmanagerOsLoginSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationmanagerOsLoginSettings resources.
type organizationmanagerOsLoginSettingsState struct {
	// The organization to manage it's OsLogin Settings.
	OrganizationId *string `pulumi:"organizationId"`
	// SSH Certificate settings.
	SshCertificateSettings *OrganizationmanagerOsLoginSettingsSshCertificateSettings `pulumi:"sshCertificateSettings"`
	// Users SSH key settings.
	UserSshKeySettings *OrganizationmanagerOsLoginSettingsUserSshKeySettings `pulumi:"userSshKeySettings"`
}

type OrganizationmanagerOsLoginSettingsState struct {
	// The organization to manage it's OsLogin Settings.
	OrganizationId pulumi.StringPtrInput
	// SSH Certificate settings.
	SshCertificateSettings OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrInput
	// Users SSH key settings.
	UserSshKeySettings OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrInput
}

func (OrganizationmanagerOsLoginSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationmanagerOsLoginSettingsState)(nil)).Elem()
}

type organizationmanagerOsLoginSettingsArgs struct {
	// The organization to manage it's OsLogin Settings.
	OrganizationId string `pulumi:"organizationId"`
	// SSH Certificate settings.
	SshCertificateSettings *OrganizationmanagerOsLoginSettingsSshCertificateSettings `pulumi:"sshCertificateSettings"`
	// Users SSH key settings.
	UserSshKeySettings *OrganizationmanagerOsLoginSettingsUserSshKeySettings `pulumi:"userSshKeySettings"`
}

// The set of arguments for constructing a OrganizationmanagerOsLoginSettings resource.
type OrganizationmanagerOsLoginSettingsArgs struct {
	// The organization to manage it's OsLogin Settings.
	OrganizationId pulumi.StringInput
	// SSH Certificate settings.
	SshCertificateSettings OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrInput
	// Users SSH key settings.
	UserSshKeySettings OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrInput
}

func (OrganizationmanagerOsLoginSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationmanagerOsLoginSettingsArgs)(nil)).Elem()
}

type OrganizationmanagerOsLoginSettingsInput interface {
	pulumi.Input

	ToOrganizationmanagerOsLoginSettingsOutput() OrganizationmanagerOsLoginSettingsOutput
	ToOrganizationmanagerOsLoginSettingsOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsOutput
}

func (*OrganizationmanagerOsLoginSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (i *OrganizationmanagerOsLoginSettings) ToOrganizationmanagerOsLoginSettingsOutput() OrganizationmanagerOsLoginSettingsOutput {
	return i.ToOrganizationmanagerOsLoginSettingsOutputWithContext(context.Background())
}

func (i *OrganizationmanagerOsLoginSettings) ToOrganizationmanagerOsLoginSettingsOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOsLoginSettingsOutput)
}

// OrganizationmanagerOsLoginSettingsArrayInput is an input type that accepts OrganizationmanagerOsLoginSettingsArray and OrganizationmanagerOsLoginSettingsArrayOutput values.
// You can construct a concrete instance of `OrganizationmanagerOsLoginSettingsArrayInput` via:
//
//	OrganizationmanagerOsLoginSettingsArray{ OrganizationmanagerOsLoginSettingsArgs{...} }
type OrganizationmanagerOsLoginSettingsArrayInput interface {
	pulumi.Input

	ToOrganizationmanagerOsLoginSettingsArrayOutput() OrganizationmanagerOsLoginSettingsArrayOutput
	ToOrganizationmanagerOsLoginSettingsArrayOutputWithContext(context.Context) OrganizationmanagerOsLoginSettingsArrayOutput
}

type OrganizationmanagerOsLoginSettingsArray []OrganizationmanagerOsLoginSettingsInput

func (OrganizationmanagerOsLoginSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (i OrganizationmanagerOsLoginSettingsArray) ToOrganizationmanagerOsLoginSettingsArrayOutput() OrganizationmanagerOsLoginSettingsArrayOutput {
	return i.ToOrganizationmanagerOsLoginSettingsArrayOutputWithContext(context.Background())
}

func (i OrganizationmanagerOsLoginSettingsArray) ToOrganizationmanagerOsLoginSettingsArrayOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOsLoginSettingsArrayOutput)
}

// OrganizationmanagerOsLoginSettingsMapInput is an input type that accepts OrganizationmanagerOsLoginSettingsMap and OrganizationmanagerOsLoginSettingsMapOutput values.
// You can construct a concrete instance of `OrganizationmanagerOsLoginSettingsMapInput` via:
//
//	OrganizationmanagerOsLoginSettingsMap{ "key": OrganizationmanagerOsLoginSettingsArgs{...} }
type OrganizationmanagerOsLoginSettingsMapInput interface {
	pulumi.Input

	ToOrganizationmanagerOsLoginSettingsMapOutput() OrganizationmanagerOsLoginSettingsMapOutput
	ToOrganizationmanagerOsLoginSettingsMapOutputWithContext(context.Context) OrganizationmanagerOsLoginSettingsMapOutput
}

type OrganizationmanagerOsLoginSettingsMap map[string]OrganizationmanagerOsLoginSettingsInput

func (OrganizationmanagerOsLoginSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (i OrganizationmanagerOsLoginSettingsMap) ToOrganizationmanagerOsLoginSettingsMapOutput() OrganizationmanagerOsLoginSettingsMapOutput {
	return i.ToOrganizationmanagerOsLoginSettingsMapOutputWithContext(context.Background())
}

func (i OrganizationmanagerOsLoginSettingsMap) ToOrganizationmanagerOsLoginSettingsMapOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationmanagerOsLoginSettingsMapOutput)
}

type OrganizationmanagerOsLoginSettingsOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOsLoginSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (o OrganizationmanagerOsLoginSettingsOutput) ToOrganizationmanagerOsLoginSettingsOutput() OrganizationmanagerOsLoginSettingsOutput {
	return o
}

func (o OrganizationmanagerOsLoginSettingsOutput) ToOrganizationmanagerOsLoginSettingsOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsOutput {
	return o
}

// The organization to manage it's OsLogin Settings.
func (o OrganizationmanagerOsLoginSettingsOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationmanagerOsLoginSettings) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// SSH Certificate settings.
func (o OrganizationmanagerOsLoginSettingsOutput) SshCertificateSettings() OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrOutput {
	return o.ApplyT(func(v *OrganizationmanagerOsLoginSettings) OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrOutput {
		return v.SshCertificateSettings
	}).(OrganizationmanagerOsLoginSettingsSshCertificateSettingsPtrOutput)
}

// Users SSH key settings.
func (o OrganizationmanagerOsLoginSettingsOutput) UserSshKeySettings() OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrOutput {
	return o.ApplyT(func(v *OrganizationmanagerOsLoginSettings) OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrOutput {
		return v.UserSshKeySettings
	}).(OrganizationmanagerOsLoginSettingsUserSshKeySettingsPtrOutput)
}

type OrganizationmanagerOsLoginSettingsArrayOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOsLoginSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (o OrganizationmanagerOsLoginSettingsArrayOutput) ToOrganizationmanagerOsLoginSettingsArrayOutput() OrganizationmanagerOsLoginSettingsArrayOutput {
	return o
}

func (o OrganizationmanagerOsLoginSettingsArrayOutput) ToOrganizationmanagerOsLoginSettingsArrayOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsArrayOutput {
	return o
}

func (o OrganizationmanagerOsLoginSettingsArrayOutput) Index(i pulumi.IntInput) OrganizationmanagerOsLoginSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationmanagerOsLoginSettings {
		return vs[0].([]*OrganizationmanagerOsLoginSettings)[vs[1].(int)]
	}).(OrganizationmanagerOsLoginSettingsOutput)
}

type OrganizationmanagerOsLoginSettingsMapOutput struct{ *pulumi.OutputState }

func (OrganizationmanagerOsLoginSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationmanagerOsLoginSettings)(nil)).Elem()
}

func (o OrganizationmanagerOsLoginSettingsMapOutput) ToOrganizationmanagerOsLoginSettingsMapOutput() OrganizationmanagerOsLoginSettingsMapOutput {
	return o
}

func (o OrganizationmanagerOsLoginSettingsMapOutput) ToOrganizationmanagerOsLoginSettingsMapOutputWithContext(ctx context.Context) OrganizationmanagerOsLoginSettingsMapOutput {
	return o
}

func (o OrganizationmanagerOsLoginSettingsMapOutput) MapIndex(k pulumi.StringInput) OrganizationmanagerOsLoginSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationmanagerOsLoginSettings {
		return vs[0].(map[string]*OrganizationmanagerOsLoginSettings)[vs[1].(string)]
	}).(OrganizationmanagerOsLoginSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOsLoginSettingsInput)(nil)).Elem(), &OrganizationmanagerOsLoginSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOsLoginSettingsArrayInput)(nil)).Elem(), OrganizationmanagerOsLoginSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationmanagerOsLoginSettingsMapInput)(nil)).Elem(), OrganizationmanagerOsLoginSettingsMap{})
	pulumi.RegisterOutputType(OrganizationmanagerOsLoginSettingsOutput{})
	pulumi.RegisterOutputType(OrganizationmanagerOsLoginSettingsArrayOutput{})
	pulumi.RegisterOutputType(OrganizationmanagerOsLoginSettingsMapOutput{})
}

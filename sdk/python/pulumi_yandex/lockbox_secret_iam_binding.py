# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['LockboxSecretIamBindingArgs', 'LockboxSecretIamBinding']

@pulumi.input_type
class LockboxSecretIamBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 role: pulumi.Input[builtins.str],
                 secret_id: pulumi.Input[builtins.str],
                 sleep_after: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a LockboxSecretIamBinding resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
               values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
               **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
               federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
               **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
               **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
               All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
               system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        :param pulumi.Input[builtins.str] role: The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        :param pulumi.Input[builtins.str] secret_id: The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "secret_id", secret_id)
        if sleep_after is not None:
            pulumi.set(__self__, "sleep_after", sleep_after)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
        values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
        **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
        federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
        **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
        **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
        All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
        system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[builtins.str]:
        """
        The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[builtins.str]:
        """
        The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="sleepAfter")
    def sleep_after(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "sleep_after")

    @sleep_after.setter
    def sleep_after(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep_after", value)


@pulumi.input_type
class _LockboxSecretIamBindingState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None,
                 sleep_after: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering LockboxSecretIamBinding resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
               values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
               **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
               federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
               **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
               **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
               All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
               system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        :param pulumi.Input[builtins.str] role: The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        :param pulumi.Input[builtins.str] secret_id: The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if sleep_after is not None:
            pulumi.set(__self__, "sleep_after", sleep_after)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
        values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
        **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
        federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
        **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
        **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
        All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
        system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="sleepAfter")
    def sleep_after(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "sleep_after")

    @sleep_after.setter
    def sleep_after(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "sleep_after", value)


@pulumi.type_token("yandex:index/lockboxSecretIamBinding:LockboxSecretIamBinding")
class LockboxSecretIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None,
                 sleep_after: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Create a LockboxSecretIamBinding resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
               values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
               **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
               federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
               **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
               **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
               All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
               system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        :param pulumi.Input[builtins.str] role: The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        :param pulumi.Input[builtins.str] secret_id: The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LockboxSecretIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LockboxSecretIamBinding resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LockboxSecretIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LockboxSecretIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None,
                 sleep_after: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LockboxSecretIamBindingArgs.__new__(LockboxSecretIamBindingArgs)

            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__.__dict__["secret_id"] = secret_id
            __props__.__dict__["sleep_after"] = sleep_after
        super(LockboxSecretIamBinding, __self__).__init__(
            'yandex:index/lockboxSecretIamBinding:LockboxSecretIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            role: Optional[pulumi.Input[builtins.str]] = None,
            secret_id: Optional[pulumi.Input[builtins.str]] = None,
            sleep_after: Optional[pulumi.Input[builtins.int]] = None) -> 'LockboxSecretIamBinding':
        """
        Get an existing LockboxSecretIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
               values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
               **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
               federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
               **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
               **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
               All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
               system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        :param pulumi.Input[builtins.str] role: The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        :param pulumi.Input[builtins.str] secret_id: The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LockboxSecretIamBindingState.__new__(_LockboxSecretIamBindingState)

        __props__.__dict__["members"] = members
        __props__.__dict__["role"] = role
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["sleep_after"] = sleep_after
        return LockboxSecretIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following
        values: * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account. *
        **serviceAccount:{service_account_id}**: A unique service account ID. * **federatedUser:{federated_user_id}**: A unique
        federated user ID. * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID. *
        **group:{group_id}**: A unique group ID. * **system:group:federation:{federation_id}:users**: All users in federation. *
        **system:group:organization:{organization_id}:users**: All users in organization. * **system:allAuthenticatedUsers**:
        All authenticated users. * **system:allUsers**: All users, including unauthenticated ones. > for more information about
        system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[builtins.str]:
        """
        The [Yandex Lockbox Secret](https://yandex.cloud/docs/lockbox/) Secret ID to apply a binding to.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="sleepAfter")
    def sleep_after(self) -> pulumi.Output[Optional[builtins.int]]:
        return pulumi.get(self, "sleep_after")


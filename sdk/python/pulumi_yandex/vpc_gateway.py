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
from . import outputs
from ._inputs import *

__all__ = ['VpcGatewayArgs', 'VpcGateway']

@pulumi.input_type
class VpcGatewayArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_egress_gateway: Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']] = None):
        """
        The set of arguments for constructing a VpcGateway resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input['VpcGatewaySharedEgressGatewayArgs'] shared_egress_gateway: Shared egress gateway configuration. Currently empty.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shared_egress_gateway is not None:
            pulumi.set(__self__, "shared_egress_gateway", shared_egress_gateway)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sharedEgressGateway")
    def shared_egress_gateway(self) -> Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']]:
        """
        Shared egress gateway configuration. Currently empty.
        """
        return pulumi.get(self, "shared_egress_gateway")

    @shared_egress_gateway.setter
    def shared_egress_gateway(self, value: Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']]):
        pulumi.set(self, "shared_egress_gateway", value)


@pulumi.input_type
class _VpcGatewayState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_egress_gateway: Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']] = None):
        """
        Input properties used for looking up and filtering VpcGateway resources.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input['VpcGatewaySharedEgressGatewayArgs'] shared_egress_gateway: Shared egress gateway configuration. Currently empty.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shared_egress_gateway is not None:
            pulumi.set(__self__, "shared_egress_gateway", shared_egress_gateway)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The creation timestamp of the resource.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sharedEgressGateway")
    def shared_egress_gateway(self) -> Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']]:
        """
        Shared egress gateway configuration. Currently empty.
        """
        return pulumi.get(self, "shared_egress_gateway")

    @shared_egress_gateway.setter
    def shared_egress_gateway(self, value: Optional[pulumi.Input['VpcGatewaySharedEgressGatewayArgs']]):
        pulumi.set(self, "shared_egress_gateway", value)


@pulumi.type_token("yandex:index/vpcGateway:VpcGateway")
class VpcGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_egress_gateway: Optional[pulumi.Input[Union['VpcGatewaySharedEgressGatewayArgs', 'VpcGatewaySharedEgressGatewayArgsDict']]] = None,
                 __props__=None):
        """
        Create a VpcGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[Union['VpcGatewaySharedEgressGatewayArgs', 'VpcGatewaySharedEgressGatewayArgsDict']] shared_egress_gateway: Shared egress gateway configuration. Currently empty.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VpcGatewayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpcGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 shared_egress_gateway: Optional[pulumi.Input[Union['VpcGatewaySharedEgressGatewayArgs', 'VpcGatewaySharedEgressGatewayArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcGatewayArgs.__new__(VpcGatewayArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["shared_egress_gateway"] = shared_egress_gateway
            __props__.__dict__["created_at"] = None
        super(VpcGateway, __self__).__init__(
            'yandex:index/vpcGateway:VpcGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            folder_id: Optional[pulumi.Input[builtins.str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            shared_egress_gateway: Optional[pulumi.Input[Union['VpcGatewaySharedEgressGatewayArgs', 'VpcGatewaySharedEgressGatewayArgsDict']]] = None) -> 'VpcGateway':
        """
        Get an existing VpcGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[Union['VpcGatewaySharedEgressGatewayArgs', 'VpcGatewaySharedEgressGatewayArgsDict']] shared_egress_gateway: Shared egress gateway configuration. Currently empty.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcGatewayState.__new__(_VpcGatewayState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["shared_egress_gateway"] = shared_egress_gateway
        return VpcGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The creation timestamp of the resource.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[builtins.str]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sharedEgressGateway")
    def shared_egress_gateway(self) -> pulumi.Output[Optional['outputs.VpcGatewaySharedEgressGateway']]:
        """
        Shared egress gateway configuration. Currently empty.
        """
        return pulumi.get(self, "shared_egress_gateway")


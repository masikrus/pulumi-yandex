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

__all__ = ['VpcSubnetArgs', 'VpcSubnet']

@pulumi.input_type
class VpcSubnetArgs:
    def __init__(__self__, *,
                 network_id: pulumi.Input[builtins.str],
                 v4_cidr_blocks: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dhcp_options: Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a VpcSubnet resource.
        :param pulumi.Input[builtins.str] network_id: ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v4_cidr_blocks: A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
               subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
               network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input['VpcSubnetDhcpOptionsArgs'] dhcp_options: Options for DHCP client.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
               subnet.
        :param pulumi.Input[builtins.str] zone: The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
               provided, the default provider zone will be used.
        """
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "v4_cidr_blocks", v4_cidr_blocks)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dhcp_options is not None:
            pulumi.set(__self__, "dhcp_options", dhcp_options)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="v4CidrBlocks")
    def v4_cidr_blocks(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
        subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
        network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        """
        return pulumi.get(self, "v4_cidr_blocks")

    @v4_cidr_blocks.setter
    def v4_cidr_blocks(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "v4_cidr_blocks", value)

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
    @pulumi.getter(name="dhcpOptions")
    def dhcp_options(self) -> Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']]:
        """
        Options for DHCP client.
        """
        return pulumi.get(self, "dhcp_options")

    @dhcp_options.setter
    def dhcp_options(self, value: Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']]):
        pulumi.set(self, "dhcp_options", value)

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
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
        subnet.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        provided, the default provider zone will be used.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _VpcSubnetState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dhcp_options: Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 v4_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 v6_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering VpcSubnet resources.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input['VpcSubnetDhcpOptionsArgs'] dhcp_options: Options for DHCP client.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
               subnet.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v4_cidr_blocks: A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
               subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
               network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v6_cidr_blocks: An optional list of blocks of IPv6 addresses that are owned by this subnet.
        :param pulumi.Input[builtins.str] zone: The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
               provided, the default provider zone will be used.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dhcp_options is not None:
            pulumi.set(__self__, "dhcp_options", dhcp_options)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)
        if v4_cidr_blocks is not None:
            pulumi.set(__self__, "v4_cidr_blocks", v4_cidr_blocks)
        if v6_cidr_blocks is not None:
            pulumi.set(__self__, "v6_cidr_blocks", v6_cidr_blocks)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

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
    @pulumi.getter(name="dhcpOptions")
    def dhcp_options(self) -> Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']]:
        """
        Options for DHCP client.
        """
        return pulumi.get(self, "dhcp_options")

    @dhcp_options.setter
    def dhcp_options(self, value: Optional[pulumi.Input['VpcSubnetDhcpOptionsArgs']]):
        pulumi.set(self, "dhcp_options", value)

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
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
        subnet.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="v4CidrBlocks")
    def v4_cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
        subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
        network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        """
        return pulumi.get(self, "v4_cidr_blocks")

    @v4_cidr_blocks.setter
    def v4_cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "v4_cidr_blocks", value)

    @property
    @pulumi.getter(name="v6CidrBlocks")
    def v6_cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        An optional list of blocks of IPv6 addresses that are owned by this subnet.
        """
        return pulumi.get(self, "v6_cidr_blocks")

    @v6_cidr_blocks.setter
    def v6_cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "v6_cidr_blocks", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        provided, the default provider zone will be used.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.type_token("yandex:index/vpcSubnet:VpcSubnet")
class VpcSubnet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dhcp_options: Optional[pulumi.Input[Union['VpcSubnetDhcpOptionsArgs', 'VpcSubnetDhcpOptionsArgsDict']]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 v4_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a VpcSubnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Union['VpcSubnetDhcpOptionsArgs', 'VpcSubnetDhcpOptionsArgsDict']] dhcp_options: Options for DHCP client.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
               subnet.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v4_cidr_blocks: A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
               subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
               network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        :param pulumi.Input[builtins.str] zone: The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
               provided, the default provider zone will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcSubnetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpcSubnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcSubnetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcSubnetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dhcp_options: Optional[pulumi.Input[Union['VpcSubnetDhcpOptionsArgs', 'VpcSubnetDhcpOptionsArgsDict']]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 v4_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcSubnetArgs.__new__(VpcSubnetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dhcp_options"] = dhcp_options
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["route_table_id"] = route_table_id
            if v4_cidr_blocks is None and not opts.urn:
                raise TypeError("Missing required property 'v4_cidr_blocks'")
            __props__.__dict__["v4_cidr_blocks"] = v4_cidr_blocks
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["v6_cidr_blocks"] = None
        super(VpcSubnet, __self__).__init__(
            'yandex:index/vpcSubnet:VpcSubnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            dhcp_options: Optional[pulumi.Input[Union['VpcSubnetDhcpOptionsArgs', 'VpcSubnetDhcpOptionsArgsDict']]] = None,
            folder_id: Optional[pulumi.Input[builtins.str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            network_id: Optional[pulumi.Input[builtins.str]] = None,
            route_table_id: Optional[pulumi.Input[builtins.str]] = None,
            v4_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            v6_cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            zone: Optional[pulumi.Input[builtins.str]] = None) -> 'VpcSubnet':
        """
        Get an existing VpcSubnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Union['VpcSubnetDhcpOptionsArgs', 'VpcSubnetDhcpOptionsArgsDict']] dhcp_options: Options for DHCP client.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
               subnet.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v4_cidr_blocks: A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
               subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
               network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] v6_cidr_blocks: An optional list of blocks of IPv6 addresses that are owned by this subnet.
        :param pulumi.Input[builtins.str] zone: The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
               provided, the default provider zone will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcSubnetState.__new__(_VpcSubnetState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["dhcp_options"] = dhcp_options
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["route_table_id"] = route_table_id
        __props__.__dict__["v4_cidr_blocks"] = v4_cidr_blocks
        __props__.__dict__["v6_cidr_blocks"] = v6_cidr_blocks
        __props__.__dict__["zone"] = zone
        return VpcSubnet(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="dhcpOptions")
    def dhcp_options(self) -> pulumi.Output[Optional['outputs.VpcSubnetDhcpOptions']]:
        """
        Options for DHCP client.
        """
        return pulumi.get(self, "dhcp_options")

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
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this
        subnet.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="v4CidrBlocks")
    def v4_cidr_blocks(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the
        subnet. For example, `10.0.0.0/22` or `192.168.0.0/16`. Blocks of addresses must be unique and non-overlapping within a
        network. Minimum subnet size is `/28`, and maximum subnet size is `/16`. Only IPv4 is supported.
        """
        return pulumi.get(self, "v4_cidr_blocks")

    @property
    @pulumi.getter(name="v6CidrBlocks")
    def v6_cidr_blocks(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        An optional list of blocks of IPv6 addresses that are owned by this subnet.
        """
        return pulumi.get(self, "v6_cidr_blocks")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[builtins.str]:
        """
        The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        provided, the default provider zone will be used.
        """
        return pulumi.get(self, "zone")


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

__all__ = ['DnsRecordsetArgs', 'DnsRecordset']

@pulumi.input_type
class DnsRecordsetArgs:
    def __init__(__self__, *,
                 datas: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 ttl: pulumi.Input[builtins.int],
                 type: pulumi.Input[builtins.str],
                 zone_id: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DnsRecordset resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] datas: The string data for the records in this record set.
        :param pulumi.Input[builtins.int] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[builtins.str] type: The DNS record set type.
        :param pulumi.Input[builtins.str] zone_id: The id of the zone in which this record set will reside.
        :param pulumi.Input[builtins.str] name: The DNS name this record set will apply to.
        """
        pulumi.set(__self__, "datas", datas)
        pulumi.set(__self__, "ttl", ttl)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone_id", zone_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def datas(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The string data for the records in this record set.
        """
        return pulumi.get(self, "datas")

    @datas.setter
    def datas(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "datas", value)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[builtins.int]:
        """
        The time-to-live of this record set (seconds).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        The DNS record set type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[builtins.str]:
        """
        The id of the zone in which this record set will reside.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The DNS name this record set will apply to.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DnsRecordsetState:
    def __init__(__self__, *,
                 datas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 zone_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DnsRecordset resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] datas: The string data for the records in this record set.
        :param pulumi.Input[builtins.str] name: The DNS name this record set will apply to.
        :param pulumi.Input[builtins.int] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[builtins.str] type: The DNS record set type.
        :param pulumi.Input[builtins.str] zone_id: The id of the zone in which this record set will reside.
        """
        if datas is not None:
            pulumi.set(__self__, "datas", datas)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The string data for the records in this record set.
        """
        return pulumi.get(self, "datas")

    @datas.setter
    def datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "datas", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The DNS name this record set will apply to.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The time-to-live of this record set (seconds).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The DNS record set type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the zone in which this record set will reside.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.type_token("yandex:index/dnsRecordset:DnsRecordset")
class DnsRecordset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 zone_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a DnsRecordset resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] datas: The string data for the records in this record set.
        :param pulumi.Input[builtins.str] name: The DNS name this record set will apply to.
        :param pulumi.Input[builtins.int] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[builtins.str] type: The DNS record set type.
        :param pulumi.Input[builtins.str] zone_id: The id of the zone in which this record set will reside.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnsRecordsetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DnsRecordset resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DnsRecordsetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnsRecordsetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 ttl: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 zone_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DnsRecordsetArgs.__new__(DnsRecordsetArgs)

            if datas is None and not opts.urn:
                raise TypeError("Missing required property 'datas'")
            __props__.__dict__["datas"] = datas
            __props__.__dict__["name"] = name
            if ttl is None and not opts.urn:
                raise TypeError("Missing required property 'ttl'")
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(DnsRecordset, __self__).__init__(
            'yandex:index/dnsRecordset:DnsRecordset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            ttl: Optional[pulumi.Input[builtins.int]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            zone_id: Optional[pulumi.Input[builtins.str]] = None) -> 'DnsRecordset':
        """
        Get an existing DnsRecordset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] datas: The string data for the records in this record set.
        :param pulumi.Input[builtins.str] name: The DNS name this record set will apply to.
        :param pulumi.Input[builtins.int] ttl: The time-to-live of this record set (seconds).
        :param pulumi.Input[builtins.str] type: The DNS record set type.
        :param pulumi.Input[builtins.str] zone_id: The id of the zone in which this record set will reside.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnsRecordsetState.__new__(_DnsRecordsetState)

        __props__.__dict__["datas"] = datas
        __props__.__dict__["name"] = name
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["zone_id"] = zone_id
        return DnsRecordset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def datas(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The string data for the records in this record set.
        """
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The DNS name this record set will apply to.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[builtins.int]:
        """
        The time-to-live of this record set (seconds).
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The DNS record set type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[builtins.str]:
        """
        The id of the zone in which this record set will reside.
        """
        return pulumi.get(self, "zone_id")


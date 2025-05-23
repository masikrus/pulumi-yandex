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

__all__ = [
    'GetVpcAddressResult',
    'AwaitableGetVpcAddressResult',
    'get_vpc_address',
    'get_vpc_address_output',
]

@pulumi.output_type
class GetVpcAddressResult:
    """
    A collection of values returned by getVpcAddress.
    """
    def __init__(__self__, address_id=None, created_at=None, deletion_protection=None, description=None, dns_records=None, external_ipv4_addresses=None, folder_id=None, id=None, labels=None, name=None, reserved=None, used=None):
        if address_id and not isinstance(address_id, str):
            raise TypeError("Expected argument 'address_id' to be a str")
        pulumi.set(__self__, "address_id", address_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dns_records and not isinstance(dns_records, list):
            raise TypeError("Expected argument 'dns_records' to be a list")
        pulumi.set(__self__, "dns_records", dns_records)
        if external_ipv4_addresses and not isinstance(external_ipv4_addresses, list):
            raise TypeError("Expected argument 'external_ipv4_addresses' to be a list")
        pulumi.set(__self__, "external_ipv4_addresses", external_ipv4_addresses)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if reserved and not isinstance(reserved, bool):
            raise TypeError("Expected argument 'reserved' to be a bool")
        pulumi.set(__self__, "reserved", reserved)
        if used and not isinstance(used, bool):
            raise TypeError("Expected argument 'used' to be a bool")
        pulumi.set(__self__, "used", used)

    @property
    @pulumi.getter(name="addressId")
    def address_id(self) -> builtins.str:
        return pulumi.get(self, "address_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> builtins.bool:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> Sequence['outputs.GetVpcAddressDnsRecordResult']:
        return pulumi.get(self, "dns_records")

    @property
    @pulumi.getter(name="externalIpv4Addresses")
    def external_ipv4_addresses(self) -> Sequence['outputs.GetVpcAddressExternalIpv4AddressResult']:
        return pulumi.get(self, "external_ipv4_addresses")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def reserved(self) -> builtins.bool:
        return pulumi.get(self, "reserved")

    @property
    @pulumi.getter
    def used(self) -> builtins.bool:
        return pulumi.get(self, "used")


class AwaitableGetVpcAddressResult(GetVpcAddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcAddressResult(
            address_id=self.address_id,
            created_at=self.created_at,
            deletion_protection=self.deletion_protection,
            description=self.description,
            dns_records=self.dns_records,
            external_ipv4_addresses=self.external_ipv4_addresses,
            folder_id=self.folder_id,
            id=self.id,
            labels=self.labels,
            name=self.name,
            reserved=self.reserved,
            used=self.used)


def get_vpc_address(address_id: Optional[builtins.str] = None,
                    folder_id: Optional[builtins.str] = None,
                    name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcAddressResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressId'] = address_id
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getVpcAddress:getVpcAddress', __args__, opts=opts, typ=GetVpcAddressResult).value

    return AwaitableGetVpcAddressResult(
        address_id=pulumi.get(__ret__, 'address_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        description=pulumi.get(__ret__, 'description'),
        dns_records=pulumi.get(__ret__, 'dns_records'),
        external_ipv4_addresses=pulumi.get(__ret__, 'external_ipv4_addresses'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        reserved=pulumi.get(__ret__, 'reserved'),
        used=pulumi.get(__ret__, 'used'))
def get_vpc_address_output(address_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcAddressResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressId'] = address_id
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getVpcAddress:getVpcAddress', __args__, opts=opts, typ=GetVpcAddressResult)
    return __ret__.apply(lambda __response__: GetVpcAddressResult(
        address_id=pulumi.get(__response__, 'address_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        description=pulumi.get(__response__, 'description'),
        dns_records=pulumi.get(__response__, 'dns_records'),
        external_ipv4_addresses=pulumi.get(__response__, 'external_ipv4_addresses'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        reserved=pulumi.get(__response__, 'reserved'),
        used=pulumi.get(__response__, 'used')))

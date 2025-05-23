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

__all__ = [
    'GetComputeDiskResult',
    'AwaitableGetComputeDiskResult',
    'get_compute_disk',
    'get_compute_disk_output',
]

@pulumi.output_type
class GetComputeDiskResult:
    """
    A collection of values returned by getComputeDisk.
    """
    def __init__(__self__, block_size=None, created_at=None, description=None, disk_id=None, disk_placement_policy=None, folder_id=None, hardware_generations=None, id=None, image_id=None, instance_ids=None, kms_key_id=None, labels=None, name=None, product_ids=None, size=None, snapshot_id=None, status=None, type=None, zone=None):
        if block_size and not isinstance(block_size, int):
            raise TypeError("Expected argument 'block_size' to be a int")
        pulumi.set(__self__, "block_size", block_size)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_id and not isinstance(disk_id, str):
            raise TypeError("Expected argument 'disk_id' to be a str")
        pulumi.set(__self__, "disk_id", disk_id)
        if disk_placement_policy and not isinstance(disk_placement_policy, dict):
            raise TypeError("Expected argument 'disk_placement_policy' to be a dict")
        pulumi.set(__self__, "disk_placement_policy", disk_placement_policy)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if hardware_generations and not isinstance(hardware_generations, list):
            raise TypeError("Expected argument 'hardware_generations' to be a list")
        pulumi.set(__self__, "hardware_generations", hardware_generations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if product_ids and not isinstance(product_ids, list):
            raise TypeError("Expected argument 'product_ids' to be a list")
        pulumi.set(__self__, "product_ids", product_ids)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> builtins.int:
        return pulumi.get(self, "block_size")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> builtins.str:
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="diskPlacementPolicy")
    def disk_placement_policy(self) -> Optional['outputs.GetComputeDiskDiskPlacementPolicyResult']:
        return pulumi.get(self, "disk_placement_policy")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="hardwareGenerations")
    def hardware_generations(self) -> Sequence['outputs.GetComputeDiskHardwareGenerationResult']:
        return pulumi.get(self, "hardware_generations")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> builtins.str:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> builtins.str:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="productIds")
    def product_ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "product_ids")

    @property
    @pulumi.getter
    def size(self) -> builtins.int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> builtins.str:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        return pulumi.get(self, "zone")


class AwaitableGetComputeDiskResult(GetComputeDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeDiskResult(
            block_size=self.block_size,
            created_at=self.created_at,
            description=self.description,
            disk_id=self.disk_id,
            disk_placement_policy=self.disk_placement_policy,
            folder_id=self.folder_id,
            hardware_generations=self.hardware_generations,
            id=self.id,
            image_id=self.image_id,
            instance_ids=self.instance_ids,
            kms_key_id=self.kms_key_id,
            labels=self.labels,
            name=self.name,
            product_ids=self.product_ids,
            size=self.size,
            snapshot_id=self.snapshot_id,
            status=self.status,
            type=self.type,
            zone=self.zone)


def get_compute_disk(disk_id: Optional[builtins.str] = None,
                     disk_placement_policy: Optional[Union['GetComputeDiskDiskPlacementPolicyArgs', 'GetComputeDiskDiskPlacementPolicyArgsDict']] = None,
                     folder_id: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeDiskResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['diskId'] = disk_id
    __args__['diskPlacementPolicy'] = disk_placement_policy
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getComputeDisk:getComputeDisk', __args__, opts=opts, typ=GetComputeDiskResult).value

    return AwaitableGetComputeDiskResult(
        block_size=pulumi.get(__ret__, 'block_size'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        disk_id=pulumi.get(__ret__, 'disk_id'),
        disk_placement_policy=pulumi.get(__ret__, 'disk_placement_policy'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        hardware_generations=pulumi.get(__ret__, 'hardware_generations'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        instance_ids=pulumi.get(__ret__, 'instance_ids'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        product_ids=pulumi.get(__ret__, 'product_ids'),
        size=pulumi.get(__ret__, 'size'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'),
        zone=pulumi.get(__ret__, 'zone'))
def get_compute_disk_output(disk_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            disk_placement_policy: Optional[pulumi.Input[Optional[Union['GetComputeDiskDiskPlacementPolicyArgs', 'GetComputeDiskDiskPlacementPolicyArgsDict']]]] = None,
                            folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeDiskResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['diskId'] = disk_id
    __args__['diskPlacementPolicy'] = disk_placement_policy
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getComputeDisk:getComputeDisk', __args__, opts=opts, typ=GetComputeDiskResult)
    return __ret__.apply(lambda __response__: GetComputeDiskResult(
        block_size=pulumi.get(__response__, 'block_size'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        disk_id=pulumi.get(__response__, 'disk_id'),
        disk_placement_policy=pulumi.get(__response__, 'disk_placement_policy'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        hardware_generations=pulumi.get(__response__, 'hardware_generations'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        instance_ids=pulumi.get(__response__, 'instance_ids'),
        kms_key_id=pulumi.get(__response__, 'kms_key_id'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        product_ids=pulumi.get(__response__, 'product_ids'),
        size=pulumi.get(__response__, 'size'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        status=pulumi.get(__response__, 'status'),
        type=pulumi.get(__response__, 'type'),
        zone=pulumi.get(__response__, 'zone')))

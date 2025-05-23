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
    'GetComputeSnapshotResult',
    'AwaitableGetComputeSnapshotResult',
    'get_compute_snapshot',
    'get_compute_snapshot_output',
]

@pulumi.output_type
class GetComputeSnapshotResult:
    """
    A collection of values returned by getComputeSnapshot.
    """
    def __init__(__self__, created_at=None, description=None, disk_size=None, folder_id=None, hardware_generations=None, id=None, kms_key_id=None, labels=None, name=None, product_ids=None, snapshot_id=None, source_disk_id=None, status=None, storage_size=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size and not isinstance(disk_size, int):
            raise TypeError("Expected argument 'disk_size' to be a int")
        pulumi.set(__self__, "disk_size", disk_size)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if hardware_generations and not isinstance(hardware_generations, list):
            raise TypeError("Expected argument 'hardware_generations' to be a list")
        pulumi.set(__self__, "hardware_generations", hardware_generations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if source_disk_id and not isinstance(source_disk_id, str):
            raise TypeError("Expected argument 'source_disk_id' to be a str")
        pulumi.set(__self__, "source_disk_id", source_disk_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_size and not isinstance(storage_size, int):
            raise TypeError("Expected argument 'storage_size' to be a int")
        pulumi.set(__self__, "storage_size", storage_size)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> builtins.int:
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="hardwareGenerations")
    def hardware_generations(self) -> Sequence['outputs.GetComputeSnapshotHardwareGenerationResult']:
        return pulumi.get(self, "hardware_generations")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> builtins.str:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="sourceDiskId")
    def source_disk_id(self) -> builtins.str:
        return pulumi.get(self, "source_disk_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> builtins.int:
        return pulumi.get(self, "storage_size")


class AwaitableGetComputeSnapshotResult(GetComputeSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeSnapshotResult(
            created_at=self.created_at,
            description=self.description,
            disk_size=self.disk_size,
            folder_id=self.folder_id,
            hardware_generations=self.hardware_generations,
            id=self.id,
            kms_key_id=self.kms_key_id,
            labels=self.labels,
            name=self.name,
            product_ids=self.product_ids,
            snapshot_id=self.snapshot_id,
            source_disk_id=self.source_disk_id,
            status=self.status,
            storage_size=self.storage_size)


def get_compute_snapshot(folder_id: Optional[builtins.str] = None,
                         name: Optional[builtins.str] = None,
                         snapshot_id: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeSnapshotResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['snapshotId'] = snapshot_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getComputeSnapshot:getComputeSnapshot', __args__, opts=opts, typ=GetComputeSnapshotResult).value

    return AwaitableGetComputeSnapshotResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        disk_size=pulumi.get(__ret__, 'disk_size'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        hardware_generations=pulumi.get(__ret__, 'hardware_generations'),
        id=pulumi.get(__ret__, 'id'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        product_ids=pulumi.get(__ret__, 'product_ids'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        source_disk_id=pulumi.get(__ret__, 'source_disk_id'),
        status=pulumi.get(__ret__, 'status'),
        storage_size=pulumi.get(__ret__, 'storage_size'))
def get_compute_snapshot_output(folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                snapshot_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeSnapshotResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['snapshotId'] = snapshot_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getComputeSnapshot:getComputeSnapshot', __args__, opts=opts, typ=GetComputeSnapshotResult)
    return __ret__.apply(lambda __response__: GetComputeSnapshotResult(
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        disk_size=pulumi.get(__response__, 'disk_size'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        hardware_generations=pulumi.get(__response__, 'hardware_generations'),
        id=pulumi.get(__response__, 'id'),
        kms_key_id=pulumi.get(__response__, 'kms_key_id'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        product_ids=pulumi.get(__response__, 'product_ids'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        source_disk_id=pulumi.get(__response__, 'source_disk_id'),
        status=pulumi.get(__response__, 'status'),
        storage_size=pulumi.get(__response__, 'storage_size')))

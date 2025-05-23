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
    'GetComputeInstanceResult',
    'AwaitableGetComputeInstanceResult',
    'get_compute_instance',
    'get_compute_instance_output',
]

@pulumi.output_type
class GetComputeInstanceResult:
    """
    A collection of values returned by getComputeInstance.
    """
    def __init__(__self__, boot_disks=None, created_at=None, description=None, filesystems=None, folder_id=None, fqdn=None, gpu_cluster_id=None, hardware_generations=None, id=None, instance_id=None, labels=None, local_disks=None, maintenance_grace_period=None, maintenance_policy=None, metadata=None, metadata_options=None, name=None, network_acceleration_type=None, network_interfaces=None, placement_policy=None, platform_id=None, resources=None, scheduling_policies=None, secondary_disks=None, service_account_id=None, status=None, zone=None):
        if boot_disks and not isinstance(boot_disks, list):
            raise TypeError("Expected argument 'boot_disks' to be a list")
        pulumi.set(__self__, "boot_disks", boot_disks)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filesystems and not isinstance(filesystems, list):
            raise TypeError("Expected argument 'filesystems' to be a list")
        pulumi.set(__self__, "filesystems", filesystems)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if gpu_cluster_id and not isinstance(gpu_cluster_id, str):
            raise TypeError("Expected argument 'gpu_cluster_id' to be a str")
        pulumi.set(__self__, "gpu_cluster_id", gpu_cluster_id)
        if hardware_generations and not isinstance(hardware_generations, list):
            raise TypeError("Expected argument 'hardware_generations' to be a list")
        pulumi.set(__self__, "hardware_generations", hardware_generations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if local_disks and not isinstance(local_disks, list):
            raise TypeError("Expected argument 'local_disks' to be a list")
        pulumi.set(__self__, "local_disks", local_disks)
        if maintenance_grace_period and not isinstance(maintenance_grace_period, str):
            raise TypeError("Expected argument 'maintenance_grace_period' to be a str")
        pulumi.set(__self__, "maintenance_grace_period", maintenance_grace_period)
        if maintenance_policy and not isinstance(maintenance_policy, str):
            raise TypeError("Expected argument 'maintenance_policy' to be a str")
        pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if metadata_options and not isinstance(metadata_options, dict):
            raise TypeError("Expected argument 'metadata_options' to be a dict")
        pulumi.set(__self__, "metadata_options", metadata_options)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_acceleration_type and not isinstance(network_acceleration_type, str):
            raise TypeError("Expected argument 'network_acceleration_type' to be a str")
        pulumi.set(__self__, "network_acceleration_type", network_acceleration_type)
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        pulumi.set(__self__, "network_interfaces", network_interfaces)
        if placement_policy and not isinstance(placement_policy, dict):
            raise TypeError("Expected argument 'placement_policy' to be a dict")
        pulumi.set(__self__, "placement_policy", placement_policy)
        if platform_id and not isinstance(platform_id, str):
            raise TypeError("Expected argument 'platform_id' to be a str")
        pulumi.set(__self__, "platform_id", platform_id)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if scheduling_policies and not isinstance(scheduling_policies, list):
            raise TypeError("Expected argument 'scheduling_policies' to be a list")
        pulumi.set(__self__, "scheduling_policies", scheduling_policies)
        if secondary_disks and not isinstance(secondary_disks, list):
            raise TypeError("Expected argument 'secondary_disks' to be a list")
        pulumi.set(__self__, "secondary_disks", secondary_disks)
        if service_account_id and not isinstance(service_account_id, str):
            raise TypeError("Expected argument 'service_account_id' to be a str")
        pulumi.set(__self__, "service_account_id", service_account_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="bootDisks")
    def boot_disks(self) -> Sequence['outputs.GetComputeInstanceBootDiskResult']:
        return pulumi.get(self, "boot_disks")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filesystems(self) -> Optional[Sequence['outputs.GetComputeInstanceFilesystemResult']]:
        return pulumi.get(self, "filesystems")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def fqdn(self) -> builtins.str:
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="gpuClusterId")
    def gpu_cluster_id(self) -> builtins.str:
        return pulumi.get(self, "gpu_cluster_id")

    @property
    @pulumi.getter(name="hardwareGenerations")
    def hardware_generations(self) -> Sequence['outputs.GetComputeInstanceHardwareGenerationResult']:
        return pulumi.get(self, "hardware_generations")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> builtins.str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="localDisks")
    def local_disks(self) -> Optional[Sequence['outputs.GetComputeInstanceLocalDiskResult']]:
        return pulumi.get(self, "local_disks")

    @property
    @pulumi.getter(name="maintenanceGracePeriod")
    def maintenance_grace_period(self) -> builtins.str:
        return pulumi.get(self, "maintenance_grace_period")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> builtins.str:
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataOptions")
    def metadata_options(self) -> 'outputs.GetComputeInstanceMetadataOptionsResult':
        return pulumi.get(self, "metadata_options")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkAccelerationType")
    def network_acceleration_type(self) -> builtins.str:
        return pulumi.get(self, "network_acceleration_type")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Sequence['outputs.GetComputeInstanceNetworkInterfaceResult']:
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="placementPolicy")
    def placement_policy(self) -> Optional['outputs.GetComputeInstancePlacementPolicyResult']:
        return pulumi.get(self, "placement_policy")

    @property
    @pulumi.getter(name="platformId")
    def platform_id(self) -> builtins.str:
        return pulumi.get(self, "platform_id")

    @property
    @pulumi.getter
    def resources(self) -> Sequence['outputs.GetComputeInstanceResourceResult']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="schedulingPolicies")
    def scheduling_policies(self) -> Sequence['outputs.GetComputeInstanceSchedulingPolicyResult']:
        return pulumi.get(self, "scheduling_policies")

    @property
    @pulumi.getter(name="secondaryDisks")
    def secondary_disks(self) -> Sequence['outputs.GetComputeInstanceSecondaryDiskResult']:
        return pulumi.get(self, "secondary_disks")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> builtins.str:
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        return pulumi.get(self, "zone")


class AwaitableGetComputeInstanceResult(GetComputeInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeInstanceResult(
            boot_disks=self.boot_disks,
            created_at=self.created_at,
            description=self.description,
            filesystems=self.filesystems,
            folder_id=self.folder_id,
            fqdn=self.fqdn,
            gpu_cluster_id=self.gpu_cluster_id,
            hardware_generations=self.hardware_generations,
            id=self.id,
            instance_id=self.instance_id,
            labels=self.labels,
            local_disks=self.local_disks,
            maintenance_grace_period=self.maintenance_grace_period,
            maintenance_policy=self.maintenance_policy,
            metadata=self.metadata,
            metadata_options=self.metadata_options,
            name=self.name,
            network_acceleration_type=self.network_acceleration_type,
            network_interfaces=self.network_interfaces,
            placement_policy=self.placement_policy,
            platform_id=self.platform_id,
            resources=self.resources,
            scheduling_policies=self.scheduling_policies,
            secondary_disks=self.secondary_disks,
            service_account_id=self.service_account_id,
            status=self.status,
            zone=self.zone)


def get_compute_instance(filesystems: Optional[Sequence[Union['GetComputeInstanceFilesystemArgs', 'GetComputeInstanceFilesystemArgsDict']]] = None,
                         folder_id: Optional[builtins.str] = None,
                         gpu_cluster_id: Optional[builtins.str] = None,
                         instance_id: Optional[builtins.str] = None,
                         local_disks: Optional[Sequence[Union['GetComputeInstanceLocalDiskArgs', 'GetComputeInstanceLocalDiskArgsDict']]] = None,
                         maintenance_grace_period: Optional[builtins.str] = None,
                         maintenance_policy: Optional[builtins.str] = None,
                         metadata_options: Optional[Union['GetComputeInstanceMetadataOptionsArgs', 'GetComputeInstanceMetadataOptionsArgsDict']] = None,
                         name: Optional[builtins.str] = None,
                         placement_policy: Optional[Union['GetComputeInstancePlacementPolicyArgs', 'GetComputeInstancePlacementPolicyArgsDict']] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeInstanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filesystems'] = filesystems
    __args__['folderId'] = folder_id
    __args__['gpuClusterId'] = gpu_cluster_id
    __args__['instanceId'] = instance_id
    __args__['localDisks'] = local_disks
    __args__['maintenanceGracePeriod'] = maintenance_grace_period
    __args__['maintenancePolicy'] = maintenance_policy
    __args__['metadataOptions'] = metadata_options
    __args__['name'] = name
    __args__['placementPolicy'] = placement_policy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getComputeInstance:getComputeInstance', __args__, opts=opts, typ=GetComputeInstanceResult).value

    return AwaitableGetComputeInstanceResult(
        boot_disks=pulumi.get(__ret__, 'boot_disks'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        filesystems=pulumi.get(__ret__, 'filesystems'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        fqdn=pulumi.get(__ret__, 'fqdn'),
        gpu_cluster_id=pulumi.get(__ret__, 'gpu_cluster_id'),
        hardware_generations=pulumi.get(__ret__, 'hardware_generations'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        labels=pulumi.get(__ret__, 'labels'),
        local_disks=pulumi.get(__ret__, 'local_disks'),
        maintenance_grace_period=pulumi.get(__ret__, 'maintenance_grace_period'),
        maintenance_policy=pulumi.get(__ret__, 'maintenance_policy'),
        metadata=pulumi.get(__ret__, 'metadata'),
        metadata_options=pulumi.get(__ret__, 'metadata_options'),
        name=pulumi.get(__ret__, 'name'),
        network_acceleration_type=pulumi.get(__ret__, 'network_acceleration_type'),
        network_interfaces=pulumi.get(__ret__, 'network_interfaces'),
        placement_policy=pulumi.get(__ret__, 'placement_policy'),
        platform_id=pulumi.get(__ret__, 'platform_id'),
        resources=pulumi.get(__ret__, 'resources'),
        scheduling_policies=pulumi.get(__ret__, 'scheduling_policies'),
        secondary_disks=pulumi.get(__ret__, 'secondary_disks'),
        service_account_id=pulumi.get(__ret__, 'service_account_id'),
        status=pulumi.get(__ret__, 'status'),
        zone=pulumi.get(__ret__, 'zone'))
def get_compute_instance_output(filesystems: Optional[pulumi.Input[Optional[Sequence[Union['GetComputeInstanceFilesystemArgs', 'GetComputeInstanceFilesystemArgsDict']]]]] = None,
                                folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                gpu_cluster_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                instance_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                local_disks: Optional[pulumi.Input[Optional[Sequence[Union['GetComputeInstanceLocalDiskArgs', 'GetComputeInstanceLocalDiskArgsDict']]]]] = None,
                                maintenance_grace_period: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                maintenance_policy: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                metadata_options: Optional[pulumi.Input[Optional[Union['GetComputeInstanceMetadataOptionsArgs', 'GetComputeInstanceMetadataOptionsArgsDict']]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                placement_policy: Optional[pulumi.Input[Optional[Union['GetComputeInstancePlacementPolicyArgs', 'GetComputeInstancePlacementPolicyArgsDict']]]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeInstanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filesystems'] = filesystems
    __args__['folderId'] = folder_id
    __args__['gpuClusterId'] = gpu_cluster_id
    __args__['instanceId'] = instance_id
    __args__['localDisks'] = local_disks
    __args__['maintenanceGracePeriod'] = maintenance_grace_period
    __args__['maintenancePolicy'] = maintenance_policy
    __args__['metadataOptions'] = metadata_options
    __args__['name'] = name
    __args__['placementPolicy'] = placement_policy
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getComputeInstance:getComputeInstance', __args__, opts=opts, typ=GetComputeInstanceResult)
    return __ret__.apply(lambda __response__: GetComputeInstanceResult(
        boot_disks=pulumi.get(__response__, 'boot_disks'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        filesystems=pulumi.get(__response__, 'filesystems'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        fqdn=pulumi.get(__response__, 'fqdn'),
        gpu_cluster_id=pulumi.get(__response__, 'gpu_cluster_id'),
        hardware_generations=pulumi.get(__response__, 'hardware_generations'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        labels=pulumi.get(__response__, 'labels'),
        local_disks=pulumi.get(__response__, 'local_disks'),
        maintenance_grace_period=pulumi.get(__response__, 'maintenance_grace_period'),
        maintenance_policy=pulumi.get(__response__, 'maintenance_policy'),
        metadata=pulumi.get(__response__, 'metadata'),
        metadata_options=pulumi.get(__response__, 'metadata_options'),
        name=pulumi.get(__response__, 'name'),
        network_acceleration_type=pulumi.get(__response__, 'network_acceleration_type'),
        network_interfaces=pulumi.get(__response__, 'network_interfaces'),
        placement_policy=pulumi.get(__response__, 'placement_policy'),
        platform_id=pulumi.get(__response__, 'platform_id'),
        resources=pulumi.get(__response__, 'resources'),
        scheduling_policies=pulumi.get(__response__, 'scheduling_policies'),
        secondary_disks=pulumi.get(__response__, 'secondary_disks'),
        service_account_id=pulumi.get(__response__, 'service_account_id'),
        status=pulumi.get(__response__, 'status'),
        zone=pulumi.get(__response__, 'zone')))

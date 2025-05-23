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
    'GetMdbRedisClusterResult',
    'AwaitableGetMdbRedisClusterResult',
    'get_mdb_redis_cluster',
    'get_mdb_redis_cluster_output',
]

@pulumi.output_type
class GetMdbRedisClusterResult:
    """
    A collection of values returned by getMdbRedisCluster.
    """
    def __init__(__self__, announce_hostnames=None, auth_sentinel=None, cluster_id=None, configs=None, created_at=None, deletion_protection=None, description=None, disk_size_autoscalings=None, environment=None, folder_id=None, health=None, hosts=None, id=None, labels=None, maintenance_windows=None, name=None, network_id=None, persistence_mode=None, resources=None, security_group_ids=None, sharded=None, status=None, tls_enabled=None):
        if announce_hostnames and not isinstance(announce_hostnames, bool):
            raise TypeError("Expected argument 'announce_hostnames' to be a bool")
        pulumi.set(__self__, "announce_hostnames", announce_hostnames)
        if auth_sentinel and not isinstance(auth_sentinel, bool):
            raise TypeError("Expected argument 'auth_sentinel' to be a bool")
        pulumi.set(__self__, "auth_sentinel", auth_sentinel)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if configs and not isinstance(configs, list):
            raise TypeError("Expected argument 'configs' to be a list")
        pulumi.set(__self__, "configs", configs)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size_autoscalings and not isinstance(disk_size_autoscalings, list):
            raise TypeError("Expected argument 'disk_size_autoscalings' to be a list")
        pulumi.set(__self__, "disk_size_autoscalings", disk_size_autoscalings)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if health and not isinstance(health, str):
            raise TypeError("Expected argument 'health' to be a str")
        pulumi.set(__self__, "health", health)
        if hosts and not isinstance(hosts, list):
            raise TypeError("Expected argument 'hosts' to be a list")
        pulumi.set(__self__, "hosts", hosts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if maintenance_windows and not isinstance(maintenance_windows, list):
            raise TypeError("Expected argument 'maintenance_windows' to be a list")
        pulumi.set(__self__, "maintenance_windows", maintenance_windows)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if persistence_mode and not isinstance(persistence_mode, str):
            raise TypeError("Expected argument 'persistence_mode' to be a str")
        pulumi.set(__self__, "persistence_mode", persistence_mode)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sharded and not isinstance(sharded, bool):
            raise TypeError("Expected argument 'sharded' to be a bool")
        pulumi.set(__self__, "sharded", sharded)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tls_enabled and not isinstance(tls_enabled, bool):
            raise TypeError("Expected argument 'tls_enabled' to be a bool")
        pulumi.set(__self__, "tls_enabled", tls_enabled)

    @property
    @pulumi.getter(name="announceHostnames")
    def announce_hostnames(self) -> builtins.bool:
        return pulumi.get(self, "announce_hostnames")

    @property
    @pulumi.getter(name="authSentinel")
    def auth_sentinel(self) -> builtins.bool:
        return pulumi.get(self, "auth_sentinel")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def configs(self) -> Sequence['outputs.GetMdbRedisClusterConfigResult']:
        return pulumi.get(self, "configs")

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
    @pulumi.getter(name="diskSizeAutoscalings")
    def disk_size_autoscalings(self) -> Sequence['outputs.GetMdbRedisClusterDiskSizeAutoscalingResult']:
        return pulumi.get(self, "disk_size_autoscalings")

    @property
    @pulumi.getter
    def environment(self) -> builtins.str:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def health(self) -> builtins.str:
        return pulumi.get(self, "health")

    @property
    @pulumi.getter
    def hosts(self) -> Sequence['outputs.GetMdbRedisClusterHostResult']:
        return pulumi.get(self, "hosts")

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
    @pulumi.getter(name="maintenanceWindows")
    def maintenance_windows(self) -> Sequence['outputs.GetMdbRedisClusterMaintenanceWindowResult']:
        return pulumi.get(self, "maintenance_windows")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> builtins.str:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="persistenceMode")
    def persistence_mode(self) -> builtins.str:
        return pulumi.get(self, "persistence_mode")

    @property
    @pulumi.getter
    def resources(self) -> Sequence['outputs.GetMdbRedisClusterResourceResult']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def sharded(self) -> builtins.bool:
        return pulumi.get(self, "sharded")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> builtins.bool:
        return pulumi.get(self, "tls_enabled")


class AwaitableGetMdbRedisClusterResult(GetMdbRedisClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMdbRedisClusterResult(
            announce_hostnames=self.announce_hostnames,
            auth_sentinel=self.auth_sentinel,
            cluster_id=self.cluster_id,
            configs=self.configs,
            created_at=self.created_at,
            deletion_protection=self.deletion_protection,
            description=self.description,
            disk_size_autoscalings=self.disk_size_autoscalings,
            environment=self.environment,
            folder_id=self.folder_id,
            health=self.health,
            hosts=self.hosts,
            id=self.id,
            labels=self.labels,
            maintenance_windows=self.maintenance_windows,
            name=self.name,
            network_id=self.network_id,
            persistence_mode=self.persistence_mode,
            resources=self.resources,
            security_group_ids=self.security_group_ids,
            sharded=self.sharded,
            status=self.status,
            tls_enabled=self.tls_enabled)


def get_mdb_redis_cluster(cluster_id: Optional[builtins.str] = None,
                          deletion_protection: Optional[builtins.bool] = None,
                          folder_id: Optional[builtins.str] = None,
                          name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMdbRedisClusterResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['deletionProtection'] = deletion_protection
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getMdbRedisCluster:getMdbRedisCluster', __args__, opts=opts, typ=GetMdbRedisClusterResult).value

    return AwaitableGetMdbRedisClusterResult(
        announce_hostnames=pulumi.get(__ret__, 'announce_hostnames'),
        auth_sentinel=pulumi.get(__ret__, 'auth_sentinel'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        configs=pulumi.get(__ret__, 'configs'),
        created_at=pulumi.get(__ret__, 'created_at'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        description=pulumi.get(__ret__, 'description'),
        disk_size_autoscalings=pulumi.get(__ret__, 'disk_size_autoscalings'),
        environment=pulumi.get(__ret__, 'environment'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        health=pulumi.get(__ret__, 'health'),
        hosts=pulumi.get(__ret__, 'hosts'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        maintenance_windows=pulumi.get(__ret__, 'maintenance_windows'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        persistence_mode=pulumi.get(__ret__, 'persistence_mode'),
        resources=pulumi.get(__ret__, 'resources'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        sharded=pulumi.get(__ret__, 'sharded'),
        status=pulumi.get(__ret__, 'status'),
        tls_enabled=pulumi.get(__ret__, 'tls_enabled'))
def get_mdb_redis_cluster_output(cluster_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 deletion_protection: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                 folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMdbRedisClusterResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['deletionProtection'] = deletion_protection
    __args__['folderId'] = folder_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getMdbRedisCluster:getMdbRedisCluster', __args__, opts=opts, typ=GetMdbRedisClusterResult)
    return __ret__.apply(lambda __response__: GetMdbRedisClusterResult(
        announce_hostnames=pulumi.get(__response__, 'announce_hostnames'),
        auth_sentinel=pulumi.get(__response__, 'auth_sentinel'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        configs=pulumi.get(__response__, 'configs'),
        created_at=pulumi.get(__response__, 'created_at'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        description=pulumi.get(__response__, 'description'),
        disk_size_autoscalings=pulumi.get(__response__, 'disk_size_autoscalings'),
        environment=pulumi.get(__response__, 'environment'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        health=pulumi.get(__response__, 'health'),
        hosts=pulumi.get(__response__, 'hosts'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        maintenance_windows=pulumi.get(__response__, 'maintenance_windows'),
        name=pulumi.get(__response__, 'name'),
        network_id=pulumi.get(__response__, 'network_id'),
        persistence_mode=pulumi.get(__response__, 'persistence_mode'),
        resources=pulumi.get(__response__, 'resources'),
        security_group_ids=pulumi.get(__response__, 'security_group_ids'),
        sharded=pulumi.get(__response__, 'sharded'),
        status=pulumi.get(__response__, 'status'),
        tls_enabled=pulumi.get(__response__, 'tls_enabled')))

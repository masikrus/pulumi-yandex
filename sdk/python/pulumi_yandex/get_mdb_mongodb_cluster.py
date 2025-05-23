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
    'GetMdbMongodbClusterResult',
    'AwaitableGetMdbMongodbClusterResult',
    'get_mdb_mongodb_cluster',
    'get_mdb_mongodb_cluster_output',
]

@pulumi.output_type
class GetMdbMongodbClusterResult:
    """
    A collection of values returned by getMdbMongodbCluster.
    """
    def __init__(__self__, cluster_config=None, cluster_id=None, created_at=None, databases=None, deletion_protection=None, description=None, disk_size_autoscaling_mongocfg=None, disk_size_autoscaling_mongod=None, disk_size_autoscaling_mongoinfra=None, disk_size_autoscaling_mongos=None, environment=None, folder_id=None, health=None, hosts=None, id=None, labels=None, maintenance_window=None, name=None, network_id=None, resources=None, resources_mongocfg=None, resources_mongod=None, resources_mongoinfra=None, resources_mongos=None, restore=None, security_group_ids=None, sharded=None, status=None, users=None):
        if cluster_config and not isinstance(cluster_config, dict):
            raise TypeError("Expected argument 'cluster_config' to be a dict")
        pulumi.set(__self__, "cluster_config", cluster_config)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size_autoscaling_mongocfg and not isinstance(disk_size_autoscaling_mongocfg, dict):
            raise TypeError("Expected argument 'disk_size_autoscaling_mongocfg' to be a dict")
        pulumi.set(__self__, "disk_size_autoscaling_mongocfg", disk_size_autoscaling_mongocfg)
        if disk_size_autoscaling_mongod and not isinstance(disk_size_autoscaling_mongod, dict):
            raise TypeError("Expected argument 'disk_size_autoscaling_mongod' to be a dict")
        pulumi.set(__self__, "disk_size_autoscaling_mongod", disk_size_autoscaling_mongod)
        if disk_size_autoscaling_mongoinfra and not isinstance(disk_size_autoscaling_mongoinfra, dict):
            raise TypeError("Expected argument 'disk_size_autoscaling_mongoinfra' to be a dict")
        pulumi.set(__self__, "disk_size_autoscaling_mongoinfra", disk_size_autoscaling_mongoinfra)
        if disk_size_autoscaling_mongos and not isinstance(disk_size_autoscaling_mongos, dict):
            raise TypeError("Expected argument 'disk_size_autoscaling_mongos' to be a dict")
        pulumi.set(__self__, "disk_size_autoscaling_mongos", disk_size_autoscaling_mongos)
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
        if maintenance_window and not isinstance(maintenance_window, dict):
            raise TypeError("Expected argument 'maintenance_window' to be a dict")
        pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if resources and not isinstance(resources, dict):
            raise TypeError("Expected argument 'resources' to be a dict")
        pulumi.set(__self__, "resources", resources)
        if resources_mongocfg and not isinstance(resources_mongocfg, dict):
            raise TypeError("Expected argument 'resources_mongocfg' to be a dict")
        pulumi.set(__self__, "resources_mongocfg", resources_mongocfg)
        if resources_mongod and not isinstance(resources_mongod, dict):
            raise TypeError("Expected argument 'resources_mongod' to be a dict")
        pulumi.set(__self__, "resources_mongod", resources_mongod)
        if resources_mongoinfra and not isinstance(resources_mongoinfra, dict):
            raise TypeError("Expected argument 'resources_mongoinfra' to be a dict")
        pulumi.set(__self__, "resources_mongoinfra", resources_mongoinfra)
        if resources_mongos and not isinstance(resources_mongos, dict):
            raise TypeError("Expected argument 'resources_mongos' to be a dict")
        pulumi.set(__self__, "resources_mongos", resources_mongos)
        if restore and not isinstance(restore, dict):
            raise TypeError("Expected argument 'restore' to be a dict")
        pulumi.set(__self__, "restore", restore)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sharded and not isinstance(sharded, bool):
            raise TypeError("Expected argument 'sharded' to be a bool")
        pulumi.set(__self__, "sharded", sharded)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="clusterConfig")
    def cluster_config(self) -> Optional['outputs.GetMdbMongodbClusterClusterConfigResult']:
        return pulumi.get(self, "cluster_config")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    @_utilities.deprecated("""to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database""")
    def databases(self) -> Sequence['outputs.GetMdbMongodbClusterDatabaseResult']:
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> builtins.bool:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSizeAutoscalingMongocfg")
    def disk_size_autoscaling_mongocfg(self) -> 'outputs.GetMdbMongodbClusterDiskSizeAutoscalingMongocfgResult':
        return pulumi.get(self, "disk_size_autoscaling_mongocfg")

    @property
    @pulumi.getter(name="diskSizeAutoscalingMongod")
    def disk_size_autoscaling_mongod(self) -> 'outputs.GetMdbMongodbClusterDiskSizeAutoscalingMongodResult':
        return pulumi.get(self, "disk_size_autoscaling_mongod")

    @property
    @pulumi.getter(name="diskSizeAutoscalingMongoinfra")
    def disk_size_autoscaling_mongoinfra(self) -> 'outputs.GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraResult':
        return pulumi.get(self, "disk_size_autoscaling_mongoinfra")

    @property
    @pulumi.getter(name="diskSizeAutoscalingMongos")
    def disk_size_autoscaling_mongos(self) -> 'outputs.GetMdbMongodbClusterDiskSizeAutoscalingMongosResult':
        return pulumi.get(self, "disk_size_autoscaling_mongos")

    @property
    @pulumi.getter
    def environment(self) -> Optional[builtins.str]:
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
    def hosts(self) -> Optional[Sequence['outputs.GetMdbMongodbClusterHostResult']]:
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
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> 'outputs.GetMdbMongodbClusterMaintenanceWindowResult':
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    @_utilities.deprecated("""to manage `resources`s, please switch to using a separate resource type `resources_mongo*`""")
    def resources(self) -> Optional['outputs.GetMdbMongodbClusterResourcesResult']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="resourcesMongocfg")
    def resources_mongocfg(self) -> Optional['outputs.GetMdbMongodbClusterResourcesMongocfgResult']:
        return pulumi.get(self, "resources_mongocfg")

    @property
    @pulumi.getter(name="resourcesMongod")
    def resources_mongod(self) -> Optional['outputs.GetMdbMongodbClusterResourcesMongodResult']:
        return pulumi.get(self, "resources_mongod")

    @property
    @pulumi.getter(name="resourcesMongoinfra")
    def resources_mongoinfra(self) -> Optional['outputs.GetMdbMongodbClusterResourcesMongoinfraResult']:
        return pulumi.get(self, "resources_mongoinfra")

    @property
    @pulumi.getter(name="resourcesMongos")
    def resources_mongos(self) -> Optional['outputs.GetMdbMongodbClusterResourcesMongosResult']:
        return pulumi.get(self, "resources_mongos")

    @property
    @pulumi.getter
    def restore(self) -> Optional['outputs.GetMdbMongodbClusterRestoreResult']:
        return pulumi.get(self, "restore")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[builtins.str]]:
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
    @pulumi.getter
    @_utilities.deprecated("""to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user""")
    def users(self) -> Sequence['outputs.GetMdbMongodbClusterUserResult']:
        return pulumi.get(self, "users")


class AwaitableGetMdbMongodbClusterResult(GetMdbMongodbClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMdbMongodbClusterResult(
            cluster_config=self.cluster_config,
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            databases=self.databases,
            deletion_protection=self.deletion_protection,
            description=self.description,
            disk_size_autoscaling_mongocfg=self.disk_size_autoscaling_mongocfg,
            disk_size_autoscaling_mongod=self.disk_size_autoscaling_mongod,
            disk_size_autoscaling_mongoinfra=self.disk_size_autoscaling_mongoinfra,
            disk_size_autoscaling_mongos=self.disk_size_autoscaling_mongos,
            environment=self.environment,
            folder_id=self.folder_id,
            health=self.health,
            hosts=self.hosts,
            id=self.id,
            labels=self.labels,
            maintenance_window=self.maintenance_window,
            name=self.name,
            network_id=self.network_id,
            resources=self.resources,
            resources_mongocfg=self.resources_mongocfg,
            resources_mongod=self.resources_mongod,
            resources_mongoinfra=self.resources_mongoinfra,
            resources_mongos=self.resources_mongos,
            restore=self.restore,
            security_group_ids=self.security_group_ids,
            sharded=self.sharded,
            status=self.status,
            users=self.users)


def get_mdb_mongodb_cluster(cluster_config: Optional[Union['GetMdbMongodbClusterClusterConfigArgs', 'GetMdbMongodbClusterClusterConfigArgsDict']] = None,
                            cluster_id: Optional[builtins.str] = None,
                            created_at: Optional[builtins.str] = None,
                            databases: Optional[Sequence[Union['GetMdbMongodbClusterDatabaseArgs', 'GetMdbMongodbClusterDatabaseArgsDict']]] = None,
                            deletion_protection: Optional[builtins.bool] = None,
                            description: Optional[builtins.str] = None,
                            disk_size_autoscaling_mongocfg: Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongocfgArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongocfgArgsDict']] = None,
                            disk_size_autoscaling_mongod: Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongodArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongodArgsDict']] = None,
                            disk_size_autoscaling_mongoinfra: Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraArgsDict']] = None,
                            disk_size_autoscaling_mongos: Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongosArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongosArgsDict']] = None,
                            environment: Optional[builtins.str] = None,
                            folder_id: Optional[builtins.str] = None,
                            health: Optional[builtins.str] = None,
                            hosts: Optional[Sequence[Union['GetMdbMongodbClusterHostArgs', 'GetMdbMongodbClusterHostArgsDict']]] = None,
                            labels: Optional[Mapping[str, builtins.str]] = None,
                            maintenance_window: Optional[Union['GetMdbMongodbClusterMaintenanceWindowArgs', 'GetMdbMongodbClusterMaintenanceWindowArgsDict']] = None,
                            name: Optional[builtins.str] = None,
                            network_id: Optional[builtins.str] = None,
                            resources: Optional[Union['GetMdbMongodbClusterResourcesArgs', 'GetMdbMongodbClusterResourcesArgsDict']] = None,
                            resources_mongocfg: Optional[Union['GetMdbMongodbClusterResourcesMongocfgArgs', 'GetMdbMongodbClusterResourcesMongocfgArgsDict']] = None,
                            resources_mongod: Optional[Union['GetMdbMongodbClusterResourcesMongodArgs', 'GetMdbMongodbClusterResourcesMongodArgsDict']] = None,
                            resources_mongoinfra: Optional[Union['GetMdbMongodbClusterResourcesMongoinfraArgs', 'GetMdbMongodbClusterResourcesMongoinfraArgsDict']] = None,
                            resources_mongos: Optional[Union['GetMdbMongodbClusterResourcesMongosArgs', 'GetMdbMongodbClusterResourcesMongosArgsDict']] = None,
                            restore: Optional[Union['GetMdbMongodbClusterRestoreArgs', 'GetMdbMongodbClusterRestoreArgsDict']] = None,
                            security_group_ids: Optional[Sequence[builtins.str]] = None,
                            sharded: Optional[builtins.bool] = None,
                            status: Optional[builtins.str] = None,
                            users: Optional[Sequence[Union['GetMdbMongodbClusterUserArgs', 'GetMdbMongodbClusterUserArgsDict']]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMdbMongodbClusterResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterConfig'] = cluster_config
    __args__['clusterId'] = cluster_id
    __args__['createdAt'] = created_at
    __args__['databases'] = databases
    __args__['deletionProtection'] = deletion_protection
    __args__['description'] = description
    __args__['diskSizeAutoscalingMongocfg'] = disk_size_autoscaling_mongocfg
    __args__['diskSizeAutoscalingMongod'] = disk_size_autoscaling_mongod
    __args__['diskSizeAutoscalingMongoinfra'] = disk_size_autoscaling_mongoinfra
    __args__['diskSizeAutoscalingMongos'] = disk_size_autoscaling_mongos
    __args__['environment'] = environment
    __args__['folderId'] = folder_id
    __args__['health'] = health
    __args__['hosts'] = hosts
    __args__['labels'] = labels
    __args__['maintenanceWindow'] = maintenance_window
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['resources'] = resources
    __args__['resourcesMongocfg'] = resources_mongocfg
    __args__['resourcesMongod'] = resources_mongod
    __args__['resourcesMongoinfra'] = resources_mongoinfra
    __args__['resourcesMongos'] = resources_mongos
    __args__['restore'] = restore
    __args__['securityGroupIds'] = security_group_ids
    __args__['sharded'] = sharded
    __args__['status'] = status
    __args__['users'] = users
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getMdbMongodbCluster:getMdbMongodbCluster', __args__, opts=opts, typ=GetMdbMongodbClusterResult).value

    return AwaitableGetMdbMongodbClusterResult(
        cluster_config=pulumi.get(__ret__, 'cluster_config'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        databases=pulumi.get(__ret__, 'databases'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        description=pulumi.get(__ret__, 'description'),
        disk_size_autoscaling_mongocfg=pulumi.get(__ret__, 'disk_size_autoscaling_mongocfg'),
        disk_size_autoscaling_mongod=pulumi.get(__ret__, 'disk_size_autoscaling_mongod'),
        disk_size_autoscaling_mongoinfra=pulumi.get(__ret__, 'disk_size_autoscaling_mongoinfra'),
        disk_size_autoscaling_mongos=pulumi.get(__ret__, 'disk_size_autoscaling_mongos'),
        environment=pulumi.get(__ret__, 'environment'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        health=pulumi.get(__ret__, 'health'),
        hosts=pulumi.get(__ret__, 'hosts'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        maintenance_window=pulumi.get(__ret__, 'maintenance_window'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        resources=pulumi.get(__ret__, 'resources'),
        resources_mongocfg=pulumi.get(__ret__, 'resources_mongocfg'),
        resources_mongod=pulumi.get(__ret__, 'resources_mongod'),
        resources_mongoinfra=pulumi.get(__ret__, 'resources_mongoinfra'),
        resources_mongos=pulumi.get(__ret__, 'resources_mongos'),
        restore=pulumi.get(__ret__, 'restore'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        sharded=pulumi.get(__ret__, 'sharded'),
        status=pulumi.get(__ret__, 'status'),
        users=pulumi.get(__ret__, 'users'))
def get_mdb_mongodb_cluster_output(cluster_config: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterClusterConfigArgs', 'GetMdbMongodbClusterClusterConfigArgsDict']]]] = None,
                                   cluster_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   created_at: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   databases: Optional[pulumi.Input[Optional[Sequence[Union['GetMdbMongodbClusterDatabaseArgs', 'GetMdbMongodbClusterDatabaseArgsDict']]]]] = None,
                                   deletion_protection: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                   description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   disk_size_autoscaling_mongocfg: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongocfgArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongocfgArgsDict']]]] = None,
                                   disk_size_autoscaling_mongod: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongodArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongodArgsDict']]]] = None,
                                   disk_size_autoscaling_mongoinfra: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraArgsDict']]]] = None,
                                   disk_size_autoscaling_mongos: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterDiskSizeAutoscalingMongosArgs', 'GetMdbMongodbClusterDiskSizeAutoscalingMongosArgsDict']]]] = None,
                                   environment: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   health: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   hosts: Optional[pulumi.Input[Optional[Sequence[Union['GetMdbMongodbClusterHostArgs', 'GetMdbMongodbClusterHostArgsDict']]]]] = None,
                                   labels: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                                   maintenance_window: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterMaintenanceWindowArgs', 'GetMdbMongodbClusterMaintenanceWindowArgsDict']]]] = None,
                                   name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   network_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   resources: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterResourcesArgs', 'GetMdbMongodbClusterResourcesArgsDict']]]] = None,
                                   resources_mongocfg: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterResourcesMongocfgArgs', 'GetMdbMongodbClusterResourcesMongocfgArgsDict']]]] = None,
                                   resources_mongod: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterResourcesMongodArgs', 'GetMdbMongodbClusterResourcesMongodArgsDict']]]] = None,
                                   resources_mongoinfra: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterResourcesMongoinfraArgs', 'GetMdbMongodbClusterResourcesMongoinfraArgsDict']]]] = None,
                                   resources_mongos: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterResourcesMongosArgs', 'GetMdbMongodbClusterResourcesMongosArgsDict']]]] = None,
                                   restore: Optional[pulumi.Input[Optional[Union['GetMdbMongodbClusterRestoreArgs', 'GetMdbMongodbClusterRestoreArgsDict']]]] = None,
                                   security_group_ids: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                   sharded: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                   status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   users: Optional[pulumi.Input[Optional[Sequence[Union['GetMdbMongodbClusterUserArgs', 'GetMdbMongodbClusterUserArgsDict']]]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMdbMongodbClusterResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterConfig'] = cluster_config
    __args__['clusterId'] = cluster_id
    __args__['createdAt'] = created_at
    __args__['databases'] = databases
    __args__['deletionProtection'] = deletion_protection
    __args__['description'] = description
    __args__['diskSizeAutoscalingMongocfg'] = disk_size_autoscaling_mongocfg
    __args__['diskSizeAutoscalingMongod'] = disk_size_autoscaling_mongod
    __args__['diskSizeAutoscalingMongoinfra'] = disk_size_autoscaling_mongoinfra
    __args__['diskSizeAutoscalingMongos'] = disk_size_autoscaling_mongos
    __args__['environment'] = environment
    __args__['folderId'] = folder_id
    __args__['health'] = health
    __args__['hosts'] = hosts
    __args__['labels'] = labels
    __args__['maintenanceWindow'] = maintenance_window
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['resources'] = resources
    __args__['resourcesMongocfg'] = resources_mongocfg
    __args__['resourcesMongod'] = resources_mongod
    __args__['resourcesMongoinfra'] = resources_mongoinfra
    __args__['resourcesMongos'] = resources_mongos
    __args__['restore'] = restore
    __args__['securityGroupIds'] = security_group_ids
    __args__['sharded'] = sharded
    __args__['status'] = status
    __args__['users'] = users
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getMdbMongodbCluster:getMdbMongodbCluster', __args__, opts=opts, typ=GetMdbMongodbClusterResult)
    return __ret__.apply(lambda __response__: GetMdbMongodbClusterResult(
        cluster_config=pulumi.get(__response__, 'cluster_config'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        databases=pulumi.get(__response__, 'databases'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        description=pulumi.get(__response__, 'description'),
        disk_size_autoscaling_mongocfg=pulumi.get(__response__, 'disk_size_autoscaling_mongocfg'),
        disk_size_autoscaling_mongod=pulumi.get(__response__, 'disk_size_autoscaling_mongod'),
        disk_size_autoscaling_mongoinfra=pulumi.get(__response__, 'disk_size_autoscaling_mongoinfra'),
        disk_size_autoscaling_mongos=pulumi.get(__response__, 'disk_size_autoscaling_mongos'),
        environment=pulumi.get(__response__, 'environment'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        health=pulumi.get(__response__, 'health'),
        hosts=pulumi.get(__response__, 'hosts'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        maintenance_window=pulumi.get(__response__, 'maintenance_window'),
        name=pulumi.get(__response__, 'name'),
        network_id=pulumi.get(__response__, 'network_id'),
        resources=pulumi.get(__response__, 'resources'),
        resources_mongocfg=pulumi.get(__response__, 'resources_mongocfg'),
        resources_mongod=pulumi.get(__response__, 'resources_mongod'),
        resources_mongoinfra=pulumi.get(__response__, 'resources_mongoinfra'),
        resources_mongos=pulumi.get(__response__, 'resources_mongos'),
        restore=pulumi.get(__response__, 'restore'),
        security_group_ids=pulumi.get(__response__, 'security_group_ids'),
        sharded=pulumi.get(__response__, 'sharded'),
        status=pulumi.get(__response__, 'status'),
        users=pulumi.get(__response__, 'users')))

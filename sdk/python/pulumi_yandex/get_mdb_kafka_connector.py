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
    'GetMdbKafkaConnectorResult',
    'AwaitableGetMdbKafkaConnectorResult',
    'get_mdb_kafka_connector',
    'get_mdb_kafka_connector_output',
]

@pulumi.output_type
class GetMdbKafkaConnectorResult:
    """
    A collection of values returned by getMdbKafkaConnector.
    """
    def __init__(__self__, cluster_id=None, connector_config_mirrormakers=None, connector_config_s3_sinks=None, id=None, name=None, properties=None, tasks_max=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if connector_config_mirrormakers and not isinstance(connector_config_mirrormakers, list):
            raise TypeError("Expected argument 'connector_config_mirrormakers' to be a list")
        pulumi.set(__self__, "connector_config_mirrormakers", connector_config_mirrormakers)
        if connector_config_s3_sinks and not isinstance(connector_config_s3_sinks, list):
            raise TypeError("Expected argument 'connector_config_s3_sinks' to be a list")
        pulumi.set(__self__, "connector_config_s3_sinks", connector_config_s3_sinks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if tasks_max and not isinstance(tasks_max, int):
            raise TypeError("Expected argument 'tasks_max' to be a int")
        pulumi.set(__self__, "tasks_max", tasks_max)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="connectorConfigMirrormakers")
    def connector_config_mirrormakers(self) -> Sequence['outputs.GetMdbKafkaConnectorConnectorConfigMirrormakerResult']:
        return pulumi.get(self, "connector_config_mirrormakers")

    @property
    @pulumi.getter(name="connectorConfigS3Sinks")
    def connector_config_s3_sinks(self) -> Sequence['outputs.GetMdbKafkaConnectorConnectorConfigS3SinkResult']:
        return pulumi.get(self, "connector_config_s3_sinks")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="tasksMax")
    def tasks_max(self) -> builtins.int:
        return pulumi.get(self, "tasks_max")


class AwaitableGetMdbKafkaConnectorResult(GetMdbKafkaConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMdbKafkaConnectorResult(
            cluster_id=self.cluster_id,
            connector_config_mirrormakers=self.connector_config_mirrormakers,
            connector_config_s3_sinks=self.connector_config_s3_sinks,
            id=self.id,
            name=self.name,
            properties=self.properties,
            tasks_max=self.tasks_max)


def get_mdb_kafka_connector(cluster_id: Optional[builtins.str] = None,
                            name: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMdbKafkaConnectorResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getMdbKafkaConnector:getMdbKafkaConnector', __args__, opts=opts, typ=GetMdbKafkaConnectorResult).value

    return AwaitableGetMdbKafkaConnectorResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        connector_config_mirrormakers=pulumi.get(__ret__, 'connector_config_mirrormakers'),
        connector_config_s3_sinks=pulumi.get(__ret__, 'connector_config_s3_sinks'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        tasks_max=pulumi.get(__ret__, 'tasks_max'))
def get_mdb_kafka_connector_output(cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                                   name: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMdbKafkaConnectorResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getMdbKafkaConnector:getMdbKafkaConnector', __args__, opts=opts, typ=GetMdbKafkaConnectorResult)
    return __ret__.apply(lambda __response__: GetMdbKafkaConnectorResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        connector_config_mirrormakers=pulumi.get(__response__, 'connector_config_mirrormakers'),
        connector_config_s3_sinks=pulumi.get(__response__, 'connector_config_s3_sinks'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        tasks_max=pulumi.get(__response__, 'tasks_max')))

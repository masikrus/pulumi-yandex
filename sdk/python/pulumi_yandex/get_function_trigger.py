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
    'GetFunctionTriggerResult',
    'AwaitableGetFunctionTriggerResult',
    'get_function_trigger',
    'get_function_trigger_output',
]

@pulumi.output_type
class GetFunctionTriggerResult:
    """
    A collection of values returned by getFunctionTrigger.
    """
    def __init__(__self__, container_registries=None, containers=None, created_at=None, data_streams=None, description=None, dlqs=None, folder_id=None, functions=None, id=None, iots=None, labels=None, log_groups=None, loggings=None, mails=None, message_queues=None, name=None, object_storages=None, timers=None, trigger_id=None):
        if container_registries and not isinstance(container_registries, list):
            raise TypeError("Expected argument 'container_registries' to be a list")
        pulumi.set(__self__, "container_registries", container_registries)
        if containers and not isinstance(containers, list):
            raise TypeError("Expected argument 'containers' to be a list")
        pulumi.set(__self__, "containers", containers)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if data_streams and not isinstance(data_streams, list):
            raise TypeError("Expected argument 'data_streams' to be a list")
        pulumi.set(__self__, "data_streams", data_streams)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dlqs and not isinstance(dlqs, list):
            raise TypeError("Expected argument 'dlqs' to be a list")
        pulumi.set(__self__, "dlqs", dlqs)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if functions and not isinstance(functions, list):
            raise TypeError("Expected argument 'functions' to be a list")
        pulumi.set(__self__, "functions", functions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iots and not isinstance(iots, list):
            raise TypeError("Expected argument 'iots' to be a list")
        pulumi.set(__self__, "iots", iots)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if log_groups and not isinstance(log_groups, list):
            raise TypeError("Expected argument 'log_groups' to be a list")
        pulumi.set(__self__, "log_groups", log_groups)
        if loggings and not isinstance(loggings, list):
            raise TypeError("Expected argument 'loggings' to be a list")
        pulumi.set(__self__, "loggings", loggings)
        if mails and not isinstance(mails, list):
            raise TypeError("Expected argument 'mails' to be a list")
        pulumi.set(__self__, "mails", mails)
        if message_queues and not isinstance(message_queues, list):
            raise TypeError("Expected argument 'message_queues' to be a list")
        pulumi.set(__self__, "message_queues", message_queues)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if object_storages and not isinstance(object_storages, list):
            raise TypeError("Expected argument 'object_storages' to be a list")
        pulumi.set(__self__, "object_storages", object_storages)
        if timers and not isinstance(timers, list):
            raise TypeError("Expected argument 'timers' to be a list")
        pulumi.set(__self__, "timers", timers)
        if trigger_id and not isinstance(trigger_id, str):
            raise TypeError("Expected argument 'trigger_id' to be a str")
        pulumi.set(__self__, "trigger_id", trigger_id)

    @property
    @pulumi.getter(name="containerRegistries")
    def container_registries(self) -> Sequence['outputs.GetFunctionTriggerContainerRegistryResult']:
        return pulumi.get(self, "container_registries")

    @property
    @pulumi.getter
    def containers(self) -> Sequence['outputs.GetFunctionTriggerContainerResult']:
        return pulumi.get(self, "containers")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dataStreams")
    def data_streams(self) -> Sequence['outputs.GetFunctionTriggerDataStreamResult']:
        return pulumi.get(self, "data_streams")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dlqs(self) -> Sequence['outputs.GetFunctionTriggerDlqResult']:
        return pulumi.get(self, "dlqs")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def functions(self) -> Sequence['outputs.GetFunctionTriggerFunctionResult']:
        return pulumi.get(self, "functions")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iots(self) -> Sequence['outputs.GetFunctionTriggerIotResult']:
        return pulumi.get(self, "iots")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, builtins.str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="logGroups")
    def log_groups(self) -> Sequence['outputs.GetFunctionTriggerLogGroupResult']:
        return pulumi.get(self, "log_groups")

    @property
    @pulumi.getter
    def loggings(self) -> Sequence['outputs.GetFunctionTriggerLoggingResult']:
        return pulumi.get(self, "loggings")

    @property
    @pulumi.getter
    def mails(self) -> Sequence['outputs.GetFunctionTriggerMailResult']:
        return pulumi.get(self, "mails")

    @property
    @pulumi.getter(name="messageQueues")
    def message_queues(self) -> Sequence['outputs.GetFunctionTriggerMessageQueueResult']:
        return pulumi.get(self, "message_queues")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectStorages")
    def object_storages(self) -> Sequence['outputs.GetFunctionTriggerObjectStorageResult']:
        return pulumi.get(self, "object_storages")

    @property
    @pulumi.getter
    def timers(self) -> Sequence['outputs.GetFunctionTriggerTimerResult']:
        return pulumi.get(self, "timers")

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "trigger_id")


class AwaitableGetFunctionTriggerResult(GetFunctionTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionTriggerResult(
            container_registries=self.container_registries,
            containers=self.containers,
            created_at=self.created_at,
            data_streams=self.data_streams,
            description=self.description,
            dlqs=self.dlqs,
            folder_id=self.folder_id,
            functions=self.functions,
            id=self.id,
            iots=self.iots,
            labels=self.labels,
            log_groups=self.log_groups,
            loggings=self.loggings,
            mails=self.mails,
            message_queues=self.message_queues,
            name=self.name,
            object_storages=self.object_storages,
            timers=self.timers,
            trigger_id=self.trigger_id)


def get_function_trigger(folder_id: Optional[builtins.str] = None,
                         name: Optional[builtins.str] = None,
                         trigger_id: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionTriggerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['triggerId'] = trigger_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getFunctionTrigger:getFunctionTrigger', __args__, opts=opts, typ=GetFunctionTriggerResult).value

    return AwaitableGetFunctionTriggerResult(
        container_registries=pulumi.get(__ret__, 'container_registries'),
        containers=pulumi.get(__ret__, 'containers'),
        created_at=pulumi.get(__ret__, 'created_at'),
        data_streams=pulumi.get(__ret__, 'data_streams'),
        description=pulumi.get(__ret__, 'description'),
        dlqs=pulumi.get(__ret__, 'dlqs'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        functions=pulumi.get(__ret__, 'functions'),
        id=pulumi.get(__ret__, 'id'),
        iots=pulumi.get(__ret__, 'iots'),
        labels=pulumi.get(__ret__, 'labels'),
        log_groups=pulumi.get(__ret__, 'log_groups'),
        loggings=pulumi.get(__ret__, 'loggings'),
        mails=pulumi.get(__ret__, 'mails'),
        message_queues=pulumi.get(__ret__, 'message_queues'),
        name=pulumi.get(__ret__, 'name'),
        object_storages=pulumi.get(__ret__, 'object_storages'),
        timers=pulumi.get(__ret__, 'timers'),
        trigger_id=pulumi.get(__ret__, 'trigger_id'))
def get_function_trigger_output(folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                trigger_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFunctionTriggerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['triggerId'] = trigger_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getFunctionTrigger:getFunctionTrigger', __args__, opts=opts, typ=GetFunctionTriggerResult)
    return __ret__.apply(lambda __response__: GetFunctionTriggerResult(
        container_registries=pulumi.get(__response__, 'container_registries'),
        containers=pulumi.get(__response__, 'containers'),
        created_at=pulumi.get(__response__, 'created_at'),
        data_streams=pulumi.get(__response__, 'data_streams'),
        description=pulumi.get(__response__, 'description'),
        dlqs=pulumi.get(__response__, 'dlqs'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        functions=pulumi.get(__response__, 'functions'),
        id=pulumi.get(__response__, 'id'),
        iots=pulumi.get(__response__, 'iots'),
        labels=pulumi.get(__response__, 'labels'),
        log_groups=pulumi.get(__response__, 'log_groups'),
        loggings=pulumi.get(__response__, 'loggings'),
        mails=pulumi.get(__response__, 'mails'),
        message_queues=pulumi.get(__response__, 'message_queues'),
        name=pulumi.get(__response__, 'name'),
        object_storages=pulumi.get(__response__, 'object_storages'),
        timers=pulumi.get(__response__, 'timers'),
        trigger_id=pulumi.get(__response__, 'trigger_id')))

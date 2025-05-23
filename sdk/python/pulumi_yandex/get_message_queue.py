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

__all__ = [
    'GetMessageQueueResult',
    'AwaitableGetMessageQueueResult',
    'get_message_queue',
    'get_message_queue_output',
]

@pulumi.output_type
class GetMessageQueueResult:
    """
    A collection of values returned by getMessageQueue.
    """
    def __init__(__self__, access_key=None, arn=None, id=None, name=None, region_id=None, secret_key=None, url=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region_id and not isinstance(region_id, str):
            raise TypeError("Expected argument 'region_id' to be a str")
        pulumi.set(__self__, "region_id", region_id)
        if secret_key and not isinstance(secret_key, str):
            raise TypeError("Expected argument 'secret_key' to be a str")
        pulumi.set(__self__, "secret_key", secret_key)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[builtins.str]:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def arn(self) -> builtins.str:
        return pulumi.get(self, "arn")

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
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[builtins.str]:
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        return pulumi.get(self, "url")


class AwaitableGetMessageQueueResult(GetMessageQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMessageQueueResult(
            access_key=self.access_key,
            arn=self.arn,
            id=self.id,
            name=self.name,
            region_id=self.region_id,
            secret_key=self.secret_key,
            url=self.url)


def get_message_queue(access_key: Optional[builtins.str] = None,
                      name: Optional[builtins.str] = None,
                      region_id: Optional[builtins.str] = None,
                      secret_key: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMessageQueueResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['accessKey'] = access_key
    __args__['name'] = name
    __args__['regionId'] = region_id
    __args__['secretKey'] = secret_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getMessageQueue:getMessageQueue', __args__, opts=opts, typ=GetMessageQueueResult).value

    return AwaitableGetMessageQueueResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        arn=pulumi.get(__ret__, 'arn'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region_id=pulumi.get(__ret__, 'region_id'),
        secret_key=pulumi.get(__ret__, 'secret_key'),
        url=pulumi.get(__ret__, 'url'))
def get_message_queue_output(access_key: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             name: Optional[pulumi.Input[builtins.str]] = None,
                             region_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             secret_key: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMessageQueueResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['accessKey'] = access_key
    __args__['name'] = name
    __args__['regionId'] = region_id
    __args__['secretKey'] = secret_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getMessageQueue:getMessageQueue', __args__, opts=opts, typ=GetMessageQueueResult)
    return __ret__.apply(lambda __response__: GetMessageQueueResult(
        access_key=pulumi.get(__response__, 'access_key'),
        arn=pulumi.get(__response__, 'arn'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region_id=pulumi.get(__response__, 'region_id'),
        secret_key=pulumi.get(__response__, 'secret_key'),
        url=pulumi.get(__response__, 'url')))

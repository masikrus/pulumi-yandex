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
    'GetContainerRepositoryResult',
    'AwaitableGetContainerRepositoryResult',
    'get_container_repository',
    'get_container_repository_output',
]

@pulumi.output_type
class GetContainerRepositoryResult:
    """
    A collection of values returned by getContainerRepository.
    """
    def __init__(__self__, id=None, name=None, repository_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if repository_id and not isinstance(repository_id, str):
            raise TypeError("Expected argument 'repository_id' to be a str")
        pulumi.set(__self__, "repository_id", repository_id)

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
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> builtins.str:
        return pulumi.get(self, "repository_id")


class AwaitableGetContainerRepositoryResult(GetContainerRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerRepositoryResult(
            id=self.id,
            name=self.name,
            repository_id=self.repository_id)


def get_container_repository(name: Optional[builtins.str] = None,
                             repository_id: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerRepositoryResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getContainerRepository:getContainerRepository', __args__, opts=opts, typ=GetContainerRepositoryResult).value

    return AwaitableGetContainerRepositoryResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        repository_id=pulumi.get(__ret__, 'repository_id'))
def get_container_repository_output(name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                    repository_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContainerRepositoryResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getContainerRepository:getContainerRepository', __args__, opts=opts, typ=GetContainerRepositoryResult)
    return __ret__.apply(lambda __response__: GetContainerRepositoryResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        repository_id=pulumi.get(__response__, 'repository_id')))

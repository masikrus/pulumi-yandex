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
    'GetIamPolicyResult',
    'AwaitableGetIamPolicyResult',
    'get_iam_policy',
    'get_iam_policy_output',
]

@pulumi.output_type
class GetIamPolicyResult:
    """
    A collection of values returned by getIamPolicy.
    """
    def __init__(__self__, bindings=None, id=None, policy_data=None):
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        pulumi.set(__self__, "bindings", bindings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter
    def bindings(self) -> Sequence['outputs.GetIamPolicyBindingResult']:
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> builtins.str:
        return pulumi.get(self, "policy_data")


class AwaitableGetIamPolicyResult(GetIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamPolicyResult(
            bindings=self.bindings,
            id=self.id,
            policy_data=self.policy_data)


def get_iam_policy(bindings: Optional[Sequence[Union['GetIamPolicyBindingArgs', 'GetIamPolicyBindingArgsDict']]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamPolicyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['bindings'] = bindings
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getIamPolicy:getIamPolicy', __args__, opts=opts, typ=GetIamPolicyResult).value

    return AwaitableGetIamPolicyResult(
        bindings=pulumi.get(__ret__, 'bindings'),
        id=pulumi.get(__ret__, 'id'),
        policy_data=pulumi.get(__ret__, 'policy_data'))
def get_iam_policy_output(bindings: Optional[pulumi.Input[Sequence[Union['GetIamPolicyBindingArgs', 'GetIamPolicyBindingArgsDict']]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIamPolicyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['bindings'] = bindings
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getIamPolicy:getIamPolicy', __args__, opts=opts, typ=GetIamPolicyResult)
    return __ret__.apply(lambda __response__: GetIamPolicyResult(
        bindings=pulumi.get(__response__, 'bindings'),
        id=pulumi.get(__response__, 'id'),
        policy_data=pulumi.get(__response__, 'policy_data')))

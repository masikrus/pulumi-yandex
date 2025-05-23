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
    'GetLockboxSecretResult',
    'AwaitableGetLockboxSecretResult',
    'get_lockbox_secret',
    'get_lockbox_secret_output',
]

@pulumi.output_type
class GetLockboxSecretResult:
    """
    A collection of values returned by getLockboxSecret.
    """
    def __init__(__self__, created_at=None, current_versions=None, deletion_protection=None, description=None, folder_id=None, id=None, kms_key_id=None, labels=None, name=None, password_payload_specifications=None, secret_id=None, status=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if current_versions and not isinstance(current_versions, list):
            raise TypeError("Expected argument 'current_versions' to be a list")
        pulumi.set(__self__, "current_versions", current_versions)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
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
        if password_payload_specifications and not isinstance(password_payload_specifications, list):
            raise TypeError("Expected argument 'password_payload_specifications' to be a list")
        pulumi.set(__self__, "password_payload_specifications", password_payload_specifications)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentVersions")
    def current_versions(self) -> Sequence['outputs.GetLockboxSecretCurrentVersionResult']:
        return pulumi.get(self, "current_versions")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> builtins.bool:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> builtins.str:
        return pulumi.get(self, "folder_id")

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
    @pulumi.getter(name="passwordPayloadSpecifications")
    def password_payload_specifications(self) -> Sequence['outputs.GetLockboxSecretPasswordPayloadSpecificationResult']:
        return pulumi.get(self, "password_payload_specifications")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        return pulumi.get(self, "status")


class AwaitableGetLockboxSecretResult(GetLockboxSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLockboxSecretResult(
            created_at=self.created_at,
            current_versions=self.current_versions,
            deletion_protection=self.deletion_protection,
            description=self.description,
            folder_id=self.folder_id,
            id=self.id,
            kms_key_id=self.kms_key_id,
            labels=self.labels,
            name=self.name,
            password_payload_specifications=self.password_payload_specifications,
            secret_id=self.secret_id,
            status=self.status)


def get_lockbox_secret(folder_id: Optional[builtins.str] = None,
                       name: Optional[builtins.str] = None,
                       secret_id: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLockboxSecretResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['secretId'] = secret_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getLockboxSecret:getLockboxSecret', __args__, opts=opts, typ=GetLockboxSecretResult).value

    return AwaitableGetLockboxSecretResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        current_versions=pulumi.get(__ret__, 'current_versions'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        description=pulumi.get(__ret__, 'description'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        id=pulumi.get(__ret__, 'id'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        password_payload_specifications=pulumi.get(__ret__, 'password_payload_specifications'),
        secret_id=pulumi.get(__ret__, 'secret_id'),
        status=pulumi.get(__ret__, 'status'))
def get_lockbox_secret_output(folder_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              secret_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLockboxSecretResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['secretId'] = secret_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getLockboxSecret:getLockboxSecret', __args__, opts=opts, typ=GetLockboxSecretResult)
    return __ret__.apply(lambda __response__: GetLockboxSecretResult(
        created_at=pulumi.get(__response__, 'created_at'),
        current_versions=pulumi.get(__response__, 'current_versions'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        description=pulumi.get(__response__, 'description'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        id=pulumi.get(__response__, 'id'),
        kms_key_id=pulumi.get(__response__, 'kms_key_id'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        password_payload_specifications=pulumi.get(__response__, 'password_payload_specifications'),
        secret_id=pulumi.get(__response__, 'secret_id'),
        status=pulumi.get(__response__, 'status')))

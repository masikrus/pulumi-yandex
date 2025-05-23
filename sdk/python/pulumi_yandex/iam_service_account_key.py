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

__all__ = ['IamServiceAccountKeyArgs', 'IamServiceAccountKey']

@pulumi.input_type
class IamServiceAccountKeyArgs:
    def __init__(__self__, *,
                 service_account_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 format: Optional[pulumi.Input[builtins.str]] = None,
                 key_algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 output_to_lockbox: Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']] = None,
                 pgp_key: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a IamServiceAccountKey resource.
        :param pulumi.Input[builtins.str] service_account_id: ID of the service account to create a pair for.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] format: The output format of the keys. `PEM_FILE` is the default format.
        :param pulumi.Input[builtins.str] key_algorithm: The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
               reference](https://yandex.cloud/docs/iam/api-ref/Key).
        :param pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs'] output_to_lockbox: option to create a Lockbox secret version from sensitive outputs
        :param pulumi.Input[builtins.str] pgp_key: An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
               keybase username in the form `keybase:keybaseusername`.
        """
        pulumi.set(__self__, "service_account_id", service_account_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if output_to_lockbox is not None:
            pulumi.set(__self__, "output_to_lockbox", output_to_lockbox)
        if pgp_key is not None:
            pulumi.set(__self__, "pgp_key", pgp_key)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the service account to create a pair for.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The output format of the keys. `PEM_FILE` is the default format.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
        reference](https://yandex.cloud/docs/iam/api-ref/Key).
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="outputToLockbox")
    def output_to_lockbox(self) -> Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']]:
        """
        option to create a Lockbox secret version from sensitive outputs
        """
        return pulumi.get(self, "output_to_lockbox")

    @output_to_lockbox.setter
    def output_to_lockbox(self, value: Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']]):
        pulumi.set(self, "output_to_lockbox", value)

    @property
    @pulumi.getter(name="pgpKey")
    def pgp_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
        keybase username in the form `keybase:keybaseusername`.
        """
        return pulumi.get(self, "pgp_key")

    @pgp_key.setter
    def pgp_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pgp_key", value)


@pulumi.input_type
class _IamServiceAccountKeyState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 encrypted_private_key: Optional[pulumi.Input[builtins.str]] = None,
                 format: Optional[pulumi.Input[builtins.str]] = None,
                 key_algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 key_fingerprint: Optional[pulumi.Input[builtins.str]] = None,
                 output_to_lockbox: Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']] = None,
                 output_to_lockbox_version_id: Optional[pulumi.Input[builtins.str]] = None,
                 pgp_key: Optional[pulumi.Input[builtins.str]] = None,
                 private_key: Optional[pulumi.Input[builtins.str]] = None,
                 public_key: Optional[pulumi.Input[builtins.str]] = None,
                 service_account_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IamServiceAccountKey resources.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] encrypted_private_key: The encrypted private key, base64 encoded. This is only populated when `pgp_key` is supplied.
        :param pulumi.Input[builtins.str] format: The output format of the keys. `PEM_FILE` is the default format.
        :param pulumi.Input[builtins.str] key_algorithm: The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
               reference](https://yandex.cloud/docs/iam/api-ref/Key).
        :param pulumi.Input[builtins.str] key_fingerprint: The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgp_key` is supplied.
        :param pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs'] output_to_lockbox: option to create a Lockbox secret version from sensitive outputs
        :param pulumi.Input[builtins.str] output_to_lockbox_version_id: ID of the Lockbox secret version that contains the value of `secret_key`. This is only populated when
               `output_to_lockbox` is supplied. This version will be destroyed when the IAM key is destroyed, or when
               `output_to_lockbox` is removed.
        :param pulumi.Input[builtins.str] pgp_key: An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
               keybase username in the form `keybase:keybaseusername`.
        :param pulumi.Input[builtins.str] private_key: The private key. This is only populated when neither `pgp_key` nor `output_to_lockbox` are provided.
        :param pulumi.Input[builtins.str] public_key: The public key.
        :param pulumi.Input[builtins.str] service_account_id: ID of the service account to create a pair for.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encrypted_private_key is not None:
            pulumi.set(__self__, "encrypted_private_key", encrypted_private_key)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if key_fingerprint is not None:
            pulumi.set(__self__, "key_fingerprint", key_fingerprint)
        if output_to_lockbox is not None:
            pulumi.set(__self__, "output_to_lockbox", output_to_lockbox)
        if output_to_lockbox_version_id is not None:
            pulumi.set(__self__, "output_to_lockbox_version_id", output_to_lockbox_version_id)
        if pgp_key is not None:
            pulumi.set(__self__, "pgp_key", pgp_key)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The creation timestamp of the resource.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptedPrivateKey")
    def encrypted_private_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The encrypted private key, base64 encoded. This is only populated when `pgp_key` is supplied.
        """
        return pulumi.get(self, "encrypted_private_key")

    @encrypted_private_key.setter
    def encrypted_private_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "encrypted_private_key", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The output format of the keys. `PEM_FILE` is the default format.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
        reference](https://yandex.cloud/docs/iam/api-ref/Key).
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="keyFingerprint")
    def key_fingerprint(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgp_key` is supplied.
        """
        return pulumi.get(self, "key_fingerprint")

    @key_fingerprint.setter
    def key_fingerprint(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_fingerprint", value)

    @property
    @pulumi.getter(name="outputToLockbox")
    def output_to_lockbox(self) -> Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']]:
        """
        option to create a Lockbox secret version from sensitive outputs
        """
        return pulumi.get(self, "output_to_lockbox")

    @output_to_lockbox.setter
    def output_to_lockbox(self, value: Optional[pulumi.Input['IamServiceAccountKeyOutputToLockboxArgs']]):
        pulumi.set(self, "output_to_lockbox", value)

    @property
    @pulumi.getter(name="outputToLockboxVersionId")
    def output_to_lockbox_version_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the Lockbox secret version that contains the value of `secret_key`. This is only populated when
        `output_to_lockbox` is supplied. This version will be destroyed when the IAM key is destroyed, or when
        `output_to_lockbox` is removed.
        """
        return pulumi.get(self, "output_to_lockbox_version_id")

    @output_to_lockbox_version_id.setter
    def output_to_lockbox_version_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "output_to_lockbox_version_id", value)

    @property
    @pulumi.getter(name="pgpKey")
    def pgp_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
        keybase username in the form `keybase:keybaseusername`.
        """
        return pulumi.get(self, "pgp_key")

    @pgp_key.setter
    def pgp_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pgp_key", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The private key. This is only populated when neither `pgp_key` nor `output_to_lockbox` are provided.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The public key.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the service account to create a pair for.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_account_id", value)


@pulumi.type_token("yandex:index/iamServiceAccountKey:IamServiceAccountKey")
class IamServiceAccountKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 format: Optional[pulumi.Input[builtins.str]] = None,
                 key_algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 output_to_lockbox: Optional[pulumi.Input[Union['IamServiceAccountKeyOutputToLockboxArgs', 'IamServiceAccountKeyOutputToLockboxArgsDict']]] = None,
                 pgp_key: Optional[pulumi.Input[builtins.str]] = None,
                 service_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a IamServiceAccountKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] format: The output format of the keys. `PEM_FILE` is the default format.
        :param pulumi.Input[builtins.str] key_algorithm: The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
               reference](https://yandex.cloud/docs/iam/api-ref/Key).
        :param pulumi.Input[Union['IamServiceAccountKeyOutputToLockboxArgs', 'IamServiceAccountKeyOutputToLockboxArgsDict']] output_to_lockbox: option to create a Lockbox secret version from sensitive outputs
        :param pulumi.Input[builtins.str] pgp_key: An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
               keybase username in the form `keybase:keybaseusername`.
        :param pulumi.Input[builtins.str] service_account_id: ID of the service account to create a pair for.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamServiceAccountKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IamServiceAccountKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IamServiceAccountKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamServiceAccountKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 format: Optional[pulumi.Input[builtins.str]] = None,
                 key_algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 output_to_lockbox: Optional[pulumi.Input[Union['IamServiceAccountKeyOutputToLockboxArgs', 'IamServiceAccountKeyOutputToLockboxArgsDict']]] = None,
                 pgp_key: Optional[pulumi.Input[builtins.str]] = None,
                 service_account_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamServiceAccountKeyArgs.__new__(IamServiceAccountKeyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["format"] = format
            __props__.__dict__["key_algorithm"] = key_algorithm
            __props__.__dict__["output_to_lockbox"] = output_to_lockbox
            __props__.__dict__["pgp_key"] = pgp_key
            if service_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_id'")
            __props__.__dict__["service_account_id"] = service_account_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["encrypted_private_key"] = None
            __props__.__dict__["key_fingerprint"] = None
            __props__.__dict__["output_to_lockbox_version_id"] = None
            __props__.__dict__["private_key"] = None
            __props__.__dict__["public_key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IamServiceAccountKey, __self__).__init__(
            'yandex:index/iamServiceAccountKey:IamServiceAccountKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            encrypted_private_key: Optional[pulumi.Input[builtins.str]] = None,
            format: Optional[pulumi.Input[builtins.str]] = None,
            key_algorithm: Optional[pulumi.Input[builtins.str]] = None,
            key_fingerprint: Optional[pulumi.Input[builtins.str]] = None,
            output_to_lockbox: Optional[pulumi.Input[Union['IamServiceAccountKeyOutputToLockboxArgs', 'IamServiceAccountKeyOutputToLockboxArgsDict']]] = None,
            output_to_lockbox_version_id: Optional[pulumi.Input[builtins.str]] = None,
            pgp_key: Optional[pulumi.Input[builtins.str]] = None,
            private_key: Optional[pulumi.Input[builtins.str]] = None,
            public_key: Optional[pulumi.Input[builtins.str]] = None,
            service_account_id: Optional[pulumi.Input[builtins.str]] = None) -> 'IamServiceAccountKey':
        """
        Get an existing IamServiceAccountKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[builtins.str] encrypted_private_key: The encrypted private key, base64 encoded. This is only populated when `pgp_key` is supplied.
        :param pulumi.Input[builtins.str] format: The output format of the keys. `PEM_FILE` is the default format.
        :param pulumi.Input[builtins.str] key_algorithm: The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
               reference](https://yandex.cloud/docs/iam/api-ref/Key).
        :param pulumi.Input[builtins.str] key_fingerprint: The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgp_key` is supplied.
        :param pulumi.Input[Union['IamServiceAccountKeyOutputToLockboxArgs', 'IamServiceAccountKeyOutputToLockboxArgsDict']] output_to_lockbox: option to create a Lockbox secret version from sensitive outputs
        :param pulumi.Input[builtins.str] output_to_lockbox_version_id: ID of the Lockbox secret version that contains the value of `secret_key`. This is only populated when
               `output_to_lockbox` is supplied. This version will be destroyed when the IAM key is destroyed, or when
               `output_to_lockbox` is removed.
        :param pulumi.Input[builtins.str] pgp_key: An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
               keybase username in the form `keybase:keybaseusername`.
        :param pulumi.Input[builtins.str] private_key: The private key. This is only populated when neither `pgp_key` nor `output_to_lockbox` are provided.
        :param pulumi.Input[builtins.str] public_key: The public key.
        :param pulumi.Input[builtins.str] service_account_id: ID of the service account to create a pair for.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamServiceAccountKeyState.__new__(_IamServiceAccountKeyState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["encrypted_private_key"] = encrypted_private_key
        __props__.__dict__["format"] = format
        __props__.__dict__["key_algorithm"] = key_algorithm
        __props__.__dict__["key_fingerprint"] = key_fingerprint
        __props__.__dict__["output_to_lockbox"] = output_to_lockbox
        __props__.__dict__["output_to_lockbox_version_id"] = output_to_lockbox_version_id
        __props__.__dict__["pgp_key"] = pgp_key
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["service_account_id"] = service_account_id
        return IamServiceAccountKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The creation timestamp of the resource.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptedPrivateKey")
    def encrypted_private_key(self) -> pulumi.Output[builtins.str]:
        """
        The encrypted private key, base64 encoded. This is only populated when `pgp_key` is supplied.
        """
        return pulumi.get(self, "encrypted_private_key")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The output format of the keys. `PEM_FILE` is the default format.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The algorithm used to generate the key. `RSA_2048` is the default algorithm. Valid values are listed in the [API
        reference](https://yandex.cloud/docs/iam/api-ref/Key).
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter(name="keyFingerprint")
    def key_fingerprint(self) -> pulumi.Output[builtins.str]:
        """
        The fingerprint of the PGP key used to encrypt the private key. This is only populated when `pgp_key` is supplied.
        """
        return pulumi.get(self, "key_fingerprint")

    @property
    @pulumi.getter(name="outputToLockbox")
    def output_to_lockbox(self) -> pulumi.Output[Optional['outputs.IamServiceAccountKeyOutputToLockbox']]:
        """
        option to create a Lockbox secret version from sensitive outputs
        """
        return pulumi.get(self, "output_to_lockbox")

    @property
    @pulumi.getter(name="outputToLockboxVersionId")
    def output_to_lockbox_version_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the Lockbox secret version that contains the value of `secret_key`. This is only populated when
        `output_to_lockbox` is supplied. This version will be destroyed when the IAM key is destroyed, or when
        `output_to_lockbox` is removed.
        """
        return pulumi.get(self, "output_to_lockbox_version_id")

    @property
    @pulumi.getter(name="pgpKey")
    def pgp_key(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a
        keybase username in the form `keybase:keybaseusername`.
        """
        return pulumi.get(self, "pgp_key")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[builtins.str]:
        """
        The private key. This is only populated when neither `pgp_key` nor `output_to_lockbox` are provided.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[builtins.str]:
        """
        The public key.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the service account to create a pair for.
        """
        return pulumi.get(self, "service_account_id")


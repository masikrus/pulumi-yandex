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

__all__ = ['KmsSecretCiphertextArgs', 'KmsSecretCiphertext']

@pulumi.input_type
class KmsSecretCiphertextArgs:
    def __init__(__self__, *,
                 key_id: pulumi.Input[builtins.str],
                 plaintext: pulumi.Input[builtins.str],
                 aad_context: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a KmsSecretCiphertext resource.
        :param pulumi.Input[builtins.str] key_id: ID of the symmetric KMS key to use for encryption.
        :param pulumi.Input[builtins.str] plaintext: Plaintext to be encrypted.
        :param pulumi.Input[builtins.str] aad_context: Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
               `SymmetricDecryptRequest`.
        """
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "plaintext", plaintext)
        if aad_context is not None:
            pulumi.set(__self__, "aad_context", aad_context)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the symmetric KMS key to use for encryption.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def plaintext(self) -> pulumi.Input[builtins.str]:
        """
        Plaintext to be encrypted.
        """
        return pulumi.get(self, "plaintext")

    @plaintext.setter
    def plaintext(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "plaintext", value)

    @property
    @pulumi.getter(name="aadContext")
    def aad_context(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
        `SymmetricDecryptRequest`.
        """
        return pulumi.get(self, "aad_context")

    @aad_context.setter
    def aad_context(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "aad_context", value)


@pulumi.input_type
class _KmsSecretCiphertextState:
    def __init__(__self__, *,
                 aad_context: Optional[pulumi.Input[builtins.str]] = None,
                 ciphertext: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 plaintext: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering KmsSecretCiphertext resources.
        :param pulumi.Input[builtins.str] aad_context: Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
               `SymmetricDecryptRequest`.
        :param pulumi.Input[builtins.str] ciphertext: Resulting CipherText, encoded with `standard` base64 alphabet as defined in RFC 4648 section 4.
        :param pulumi.Input[builtins.str] key_id: ID of the symmetric KMS key to use for encryption.
        :param pulumi.Input[builtins.str] plaintext: Plaintext to be encrypted.
        """
        if aad_context is not None:
            pulumi.set(__self__, "aad_context", aad_context)
        if ciphertext is not None:
            pulumi.set(__self__, "ciphertext", ciphertext)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if plaintext is not None:
            pulumi.set(__self__, "plaintext", plaintext)

    @property
    @pulumi.getter(name="aadContext")
    def aad_context(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
        `SymmetricDecryptRequest`.
        """
        return pulumi.get(self, "aad_context")

    @aad_context.setter
    def aad_context(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "aad_context", value)

    @property
    @pulumi.getter
    def ciphertext(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resulting CipherText, encoded with `standard` base64 alphabet as defined in RFC 4648 section 4.
        """
        return pulumi.get(self, "ciphertext")

    @ciphertext.setter
    def ciphertext(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ciphertext", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the symmetric KMS key to use for encryption.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def plaintext(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Plaintext to be encrypted.
        """
        return pulumi.get(self, "plaintext")

    @plaintext.setter
    def plaintext(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "plaintext", value)


@pulumi.type_token("yandex:index/kmsSecretCiphertext:KmsSecretCiphertext")
class KmsSecretCiphertext(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aad_context: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 plaintext: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a KmsSecretCiphertext resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aad_context: Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
               `SymmetricDecryptRequest`.
        :param pulumi.Input[builtins.str] key_id: ID of the symmetric KMS key to use for encryption.
        :param pulumi.Input[builtins.str] plaintext: Plaintext to be encrypted.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KmsSecretCiphertextArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a KmsSecretCiphertext resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param KmsSecretCiphertextArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KmsSecretCiphertextArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aad_context: Optional[pulumi.Input[builtins.str]] = None,
                 key_id: Optional[pulumi.Input[builtins.str]] = None,
                 plaintext: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KmsSecretCiphertextArgs.__new__(KmsSecretCiphertextArgs)

            __props__.__dict__["aad_context"] = aad_context
            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            if plaintext is None and not opts.urn:
                raise TypeError("Missing required property 'plaintext'")
            __props__.__dict__["plaintext"] = None if plaintext is None else pulumi.Output.secret(plaintext)
            __props__.__dict__["ciphertext"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["plaintext"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(KmsSecretCiphertext, __self__).__init__(
            'yandex:index/kmsSecretCiphertext:KmsSecretCiphertext',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aad_context: Optional[pulumi.Input[builtins.str]] = None,
            ciphertext: Optional[pulumi.Input[builtins.str]] = None,
            key_id: Optional[pulumi.Input[builtins.str]] = None,
            plaintext: Optional[pulumi.Input[builtins.str]] = None) -> 'KmsSecretCiphertext':
        """
        Get an existing KmsSecretCiphertext resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] aad_context: Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
               `SymmetricDecryptRequest`.
        :param pulumi.Input[builtins.str] ciphertext: Resulting CipherText, encoded with `standard` base64 alphabet as defined in RFC 4648 section 4.
        :param pulumi.Input[builtins.str] key_id: ID of the symmetric KMS key to use for encryption.
        :param pulumi.Input[builtins.str] plaintext: Plaintext to be encrypted.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KmsSecretCiphertextState.__new__(_KmsSecretCiphertextState)

        __props__.__dict__["aad_context"] = aad_context
        __props__.__dict__["ciphertext"] = ciphertext
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["plaintext"] = plaintext
        return KmsSecretCiphertext(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aadContext")
    def aad_context(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the
        `SymmetricDecryptRequest`.
        """
        return pulumi.get(self, "aad_context")

    @property
    @pulumi.getter
    def ciphertext(self) -> pulumi.Output[builtins.str]:
        """
        Resulting CipherText, encoded with `standard` base64 alphabet as defined in RFC 4648 section 4.
        """
        return pulumi.get(self, "ciphertext")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the symmetric KMS key to use for encryption.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def plaintext(self) -> pulumi.Output[builtins.str]:
        """
        Plaintext to be encrypted.
        """
        return pulumi.get(self, "plaintext")


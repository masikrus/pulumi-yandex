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

__all__ = ['VpcSecurityGroupArgs', 'VpcSecurityGroup']

@pulumi.input_type
class VpcSecurityGroupArgs:
    def __init__(__self__, *,
                 network_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a VpcSecurityGroup resource.
        :param pulumi.Input[builtins.str] network_id: ID of the network this security group belongs to.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]] egresses: A list of egress rules.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]] ingresses: A list of ingress rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        """
        pulumi.set(__self__, "network_id", network_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the network this security group belongs to.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "network_id", value)

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
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]]:
        """
        A list of egress rules.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]]:
        """
        A list of ingress rules.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _VpcSecurityGroupState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering VpcSecurityGroup resources.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]] egresses: A list of egress rules.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]] ingresses: A list of ingress rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this security group belongs to.
        :param pulumi.Input[builtins.str] status: Status of this security group.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

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
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]]:
        """
        A list of egress rules.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]]:
        """
        A list of ingress rules.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcSecurityGroupIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the network this security group belongs to.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Status of this security group.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


@pulumi.type_token("yandex:index/vpcSecurityGroup:VpcSecurityGroup")
class VpcSecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupEgressArgs', 'VpcSecurityGroupEgressArgsDict']]]]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupIngressArgs', 'VpcSecurityGroupIngressArgsDict']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a VpcSecurityGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupEgressArgs', 'VpcSecurityGroupEgressArgsDict']]]] egresses: A list of egress rules.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupIngressArgs', 'VpcSecurityGroupIngressArgsDict']]]] ingresses: A list of ingress rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this security group belongs to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcSecurityGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpcSecurityGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcSecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcSecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupEgressArgs', 'VpcSecurityGroupEgressArgsDict']]]]] = None,
                 folder_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupIngressArgs', 'VpcSecurityGroupIngressArgsDict']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcSecurityGroupArgs.__new__(VpcSecurityGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["egresses"] = egresses
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["ingresses"] = ingresses
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
        super(VpcSecurityGroup, __self__).__init__(
            'yandex:index/vpcSecurityGroup:VpcSecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            egresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupEgressArgs', 'VpcSecurityGroupEgressArgsDict']]]]] = None,
            folder_id: Optional[pulumi.Input[builtins.str]] = None,
            ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupIngressArgs', 'VpcSecurityGroupIngressArgsDict']]]]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            network_id: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'VpcSecurityGroup':
        """
        Get an existing VpcSecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The creation timestamp of the resource.
        :param pulumi.Input[builtins.str] description: The resource description.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupEgressArgs', 'VpcSecurityGroupEgressArgsDict']]]] egresses: A list of egress rules.
        :param pulumi.Input[builtins.str] folder_id: The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VpcSecurityGroupIngressArgs', 'VpcSecurityGroupIngressArgsDict']]]] ingresses: A list of ingress rules.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] labels: A set of key/value label pairs which assigned to resource.
        :param pulumi.Input[builtins.str] name: The resource name.
        :param pulumi.Input[builtins.str] network_id: ID of the network this security group belongs to.
        :param pulumi.Input[builtins.str] status: Status of this security group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcSecurityGroupState.__new__(_VpcSecurityGroupState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["egresses"] = egresses
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["ingresses"] = ingresses
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["status"] = status
        return VpcSecurityGroup(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter
    def egresses(self) -> pulumi.Output[Sequence['outputs.VpcSecurityGroupEgress']]:
        """
        A list of egress rules.
        """
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[builtins.str]:
        """
        The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def ingresses(self) -> pulumi.Output[Sequence['outputs.VpcSecurityGroupIngress']]:
        """
        A list of ingress rules.
        """
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        A set of key/value label pairs which assigned to resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the network this security group belongs to.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Status of this security group.
        """
        return pulumi.get(self, "status")


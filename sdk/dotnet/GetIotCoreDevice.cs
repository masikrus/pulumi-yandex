// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetIotCoreDevice
    {
        public static Task<GetIotCoreDeviceResult> InvokeAsync(GetIotCoreDeviceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIotCoreDeviceResult>("yandex:index/getIotCoreDevice:getIotCoreDevice", args ?? new GetIotCoreDeviceArgs(), options.WithDefaults());

        public static Output<GetIotCoreDeviceResult> Invoke(GetIotCoreDeviceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotCoreDeviceResult>("yandex:index/getIotCoreDevice:getIotCoreDevice", args ?? new GetIotCoreDeviceInvokeArgs(), options.WithDefaults());

        public static Output<GetIotCoreDeviceResult> Invoke(GetIotCoreDeviceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotCoreDeviceResult>("yandex:index/getIotCoreDevice:getIotCoreDevice", args ?? new GetIotCoreDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIotCoreDeviceArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceId")]
        public string? DeviceId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetIotCoreDeviceArgs()
        {
        }
        public static new GetIotCoreDeviceArgs Empty => new GetIotCoreDeviceArgs();
    }

    public sealed class GetIotCoreDeviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetIotCoreDeviceInvokeArgs()
        {
        }
        public static new GetIotCoreDeviceInvokeArgs Empty => new GetIotCoreDeviceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIotCoreDeviceResult
    {
        public readonly ImmutableDictionary<string, string> Aliases;
        public readonly ImmutableArray<string> Certificates;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string? DeviceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string? Name;
        public readonly ImmutableArray<string> Passwords;
        public readonly string RegistryId;

        [OutputConstructor]
        private GetIotCoreDeviceResult(
            ImmutableDictionary<string, string> aliases,

            ImmutableArray<string> certificates,

            string createdAt,

            string description,

            string? deviceId,

            string id,

            ImmutableDictionary<string, string> labels,

            string? name,

            ImmutableArray<string> passwords,

            string registryId)
        {
            Aliases = aliases;
            Certificates = certificates;
            CreatedAt = createdAt;
            Description = description;
            DeviceId = deviceId;
            Id = id;
            Labels = labels;
            Name = name;
            Passwords = passwords;
            RegistryId = registryId;
        }
    }
}

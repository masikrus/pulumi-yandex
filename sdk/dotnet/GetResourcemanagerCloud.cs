// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetResourcemanagerCloud
    {
        public static Task<GetResourcemanagerCloudResult> InvokeAsync(GetResourcemanagerCloudArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourcemanagerCloudResult>("yandex:index/getResourcemanagerCloud:getResourcemanagerCloud", args ?? new GetResourcemanagerCloudArgs(), options.WithDefaults());

        public static Output<GetResourcemanagerCloudResult> Invoke(GetResourcemanagerCloudInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourcemanagerCloudResult>("yandex:index/getResourcemanagerCloud:getResourcemanagerCloud", args ?? new GetResourcemanagerCloudInvokeArgs(), options.WithDefaults());

        public static Output<GetResourcemanagerCloudResult> Invoke(GetResourcemanagerCloudInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourcemanagerCloudResult>("yandex:index/getResourcemanagerCloud:getResourcemanagerCloud", args ?? new GetResourcemanagerCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcemanagerCloudArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudId")]
        public string? CloudId { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetResourcemanagerCloudArgs()
        {
        }
        public static new GetResourcemanagerCloudArgs Empty => new GetResourcemanagerCloudArgs();
    }

    public sealed class GetResourcemanagerCloudInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetResourcemanagerCloudInvokeArgs()
        {
        }
        public static new GetResourcemanagerCloudInvokeArgs Empty => new GetResourcemanagerCloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourcemanagerCloudResult
    {
        public readonly string CloudId;
        public readonly string CreatedAt;
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetResourcemanagerCloudResult(
            string cloudId,

            string createdAt,

            string? description,

            string id,

            string name)
        {
            CloudId = cloudId;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
        }
    }
}

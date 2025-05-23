// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetServerlessEventrouterBus
    {
        public static Task<GetServerlessEventrouterBusResult> InvokeAsync(GetServerlessEventrouterBusArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerlessEventrouterBusResult>("yandex:index/getServerlessEventrouterBus:getServerlessEventrouterBus", args ?? new GetServerlessEventrouterBusArgs(), options.WithDefaults());

        public static Output<GetServerlessEventrouterBusResult> Invoke(GetServerlessEventrouterBusInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerlessEventrouterBusResult>("yandex:index/getServerlessEventrouterBus:getServerlessEventrouterBus", args ?? new GetServerlessEventrouterBusInvokeArgs(), options.WithDefaults());

        public static Output<GetServerlessEventrouterBusResult> Invoke(GetServerlessEventrouterBusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerlessEventrouterBusResult>("yandex:index/getServerlessEventrouterBus:getServerlessEventrouterBus", args ?? new GetServerlessEventrouterBusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerlessEventrouterBusArgs : global::Pulumi.InvokeArgs
    {
        [Input("busId")]
        public string? BusId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetServerlessEventrouterBusArgs()
        {
        }
        public static new GetServerlessEventrouterBusArgs Empty => new GetServerlessEventrouterBusArgs();
    }

    public sealed class GetServerlessEventrouterBusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("busId")]
        public Input<string>? BusId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetServerlessEventrouterBusInvokeArgs()
        {
        }
        public static new GetServerlessEventrouterBusInvokeArgs Empty => new GetServerlessEventrouterBusInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerlessEventrouterBusResult
    {
        public readonly string? BusId;
        public readonly string CloudId;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string? Name;

        [OutputConstructor]
        private GetServerlessEventrouterBusResult(
            string? busId,

            string cloudId,

            string createdAt,

            bool deletionProtection,

            string description,

            string folderId,

            string id,

            ImmutableDictionary<string, string> labels,

            string? name)
        {
            BusId = busId;
            CloudId = cloudId;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            FolderId = folderId;
            Id = id;
            Labels = labels;
            Name = name;
        }
    }
}

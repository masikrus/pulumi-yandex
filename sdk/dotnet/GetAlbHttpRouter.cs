// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetAlbHttpRouter
    {
        public static Task<GetAlbHttpRouterResult> InvokeAsync(GetAlbHttpRouterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlbHttpRouterResult>("yandex:index/getAlbHttpRouter:getAlbHttpRouter", args ?? new GetAlbHttpRouterArgs(), options.WithDefaults());

        public static Output<GetAlbHttpRouterResult> Invoke(GetAlbHttpRouterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlbHttpRouterResult>("yandex:index/getAlbHttpRouter:getAlbHttpRouter", args ?? new GetAlbHttpRouterInvokeArgs(), options.WithDefaults());

        public static Output<GetAlbHttpRouterResult> Invoke(GetAlbHttpRouterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlbHttpRouterResult>("yandex:index/getAlbHttpRouter:getAlbHttpRouter", args ?? new GetAlbHttpRouterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlbHttpRouterArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("httpRouterId")]
        public string? HttpRouterId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetAlbHttpRouterArgs()
        {
        }
        public static new GetAlbHttpRouterArgs Empty => new GetAlbHttpRouterArgs();
    }

    public sealed class GetAlbHttpRouterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("httpRouterId")]
        public Input<string>? HttpRouterId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetAlbHttpRouterInvokeArgs()
        {
        }
        public static new GetAlbHttpRouterInvokeArgs Empty => new GetAlbHttpRouterInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlbHttpRouterResult
    {
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string FolderId;
        public readonly string HttpRouterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetAlbHttpRouterRouteOptionResult> RouteOptions;

        [OutputConstructor]
        private GetAlbHttpRouterResult(
            string createdAt,

            string description,

            string folderId,

            string httpRouterId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.GetAlbHttpRouterRouteOptionResult> routeOptions)
        {
            CreatedAt = createdAt;
            Description = description;
            FolderId = folderId;
            HttpRouterId = httpRouterId;
            Id = id;
            Labels = labels;
            Name = name;
            RouteOptions = routeOptions;
        }
    }
}

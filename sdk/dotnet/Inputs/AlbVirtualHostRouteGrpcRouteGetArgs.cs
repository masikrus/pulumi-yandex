// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRouteGrpcRouteGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("grpcMatches")]
        private InputList<Inputs.AlbVirtualHostRouteGrpcRouteGrpcMatchGetArgs>? _grpcMatches;

        /// <summary>
        /// Checks `/` prefix by default.
        /// </summary>
        public InputList<Inputs.AlbVirtualHostRouteGrpcRouteGrpcMatchGetArgs> GrpcMatches
        {
            get => _grpcMatches ?? (_grpcMatches = new InputList<Inputs.AlbVirtualHostRouteGrpcRouteGrpcMatchGetArgs>());
            set => _grpcMatches = value;
        }

        /// <summary>
        /// gRPC route action resource.
        /// 
        /// &gt; Only one type of host rewrite specifiers `host_rewrite` or `auto_host_rewrite` should be specified.
        /// </summary>
        [Input("grpcRouteAction")]
        public Input<Inputs.AlbVirtualHostRouteGrpcRouteGrpcRouteActionGetArgs>? GrpcRouteAction { get; set; }

        /// <summary>
        /// gRPC status response action resource.
        /// </summary>
        [Input("grpcStatusResponseAction")]
        public Input<Inputs.AlbVirtualHostRouteGrpcRouteGrpcStatusResponseActionGetArgs>? GrpcStatusResponseAction { get; set; }

        public AlbVirtualHostRouteGrpcRouteGetArgs()
        {
        }
        public static new AlbVirtualHostRouteGrpcRouteGetArgs Empty => new AlbVirtualHostRouteGrpcRouteGetArgs();
    }
}

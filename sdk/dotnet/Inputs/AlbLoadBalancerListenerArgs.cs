// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbLoadBalancerListenerArgs : global::Pulumi.ResourceArgs
    {
        [Input("endpoints")]
        private InputList<Inputs.AlbLoadBalancerListenerEndpointArgs>? _endpoints;

        /// <summary>
        /// Network endpoint (addresses and ports) of the listener.
        /// </summary>
        public InputList<Inputs.AlbLoadBalancerListenerEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.AlbLoadBalancerListenerEndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// HTTP handler that sets plain text HTTP router.
        /// </summary>
        [Input("http")]
        public Input<Inputs.AlbLoadBalancerListenerHttpArgs>? Http { get; set; }

        /// <summary>
        /// Name of the listener.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Stream configuration
        /// </summary>
        [Input("stream")]
        public Input<Inputs.AlbLoadBalancerListenerStreamArgs>? Stream { get; set; }

        /// <summary>
        /// TLS configuration
        /// </summary>
        [Input("tls")]
        public Input<Inputs.AlbLoadBalancerListenerTlsArgs>? Tls { get; set; }

        public AlbLoadBalancerListenerArgs()
        {
        }
        public static new AlbLoadBalancerListenerArgs Empty => new AlbLoadBalancerListenerArgs();
    }
}

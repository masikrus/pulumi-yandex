// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetServerlessContainerConnectivityInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        public GetServerlessContainerConnectivityInputArgs()
        {
        }
        public static new GetServerlessContainerConnectivityInputArgs Empty => new GetServerlessContainerConnectivityInputArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbBackendGroupGrpcBackendTlsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [SNI](https://en.wikipedia.org/wiki/Server_Name_Indication) string for TLS connections.
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        [Input("validationContext")]
        public Input<Inputs.AlbBackendGroupGrpcBackendTlsValidationContextGetArgs>? ValidationContext { get; set; }

        public AlbBackendGroupGrpcBackendTlsGetArgs()
        {
        }
        public static new AlbBackendGroupGrpcBackendTlsGetArgs Empty => new AlbBackendGroupGrpcBackendTlsGetArgs();
    }
}

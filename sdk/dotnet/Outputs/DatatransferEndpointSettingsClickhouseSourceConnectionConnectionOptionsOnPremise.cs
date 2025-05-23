// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremise
    {
        public readonly int? HttpPort;
        public readonly int? NativePort;
        public readonly ImmutableArray<Outputs.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseShard> Shards;
        public readonly Outputs.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsMode? TlsMode;

        [OutputConstructor]
        private DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremise(
            int? httpPort,

            int? nativePort,

            ImmutableArray<Outputs.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseShard> shards,

            Outputs.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsOnPremiseTlsMode? tlsMode)
        {
            HttpPort = httpPort;
            NativePort = nativePort;
            Shards = shards;
            TlsMode = tlsMode;
        }
    }
}

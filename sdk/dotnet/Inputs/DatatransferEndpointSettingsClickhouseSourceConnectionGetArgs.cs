// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsClickhouseSourceConnectionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectionOptions")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseSourceConnectionConnectionOptionsGetArgs>? ConnectionOptions { get; set; }

        public DatatransferEndpointSettingsClickhouseSourceConnectionGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsClickhouseSourceConnectionGetArgs Empty => new DatatransferEndpointSettingsClickhouseSourceConnectionGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeArgs : global::Pulumi.ResourceArgs
    {
        [Input("disabled")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeDisabledArgs>? Disabled { get; set; }

        [Input("enabled")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeEnabledArgs>? Enabled { get; set; }

        public DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeArgs()
        {
        }
        public static new DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeArgs Empty => new DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseTlsModeArgs();
    }
}

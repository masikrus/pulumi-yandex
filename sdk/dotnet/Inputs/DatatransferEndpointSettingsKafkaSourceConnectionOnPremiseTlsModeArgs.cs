// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Empty block designating that the connection is not secured, i.e. plaintext connection.
        /// </summary>
        [Input("disabled")]
        public Input<Inputs.DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeDisabledArgs>? Disabled { get; set; }

        /// <summary>
        /// If this attribute is not an empty block, then TLS is used for the server connection.
        /// </summary>
        [Input("enabled")]
        public Input<Inputs.DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeEnabledArgs>? Enabled { get; set; }

        public DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeArgs()
        {
        }
        public static new DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeArgs Empty => new DatatransferEndpointSettingsKafkaSourceConnectionOnPremiseTlsModeArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsKafkaSourceAuthSaslArgs : global::Pulumi.ResourceArgs
    {
        [Input("mechanism")]
        public Input<string>? Mechanism { get; set; }

        [Input("password")]
        public Input<Inputs.DatatransferEndpointSettingsKafkaSourceAuthSaslPasswordArgs>? Password { get; set; }

        [Input("user")]
        public Input<string>? User { get; set; }

        public DatatransferEndpointSettingsKafkaSourceAuthSaslArgs()
        {
        }
        public static new DatatransferEndpointSettingsKafkaSourceAuthSaslArgs Empty => new DatatransferEndpointSettingsKafkaSourceAuthSaslArgs();
    }
}

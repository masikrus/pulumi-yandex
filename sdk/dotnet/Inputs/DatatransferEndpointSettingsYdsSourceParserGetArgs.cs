// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsYdsSourceParserGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parse Audit Trails data. Empty struct.
        /// </summary>
        [Input("auditTrailsV1Parser")]
        public Input<Inputs.DatatransferEndpointSettingsYdsSourceParserAuditTrailsV1ParserGetArgs>? AuditTrailsV1Parser { get; set; }

        /// <summary>
        /// Parse Cloud Logging data. Empty struct.
        /// </summary>
        [Input("cloudLoggingParser")]
        public Input<Inputs.DatatransferEndpointSettingsYdsSourceParserCloudLoggingParserGetArgs>? CloudLoggingParser { get; set; }

        /// <summary>
        /// Parse data in json format.
        /// </summary>
        [Input("jsonParser")]
        public Input<Inputs.DatatransferEndpointSettingsYdsSourceParserJsonParserGetArgs>? JsonParser { get; set; }

        [Input("tskvParser")]
        public Input<Inputs.DatatransferEndpointSettingsYdsSourceParserTskvParserGetArgs>? TskvParser { get; set; }

        public DatatransferEndpointSettingsYdsSourceParserGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsYdsSourceParserGetArgs Empty => new DatatransferEndpointSettingsYdsSourceParserGetArgs();
    }
}

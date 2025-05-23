// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsKafkaSourceParserJsonParserDataSchemaGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("fields")]
        public Input<Inputs.DatatransferEndpointSettingsKafkaSourceParserJsonParserDataSchemaFieldsGetArgs>? Fields { get; set; }

        /// <summary>
        /// Description of the data schema as JSON specification.
        /// </summary>
        [Input("jsonFields")]
        public Input<string>? JsonFields { get; set; }

        public DatatransferEndpointSettingsKafkaSourceParserJsonParserDataSchemaGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsKafkaSourceParserJsonParserDataSchemaGetArgs Empty => new DatatransferEndpointSettingsKafkaSourceParserJsonParserDataSchemaGetArgs();
    }
}

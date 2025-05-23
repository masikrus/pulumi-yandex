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
    public sealed class GetMdbKafkaConnectorConnectorConfigS3SinkResult
    {
        /// <summary>
        /// Compression type for messages. Cannot be changed.
        /// </summary>
        public readonly string FileCompressionType;
        /// <summary>
        /// Max records per file.
        /// </summary>
        public readonly int FileMaxRecords;
        /// <summary>
        /// Settings for connection to s3-compatible storage.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionResult> S3Connections;
        /// <summary>
        /// The pattern for topic names to be copied to s3 bucket.
        /// </summary>
        public readonly string Topics;

        [OutputConstructor]
        private GetMdbKafkaConnectorConnectorConfigS3SinkResult(
            string fileCompressionType,

            int fileMaxRecords,

            ImmutableArray<Outputs.GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionResult> s3Connections,

            string topics)
        {
            FileCompressionType = fileCompressionType;
            FileMaxRecords = fileMaxRecords;
            S3Connections = s3Connections;
            Topics = topics;
        }
    }
}

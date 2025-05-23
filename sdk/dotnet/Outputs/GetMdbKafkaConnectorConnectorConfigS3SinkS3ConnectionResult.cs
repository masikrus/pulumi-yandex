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
    public sealed class GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionResult
    {
        /// <summary>
        /// Name of the bucket in s3-compatible storage.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Connection params for external s3-compatible storage.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionExternalS3Result> ExternalS3s;

        [OutputConstructor]
        private GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionResult(
            string bucketName,

            ImmutableArray<Outputs.GetMdbKafkaConnectorConnectorConfigS3SinkS3ConnectionExternalS3Result> externalS3s)
        {
            BucketName = bucketName;
            ExternalS3s = externalS3s;
        }
    }
}

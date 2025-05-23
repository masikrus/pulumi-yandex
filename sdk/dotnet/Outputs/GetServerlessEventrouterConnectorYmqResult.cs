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
    public sealed class GetServerlessEventrouterConnectorYmqResult
    {
        /// <summary>
        /// Batch size for polling
        /// </summary>
        public readonly int BatchSize;
        /// <summary>
        /// Queue polling timeout
        /// </summary>
        public readonly string PollingTimeout;
        /// <summary>
        /// Required field. Queue ARN. Example: yrn:yc:ymq:ru-central1:aoe***:test
        /// </summary>
        public readonly string QueueArn;
        /// <summary>
        /// Service account which has read access to the queue
        /// </summary>
        public readonly string ServiceAccountId;
        /// <summary>
        /// Queue visibility timeout override
        /// </summary>
        public readonly string VisibilityTimeout;

        [OutputConstructor]
        private GetServerlessEventrouterConnectorYmqResult(
            int batchSize,

            string pollingTimeout,

            string queueArn,

            string serviceAccountId,

            string visibilityTimeout)
        {
            BatchSize = batchSize;
            PollingTimeout = pollingTimeout;
            QueueArn = queueArn;
            ServiceAccountId = serviceAccountId;
            VisibilityTimeout = visibilityTimeout;
        }
    }
}

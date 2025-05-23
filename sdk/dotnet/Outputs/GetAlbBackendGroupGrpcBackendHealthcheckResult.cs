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
    public sealed class GetAlbBackendGroupGrpcBackendHealthcheckResult
    {
        public readonly Outputs.GetAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckResult GrpcHealthcheck;
        public readonly int HealthcheckPort;
        public readonly int HealthyThreshold;
        public readonly Outputs.GetAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckResult HttpHealthcheck;
        public readonly string Interval;
        public readonly double IntervalJitterPercent;
        public readonly Outputs.GetAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheckResult StreamHealthcheck;
        public readonly string Timeout;
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private GetAlbBackendGroupGrpcBackendHealthcheckResult(
            Outputs.GetAlbBackendGroupGrpcBackendHealthcheckGrpcHealthcheckResult grpcHealthcheck,

            int healthcheckPort,

            int healthyThreshold,

            Outputs.GetAlbBackendGroupGrpcBackendHealthcheckHttpHealthcheckResult httpHealthcheck,

            string interval,

            double intervalJitterPercent,

            Outputs.GetAlbBackendGroupGrpcBackendHealthcheckStreamHealthcheckResult streamHealthcheck,

            string timeout,

            int unhealthyThreshold)
        {
            GrpcHealthcheck = grpcHealthcheck;
            HealthcheckPort = healthcheckPort;
            HealthyThreshold = healthyThreshold;
            HttpHealthcheck = httpHealthcheck;
            Interval = interval;
            IntervalJitterPercent = intervalJitterPercent;
            StreamHealthcheck = streamHealthcheck;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}

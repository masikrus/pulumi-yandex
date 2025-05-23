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
    public sealed class GetMdbMysqlUserConnectionLimitResult
    {
        public readonly int MaxConnectionsPerHour;
        public readonly int MaxQuestionsPerHour;
        public readonly int MaxUpdatesPerHour;
        public readonly int MaxUserConnections;

        [OutputConstructor]
        private GetMdbMysqlUserConnectionLimitResult(
            int maxConnectionsPerHour,

            int maxQuestionsPerHour,

            int maxUpdatesPerHour,

            int maxUserConnections)
        {
            MaxConnectionsPerHour = maxConnectionsPerHour;
            MaxQuestionsPerHour = maxQuestionsPerHour;
            MaxUpdatesPerHour = maxUpdatesPerHour;
            MaxUserConnections = maxUserConnections;
        }
    }
}

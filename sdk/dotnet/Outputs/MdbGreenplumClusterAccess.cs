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
    public sealed class MdbGreenplumClusterAccess
    {
        /// <summary>
        /// Allow access for [Yandex DataLens](https://yandex.cloud/services/datalens).
        /// </summary>
        public readonly bool? DataLens;
        /// <summary>
        /// Allow access for [DataTransfer](https://yandex.cloud/services/data-transfer)
        /// </summary>
        public readonly bool? DataTransfer;
        /// <summary>
        /// Allows access for [SQL queries in the management console](https://yandex.cloud/docs/managed-mysql/operations/web-sql-query).
        /// </summary>
        public readonly bool? WebSql;
        /// <summary>
        /// Allow access for [Yandex Query](https://yandex.cloud/services/query)
        /// </summary>
        public readonly bool? YandexQuery;

        [OutputConstructor]
        private MdbGreenplumClusterAccess(
            bool? dataLens,

            bool? dataTransfer,

            bool? webSql,

            bool? yandexQuery)
        {
            DataLens = dataLens;
            DataTransfer = dataTransfer;
            WebSql = webSql;
            YandexQuery = yandexQuery;
        }
    }
}

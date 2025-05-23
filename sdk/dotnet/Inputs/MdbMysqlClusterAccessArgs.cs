// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMysqlClusterAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow access for [Yandex DataLens](https://yandex.cloud/services/datalens).
        /// </summary>
        [Input("dataLens")]
        public Input<bool>? DataLens { get; set; }

        /// <summary>
        /// Allow access for [DataTransfer](https://yandex.cloud/services/data-transfer).
        /// </summary>
        [Input("dataTransfer")]
        public Input<bool>? DataTransfer { get; set; }

        /// <summary>
        /// Allows access for [SQL queries in the management console](https://yandex.cloud/docs/managed-mysql/operations/web-sql-query).
        /// </summary>
        [Input("webSql")]
        public Input<bool>? WebSql { get; set; }

        public MdbMysqlClusterAccessArgs()
        {
        }
        public static new MdbMysqlClusterAccessArgs Empty => new MdbMysqlClusterAccessArgs();
    }
}

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
    public sealed class MonitoringDashboardWidgetPosition
    {
        /// <summary>
        /// Height.
        /// </summary>
        public readonly int? H;
        /// <summary>
        /// Weight.
        /// </summary>
        public readonly int? W;
        /// <summary>
        /// X-axis top-left corner coordinate.
        /// </summary>
        public readonly int? X;
        /// <summary>
        /// Y-axis top-left corner coordinate.
        /// </summary>
        public readonly int? Y;

        [OutputConstructor]
        private MonitoringDashboardWidgetPosition(
            int? h,

            int? w,

            int? x,

            int? y)
        {
            H = h;
            W = w;
            X = x;
            Y = y;
        }
    }
}

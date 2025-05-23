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
    public sealed class MonitoringDashboardWidgetChartVisualizationSettingHeatmapSetting
    {
        /// <summary>
        /// Heatmap green value
        /// </summary>
        public readonly string? GreenValue;
        /// <summary>
        /// Heatmap red value
        /// </summary>
        public readonly string? RedValue;
        /// <summary>
        /// Heatmap violet_value
        /// </summary>
        public readonly string? VioletValue;
        /// <summary>
        /// Heatmap yellow value
        /// </summary>
        public readonly string? YellowValue;

        [OutputConstructor]
        private MonitoringDashboardWidgetChartVisualizationSettingHeatmapSetting(
            string? greenValue,

            string? redValue,

            string? violetValue,

            string? yellowValue)
        {
            GreenValue = greenValue;
            RedValue = redValue;
            VioletValue = violetValue;
            YellowValue = yellowValue;
        }
    }
}

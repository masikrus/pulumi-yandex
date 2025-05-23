// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingLeftGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max value in extended number format or empty
        /// </summary>
        [Input("max")]
        public Input<string>? Max { get; set; }

        /// <summary>
        /// Min value in extended number format or empty
        /// </summary>
        [Input("min")]
        public Input<string>? Min { get; set; }

        /// <summary>
        /// Tick value precision (null as default, 0-7 in other cases)
        /// </summary>
        [Input("precision")]
        public Input<int>? Precision { get; set; }

        /// <summary>
        /// Title or empty
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Unit format
        /// </summary>
        [Input("unitFormat")]
        public Input<string>? UnitFormat { get; set; }

        public MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingLeftGetArgs()
        {
        }
        public static new MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingLeftGetArgs Empty => new MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingLeftGetArgs();
    }
}

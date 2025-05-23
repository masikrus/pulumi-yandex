// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MonitoringDashboardWidgetChartVisualizationSettingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Aggregation
        /// </summary>
        [Input("aggregation")]
        public Input<string>? Aggregation { get; set; }

        [Input("colorSchemeSettings")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGetArgs>? _colorSchemeSettings;

        /// <summary>
        /// Color scheme settings
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGetArgs> ColorSchemeSettings
        {
            get => _colorSchemeSettings ?? (_colorSchemeSettings = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGetArgs>());
            set => _colorSchemeSettings = value;
        }

        [Input("heatmapSettings")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingHeatmapSettingGetArgs>? _heatmapSettings;

        /// <summary>
        /// Heatmap settings
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingHeatmapSettingGetArgs> HeatmapSettings
        {
            get => _heatmapSettings ?? (_heatmapSettings = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingHeatmapSettingGetArgs>());
            set => _heatmapSettings = value;
        }

        /// <summary>
        /// Interpolate
        /// </summary>
        [Input("interpolate")]
        public Input<string>? Interpolate { get; set; }

        /// <summary>
        /// Normalize
        /// </summary>
        [Input("normalize")]
        public Input<bool>? Normalize { get; set; }

        /// <summary>
        /// Show chart labels
        /// </summary>
        [Input("showLabels")]
        public Input<bool>? ShowLabels { get; set; }

        /// <summary>
        /// Inside chart title
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Visualization type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("yaxisSettings")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingGetArgs>? _yaxisSettings;

        /// <summary>
        /// Y axis settings
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingGetArgs> YaxisSettings
        {
            get => _yaxisSettings ?? (_yaxisSettings = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingYaxisSettingGetArgs>());
            set => _yaxisSettings = value;
        }

        public MonitoringDashboardWidgetChartVisualizationSettingGetArgs()
        {
        }
        public static new MonitoringDashboardWidgetChartVisualizationSettingGetArgs Empty => new MonitoringDashboardWidgetChartVisualizationSettingGetArgs();
    }
}

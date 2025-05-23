// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingArgs : global::Pulumi.ResourceArgs
    {
        [Input("automatics")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingAutomaticArgs>? _automatics;

        /// <summary>
        /// Automatic color scheme
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingAutomaticArgs> Automatics
        {
            get => _automatics ?? (_automatics = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingAutomaticArgs>());
            set => _automatics = value;
        }

        [Input("gradients")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGradientArgs>? _gradients;

        /// <summary>
        /// Gradient color scheme
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGradientArgs> Gradients
        {
            get => _gradients ?? (_gradients = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingGradientArgs>());
            set => _gradients = value;
        }

        [Input("standards")]
        private InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingStandardArgs>? _standards;

        /// <summary>
        /// Standard color scheme
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingStandardArgs> Standards
        {
            get => _standards ?? (_standards = new InputList<Inputs.MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingStandardArgs>());
            set => _standards = value;
        }

        public MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingArgs()
        {
        }
        public static new MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingArgs Empty => new MonitoringDashboardWidgetChartVisualizationSettingColorSchemeSettingArgs();
    }
}

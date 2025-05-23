// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MonitoringDashboardWidgetChartQueryArgs : global::Pulumi.ResourceArgs
    {
        [Input("downsamplings")]
        private InputList<Inputs.MonitoringDashboardWidgetChartQueryDownsamplingArgs>? _downsamplings;

        /// <summary>
        /// Downsampling settings
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartQueryDownsamplingArgs> Downsamplings
        {
            get => _downsamplings ?? (_downsamplings = new InputList<Inputs.MonitoringDashboardWidgetChartQueryDownsamplingArgs>());
            set => _downsamplings = value;
        }

        [Input("targets")]
        private InputList<Inputs.MonitoringDashboardWidgetChartQueryTargetArgs>? _targets;

        /// <summary>
        /// Downsampling settings
        /// </summary>
        public InputList<Inputs.MonitoringDashboardWidgetChartQueryTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MonitoringDashboardWidgetChartQueryTargetArgs>());
            set => _targets = value;
        }

        public MonitoringDashboardWidgetChartQueryArgs()
        {
        }
        public static new MonitoringDashboardWidgetChartQueryArgs Empty => new MonitoringDashboardWidgetChartQueryArgs();
    }
}

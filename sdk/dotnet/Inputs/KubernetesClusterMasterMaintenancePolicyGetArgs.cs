// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesClusterMasterMaintenancePolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.
        /// </summary>
        [Input("autoUpgrade", required: true)]
        public Input<bool> AutoUpgrade { get; set; } = null!;

        [Input("maintenanceWindows")]
        private InputList<Inputs.KubernetesClusterMasterMaintenancePolicyMaintenanceWindowGetArgs>? _maintenanceWindows;

        /// <summary>
        /// This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time. To specify time of day interval, for all days, one element should be provided, with two fields set, `start_time` and `duration`. Please see `zonal_cluster_resource_name` config example.
        /// 
        /// To allow maintenance only on specific days of week, please provide list of elements, with all fields set. Only one time interval (`duration`) is allowed for each day of week. Please see `regional_cluster_resource_name` config example
        /// </summary>
        public InputList<Inputs.KubernetesClusterMasterMaintenancePolicyMaintenanceWindowGetArgs> MaintenanceWindows
        {
            get => _maintenanceWindows ?? (_maintenanceWindows = new InputList<Inputs.KubernetesClusterMasterMaintenancePolicyMaintenanceWindowGetArgs>());
            set => _maintenanceWindows = value;
        }

        public KubernetesClusterMasterMaintenancePolicyGetArgs()
        {
        }
        public static new KubernetesClusterMasterMaintenancePolicyGetArgs Empty => new KubernetesClusterMasterMaintenancePolicyGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class BackupPolicySchedulingArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupSets")]
        private InputList<Inputs.BackupPolicySchedulingBackupSetArgs>? _backupSets;

        /// <summary>
        /// A list of schedules with backup sets that compose the whole scheme.
        /// </summary>
        public InputList<Inputs.BackupPolicySchedulingBackupSetArgs> BackupSets
        {
            get => _backupSets ?? (_backupSets = new InputList<Inputs.BackupPolicySchedulingBackupSetArgs>());
            set => _backupSets = value;
        }

        /// <summary>
        /// Enables or disables scheduling. Default `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Perform backup by interval, since last backup of the host. Maximum value is: 9999 days. See `interval_type` for available values. Exactly on of options should be set: `execute_by_interval` or `execute_by_time`.
        /// </summary>
        [Input("executeByInterval")]
        public Input<int>? ExecuteByInterval { get; set; }

        [Input("executeByTimes")]
        private InputList<Inputs.BackupPolicySchedulingExecuteByTimeArgs>? _executeByTimes;

        /// <summary>
        /// Perform backup periodically at specific time. Exactly on of options should be set: `execute_by_interval` or `execute_by_time`.
        /// </summary>
        [Obsolete(@"The 'execute_by_time' field has been deprecated. Please use 'backup_sets' instead.")]
        public InputList<Inputs.BackupPolicySchedulingExecuteByTimeArgs> ExecuteByTimes
        {
            get => _executeByTimes ?? (_executeByTimes = new InputList<Inputs.BackupPolicySchedulingExecuteByTimeArgs>());
            set => _executeByTimes = value;
        }

        /// <summary>
        /// Maximum number of backup processes allowed to run in parallel. 0 for unlimited. Default `0`.
        /// </summary>
        [Input("maxParallelBackups")]
        public Input<int>? MaxParallelBackups { get; set; }

        /// <summary>
        /// Configuration of the random delay between the execution of parallel tasks. See `interval_type` for available values. Default `30m`.
        /// </summary>
        [Input("randomMaxDelay")]
        public Input<string>? RandomMaxDelay { get; set; }

        /// <summary>
        /// Scheme of the backups. Available values are: `ALWAYS_INCREMENTAL`, `ALWAYS_FULL`, `WEEKLY_FULL_DAILY_INCREMENTAL`, `WEEKLY_INCREMENTAL`. Default `ALWAYS_INCREMENTAL`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// A day of week to start weekly backups. See `day_type` for available values. Default `MONDAY`.
        /// </summary>
        [Input("weeklyBackupDay")]
        public Input<string>? WeeklyBackupDay { get; set; }

        public BackupPolicySchedulingArgs()
        {
        }
        public static new BackupPolicySchedulingArgs Empty => new BackupPolicySchedulingArgs();
    }
}

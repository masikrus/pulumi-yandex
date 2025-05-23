// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class BackupPolicySchedulingBackupSetExecuteByTimeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, schedule will be applied on the last day of month. See `day_type` for available values. Default `true`.
        /// </summary>
        [Input("includeLastDayOfMonth")]
        public Input<bool>? IncludeLastDayOfMonth { get; set; }

        [Input("monthdays")]
        private InputList<int>? _monthdays;

        /// <summary>
        /// List of days when schedule applies. Used in `MONTHLY` type.
        /// </summary>
        public InputList<int> Monthdays
        {
            get => _monthdays ?? (_monthdays = new InputList<int>());
            set => _monthdays = value;
        }

        [Input("months")]
        private InputList<int>? _months;

        /// <summary>
        /// Set of values. Allowed values form 1 to 12.
        /// </summary>
        public InputList<int> Months
        {
            get => _months ?? (_months = new InputList<int>());
            set => _months = value;
        }

        [Input("repeatAts")]
        private InputList<string>? _repeatAts;

        /// <summary>
        /// List of time in format `HH:MM` (24-hours format), when the schedule applies.
        /// </summary>
        public InputList<string> RepeatAts
        {
            get => _repeatAts ?? (_repeatAts = new InputList<string>());
            set => _repeatAts = value;
        }

        /// <summary>
        /// Frequency of backup repetition. See `interval_type` for available values.
        /// </summary>
        [Input("repeatEvery")]
        public Input<string>? RepeatEvery { get; set; }

        /// <summary>
        /// Type of the scheduling. Available values are: `HOURLY`, `DAILY`, `WEEKLY`, `MONTHLY`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("weekdays")]
        private InputList<string>? _weekdays;

        /// <summary>
        /// List of weekdays when the backup will be applied. Used in `WEEKLY` type.
        /// </summary>
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public BackupPolicySchedulingBackupSetExecuteByTimeGetArgs()
        {
        }
        public static new BackupPolicySchedulingBackupSetExecuteByTimeGetArgs Empty => new BackupPolicySchedulingBackupSetExecuteByTimeGetArgs();
    }
}

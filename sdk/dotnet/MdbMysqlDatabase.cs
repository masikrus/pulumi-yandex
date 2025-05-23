// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/mdbMysqlDatabase:MdbMysqlDatabase")]
    public partial class MdbMysqlDatabase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The MySQL cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MdbMysqlDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbMysqlDatabase(string name, MdbMysqlDatabaseArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbMysqlDatabase:MdbMysqlDatabase", name, args ?? new MdbMysqlDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbMysqlDatabase(string name, Input<string> id, MdbMysqlDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbMysqlDatabase:MdbMysqlDatabase", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbMysqlDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbMysqlDatabase Get(string name, Input<string> id, MdbMysqlDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbMysqlDatabase(name, id, state, options);
        }
    }

    public sealed class MdbMysqlDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The MySQL cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MdbMysqlDatabaseArgs()
        {
        }
        public static new MdbMysqlDatabaseArgs Empty => new MdbMysqlDatabaseArgs();
    }

    public sealed class MdbMysqlDatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The MySQL cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MdbMysqlDatabaseState()
        {
        }
        public static new MdbMysqlDatabaseState Empty => new MdbMysqlDatabaseState();
    }
}

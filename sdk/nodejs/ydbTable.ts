// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class YdbTable extends pulumi.CustomResource {
    /**
     * Get an existing YdbTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: YdbTableState, opts?: pulumi.CustomResourceOptions): YdbTable {
        return new YdbTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/ydbTable:YdbTable';

    /**
     * Returns true if the given object is an instance of YdbTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is YdbTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === YdbTable.__pulumiType;
    }

    public readonly attributes!: pulumi.Output<{[key: string]: string}>;
    public readonly columns!: pulumi.Output<outputs.YdbTableColumn[]>;
    public readonly connectionString!: pulumi.Output<string>;
    public readonly families!: pulumi.Output<outputs.YdbTableFamily[] | undefined>;
    public readonly keyBloomFilter!: pulumi.Output<boolean>;
    public readonly partitioningSettings!: pulumi.Output<outputs.YdbTablePartitioningSettings>;
    public readonly path!: pulumi.Output<string>;
    public readonly primaryKeys!: pulumi.Output<string[]>;
    public readonly readReplicasSettings!: pulumi.Output<string>;
    public readonly ttl!: pulumi.Output<outputs.YdbTableTtl | undefined>;

    /**
     * Create a YdbTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: YdbTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: YdbTableArgs | YdbTableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as YdbTableState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["columns"] = state ? state.columns : undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["families"] = state ? state.families : undefined;
            resourceInputs["keyBloomFilter"] = state ? state.keyBloomFilter : undefined;
            resourceInputs["partitioningSettings"] = state ? state.partitioningSettings : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["primaryKeys"] = state ? state.primaryKeys : undefined;
            resourceInputs["readReplicasSettings"] = state ? state.readReplicasSettings : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as YdbTableArgs | undefined;
            if ((!args || args.columns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'columns'");
            }
            if ((!args || args.connectionString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionString'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.primaryKeys === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryKeys'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["columns"] = args ? args.columns : undefined;
            resourceInputs["connectionString"] = args ? args.connectionString : undefined;
            resourceInputs["families"] = args ? args.families : undefined;
            resourceInputs["keyBloomFilter"] = args ? args.keyBloomFilter : undefined;
            resourceInputs["partitioningSettings"] = args ? args.partitioningSettings : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["primaryKeys"] = args ? args.primaryKeys : undefined;
            resourceInputs["readReplicasSettings"] = args ? args.readReplicasSettings : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(YdbTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering YdbTable resources.
 */
export interface YdbTableState {
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    columns?: pulumi.Input<pulumi.Input<inputs.YdbTableColumn>[]>;
    connectionString?: pulumi.Input<string>;
    families?: pulumi.Input<pulumi.Input<inputs.YdbTableFamily>[]>;
    keyBloomFilter?: pulumi.Input<boolean>;
    partitioningSettings?: pulumi.Input<inputs.YdbTablePartitioningSettings>;
    path?: pulumi.Input<string>;
    primaryKeys?: pulumi.Input<pulumi.Input<string>[]>;
    readReplicasSettings?: pulumi.Input<string>;
    ttl?: pulumi.Input<inputs.YdbTableTtl>;
}

/**
 * The set of arguments for constructing a YdbTable resource.
 */
export interface YdbTableArgs {
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    columns: pulumi.Input<pulumi.Input<inputs.YdbTableColumn>[]>;
    connectionString: pulumi.Input<string>;
    families?: pulumi.Input<pulumi.Input<inputs.YdbTableFamily>[]>;
    keyBloomFilter?: pulumi.Input<boolean>;
    partitioningSettings?: pulumi.Input<inputs.YdbTablePartitioningSettings>;
    path: pulumi.Input<string>;
    primaryKeys: pulumi.Input<pulumi.Input<string>[]>;
    readReplicasSettings?: pulumi.Input<string>;
    ttl?: pulumi.Input<inputs.YdbTableTtl>;
}

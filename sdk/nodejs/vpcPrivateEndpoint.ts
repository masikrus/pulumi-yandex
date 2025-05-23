// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class VpcPrivateEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcPrivateEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPrivateEndpointState, opts?: pulumi.CustomResourceOptions): VpcPrivateEndpoint {
        return new VpcPrivateEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint';

    /**
     * Returns true if the given object is an instance of VpcPrivateEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPrivateEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPrivateEndpoint.__pulumiType;
    }

    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Private endpoint DNS options block.
     */
    public readonly dnsOptions!: pulumi.Output<outputs.VpcPrivateEndpointDnsOptions>;
    /**
     * Private endpoint address specification block. > Only one of `addressId` or `subnetId` + `address` arguments can be
     * specified.
     */
    public readonly endpointAddress!: pulumi.Output<outputs.VpcPrivateEndpointEndpointAddress>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the network which private endpoint belongs to.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Private endpoint for Object Storage.
     */
    public readonly objectStorage!: pulumi.Output<outputs.VpcPrivateEndpointObjectStorage>;
    /**
     * Status of the private endpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a VpcPrivateEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcPrivateEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPrivateEndpointArgs | VpcPrivateEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcPrivateEndpointState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsOptions"] = state ? state.dnsOptions : undefined;
            resourceInputs["endpointAddress"] = state ? state.endpointAddress : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["objectStorage"] = state ? state.objectStorage : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as VpcPrivateEndpointArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.objectStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectStorage'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsOptions"] = args ? args.dnsOptions : undefined;
            resourceInputs["endpointAddress"] = args ? args.endpointAddress : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["objectStorage"] = args ? args.objectStorage : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcPrivateEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPrivateEndpoint resources.
 */
export interface VpcPrivateEndpointState {
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Private endpoint DNS options block.
     */
    dnsOptions?: pulumi.Input<inputs.VpcPrivateEndpointDnsOptions>;
    /**
     * Private endpoint address specification block. > Only one of `addressId` or `subnetId` + `address` arguments can be
     * specified.
     */
    endpointAddress?: pulumi.Input<inputs.VpcPrivateEndpointEndpointAddress>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network which private endpoint belongs to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Private endpoint for Object Storage.
     */
    objectStorage?: pulumi.Input<inputs.VpcPrivateEndpointObjectStorage>;
    /**
     * Status of the private endpoint.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPrivateEndpoint resource.
 */
export interface VpcPrivateEndpointArgs {
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Private endpoint DNS options block.
     */
    dnsOptions?: pulumi.Input<inputs.VpcPrivateEndpointDnsOptions>;
    /**
     * Private endpoint address specification block. > Only one of `addressId` or `subnetId` + `address` arguments can be
     * specified.
     */
    endpointAddress?: pulumi.Input<inputs.VpcPrivateEndpointEndpointAddress>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the network which private endpoint belongs to.
     */
    networkId: pulumi.Input<string>;
    /**
     * Private endpoint for Object Storage.
     */
    objectStorage: pulumi.Input<inputs.VpcPrivateEndpointObjectStorage>;
}

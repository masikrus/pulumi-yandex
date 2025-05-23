// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the yandex package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'yandex';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The ID of the [Cloud](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#cloud) to apply any
     * resources to. This can also be specified using environment variable `YC_CLOUD_ID`.
     */
    public readonly cloudId!: pulumi.Output<string | undefined>;
    /**
     * The endpoint for API calls, default value is **api.cloud.yandex.net:443**. This can also be defined by environment
     * variable `YC_ENDPOINT`.
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * The ID of the [Folder](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#folder) to operate under,
     * if not specified by a given resource. This can also be specified using environment variable `YC_FOLDER_ID`.
     */
    public readonly folderId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the [Cloud Organization](https://yandex.cloud/docs/organization/quickstart) to operate under.
     */
    public readonly organizationId!: pulumi.Output<string | undefined>;
    /**
     * Profile name to use in the shared credentials file. Default value is `default`.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * [The region](https://yandex.cloud/docs/overview/concepts/region) where operations will take place. For example
     * `ru-central1`.
     */
    public readonly regionId!: pulumi.Output<string | undefined>;
    /**
     * Contains either a path to or the contents of the [Service Account
     * file](https://yandex.cloud/docs/iam/concepts/authorization/key) in JSON format. This can also be specified using
     * environment variable `YC_SERVICE_ACCOUNT_KEY_FILE`. You can read how to create service account key file
     * [here](https://yandex.cloud/docs/iam/operations/iam-token/create-for-sa#keys-create). > Only one of `token` or
     * `serviceAccountKeyFile` must be specified. > One can authenticate via instance service account from inside a compute
     * instance. In order to use this method, omit both `token`/`serviceAccountKeyFile` and attach service account to the
     * instance. [Working with Yandex Cloud from inside an
     * instance](https://yandex.cloud/docs/compute/operations/vm-connect/auth-inside-vm).
     */
    public readonly serviceAccountKeyFile!: pulumi.Output<string | undefined>;
    /**
     * Shared credentials file path. Supported keys: `storageAccessKey` and `storageSecretKey`. > The `storageAccessKey` and
     * `storageSecretKey` attributes from the shared credentials file are used only when the provider and a storage
     * data/resource do not have an access/secret keys explicitly specified.
     */
    public readonly sharedCredentialsFile!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud Object Storage access key, which is used when a storage data/resource doesn't have an access key explicitly
     * specified. This can also be specified using environment variable `YC_STORAGE_ACCESS_KEY`.
     */
    public readonly storageAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud [Object Storage Endpoint](https://yandex.cloud/docs/storage/s3/#request-url), which is used to connect to
     * `S3 API`. Default value is **storage.yandexcloud.net**.
     */
    public readonly storageEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud Object Storage secret key, which is used when a storage data/resource doesn't have a secret key explicitly
     * specified. This can also be specified using environment variable `YC_STORAGE_SECRET_KEY`.
     */
    public readonly storageSecretKey!: pulumi.Output<string | undefined>;
    /**
     * Security token or IAM token used for authentication in Yandex Cloud. Check
     * [documentation](https://yandex.cloud/docs/iam/operations/iam-token/create) about how to create IAM token. This can also
     * be specified using environment variable `YC_TOKEN`.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud Message Queue service access key, which is used when a YMQ queue resource doesn't have an access key
     * explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_ACCESS_KEY`.
     */
    public readonly ymqAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud Message Queue service endpoint. Default value is **message-queue.api.cloud.yandex.net**.
     */
    public readonly ymqEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Yandex Cloud Message Queue service secret key, which is used when a YMQ queue resource doesn't have a secret key
     * explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_SECRET_KEY`.
     */
    public readonly ymqSecretKey!: pulumi.Output<string | undefined>;
    /**
     * The default [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) to operate under, if not
     * specified by a given resource. This can also be specified using environment variable `YC_ZONE`.
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["cloudId"] = (args ? args.cloudId : undefined) ?? utilities.getEnv("YANDEX_CLOUD_CLOUD_ID");
            resourceInputs["endpoint"] = (args ? args.endpoint : undefined) ?? utilities.getEnv("YANDEX_CLOUD_ENDPOINT");
            resourceInputs["folderId"] = (args ? args.folderId : undefined) ?? utilities.getEnv("YANDEX_CLOUD_FOLDER_ID");
            resourceInputs["insecure"] = pulumi.output((args ? args.insecure : undefined) ?? utilities.getEnvBoolean("YANDEX_CLOUD_INSECURE")).apply(JSON.stringify);
            resourceInputs["maxRetries"] = pulumi.output((args ? args.maxRetries : undefined) ?? utilities.getEnvNumber("YANDEX_CLOUD_MAX_RETRIES")).apply(JSON.stringify);
            resourceInputs["organizationId"] = (args ? args.organizationId : undefined) ?? utilities.getEnv("YANDEX_CLOUD_ORGANIZATION_ID");
            resourceInputs["plaintext"] = pulumi.output((args ? args.plaintext : undefined) ?? utilities.getEnvBoolean("YANDEX_CLOUD_PLAINTEXT")).apply(JSON.stringify);
            resourceInputs["profile"] = (args ? args.profile : undefined) ?? utilities.getEnv("YANDEX_CLOUD_PROFILE");
            resourceInputs["regionId"] = (args ? args.regionId : undefined) ?? utilities.getEnv("YANDEX_CLOUD_REGION_ID");
            resourceInputs["serviceAccountKeyFile"] = (args ? args.serviceAccountKeyFile : undefined) ?? utilities.getEnv("YANDEX_CLOUD_SERVICE_ACCOUNT_KEY_FILE");
            resourceInputs["sharedCredentialsFile"] = (args ? args.sharedCredentialsFile : undefined) ?? utilities.getEnv("YANDEX_CLOUD_SHARED_CREDENTIALS_FILE");
            resourceInputs["storageAccessKey"] = (args ? args.storageAccessKey : undefined) ?? utilities.getEnv("YANDEX_CLOUD_STORAGE_ACCESS_KEY");
            resourceInputs["storageEndpoint"] = (args ? args.storageEndpoint : undefined) ?? utilities.getEnv("YANDEX_CLOUD_STORAGE_ENDPOINT");
            resourceInputs["storageSecretKey"] = (args?.storageSecretKey ? pulumi.secret(args.storageSecretKey) : undefined) ?? utilities.getEnv("YANDEX_CLOUD_STORAGE_SECRET_KEY");
            resourceInputs["token"] = (args?.token ? pulumi.secret(args.token) : undefined) ?? utilities.getEnv("YANDEX_CLOUD_TOKEN");
            resourceInputs["ymqAccessKey"] = (args ? args.ymqAccessKey : undefined) ?? utilities.getEnv("YANDEX_CLOUD_YMQ_ACCESS_KEY");
            resourceInputs["ymqEndpoint"] = (args ? args.ymqEndpoint : undefined) ?? utilities.getEnv("YANDEX_CLOUD_YMQ_ENDPOINT");
            resourceInputs["ymqSecretKey"] = (args?.ymqSecretKey ? pulumi.secret(args.ymqSecretKey) : undefined) ?? utilities.getEnv("YANDEX_CLOUD_YMQ_SECRET_KEY");
            resourceInputs["zone"] = (args ? args.zone : undefined) ?? utilities.getEnv("YANDEX_CLOUD_ZONE");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["storageSecretKey", "token", "ymqSecretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }

    /**
     * This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
     */
    terraformConfig(): pulumi.Output<Provider.TerraformConfigResult> {
        return pulumi.runtime.call("pulumi:providers:yandex/terraformConfig", {
            "__self__": this,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The ID of the [Cloud](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#cloud) to apply any
     * resources to. This can also be specified using environment variable `YC_CLOUD_ID`.
     */
    cloudId?: pulumi.Input<string>;
    /**
     * The endpoint for API calls, default value is **api.cloud.yandex.net:443**. This can also be defined by environment
     * variable `YC_ENDPOINT`.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The ID of the [Folder](https://yandex.cloud/docs/resource-manager/concepts/resources-hierarchy#folder) to operate under,
     * if not specified by a given resource. This can also be specified using environment variable `YC_FOLDER_ID`.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`.
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * This is the maximum number of times an API call is retried, in the case where requests are being throttled or
     * experiencing transient failures. The delay between the subsequent API calls increases exponentially.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * The ID of the [Cloud Organization](https://yandex.cloud/docs/organization/quickstart) to operate under.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Disable use of TLS. Default value is `false`.
     */
    plaintext?: pulumi.Input<boolean>;
    /**
     * Profile name to use in the shared credentials file. Default value is `default`.
     */
    profile?: pulumi.Input<string>;
    /**
     * [The region](https://yandex.cloud/docs/overview/concepts/region) where operations will take place. For example
     * `ru-central1`.
     */
    regionId?: pulumi.Input<string>;
    /**
     * Contains either a path to or the contents of the [Service Account
     * file](https://yandex.cloud/docs/iam/concepts/authorization/key) in JSON format. This can also be specified using
     * environment variable `YC_SERVICE_ACCOUNT_KEY_FILE`. You can read how to create service account key file
     * [here](https://yandex.cloud/docs/iam/operations/iam-token/create-for-sa#keys-create). > Only one of `token` or
     * `serviceAccountKeyFile` must be specified. > One can authenticate via instance service account from inside a compute
     * instance. In order to use this method, omit both `token`/`serviceAccountKeyFile` and attach service account to the
     * instance. [Working with Yandex Cloud from inside an
     * instance](https://yandex.cloud/docs/compute/operations/vm-connect/auth-inside-vm).
     */
    serviceAccountKeyFile?: pulumi.Input<string>;
    /**
     * Shared credentials file path. Supported keys: `storageAccessKey` and `storageSecretKey`. > The `storageAccessKey` and
     * `storageSecretKey` attributes from the shared credentials file are used only when the provider and a storage
     * data/resource do not have an access/secret keys explicitly specified.
     */
    sharedCredentialsFile?: pulumi.Input<string>;
    /**
     * Yandex Cloud Object Storage access key, which is used when a storage data/resource doesn't have an access key explicitly
     * specified. This can also be specified using environment variable `YC_STORAGE_ACCESS_KEY`.
     */
    storageAccessKey?: pulumi.Input<string>;
    /**
     * Yandex Cloud [Object Storage Endpoint](https://yandex.cloud/docs/storage/s3/#request-url), which is used to connect to
     * `S3 API`. Default value is **storage.yandexcloud.net**.
     */
    storageEndpoint?: pulumi.Input<string>;
    /**
     * Yandex Cloud Object Storage secret key, which is used when a storage data/resource doesn't have a secret key explicitly
     * specified. This can also be specified using environment variable `YC_STORAGE_SECRET_KEY`.
     */
    storageSecretKey?: pulumi.Input<string>;
    /**
     * Security token or IAM token used for authentication in Yandex Cloud. Check
     * [documentation](https://yandex.cloud/docs/iam/operations/iam-token/create) about how to create IAM token. This can also
     * be specified using environment variable `YC_TOKEN`.
     */
    token?: pulumi.Input<string>;
    /**
     * Yandex Cloud Message Queue service access key, which is used when a YMQ queue resource doesn't have an access key
     * explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_ACCESS_KEY`.
     */
    ymqAccessKey?: pulumi.Input<string>;
    /**
     * Yandex Cloud Message Queue service endpoint. Default value is **message-queue.api.cloud.yandex.net**.
     */
    ymqEndpoint?: pulumi.Input<string>;
    /**
     * Yandex Cloud Message Queue service secret key, which is used when a YMQ queue resource doesn't have a secret key
     * explicitly specified. This can also be specified using environment variable `YC_MESSAGE_QUEUE_SECRET_KEY`.
     */
    ymqSecretKey?: pulumi.Input<string>;
    /**
     * The default [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) to operate under, if not
     * specified by a given resource. This can also be specified using environment variable `YC_ZONE`.
     */
    zone?: pulumi.Input<string>;
}

export namespace Provider {
    /**
     * The results of the Provider.terraformConfig method.
     */
    export interface TerraformConfigResult {
        readonly result: {[key: string]: any};
    }

}

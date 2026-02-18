# Yandex Cloud Pulumi Provider Configuration

This document describes the available configuration parameters for the Yandex Cloud Pulumi provider.

## Authentication Parameters

- **token** - Yandex Cloud OAuth token or IAM token for authentication  
  Env: `YANDEX_CLOUD_TOKEN`

- **service_account_key_file** - Path to service account key file in JSON format  
  Env: `YANDEX_CLOUD_SERVICE_ACCOUNT_KEY_FILE`

- **shared_credentials_file** - Path to shared credentials file  
  Env: `YANDEX_CLOUD_SHARED_CREDENTIALS_FILE`

- **profile** - Profile name from shared credentials file  
  Env: `YANDEX_CLOUD_PROFILE`

## Resource Location Parameters

- **folder_id** - Default folder ID for resources  
  Env: `YANDEX_CLOUD_FOLDER_ID`

- **cloud_id** - Default cloud ID for resources  
  Env: `YANDEX_CLOUD_CLOUD_ID`

- **organization_id** - Default organization ID  
  Env: `YANDEX_CLOUD_ORGANIZATION_ID`

- **region_id** - Default region (e.g., "ru-central1")  
  Env: `YANDEX_CLOUD_REGION_ID`

- **zone** - Default availability zone (e.g., "ru-central1-a")  
  Env: `YANDEX_CLOUD_ZONE`

## API Configuration

- **endpoint** - Custom API endpoint URL  
  Env: `YANDEX_CLOUD_ENDPOINT`

- **insecure** - Allow insecure connections to API (true/false)  
  Env: `YANDEX_CLOUD_INSECURE`

- **plaintext** - Disable TLS for API connections (true/false)  
  Env: `YANDEX_CLOUD_PLAINTEXT`

- **max_retries** - Maximum number of API retries  
  Env: `YANDEX_CLOUD_MAX_RETRIES`

## Storage Configuration

- **storage_endpoint** - Custom Object Storage endpoint  
  Env: `YANDEX_CLOUD_STORAGE_ENDPOINT`

- **storage_access_key** - Object Storage access key  
  Env: `YANDEX_CLOUD_STORAGE_ACCESS_KEY`

- **storage_secret_key** - Object Storage secret key  
  Env: `YANDEX_CLOUD_STORAGE_SECRET_KEY`

## Message Queue (YMQ) Configuration

- **ymq_endpoint** - Custom Yandex Message Queue endpoint  
  Env: `YANDEX_CLOUD_YMQ_ENDPOINT`

- **ymq_access_key** - YMQ access key  
  Env: `YANDEX_CLOUD_YMQ_ACCESS_KEY`

- **ymq_secret_key** - YMQ secret key  
  Env: `YANDEX_CLOUD_YMQ_SECRET_KEY`

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/yandex
```

or `yarn`:

```bash
yarn add @pulumi/yandex
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_yandex
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-yandex/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Yandex
```

## Configuration

The following configuration points are available for the `yandex` provider:

- `yandex:region` (environment: `YANDEX_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/yandex/api-docs/).

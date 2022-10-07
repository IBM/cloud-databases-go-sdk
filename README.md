[![Build Status](https://travis-ci.com/IBM/cloud-databases-go-sdk.svg?branch=main)](https://travis-ci.com/IBM/cloud-databases-go-sdk)
[![Release](https://img.shields.io/github/v/release/IBM/cloud-databases-go-sdk)](https://github.com/IBM/cloud-databases-go-sdk/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/IBM/cloud-databases-go-sdk.svg)](https://pkg.go.dev/github.com/IBM/cloud-databases-go-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/cloud-databases-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![CLA assistant](https://cla-assistant.io/readme/badge/ibm/cloud-databases-go-sdk)](https://cla-assistant.io/ibm/cloud-databases-go-sdk)

# IBM Cloud Databases Go SDK 0.2.0
Go client library to interact with the various [IBM Cloud Cloud Databases APIs](https://cloud.ibm.com/apidocs?category=cloud-databases).

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Running the Integration Tests](#running-the-integration-tests)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Cloud Databases Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Cloud Databases](https://cloud.ibm.com/apidocs/cloud-databases-api/cloud-databases-api-v5) | clouddatabasesv5

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.14 or above.

## Installation
The current version of this SDK: 0.2.0

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `clouddatabasesv5` part of the import path is the package name
associated with the IBM Cloud Databases service.
See the service table above to find the approprate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5
```
Be sure to use the appropriate package name from the service table above for the services used by your application.


## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Running the Integration Tests
To run the integration tests run the `make test-int` command in the root directory. If you wish to run all the integration tests
you can configure your `cloud_databases.env` file as seen below:

```
CLOUD_DATABASES_URL=<service base url> (e.g. https://api.eu-gb.databases.cloud.ibm.com/v5/ibm)
CLOUD_DATABASES_AUTH_TYPE=iam
CLOUD_DATABASES_APIKEY=<IAM apikey>
CLOUD_DATABASES_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
CLOUD_DATABASES_DEPLOYMENT_ID=<ID of an example deployment>
CLOUD_DATABASES_REPLICA_ID=<ID of an example replica>
```

Make sure your env file is in the root directory.

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](github.com/IBM/cloud-databases-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

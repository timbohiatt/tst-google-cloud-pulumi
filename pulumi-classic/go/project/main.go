// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		organizationId := cfg.Require("organizationId")
		projectId := cfg.Require("projectId")
		name := cfg.Require("name")
		billingAccount := cfg.Require("billingAccount")
		_, err := organizations.NewProject(ctx, "project", &organizations.ProjectArgs{
			OrgId:          pulumi.String(organizationId),
			ProjectId:      pulumi.String(projectId),
			BillingAccount: pulumi.String(billingAccount),
			Name:           pulumi.String(name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

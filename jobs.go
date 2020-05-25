package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
)

var jobs = sdk.Jobs{
	sdk.Job{
		Handler:     CreateUser,
		Title:       "Create DB User",
		Description: "Creates a database user with least privileged permissions.",
	},
	sdk.Job{
		Handler:     MigrateDB,
		Title:       "DB Migration",
		Description: "Imports newest test data dump and migrates to newest version.",
		DependsOn:   []string{"Create DB User"},
	},
	sdk.Job{
		Handler:     CreateNamespace,
		Title:       "Create K8S Namespace",
		Description: "Creates a new Kubernetes namespace for the new test environment.",
		DependsOn:   []string{"DB Migration"},
	},
	sdk.Job{
		Handler:     CreateDeployment,
		Title:       "Create K8S Deployment",
		Description: "Creates a new Kubernetes deployment for the new test environment.",
		DependsOn:   []string{"Create K8S Namespace"},
	},
}

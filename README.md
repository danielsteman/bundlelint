# BundleLint ✨

An opinionated validation on top of [databricks-cli](https://docs.gcp.databricks.com/en/dev-tools/cli/bundle-commands.html)'s `bundle validate`.

When the number of asset bundles in your Databricks instance grows, it becomes harder and harder to keep their configurations consistent and complete. For example, as a data platform engineer, you might want to enforce that all jobs that are deployed in the production workspace have notifications configured. Or you want to enforce a specific range of compute types for specific workloads. These types of validations are highly opinionated so it's understandable that the `databricks-cli bundle` API group doesn't cater this use case. This is why `bundlelint` saw the light.

A CLI tool build with Go, much like the `databricks-cli`, is easily distributed across several platforms and can be ran on demand, as a pre-commit hook or in a CI lint task.

## Example usage

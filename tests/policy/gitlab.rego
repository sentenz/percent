package tests.policy

import rego.v1

is_gitlab_ci if {
	input.include
}

# METADATA
# title: GitLab CI component version pinning
# description: Ensure GitLab CI components are pinned to a specific version
# custom:
#   severity: medium
# entrypoint: true
deny_gitlab_ci contains msg if {
	is_gitlab_ci

	some include in input.include
	component := include.component
	parts := split(component, "@")
	version := parts[count(parts) - 1]

	not is_strict_semver(version)
	msg := $"Component '{component}' uses label tag '{version}'. Pin to an exact version like 'v1.2.3'."
}

is_strict_semver(v) if {
	is_string(v)
	regex.match(`^v?\d+\.\d+\.\d+$`, v)
}

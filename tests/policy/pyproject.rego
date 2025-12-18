package main

import rego.v1

is_pyproject if {
	input.tool.poetry # Specific to Poetry, adjust for Flit/Hatch if needed
}

deny contains msg if {
	is_pyproject

	dependencies := input.tool.poetry.dependencies
	some name, version in dependencies

	not is_strict_semver(version)
	msg := sprintf("Dependency '%s' uses ranged version specifier '%s'. Pin to the exact semantic versioning tag like 'v1.2.3'.", [name, version])
}

deny contains msg if {
	is_pyproject

	dependencies := input.poetry.group.dev.dependencies
	some name, version in dependencies

	not is_strict_semver(version)
	msg := sprintf("Dev Dependency '%s' uses ranged version specifier '%s'. Pin to the exact semantic versioning tag like 'v1.2.3'.", [name, version])
}

is_strict_semver(v) if {
	is_string(v)
	regex.match(`^v?\d+\.\d+\.\d+$`, v)
}

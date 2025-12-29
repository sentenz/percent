# Changelog

## [2.1.8](https://github.com/sentenz/percent/compare/2.1.7...2.1.8) (2025-12-29)


### Bug Fixes

* ensure SBOM generation job depends on semantic-release completion ([6c57639](https://github.com/sentenz/percent/commit/6c57639f9cf4904578f2b7b03eba4d3a3c1182c2))

## [2.1.7](https://github.com/sentenz/percent/compare/2.1.6...2.1.7) (2025-12-29)


### Bug Fixes

* remove conditional trigger for SBOM generation on release events ([755d5a4](https://github.com/sentenz/percent/commit/755d5a4bcaac44a609fb3f350d8c0e78173f4d70))

## [2.1.6](https://github.com/sentenz/percent/compare/2.1.5...2.1.6) (2025-12-29)


### Bug Fixes

* update semantic-release workflow to remove release trigger and adjust SBOM publishing ([b1c0ddc](https://github.com/sentenz/percent/commit/b1c0ddc94046465354deb1835d2a3a7738bb6086))

## [2.1.5](https://github.com/sentenz/percent/compare/2.1.4...2.1.5) (2025-12-29)


### Bug Fixes

* publish SBOM to release notes ([170527c](https://github.com/sentenz/percent/commit/170527c339c7d7f68eff151e76fb47772d83500a))

## [2.1.4](https://github.com/sentenz/percent/compare/2.1.3...2.1.4) (2025-12-29)


### Bug Fixes

* add release trigger types for semantic-release workflow ([e141df3](https://github.com/sentenz/percent/commit/e141df3d99d2bb1c9a646ff398d3109b4c0ab2ff))

## [2.1.3](https://github.com/sentenz/percent/compare/2.1.2...2.1.3) (2025-12-29)


### Bug Fixes

* update release trigger type from 'created' to 'published' in Trivy workflow ([cfe9bc6](https://github.com/sentenz/percent/commit/cfe9bc66da70e502f27551a5d8e1ee3b7aeaede1))

## [2.1.2](https://github.com/sentenz/percent/compare/2.1.1...2.1.2) (2025-12-29)


### Bug Fixes

* add Trivy action for SBOM generation and publishing to releases ([dbf7968](https://github.com/sentenz/percent/commit/dbf796851627daaa649cd414e14885f27e757f84))

## [2.1.1](https://github.com/sentenz/percent/compare/2.1.0...2.1.1) (2025-12-29)


### Bug Fixes

* add benchmark comparison task to Makefile and update AGENTS.md ([42707eb](https://github.com/sentenz/percent/commit/42707ebc643b03494d2e3dcacdde20d6282482dc))

# [2.1.0](https://github.com/sentenz/percent/compare/2.0.2...2.1.0) (2025-12-20)


### Bug Fixes

* resolve plugin configuration in .releaserc.json ([cffd6f3](https://github.com/sentenz/percent/commit/cffd6f351d792cde7856972fce93f528a151643e))


### Features

* add AGENTS.md for automated unit testing guidelines ([812e990](https://github.com/sentenz/percent/commit/812e990ec324ae3106e73c632824c3142b785891))

## <small>2.0.2 (2025-12-20)</small>

* fix: remove hidden types from release notes configuration ([c15a938](https://github.com/sentenz/percent/commit/c15a938))

## <small>2.0.1 (2025-12-20)</small>

* fix: resolve semantic-release configuration to adhere conventional commits ([c377588](https://github.com/sentenz/percent/commit/c377588))


## 2.0.0 (2025-12-20)

* chore!: change MIT license with Apache 2.0 using SPDX identifiers (#12) ([bca28c8](https://github.com/sentenz/percent/commit/bca28c8)), closes [#12](https://github.com/sentenz/percent/issues/12) [#11](https://github.com/sentenz/percent/issues/11)


### BREAKING CHANGE

* The project's license has changed from MIT to Apache 2.0. Users must ensure compliance with Apache 2.0 requirements.

# 1.0.0 (2023-07-04)


### Features

* initial commit ([95620cb](https://github.com/sentenz/percent/commit/95620cb22e912b18b4af32dbb23be1847a2d9afe))

# Changelog

## [3.0.10](https://github.com/sentenz/percent/compare/v3.0.9...v3.0.10) (2026-08-23)

### Bug Fixes

* **deps:** update go dependencies ([#147](https://github.com/sentenz/percent/issues/147)) ([18c4873](https://github.com/sentenz/percent/commit/18c48732150fd0e856bf409e8b9eb52548796302))

## [3.0.9](https://github.com/sentenz/percent/compare/v3.0.8...v3.0.9) (2026-08-16)

### Bug Fixes

* **deps:** update go dependencies ([#143](https://github.com/sentenz/percent/issues/143)) ([8cbb685](https://github.com/sentenz/percent/commit/8cbb6856cc9301fa4a92155a563e0f1d646082ca))

## [3.0.8](https://github.com/sentenz/percent/compare/v3.0.7...v3.0.8) (2026-08-16)

### Bug Fixes

* **renovate:** tidy Go modules before vendoring ([8c75499](https://github.com/sentenz/percent/commit/8c75499730ba6eb07d9a07aeb3621cd9affed59a))

## [3.0.7](https://github.com/sentenz/percent/compare/v3.0.6...v3.0.7) (2026-08-12)

### Bug Fixes

* **ci:** repair semantic-release workflow ([724b4c2](https://github.com/sentenz/percent/commit/724b4c21057b2e5334def6a4200407bc9777b1b6))
* generating empty Changelog entries resolved in semantic-release.yml ([a996296](https://github.com/sentenz/percent/commit/a996296fa4e012f003b43491742bdee12a033e9a))

## [3.0.6](https://github.com/sentenz/percent/compare/v3.0.5...v3.0.6) (2026-08-10)

### Bug Fixes

* resolve semantic-release notes generation ([#138](https://github.com/sentenz/percent/issues/138)) ([6e8d529](https://github.com/sentenz/percent/commit/6e8d5293e9f8792bb46c5889814f7fb85ff84b7b))

## [3.0.5](https://github.com/sentenz/percent/compare/3.0.4...3.0.5) (2026-08-10)

## [3.0.4](https://github.com/sentenz/percent/compare/v3.0.3...3.0.4) (2026-08-09)

## [3.0.3](https://github.com/sentenz/percent/compare/v3.0.2...v3.0.3) (2026-02-04)


### Bug Fixes

* update module path and imports to v3 for Go modules compatibility ([b5e9531](https://github.com/sentenz/percent/commit/b5e953137f828bf4f295f308300db62c4fb2dac4))

## [3.0.2](https://github.com/sentenz/percent/compare/v3.0.1...v3.0.2) (2026-02-04)


### Bug Fixes

* remove version suffix from module path and imports ([042b65b](https://github.com/sentenz/percent/commit/042b65b512c848bb1114f0bc6cf0f311ed2ebb2c))

## [3.0.1](https://github.com/sentenz/percent/compare/v3.0.0...v3.0.1) (2026-02-03)


### Bug Fixes

* prepend 'v' to tag name for SBOM artifact publishing ([f04d646](https://github.com/sentenz/percent/commit/f04d6466adc8548207e761e90cd898426243089a))

# [3.0.0](https://github.com/sentenz/percent/compare/v2.0.0...v3.0.0) (2026-02-03)


* fix!: update module path to v2 for Go modules compatibility ([113c07a](https://github.com/sentenz/percent/commit/113c07a7a946b11f2ca59223b0c9dad4e86bcc80))


### BREAKING CHANGES

* Module path changed from github.com/sentenz/percent to github.com/sentenz/percent/v2

This change follows Go modules semantic import versioning requirements for v2+ modules.

# [2.0.0](https://github.com/sentenz/percent/compare/v1.0.0...v2.0.0) (2026-02-03)


* chore!: change MIT license with Apache 2.0 using SPDX identifiers ([#12](https://github.com/sentenz/percent/issues/12)) ([bca28c8](https://github.com/sentenz/percent/commit/bca28c8f18a05ba8e724cab73f8041c38f8905d5)), closes [#11](https://github.com/sentenz/percent/issues/11)


### Bug Fixes

* add benchmark comparison task to Makefile and update AGENTS.md ([42707eb](https://github.com/sentenz/percent/commit/42707ebc643b03494d2e3dcacdde20d6282482dc))
* add release trigger types for semantic-release workflow ([e141df3](https://github.com/sentenz/percent/commit/e141df3d99d2bb1c9a646ff398d3109b4c0ab2ff))
* add tag-format parameter to semantic-release GitHub Action ([f2917d0](https://github.com/sentenz/percent/commit/f2917d07d0dcc2cfdcdb50f7924297c601d682ef))
* add Trivy action for SBOM generation and publishing to releases ([dbf7968](https://github.com/sentenz/percent/commit/dbf796851627daaa649cd414e14885f27e757f84))
* correct tag_name reference for SBOM artifact publishing ([49ef8ed](https://github.com/sentenz/percent/commit/49ef8ed1eed08228ad12090386aca74233ae83b8))
* ensure SBOM generation job depends on semantic-release completion ([6c57639](https://github.com/sentenz/percent/commit/6c57639f9cf4904578f2b7b03eba4d3a3c1182c2))
* ensure tag_name is set for SBOM artifact publishing ([b522d47](https://github.com/sentenz/percent/commit/b522d47df2a314e460d2ebb7c2d1bbd3f565d243))
* publish SBOM artifacts to GitHub releases ([#58](https://github.com/sentenz/percent/issues/58)) ([c96e509](https://github.com/sentenz/percent/commit/c96e509a561bf27c4dcb3bb3120d51f67cd25240)), closes [#57](https://github.com/sentenz/percent/issues/57)
* publish SBOM to release notes ([170527c](https://github.com/sentenz/percent/commit/170527c339c7d7f68eff151e76fb47772d83500a))
* refine tag format of semantic-release for semantic versioning ([cc0cb68](https://github.com/sentenz/percent/commit/cc0cb681a4cfcdee485acc0f2d274a2187db8078))
* remove conditional trigger for SBOM generation on release events ([755d5a4](https://github.com/sentenz/percent/commit/755d5a4bcaac44a609fb3f350d8c0e78173f4d70))
* remove hidden types from release notes configuration ([c15a938](https://github.com/sentenz/percent/commit/c15a9386fd5c71aa657f5ad54c1d4724ac55cbb1))
* remove redundant SBOM generation job and correct tag_name reference ([2b77f3e](https://github.com/sentenz/percent/commit/2b77f3e69ced7032efd22eab66e1cc3e6182e97e))
* resolve plugin configuration in .releaserc.json ([cffd6f3](https://github.com/sentenz/percent/commit/cffd6f351d792cde7856972fce93f528a151643e))
* resolve semantic-release configuration to adhere conventional commits ([c377588](https://github.com/sentenz/percent/commit/c37758837012305554d9fcef96e6ae164dcfd2f7))
* standardize quotes for condition checks in SBOM generation steps ([eb753d4](https://github.com/sentenz/percent/commit/eb753d481a6326e792815c86d53ed6c5ca704fc7))
* update documentation and improve clarity in SKILL.md and percent.go ([1e82d5d](https://github.com/sentenz/percent/commit/1e82d5dc1ae269e1371a99b986a14804d555d781))
* update release trigger type from 'created' to 'published' in Trivy workflow ([cfe9bc6](https://github.com/sentenz/percent/commit/cfe9bc66da70e502f27551a5d8e1ee3b7aeaede1))
* update semantic-release workflow to remove release trigger and adjust SBOM publishing ([b1c0ddc](https://github.com/sentenz/percent/commit/b1c0ddc94046465354deb1835d2a3a7738bb6086))
* update tag_name reference for SBOM artifact publishing ([1ec77e6](https://github.com/sentenz/percent/commit/1ec77e655bd6967cf4409701cf0bda1cdd96d61d))


### Features

* add Agent Skills documentation for unit, benchmark, fuzz testing, and API documentation ([6017d90](https://github.com/sentenz/percent/commit/6017d902ca355ae1813c152016e6ddef5c7ae242))
* add AGENTS.md for automated unit testing guidelines ([812e990](https://github.com/sentenz/percent/commit/812e990ec324ae3106e73c632824c3142b785891))
* update SKILL.md and percent.go with enhanced formula documentatio ([17769bc](https://github.com/sentenz/percent/commit/17769bc2dd2fbd31a5211166b2e6bb0ab1c8a66a))


### BREAKING CHANGES

* The project's license has changed from MIT to Apache 2.0. Users must ensure compliance with Apache 2.0 requirements.

## [2.3.2](https://github.com/sentenz/percent/compare/2.3.1...2.3.2) (2026-02-03)


### Bug Fixes

* refine tag format of semantic-release for semantic versioning ([cc0cb68](https://github.com/sentenz/percent/commit/cc0cb681a4cfcdee485acc0f2d274a2187db8078))

## [2.3.1](https://github.com/sentenz/percent/compare/2.3.0...2.3.1) (2026-02-03)


### Bug Fixes

* update documentation and improve clarity in SKILL.md and percent.go ([1e82d5d](https://github.com/sentenz/percent/commit/1e82d5dc1ae269e1371a99b986a14804d555d781))

# [2.3.0](https://github.com/sentenz/percent/compare/2.2.0...2.3.0) (2026-02-03)


### Features

* update SKILL.md and percent.go with enhanced formula documentatio ([17769bc](https://github.com/sentenz/percent/commit/17769bc2dd2fbd31a5211166b2e6bb0ab1c8a66a))

# [2.2.0](https://github.com/sentenz/percent/compare/2.1.12...2.2.0) (2026-02-01)


### Features

* add Agent Skills documentation for unit, benchmark, fuzz testing, and API documentation ([6017d90](https://github.com/sentenz/percent/commit/6017d902ca355ae1813c152016e6ddef5c7ae242))

## [2.1.12](https://github.com/sentenz/percent/compare/2.1.11...2.1.12) (2025-12-30)


### Bug Fixes

* update tag_name reference for SBOM artifact publishing ([1ec77e6](https://github.com/sentenz/percent/commit/1ec77e655bd6967cf4409701cf0bda1cdd96d61d))

## [2.1.11](https://github.com/sentenz/percent/compare/2.1.10...2.1.11) (2025-12-30)


### Bug Fixes

* correct tag_name reference for SBOM artifact publishing ([49ef8ed](https://github.com/sentenz/percent/commit/49ef8ed1eed08228ad12090386aca74233ae83b8))
* remove redundant SBOM generation job and correct tag_name reference ([2b77f3e](https://github.com/sentenz/percent/commit/2b77f3e69ced7032efd22eab66e1cc3e6182e97e))
* standardize quotes for condition checks in SBOM generation steps ([eb753d4](https://github.com/sentenz/percent/commit/eb753d481a6326e792815c86d53ed6c5ca704fc7))

## [2.1.10](https://github.com/sentenz/percent/compare/2.1.9...2.1.10) (2025-12-30)


### Bug Fixes

* ensure tag_name is set for SBOM artifact publishing ([b522d47](https://github.com/sentenz/percent/commit/b522d47df2a314e460d2ebb7c2d1bbd3f565d243))

## [2.1.9](https://github.com/sentenz/percent/compare/2.1.8...2.1.9) (2025-12-30)


### Bug Fixes

* publish SBOM artifacts to GitHub releases ([#58](https://github.com/sentenz/percent/issues/58)) ([c96e509](https://github.com/sentenz/percent/commit/c96e509a561bf27c4dcb3bb3120d51f67cd25240)), closes [#57](https://github.com/sentenz/percent/issues/57)

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

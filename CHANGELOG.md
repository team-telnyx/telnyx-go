# Changelog

## 3.7.0 (2025-12-01)

Full Changelog: [v3.6.0...v3.7.0](https://github.com/team-telnyx/telnyx-go/compare/v3.6.0...v3.7.0)

### Features

* Ai 1967 ([75b3cd8](https://github.com/team-telnyx/telnyx-go/commit/75b3cd8c3fb25120c46a13b94285ba71e5b32487))
* Ai 1967 part 2 ([778ec4f](https://github.com/team-telnyx/telnyx-go/commit/778ec4f83f6ae0a3cb4216f446801af3efc5979b))
* ENGDESK-47508 - part 2 shared schema fixes ([f93c6b9](https://github.com/team-telnyx/telnyx-go/commit/f93c6b95d4dfd1e500521e86a32dc25ae59bb538))
* ENGDESK-47518 document mobile number and mobile voice connection endpoints ([638701b](https://github.com/team-telnyx/telnyx-go/commit/638701b1a11ba8d7f0664740f96d0c5b5b06b22f))
* Fix invalid responses ([ca7c8c5](https://github.com/team-telnyx/telnyx-go/commit/ca7c8c55c0e1ce1e07785cd6a35785f100453da1))
* Fix Redocly linting errors and warnings in TDA reporting specs ([97d2790](https://github.com/team-telnyx/telnyx-go/commit/97d2790b41bcd4e1ed1c8bc66ea1ba249871cd0a))
* Fix Redocly linting warnings in Number Lookup spec ([14a0460](https://github.com/team-telnyx/telnyx-go/commit/14a0460b6d09328cab63167546d4d49f99b415f5))
* Fix Redocly linting warnings in OAuth and Integration Secrets specs ([32bfec3](https://github.com/team-telnyx/telnyx-go/commit/32bfec30fe2554d884c9bbc8177202b5f7109bfa))
* Fixing lint errors ([5ff41bb](https://github.com/team-telnyx/telnyx-go/commit/5ff41bbe5e9af1a1f9d9e133465b9d61596253e4))
* MSG-6076: webhook event for 10DLC campaign suspended status ([eabeebe](https://github.com/team-telnyx/telnyx-go/commit/eabeebedf20be16941a3fe981927252e5550f680))
* Refactored README to only contain useful information and reflect accu… ([ca69114](https://github.com/team-telnyx/telnyx-go/commit/ca6911479dfea4cf8e40ebd5dd31732c8cd44ede))
* TELAPPS-5459: Add Azure to transcription start ([b12b8e4](https://github.com/team-telnyx/telnyx-go/commit/b12b8e43a45cf3490706bceb2eb8e6ceb40e837b))
* Updated README to include the step for make buildcontainer bundle to … ([8c6c1df](https://github.com/team-telnyx/telnyx-go/commit/8c6c1df6ffef43ec19916139f5b0cee0f3bd624e))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([938766d](https://github.com/team-telnyx/telnyx-go/commit/938766d6634cb58bb4f8dd49b9e531195c0dd4d6))
* remove readonly parameters from request params ([882538a](https://github.com/team-telnyx/telnyx-go/commit/882538a8830065cea3be9a19dfee2fef8fda2f5c))


### Chores

* bump gjson version ([b2d481b](https://github.com/team-telnyx/telnyx-go/commit/b2d481b629d03ea2b16952aa3f3f4748e8c4e6bf))
* fix empty interfaces ([9bd32a7](https://github.com/team-telnyx/telnyx-go/commit/9bd32a73a620e729412095dbc96bb8e54356067d))

## 3.6.0 (2025-11-03)

Full Changelog: [v3.5.0...v3.6.0](https://github.com/team-telnyx/telnyx-go/compare/v3.5.0...v3.6.0)

### Features

* AI-1842: Add MCP Servers and Integrations sections ([9b90090](https://github.com/team-telnyx/telnyx-go/commit/9b90090aefad621c428224dab71bb84774ea57ae))
* ENGDESK-44767 - Document force remove calls from queue ([7bafd94](https://github.com/team-telnyx/telnyx-go/commit/7bafd940693e16995a976696e548d9570dbe53c4))
* ENGDESK-45429 - Add sip_region documentation for dial and transfer command ([676e3e6](https://github.com/team-telnyx/telnyx-go/commit/676e3e6a179adecd3b1112ea5777af5a49433578))
* ENGDESK-46399 - Add sip_call_id filter for retreiving recordings ([b7fa58d](https://github.com/team-telnyx/telnyx-go/commit/b7fa58da7364610ae1dd79598c59431296c0e7a5))
* TELAPPS-5399 Add region to conference commands ([f81c9c7](https://github.com/team-telnyx/telnyx-go/commit/f81c9c7705fd1805ec6f15121d75f5df857a75b9))
* TELAPPS-ENGDESK-46395 Add keep_after_hangup to enqueue command ([1105462](https://github.com/team-telnyx/telnyx-go/commit/1105462342509db20c0d50f6480f966234e92bbb))
* TELAPPS-ENGDESK-46395 Add PATCH /queues/{queue_name}/calls/{call_control_id} endpoint ([11a6b03](https://github.com/team-telnyx/telnyx-go/commit/11a6b03e6d95857c1ec882b1944ccd9c73c98fd2))


### Bug Fixes

* **client:** fix issue with example webhook payload ([e5a9631](https://github.com/team-telnyx/telnyx-go/commit/e5a9631ed59080df3fa2edb14de6b0f8169633fc))


### Chores

* **internal:** grammar fix (it's -&gt; its) ([64a41b0](https://github.com/team-telnyx/telnyx-go/commit/64a41b02454d68e38f48a65170e672d592e2d605))

## 3.5.0 (2025-10-27)

Full Changelog: [v3.4.0...v3.5.0](https://github.com/team-telnyx/telnyx-go/compare/v3.4.0...v3.5.0)

### Features

* **api:** add method to upload JSON documents ([615048c](https://github.com/team-telnyx/telnyx-go/commit/615048c0379fc70ec04590b6bf5a46cc340c1829))
* **api:** added webhook public key ([cd05c9d](https://github.com/team-telnyx/telnyx-go/commit/cd05c9d02c81da3b9d0453ec26cd58f472b203ab))
* **api:** manual updates ([840698d](https://github.com/team-telnyx/telnyx-go/commit/840698db717d114e98e6a6f475721429ea259b54))
* **api:** manual updates ([f719b16](https://github.com/team-telnyx/telnyx-go/commit/f719b168ca5bf37cd9854eb2366d3832f2f19528))
* **api:** manual updates ([15fba7a](https://github.com/team-telnyx/telnyx-go/commit/15fba7a16dabaefa71fe2f9feafab3649bee9b9d))
* define more models ([7bfb5c4](https://github.com/team-telnyx/telnyx-go/commit/7bfb5c4c22caa21403f7d42123c069c5c5af3561))


### Documentation

* update README description with comprehensive SDK overview ([#3](https://github.com/team-telnyx/telnyx-go/issues/3)) ([636b60c](https://github.com/team-telnyx/telnyx-go/commit/636b60c80e542d309de4dc65416c125d4e79e613))

## 3.4.0 (2025-10-16)

Full Changelog: [v3.3.0...v3.4.0](https://github.com/team-telnyx/telnyx-go/compare/v3.3.0...v3.4.0)

### Features

* ENGDESK-45836: Document private endpoint for republishing account events ([67d8410](https://github.com/team-telnyx/telnyx-go/commit/67d84108da524ae639a03ba1129fd3e3ab695d20))
* Fix broken link to List SIM card action ([3900660](https://github.com/team-telnyx/telnyx-go/commit/390066093def39560747c4fa07f4e26fd6cac499))
* MSG-5978: Add BRN fields to toll-free verification OpenAPI specs ([8608b0e](https://github.com/team-telnyx/telnyx-go/commit/8608b0ec46429e8e4583c23eb051d94c0084e414))

## 3.3.0 (2025-10-09)

Full Changelog: [v3.2.0...v3.3.0](https://github.com/team-telnyx/telnyx-go/compare/v3.2.0...v3.3.0)

### Features

* Retell assistants import ([f44fb16](https://github.com/team-telnyx/telnyx-go/commit/f44fb16c42564811398675d81458d7ad9c3fd5d3))


### Bug Fixes

* reference models more accurately ([1a21833](https://github.com/team-telnyx/telnyx-go/commit/1a2183324599d907964e597f528bd760c16eb28f))

## 3.2.0 (2025-10-06)

Full Changelog: [v3.1.0...v3.2.0](https://github.com/team-telnyx/telnyx-go/compare/v3.1.0...v3.2.0)

### Features

* ENGDESK-45343: Add CustomHeaders documentation to TeXML Dial endpoints ([d88d857](https://github.com/team-telnyx/telnyx-go/commit/d88d8571a2e88f4616bd984278b647703115bd5c))
* MSG-5944: added mobile_only field to messaging profiles ([ff50db2](https://github.com/team-telnyx/telnyx-go/commit/ff50db27726cec5faacda9df532d9fa103ca8260))

## 3.1.0 (2025-10-02)

Full Changelog: [v3.0.0...v3.1.0](https://github.com/team-telnyx/telnyx-go/compare/v3.0.0...v3.1.0)

### Features

* **api:** manual updates ([4b1795c](https://github.com/team-telnyx/telnyx-go/commit/4b1795c55b366192a56adcace1040571ee52c050))
* Engdesk 44932 ([1f598f1](https://github.com/team-telnyx/telnyx-go/commit/1f598f1d138ea06138e6864082796c093919010c))
* Fix listing deepgram languages for transcription start ([26297d6](https://github.com/team-telnyx/telnyx-go/commit/26297d6da9ea37e8acd38f8b3f668169eb37d2c8))
* TELAPPS-5367: Update transcription start docs ([ed74c02](https://github.com/team-telnyx/telnyx-go/commit/ed74c02d5df98f587ea8ca51fcb60d9050beb007))


### Chores

* configure new SDK language ([621692e](https://github.com/team-telnyx/telnyx-go/commit/621692e2c21b74541d6167a5748729bab2f58832))
* sync repo ([301e6e0](https://github.com/team-telnyx/telnyx-go/commit/301e6e0f3028bbdd6bbefc0edac334af09a677aa))

## 3.0.0 (2025-09-24)

Full Changelog: [v0.0.1...v3.0.0](https://github.com/team-telnyx/telnyx-go/compare/v0.0.1...v3.0.0)

### ⚠ BREAKING CHANGES

* **api:** extract APIError to shared models

### Features

* **api:** extract APIError to shared models ([047df10](https://github.com/team-telnyx/telnyx-go/commit/047df107a0e129a71db1a41f4edbd556acf1219a))


### Bug Fixes

* **client:** elide enum validators for non-json structs ([1fe6844](https://github.com/team-telnyx/telnyx-go/commit/1fe6844f47a0cdc99db123a31cf3ee04dcfe06a9))


### Chores

* sync repo ([47c65ce](https://github.com/team-telnyx/telnyx-go/commit/47c65ce66fb46253601d34ebebf57848270a82a3))

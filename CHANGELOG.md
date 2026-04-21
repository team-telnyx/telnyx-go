# Changelog

## 4.60.0 (2026-04-20)

Full Changelog: [v4.59.0...v4.60.0](https://github.com/team-telnyx/telnyx-go/compare/v4.59.0...v4.60.0)

### Features

* Add keyterm field to TranscriptionSettingsConfig ([f0f24fe](https://github.com/team-telnyx/telnyx-go/commit/f0f24fe17b8ed971b1f79cfaffea35e3d61a198e))

## 4.59.0 (2026-04-17)

Full Changelog: [v4.58.0...v4.59.0](https://github.com/team-telnyx/telnyx-go/compare/v4.58.0...v4.59.0)

### Features

* Add user_idle_reply_secs to TelephonySettings spec ([0d8ed6f](https://github.com/team-telnyx/telnyx-go/commit/0d8ed6f8b456ab3c7e9c9c3715c966fa072d75a9))
* **client:** update apidata interface ([13dc2fb](https://github.com/team-telnyx/telnyx-go/commit/13dc2fbca4fb88b5d4dfab517119c5787ddccb6f))
* Lower user_idle_timeout_secs minimum from 30s to 10s ([8b822f8](https://github.com/team-telnyx/telnyx-go/commit/8b822f82c794aafb7085c04b78aa4b825b475ca8))
* TELAPPS Provide description what params can be used for premium amd ([12b7550](https://github.com/team-telnyx/telnyx-go/commit/12b755046ce6dd1c372731937a9e9824a1d55dcb))

## 4.58.0 (2026-04-16)

Full Changelog: [v4.57.1...v4.58.0](https://github.com/team-telnyx/telnyx-go/compare/v4.57.1...v4.58.0)

### Features

* [TDA-6425] Fix session analysis API spec: relaxed date_time, remove status & completed_at ([cff36ab](https://github.com/team-telnyx/telnyx-go/commit/cff36abea565824e25ed18bbf77ea1b953a1d632))
* Add webhook_urls, webhook_urls_method, webhook_retries_policies to Dial endpoint ([cd9077d](https://github.com/team-telnyx/telnyx-go/commit/cd9077db39767a3c8eb92b06649d3d7f4fa39f65))
* MSG-6868: document whitelisted_destinations as conditionally required ([8519e0b](https://github.com/team-telnyx/telnyx-go/commit/8519e0b3fc86324ec901be234c333c024bc15aa0))
* TELAPPS-5712: Add deepfake detection to call-control API spec ([d48d0d1](https://github.com/team-telnyx/telnyx-go/commit/d48d0d192ab3b046985c494ce066790eaeac2ee8))


### Bug Fixes

* correct VoiceCloneNewFromUploadParams field reference in MarshalMultipart ([d70642b](https://github.com/team-telnyx/telnyx-go/commit/d70642b6e269519926c9255ada0fc1da9ad6e931))


### Documentation

* add pagination params to conversation messages endpoint ([e32fae9](https://github.com/team-telnyx/telnyx-go/commit/e32fae95eb64378b3e92a59fd170fb0dd7908c55))

## 4.57.1 (2026-04-12)

Full Changelog: [v4.57.0...v4.57.1](https://github.com/team-telnyx/telnyx-go/compare/v4.57.0...v4.57.1)

### Bug Fixes

* set additionalProperties=false on VoiceCloneUploadRequest to prevent codegen errors ([ba0aa93](https://github.com/team-telnyx/telnyx-go/commit/ba0aa9333684a82de760e2decf8cfaeb15095f39))


### Refactors

* separate custom webhook methods from generated code ([e0895e5](https://github.com/team-telnyx/telnyx-go/commit/e0895e52cecaebde3a67ad6f9a3af608f40cca47))

## 4.57.0 (2026-04-11)

Full Changelog: [v4.56.0...v4.57.0](https://github.com/team-telnyx/telnyx-go/compare/v4.56.0...v4.57.0)

### Features

* add shared CallAssistantRequest schema for call-control assistant object ([d7f4420](https://github.com/team-telnyx/telnyx-go/commit/d7f4420c6537c4c779512deb0f380729e1d3b655))
* **api:** manual updates ([719d033](https://github.com/team-telnyx/telnyx-go/commit/719d033cda6244cde115f9497747e415ef34d9d2))
* **api:** Merge pull request [#46](https://github.com/team-telnyx/telnyx-go/issues/46) from stainless-sdks/FixModelRecommendation ([35f93a3](https://github.com/team-telnyx/telnyx-go/commit/35f93a31250aab1b11b5073108e2a4557883b96f))
* CW-3815 fix PATCH /wirelss_blocklists/{id} endpoint ([a4c0b0d](https://github.com/team-telnyx/telnyx-go/commit/a4c0b0d4b9ab4171ed66fcd6e0f7c44c96762c66))
* MSG-6846: add GET /profile/photo docs for whatsapp API ([f19fac6](https://github.com/team-telnyx/telnyx-go/commit/f19fac6442ba4b8830ced2211c024436e1962f4e))
* TELAPPS-5689: Pronunciation dictionaries API docs ([1f643ea](https://github.com/team-telnyx/telnyx-go/commit/1f643ea419110596d56aa028a0a357492d60a1a3))
* TELAPPS-5707: Add privacy parameter to Call Control dial and transfer ([769f3aa](https://github.com/team-telnyx/telnyx-go/commit/769f3aaf33d9fb90ab8ad8a4a20e14a170fb7600))


### Reverts

* restore stainless.yml to pre-6a6df5b state ([891f9f4](https://github.com/team-telnyx/telnyx-go/commit/891f9f44e0fb44045ab51954e11b31f3deeff0b7))
* revert stainless.yml changes from 9c5e8d8 ([4b0926f](https://github.com/team-telnyx/telnyx-go/commit/4b0926f13c015154da5eace34746c0ddcd8f59a1))
* revert stainless.yml changes from pronunciation dictionaries commit ([d4b2d27](https://github.com/team-telnyx/telnyx-go/commit/d4b2d27977d4fedc5d45e65e6b7e3b5072cce75a))


### Chores

* **internal:** version bump ([676c7c4](https://github.com/team-telnyx/telnyx-go/commit/676c7c40d79a813eeb3337f88776bc28e2b5f8dc))


### Documentation

* improve examples ([a6c09d6](https://github.com/team-telnyx/telnyx-go/commit/a6c09d60e63b3d209f8b53e8c6b265f9b956efb9))
* update voice clone schemas to match Ultra/model_id implementation ([2118703](https://github.com/team-telnyx/telnyx-go/commit/2118703868d58f253b1b86db54b4314065fd9162))

## 4.57.0 (2026-04-09)

Full Changelog: [v4.56.0...v4.57.0](https://github.com/team-telnyx/telnyx-go/compare/v4.56.0...v4.57.0)

### Features

* Add oneOf constraint for Url/Texml mutual exclusivity in InitiateCallRequest ([190ec78](https://github.com/team-telnyx/telnyx-go/commit/190ec780c665897724a757bc2bbb7510bf61ae22))

## 4.56.0 (2026-04-08)

Full Changelog: [v4.55.2...v4.56.0](https://github.com/team-telnyx/telnyx-go/compare/v4.55.2...v4.56.0)

### Features

* Add ai_calls endpoint documentation to OpenAPI spec ([fda8750](https://github.com/team-telnyx/telnyx-go/commit/fda875063c485f442f6558e2ae5d89f121570621))
* add enabled boolean to recording_settings [AI-2178] ([5492e73](https://github.com/team-telnyx/telnyx-go/commit/5492e735ab2d9120b337308cdb8465c5bb895999))
* Add oneOf constraint for Url/Texml mutual exclusivity in InitiateCallRequest ([8bebb65](https://github.com/team-telnyx/telnyx-go/commit/8bebb65ea7e4d9e31e1d8bc2e0a4c94d5c6d4adc))
* AI-2180: Add message_template to SendMessageTool schema ([5ac41fa](https://github.com/team-telnyx/telnyx-go/commit/5ac41faeddad4d27a86a20772ad2b73c9e52c52f))
* **api:** manual updates ([c8d37fe](https://github.com/team-telnyx/telnyx-go/commit/c8d37fe5b28455d1d5872d693b1aa37387fb58e6))
* **api:** Merge pull request [#39](https://github.com/team-telnyx/telnyx-go/issues/39) from stainless-sdks/revert-a988c49-stainless-changes ([f2c0f4f](https://github.com/team-telnyx/telnyx-go/commit/f2c0f4f1ba523ce5038caeec99d66ca330f91a75))
* Assistants: add observability ([5318f99](https://github.com/team-telnyx/telnyx-go/commit/5318f99fe333cd662db7eefe4fcc23e99c670134))
* MSG-6666: Add template and text properties to WhatsApp send message schema ([5c7b934](https://github.com/team-telnyx/telnyx-go/commit/5c7b9345be30ead7fa1619bad70c74ff77cdb868))
* MSG-6673: Add WhatsApp verification endpoint and profile settings ([cbb9984](https://github.com/team-telnyx/telnyx-go/commit/cbb9984951767dba886f028ff5460daf1e01da1c))


### Reverts

* restore stainless.yml SDK generation fixes ([f2c0f4f](https://github.com/team-telnyx/telnyx-go/commit/f2c0f4f1ba523ce5038caeec99d66ca330f91a75))

## 4.55.2 (2026-04-02)

Full Changelog: [v4.55.1...v4.55.2](https://github.com/team-telnyx/telnyx-go/compare/v4.55.1...v4.55.2)

### Bug Fixes

* fixes for pagination and iteration, plus iter.Seq support ([22c5ffc](https://github.com/team-telnyx/telnyx-go/commit/22c5ffcbf9b032f12fe4770abc7b69018cfa9942))

## 4.55.1 (2026-03-31)

Full Changelog: [v4.55.0...v4.55.1](https://github.com/team-telnyx/telnyx-go/compare/v4.55.0...v4.55.1)

### Bug Fixes

* fix issue with unmarshaling in some cases ([0d4afda](https://github.com/team-telnyx/telnyx-go/commit/0d4afda8991038c2b9676360eeff676dcc7d6978))


### Chores

* **ci:** support opting out of skipping builds on metadata-only commits ([3406f7d](https://github.com/team-telnyx/telnyx-go/commit/3406f7d7b1a732ab692ca939c0bab332ed7533d7))
* update docs for api:"required" ([49e614c](https://github.com/team-telnyx/telnyx-go/commit/49e614cd13c55041760c8993e2c95b05073db954))


### Documentation

* fix voice settings available voices link ([c4d6bfa](https://github.com/team-telnyx/telnyx-go/commit/c4d6bfa48f6b898fbc8f8ca0bf0cff4d40c12e3d))

## 4.55.0 (2026-03-26)

Full Changelog: [v4.54.0...v4.55.0](https://github.com/team-telnyx/telnyx-go/compare/v4.54.0...v4.55.0)

### Features

* **internal:** support comma format in multipart form encoding ([df166a3](https://github.com/team-telnyx/telnyx-go/commit/df166a35358164b24e810ce12ea13b1b0bd832e2))


### Bug Fixes

* prevent duplicate ? in query params ([89e80d1](https://github.com/team-telnyx/telnyx-go/commit/89e80d12547e98b2ed8ccbadd5e3ae40e88072fa))


### Chores

* remove unnecessary error check for url parsing ([56a5052](https://github.com/team-telnyx/telnyx-go/commit/56a505257d97beea70839c2f1c842535e4b95c58))

## 4.54.0 (2026-03-25)

Full Changelog: [v4.53.0...v4.54.0](https://github.com/team-telnyx/telnyx-go/compare/v4.53.0...v4.54.0)

### Features

* **websocket:** add STT/TTS WebSocket streaming support ([4f59511](https://github.com/team-telnyx/telnyx-go/commit/4f595110ac958c275c84fa9dd16a272ba34b11b5))


### Bug Fixes

* rename number-reputation ToS route to use underscores ([694fb77](https://github.com/team-telnyx/telnyx-go/commit/694fb7792a1822affd01bf51b11bbab3236eb345))


### Chores

* **client:** fix multipart serialisation of Default() fields ([9e422c9](https://github.com/team-telnyx/telnyx-go/commit/9e422c9c2c833e3dd3f22247aeb9e5a3defe00d0))
* **internal:** support default value struct tag ([3649c34](https://github.com/team-telnyx/telnyx-go/commit/3649c343caf487301d1984e00e1e72190905316f))


### Documentation

* **branded-calling:** add Number Reputation API specs ([70e90a3](https://github.com/team-telnyx/telnyx-go/commit/70e90a30b8a0bb27bde20ba3969de0f21ae8788c))

## 4.53.0 (2026-03-25)

Full Changelog: [v4.52.0...v4.53.0](https://github.com/team-telnyx/telnyx-go/compare/v4.52.0...v4.53.0)

### Features

* **api:** manual updates ([39ed85c](https://github.com/team-telnyx/telnyx-go/commit/39ed85cbb4cbc79057d65184ff85903101639e02))
* **api:** manual updates ([e7d17ab](https://github.com/team-telnyx/telnyx-go/commit/e7d17ab7851961c43b2a22133c2bbc91f44a0b9b))
* **api:** Merge pull request [#30](https://github.com/team-telnyx/telnyx-go/issues/30) from stainless-sdks/fix-schemaUnionDiscriminatorMissing ([c0c650e](https://github.com/team-telnyx/telnyx-go/commit/c0c650ec463a2087d393072df14d57bc78853abf))
* New tools api ([fa8d724](https://github.com/team-telnyx/telnyx-go/commit/fa8d724e524641f80e8845aeb9cdf6224c870451))
* TELAPPS-5685: Add store_fields_as_variables to WebhookToolParams ([3c0f71a](https://github.com/team-telnyx/telnyx-go/commit/3c0f71acaea7e95dfd0c527743ff476bf9b99d30))


### Chores

* **ci:** skip lint on metadata-only changes ([54db857](https://github.com/team-telnyx/telnyx-go/commit/54db857938ef8514d66199e06a720322f4d06a5c))

## 4.52.0 (2026-03-23)

Full Changelog: [v4.51.0...v4.52.0](https://github.com/team-telnyx/telnyx-go/compare/v4.51.0...v4.52.0)

### Features

* **api:** manual updates ([5dae930](https://github.com/team-telnyx/telnyx-go/commit/5dae930d20a986f528379d475b97255b7e0aceda))


### Chores

* **internal:** update gitignore ([4c9fe9f](https://github.com/team-telnyx/telnyx-go/commit/4c9fe9f642f92f4e91a72fd479c86ea8498fe797))

## 4.51.0 (2026-03-20)

Full Changelog: [v4.50.0...v4.51.0](https://github.com/team-telnyx/telnyx-go/compare/v4.50.0...v4.51.0)

### Features

* Add Minimax provider support to Voice Designs and Voice Clones API spec ([bec9d43](https://github.com/team-telnyx/telnyx-go/commit/bec9d43284fa65f425598be10bfa375f252efee7))
* **api:** manual updates ([43d12db](https://github.com/team-telnyx/telnyx-go/commit/43d12db0a5b1f4861de113a682cdbc72c0731a4c))
* **api:** manual updates ([ece1771](https://github.com/team-telnyx/telnyx-go/commit/ece177187db12ed07b9d11c29876fb7087447239))
* **api:** manual updates ([3e9dd03](https://github.com/team-telnyx/telnyx-go/commit/3e9dd039e2c9ca0039783dc53321fa26e5c241cc))
* **api:** manual updates ([76696ea](https://github.com/team-telnyx/telnyx-go/commit/76696ea5fc52ce63e1f26b8d176ea9a05099eece))
* **wireless:** add traffic policy profiles endpoints to OpenAPI spec ([87f32c2](https://github.com/team-telnyx/telnyx-go/commit/87f32c2d4ab04e80585f420f3363f5d54156fd0d))


### Documentation

* WhatsApp template components schema ([25dfdda](https://github.com/team-telnyx/telnyx-go/commit/25dfddac0f5175bdd1eb687b4eb6cde364ffe43f))


### Refactors

* move webhook verification to lib package to avoid merge conflicts ([c238d65](https://github.com/team-telnyx/telnyx-go/commit/c238d655ed20417cbb88e9dad6bd1b3b2db725ad))

## 4.44.0 (2026-03-10)

Full Changelog: [v4.43.0...v4.44.0](https://github.com/team-telnyx/telnyx-go/compare/v4.43.0...v4.44.0)

### Features

* **api:** manual updates ([156c285](https://github.com/team-telnyx/telnyx-go/commit/156c2855302082ce14bf76a1b4f3197c26ed6ab5))
* Assistant tags ([3c74d33](https://github.com/team-telnyx/telnyx-go/commit/3c74d33e99c36c8215bca6ce25ccf988bdaa47a2))
* CW-2881 publish wireless VoLTE docs to prod ([9c16512](https://github.com/team-telnyx/telnyx-go/commit/9c1651289e816252a022ca6c24b80f06531a7dd8))
* MSG-6418: remove flashcall from hosted number verification codes endpoint ([64284b4](https://github.com/team-telnyx/telnyx-go/commit/64284b46640ade8533493156897542dd3a17a0cf))
* TELAPPS-ENGDESK-49737 Add prevent_double_bridge param to dial ([a455f5c](https://github.com/team-telnyx/telnyx-go/commit/a455f5ca425a3ece6f476bc25eb3df87f8f0a396))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([626f387](https://github.com/team-telnyx/telnyx-go/commit/626f387fe9fde8b69874f219251f5f79af2c897e))
* **internal:** minor cleanup ([5424a9f](https://github.com/team-telnyx/telnyx-go/commit/5424a9ff200b05c8816f46327ecaff7f9f1ff49c))
* **internal:** use explicit returns ([ae07632](https://github.com/team-telnyx/telnyx-go/commit/ae0763212ccbe13761dd7b653f053f8caa461cc5))
* update placeholder string ([0175443](https://github.com/team-telnyx/telnyx-go/commit/01754435c64ece5a3011fe9aa25feac8c6e7cca7))

## 4.43.0 (2026-03-05)

Full Changelog: [v4.42.0...v4.43.0](https://github.com/team-telnyx/telnyx-go/compare/v4.42.0...v4.43.0)

### Features

* Changing a tag for filebased STT endpoint ([ba25f80](https://github.com/team-telnyx/telnyx-go/commit/ba25f8051f3c296c15d23ec7c4ab204d3e589d49))


### Documentation

* **messaging:** Add wait_seconds to message response schemas ([0145fd3](https://github.com/team-telnyx/telnyx-go/commit/0145fd32934dddd7df16360b7cd61042fa5a1d72))

## 4.42.0 (2026-03-05)

Full Changelog: [v4.41.0...v4.42.0](https://github.com/team-telnyx/telnyx-go/compare/v4.41.0...v4.42.0)

### Features

* **api:** manual updates ([55dd97a](https://github.com/team-telnyx/telnyx-go/commit/55dd97aeaaf1a2cf266b2214d983dcce6e30d0e7))
* **api:** manual updates ([7f276c7](https://github.com/team-telnyx/telnyx-go/commit/7f276c75ab6c935ef331006edc90e179d7c56ba2))
* **api:** manual updates ([50f631d](https://github.com/team-telnyx/telnyx-go/commit/50f631ddce3328ed5ae963f929d87dc03576b65d))
* **api:** manual updates ([46a26ed](https://github.com/team-telnyx/telnyx-go/commit/46a26ed340f5cb1a50ef75bd3c0880fff1bd1b2b))
* **api:** manual updates ([090607f](https://github.com/team-telnyx/telnyx-go/commit/090607fe5a6f4b9de15c377ae09e56425fba157f))
* **api:** manual updates ([6953d48](https://github.com/team-telnyx/telnyx-go/commit/6953d48daf8d0f062578bdc69167040a3ef55d56))
* **stt:** add WebSocket event schemas for Stainless SDK generation ([30858ac](https://github.com/team-telnyx/telnyx-go/commit/30858ac548922104708e0e6f4ac2c7f5311e6610))


### Bug Fixes

* fix request delays for retrying to be more respectful of high requested delays ([07c4cec](https://github.com/team-telnyx/telnyx-go/commit/07c4cecfd18d885b47a8ce44d4dd0af2d0fc37a0))

## 4.41.0 (2026-03-04)

Full Changelog: [v4.40.0...v4.41.0](https://github.com/team-telnyx/telnyx-go/compare/v4.40.0...v4.41.0)

### Features

* **api:** manual updates ([5c36d94](https://github.com/team-telnyx/telnyx-go/commit/5c36d94b80b018ff8fb5cd2e551254508ab8c374))
* **api:** manual updates ([45aec2d](https://github.com/team-telnyx/telnyx-go/commit/45aec2d45dd58e0f36b92e5916b59df4017019e8))
* **api:** manual updates ([4a2a2a1](https://github.com/team-telnyx/telnyx-go/commit/4a2a2a1ffd0ce15f020a05276479e8e8d087a7ff))
* **api:** manual updates ([c30b2d3](https://github.com/team-telnyx/telnyx-go/commit/c30b2d3cad4a4470cbdc3fb4141d9d399694bf8f))

## 4.40.0 (2026-03-04)

Full Changelog: [v4.39.0...v4.40.0](https://github.com/team-telnyx/telnyx-go/compare/v4.39.0...v4.40.0)

### Features

* **api:** manual updates ([24bbede](https://github.com/team-telnyx/telnyx-go/commit/24bbede1f44d1ddf01de4cd728cfd5bc94aff156))


### Bug Fixes

* add discriminator to TtsServerEvent for Stainless SDK generation ([cb741ee](https://github.com/team-telnyx/telnyx-go/commit/cb741ee3032ca1aa6753eb32777d731c6bfb90ab))

## 4.39.0 (2026-03-03)

Full Changelog: [v4.38.0...v4.39.0](https://github.com/team-telnyx/telnyx-go/compare/v4.38.0...v4.39.0)

### Features

* [TDA-6425] Add Session Analysis API spec to public docs ([19c208e](https://github.com/team-telnyx/telnyx-go/commit/19c208ec8a07642c4536c21927d7cfd872027389))
* **api:** manual updates ([ba56cae](https://github.com/team-telnyx/telnyx-go/commit/ba56cae11a527d4a47af1e3a988177f3808c7ef2))
* **api:** manual updates ([f039a3b](https://github.com/team-telnyx/telnyx-go/commit/f039a3b2e64b14e9a6358de81ebe983c6af0a5a6))

## 4.38.0 (2026-03-03)

Full Changelog: [v4.37.0...v4.38.0](https://github.com/team-telnyx/telnyx-go/compare/v4.37.0...v4.38.0)

### Features

* AI-2106: Add invite tool schema to inference OpenAPI spec ([dc022c5](https://github.com/team-telnyx/telnyx-go/commit/dc022c5b9b84e25f82fed040fd42eb7636474548))
* **api:** manual updates ([a276b8c](https://github.com/team-telnyx/telnyx-go/commit/a276b8ce08732afcd75db68fdde6c6ba5319e5ad))
* **api:** manual updates ([d80b22f](https://github.com/team-telnyx/telnyx-go/commit/d80b22f320dd8bd64f1493dd09be9020aaa21e9a))
* Changing the tag for TTS endpoint ([2696d43](https://github.com/team-telnyx/telnyx-go/commit/2696d43be1a53528713a5bb3267852e06f0259ce))


### Chores

* **tests:** update webhook tests ([df6882c](https://github.com/team-telnyx/telnyx-go/commit/df6882c2f53d85bbee9789241b36180dd7d6f023))

## 4.37.0 (2026-03-02)

Full Changelog: [v4.36.1...v4.37.0](https://github.com/team-telnyx/telnyx-go/compare/v4.36.1...v4.37.0)

### Features

* Merge TTS file-based spec into text-to-speech.json ([8368964](https://github.com/team-telnyx/telnyx-go/commit/8368964be47f11fb5121fd980735617eee12d8f4))


### Bug Fixes

* narrow porting event_type enums for SDK discriminator support ([64c4319](https://github.com/team-telnyx/telnyx-go/commit/64c431986412ab6bbfd3f9fa34071507fc4bf0d4))

## 4.36.1 (2026-03-02)

Full Changelog: [v4.36.0...v4.36.1](https://github.com/team-telnyx/telnyx-go/compare/v4.36.0...v4.36.1)

### Bug Fixes

* fix test cases where required client options were not set ([d1d383a](https://github.com/team-telnyx/telnyx-go/commit/d1d383acc14162f74079e26563501efde22ea91a))

## 4.36.0 (2026-02-28)

Full Changelog: [v4.35.0...v4.36.0](https://github.com/team-telnyx/telnyx-go/compare/v4.35.0...v4.36.0)

### Features

* **api:** manual updates ([ed0509f](https://github.com/team-telnyx/telnyx-go/commit/ed0509f68bbe1fe076567213c52930b6b713dac2))


### Chores

* **ci:** add build step ([b4a183e](https://github.com/team-telnyx/telnyx-go/commit/b4a183e894e06cf09b77dfce05b083acc6f33eaf))
* **docs:** add missing descriptions ([fc30b09](https://github.com/team-telnyx/telnyx-go/commit/fc30b09a9b599aac96c923a9f24a35b981da8bc6))

## 4.35.0 (2026-02-27)

Full Changelog: [v4.34.0...v4.35.0](https://github.com/team-telnyx/telnyx-go/compare/v4.34.0...v4.35.0)

### Features

* Add TTS file-based endpoint spec ([26587e5](https://github.com/team-telnyx/telnyx-go/commit/26587e5b35c6edf2bf325a80e019a257e9f9c45b))
* **api:** manual updates ([bb9f4ad](https://github.com/team-telnyx/telnyx-go/commit/bb9f4ad3954d0ecb0327add86d27890e7b2236c3))
* **api:** manual updates ([b6ceabe](https://github.com/team-telnyx/telnyx-go/commit/b6ceabe43f90fdff66dc3e93988dd21b504b53b6))
* **api:** manual updates ([69d8bcf](https://github.com/team-telnyx/telnyx-go/commit/69d8bcf92c50627f562bb89ab19661c45d42a139))

## 4.34.0 (2026-02-26)

Full Changelog: [v4.33.0...v4.34.0](https://github.com/team-telnyx/telnyx-go/compare/v4.33.0...v4.34.0)

### Features

* **api:** manual updates ([9bd40c8](https://github.com/team-telnyx/telnyx-go/commit/9bd40c855433d80a96e498f5461857537fbd05a5))
* **api:** manual updates ([9b02b25](https://github.com/team-telnyx/telnyx-go/commit/9b02b25475d882df4dacb901e52d6dda3cb88b88))


### Chores

* bring back other changes ([3d5b673](https://github.com/team-telnyx/telnyx-go/commit/3d5b673b6a2a3b84fda81d3810ed1d22edad7ddc))

## 4.33.0 (2026-02-26)

Full Changelog: [v4.32.0...v4.33.0](https://github.com/team-telnyx/telnyx-go/compare/v4.32.0...v4.33.0)

### Features

* Add text-to-speech WebSocket streaming OpenAPI spec ([843b234](https://github.com/team-telnyx/telnyx-go/commit/843b23495b47df199ccdfadaf2899904016ca0c9))
* **api:** manual updates ([23b6c36](https://github.com/team-telnyx/telnyx-go/commit/23b6c36458390ffd251b33c3e4fc987ba2a1c0f6))
* **api:** manual updates ([f84640f](https://github.com/team-telnyx/telnyx-go/commit/f84640f161d7b2ceb397c4f8feb09adabce290f8))
* TELAPPS-ENGDESK-48951 add channels to conf record start ([20421a4](https://github.com/team-telnyx/telnyx-go/commit/20421a4cd8cc8396ce18c2efabc7584512859669))

## 4.32.0 (2026-02-25)

Full Changelog: [v4.31.0...v4.32.0](https://github.com/team-telnyx/telnyx-go/compare/v4.31.0...v4.32.0)

### Features

* PORTAL-5923: Add stored_payment_transactions endpoint to OpenAPI docs ([5c565eb](https://github.com/team-telnyx/telnyx-go/commit/5c565eb70d4b7ed940f2aff53cff624e69b74f7c))


### Documentation

* **call-control:** Add missing params to hangup, bridge, answer ([c3a6ea4](https://github.com/team-telnyx/telnyx-go/commit/c3a6ea4093368e8fb7c17882fa55f2f29d11a866))
* **call-control:** Add queue CRUD endpoints ([1497f4a](https://github.com/team-telnyx/telnyx-go/commit/1497f4a3d382aef18ecfb1a4cacdc5699e25a715))

## 4.31.0 (2026-02-25)

Full Changelog: [v4.30.1...v4.31.0](https://github.com/team-telnyx/telnyx-go/compare/v4.30.1...v4.31.0)

### Features

* Add missing TTS voice settings schemas and update voice descriptions ([5764583](https://github.com/team-telnyx/telnyx-go/commit/576458359f9b3c27c18df38fcb9ed4f0c282ec01))
* **api:** manual updates ([caa6b17](https://github.com/team-telnyx/telnyx-go/commit/caa6b173688e5e87466eb9f0125cb87fd7e5b832))
* TELAPPS Add prevent_double_bridge to bridge command ([41f35ab](https://github.com/team-telnyx/telnyx-go/commit/41f35abf9e7794c8355f45fef386a926cc9e48b1))

## 4.30.1 (2026-02-24)

Full Changelog: [v4.30.0...v4.30.1](https://github.com/team-telnyx/telnyx-go/compare/v4.30.0...v4.30.1)

### Chores

* **internal:** move custom custom `json` tags to `api` ([26fdddf](https://github.com/team-telnyx/telnyx-go/commit/26fdddf04417aa259dfd56758bc429692e5818e2))

## 4.30.0 (2026-02-24)

Full Changelog: [v4.29.0...v4.30.0](https://github.com/team-telnyx/telnyx-go/compare/v4.29.0...v4.30.0)

### Features

* AI-2093: Add language_boost to MiniMax voice settings ([fcfef69](https://github.com/team-telnyx/telnyx-go/commit/fcfef6932be3d64e98e925d524709a035bdefaec))
* **api:** manual updates ([d58966b](https://github.com/team-telnyx/telnyx-go/commit/d58966b0d04598527d91f9058beea61cf8985bb0))
* **api:** manual updates ([fb6e833](https://github.com/team-telnyx/telnyx-go/commit/fb6e833f89d05086625deafff87a8f1b9a2b2458))
* **api:** manual updates ([bba09e9](https://github.com/team-telnyx/telnyx-go/commit/bba09e9b73ab11a3a6d8df095d8502a85c23f806))
* **api:** manual updates ([61d85b1](https://github.com/team-telnyx/telnyx-go/commit/61d85b11fee99df526fc8c734ad782d88690ab16))
* **api:** manual updates ([4240e9e](https://github.com/team-telnyx/telnyx-go/commit/4240e9e081050f68b30af2bbeb48be285fa31f82))
* **api:** manual updates ([afc7ce6](https://github.com/team-telnyx/telnyx-go/commit/afc7ce6631df7d7d43a7e8dd4685141b2f0ef507))
* **api:** manual updates ([b9d5ab6](https://github.com/team-telnyx/telnyx-go/commit/b9d5ab611f171528ba5b8609aea45e37fbcc5932))
* **api:** manual updates ([8b25a82](https://github.com/team-telnyx/telnyx-go/commit/8b25a82353a61633ebe84d5ee27334feb74a7961))
* **api:** manual updates ([fefd67b](https://github.com/team-telnyx/telnyx-go/commit/fefd67b109d2c5c012d8a47216b36759460bfdc6))
* **api:** manual updates ([1f95500](https://github.com/team-telnyx/telnyx-go/commit/1f9550001206d12ca2c9c5b4ca58bb0e3b6e5753))
* **api:** manual updates ([e625099](https://github.com/team-telnyx/telnyx-go/commit/e625099f707f14967c2bea84ed6d282c26e0fea4))
* **api:** manual updates ([e0b0f1b](https://github.com/team-telnyx/telnyx-go/commit/e0b0f1b8e4ec762ae1bcb350979af96b5568650a))
* **api:** manual updates ([cd3b9e3](https://github.com/team-telnyx/telnyx-go/commit/cd3b9e39868c81a76e902f1364b9823f9e9e4c34))
* **api:** manual updates ([c79e174](https://github.com/team-telnyx/telnyx-go/commit/c79e1748f6ae09c7bc187221b32b320ff2357aaa))
* fix-stainless-sdk-timeout ([0539121](https://github.com/team-telnyx/telnyx-go/commit/0539121c5bd1607d05452571688d31b20fe01c0a))
* TELAPPS Add ApplicationSid param ([102472f](https://github.com/team-telnyx/telnyx-go/commit/102472f5311ae02006965c37019b54fb25928908))
* TELAPPS Add interim_results to deepgram config ([56a7d7d](https://github.com/team-telnyx/telnyx-go/commit/56a7d7dda5b1e1e8eafabd57f1ed9032f9757e6d))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([09da6eb](https://github.com/team-telnyx/telnyx-go/commit/09da6eb09775e6ae1ddf9f581a2ee808b673cda1))
* **client:** use correct format specifier for header serialization ([49fed3a](https://github.com/team-telnyx/telnyx-go/commit/49fed3abfb9918d7bf3dd3ec7e23d6690cd57196))
* **internal:** skip tests that depend on mock server ([23c9bec](https://github.com/team-telnyx/telnyx-go/commit/23c9bec1d47e3950698904c6abd53d5010a57c46))
* move unsupported string formats to x-format ([e013fcc](https://github.com/team-telnyx/telnyx-go/commit/e013fccb33d3a96e6e489aa700d5228df172ced3))
* OAS drift — 10dlc.json (messaging-campaign-registry) ([fc6998b](https://github.com/team-telnyx/telnyx-go/commit/fc6998b7869607d134cf91903d60018528b2c3a1))
* OAS drift — messaging.json (messaging-settings + messaging-outbound) ([1fc474c](https://github.com/team-telnyx/telnyx-go/commit/1fc474cae78973892d96ec6e81b832b712391e0f))
* OAS drift — toll-free-verification.json (messaging-tf-verify) ([7b9e913](https://github.com/team-telnyx/telnyx-go/commit/7b9e91365cd6210db6c39f9ecedb56b3613cc67b))
* OAS drift — verify.json (messaging-2fa) ([bcbc092](https://github.com/team-telnyx/telnyx-go/commit/bcbc09281e2d908ce311b8ff200a46c1d9d216f2))
* StringFormatNotSupported ([35ad893](https://github.com/team-telnyx/telnyx-go/commit/35ad893f7e8865c595a3bde6f402894dde410fc9))


### Chores

* **internal:** codegen related update ([01b9ca6](https://github.com/team-telnyx/telnyx-go/commit/01b9ca615baa27063a1c245c087a5fe4c8615228))
* **internal:** codegen related update ([3c657ff](https://github.com/team-telnyx/telnyx-go/commit/3c657ff23a59786d0ee198da41726bf2ac48334d))
* **internal:** remove mock server code ([a087122](https://github.com/team-telnyx/telnyx-go/commit/a087122515a694d6480bcac23842ba6d67629f70))
* update mock server docs ([d036d12](https://github.com/team-telnyx/telnyx-go/commit/d036d1256c436129c404cf9b260a603daafb84b5))


### Documentation

* add service_provider_login_url to authentication provider settings ([3886cab](https://github.com/team-telnyx/telnyx-go/commit/3886cabb888880a690f4d43be249b3637d231787))
* **call-control:** Add missing conference endpoints ([d3f4cbb](https://github.com/team-telnyx/telnyx-go/commit/d3f4cbb0d5ae20b79d7f724fd0e257eb2ee339e4))
* **call-control:** Add missing parameters to call control endpoints ([622b420](https://github.com/team-telnyx/telnyx-go/commit/622b4209934b5903e5964f3353dd0189c4bb65d5))
* **call-scripting:** add Timeout and TimeLimit to InitiateTexmlCall ([293b294](https://github.com/team-telnyx/telnyx-go/commit/293b29447e7b9d9384132777ab9938f56b9428c3))

## 4.29.0 (2026-02-18)

Full Changelog: [v4.28.0...v4.29.0](https://github.com/team-telnyx/telnyx-go/compare/v4.28.0...v4.29.0)

### Features

* Add smart encoding fields to messaging API spec ([109b506](https://github.com/team-telnyx/telnyx-go/commit/109b5064f3d793b13daa1a0c321c44078f5342d9))
* **api:** manual updates ([3a79955](https://github.com/team-telnyx/telnyx-go/commit/3a799553c1d33bb3f7e426d29ffc8b81b13735b7))
* **api:** manual updates ([a5a42ff](https://github.com/team-telnyx/telnyx-go/commit/a5a42ffc85a8e41462ff83f005ae22c890310aa1))
* **api:** manual updates ([2893c31](https://github.com/team-telnyx/telnyx-go/commit/2893c3117ead11141e37121ff7eea8425718a6b9))
* **api:** manual updates ([99ede75](https://github.com/team-telnyx/telnyx-go/commit/99ede75268ac63f6a956fff647823b402ee23299))


### Bug Fixes

* add discriminator to TranscriptionEngineDeepgramConfig for Stainless SDK ([ec0f1c5](https://github.com/team-telnyx/telnyx-go/commit/ec0f1c5e632fe44478f75f795025980c10fd40d7))

## 4.28.0 (2026-02-18)

Full Changelog: [v4.27.0...v4.28.0](https://github.com/team-telnyx/telnyx-go/compare/v4.27.0...v4.28.0)

### Features

* Align transfer tool AMD spec with portal: premium mode, drop continue actions ([a5ed96c](https://github.com/team-telnyx/telnyx-go/commit/a5ed96c6e91a004b31ef5774a3d6b795bebc5c23))
* **api:** manual updates ([a1daef1](https://github.com/team-telnyx/telnyx-go/commit/a1daef1cb2675105a83e102963caa4fe71de993e))
* **api:** manual updates ([6f07e90](https://github.com/team-telnyx/telnyx-go/commit/6f07e907c4f6320f00a5077a1b1128518fd721cc))
* **api:** manual updates ([474b100](https://github.com/team-telnyx/telnyx-go/commit/474b10095b0e1dd9d7f0e56b31dbd728d4cb4945))

## 4.27.0 (2026-02-17)

Full Changelog: [v4.26.0...v4.27.0](https://github.com/team-telnyx/telnyx-go/compare/v4.26.0...v4.27.0)

### Features

* Add Minimax and Resemble voice providers for speak commands ([0a2e598](https://github.com/team-telnyx/telnyx-go/commit/0a2e59879fe735297a882308a9d4bc1a0aac02c5))

## 4.26.0 (2026-02-13)

Full Changelog: [v4.25.0...v4.26.0](https://github.com/team-telnyx/telnyx-go/compare/v4.25.0...v4.26.0)

### Features

* Add Label parameter to Dial Conference Participant endpoint ([dbb5e1a](https://github.com/team-telnyx/telnyx-go/commit/dbb5e1ad4955b4fe9b9843eb5806fa4330da166e))
* AI-2090: Add skip_turn tool type to assistant config schema ([466e3d6](https://github.com/team-telnyx/telnyx-go/commit/466e3d68b14f56d4370d6dbe8c1f00907ff752f1))

## 4.25.0 (2026-02-13)

Full Changelog: [v4.24.0...v4.25.0](https://github.com/team-telnyx/telnyx-go/compare/v4.24.0...v4.25.0)

### Features

* Add dynamic_variables field to scheduled event schemas ([d79464a](https://github.com/team-telnyx/telnyx-go/commit/d79464ae36cdcca5eaf0686b7d36d81f91b06db9))
* Add OpenAI-compatible embeddings API spec ([77b88e4](https://github.com/team-telnyx/telnyx-go/commit/77b88e4b8e1550086aafd8c66e2ce6933f57f49a))
* AI-2086: Add AI Missions endpoints to inference spec ([103af31](https://github.com/team-telnyx/telnyx-go/commit/103af3120228b9e43d04be8c82e3f33b4f92fa03))
* **api:** manual updates ([2f8fb9c](https://github.com/team-telnyx/telnyx-go/commit/2f8fb9c8a445c6f45d7a8bb22c15b7621008ce3e))
* **api:** manual updates ([fdb4a04](https://github.com/team-telnyx/telnyx-go/commit/fdb4a047e638f8bd45a9aee7769b64230a2ec4e3))
* **api:** manual updates to include models ([64e8daa](https://github.com/team-telnyx/telnyx-go/commit/64e8daa4702121946129b908fceb0b04bb3180e8))
* ENGDESK-49554: Add quail_voice_focus to noise suppression engine enums ([fb1136a](https://github.com/team-telnyx/telnyx-go/commit/fb1136a6691b498129c57dae06d7c27a9cdd735d))


### Bug Fixes

* **client:** revert change to certain pagination metadata types ([2512453](https://github.com/team-telnyx/telnyx-go/commit/25124530e6c4e05e01c9e77e4d2483e4d67cf6d6))

## 4.24.0 (2026-02-11)

Full Changelog: [v4.23.0...v4.24.0](https://github.com/team-telnyx/telnyx-go/compare/v4.23.0...v4.24.0)

### Features

* fix schema in emergency address schema ([ef38477](https://github.com/team-telnyx/telnyx-go/commit/ef384775f54c1fc95189928cc1da3406aecd9d26))


### Bug Fixes

* remove invalid discriminators from string enum schemas ([6810353](https://github.com/team-telnyx/telnyx-go/commit/6810353e6ca06386d7424531bf81c42e58755088))

## 4.23.0 (2026-02-11)

Full Changelog: [v4.22.0...v4.23.0](https://github.com/team-telnyx/telnyx-go/compare/v4.22.0...v4.23.0)

### Features

* Limit detection_mode enum to disabled and detect only ([0909bc5](https://github.com/team-telnyx/telnyx-go/commit/0909bc56e56aa2b71201c987db4cf31da15aee47))


### Bug Fixes

* **encoder:** correctly serialize NullStruct ([d4748ab](https://github.com/team-telnyx/telnyx-go/commit/d4748ab07f4d7fca78d3ab75319e774cc108b8dd))

## 4.22.0 (2026-02-09)

Full Changelog: [v4.21.0...v4.22.0](https://github.com/team-telnyx/telnyx-go/compare/v4.21.0...v4.22.0)

### Features

* AI-2078 Feature: API endpoint to disable AI assistant mid-conversation ([472ce6b](https://github.com/team-telnyx/telnyx-go/commit/472ce6b0d45a47dcc39398ef7337d5bda65151d5))

## 4.21.0 (2026-02-09)

Full Changelog: [v4.20.0...v4.21.0](https://github.com/team-telnyx/telnyx-go/compare/v4.20.0...v4.21.0)

### Features

* add Telnyx webhook verification with ED25519 signatures ([f634059](https://github.com/team-telnyx/telnyx-go/commit/f6340596c3a0867f4537fa3737671fa48cdce281))

## 4.20.0 (2026-02-06)

Full Changelog: [v4.19.0...v4.20.0](https://github.com/team-telnyx/telnyx-go/compare/v4.19.0...v4.20.0)

### Features

* Revert "fix emergency settings -schema" ([d9e25f2](https://github.com/team-telnyx/telnyx-go/commit/d9e25f298581f5462b1a2b2979eba98c1c5b2ca2))

## 4.19.0 (2026-02-05)

Full Changelog: [v4.18.0...v4.19.0](https://github.com/team-telnyx/telnyx-go/compare/v4.18.0...v4.19.0)

### Features

* **api:** manual updates ([f1dd32d](https://github.com/team-telnyx/telnyx-go/commit/f1dd32dd6f7048ceeaca99f3682a4b12416f0be2))
* **api:** Merge pull request [#23](https://github.com/team-telnyx/telnyx-go/issues/23) from stainless-sdks/fix/deepgram-nova3-enum-duplicates ([57c96d0](https://github.com/team-telnyx/telnyx-go/commit/57c96d0f6177157f78d51d987b456742721b7ffe))

## 4.18.0 (2026-02-05)

Full Changelog: [v4.17.0...v4.18.0](https://github.com/team-telnyx/telnyx-go/compare/v4.17.0...v4.18.0)

### Features

* **api:** Merge pull request [#22](https://github.com/team-telnyx/telnyx-go/issues/22) from stainless-sdks/add-all-webhook-models ([1bf76af](https://github.com/team-telnyx/telnyx-go/commit/1bf76af3b74184ac3f074272f790aecaa1f9cf93))

## 4.17.0 (2026-02-04)

Full Changelog: [v4.16.1...v4.17.0](https://github.com/team-telnyx/telnyx-go/compare/v4.16.1...v4.17.0)

### Features

* Add Texml parameter to create call endpoint [ENGDESK-49187] ([f987ad8](https://github.com/team-telnyx/telnyx-go/commit/f987ad87fe9b6dde8dbdf3d39d8022d7ce733fcf))

## 4.16.1 (2026-02-02)

Full Changelog: [v4.16.0...v4.16.1](https://github.com/team-telnyx/telnyx-go/compare/v4.16.0...v4.16.1)

### Bug Fixes

* **client/oauth:** send grant_type in the right location ([59c366b](https://github.com/team-telnyx/telnyx-go/commit/59c366b03572001cece12ea751af3a90bc72fede))

## 4.16.0 (2026-01-30)

Full Changelog: [v4.15.1...v4.16.0](https://github.com/team-telnyx/telnyx-go/compare/v4.15.1...v4.16.0)

### Features

* **api:** manual updates ([5479ea0](https://github.com/team-telnyx/telnyx-go/commit/5479ea06526da84eaef99f9e6944664963894409))

## 4.15.1 (2026-01-30)

Full Changelog: [v4.15.0...v4.15.1](https://github.com/team-telnyx/telnyx-go/compare/v4.15.0...v4.15.1)

### Bug Fixes

* remove deprecated TeXML API endpoints from OpenAPI spec ([91ddee4](https://github.com/team-telnyx/telnyx-go/commit/91ddee4fbccdd9fefa6eba550ac6d306e91ac416))
* use PaginationMeta schema for ListFaxesResponse.meta ([3651a7e](https://github.com/team-telnyx/telnyx-go/commit/3651a7e578ef69d772c1d5952528df33e275dcb8))

## 4.15.0 (2026-01-29)

Full Changelog: [v4.14.0...v4.15.0](https://github.com/team-telnyx/telnyx-go/compare/v4.14.0...v4.15.0)

### Features

* **api:** revert bad update to spec ([5155d01](https://github.com/team-telnyx/telnyx-go/commit/5155d01033d71ff661aa25dadc3a66fc79d40780))

## 4.14.0 (2026-01-29)

Full Changelog: [v4.13.0...v4.14.0](https://github.com/team-telnyx/telnyx-go/compare/v4.13.0...v4.14.0)

### Features

* Add deepgram/nova-3 transcription engine option to record_start ([51916b7](https://github.com/team-telnyx/telnyx-go/commit/51916b7248951efa08b945ff171c3558b2434467))

## 4.13.0 (2026-01-28)

Full Changelog: [v4.12.1...v4.13.0](https://github.com/team-telnyx/telnyx-go/compare/v4.12.1...v4.13.0)

### Features

* Fix broken documentation links ([d16d7ec](https://github.com/team-telnyx/telnyx-go/commit/d16d7ec304175ae29dc7c78b4ad40ad941612c38))

## 4.12.1 (2026-01-28)

Full Changelog: [v4.12.0...v4.12.1](https://github.com/team-telnyx/telnyx-go/compare/v4.12.0...v4.12.1)

### Bug Fixes

* **docs:** fix mcp installation instructions for remote servers ([38cf582](https://github.com/team-telnyx/telnyx-go/commit/38cf5826fdc9360343f8d85c7c8b2f7d6ec2224d))

## 4.12.0 (2026-01-27)

Full Changelog: [v4.11.0...v4.12.0](https://github.com/team-telnyx/telnyx-go/compare/v4.11.0...v4.12.0)

### Features

* Deploy dev/mc vady wip ([a10129d](https://github.com/team-telnyx/telnyx-go/commit/a10129db5b9e7f7a67d4399fa57e59fdef49d2c2))
* jira-engdesk-49089 add new connection jitter buffer related fields ([8601fcf](https://github.com/team-telnyx/telnyx-go/commit/8601fcf8da3f19a99061f9ba791a1917a62ba6c3))

## 4.11.0 (2026-01-23)

Full Changelog: [v4.10.0...v4.11.0](https://github.com/team-telnyx/telnyx-go/compare/v4.10.0...v4.11.0)

### Features

* **api:** manual updates ([18d292c](https://github.com/team-telnyx/telnyx-go/commit/18d292c4e30d72a7f1e15ee048821cef135f8b46))
* **client:** add a convenient param.SetJSON helper ([72ee0e5](https://github.com/team-telnyx/telnyx-go/commit/72ee0e5698455866f70ef66fa84fccd99f70d990))

## 4.10.0 (2026-01-22)

Full Changelog: [v4.9.0...v4.10.0](https://github.com/team-telnyx/telnyx-go/compare/v4.9.0...v4.10.0)

### Features

* Deploy dev/mc vady wip ([3f2d6a3](https://github.com/team-telnyx/telnyx-go/commit/3f2d6a3df9a4d09693b47eff0ec5da0e6ee1bcee))

## 4.9.0 (2026-01-21)

Full Changelog: [v4.8.0...v4.9.0](https://github.com/team-telnyx/telnyx-go/compare/v4.8.0...v4.9.0)

### Features

* fix-redocly-lint-issues ([45d63d8](https://github.com/team-telnyx/telnyx-go/commit/45d63d8ac47fd3830ae09e2e7f32747d1e38e0a1))

## 4.8.0 (2026-01-20)

Full Changelog: [v4.7.0...v4.8.0](https://github.com/team-telnyx/telnyx-go/compare/v4.7.0...v4.8.0)

### Features

* Add Post /v2/calls/:call_control_id/actions/ai_assistant_add_messages ([2f0be16](https://github.com/team-telnyx/telnyx-go/commit/2f0be16ab18549b188b2a4698dc2aad0ccefd7fc))

## 4.7.0 (2026-01-20)

Full Changelog: [v4.6.0...v4.7.0](https://github.com/team-telnyx/telnyx-go/compare/v4.6.0...v4.7.0)

### Features

* Update voicemail_detection description with AMD enablement info ([b68557b](https://github.com/team-telnyx/telnyx-go/commit/b68557bc78d56a7fb6b4057c74eac4bf62d374f6))


### Bug Fixes

* correct broken link to List SIM Card Actions endpoint in SIM car… ([6927bbd](https://github.com/team-telnyx/telnyx-go/commit/6927bbd57669a6e1d13a987baa3acbf39fd77a5f))

## 4.6.0 (2026-01-19)

Full Changelog: [v4.5.1...v4.6.0](https://github.com/team-telnyx/telnyx-go/compare/v4.5.1...v4.6.0)

### Features

* Add AI Assistant spec updates for FE tickets ([956d588](https://github.com/team-telnyx/telnyx-go/commit/956d588e576efdf85e5769f45c18d7c88ae16f1c))
* fix-external-connection-link ([7d5af93](https://github.com/team-telnyx/telnyx-go/commit/7d5af9395216e821bdcc325218c2f61d7ac6f5fc))

## 4.5.1 (2026-01-16)

Full Changelog: [v4.5.0...v4.5.1](https://github.com/team-telnyx/telnyx-go/compare/v4.5.0...v4.5.1)

### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([13fd1fc](https://github.com/team-telnyx/telnyx-go/commit/13fd1fc13f04b602c427a2415753d21a3b2d9cb5))

## 4.5.0 (2026-01-16)

Full Changelog: [v4.4.0...v4.5.0](https://github.com/team-telnyx/telnyx-go/compare/v4.4.0...v4.5.0)

### Features

* TELAPPS-5507: Add Krisp engine description for noise suppression ([bfe119a](https://github.com/team-telnyx/telnyx-go/commit/bfe119abb2ee2d660666ddd5659cd6122161f247))


### Bug Fixes

* update broken MDR report link in GetMessage endpoint ([94302bd](https://github.com/team-telnyx/telnyx-go/commit/94302bd7a7bb1e1a2a694e01b19e693aa4f68422))

## 4.4.0 (2026-01-16)

Full Changelog: [v4.3.0...v4.4.0](https://github.com/team-telnyx/telnyx-go/compare/v4.3.0...v4.4.0)

### Features

* fix links ([1895c7d](https://github.com/team-telnyx/telnyx-go/commit/1895c7dd27347dea734cc9e310698a99f4afdd72))


### Chores

* **internal:** update `actions/checkout` version ([45e94d7](https://github.com/team-telnyx/telnyx-go/commit/45e94d7f80a619535edd62d379281fc46dfa5689))

## 4.3.0 (2026-01-15)

Full Changelog: [v4.2.1...v4.3.0](https://github.com/team-telnyx/telnyx-go/compare/v4.2.1...v4.3.0)

### Features

* jira-engdesk-48800 add organizations-related docs to the external api… ([87dd437](https://github.com/team-telnyx/telnyx-go/commit/87dd437c7df7d63e6a13028793edf3069da2e576))
* MSG-6148: adding the new campaignVerifyAuthorizationToken field and missing GET OTP endpoint ([ff821b7](https://github.com/team-telnyx/telnyx-go/commit/ff821b75c902a9cfab7a80d9e4a6165cb70a3730))
* MSG-6228: MSG-6228: Add smart_encoding option for SMS character encoding optimization ([796ab2b](https://github.com/team-telnyx/telnyx-go/commit/796ab2b4d643f934336ebb9489b798659ae52e57))
* TELAPPS-ENGDESK-48790 Remove duplicated webhooks ([d3615f5](https://github.com/team-telnyx/telnyx-go/commit/d3615f5e2dcecbf959ccaa428968e0fa595c05ca))

## 4.2.1 (2026-01-14)

Full Changelog: [v4.2.0...v4.2.1](https://github.com/team-telnyx/telnyx-go/compare/v4.2.0...v4.2.1)

### Bug Fixes

* **client:** invalid URL ([4895700](https://github.com/team-telnyx/telnyx-go/commit/48957003021a0cb61f582c90630d7c16c46fd1ef))

## 4.2.0 (2026-01-14)

Full Changelog: [v4.1.0...v4.2.0](https://github.com/team-telnyx/telnyx-go/compare/v4.1.0...v4.2.0)

### Features

* Add widget_settings to AI Assistant and import_ids to ImportAssistant… ([1540c3b](https://github.com/team-telnyx/telnyx-go/commit/1540c3b35a65ac050a6c8b6c1c2e10a4b02344df))
* **api:** fix default pagination by correctly using nested params ([bf475c0](https://github.com/team-telnyx/telnyx-go/commit/bf475c02def137869f3226f845fa2c03b1fff9b2))

## 4.1.0 (2026-01-13)

Full Changelog: [v4.0.0...v4.1.0](https://github.com/team-telnyx/telnyx-go/compare/v4.0.0...v4.1.0)

### Features

* **api:** manual updates ([04a604f](https://github.com/team-telnyx/telnyx-go/commit/04a604f6eb9db097d3e7c19621817e81aa278f69))
* TELAPPS Add GET /texml/Accounts/{account_sid}/Queues endpoint ([30157cd](https://github.com/team-telnyx/telnyx-go/commit/30157cdc59e45769a71fb672a88bfdf1211b6def))
* TELAPPS-ENGDESK-47967 Add black_threshold parameter to send_fax request ([e975cbb](https://github.com/team-telnyx/telnyx-go/commit/e975cbb078ebde59775c24319f2751da16fdadd4))

## 4.0.0 (2026-01-09)

Full Changelog: [v3.7.0...v4.0.0](https://github.com/team-telnyx/telnyx-go/compare/v3.7.0...v4.0.0)

### ⚠ BREAKING CHANGES

* Resolved all codegen errors

### Features

* (draft/don't review) ENGDESK-38070-c: add deepgram keyword documentation ([8811cea](https://github.com/team-telnyx/telnyx-go/commit/8811cead0b05702304355e21cd68f72d5a12ed2d))
* [PORT-4538] Fix ambiguous oneOf instances on porting service and documents ([087050e](https://github.com/team-telnyx/telnyx-go/commit/087050ee14af17a68273c4de587fb8b80405ba0f))
* Add AI assistant voice settings, telephony config, and tools updates ([a8a10ac](https://github.com/team-telnyx/telnyx-go/commit/a8a10acbff7d54788d21bdd5f103e2157cbaca75))
* Add response schemas for telco data usage report endpoints ([7892672](https://github.com/team-telnyx/telnyx-go/commit/78926727f02fa66f2eb6c829e59898a52627c522))
* alright, shut up redocly ([5bbf314](https://github.com/team-telnyx/telnyx-go/commit/5bbf3144a502222199a8d94c8f4247c05a84dbf6))
* **api:** join all 10dlc operations into messaging_10dlc group ([91deaaf](https://github.com/team-telnyx/telnyx-go/commit/91deaaf0f9889b4d0a1abe1a5c4c1d128f0d1cc3))
* **api:** manual updates ([f9162bc](https://github.com/team-telnyx/telnyx-go/commit/f9162bc4f08c7277e9726da32032ecf6a1d2c274))
* **api:** manual updates ([3d42616](https://github.com/team-telnyx/telnyx-go/commit/3d426165e11afe5f2f1a8a9279f4c8143738b187))
* **api:** manual updates ([62492df](https://github.com/team-telnyx/telnyx-go/commit/62492df56c80ead1b6ebbcc512ad37b4cf4419fb))
* **api:** manual updates ([88e5593](https://github.com/team-telnyx/telnyx-go/commit/88e559375c32de01ddf57e7cc0a3ab014acfd2c8))
* **api:** manual updates ([38fcd33](https://github.com/team-telnyx/telnyx-go/commit/38fcd33cfea8f999ccbadcbfc32331b22cc07cfd))
* **api:** manual updates ([abcf0cb](https://github.com/team-telnyx/telnyx-go/commit/abcf0cbbdf4c205b998c1bc1dd4d0bf2e36ef512))
* **api:** manual updates ([8f2014a](https://github.com/team-telnyx/telnyx-go/commit/8f2014a8606f6f3b3c12aa75ad2991c9af1a7ea2))
* **api:** messaging_10dlc group with all their endpoints ([c6ab42a](https://github.com/team-telnyx/telnyx-go/commit/c6ab42a9392f41fe385a897fd02fc66aed461098))
* **api:** PHP codegen error fixes ([3926262](https://github.com/team-telnyx/telnyx-go/commit/392626257db3a853bf4b3feacef9e1361f10010b))
* **api:** reverted previous commit ([538f71b](https://github.com/team-telnyx/telnyx-go/commit/538f71b270f7d6b3df41963272fc5c057652a8d4))
* Chat completions response schema update ([6a3357a](https://github.com/team-telnyx/telnyx-go/commit/6a3357a9e1b6d8505fb1a0966bf70aac7d5cac0d))
* **client:** add separate models for 2 events ([5138e49](https://github.com/team-telnyx/telnyx-go/commit/5138e49cd0ce938ab3b16c9c93aca6a636ff98ed))
* Document supervising leg of call ([2b82677](https://github.com/team-telnyx/telnyx-go/commit/2b82677e5b583130102c9026765b6790cb49ac7a))
* DOTCOM-5145. Update redocly lint to block new lint errors or warning being introduced ([aee1f3d](https://github.com/team-telnyx/telnyx-go/commit/aee1f3d0c62426229b5a22f5d23cd34ec83ca14d))
* DOTCOM-5179. Fix Redocly errors in outbound-voice-profiles.json ([0d30374](https://github.com/team-telnyx/telnyx-go/commit/0d303740a5f6abd4a007bdc56dbc8152e519ea80))
* Draft. DOTCOM-5184. Fix 44 errors in the spec as reported by Redocly on video ([50d0826](https://github.com/team-telnyx/telnyx-go/commit/50d082614d78c58a1dd1f40b75d98e29ee7cb1a0))
* **encoder:** support bracket encoding form-data object members ([1162825](https://github.com/team-telnyx/telnyx-go/commit/1162825e37115d6db6a229df55909495f810a558))
* Engdesk 47920/wireless cleanup ([b4fc4eb](https://github.com/team-telnyx/telnyx-go/commit/b4fc4eb90976339daa46bc5702cb90fdab6131d7))
* ENGDESK-47580: Add quickship and exclude_held_numbers filters to inexplicit number order API ([21059b3](https://github.com/team-telnyx/telnyx-go/commit/21059b3fa97055d4e091ef6246faa5e0d9a40ac8))
* ENGDESK-47580: Add quickship and exclude_held_numbers to InexplicitNumberOrderResponse ([296dba9](https://github.com/team-telnyx/telnyx-go/commit/296dba9550d001d9bcc2a89e28991d5ae2e3100d))
* ENGDESK-47706: Update TranscriptionEngineDeepgramConfig Schema ([197e351](https://github.com/team-telnyx/telnyx-go/commit/197e351e687400a06c1dd96553631c1e9704bcd5))
* ENGDESK-47736: added discriminator fields to oneOffs that were missing them ([04d2142](https://github.com/team-telnyx/telnyx-go/commit/04d21427ce30eaaf62145d8921e7f478575875e8))
* ENGDESK-47759 - fix missing meta definition in authorized ips spec ([dab1e10](https://github.com/team-telnyx/telnyx-go/commit/dab1e10405fa3540f4209760191591ff13954c01))
* ENGDESK-47883: Fix all lint errors in telapps owned APIs ([e6947ed](https://github.com/team-telnyx/telnyx-go/commit/e6947ed75631aa8f0dfa8e2b356537ca99930a9e))
* ENGDESK-47886: Fix API spec for emergency.json ([97ab941](https://github.com/team-telnyx/telnyx-go/commit/97ab941d3b98516bb7554737e95e8817f7466bb8))
* ENGDESK-47914 - fix warnings in numbers.json file ([2b6c3ae](https://github.com/team-telnyx/telnyx-go/commit/2b6c3aeb9f47f2ee1b4b3cf673d4e3cae68f12b6))
* ENGDESK-47947 - fix wrong type on user-addresses request object ([0ea9e53](https://github.com/team-telnyx/telnyx-go/commit/0ea9e53e7d8b9797087405d93b5d8ca950aa28d4))
* ENGDESK-48016 - document simultaneous ringing for CredentialConnections ([565870e](https://github.com/team-telnyx/telnyx-go/commit/565870e939212e535c08317e1de38f4af66f034c))
* ENGDESK-48254: Release noise suppression details docs to prod ([cd230f8](https://github.com/team-telnyx/telnyx-go/commit/cd230f8c70ebb2e2932aab2526fbb5bb79656309))
* FILE-1066: presigned url doc strings ([b6ac211](https://github.com/team-telnyx/telnyx-go/commit/b6ac2118b5286e6a929618de2ad50012bf6ecabd))
* Fix campaign usecase endpoint: /registry/enum/usecase → /10dlc/enum/usecase ([badbde6](https://github.com/team-telnyx/telnyx-go/commit/badbde6a27385994d22574f1934dc14b022af9f0))
* hotfix: restore 10dlc prefixes ([f92c64e](https://github.com/team-telnyx/telnyx-go/commit/f92c64e73eec6f3c4d42a0a4636702a2e701c92d))
* Improve messaging API naming and navigation ([715899f](https://github.com/team-telnyx/telnyx-go/commit/715899f33506bc25ee5e0b8980f3989086230fd4))
* messaging meta object with required fields ([f774d2a](https://github.com/team-telnyx/telnyx-go/commit/f774d2a5397f995c7ee78f894c827429e65d4d3c))
* Msg 6152 ([86fa49b](https://github.com/team-telnyx/telnyx-go/commit/86fa49bc3bb0188ecb8c7170eba3dce5b076674b))
* MSG-6140: Add SMS OTP endpoints for Sole Prop brands ([3cb35a2](https://github.com/team-telnyx/telnyx-go/commit/3cb35a2a627ba73c41710cdd41e54394834b2b33))
* MSG-6160 fix messaging lint issues ([28d72ee](https://github.com/team-telnyx/telnyx-go/commit/28d72ee5659feb1245b9188049f329ae478e0284))
* MSG-6166 fix empty schema responses ([02b984d](https://github.com/team-telnyx/telnyx-go/commit/02b984da6a55e1b9f0d6d4d3e1663818a06ac0ac))
* MSG-6179: Add discriminator fields to Messaging API schemas for improved SDK performance ([d86ef3e](https://github.com/team-telnyx/telnyx-go/commit/d86ef3eff863923af804013c028b8f7593c1a7b8))
* MSG-6181: Reorganize mobile phone number messaging endpoints and fix … ([9d625d9](https://github.com/team-telnyx/telnyx-go/commit/9d625d9750dca87d871d6001c53bc947e8726000))
* NETAPPS_687: Fixed IGW spec to match current API. ([2513c87](https://github.com/team-telnyx/telnyx-go/commit/2513c875c1151c0494c6dc191168358f1f2497b2))
* NUM-6334/NUM-6335 - fix redocly lint errors ([48963c0](https://github.com/team-telnyx/telnyx-go/commit/48963c0008852ea0b23301adf192763c248602fb))
* PORT-4528: Fix lint errors for porting ([f69db21](https://github.com/team-telnyx/telnyx-go/commit/f69db2187926939cb9a50150cba664aa8352cb29))
* port-4551: remove CustomerServiceRecordStatusChanged webhook doc ([c8033c4](https://github.com/team-telnyx/telnyx-go/commit/c8033c44969f00dcc1722db506afee1d8a20f4d1))
* PORT-4553: Add a discriminator to portout webhook ([96656a6](https://github.com/team-telnyx/telnyx-go/commit/96656a6bb9be80badcc51448d82e896be1131d3f))
* PORTAL-5787 - document query parameter to handle messaging service error ([aed95f1](https://github.com/team-telnyx/telnyx-go/commit/aed95f1ac03fea86e0b11566af148ac3478f6916))
* TBS-3422: Fix redocly errors ([b97a5f5](https://github.com/team-telnyx/telnyx-go/commit/b97a5f562ebec99566e436c0823df1b17c893f95))
* TBS-3422: Fix TBS redocly errors ([8192708](https://github.com/team-telnyx/telnyx-go/commit/8192708076140bdcdd2aed7032c7bf78be2fd61c))
* TELAPPS-47889 Add texml queue endpoint ([70c4c05](https://github.com/team-telnyx/telnyx-go/commit/70c4c05292f293aa7ba0bbcbe9513c6cd5133d69))
* TELAPPS-5428 Add speech-to-text WS endpoint ([04d19c3](https://github.com/team-telnyx/telnyx-go/commit/04d19c37320be4ee0f174b8b68ef643baa852873))
* Update default GATHER_USING_AI_MODEL ([6312938](https://github.com/team-telnyx/telnyx-go/commit/6312938cba6bdb19636cbe88150705dc3028b952))


### Bug Fixes

* **api:** 10dlc prefixes ([6a9ee58](https://github.com/team-telnyx/telnyx-go/commit/6a9ee584ac02c332fce6ed0059ce707198009bed))
* **client:** move MessagingError to shared and fix time in tests ([4776336](https://github.com/team-telnyx/telnyx-go/commit/47763367212c1bf51d271cb2897cde15dbf1a418))
* **client:** properly marshal embedded structs ([1dd2a57](https://github.com/team-telnyx/telnyx-go/commit/1dd2a57c40dc4cc209d56e4a9a19a72ea02e52b8))
* correct broken hyperlinks in Submit Campaign endpoint description ([075df52](https://github.com/team-telnyx/telnyx-go/commit/075df526d83d524362d14e88ed8f4c9f2de5c70a))
* make text field optional in AssistantSmsChatReq schema ([7bf540e](https://github.com/team-telnyx/telnyx-go/commit/7bf540eb8901f1eea2609640e641096d10298d64))
* **mcp:** correct code tool API endpoint ([3b1e98a](https://github.com/team-telnyx/telnyx-go/commit/3b1e98a040eab4a2c0ef2149e43703d8be39334f))
* rename param to avoid collision ([63d16d5](https://github.com/team-telnyx/telnyx-go/commit/63d16d5a5f678333a700071f24814088dfa4e51a))
* skip usage tests that don't work with Prism ([d2804fa](https://github.com/team-telnyx/telnyx-go/commit/d2804fa27b7151b0df0bb6eb89fdf65b2a570a1b))
* **stainless:** fixes the messsages typo ([a3d46f1](https://github.com/team-telnyx/telnyx-go/commit/a3d46f1b4ec076df7b515f3772abd4837ec49383))


### Chores

* add float64 to valid types for RegisterFieldValidator ([f3e3820](https://github.com/team-telnyx/telnyx-go/commit/f3e3820291f0d792fb7f1fad975e3fd1d3e10f3f))
* elide duplicate aliases ([46adbee](https://github.com/team-telnyx/telnyx-go/commit/46adbeed464fd18d126ce490a5eabea5ef8fb5e7))
* **internal:** codegen related update ([88cd4e6](https://github.com/team-telnyx/telnyx-go/commit/88cd4e6a0a0a78e85c8a4b5d83a3f995578a8459))
* **internal:** codegen related update ([598345a](https://github.com/team-telnyx/telnyx-go/commit/598345a2e934d4fc4018b7ea3f0a25235cac7c39))
* Resolved all codegen errors ([606037d](https://github.com/team-telnyx/telnyx-go/commit/606037dcfb8ec180eb02bccfc56caaf1e6d25498))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([96db56a](https://github.com/team-telnyx/telnyx-go/commit/96db56a36dcb30b26722a9821854fe2c774f3eab))

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

# Changelog

## [4.84.0](https://github.com/team-telnyx/telnyx-go/compare/v4.84.0...v4.84.0) (2026-07-02)


### release

* 4.84.0 ([#131](https://github.com/team-telnyx/telnyx-go/issues/131)) ([cf610cf](https://github.com/team-telnyx/telnyx-go/commit/cf610cf13c90004586e2c33935150d14e8e26ee4))


### Features

* [TDA-6425] Add Session Analysis API spec to public docs ([19c208e](https://github.com/team-telnyx/telnyx-go/commit/19c208ec8a07642c4536c21927d7cfd872027389))
* Add AI Assistant spec updates for FE tickets ([c43f5c7](https://github.com/team-telnyx/telnyx-go/commit/c43f5c7075ba1f7ecd3e409f174bc297204ce4a7))
* add ai_assistant_join call control command OpenAPI spec ([f0cecc4](https://github.com/team-telnyx/telnyx-go/commit/f0cecc407c836a9de9d1d12caf1cc430a05f6af8))
* Add deepgram/nova-3 transcription engine option to record_start ([7edfc1f](https://github.com/team-telnyx/telnyx-go/commit/7edfc1f49829fc61cec5b3be21c0ec4f644b009e))
* Add dynamic_variables field to scheduled event schemas ([433201e](https://github.com/team-telnyx/telnyx-go/commit/433201e5320446c2d3d3411aa80a7f82e18ad197))
* Add enable_thinking parameter to chat completions API ([eff37c5](https://github.com/team-telnyx/telnyx-go/commit/eff37c531dddec8c18cf919fc1f437c86b7826df))
* Add Label parameter to Dial Conference Participant endpoint ([597a2df](https://github.com/team-telnyx/telnyx-go/commit/597a2df071d84bd443bef818ae1fc9f5fce4d302))
* add message_history, send_message_history_updates, participants to AIAssistantStartRequest ([8c5299b](https://github.com/team-telnyx/telnyx-go/commit/8c5299b8b810e465fe17e8b20aed76a36a291aa4))
* Add Minimax and Resemble voice providers for speak commands ([ffa1019](https://github.com/team-telnyx/telnyx-go/commit/ffa1019e8a2f80402d0e4a1dea8db166d73aebf1))
* Add Minimax provider support to Voice Designs and Voice Clones API spec ([f3e0f6e](https://github.com/team-telnyx/telnyx-go/commit/f3e0f6e2c65745bd38ab7d370148c2195de1e993))
* Add missing TTS voice settings schemas and update voice descriptions ([7d9f470](https://github.com/team-telnyx/telnyx-go/commit/7d9f47029b60729442692f492f695f4657c723a2))
* Add OpenAI-compatible embeddings API spec ([0aea039](https://github.com/team-telnyx/telnyx-go/commit/0aea039674d989a0afaa91cdadb492175796c0c1))
* Add Post /v2/calls/:call_control_id/actions/ai_assistant_add_messages ([b198b1f](https://github.com/team-telnyx/telnyx-go/commit/b198b1f262d4a3bebef2b05d55ce0226af0446b5))
* add public x402 payment endpoints to external specs ([2dd0be8](https://github.com/team-telnyx/telnyx-go/commit/2dd0be82a2fcc58973375df4abfe6cdfc97d9596))
* Add smart encoding fields to messaging API spec ([109b506](https://github.com/team-telnyx/telnyx-go/commit/109b5064f3d793b13daa1a0c321c44078f5342d9))
* add Telnyx webhook verification with ED25519 signatures ([45afa66](https://github.com/team-telnyx/telnyx-go/commit/45afa6639874edaffe6b53f9c69760519de026fd))
* add Telnyx webhook verification with ED25519 signatures ([f634059](https://github.com/team-telnyx/telnyx-go/commit/f6340596c3a0867f4537fa3737671fa48cdce281))
* Add Texml parameter to create call endpoint [ENGDESK-49187] ([6a975ea](https://github.com/team-telnyx/telnyx-go/commit/6a975ead69cb9f693d247d474dd606bc37e3b0aa))
* Add text-to-speech WebSocket streaming OpenAPI spec ([44aa7ec](https://github.com/team-telnyx/telnyx-go/commit/44aa7ec628935f4cfe23e7c7115d6886011e98a5))
* Add TTS file-based endpoint spec ([7f39dbe](https://github.com/team-telnyx/telnyx-go/commit/7f39dbeae28f5f386d4c1d552e553c2c15f4e530))
* Add Voice Designs and Voice Clones API specification ([ba5f82c](https://github.com/team-telnyx/telnyx-go/commit/ba5f82caacbeb020c9b2e0533ac6405157c33e3c))
* Add widget_settings to AI Assistant and import_ids to ImportAssistant… ([5b7e4d2](https://github.com/team-telnyx/telnyx-go/commit/5b7e4d2083fb04147539ba36d7ef5a6ca5a46742))
* Ai 1967 ([4fdc338](https://github.com/team-telnyx/telnyx-go/commit/4fdc3389c3eb84a5384651207931d8e076948547))
* Ai 1967 part 2 ([0a4820e](https://github.com/team-telnyx/telnyx-go/commit/0a4820eeebd71b3dcbd29b91bf376eacca062e46))
* AI-1842: Add MCP Servers and Integrations sections ([d3c4c7f](https://github.com/team-telnyx/telnyx-go/commit/d3c4c7f09c835f8274bdff0b0b8256804f6a9ae1))
* AI-2078 Feature: API endpoint to disable AI assistant mid-conversation ([ec32417](https://github.com/team-telnyx/telnyx-go/commit/ec32417709a98ea0ffbdb09b01f0e5552a8dc1bc))
* AI-2086: Add AI Missions endpoints to inference spec ([7d711ef](https://github.com/team-telnyx/telnyx-go/commit/7d711ef8abc38b1ecba48f6c558a0611d3a27557))
* AI-2090: Add skip_turn tool type to assistant config schema ([f1cf799](https://github.com/team-telnyx/telnyx-go/commit/f1cf7998f94ca6d4854235ade569b97c425cb528))
* AI-2093: Add language_boost to MiniMax voice settings ([fcfef69](https://github.com/team-telnyx/telnyx-go/commit/fcfef6932be3d64e98e925d524709a035bdefaec))
* AI-2106: Add invite tool schema to inference OpenAPI spec ([dc022c5](https://github.com/team-telnyx/telnyx-go/commit/dc022c5b9b84e25f82fed040fd42eb7636474548))
* AI-2131: Add expressive_mode boolean to VoiceSettings ([fa8e60c](https://github.com/team-telnyx/telnyx-go/commit/fa8e60c2ecc65c7d856655dabae364d3aa9fa469))
* AI-2132: Add warm_message_delay_ms to transfer tool OpenAPI spec ([9be1828](https://github.com/team-telnyx/telnyx-go/commit/9be18283eb74c7c99b8166d089ebd5afe36c55e1))
* Align transfer tool AMD spec with portal: premium mode, drop continue actions ([14dc147](https://github.com/team-telnyx/telnyx-go/commit/14dc1474ba4ee78cbe3186f9094dda4a3d60f722))
* **api:** fix default pagination by correctly using nested params ([c8fb1cd](https://github.com/team-telnyx/telnyx-go/commit/c8fb1cdff56d5f3df6307217731f481c45afb48a))
* **api:** manual updates ([63ee662](https://github.com/team-telnyx/telnyx-go/commit/63ee66266ea278bf34b32158b4de5b4ed309520f))
* **api:** manual updates ([870d363](https://github.com/team-telnyx/telnyx-go/commit/870d363e4a63e09b10143e119336ace8c00089ec))
* **api:** manual updates ([5df6d63](https://github.com/team-telnyx/telnyx-go/commit/5df6d636dd4b6f4712489c27bca58f708f9faab2))
* **api:** manual updates ([ce49b95](https://github.com/team-telnyx/telnyx-go/commit/ce49b9567d4419735e3a194e3a8cd5c8840abcfa))
* **api:** manual updates ([4081249](https://github.com/team-telnyx/telnyx-go/commit/40812491036b0d47232a63bfbf286a22f470d598))
* **api:** manual updates ([38501b1](https://github.com/team-telnyx/telnyx-go/commit/38501b1b02a4de16dd5136e08171d946f85a5606))
* **api:** manual updates ([40bbd69](https://github.com/team-telnyx/telnyx-go/commit/40bbd69f60eedbc038931ae0af7bad940b5534d5))
* **api:** manual updates ([f410c66](https://github.com/team-telnyx/telnyx-go/commit/f410c661be1f1127b84e50f262dd77bf0adb8f73))
* **api:** manual updates ([0f82678](https://github.com/team-telnyx/telnyx-go/commit/0f8267856b844ef6013438833f41124ccca8a389))
* **api:** manual updates ([cb17191](https://github.com/team-telnyx/telnyx-go/commit/cb1719158bf561fd8873ee37b5751a206c9dd4de))
* **api:** manual updates ([a09a3dd](https://github.com/team-telnyx/telnyx-go/commit/a09a3dda033e78ed6900119ec1d655cd43b56433))
* **api:** manual updates ([6af49e0](https://github.com/team-telnyx/telnyx-go/commit/6af49e0b290c54e1df93de47b91b12c2e86cc83e))
* **api:** manual updates ([43d6e9c](https://github.com/team-telnyx/telnyx-go/commit/43d6e9c03f3292b31891e9a42ce3010db17b2ed3))
* **api:** manual updates ([da187c8](https://github.com/team-telnyx/telnyx-go/commit/da187c8ce0f9b7ea4cee5534f3f5bb1a84995730))
* **api:** manual updates ([96235d8](https://github.com/team-telnyx/telnyx-go/commit/96235d895fa228600ff142821d7307631ff46724))
* **api:** manual updates ([e4fa420](https://github.com/team-telnyx/telnyx-go/commit/e4fa420246a2eedf7cc3404ba1ec461e8dfd75fb))
* **api:** manual updates ([2849feb](https://github.com/team-telnyx/telnyx-go/commit/2849febcd3e99e790c9d7407f42aad7c389ca96f))
* **api:** manual updates ([65ff32c](https://github.com/team-telnyx/telnyx-go/commit/65ff32c76efa95022039be31e8d256f3bc946aa0))
* **api:** manual updates ([330ed0c](https://github.com/team-telnyx/telnyx-go/commit/330ed0c9f90ab108b6b0dce7c587000a7d0b6394))
* **api:** manual updates ([7bc6a84](https://github.com/team-telnyx/telnyx-go/commit/7bc6a84846a06ef3fc6a11054d8ea92055825ea0))
* **api:** manual updates ([9c2d4ed](https://github.com/team-telnyx/telnyx-go/commit/9c2d4ed0630e20d586e625fa392ef2995b41b14b))
* **api:** manual updates ([8e988be](https://github.com/team-telnyx/telnyx-go/commit/8e988bead528fc7b7a5a5bf05538d93ba7ae4a18))
* **api:** manual updates ([3a4567c](https://github.com/team-telnyx/telnyx-go/commit/3a4567c820ca2c1e568d2d4297827740cab65da9))
* **api:** manual updates ([05e4284](https://github.com/team-telnyx/telnyx-go/commit/05e4284a15ffbede1a94aef1276a2cfb3fe0a672))
* **api:** manual updates ([ba56cae](https://github.com/team-telnyx/telnyx-go/commit/ba56cae11a527d4a47af1e3a988177f3808c7ef2))
* **api:** manual updates ([f039a3b](https://github.com/team-telnyx/telnyx-go/commit/f039a3b2e64b14e9a6358de81ebe983c6af0a5a6))
* **api:** manual updates ([a276b8c](https://github.com/team-telnyx/telnyx-go/commit/a276b8ce08732afcd75db68fdde6c6ba5319e5ad))
* **api:** manual updates ([d80b22f](https://github.com/team-telnyx/telnyx-go/commit/d80b22f320dd8bd64f1493dd09be9020aaa21e9a))
* **api:** manual updates ([2e49cc1](https://github.com/team-telnyx/telnyx-go/commit/2e49cc15399b2f2364f4b81e2c6fdf936dbbdcfc))
* **api:** manual updates ([8a68f65](https://github.com/team-telnyx/telnyx-go/commit/8a68f65ac73667978ba09020b8207656252e348f))
* **api:** manual updates ([74fc7a6](https://github.com/team-telnyx/telnyx-go/commit/74fc7a688015481f6cd3885a30ef7cbd1f5fa99e))
* **api:** manual updates ([fa978ee](https://github.com/team-telnyx/telnyx-go/commit/fa978ee35787e461389a63183edd1aee3895dfd9))
* **api:** manual updates ([588316a](https://github.com/team-telnyx/telnyx-go/commit/588316af4853eb016c89fd485e36f11a7f616658))
* **api:** manual updates ([04a8826](https://github.com/team-telnyx/telnyx-go/commit/04a88263d169bd589cec0979e614b24de9a09823))
* **api:** manual updates ([36b9d7c](https://github.com/team-telnyx/telnyx-go/commit/36b9d7c9c5a8db9c848c5c3dd717d358651855fb))
* **api:** manual updates ([ce79b9b](https://github.com/team-telnyx/telnyx-go/commit/ce79b9ba2cb9ab1bb20571f29901412775476ec4))
* **api:** manual updates ([3ceb652](https://github.com/team-telnyx/telnyx-go/commit/3ceb6528687354b16b6315291c255bffd6e75b34))
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
* **api:** manual updates ([3a79955](https://github.com/team-telnyx/telnyx-go/commit/3a799553c1d33bb3f7e426d29ffc8b81b13735b7))
* **api:** manual updates ([a5a42ff](https://github.com/team-telnyx/telnyx-go/commit/a5a42ffc85a8e41462ff83f005ae22c890310aa1))
* **api:** manual updates ([2893c31](https://github.com/team-telnyx/telnyx-go/commit/2893c3117ead11141e37121ff7eea8425718a6b9))
* **api:** manual updates ([99ede75](https://github.com/team-telnyx/telnyx-go/commit/99ede75268ac63f6a956fff647823b402ee23299))
* **api:** manual updates ([f052988](https://github.com/team-telnyx/telnyx-go/commit/f0529880e4cebf4dcc2c66ff617411842fc8e70d))
* **api:** manual updates ([c47f3df](https://github.com/team-telnyx/telnyx-go/commit/c47f3dfd74fbfbbd9d9d793ee54e54f8fddba347))
* **api:** manual updates ([8b42453](https://github.com/team-telnyx/telnyx-go/commit/8b4245313b1956bf1717d87ce7f2327244dadf47))
* **api:** manual updates ([0b31743](https://github.com/team-telnyx/telnyx-go/commit/0b317439b1eeacc510eaebd28dfb33210e9f635f))
* **api:** manual updates ([0a83d13](https://github.com/team-telnyx/telnyx-go/commit/0a83d13f80fcd2ec2fe0b9bf79086bf35f472680))
* **api:** manual updates ([78932f5](https://github.com/team-telnyx/telnyx-go/commit/78932f5f835ad423109ea5f4c02996662ba84d83))
* **api:** manual updates ([93c71a2](https://github.com/team-telnyx/telnyx-go/commit/93c71a2e755c61b8bc438896aa435ae1c889e32d))
* **api:** manual updates ([b083a99](https://github.com/team-telnyx/telnyx-go/commit/b083a99c75c8ece1ae460fd3016650d175476425))
* **api:** manual updates ([dde48eb](https://github.com/team-telnyx/telnyx-go/commit/dde48ebf8431b668165f85644a7bb65a46ee5949))
* **api:** manual updates ([187f3ae](https://github.com/team-telnyx/telnyx-go/commit/187f3ae3be79e8ba806f20f750ed39187d58e34c))
* **api:** manual updates to include models ([3df9123](https://github.com/team-telnyx/telnyx-go/commit/3df9123f402b28acb9bcab99cbf138f6505ccdef))
* **api:** Merge pull request [#22](https://github.com/team-telnyx/telnyx-go/issues/22) from stainless-sdks/add-all-webhook-models ([ace72ab](https://github.com/team-telnyx/telnyx-go/commit/ace72abaeeb0fb0843342f491d6a7e74b30ef5d0))
* **api:** Merge pull request [#23](https://github.com/team-telnyx/telnyx-go/issues/23) from stainless-sdks/fix/deepgram-nova3-enum-duplicates ([4be0faf](https://github.com/team-telnyx/telnyx-go/commit/4be0faf59a2dd2e6cc0d53a145fb29bcfceed367))
* **api:** Merge pull request [#27](https://github.com/team-telnyx/telnyx-go/issues/27) from stainless-sdks/fix/whatsapp-message-templates-conflict ([5315a88](https://github.com/team-telnyx/telnyx-go/commit/5315a8802e0d7c6227888e19aeb4c38f821c23e4))
* **api:** Merge pull request [#29](https://github.com/team-telnyx/telnyx-go/issues/29) from stainless-sdks/fix-add-voice-model ([f467ebe](https://github.com/team-telnyx/telnyx-go/commit/f467ebe8ca34038ddfe1b4f65da0322623006c57))
* **api:** Merge pull request [#30](https://github.com/team-telnyx/telnyx-go/issues/30) from stainless-sdks/fix-schemaUnionDiscriminatorMissing ([2703e13](https://github.com/team-telnyx/telnyx-go/commit/2703e13a80654de2f5c6532bb0ccd3b2af6a11cd))
* **api:** revert bad update to spec ([a7ee2ac](https://github.com/team-telnyx/telnyx-go/commit/a7ee2ac7a5c0682948f56e6fb4c78b4c2ab428aa))
* **api:** update OpenAPI spec or Stainless config ([f467ebe](https://github.com/team-telnyx/telnyx-go/commit/f467ebe8ca34038ddfe1b4f65da0322623006c57))
* Assistant tags ([e306258](https://github.com/team-telnyx/telnyx-go/commit/e30625829873a1548ddc547a2b821405aa0fb4e2))
* Changing a tag for filebased STT endpoint ([f25641d](https://github.com/team-telnyx/telnyx-go/commit/f25641dac47162789c2cf2cc75c2d2ff458884c4))
* Changing the tag for TTS endpoint ([2696d43](https://github.com/team-telnyx/telnyx-go/commit/2696d43be1a53528713a5bb3267852e06f0259ce))
* **client:** add a convenient param.SetJSON helper ([3f86c04](https://github.com/team-telnyx/telnyx-go/commit/3f86c04b44645675c134f180fdb759bb662ccc6e))
* CW-2881 publish wireless VoLTE docs to prod ([1bb93bd](https://github.com/team-telnyx/telnyx-go/commit/1bb93bdf9f94f67f755fad3166ebb8660d9b9e22))
* CW-2881 update `filter[action_type]` enum ([ec67504](https://github.com/team-telnyx/telnyx-go/commit/ec675041339b94d93dc3f5ccb12bfbc6c6a7847a))
* Deploy dev/mc vady wip ([6a73d0d](https://github.com/team-telnyx/telnyx-go/commit/6a73d0d6935178d50be52b5817451dc2ff65195d))
* Deploy dev/mc vady wip ([564963e](https://github.com/team-telnyx/telnyx-go/commit/564963e91a7c840de81f446044e3a0a45434b882))
* Engdesk 44932 ([e71c63b](https://github.com/team-telnyx/telnyx-go/commit/e71c63b2a06d8771dbd8e6eb7a4b3be08db7e61d))
* ENGDESK-44767 - Document force remove calls from queue ([b3a98a3](https://github.com/team-telnyx/telnyx-go/commit/b3a98a317d7fd1dd05a190a1a114a7eb906af552))
* ENGDESK-45429 - Add sip_region documentation for dial and transfer command ([843bd25](https://github.com/team-telnyx/telnyx-go/commit/843bd2589807970eae88209a4ee83f16445d976c))
* ENGDESK-46399 - Add sip_call_id filter for retreiving recordings ([e991c87](https://github.com/team-telnyx/telnyx-go/commit/e991c874c6e3b135ecb64fbd5d62fbab3f1a30d9))
* ENGDESK-47508 - part 2 shared schema fixes ([a609ed4](https://github.com/team-telnyx/telnyx-go/commit/a609ed49bbb6e3a8ec5d87ec252451096b686dc6))
* ENGDESK-47518 document mobile number and mobile voice connection endpoints ([8d5daa1](https://github.com/team-telnyx/telnyx-go/commit/8d5daa1baaf3d4e9872a4df519051f57e2faf971))
* ENGDESK-49554: Add quail_voice_focus to noise suppression engine enums ([4153ed3](https://github.com/team-telnyx/telnyx-go/commit/4153ed36aece7e5cf7f166142344a414cfb485d7))
* Fix broken documentation links ([780485e](https://github.com/team-telnyx/telnyx-go/commit/780485e3ad168f47464b6d89487d06d5600bdcb1))
* Fix invalid responses ([a12be41](https://github.com/team-telnyx/telnyx-go/commit/a12be4171bf128f2ff75c0a15d9a48a0103b5318))
* fix links ([6f34e3b](https://github.com/team-telnyx/telnyx-go/commit/6f34e3bdf6f96087b044a9bc3ba914115b40d4e1))
* Fix listing deepgram languages for transcription start ([e319d9f](https://github.com/team-telnyx/telnyx-go/commit/e319d9ff3ce7f7292286d96c4448e14b1963d25a))
* Fix Redocly linting errors and warnings in TDA reporting specs ([db057fc](https://github.com/team-telnyx/telnyx-go/commit/db057fc7cd94427caefb07d51a146e6f9c302213))
* Fix Redocly linting warnings in Number Lookup spec ([9989e5f](https://github.com/team-telnyx/telnyx-go/commit/9989e5f607b017dcf499f9fc66941eff9513aee7))
* Fix Redocly linting warnings in OAuth and Integration Secrets specs ([f928f24](https://github.com/team-telnyx/telnyx-go/commit/f928f2414b11a1f2727187ff7c49963e87601286))
* fix schema in emergency address schema ([057f28e](https://github.com/team-telnyx/telnyx-go/commit/057f28e2ee1d6d5653f832c7ae95bcfbc46bdaec))
* fix-external-connection-link ([8ee88dd](https://github.com/team-telnyx/telnyx-go/commit/8ee88ddbfb0808792eecfd634437387aab4956d3))
* fix-redocly-lint-issues ([6bf8f43](https://github.com/team-telnyx/telnyx-go/commit/6bf8f43f11a089047b50f595e0b9c3726125d307))
* fix-stainless-sdk-timeout ([0539121](https://github.com/team-telnyx/telnyx-go/commit/0539121c5bd1607d05452571688d31b20fe01c0a))
* Fixing lint errors ([f1c67db](https://github.com/team-telnyx/telnyx-go/commit/f1c67db62e1da5b08ef8b8125e22ec168568b4ab))
* **internal:** support comma format in multipart form encoding ([df166a3](https://github.com/team-telnyx/telnyx-go/commit/df166a35358164b24e810ce12ea13b1b0bd832e2))
* jira-engdesk-48800 add organizations-related docs to the external api… ([9db7459](https://github.com/team-telnyx/telnyx-go/commit/9db745916308e3ada360be4a8db9d4badf77106e))
* jira-engdesk-49089 add new connection jitter buffer related fields ([ca56153](https://github.com/team-telnyx/telnyx-go/commit/ca56153cb568be3bfc459cf55fa1edd28d03f5ca))
* Limit detection_mode enum to disabled and detect only ([d7025cc](https://github.com/team-telnyx/telnyx-go/commit/d7025cc29ac8d294786c0c18a523a392dd8929ba))
* Merge TTS file-based spec into text-to-speech.json ([63d1363](https://github.com/team-telnyx/telnyx-go/commit/63d1363bb82cc816534ba51acfe0d844e8951470))
* **messaging:** add wait_seconds to OutboundMessagePayload example ([db135bd](https://github.com/team-telnyx/telnyx-go/commit/db135bd23a00b7e1389b8e4e170c8c1c616e9a2c))
* MSG-6076: webhook event for 10DLC campaign suspended status ([dd97898](https://github.com/team-telnyx/telnyx-go/commit/dd97898a1c4dab08927344fdab21ef426398e1e3))
* MSG-6148: adding the new campaignVerifyAuthorizationToken field and missing GET OTP endpoint ([f61e3d0](https://github.com/team-telnyx/telnyx-go/commit/f61e3d006842f811275164ef559fcc2b107a5b67))
* MSG-6228: MSG-6228: Add smart_encoding option for SMS character encoding optimization ([bc558ea](https://github.com/team-telnyx/telnyx-go/commit/bc558ea17d0591a918f7cfb4cc5d46bc996b8fd2))
* MSG-6418: remove flashcall from hosted number verification codes endpoint ([e63065b](https://github.com/team-telnyx/telnyx-go/commit/e63065b60d630e78ba7c2e26e0eb922d2e753d61))
* New tools api ([e26a190](https://github.com/team-telnyx/telnyx-go/commit/e26a1904e60addbc999d6619f31dabab34a05ba0))
* port-4690: fix LOA configuration preview path (singular → plural) ([6883e97](https://github.com/team-telnyx/telnyx-go/commit/6883e97470842a64acf5ea216ab6e01819e786c9))
* PORTAL-5923: Add stored_payment_transactions endpoint to OpenAPI docs ([5c565eb](https://github.com/team-telnyx/telnyx-go/commit/5c565eb70d4b7ed940f2aff53cff624e69b74f7c))
* Refactored README to only contain useful information and reflect accu… ([f55a615](https://github.com/team-telnyx/telnyx-go/commit/f55a615d3673e60729593427a8bf28e184be65fa))
* Retell assistants import ([a9ad58c](https://github.com/team-telnyx/telnyx-go/commit/a9ad58c0acf2389975fd6f1f835a229bf17b46ad))
* Revert "fix emergency settings -schema" ([e249c20](https://github.com/team-telnyx/telnyx-go/commit/e249c20088e277ada7ef48af9b51be59bdb05d36))
* **stt:** add WebSocket event schemas for Stainless SDK generation ([2d61b11](https://github.com/team-telnyx/telnyx-go/commit/2d61b1195d18374a0eda700f6ca8d7a0fd3a2f17))
* TELAPPS Add ApplicationSid param ([102472f](https://github.com/team-telnyx/telnyx-go/commit/102472f5311ae02006965c37019b54fb25928908))
* TELAPPS Add GET /texml/Accounts/{account_sid}/Queues endpoint ([84dd1c7](https://github.com/team-telnyx/telnyx-go/commit/84dd1c799cd10bb9f59667b07e754d362453edcf))
* TELAPPS Add interim_results to deepgram config ([56a7d7d](https://github.com/team-telnyx/telnyx-go/commit/56a7d7dda5b1e1e8eafabd57f1ed9032f9757e6d))
* TELAPPS Add prevent_double_bridge to bridge command ([ff9c14b](https://github.com/team-telnyx/telnyx-go/commit/ff9c14b06c3ba5a657a676a7c34751fafb87310b))
* TELAPPS-5367: Update transcription start docs ([0f3bc14](https://github.com/team-telnyx/telnyx-go/commit/0f3bc14b02902d358ddd93e58fc3952f943f7aba))
* TELAPPS-5399 Add region to conference commands ([781cf65](https://github.com/team-telnyx/telnyx-go/commit/781cf6551ed68de15b219e75cd0e8272103a55ef))
* TELAPPS-5459: Add Azure to transcription start ([7aedf5b](https://github.com/team-telnyx/telnyx-go/commit/7aedf5b0da57d85573ca609adf7c3edc415127b1))
* TELAPPS-5507: Add Krisp engine description for noise suppression ([caf9b12](https://github.com/team-telnyx/telnyx-go/commit/caf9b12eee9c49f880e67e4df8306e9541d57ae1))
* TELAPPS-5668: Add call.cost webhook event documentation ([2bfe79a](https://github.com/team-telnyx/telnyx-go/commit/2bfe79a53c35433bc62a469161eab66ff968199f))
* TELAPPS-5685: Add store_fields_as_variables to WebhookToolParams ([4adf279](https://github.com/team-telnyx/telnyx-go/commit/4adf2792dc32b8a53906382888d6e76ed2f4f5b5))
* TELAPPS-ENGDESK-46395 Add keep_after_hangup to enqueue command ([abf7cda](https://github.com/team-telnyx/telnyx-go/commit/abf7cda0d1cad1288c43337697b8220cec4be7d0))
* TELAPPS-ENGDESK-46395 Add PATCH /queues/{queue_name}/calls/{call_control_id} endpoint ([b020913](https://github.com/team-telnyx/telnyx-go/commit/b020913fb45e6f3448170e2966d7b63d8472b941))
* TELAPPS-ENGDESK-47967 Add black_threshold parameter to send_fax request ([fa64831](https://github.com/team-telnyx/telnyx-go/commit/fa64831aebc7cc5dc19b9d7cd68cb471f8bdcf24))
* TELAPPS-ENGDESK-48790 Remove duplicated webhooks ([e562f9e](https://github.com/team-telnyx/telnyx-go/commit/e562f9e5a332ac7e2ad9d18c9b59bc7b9c1247e3))
* TELAPPS-ENGDESK-48951 add channels to conf record start ([027f1a6](https://github.com/team-telnyx/telnyx-go/commit/027f1a6c8c96e41c9c2d6ffd2f18cb24c5341074))
* TELAPPS-ENGDESK-49737 Add prevent_double_bridge param to dial ([ebfa5ed](https://github.com/team-telnyx/telnyx-go/commit/ebfa5ed5baea518a69b229a50c83e70dd7cb82cb))
* Update voicemail_detection description with AMD enablement info ([b30ddbb](https://github.com/team-telnyx/telnyx-go/commit/b30ddbba1d26436a7ded8843dc3b03eb81822fc1))
* Updated README to include the step for make buildcontainer bundle to … ([b5be3dc](https://github.com/team-telnyx/telnyx-go/commit/b5be3dc8e6e18bd4a82e44aaca252fc5c0d444ce))
* **websocket:** add STT/TTS WebSocket streaming support ([a18dc69](https://github.com/team-telnyx/telnyx-go/commit/a18dc690ac9eae51647aa59fb3deded83de13280))
* **wireless:** add traffic policy profiles endpoints to OpenAPI spec ([d391fc9](https://github.com/team-telnyx/telnyx-go/commit/d391fc99ca58d7dd126b112666f61b7607fe3636))


### Bug Fixes

* add discriminator to TranscriptionEngineDeepgramConfig for Stainless SDK ([ec0f1c5](https://github.com/team-telnyx/telnyx-go/commit/ec0f1c5e632fe44478f75f795025980c10fd40d7))
* add discriminator to TtsServerEvent for Stainless SDK generation ([dc56b05](https://github.com/team-telnyx/telnyx-go/commit/dc56b052c5ea2c3030bf3acd1121fcf4906b0689))
* add missing vertical enum values for 10DLC brand creation (ENGDESK-49040) ([451baf3](https://github.com/team-telnyx/telnyx-go/commit/451baf33567ae89cf5e7cd9cdd5e184cd768265f))
* add title to InviteTool.invite for Stainless SDK ([156a718](https://github.com/team-telnyx/telnyx-go/commit/156a7183ce71a8f9df09d58c74296f4bdefbf777))
* allow canceling a request while it is waiting to retry ([09da6eb](https://github.com/team-telnyx/telnyx-go/commit/09da6eb09775e6ae1ddf9f581a2ee808b673cda1))
* **call-recordings:** align OpenAPI spec with implementation ([868ec49](https://github.com/team-telnyx/telnyx-go/commit/868ec49125b4bfb36f3a39309c6f86cd6689d47d))
* **client/oauth:** send grant_type in the right location ([43b8e4f](https://github.com/team-telnyx/telnyx-go/commit/43b8e4fd847029a752796d0df07956e5ee0600ef))
* **client:** correctly specify Accept header with */* instead of empty ([d633f03](https://github.com/team-telnyx/telnyx-go/commit/d633f039cd21a3f7a6ef78a86b3f4c69ad00bfee))
* **client:** fix issue with example webhook payload ([6f53e1d](https://github.com/team-telnyx/telnyx-go/commit/6f53e1dfbe1363126d2e46c4a0454ecf24f0a121))
* **client:** invalid URL ([3b07951](https://github.com/team-telnyx/telnyx-go/commit/3b0795170041a6b8e6e0f717d101e6ca4f822515))
* **client:** revert change to certain pagination metadata types ([d5cb536](https://github.com/team-telnyx/telnyx-go/commit/d5cb5368b2f66130420dde980df45d9b9afd0528))
* **client:** use correct format specifier for header serialization ([49fed3a](https://github.com/team-telnyx/telnyx-go/commit/49fed3abfb9918d7bf3dd3ec7e23d6690cd57196))
* correct broken link to List SIM Card Actions endpoint in SIM car… ([b4fe5fa](https://github.com/team-telnyx/telnyx-go/commit/b4fe5fad4bfa5a192c0346dba9ee81121c7d9c5e))
* **docs:** add missing pointer prefix to api.md return types ([681874c](https://github.com/team-telnyx/telnyx-go/commit/681874c00a9e4f4d8333dc2be629f85d80179465))
* **docs:** fix mcp installation instructions for remote servers ([c53abe5](https://github.com/team-telnyx/telnyx-go/commit/c53abe57c7cf0cef01e5aaf318063f7f4f096ef6))
* **encoder:** correctly serialize NullStruct ([37388a9](https://github.com/team-telnyx/telnyx-go/commit/37388a9111cd901c315c8a71f257b99c72a4d022))
* fix issue with unmarshaling in some cases ([16d2d96](https://github.com/team-telnyx/telnyx-go/commit/16d2d96e436579895f7275d077a22f044f843008))
* fix request delays for retrying to be more respectful of high requested delays ([97f3739](https://github.com/team-telnyx/telnyx-go/commit/97f3739398b0363c299ee8b04727c42b53968bf0))
* fix test cases where required client options were not set ([87ba3fc](https://github.com/team-telnyx/telnyx-go/commit/87ba3fc26c4bd83141e9efd6ac5fc8f53649b4e4))
* fixes for pagination and iteration, plus iter.Seq support ([ff6b2c5](https://github.com/team-telnyx/telnyx-go/commit/ff6b2c55227ff881bee8c133d33ea4ee69fefd74))
* **internal:** skip tests that depend on mock server ([23c9bec](https://github.com/team-telnyx/telnyx-go/commit/23c9bec1d47e3950698904c6abd53d5010a57c46))
* move unsupported string formats to x-format ([e013fcc](https://github.com/team-telnyx/telnyx-go/commit/e013fccb33d3a96e6e489aa700d5228df172ced3))
* narrow porting event_type enums for SDK discriminator support ([ec66098](https://github.com/team-telnyx/telnyx-go/commit/ec660985877c4a68c0ae35dd1c388fadc8fde8af))
* OAS drift — 10dlc.json (messaging-campaign-registry) ([fc6998b](https://github.com/team-telnyx/telnyx-go/commit/fc6998b7869607d134cf91903d60018528b2c3a1))
* OAS drift — messaging.json (messaging-settings + messaging-outbound) ([1fc474c](https://github.com/team-telnyx/telnyx-go/commit/1fc474cae78973892d96ec6e81b832b712391e0f))
* OAS drift — toll-free-verification.json (messaging-tf-verify) ([7b9e913](https://github.com/team-telnyx/telnyx-go/commit/7b9e91365cd6210db6c39f9ecedb56b3613cc67b))
* OAS drift — verify.json (messaging-2fa) ([bcbc092](https://github.com/team-telnyx/telnyx-go/commit/bcbc09281e2d908ce311b8ff200a46c1d9d216f2))
* prevent duplicate ? in query params ([89e80d1](https://github.com/team-telnyx/telnyx-go/commit/89e80d12547e98b2ed8ccbadd5e3ae40e88072fa))
* reference models more accurately ([ce82ddf](https://github.com/team-telnyx/telnyx-go/commit/ce82ddfa825000f123c20747494ca1bf9ed19b64))
* remove deprecated TeXML API endpoints from OpenAPI spec ([ee92cc4](https://github.com/team-telnyx/telnyx-go/commit/ee92cc48a333675c5cdae1c247cc42bddb67ad8c))
* remove invalid discriminators from string enum schemas ([22dc0b8](https://github.com/team-telnyx/telnyx-go/commit/22dc0b8802420a01e1d91d9c1ede1d430041413a))
* remove readonly parameters from request params ([239b1fe](https://github.com/team-telnyx/telnyx-go/commit/239b1fe3dd9653852ed5afc9cba47237ac7f2b4b))
* rename number-reputation ToS route to use underscores ([3b43ea9](https://github.com/team-telnyx/telnyx-go/commit/3b43ea9c91c41d65b96249e527744acee481d84d))
* rename whatsapp.message_templates to whatsapp.templates to avoid conflict ([5315a88](https://github.com/team-telnyx/telnyx-go/commit/5315a8802e0d7c6227888e19aeb4c38f821c23e4))
* StringFormatNotSupported ([35ad893](https://github.com/team-telnyx/telnyx-go/commit/35ad893f7e8865c595a3bde6f402894dde410fc9))
* update broken MDR report link in GetMessage endpoint ([1c9a740](https://github.com/team-telnyx/telnyx-go/commit/1c9a7404e73026603e27d803ca28ea2de6f2edc4))
* update wait_seconds example to 0.5 ([122039b](https://github.com/team-telnyx/telnyx-go/commit/122039b13d398e38c6b2e6c06d924b209175c1ba))
* use PaginationMeta schema for ListFaxesResponse.meta ([65bd7f1](https://github.com/team-telnyx/telnyx-go/commit/65bd7f1dac7ae831cec9f929fb1ff4f1913451f5))


### Reverts

* restore stainless.yml changes removed in 1de6067 ([14d71cb](https://github.com/team-telnyx/telnyx-go/commit/14d71cb7e4c18059341a6fdb2e9aec862ed1f1e6))
* restore stainless.yml SDK generation fixes ([ff4d9bd](https://github.com/team-telnyx/telnyx-go/commit/ff4d9bdf3516493819710fc80d2741ce1b35ccd0))


### Chores

* add release-please workflow + fix config for STLC cutover ([#111](https://github.com/team-telnyx/telnyx-go/issues/111)) ([ef499b8](https://github.com/team-telnyx/telnyx-go/commit/ef499b868d7b8bb8e56f2b4018f4a3190f2b769b))
* bring back other changes ([81f62a8](https://github.com/team-telnyx/telnyx-go/commit/81f62a8ffb1af5998aa4fdd25ae29d5a251f8ee1))
* bump gjson version ([3522560](https://github.com/team-telnyx/telnyx-go/commit/35225603ffc2f3658c8ee40742beb9660c0df7ad))
* **ci:** add build step ([2663029](https://github.com/team-telnyx/telnyx-go/commit/266302903f0a11fd8879d104fda00e51eabd128f))
* **ci:** skip lint on metadata-only changes ([0ad9731](https://github.com/team-telnyx/telnyx-go/commit/0ad9731e6bc3cdedeee5256044c9e266ddaaf3e5))
* **ci:** skip uploading artifacts on stainless-internal branches ([e97ec8e](https://github.com/team-telnyx/telnyx-go/commit/e97ec8e8dd9e6ebfa821db8072dad88e695c2c37))
* **ci:** support opting out of skipping builds on metadata-only commits ([7b21a38](https://github.com/team-telnyx/telnyx-go/commit/7b21a3869a1e5d30343df86dbaf5e4a37dc2fa8d))
* **client:** fix multipart serialisation of Default() fields ([bd55ee3](https://github.com/team-telnyx/telnyx-go/commit/bd55ee34f3ea7fc6ac5ca9e6fea5236efd3697c2))
* configure new SDK language ([ec4fa60](https://github.com/team-telnyx/telnyx-go/commit/ec4fa609d6e1cd0d061975e12596f1e4cc8e9bac))
* **docs:** add missing descriptions ([004cb96](https://github.com/team-telnyx/telnyx-go/commit/004cb967e6ee668fd2e8e03338947957b619f575))
* fix empty interfaces ([2f6f406](https://github.com/team-telnyx/telnyx-go/commit/2f6f4065674e7e40141c707bbd630f86cd04f9a6))
* **internal:** codegen related update ([01b9ca6](https://github.com/team-telnyx/telnyx-go/commit/01b9ca615baa27063a1c245c087a5fe4c8615228))
* **internal:** codegen related update ([3c657ff](https://github.com/team-telnyx/telnyx-go/commit/3c657ff23a59786d0ee198da41726bf2ac48334d))
* **internal:** grammar fix (it's -&gt; its) ([c9c78d6](https://github.com/team-telnyx/telnyx-go/commit/c9c78d643f660d07aa1400eabd3549f08f575841))
* **internal:** minor cleanup ([6cfbfb4](https://github.com/team-telnyx/telnyx-go/commit/6cfbfb40a0d9c78cdce226310f547fe1de5f04fe))
* **internal:** move custom custom `json` tags to `api` ([26fdddf](https://github.com/team-telnyx/telnyx-go/commit/26fdddf04417aa259dfd56758bc429692e5818e2))
* **internal:** remove mock server code ([a087122](https://github.com/team-telnyx/telnyx-go/commit/a087122515a694d6480bcac23842ba6d67629f70))
* **internal:** support default value struct tag ([feef5f4](https://github.com/team-telnyx/telnyx-go/commit/feef5f4e36b452b354cbd24f9fbc4729b4aa3faf))
* **internal:** tweak CI branches ([252cc70](https://github.com/team-telnyx/telnyx-go/commit/252cc701a02d9625c8deb47e728f52cf22b74113))
* **internal:** update `actions/checkout` version ([2ea69ea](https://github.com/team-telnyx/telnyx-go/commit/2ea69eadcc78e92bffb76e3b5d17319423a26b0d))
* **internal:** update gitignore ([0edd0eb](https://github.com/team-telnyx/telnyx-go/commit/0edd0eb29c4eadf794b3c4512e974d0ac5704dea))
* **internal:** use explicit returns ([4926d1b](https://github.com/team-telnyx/telnyx-go/commit/4926d1b6241fa93277143e7eb27f50412417b357))
* **internal:** use explicit returns in more places ([3ea4295](https://github.com/team-telnyx/telnyx-go/commit/3ea4295e5243761699d3af3a5976b63839ce0ad1))
* remove leaked staging-only promote-to-prod.yml from prod ([#115](https://github.com/team-telnyx/telnyx-go/issues/115)) ([b387bb0](https://github.com/team-telnyx/telnyx-go/commit/b387bb0022da6c5d2460f92feda6ec7bd046c8cf))
* remove unnecessary error check for url parsing ([56a5052](https://github.com/team-telnyx/telnyx-go/commit/56a505257d97beea70839c2f1c842535e4b95c58))
* sync repo ([104336e](https://github.com/team-telnyx/telnyx-go/commit/104336e664f1be532131fe3fb5c2b3edd31d86cc))
* **tests:** update webhook tests ([df6882c](https://github.com/team-telnyx/telnyx-go/commit/df6882c2f53d85bbee9789241b36180dd7d6f023))
* update docs for api:"required" ([e75977d](https://github.com/team-telnyx/telnyx-go/commit/e75977dc5f544d929d7782d97929d95fe58361d7))
* update mock server docs ([d036d12](https://github.com/team-telnyx/telnyx-go/commit/d036d1256c436129c404cf9b260a603daafb84b5))
* update placeholder string ([db53e98](https://github.com/team-telnyx/telnyx-go/commit/db53e9896d748cb3adfe0b449337552c4e7b6be9))


### Documentation

* add service_provider_login_url to authentication provider settings ([3886cab](https://github.com/team-telnyx/telnyx-go/commit/3886cabb888880a690f4d43be249b3637d231787))
* **branded-calling:** add Number Reputation API specs ([a115364](https://github.com/team-telnyx/telnyx-go/commit/a115364f426f63498368b51d7a9939fb300168cd))
* **call-control:** Add missing conference endpoints ([d3f4cbb](https://github.com/team-telnyx/telnyx-go/commit/d3f4cbb0d5ae20b79d7f724fd0e257eb2ee339e4))
* **call-control:** Add missing parameters to call control endpoints ([622b420](https://github.com/team-telnyx/telnyx-go/commit/622b4209934b5903e5964f3353dd0189c4bb65d5))
* **call-control:** Add missing params to hangup, bridge, answer ([c3a6ea4](https://github.com/team-telnyx/telnyx-go/commit/c3a6ea4093368e8fb7c17882fa55f2f29d11a866))
* **call-control:** Add queue CRUD endpoints ([1497f4a](https://github.com/team-telnyx/telnyx-go/commit/1497f4a3d382aef18ecfb1a4cacdc5699e25a715))
* **call-scripting:** add Timeout and TimeLimit to InitiateTexmlCall ([293b294](https://github.com/team-telnyx/telnyx-go/commit/293b29447e7b9d9384132777ab9938f56b9428c3))
* fix voice settings available voices link ([4ffdd82](https://github.com/team-telnyx/telnyx-go/commit/4ffdd8250936550b5dd96ae8cfdcddf04e1eb21d))
* **messaging:** Add wait_seconds to message response schemas ([95a7442](https://github.com/team-telnyx/telnyx-go/commit/95a744295629ff2649d18ff83e999073c5903f32))
* **tts:** Add Telnyx.Ultra model documentation ([296412a](https://github.com/team-telnyx/telnyx-go/commit/296412a2ee2e41511a8f90c718d7565ee125ca6a))
* WhatsApp template components schema ([9db4c78](https://github.com/team-telnyx/telnyx-go/commit/9db4c78e415145f8e7a9c11b6e3566d90859e156))


### Refactors

* move webhook verification to lib package to avoid merge conflicts ([b7a0fec](https://github.com/team-telnyx/telnyx-go/commit/b7a0fec12ea6542b37876c798b86487a2d69a365))

## [4.107.1](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.107.0...v4.107.1) (2026-06-26)


### Chores

* add promote-to-prod workflow for STLC cutover ([b60c57b](https://github.com/team-telnyx/telnyx-go-staging/commit/b60c57b19c578f84e0c2eb3c03d4263475f7a092))

## [4.107.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.106.0...v4.107.0) (2026-06-26)


### Chores

* preserve repo-owned files not part of SDK generation ([d8a2fb0](https://github.com/team-telnyx/telnyx-go-staging/commit/d8a2fb03ab7b0588ba9f5546dbaf7333cd3f64f0))
* release go 4.107.0 ([b59882b](https://github.com/team-telnyx/telnyx-go-staging/commit/b59882bd25e457a1b0490db3d5cf0d8981815f4d))
* sync OpenAPI spec from 9b36eb0 ([dd93a58](https://github.com/team-telnyx/telnyx-go-staging/commit/dd93a58cec7ea5b5ca32d1572e3d862665531dec))
* sync OpenAPI spec from 9b36eb0 ([78b6f40](https://github.com/team-telnyx/telnyx-go-staging/commit/78b6f40c295f8ee9aed43f873846739ca8a0b77f))

## [4.106.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.105.0...v4.106.0) (2026-06-26)


### Bug Fixes

* **go:** enable V1 back-compat naming to eliminate 732 breaking type renames ([d52e43e](https://github.com/team-telnyx/telnyx-go-staging/commit/d52e43edd8089acf336be575e9149ec06eb4cbd3))
* **go:** enable V1 back-compat naming to eliminate 732 breaking type renames ([b9ea432](https://github.com/team-telnyx/telnyx-go-staging/commit/b9ea432c373fb116cd564112d0e83a681c32de1f))


### Chores

* preserve repo-owned files not part of SDK generation ([8756b81](https://github.com/team-telnyx/telnyx-go-staging/commit/8756b81f581a6e19457dd99b76d553faefe3bdfe))
* release go 4.106.0 ([ef3952b](https://github.com/team-telnyx/telnyx-go-staging/commit/ef3952b02f3b87de731d8f27916d1f6f5ec3141c))

## [4.105.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.104.0...v4.105.0) (2026-06-24)


### Chores

* preserve repo-owned files not part of SDK generation ([6fd1b3f](https://github.com/team-telnyx/telnyx-go-staging/commit/6fd1b3ffbabc426fa27daf5fec81f4cbf191578a))
* release go 4.105.0 ([3f05ee3](https://github.com/team-telnyx/telnyx-go-staging/commit/3f05ee347389bcfbba3262d8f95ad6c1e21826ea))

## [4.104.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.103.0...v4.104.0) (2026-06-24)


### Bug Fixes

* type additionalProperties for 10dlc oneOf variants to fix Java deserialization ([32bba28](https://github.com/team-telnyx/telnyx-go-staging/commit/32bba28f3d4efcff0d804578608f1388de69918d))
* type additionalProperties for 10dlc oneOf variants to fix Java deserialization ([c33d6b0](https://github.com/team-telnyx/telnyx-go-staging/commit/c33d6b0cea4085bd721a078f01a6af33ab7eb565))


### Chores

* preserve repo-owned files not part of SDK generation ([4ea6d8b](https://github.com/team-telnyx/telnyx-go-staging/commit/4ea6d8bcb1a1607de901a980f6cee35977614bba))
* release go 4.104.0 ([c3f5056](https://github.com/team-telnyx/telnyx-go-staging/commit/c3f505650c2a57f8c8a6ace46c38e5d1f8c02c20))

## [4.103.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.102.0...v4.103.0) (2026-06-23)


### Chores

* preserve repo-owned files not part of SDK generation ([4cda026](https://github.com/team-telnyx/telnyx-go-staging/commit/4cda0260ec9be9b5f6c4521f31216ef26a7336cc))
* release go 4.103.0 ([c65f052](https://github.com/team-telnyx/telnyx-go-staging/commit/c65f0522518fd2697cd3dfb7c677a23e67ad6d38))

## [4.102.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.101.0...v4.102.0) (2026-06-23)


### Bug Fixes

* register TranscriptionEngineParakeetConfig as shared model in calls resource ([23efc8d](https://github.com/team-telnyx/telnyx-go-staging/commit/23efc8d9f3f6e72507e18619bd2a1cbe4f0f6b08))
* register TranscriptionEngineParakeetConfig as shared model in calls resource ([ee4ba39](https://github.com/team-telnyx/telnyx-go-staging/commit/ee4ba3976ddac67aea381ebd79e2af7cd078949b))


### Chores

* preserve repo-owned files not part of SDK generation ([6a7a4ef](https://github.com/team-telnyx/telnyx-go-staging/commit/6a7a4ef4146ee11610c1906088055e979249fef9))
* release go 4.102.0 ([0f82795](https://github.com/team-telnyx/telnyx-go-staging/commit/0f82795bc00e95129613d3512d1d5a8f08f0f994))

## [4.101.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.100.0...v4.101.0) (2026-06-23)


### Chores

* preserve repo-owned files not part of SDK generation ([e62e7bf](https://github.com/team-telnyx/telnyx-go-staging/commit/e62e7bfde078653a55684f917e20989a89f6eb8c))
* release go 4.101.0 ([2cdcf6a](https://github.com/team-telnyx/telnyx-go-staging/commit/2cdcf6aaeb2d7940f83f68385ae8a7b8c0bf5aee))

## [4.100.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.99.0...v4.100.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([325f06c](https://github.com/team-telnyx/telnyx-go-staging/commit/325f06cf0d762b6b7b1da3646fb0ab8f2aac1d2e))
* release go 4.100.0 ([daed440](https://github.com/team-telnyx/telnyx-go-staging/commit/daed440b3663fce8f1b3933fad7bde61c56f5a5e))

## [4.99.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.98.0...v4.99.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([28471bb](https://github.com/team-telnyx/telnyx-go-staging/commit/28471bbe2785d03cf1f42ab6a19a32c6035a775c))
* release go 4.99.0 ([afad10a](https://github.com/team-telnyx/telnyx-go-staging/commit/afad10aa21410f183bf9f5d0c632511c6317af0a))

## [4.98.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.97.0...v4.98.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([6b45b48](https://github.com/team-telnyx/telnyx-go-staging/commit/6b45b484123a83cc53e19dd4194def1259597d7f))
* release go 4.98.0 ([d23b37c](https://github.com/team-telnyx/telnyx-go-staging/commit/d23b37cc92ee212f217610de80cadda173a199da))

## [4.97.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.96.0...v4.97.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([1d75f8f](https://github.com/team-telnyx/telnyx-go-staging/commit/1d75f8f6bbf34a5616df8079dee673cca7633337))
* release go 4.97.0 ([a120669](https://github.com/team-telnyx/telnyx-go-staging/commit/a1206699883dc2d45d38624ccd80cde79770cee5))

## [4.96.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.95.0...v4.96.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([72830d6](https://github.com/team-telnyx/telnyx-go-staging/commit/72830d685e1e64a7a377432402094c0f7eb1da10))
* release go 4.96.0 ([00ef7db](https://github.com/team-telnyx/telnyx-go-staging/commit/00ef7db16823096e9f472c19f073a9af42d12cce))

## [4.95.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.94.0...v4.95.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([35dd8fb](https://github.com/team-telnyx/telnyx-go-staging/commit/35dd8fbb0fb59e75db297f4bfb160a67f4dc1f0a))
* release go 4.95.0 ([f8b3dc5](https://github.com/team-telnyx/telnyx-go-staging/commit/f8b3dc55b04c7f75ce25051b04fec732eb93068e))

## [4.94.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.93.0...v4.94.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([5e9ec83](https://github.com/team-telnyx/telnyx-go-staging/commit/5e9ec83176f342d9a74f15c04a81ab3c2b24da09))
* release go 4.94.0 ([01f8780](https://github.com/team-telnyx/telnyx-go-staging/commit/01f8780e16b924345faeb1ed44038b82ea55a564))

## [4.93.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.92.0...v4.93.0) (2026-06-22)


### Chores

* preserve repo-owned files not part of SDK generation ([8317b5d](https://github.com/team-telnyx/telnyx-go-staging/commit/8317b5d38b324806458ca774adecebfd5e886746))
* release go 4.93.0 ([2299d1c](https://github.com/team-telnyx/telnyx-go-staging/commit/2299d1ced3fe566537b974e023e5bee3a297548b))

## [4.92.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.91.0...v4.92.0) (2026-06-21)


### Chores

* preserve repo-owned files not part of SDK generation ([ddd5efe](https://github.com/team-telnyx/telnyx-go-staging/commit/ddd5efee7de78bae5e0ffe27f82c30d443c4a87d))
* release go 4.92.0 ([3518968](https://github.com/team-telnyx/telnyx-go-staging/commit/3518968f7e8601d7c25896881d8935c5c61a279f))

## [4.91.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.90.0...v4.91.0) (2026-06-21)


### Chores

* preserve repo-owned files not part of SDK generation ([1882842](https://github.com/team-telnyx/telnyx-go-staging/commit/18828428509305c27cfe8569e36b2fe1e1e5faaa))
* release go 4.91.0 ([d4e6b81](https://github.com/team-telnyx/telnyx-go-staging/commit/d4e6b8120cadaee291349103f88e62fcc3e763b2))

## [4.90.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.89.0...v4.90.0) (2026-06-20)


### Chores

* preserve repo-owned files not part of SDK generation ([14ba749](https://github.com/team-telnyx/telnyx-go-staging/commit/14ba7497dc52c448a049db530ddd88900e59f850))
* release go 4.90.0 ([9f265b6](https://github.com/team-telnyx/telnyx-go-staging/commit/9f265b6736114348b7315786997d08beaa479762))

## [4.89.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.88.0...v4.89.0) (2026-06-20)


### Chores

* preserve repo-owned files not part of SDK generation ([00f7a66](https://github.com/team-telnyx/telnyx-go-staging/commit/00f7a66dcf3335236ad292daec0367d2ebfa9d4d))
* release go 4.89.0 ([06b5e10](https://github.com/team-telnyx/telnyx-go-staging/commit/06b5e10f7ab8d90c787fc5cb77fe5ba6468893c6))

## [4.88.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.87.0...v4.88.0) (2026-06-18)


### Chores

* preserve repo-owned files not part of SDK generation ([94d4187](https://github.com/team-telnyx/telnyx-go-staging/commit/94d4187ec2f279fc5bc9ed56ec1bf1a48f454eb5))
* release go 4.88.0 ([f78c833](https://github.com/team-telnyx/telnyx-go-staging/commit/f78c833d0fb54dc0f3102d26e376142e76c43f26))
* sync OpenAPI spec from cde0bc4 ([a694f05](https://github.com/team-telnyx/telnyx-go-staging/commit/a694f05bbf0359ade1e31e504b8c5188cba6ee65))
* sync OpenAPI spec from cde0bc4 ([be891bc](https://github.com/team-telnyx/telnyx-go-staging/commit/be891bca7dc20ca48fbd8d6b2b200b43324c0d85))

## [4.87.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.86.0...v4.87.0) (2026-06-18)


### Chores

* preserve repo-owned files not part of SDK generation ([bce95aa](https://github.com/team-telnyx/telnyx-go-staging/commit/bce95aacee0b6ea56d7158d8fe30de3d4b90fc96))
* release go 4.87.0 ([74312c6](https://github.com/team-telnyx/telnyx-go-staging/commit/74312c6fe48a80b88cb70c29146fcc5a39501414))
* sync OpenAPI spec from 18f622e ([2ec3879](https://github.com/team-telnyx/telnyx-go-staging/commit/2ec3879aef13a95e96e0e4d09ee534951dc2ca4d))
* sync OpenAPI spec from 18f622e ([d6e9846](https://github.com/team-telnyx/telnyx-go-staging/commit/d6e9846007858ae48bb67c606589685523dc0c2e))

## [4.86.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.85.0...v4.86.0) (2026-06-18)


### Bug Fixes

* update transforms for inlined response schemas ([badf91e](https://github.com/team-telnyx/telnyx-go-staging/commit/badf91ef38fb52ac5e64092fff94be8edfaafa00))
* update transforms for inlined response schemas ([2a50c55](https://github.com/team-telnyx/telnyx-go-staging/commit/2a50c5556c175df70282d25ea431289e00a5e41c))


### Chores

* preserve repo-owned files not part of SDK generation ([4e832a2](https://github.com/team-telnyx/telnyx-go-staging/commit/4e832a29fb1105d8c03f90676fcde349f26202b5))
* release go 4.86.0 ([9b0f431](https://github.com/team-telnyx/telnyx-go-staging/commit/9b0f431e3cd688c4a7ea9e34e6f768f0c95b4bda))

## [4.85.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.84.0...v4.85.0) (2026-06-16)


### Chores

* preserve repo-owned files not part of SDK generation ([87613ad](https://github.com/team-telnyx/telnyx-go-staging/commit/87613ad7a327ad4fb1f66de5a4d5a2b41640c6f6))
* release go 4.85.0 ([c9cbebd](https://github.com/team-telnyx/telnyx-go-staging/commit/c9cbebda9a60a1b2541718d5977367de4ff37cdc))

## [4.84.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.83.0...v4.84.0) (2026-06-16)


### Chores

* preserve repo-owned files not part of SDK generation ([1ea7a69](https://github.com/team-telnyx/telnyx-go-staging/commit/1ea7a69fcbfdea47cc8493df41409de9985a6876))
* release go 4.84.0 ([b813d18](https://github.com/team-telnyx/telnyx-go-staging/commit/b813d18fd38caf47ef3b03e57866432ff8c9d9a7))

## [4.83.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.82.0...v4.83.0) (2026-06-16)


### Chores

* preserve repo-owned files not part of SDK generation ([09e0fdb](https://github.com/team-telnyx/telnyx-go-staging/commit/09e0fdbc0e4f25d93c84d9998fc9204cfef8a2ed))
* release go 4.83.0 ([aa9e507](https://github.com/team-telnyx/telnyx-go-staging/commit/aa9e50704a553d3723b2fb8e7ccb502fd1cdf820))
* sync OpenAPI spec from e44bbbc ([e4705b9](https://github.com/team-telnyx/telnyx-go-staging/commit/e4705b9639db3691828238aa87eaf10ca78f1916))
* sync OpenAPI spec from e44bbbc ([27ab863](https://github.com/team-telnyx/telnyx-go-staging/commit/27ab8638535c9e1ccf820a41a8ec910cbf4924d9))

## [4.82.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.81.0...v4.82.0) (2026-06-15)


### Bug Fixes

* rename DIR method names to match published SDKs ([20bb17c](https://github.com/team-telnyx/telnyx-go-staging/commit/20bb17cff0323a25489c09dcd761ca357227cf14))
* rename DIR method names to match published SDKs ([b1a8632](https://github.com/team-telnyx/telnyx-go-staging/commit/b1a8632d2a06302a34907c40be85b1baa091db48))


### Chores

* preserve repo-owned files not part of SDK generation ([51416c4](https://github.com/team-telnyx/telnyx-go-staging/commit/51416c41f4f4819ec8be3913772408e83750693e))
* release go 4.82.0 ([c1fc0b4](https://github.com/team-telnyx/telnyx-go-staging/commit/c1fc0b49c8f51b36ff06262a2fa91a44e410a90e))

## [4.81.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.80.0...v4.81.0) (2026-06-15)


### Chores

* preserve repo-owned files not part of SDK generation ([2e1553d](https://github.com/team-telnyx/telnyx-go-staging/commit/2e1553d8edd34ca4530f258691cefadee2770bb4))
* release go 4.81.0 ([b8bfa6f](https://github.com/team-telnyx/telnyx-go-staging/commit/b8bfa6f507923c2723fa0906217124ad2f21e664))
* sync OpenAPI spec from 3b5c722 ([ca6d1ba](https://github.com/team-telnyx/telnyx-go-staging/commit/ca6d1ba5471b0fc33b7feb6330300c24d2de7dc8))
* sync OpenAPI spec from 3b5c722 ([5479942](https://github.com/team-telnyx/telnyx-go-staging/commit/5479942e8229fa4d031d204e8f8751e47ef43c71))

## [4.80.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.79.0...v4.80.0) (2026-06-10)


### Bug Fixes

* add missing shimjson import to advancedorder.go ([a1e6d2d](https://github.com/team-telnyx/telnyx-go-staging/commit/a1e6d2d5d9162381ebebe1adf0a27541c6442e31))


### Chores

* preserve repo-owned files not part of SDK generation ([ebc9214](https://github.com/team-telnyx/telnyx-go-staging/commit/ebc921427b140a6b11141eaade5c4ad5cbe38811))
* release go 4.80.0 ([f8df6df](https://github.com/team-telnyx/telnyx-go-staging/commit/f8df6df276454f3fb472e747a86f7bd881f4408a))
* sync OpenAPI spec from b9f127e ([350fb05](https://github.com/team-telnyx/telnyx-go-staging/commit/350fb056a11b7a036349a4666bd3ab3d771b55ba))
* sync OpenAPI spec from c319cee ([5136e51](https://github.com/team-telnyx/telnyx-go-staging/commit/5136e517f31d902475aa921f06b0ed1d9e832767))

## [4.79.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.78.0...v4.79.0) (2026-06-08)


### Chores

* preserve repo-owned files not part of SDK generation ([63bbefe](https://github.com/team-telnyx/telnyx-go-staging/commit/63bbefef0979fa3e8440169fe87467d83c4e8781))
* release go 4.79.0 ([e9154fc](https://github.com/team-telnyx/telnyx-go-staging/commit/e9154fcdd1d0908d7c9c5453a2ccc7541486d47d))
* sync OpenAPI spec from aae7c19 ([9995b4f](https://github.com/team-telnyx/telnyx-go-staging/commit/9995b4fcc01d521427f00bb83c7f6a2bcb9560c1))
* sync OpenAPI spec from aae7c19 ([4cae21d](https://github.com/team-telnyx/telnyx-go-staging/commit/4cae21d803ef82ea37130df76a7ce9ae5b5534b1))

## [4.78.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.77.0...v4.78.0) (2026-06-08)


### Chores

* preserve repo-owned files not part of SDK generation ([9a20f45](https://github.com/team-telnyx/telnyx-go-staging/commit/9a20f45b2aa9a24662e92d46656e2fabb584661c))
* release go 4.78.0 ([95fff8b](https://github.com/team-telnyx/telnyx-go-staging/commit/95fff8bb863f011d4b5f82daaa6283efa3d13c14))
* sync OpenAPI spec from a13d4b1 ([ab962b4](https://github.com/team-telnyx/telnyx-go-staging/commit/ab962b49e1404bf68549bdbb3521a6c35a183130))
* sync OpenAPI spec from a13d4b1 ([bc16c15](https://github.com/team-telnyx/telnyx-go-staging/commit/bc16c154082e091c536551ca068d82326cf56251))

## [4.77.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.76.0...v4.77.0) (2026-06-08)


### Bug Fixes

* restore go.mod with gorilla/websocket dependency from main ([1a8bfe8](https://github.com/team-telnyx/telnyx-go-staging/commit/1a8bfe8713865d1f40de89d50b489926861eba5c))
* restore go.sum with gorilla/websocket from main ([7a42c26](https://github.com/team-telnyx/telnyx-go-staging/commit/7a42c26fbc450ae601d860e67910791ea0b572bb))


### Chores

* preserve repo-owned files not part of SDK generation ([3d5ee59](https://github.com/team-telnyx/telnyx-go-staging/commit/3d5ee5966e5a7aff0ff81e5c47f236f622132efe))
* release go 4.77.0 ([e84350f](https://github.com/team-telnyx/telnyx-go-staging/commit/e84350f39cda10888c13a1e0102e227b81b942e7))
* sync OpenAPI spec from 9f5f345 ([4edc2e1](https://github.com/team-telnyx/telnyx-go-staging/commit/4edc2e1fd51b9b88b0902b5b73fe9acd85ea9618))

## [4.76.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.75.1...v4.76.0) (2026-06-08)


### Features

* add examples/speech-to-text-websocket/main.go ([685b2b5](https://github.com/team-telnyx/telnyx-go-staging/commit/685b2b5395d949e5e720dc772db412de9d8420d6))
* add examples/text-to-speech-websocket/main.go ([e5bf0a8](https://github.com/team-telnyx/telnyx-go-staging/commit/e5bf0a8c9ba42758418f2f7654ff24d27d8f5f0d))
* add gorilla/websocket dependency for custom lib/ code ([30bd3f5](https://github.com/team-telnyx/telnyx-go-staging/commit/30bd3f5334de019bc69d0f3af875da9f7e966725))
* add lib/speech_to_text_ws_test.go ([6f7df4d](https://github.com/team-telnyx/telnyx-go-staging/commit/6f7df4d647973b2924d7feb0d0e26a7635c1fe22))
* add lib/speech_to_text_ws.go ([746dd9c](https://github.com/team-telnyx/telnyx-go-staging/commit/746dd9c24cfaaab660b696d845628a8eacd615b3))
* add lib/text_to_speech_ws_test.go ([1264e75](https://github.com/team-telnyx/telnyx-go-staging/commit/1264e75e261dbc6771d213bbf7330612ae2123ed))
* add lib/text_to_speech_ws.go ([d3316fa](https://github.com/team-telnyx/telnyx-go-staging/commit/d3316fac984369cf941c172a2408d497008be28a))
* add lib/webhook_verification_test.go ([25ca8fc](https://github.com/team-telnyx/telnyx-go-staging/commit/25ca8fc50cc10cbf6661995e04972c05587c05e5))
* add lib/webhook_verification.go ([c40a9e1](https://github.com/team-telnyx/telnyx-go-staging/commit/c40a9e1ed1c1bfec43afddda9185f4a9a9d7ad09))
* add lib/websocket.go ([04b0925](https://github.com/team-telnyx/telnyx-go-staging/commit/04b09252c8f281e2bede25cb62a7c208a6d128f2))
* add webhook_custom.go for signature verification ([a56c74d](https://github.com/team-telnyx/telnyx-go-staging/commit/a56c74dd8ccd84a9177eff5fe6e3b26d67c01b5f))
* add webhook_custom.go for signature verification ([0c46445](https://github.com/team-telnyx/telnyx-go-staging/commit/0c464458b2877e178e4ac5af01a1142eb8dc53a5))
* update go.sum with gorilla/websocket checksums ([0fd3328](https://github.com/team-telnyx/telnyx-go-staging/commit/0fd3328db75fe7a4561cf3c8263db69e55200bef))


### Bug Fixes

* **ci:** pass release-please pr output via env to avoid shell injection ([a0d9ea0](https://github.com/team-telnyx/telnyx-go-staging/commit/a0d9ea02e5c2533e923134abcafd5722368d3d6c))
* **ci:** pass release-please pr output via env to avoid shell injection ([643db92](https://github.com/team-telnyx/telnyx-go-staging/commit/643db92418836146b864869e4846df499ca918b7))
* remove duplicate NewWebhookService from webhook_custom.go ([e6dde46](https://github.com/team-telnyx/telnyx-go-staging/commit/e6dde46ba1e748d7c7e11b6c8a4a6490315010d4))


### Chores

* preserve repo-owned files not part of SDK generation ([f20ed66](https://github.com/team-telnyx/telnyx-go-staging/commit/f20ed6619430d61061e759fb7df4bb9148da427d))
* release go 4.76.0 ([aef2bb3](https://github.com/team-telnyx/telnyx-go-staging/commit/aef2bb3229b6b256750b1bece7a52c831fe7e49c))
* sync OpenAPI spec from 6eae6a5 ([0a108cc](https://github.com/team-telnyx/telnyx-go-staging/commit/0a108cc9b745d6092d6ddd1367c1b3705efa5c63))
* sync OpenAPI spec from 6eae6a5 ([0c73d6c](https://github.com/team-telnyx/telnyx-go-staging/commit/0c73d6c3ace5e6d515abe116c3619ba78834abf6))

## [4.75.1](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.75.0...v4.75.1) (2026-06-07)


### Bug Fixes

* extract PR number from JSON output for auto-merge ([ac07fa4](https://github.com/team-telnyx/telnyx-go-staging/commit/ac07fa43af0b1939fd3bce93fddda8da01f3c9b2))
* extract PR number from JSON output for auto-merge ([4755930](https://github.com/team-telnyx/telnyx-go-staging/commit/4755930fd1fd31fec28ee654f3c57448a6e2d588))
* use release-please output directly for auto-merge (avoid race condition) ([6366967](https://github.com/team-telnyx/telnyx-go-staging/commit/6366967824f463d183837278a3bde76791e72ad5))
* use release-please output directly for auto-merge (avoid race condition) ([5b34b8b](https://github.com/team-telnyx/telnyx-go-staging/commit/5b34b8b8105da1ee28daba9d4f6f5d5674789f54))

## [4.75.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.74.0...v4.75.0) (2026-06-07)


### Chores

* preserve repo-owned files not part of SDK generation ([e4149d0](https://github.com/team-telnyx/telnyx-go-staging/commit/e4149d06d20d4100cc0068a49ee6f2d58aafd9f9))
* preserve repo-owned files not part of SDK generation ([bf735ed](https://github.com/team-telnyx/telnyx-go-staging/commit/bf735ede574a9035eaf8e0fe398c34101fa80ec2))
* release go 4.75.0 ([22a14b5](https://github.com/team-telnyx/telnyx-go-staging/commit/22a14b558704b12e2157289abb3a796c286ca7fe))
* release go 4.75.0 ([c8ceb40](https://github.com/team-telnyx/telnyx-go-staging/commit/c8ceb4094423cc03144d2e967cb089adf78cc8e1))
* sync OpenAPI spec from 0193002 ([8d6fb63](https://github.com/team-telnyx/telnyx-go-staging/commit/8d6fb637015c9e2377e7e259fe20040758ff5827))
* sync OpenAPI spec from 0193002 ([ac30f37](https://github.com/team-telnyx/telnyx-go-staging/commit/ac30f37ff1ef81a72bbe0e12af079a68402ac888))
* sync OpenAPI spec from 8faa4be ([ccb540e](https://github.com/team-telnyx/telnyx-go-staging/commit/ccb540ec9937709fe4bb9db0bb5ffeaf3ef37614))
* sync OpenAPI spec from 8faa4be ([8376ab3](https://github.com/team-telnyx/telnyx-go-staging/commit/8376ab3c23a8c99b1949faacdec691533d4a0afa))

## [4.74.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.73.0...v4.74.0) (2026-06-07)


### Features

* enable GitHub auto-merge on release PRs ([216ef15](https://github.com/team-telnyx/telnyx-go-staging/commit/216ef15d1d0cd888e7c0570ac719154ae5f0072e))
* enable GitHub auto-merge on release PRs ([898b901](https://github.com/team-telnyx/telnyx-go-staging/commit/898b901eed0d0d07b669b7e8989fb52352682cfc))


### Bug Fixes

* correct auto-merge output name and PR search pattern ([f5fc9f4](https://github.com/team-telnyx/telnyx-go-staging/commit/f5fc9f43081700de8464a0d4c38b2503075c6c22))
* correct auto-merge output name and PR search pattern ([d9daa5c](https://github.com/team-telnyx/telnyx-go-staging/commit/d9daa5c355189b1e555af27f7db4401fba25ce44))
* restore ${{ }} expressions in release-please workflow ([b575dfc](https://github.com/team-telnyx/telnyx-go-staging/commit/b575dfc9618f1006ba64434b78a583604698bad2))
* restore ${{ }} expressions in release-please workflow ([eddb732](https://github.com/team-telnyx/telnyx-go-staging/commit/eddb7323f819dbf06dc2305c872763b4263a85cb))

## [4.73.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.72.0...v4.73.0) (2026-06-07)


### Features

* enable automerge for release PRs ([9292127](https://github.com/team-telnyx/telnyx-go-staging/commit/92921271d8bbd1f5f22206a63a259fbd701fc4d0))
* enable automerge for release PRs ([70e8a7c](https://github.com/team-telnyx/telnyx-go-staging/commit/70e8a7cb53d41f7b1457737f18cf33369a33af38))

## [4.72.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.71.0...v4.72.0) (2026-06-07)


### Bug Fixes

* remove text_to_speech retrieve_speech mapping ([b047f24](https://github.com/team-telnyx/telnyx-go-staging/commit/b047f24b2ff1727f19bba78f26cee5b2d5ffda0c))
* use PAT for release-please to trigger CI ([f0a4ae6](https://github.com/team-telnyx/telnyx-go-staging/commit/f0a4ae61714bca8ef9b4d2483d6791176c203f8a))
* use PAT for release-please to trigger CI ([6adc265](https://github.com/team-telnyx/telnyx-go-staging/commit/6adc2652f2d4e5a10e22948dcc1cfe27c335cb4a))


### Chores

* preserve repo-owned files not part of SDK generation ([9b4f223](https://github.com/team-telnyx/telnyx-go-staging/commit/9b4f22363bc8cc505aa6e3fc16b91a4ece80f6cc))
* preserve repo-owned files not part of SDK generation ([0d8eaee](https://github.com/team-telnyx/telnyx-go-staging/commit/0d8eaee22cb8c503dab223a2460745588f9c7cff))
* release go 4.72.0 ([26ea458](https://github.com/team-telnyx/telnyx-go-staging/commit/26ea45820c5c8b5368e48101a3dbfcaa544e1b09))
* release go 4.72.0 ([24339d1](https://github.com/team-telnyx/telnyx-go-staging/commit/24339d12372f0b299432e2596a9dda22d95f3cba))
* sync OpenAPI spec from c46b2c0 ([df29451](https://github.com/team-telnyx/telnyx-go-staging/commit/df29451b1c1b7301e740fa32aba20f296fdbb162))

## [4.71.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.70.1...v4.71.0) (2026-06-06)


### Bug Fixes

* correct YAML quoting in CI if condition ([f926ab4](https://github.com/team-telnyx/telnyx-go-staging/commit/f926ab47a621704dff74c834482eca5e5c024986))
* **go:** use model mapping for DirPhoneNumberStatus dedup ([4fb7464](https://github.com/team-telnyx/telnyx-go-staging/commit/4fb7464f4bb17046f13086506cd38b0b8ed26ffc))
* **go:** use model mapping for DirPhoneNumberStatus dedup ([a496e11](https://github.com/team-telnyx/telnyx-go-staging/commit/a496e11c44fa690c0eeb3b008b3bf0dbe1b60854))


### Chores

* preserve repo-owned files not part of SDK generation ([f8a2c88](https://github.com/team-telnyx/telnyx-go-staging/commit/f8a2c88e33260cdfa806b9b4fcaf61395534cdc1))
* release go 4.71.0 ([eee9583](https://github.com/team-telnyx/telnyx-go-staging/commit/eee95834717fe65217e2398247b141599ad5e4d3))

## [4.70.1](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.70.0...v4.70.1) (2026-06-05)


### Bug Fixes

* run CI build/lint on internal PRs not just forks ([f601f82](https://github.com/team-telnyx/telnyx-go-staging/commit/f601f8274af9ce59d136e11f453a80a5ddf8c3c9))

## [4.70.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.69.0...v4.70.0) (2026-06-04)


### Bug Fixes

* **ai:** name expression op enum members (CannotInferEnumMemberName) ([efd8f3f](https://github.com/team-telnyx/telnyx-go-staging/commit/efd8f3f019e80d1f1a2e830f623a8c9677f3e4c2))


### Chores

* preserve repo-owned files (release-please runner, CHANGELOG) ([c7bcc01](https://github.com/team-telnyx/telnyx-go-staging/commit/c7bcc0172fc2db6fde01b200099104e84fbd885f))
* preserve repo-owned files not part of SDK generation ([b72a74b](https://github.com/team-telnyx/telnyx-go-staging/commit/b72a74b9aec40984ee5c4a03afc7bc0376909bd1))
* preserve repo-owned files not part of SDK generation ([0ac9d09](https://github.com/team-telnyx/telnyx-go-staging/commit/0ac9d096f31c7eb9eb594f3773169a9e3bfda068))
* preserve repo-owned files not part of SDK generation ([29601b1](https://github.com/team-telnyx/telnyx-go-staging/commit/29601b1a7242f6b42fc911b234901763e297d202))
* release go 4.70.0 ([3643748](https://github.com/team-telnyx/telnyx-go-staging/commit/36437488f2ebc2a667a2ea2a589b755ca2a359bc))
* release go 4.70.0 ([c85f1bb](https://github.com/team-telnyx/telnyx-go-staging/commit/c85f1bb3df720e428138f07e479a2658e9fcf09c))
* release go 4.70.0 ([b9877ab](https://github.com/team-telnyx/telnyx-go-staging/commit/b9877ab137ae160eab194b2d3fb12edbaa5bb179))
* sync OpenAPI spec from 2c28e93 ([3045e81](https://github.com/team-telnyx/telnyx-go-staging/commit/3045e81f783161258b25353119c917ab5d4368cf))
* sync OpenAPI spec from 404ef48 ([e6b21fe](https://github.com/team-telnyx/telnyx-go-staging/commit/e6b21fe7252108a5bcd17642616a81df4e252ad7))

## [4.69.0](https://github.com/team-telnyx/telnyx-go-staging/compare/v4.68.0...v4.69.0) (2026-06-03)


### Chores

* release go 4.69.0 ([4ba2b75](https://github.com/team-telnyx/telnyx-go-staging/commit/4ba2b757bd5447dd6045cbe247169b2328379d82))

# Details

Date : 2024-05-08 14:04:42

Directory /Users/carterrath/Documents/Fall2023/SE490/ServerlessOrchestrator

Total : 134 files,  14215 codes, 438 comments, 816 blanks, all 15469 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [.github/workflows/deployApplication.yml](/.github/workflows/deployApplication.yml) | YAML | 62 | 0 | 9 | 71 |
| [.github/workflows/tests.yml](/.github/workflows/tests.yml) | YAML | 42 | 2 | 9 | 53 |
| [.golangci.yml](/.golangci.yml) | YAML | 23 | 1 | 3 | 27 |
| [.prettierrc](/.prettierrc) | JSON | 5 | 0 | 1 | 6 |
| [README.md](/README.md) | Markdown | 27 | 0 | 12 | 39 |
| [application/api.go](/application/api.go) | Go | 76 | 31 | 12 | 119 |
| [application/dockerhub/dockerhubAPI.go](/application/dockerhub/dockerhubAPI.go) | Go | 79 | 13 | 25 | 117 |
| [application/elasticcontainerservice/elasticcontainerserviceAPI.go](/application/elasticcontainerservice/elasticcontainerserviceAPI.go) | Go | 77 | 0 | 6 | 83 |
| [application/github/githubAPI.go](/application/github/githubAPI.go) | Go | 33 | 5 | 11 | 49 |
| [application/microholder/README.md](/application/microholder/README.md) | Markdown | 1 | 0 | 1 | 2 |
| [application/routes/microservice/getAllMicroservices.go](/application/routes/microservice/getAllMicroservices.go) | Go | 14 | 0 | 5 | 19 |
| [application/routes/microservice/uploadMicroservice.go](/application/routes/microservice/uploadMicroservice.go) | Go | 56 | 9 | 15 | 80 |
| [application/routes/runmicroservice/runMicroservice.go](/application/routes/runmicroservice/runMicroservice.go) | Go | 25 | 6 | 8 | 39 |
| [application/routes/stopmicroservice/stopMicroservice.go](/application/routes/stopmicroservice/stopMicroservice.go) | Go | 23 | 5 | 10 | 38 |
| [application/routes/users/authentication.go](/application/routes/users/authentication.go) | Go | 54 | 11 | 11 | 76 |
| [application/routes/users/createAccount.go](/application/routes/users/createAccount.go) | Go | 71 | 8 | 16 | 95 |
| [application/routes/users/getUserDetails.go](/application/routes/users/getUserDetails.go) | Go | 116 | 13 | 17 | 146 |
| [application/routes/users/loginAccount.go](/application/routes/users/loginAccount.go) | Go | 41 | 9 | 11 | 61 |
| [application/routes/users/recovery.go](/application/routes/users/recovery.go) | Go | 78 | 18 | 21 | 117 |
| [application/routes/users/reset.go](/application/routes/users/reset.go) | Go | 28 | 3 | 8 | 39 |
| [application/services/executeService.go](/application/services/executeService.go) | Go | 124 | 18 | 22 | 164 |
| [application/services/stopService.go](/application/services/stopService.go) | Go | 31 | 4 | 9 | 44 |
| [application/services/uploadService.go](/application/services/uploadService.go) | Go | 164 | 38 | 45 | 247 |
| [business/BuildScript.go](/business/BuildScript.go) | Go | 3 | 1 | 2 | 6 |
| [business/DAO_IF.go](/business/DAO_IF.go) | Go | 6 | 4 | 4 | 14 |
| [business/DockerBuild.go](/business/DockerBuild.go) | Go | 1 | 0 | 1 | 2 |
| [business/Input.go](/business/Input.go) | Go | 7 | 0 | 2 | 9 |
| [business/Microservice.go](/business/Microservice.go) | Go | 16 | 1 | 3 | 20 |
| [business/User.go](/business/User.go) | Go | 11 | 1 | 3 | 15 |
| [dataaccess/dbSetup.go](/dataaccess/dbSetup.go) | Go | 47 | 16 | 16 | 79 |
| [dataaccess/microservicesDAO.go](/dataaccess/microservicesDAO.go) | Go | 151 | 16 | 29 | 196 |
| [dataaccess/userDAO.go](/dataaccess/userDAO.go) | Go | 70 | 13 | 16 | 99 |
| [go.mod](/go.mod) | Go Module File | 84 | 0 | 5 | 89 |
| [go.sum](/go.sum) | Go Checksum File | 264 | 0 | 1 | 265 |
| [main.go](/main.go) | Go | 31 | 4 | 10 | 45 |
| [package-lock.json](/package-lock.json) | JSON | 1,613 | 0 | 1 | 1,614 |
| [package.json](/package.json) | JSON | 21 | 0 | 1 | 22 |
| [presentation/.eslintrc.cjs](/presentation/.eslintrc.cjs) | JavaScript | 24 | 0 | 1 | 25 |
| [presentation/README.md](/presentation/README.md) | Markdown | 22 | 0 | 9 | 31 |
| [presentation/index.html](/presentation/index.html) | HTML | 13 | 0 | 1 | 14 |
| [presentation/package-lock.json](/presentation/package-lock.json) | JSON | 6,887 | 0 | 1 | 6,888 |
| [presentation/package.json](/presentation/package.json) | JSON | 45 | 0 | 1 | 46 |
| [presentation/postcss.config.js](/presentation/postcss.config.js) | JavaScript | 6 | 0 | 1 | 7 |
| [presentation/src/Main.tsx](/presentation/src/Main.tsx) | TypeScript JSX | 8 | 0 | 2 | 10 |
| [presentation/src/assets/styles/styles.css](/presentation/src/assets/styles/styles.css) | CSS | 3 | 0 | 1 | 4 |
| [presentation/src/assets/svg/box.svg](/presentation/src/assets/svg/box.svg) | SVG | 15 | 0 | 0 | 15 |
| [presentation/src/assets/svg/builder.svg](/presentation/src/assets/svg/builder.svg) | SVG | 91 | 0 | 0 | 91 |
| [presentation/src/assets/svg/cart.svg](/presentation/src/assets/svg/cart.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/close.svg](/presentation/src/assets/svg/close.svg) | SVG | 10 | 0 | 0 | 10 |
| [presentation/src/assets/svg/code.svg](/presentation/src/assets/svg/code.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/docker-logo-white.svg](/presentation/src/assets/svg/docker-logo-white.svg) | SVG | 10 | 0 | 0 | 10 |
| [presentation/src/assets/svg/filter.svg](/presentation/src/assets/svg/filter.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/github-black.svg](/presentation/src/assets/svg/github-black.svg) | SVG | 18 | 0 | 0 | 18 |
| [presentation/src/assets/svg/github-white.svg](/presentation/src/assets/svg/github-white.svg) | SVG | 19 | 0 | 0 | 19 |
| [presentation/src/assets/svg/go-logo-white.svg](/presentation/src/assets/svg/go-logo-white.svg) | SVG | 12 | 0 | 0 | 12 |
| [presentation/src/assets/svg/kubernetes-logo-white.svg](/presentation/src/assets/svg/kubernetes-logo-white.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/linkedIn.svg](/presentation/src/assets/svg/linkedIn.svg) | SVG | 18 | 0 | 0 | 18 |
| [presentation/src/assets/svg/loading-black.svg](/presentation/src/assets/svg/loading-black.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/loading-white.svg](/presentation/src/assets/svg/loading-white.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/login.svg](/presentation/src/assets/svg/login.svg) | SVG | 14 | 0 | 0 | 14 |
| [presentation/src/assets/svg/minus.svg](/presentation/src/assets/svg/minus.svg) | SVG | 8 | 0 | 0 | 8 |
| [presentation/src/assets/svg/network.svg](/presentation/src/assets/svg/network.svg) | SVG | 10 | 0 | 0 | 10 |
| [presentation/src/assets/svg/orca-black.svg](/presentation/src/assets/svg/orca-black.svg) | SVG | 26 | 0 | 0 | 26 |
| [presentation/src/assets/svg/orca-white.svg](/presentation/src/assets/svg/orca-white.svg) | SVG | 26 | 0 | 0 | 26 |
| [presentation/src/assets/svg/output.svg](/presentation/src/assets/svg/output.svg) | SVG | 17 | 0 | 0 | 17 |
| [presentation/src/assets/svg/play.svg](/presentation/src/assets/svg/play.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/plus.svg](/presentation/src/assets/svg/plus.svg) | SVG | 8 | 0 | 0 | 8 |
| [presentation/src/assets/svg/react-logo-white.svg](/presentation/src/assets/svg/react-logo-white.svg) | SVG | 11 | 0 | 0 | 11 |
| [presentation/src/assets/svg/react.svg](/presentation/src/assets/svg/react.svg) | SVG | 6 | 0 | 0 | 6 |
| [presentation/src/assets/svg/scale.svg](/presentation/src/assets/svg/scale.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/search.svg](/presentation/src/assets/svg/search.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/shield.svg](/presentation/src/assets/svg/shield.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/shopping-bag.svg](/presentation/src/assets/svg/shopping-bag.svg) | SVG | 15 | 0 | 0 | 15 |
| [presentation/src/assets/svg/stack.svg](/presentation/src/assets/svg/stack.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/stop.svg](/presentation/src/assets/svg/stop.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/tailwind-logo-white.svg](/presentation/src/assets/svg/tailwind-logo-white.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/tools-black.svg](/presentation/src/assets/svg/tools-black.svg) | SVG | 27 | 0 | 0 | 27 |
| [presentation/src/assets/svg/tools-white.svg](/presentation/src/assets/svg/tools-white.svg) | SVG | 27 | 0 | 0 | 27 |
| [presentation/src/assets/svg/typescript-logo-white.svg](/presentation/src/assets/svg/typescript-logo-white.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/update.svg](/presentation/src/assets/svg/update.svg) | SVG | 9 | 0 | 0 | 9 |
| [presentation/src/assets/svg/upload.svg](/presentation/src/assets/svg/upload.svg) | SVG | 12 | 0 | 0 | 12 |
| [presentation/src/assets/svg/vite.svg](/presentation/src/assets/svg/vite.svg) | SVG | 20 | 0 | 0 | 20 |
| [presentation/src/assets/svg/worker-man.svg](/presentation/src/assets/svg/worker-man.svg) | SVG | 106 | 0 | 14 | 120 |
| [presentation/src/components/AboutSection/index.tsx](/presentation/src/components/AboutSection/index.tsx) | TypeScript JSX | 115 | 0 | 3 | 118 |
| [presentation/src/components/BackgroundGradient/index.tsx](/presentation/src/components/BackgroundGradient/index.tsx) | TypeScript JSX | 13 | 0 | 2 | 15 |
| [presentation/src/components/BackgroundImage/index.tsx](/presentation/src/components/BackgroundImage/index.tsx) | TypeScript JSX | 18 | 0 | 3 | 21 |
| [presentation/src/components/DialogueMessage/index.tsx](/presentation/src/components/DialogueMessage/index.tsx) | TypeScript JSX | 54 | 1 | 9 | 64 |
| [presentation/src/components/FeaturesSection/index.tsx](/presentation/src/components/FeaturesSection/index.tsx) | TypeScript JSX | 73 | 1 | 4 | 78 |
| [presentation/src/components/HeaderSection/index.tsx](/presentation/src/components/HeaderSection/index.tsx) | TypeScript JSX | 45 | 0 | 3 | 48 |
| [presentation/src/components/Loading/index.tsx](/presentation/src/components/Loading/index.tsx) | TypeScript JSX | 15 | 0 | 2 | 17 |
| [presentation/src/components/NavBar/index.tsx](/presentation/src/components/NavBar/index.tsx) | TypeScript JSX | 64 | 0 | 6 | 70 |
| [presentation/src/components/OtherMaterialsSection/index.tsx](/presentation/src/components/OtherMaterialsSection/index.tsx) | TypeScript JSX | 34 | 0 | 2 | 36 |
| [presentation/src/components/ScrollAssistant/index.tsx](/presentation/src/components/ScrollAssistant/index.tsx) | TypeScript JSX | 36 | 12 | 10 | 58 |
| [presentation/src/components/TeamSection/index.tsx](/presentation/src/components/TeamSection/index.tsx) | TypeScript JSX | 84 | 0 | 3 | 87 |
| [presentation/src/components/TechStackSection/index.tsx](/presentation/src/components/TechStackSection/index.tsx) | TypeScript JSX | 64 | 0 | 3 | 67 |
| [presentation/src/constants/index.ts](/presentation/src/constants/index.ts) | TypeScript | 1 | 0 | 1 | 2 |
| [presentation/src/hooks/useAuth.tsx](/presentation/src/hooks/useAuth.tsx) | TypeScript JSX | 125 | 6 | 21 | 152 |
| [presentation/src/index.tsx](/presentation/src/index.tsx) | TypeScript JSX | 10 | 1 | 3 | 14 |
| [presentation/src/pages/DeveloperOptions/index.tsx](/presentation/src/pages/DeveloperOptions/index.tsx) | TypeScript JSX | 40 | 2 | 5 | 47 |
| [presentation/src/pages/GetStarted/ConsumerLogin.tsx](/presentation/src/pages/GetStarted/ConsumerLogin.tsx) | TypeScript JSX | 105 | 0 | 13 | 118 |
| [presentation/src/pages/GetStarted/ConsumerSignup.tsx](/presentation/src/pages/GetStarted/ConsumerSignup.tsx) | TypeScript JSX | 114 | 1 | 12 | 127 |
| [presentation/src/pages/GetStarted/DeveloperLogin.tsx](/presentation/src/pages/GetStarted/DeveloperLogin.tsx) | TypeScript JSX | 105 | 0 | 12 | 117 |
| [presentation/src/pages/GetStarted/DeveloperSignup.tsx](/presentation/src/pages/GetStarted/DeveloperSignup.tsx) | TypeScript JSX | 114 | 1 | 12 | 127 |
| [presentation/src/pages/GetStarted/Recover.tsx](/presentation/src/pages/GetStarted/Recover.tsx) | TypeScript JSX | 122 | 0 | 14 | 136 |
| [presentation/src/pages/GetStarted/ResetPassword.tsx](/presentation/src/pages/GetStarted/ResetPassword.tsx) | TypeScript JSX | 84 | 4 | 10 | 98 |
| [presentation/src/pages/GetStarted/index.tsx](/presentation/src/pages/GetStarted/index.tsx) | TypeScript JSX | 60 | 3 | 6 | 69 |
| [presentation/src/pages/Home/index.tsx](/presentation/src/pages/Home/index.tsx) | TypeScript JSX | 30 | 0 | 4 | 34 |
| [presentation/src/pages/Microservices/MicroserviceCard.tsx](/presentation/src/pages/Microservices/MicroserviceCard.tsx) | TypeScript JSX | 138 | 2 | 9 | 149 |
| [presentation/src/pages/Microservices/MicroserviceCards.tsx](/presentation/src/pages/Microservices/MicroserviceCards.tsx) | TypeScript JSX | 35 | 4 | 5 | 44 |
| [presentation/src/pages/Microservices/index.tsx](/presentation/src/pages/Microservices/index.tsx) | TypeScript JSX | 78 | 0 | 7 | 85 |
| [presentation/src/pages/UploadMicroservice/Input.tsx](/presentation/src/pages/UploadMicroservice/Input.tsx) | TypeScript JSX | 53 | 0 | 3 | 56 |
| [presentation/src/pages/UploadMicroservice/index.tsx](/presentation/src/pages/UploadMicroservice/index.tsx) | TypeScript JSX | 216 | 10 | 16 | 242 |
| [presentation/src/router/index.tsx](/presentation/src/router/index.tsx) | TypeScript JSX | 38 | 1 | 1 | 40 |
| [presentation/src/types/input.d.ts](/presentation/src/types/input.d.ts) | TypeScript | 7 | 0 | 2 | 9 |
| [presentation/src/types/microservice-data.d.ts](/presentation/src/types/microservice-data.d.ts) | TypeScript | 16 | 0 | 2 | 18 |
| [presentation/src/types/microservice-upload.d.ts](/presentation/src/types/microservice-upload.d.ts) | TypeScript | 6 | 0 | 1 | 7 |
| [presentation/src/types/password-reset.d.ts](/presentation/src/types/password-reset.d.ts) | TypeScript | 4 | 0 | 1 | 5 |
| [presentation/src/types/recovery-data.d.ts](/presentation/src/types/recovery-data.d.ts) | TypeScript | 4 | 0 | 1 | 5 |
| [presentation/src/types/type.d.ts](/presentation/src/types/type.d.ts) | TypeScript | 1 | 0 | 1 | 2 |
| [presentation/src/types/user-data.d.ts](/presentation/src/types/user-data.d.ts) | TypeScript | 9 | 0 | 1 | 10 |
| [presentation/src/types/user-upload.d.ts](/presentation/src/types/user-upload.d.ts) | TypeScript | 6 | 0 | 1 | 7 |
| [presentation/src/vite-env.d.ts](/presentation/src/vite-env.d.ts) | TypeScript | 0 | 1 | 1 | 2 |
| [presentation/tailwind.config.js](/presentation/tailwind.config.js) | JavaScript | 31 | 1 | 1 | 33 |
| [presentation/tsconfig.json](/presentation/tsconfig.json) | JSON with Comments | 21 | 2 | 3 | 26 |
| [presentation/tsconfig.node.json](/presentation/tsconfig.node.json) | JSON | 10 | 0 | 1 | 11 |
| [presentation/vite.config.ts](/presentation/vite.config.ts) | TypeScript | 5 | 1 | 2 | 8 |
| [testing/applicationtests/dockerhubAPI_test.go](/testing/applicationtests/dockerhubAPI_test.go) | Go | 55 | 22 | 13 | 90 |
| [testing/applicationtests/executeService_test.go](/testing/applicationtests/executeService_test.go) | Go | 87 | 13 | 21 | 121 |
| [testing/applicationtests/githubAPI_test.go](/testing/applicationtests/githubAPI_test.go) | Go | 68 | 10 | 19 | 97 |
| [testing/applicationtests/stopService_test.go](/testing/applicationtests/stopService_test.go) | Go | 57 | 11 | 15 | 83 |
| [testing/applicationtests/uploadService_test.go](/testing/applicationtests/uploadService_test.go) | Go | 237 | 26 | 47 | 310 |
| [testing/dataaccesstests/dbSetup_test.go](/testing/dataaccesstests/dbSetup_test.go) | Go | 49 | 8 | 14 | 71 |
| [testing/dataaccesstests/microservicesDAO_test.go](/testing/dataaccesstests/microservicesDAO_test.go) | Go | 124 | 28 | 29 | 181 |
| [testing/dataaccesstests/userDAO_test.go](/testing/dataaccesstests/userDAO_test.go) | Go | 91 | 16 | 24 | 131 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)
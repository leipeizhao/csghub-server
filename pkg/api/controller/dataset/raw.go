package dataset

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) FileRaw(ctx *gin.Context) (string, error) {
	markdownContent := "# Gitness\nGitness is an open source development platform packed with the power of code hosting and automated DevOps pipelines.\n\n## Overview\nGitness is an open source development platform packed with the power of code hosting and automated continuous integration pipelines.\n\n## Running Gitness locally\n> The latest publicly released docker image can be found on [harness/gitness](https://hub.docker.com/r/harness/gitness).\n\nTo install Gitness yourself, simply run the command below. Once the container is up, you can visit http://localhost:3000 in your browser.\n\n```bash\ndocker run -d \\\n  -p 3000:3000 \\\n  -v /var/run/docker.sock:/var/run/docker.sock \\\n  -v /tmp/gitness:/data \\\n  --name gitness \\\n  --restart always \\\n  harness/gitness\n```\n> The Gitness image uses a volume to store the database and repositories. It is highly recommended to use a bind mount or named volume as otherwise all data will be lost once the container is stopped.\n\nSee [docs.gitness.com](https://docs.gitness.com) to learn how to get the most out of Gitness.\n\n## Where is Drone?\n\nGitness represents a massive investment in the next generation of Drone. Where Drone focused on continuous integration, Gitness adds source code hosting, bringing code management and pipelines closer together.\n\nThe goal is for Gitness to eventually be at full parity with Drone in terms of pipeline capabilities, allowing users to seemlessly migrate from Drone to Gitness.\n\nBut, we expect this to take some time, which is why we took a snapshot of Drone as a feature branch [drone](https://github.com/harness/gitness/tree/drone) ([README](https://github.com/harness/gitness/blob/drone/.github/readme.md)) so it can continue development.\n\nAs for Gitness, the development is taking place on the [main](https://github.com/harness/gitness/tree/main) branch.\n\nFor more information on Gitness, please visit [gitness.com](https://gitness.com/).\n\nFor more information on Drone, please visit [drone.io](https://www.drone.io/).\n\n## Gitness Development\n### Pre-Requisites\n\nInstall the latest stable version of Node and Go version 1.19 or higher, and then install the below Go programs. Ensure the GOPATH [bin directory](https://go.dev/doc/gopath_code#GOPATH) is added to your PATH.\n\nInstall protobuf\n- Check if you've already installed protobuf ```protoc --version```\n- If your version is different than v3.21.11, run ```brew unlink protobuf```\n- Get v3.21.11 ```curl -s https://raw.githubusercontent.com/Homebrew/homebrew-core/9de8de7a533609ebfded833480c1f7c05a3448cb/Formula/protobuf.rb > /tmp/protobuf.rb```\n- Install it ```brew install /tmp/protobuf.rb```\n- Check out your version ```protoc --version```\n\nInstall protoc-gen-go and protoc-gen-go-rpc:\n\n- Install protoc-gen-go v1.28.1 ```go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1```\n(Note that this will install a binary in $GOBIN so make sure $GOBIN is in your $PATH)\n\n- Install protoc-gen-go-grpc v1.2.0 ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0```\n\n```bash\n$ make dep\n$ make tools\n```\n\n### Build\n\nFirst step is to build the user interface artifacts:\n\n```bash\n$ pushd web\n$ yarn install\n$ yarn build\n$ popd\n```\n\nAfter that, you can build the gitness binary:\n\n```bash\n$ make build\n```\n\n### Run\n\nThis project supports all operating systems and architectures supported by Go.  This means you can build and run the system on your machine; docker containers are not required for local development and testing.\n\nTo start the server at `localhost:3000`, simply run the following command:\n\n```bash\n./gitness server .local.env\n```\n\n### Auto-Generate Gitness API Client used by UI using Swagger\nPlease make sure to update the autogenerated client code used by the UI when adding new rest APIs.\n\nTo regenerate the code, please execute the following steps:\n- Run local gitness instance with latest changes\n- Get latest OpenAPI specs from `http://localhost:3000/openapi.yaml` and store it in `web/src/services/code/swagger.yaml`\n- navigate into the `web` folder and run `yarn services`\n\nThe latest API changes should now be reflected in `web/src/services/code/index.tsx`\n\n\n## User Interface\n\nThis project includes a full user interface for interacting with the system. When you run the application, you can access the user interface by navigating to `http://localhost:3000` in your browser.\n\n## REST API\n\nThis project includes a swagger specification. When you run the application, you can access the swagger specification by navigating to `http://localhost:3000/swagger` in your browser (for raw yaml see `http://localhost:3000/openapi.yaml`).\n\n\nFor testing, it's simplest to just use the cli to create a token (this requires gitness server to run):\n```bash\n# LOGIN (user: admin, pw: changeit)\n$ ./gitness login\n\n# GENERATE PAT (1 YEAR VALIDITY)\n$ ./gitness user pat \"my-pat-uid\" 2592000\n```\n\nThe command outputs a valid PAT that has been granted full access as the user.\nThe token can then be send as part of the `Authorization` header with Postman or curl:\n\n```bash\n$ curl http://localhost:3000/api/v1/user \\\n-H \"Authorization: Bearer $TOKEN\"\n```\n\n\n## CLI\nThis project includes VERY basic command line tools for development and running the service. Please remember that you must start the server before you can execute commands.\n\nFor a full list of supported operations, please see\n```bash\n$ ./gitness --help\n```"
	return markdownContent, nil
}

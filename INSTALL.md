# Installing Plutono

Plutono and its packages can be installed to your local machine or run from a Docker container.

## Getting started

Plutono has custom depdendencies that are published to the Github package registry. In order to install these packages, you must generate a personal access token and save it to an environment variable.

1. Clone this repository.
1. Visit [github.com/settings/tokens](https://github.com/settings/tokens) to generate a new **classic** Personal Access Token with the `read:packages` scope.
1. Add this new token to your environment variables by running `export NODE_AUTH_TOKEN=your_github_personal_access_token`

### Docker

1. To build the Docker image, run `DOCKER_BUILDKIT=1 docker build --secret id=NODE_AUTH_TOKEN --tag plutono .`
1. Run the application: `docker run -p 3000:3000 -it plutono`

### Locally

For consistency and reliability, it's best to use the same version of node that is defined in the [Dockerfile](Dockerfile). This can be easily accomplished using [Node Version Manager](https://github.com/nvm-sh/nvm).

1. Install [Node Version Manager](https://github.com/nvm-sh/nvm)
1. Install the version of node from the Dockerfile: `nvm install v14.15.5`
1. Install yarn: `npm install --global yarn`
1. Create an `.npmrc` file in the root of the `/plutono` directory and add the following contents:
   ```
   //npm.pkg.github.com/:_authToken=${NODE_AUTH_TOKEN}
   @credativ:registry=https://npm.pkg.github.com/
   ```
1. To build the packages, run `yarn install`

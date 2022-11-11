# Plutono UI components library

> **@credativ/plutono-ui is currently in BETA**.

@credativ/plutono-ui is a collection of components used by [Plutono](https://github.com/grafana/grafana)

Our goal is to deliver Plutono's common UI elements for plugins developers and contributors.

See [package source](https://github.com/credativ/plutono/tree/master/packages/plutono-ui) for more details.

## Installation

`yarn add @credativ/plutono-ui`

`npm install @credativ/plutono-ui`

## Development

For development purposes we suggest using `yarn link` that will create symlink to @credativ/plutono-ui lib. To do so navigate to `packages/plutono-ui` and run `yarn link`. Then, navigate to your project and run `yarn link @credativ/plutono-ui` to use the linked version of the lib. To unlink follow the same procedure, but use `yarn unlink` instead.

### Storybook 6.x migration

We've upgraded Storybook to version 6 and with that we will convert to using [controls](https://storybook.js.org/docs/react/essentials/controls) instead of knobs for manipulating components. Controls will not require as much coding as knobs do. Please refer to the [storybook style-guide](https://github.com/credativ/plutono/blob/master/contribute/style-guides/storybook.md#contrls) for further information.

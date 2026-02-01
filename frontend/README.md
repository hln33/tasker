# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project
npx sv create my-app
```

To recreate this project with the same configuration:

```sh
# recreate this project
npx sv create --template minimal --types ts --add prettier eslint vitest="usages:unit,component" tailwindcss="plugins:none" mcp="ide:claude-code+setup:remote" --install npm my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Testing

The project uses Vitest with Playwright for browser-based testing.

### Run all tests

```sh
npm test
```

### Run tests in watch mode

```sh
npm run test:unit
```

### Run a specific test file

```sh
npm run test:unit -- path/to/file.spec.ts
```

### Testing setup

- **Test runner**: Vitest v4.0.18
- **Browser automation**: Playwright with Chromium (headless)
- **Svelte testing**: `vitest-browser-svelte` for component rendering
- **Test location**: Co-located with source files

### Test pattern

- Test files follow pattern: `{filename}.svelte.spec.ts`
- Located alongside the component being tested
- Uses `render()` from `vitest-browser-svelte`
- Uses `page` from `vitest/browser` for querying DOM

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.

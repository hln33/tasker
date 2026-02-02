You are able to use the Svelte MCP server, where you have access to comprehensive Svelte 5 and SvelteKit documentation. Here's how to use the available tools effectively:

## Available MCP Tools:

### 1. list-sections

Use this FIRST to discover all available documentation sections. Returns a structured list with titles, use_cases, and paths.
When asked about Svelte or SvelteKit topics, ALWAYS use this tool at the start of the chat to find relevant sections.

### 2. get-documentation

Retrieves full documentation content for specific sections. Accepts single or multiple sections.
After calling the list-sections tool, you MUST analyze the returned documentation sections (especially the use_cases field) and then use the get-documentation tool to fetch ALL documentation sections that are relevant for the user's task.

### 3. svelte-autofixer

Analyzes Svelte code and returns issues and suggestions.
You MUST use this tool whenever writing Svelte code before sending it to the user. Keep calling it until no issues or suggestions are returned.

### 4. playground-link

Generates a Svelte Playground link with the provided code.
After completing the code, ask the user if they want a playground link. Only call this tool after user confirmation and NEVER if code was written to files in their project.

---

## Frontend Development Guidelines

### Button Styling Requirements

When creating or modifying buttons in this project, **ALL buttons must include the `cursor-pointer` class** to ensure proper user experience and accessibility.

#### Required Classes Checklist
For every button element, ensure it has:
- ✅ `cursor-pointer` - **MANDATORY** for all interactive buttons
- ✅ Hover state (`hover:bg-*` or `hover:text-*`)
- ✅ Padding (`px-* py-*` or `p-*`)
- ✅ Border radius (`rounded-lg`)
- ✅ ARIA label for icon-only buttons (`aria-label="..."`)

#### Common Mistakes to Avoid
- ❌ Omitting `cursor-pointer` class
- ❌ Using `cursor-default` on buttons
- ❌ Missing hover states for better UX
- ❌ Forgetting ARIA labels on icon buttons

---

## Git Commit Guidelines

### Conventional Commits Standard

All git commit messages **must** follow the [Conventional Commits] specification.

#### Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Commit Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests or correcting existing tests
- **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation
- **revert**: Reverts a previous commit

#### Examples

```
feat: add error page component with playful messages
fix: handle backend fetch errors gracefully
test: add unit tests for error page component
docs: update README with setup instructions
refactor: extract error messages into constants
style: format code with prettier
chore: update dependencies
```

#### Important Rules

- Use lowercase for type and description
- Keep the description under 72 characters when possible
- Use the imperative mood ("add" not "added" or "adds")
- Do not end the description with a period
- Add a scope (in parentheses) after the type for more context: `feat(button): add hover state`
- Add a body for explaining **what** and **why** (not **how**) for complex changes
- Reference issues in the footer: `Closes #123` or `Refs #456`

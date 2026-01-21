We want to create a CLI command `zw` using Golang that would work similar to the `z` command, but with focus on working with git worktrees.
The idea is that in the current world of LLM assisted coding git worktrees are a great way to be able to work on multiple features, bug-fixes, etc using multiple agents in parallel, while also being able to easily jump between different git branches. The main problem is that git worktrees workflow is very cumbersome and error prone.
We should be able to use `zw` to make this workflow faster and more intuitive, eliminating guess work and provided good fallback. Here are some examples of how `zw` would be used:

## Core functionality

Running a command like `zw ft/create-checkout-endpoint` inside a repo should be able to find the worktree for that branch and navigate you in that worktree.
It should also be able to smartly determinate the best candidate if the given name doesn't completely match an existing branch. For example running `zw che` should be able to still navigate you to the work tree of ft/create-checkout-endpoint (if that the best candidate of course). At first it should do that using fuzzy search but later on we could make it smarter, for example go to the best match order by the last visited worktrees.

### Good fallbacks

If I run a command like `zw refactor-to-base-ui` and no matching worktree exists it should be able to:

- check if a remote branch with that name exists and if so ask if it should create a worktree for that branch and navigate to it
- if no matching worktree or remote branch exists it should prompt the user if the want to create a new branch and worktree with that name and navigate to it. This prompt should also ask from what branch that new branch should be checked out, and also allow editing the name of the branch/worktree before submitting. Later on we could allow users to config the main branch of the repo (or infer it automatically if we can) and always suggest that as the branch to checkout from. And even maybe allowing configure specifically what the default branch for these operations should be, for example maybe the main branch is master but we always create new branches for the develop branch.

## Other features

### Jump to "main" branch

There should be a command that always navigate us to the "main" branch of the repo.

### Exact match

There should be a flag that will bypass the smart search and only look for an exact match.

### Interactive selection

When running a command like `zw us` we should be able to provide a flag that instead of automatically navigating use to the best candidate it will list all the possible candidates of worktrees and remote branches without a worktree and we should be able to interactively select the one we want.
Also if no "query" is provided it should start the interactive selection with every option available.
Interactive selection should support fuzzy searching.

### Deletion

We should be able to add a flag that will find the best candidate and prompt user if they want to delete that worktree. This should also warns if there are uncommitted changes, unpushed commits, etc.

### Cleanup

There should be a flag, for example `zw --prune` or something similar, that should list worktrees that could be deleted. For example worktrees that their branches have been merged, or orphan. This can be interactive. This should also warns if there are uncommitted changes, unpushed commits, etc.

### Configuration

We should be able to configure `zw` globally and locally on repo's root. We should be able to configure:

- what the main branch is
- what the default branch to create new worktrees from should be, etc
- in what location should the worktrees be created
- what the preferred way of matching should be, for example the most recent one, etc.

### Copying files

One of the pain points of using worktrees is that when you create a worktree files that are not tracked by git do not exists in the created worktree. That results in having to manually copy files like `.env` for example. We should be able to configure files that we want to automatically get copied when creating a new worktree.

### Running commands

One other pain point when creating or moving to a worktree is that you need to install dependencies. User should be able to define commands that want to be automatically run when creating a new worktree or moving to one.

### Aliases

We should be able to define aliases for common branch to fast jump to them.

### Dry Run

There should be a `--dry-run` flag that will only explain what would happen.

### Bypass prompts

There should be flag that will bypass any prompt when possible with accepting the default option.

### Init

There should be flag that will "init" `zw` in the current repo, creating a local config.

### Jump to previous worktree

There should be flag that will allow to quickly jump on the previous worktree so you can easily go back and forth worktrees.

### Showing "parent" branch of worktree

We should be able to have a setting that when enabled everytime we show a worktree or branch we also show the branch it was checked out from.

### Only "children" of branch

There should be flag that will narrow down all the operations to worktrees/branches that are parent of the provided branch.

### Protected worktrees

We should be able to "lock" worktrees so that they cannot be delete except it's explicit override.

### Debug/Verbose

Support `--verbose` flag.

### Tab completion

Support tab completion like `z`. For example writing `zw user` and hitting tab should complete the best candidate, etc.

## Technical considerations

We are going to use Golang and any required dependencies.
The tool should have proper documentation/--help command for easy of use and discoverability.
We should use only well-known actively maintained dependencies.
Code should be thoroughly tested.
We should use well-known methods of distribution.
UX and performance are critical.
This is going to be an open source tool hosted on github. We should have action for testing, linting etc to ensure code quality. Proper versioning with changelogs etc.

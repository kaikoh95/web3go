# web3go
Web3 + Go

### Collaborators

1. Ryan Tan

2. Kai Koh

### Development Workflow

When creating a new feature, ensure that you are on `main` branch

```
$ git branch
# should expect to see * main

$ git checkout main
# else checkout to main branch
```

Checkout from main branch to a new branch for your feature.

The new feature branch should be named with the prefix `feature/` and your feature's title.

```
$ git checkout -b feature/my-feature
```

Pushing your feature branch to repo (first time)

```
$ git push --set-upstream origin feature/my-feature
```

Pushing your feature branch to repo (subsequently)

```
$ git push
```

When ready for testing on development environment,

Go to Github, create a Pull Request (PR) to `develop` branch

Add other collaborators as the Reviewers.

Once your PR has been approved and merged, you should see it appear on the `develop` branch.

### Release Workflow

When your feature branch is ready to be deployed

Create a PR back to `main` branch

Add other collaborators as the Reviewers.

Once your PR has been approved and merged, you should see it appear on the `main` branch.
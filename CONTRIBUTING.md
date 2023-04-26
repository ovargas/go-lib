# Contributing

## Development environment

* [Go Lang](https://golang.org) (minimum version 1.19)
  *  Ubuntu: `sudo apt install golang`
  *  OS X: `brew install go`
  *  Windows: [Download](https://golang.org/dl/) and install the MSI package

## Check out the repository

Clone the repository in your preferred location:

```sh
git clone https://github.com/ovargas/go-lib.git
```

## License headers

The license headers are automatically added to all source files. To add the license headers to all files, run:

```sh
make copyright
```

## Commit messages and pull requests titles

Follow the [semantic release](https://github.com/semantic-release/semantic-release) format for commit messages and pull 
requests titles, following the [Angular Commit Message Conventions](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format)

```text
<type>(<scope>): <short summary>
  │       │             │
  │       │             └─⫸ Summary in present tense. Not capitalized. No period at the end.
  │       │
  │       └─⫸ Commit Scope: animations|bazel|benchpress|common|compiler|compiler-cli|core|
  │                          elements|forms|http|language-service|localize|platform-browser|
  │                          platform-browser-dynamic|platform-server|router|service-worker|
  │                          upgrade|zone.js|packaging|changelog|docs-infra|migrations|
  │                          devtools
  │
  └─⫸ Commit Type: build|ci|docs|feat|fix|perf|refactor|test
```

In the scope, use the name of the package that you are modifying. For example, if you are modifying the `collections` 
package the commit message should look like:

```text
feat(collections): add new method to the collections package
```

## Testing

Run unit tests:
```sh
make test
```
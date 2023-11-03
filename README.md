# git-composition
`git-composition` is a tool for simplifying composition of standardized git
commit messages.  It is intended to be used with a git commit-msg hook which
will automatically compose commit messages for you.

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

## Installation
`git-comp` is a Go program, so you can install it with `go get`:

```bash
go get github.com/hsoj/git-composition
```

## Usage
`git-comp` needs to be initialized before it can be used.  To initialize it,
run the following command:

```bash
git comp init
```

This will create a `.git-comp.yaml` file in your home directory.  This file
contains the configuration for `git-comp`.  You can edit this file to change
the configuration or use the `git comp config` command to change the 
configuration.

Once `git-comp` is initialized, you can use it to compose commit messages.  To
compose a commit message, run the following command:

```bash
git comp
```

## Configuration
`git-comp` is configured using a YAML file.  The default configuration file is
located at `~/.git-comp.yaml`.  You can change the location of the
configuration file by using the `--config` flag.  For example:

```bash
git comp --config /path/to/config.yaml
```

The configuration file has the following structure:

```yaml
# The version the configuration file was created with.
version: 1
# The default commit message template.
template: |
  {{.Type}}({{.Scope}}): {{.Summary}}

  {{.Body}}

  {{.Coauthors}}
  {{.Footer}}
# The default commit message type.
type: feat
# The default commit message scope.
scope: ""
# List of known authors.
authors:
  # The name of the author.
  - name: John Smith
    # The email of the author.
    email: john.smith@email.com
```

## Semantic Commit Messages
`git-comp` uses [semantic commit messages](https://seesparkbox.com/foundry/semantic_commit_messages)
to compose commit messages.  The commit message is composed of the following
parts:

* Type
* Scope
* Summary
* Body
* Coauthors
* Footer

### Type
The type of the commit message.  The type is used to categorize the commit
message.  The type is required and must be one of the following:

* feat
* fix
* docs
* style
* refactor
* perf
* test
* build
* ci
* chore
* revert

### Scope
The scope of the commit message.  The scope is used to specify what part of the
project the commit message is for.  The scope is optional.

### Summary
The summary of the commit message.  The summary is a short description of the
commit message.  The summary is required and must be less than 50 characters.

### Body
The body of the commit message.  The body is a longer description of the commit
message.  The body is optional.

### Coauthors
The coauthors of the commit message.  The coauthors are people who contributed
to the commit message.  The coauthors are optional.

### Footer
The footer of the commit message.  The footer is used to specify any issues
that are fixed by the commit message.  The footer is optional.

## License
`git-composition` is licensed under the MIT license.  See the LICENSE file for
more information.

## Contributing
Contributions are welcome!  Feel free to open an issue or submit a pull request
at any time.

## Code of Conduct
This project has adopted the [Contributor Covenant](https://www.contributor-covenant.org/)
as its code of conduct.  See the CODE_OF_CONDUCT file for more information.

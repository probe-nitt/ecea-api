## Contributing to ecea-Server

Here are the guidelines we'd like you to follow:

- [Coding Rules](#rules)
- [Commit Message Guidelines](#commit)

---

**NOTE:**

Never push directly to main repository (upstream). Only push to your forked repo (origin) and send a pull request to
the main repository

---

### <a id="rules"></a> Coding Rules

To ensure consistency throughout the source code, keep these rules in mind as you are working:

- Follow the suggested vscode settings mentioned in Readme.md 


### <a id="commit"></a> Git Commit Guidelines

#### Commit Message Format

Each commit message consists of a **header**, a **body** and a **footer**. The header has a special
format that includes a **type**, a **scope** and a **subject**:

```bash
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

Any line of the commit message cannot be longer 100 characters! This allows the message to be easier to read on github
as well as in various git tools.

#### Example Commit Message

```bash
feat(Auth): Implement Login and Register 

create login and registration route with respective logics in controllers

```

Please follow the conventions followed [here](http://karma-runner.github.io/latest/dev/git-commit-msg.html).

Also, refer [this page](https://chris.beams.io/posts/git-commit/) on how to write the body
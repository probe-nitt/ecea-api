# Probe-Server
___

### Requirements
* [Go](https://go.dev/)
* [golangci](https://golangci-lint.run/usage/install/)
* [Docker](https://www.docker.com/)

### Setup
* Configure .vscode/settings.json
    ```
    {
        "go.lintTool":"golangci-lint",
        "go.lintFlags": [
        "--fast"
        ],
        "go.lintOnSave": "package",
        "go.formatTool": "goimports",
        "go.useLanguageServer": true,
        "[go]": {
            "editor.formatOnSave": true,
            "editor.codeActionsOnSave": {
                "source.organizeImports": true
            }
        },
        "go.docsTool": "gogetdoc"
    }
    ```
* Enable githooks
    ``` sh
    git config core.hooksPath .githooks
    ``` 
* Create config.json and .env files
    ``` sh
    cp config.example.json config.json
    ```
    ``` sh
    cp .env.example .env
    ```

* Install reflex
   ``` sh
   go install github.com/cespare/reflex@latest
   ```


* Configure database settings in config.json

### Run
- #### On Local
    ``` sh
    make run
    ```
   
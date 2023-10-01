# litterrobot
A CLI for the Litter Robot 3.

## Usage
### Top Level Commands

```console
$ litterrobot
Control your Litter Robot 3 Connect

Usage:
  litterrobot [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Manage local configuration for the CLI
  help        Help about any command
  version     print the version information

Flags:
  -h, --help   help for litterrobot

Use "litterrobot [command] --help" for more information about a command.
```

### Config
> The config file is stored at `~/.litterrobot/config.yaml`. This location is not currently configurable.

Use the `config` subcommands to configure the CLI.

#### Initialize the config file
```console
$ litterrobot config init
```

#### Set a Config Value
```console
$ litterrobot config set password changeIt
```

#### Get a Config Value
```console
$ litterrobot config get password
*******

# To show sensitive values, use the --show-secret flag
$ litterrobot config get password --show-secrets
changeIt
```

#### Show all Config Values
```console
$ litterrobot config show                 
email: email@example.com
password: *******

# To show sensitive values, use the --show-secrets flag
$ litterrobot config show --show-secrets
email: email@example.com
password: changeIt
```

### Login
Use the `login` command to authenticate to Litter Robot.

This command writes the `token` to the config file. Keep it secret!

```console
$ litterobot login
```

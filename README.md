# qbit tools

## Install

```sh
CGO_ENABLED=0 go build
sudo mv qbit-tools /usr/local/bin/qbit-tools
sudo chmod +x /usr/local/bin/qbit-tools
```

## Usage

```sh
$ qbit-tools --help

Usage: qbit-tools <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  tracker-status
    Label tracker errors

  backup
    Backup torrents as zip

Run "qbit-tools <command> --help" for more information on a command.
```


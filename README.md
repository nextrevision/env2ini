# env2ini

[![Circle CI](https://circleci.com/gh/nextrevision/env2ini.svg?style=svg)](https://circleci.com/gh/nextrevision/env2ini)

Creates and updates INI files from environment variables.

## Installation

Install with go:

```
go get github.com/nextrevision/env2ini
```

With curl using a specific release:

```
curl -o env2ini \
  https://github.com/nextrevision/env2ini/releases/download/0.2.0/env2ini_linux_amd64
chmod +x env2ini
```

## Usage

```
Usage of env2ini:
  -debug
        enable debug logging
  -filename string
        destination filename for writing settings (required)
  -prefix string
        environment prefix to look for keys (required)
```

To configure the file `/etc/config.ini` with environment variables starting
with `MYAPP`, you would run:

```
$ export MYAPP__DEFAULT__conn="mysql://user:pass@mysql/myapp"
$ export MYAPP__section1__key1=myapp_value1

$ ./env2ini -filename /etc/config.ini -prefix MYAPP
INFO[0000] Updated setting                               key=conn section=DEFAULT
INFO[0000] Updated setting                               key=key1 section=section1
```

Which would result in `/etc/config.ini` looking like:

```
conn = mysql://user:pass@mysql/myapp

[section1]
key1 = myapp_value1
```

## Key Syntax

For a key to be added into an INI configuration, the following must be true:

* The variable matches the following syntax: `PREFIX__SECTION__KEY=value`
* There is a support for the config file listed in the table below

The double underscore (`__`) serves as a delimiter for the prefix, section, and
key. Each of those components can contain underscores themselves, just not two
consecutive.

## Examples

### Simple

To update a the key `connection` in section `database` to the value
`mysql://user:pass@host:3306/db` in file `/etc/config.ini`, a variable
would be exported in the environment as:

```
MYAPP__database__connection=mysql://user:pass@host:3306/db
```

After running `env2ini`, we could expect the following in
`/etc/config.ini`:

```
[database]
connection = mysql://user:pass@host:3306/db
```

### Complex Section/Key Names

To update a section or key with a colon, period (dot), or forward slash, we can
make use of a few reserved words: `COLON`, `DOT`, and `SLASH`. If we wanted to
reverse engineer the following output in `/etc/config.ini`:

```
[composite:main]
use = egg:Paste#urlmap
/v2.0 = public_api
/v3 = api_v3
/ = public_version_api
```

We can use the following exported variables:

```
MYAPP__COMPOSITE_COLON_MAIN__USE='egg:Paste#urlmap'
MYAPP__COMPOSITE_COLON_MAIN__SLASH_V2_DOT_0='public_api'
MYAPP__COMPOSITE_COLON_MAIN__SLASH_V3='api_v3'
MYAPP__COMPOSITE_COLON_MAIN__SLASH='public_version_api'
```

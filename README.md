# osenv2ini

[![Circle CI](https://circleci.com/gh/nextrevision/osenv2ini.svg?style=svg)](https://circleci.com/gh/nextrevision/osenv2ini)

Converts shell environment variables to OpenStack INI configuration files.

## Usage

```
Usage of osenv2ini:
  -debug
        enable debug logging
```

## Key Syntax

For a key to be added into an INI configuration, the following must be true:

* The variable matches the following syntax: `PREFIX__SECTION__KEY=value`
* There is a support for the config file listed in the table below

The double underscore (`__`) serves as a delimiter for the prefix, section, and
key. Each of those components can contain underscores themselves, just not two
consecutive.

## Supported Configs

| Prefix         | Config File                      | Defaults |
|----------------|----------------------------------|----------|
| KEYSTONE       | /etc/keystone/keystone.conf      | DEFAULTS |
| KEYSTONE_PASTE | /etc/keystone/keystone-paste.ini | DEFAULTS |

## Examples

### Simple

To update a SQL connection string in `/etc/keystone/keystone.conf`, a variable
would be exported as

```
KEYSTONE__DATABASE__CONNECTION=mysql://user:pass@host:3306/db
```

After running `osenv2ini`, we could expect the following in
`/etc/keystone/keystone.conf`:

```
[database]
connection = mysql://user:pass@host:3306/db
```

### Complex Section/Key Names

To update a section or key with a colon, period (dot), or forward slash, we can
make use of a few reserved words: `COLON`, `DOT`, and `SLASH`. If we wanted to
reverse engineer the following output in `/etc/keystone/keystone-paste.ini`:

```
[composite:main]
use = egg:Paste#urlmap
/v2.0 = public_api
/v3 = api_v3
/ = public_version_api
```

We can use the following exported variables:

```
KEYSTONE_PASTE__COMPOSITE_COLON_MAIN__USE='egg:Paste#urlmap'
KEYSTONE_PASTE__COMPOSITE_COLON_MAIN__SLASH_V2_DOT_0='public_api'
KEYSTONE_PASTE__COMPOSITE_COLON_MAIN__SLASH_V3='api_v3'
KEYSTONE_PASTE__COMPOSITE_COLON_MAIN__SLASH='public_version_api'
```

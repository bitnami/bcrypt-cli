[![CI](https://github.com/bitnami/bcrypt-cli/actions/workflows/main.yml/badge.svg)](https://github.com/bitnami/bcrypt-cli/actions/workflows/main.yml)

# bcrypt

This tool allows generating the bcrypt hash for the provided password through stdin 

# Basic usage

~~~bash
$> bcrypt --help
Usage:
  bcrypt [OPTIONS]

Application Options:
  -c, --cost=COST    The cost weight, range of 4-31 (default: 10)

Help Options:
  -h, --help         Show this help message
~~~

# Examples

## Hash the provided password with default cost (10)

~~~bash
$> echo -n supersecret | bcrypt
$2a$10$m8hh/nNMb.2krzysVyoRVOS3LoSZha7rwV4lfBnCvasTBKAGC.X4i
~~~

## Hash the provided password with higher cost

~~~bash
$> echo -n supersecret | bcrypt --cost=16
$2a$16$qYLz4PvNXK4tS3LZeud0OutgUDE3yX0KwJgMq3zlg/uOKjaUnrdwy
~~~

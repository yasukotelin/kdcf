# kdcf

```
kdcf is a CLI tool for formatting Kotlin data classes.

e.g.
$ kdcf 'User(name="user name", age=20)'

Usage:
  kdcf [flags]

Flags:
  -h, --help      help for kdcf
  -v, --version   version for kdcf
```

## Install

```bash
go install github.com/yasukotelin/kdcf@latest
```

## Usage

```bash
kdcf 'MyPage(user=User(name=name, age=20), coin=100, posts=4)'
```

```bash
MyPage(
    user = User(
        name = name,
        age = 20
    ),
    coin = 100,
    posts = 4
)
```

```bash
kdcf '[Category(id=1, name=category1, locations=[Location(id=1, name=location1)]), Category(id=2, name=category2, locations=[Location(id=2, name=location2), Location(id=3, name=location3)]), Category(id=3, name=category3, locations=[Location(id=4, name=location4)])]'
```

```bash
[
    Category(
        id = 1,
        name = category1,
        locations = [
            Location(
                id = 1,
                name = location1
            )
        ]
    ),
    Category(
        id = 2,
        name = category2,
        locations = [
            Location(
                id = 2,
                name = location2
            ),
            Location(
                id = 3,
                name = location3
            )
        ]
    ),
    Category(
        id = 3,
        name = category3,
        locations = [
            Location(
                id = 4,
                name = location4
            )
        ]
    )
]
```

## LICENSE

MIT License

## Author

Telin. (github.com/yasukotelin)
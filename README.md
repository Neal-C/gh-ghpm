# A gh-extension to manage privacy on github

> [!WARNING]
> Turning a starred repository into a private repository will lose all the stars

> [!IMPORTANT]
> if it has >= 1 stars, ghpm does not turn a starred repository into a private repository.  
> It does not turn your README repository (username/username) private because it's a special repository meant for public display

> [!NOTE]
> I am not sponsored by github, nor affiliated, but you can change that by pinging them on social media. And ask for this functionality has been integrated directly into the gh CLI

## Requirements 

- The Github CLI https://cli.github.com/

## Installation

```bash
gh extension install Neal-C/gh-ghpm
```

or fancier

```bash
gh extension install https://git.github.com/Neal-C/gh-ghpm
```

List your extensions

```bash
gh extension list
```

Upgrade

```bash
gh extension upgrade ghpm
```

Uninstall

```bash
gh extension remove ghpm
```

## Usage

```bash
# prints help message
gh ghpm --help
```

```bash
# turns all your repositories private (except starred repos)
gh ghpm thanos_snap
```

## Roadmap

- [x] switch every repositories to private (excluding repos with >= 1 stars)

## Contributing

I am open to random pull requests that do at least 1 of the following :
- cross items off the roadmap
- fix typos
- add tooling
- add tests
- add/improve documentation
- improve CI/CD


## How to permanently delete or hide data from a repository ?

Only sure way is to contact github support : https://support.github.com/

When in doubt, revoke and rotate your keys. Or better yet, automate it.

---

Made with 💞 love 💞 for developers by a developer ❤️




# Git ToolKit

This is my personal git toolkit, built off [mritd's GitFlow Toolkit](https://github.com/mritd/gitflow-toolkit)

> GitFlow Toolkit is a gitflow commit tool written by go, used to standardize the format of git commit message and quickly create gitflow branches,
> It should be noted that GitFlow Toolkit currently only supports the generation of the commit message style of the [Angular community specification](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#heading=h.greljkmo14y0).


|                                                                                                                                               |                                                                                                                                              | 
|:---------------------------------------------------------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------------------------------------------------:|
| Install | Uninstall |
| <img width="500" src="https://user-images.githubusercontent.com/13043245/134647305-a1df0023-972b-48c3-a6bf-668e96094df9.gif"> | <img width="500" src="https://user-images.githubusercontent.com/13043245/134646600-976f01b4-6000-41e7-8389-0d0e761e15c9.gif"> |
| Commit Success | Commit Failed |
| <img width="500" src="https://user-images.githubusercontent.com/13043245/134485491-993ef0cb-7438-4c42-9a2e-16db05503a0b.gif"> | <img width="500" src="https://user-images.githubusercontent.com/13043245/134485537-6375d280-10d2-4475-a834-7d0ad72248aa.gif"> |
| Push Success | Push Failed |
| <img width="500" src="https://user-images.githubusercontent.com/13043245/134485533-3a01d3be-0912-45cb-9e63-d343a7bad847.gif"> |  <img width="500" src="https://user-images.githubusercontent.com/13043245/134485503-f7de0493-6d2d-403d-aa4d-79a62a83c048.gif"> |
| Create Branch | |
| <img width="500" src="https://user-images.githubusercontent.com/13043245/134485549-5ee7853d-1cc7-4a0f-b083-03514045f8eb.gif">  | |

## Installation

### x86_64

Currently, I'm only building binaries for `x86_64`, you can download and install as follows:

```sh

export VERSION='v0.5'

# download bin file
wget https://github.com/BenDundon/git-toolkit/releases/download/${VERSION}/gitflow-toolkit-x86_64

# add permissions
chmod +x gitflow-toolkit-x86_64

# install
sudo ./gitflow-toolkit-x86_64 install

```

### Other arch/Install from source

Requires:

- `golang`

```sh
# Clone the repository
git clone https://github.com/BenDundon/git-toolkit.git

cd git-toolkit

# build the binary
go build

sudo ./git-toolkit install
```

After the installation is complete, you can delete the bin file.


## Comands

| cmd                 | desc                                                      |
|---------------------|-----------------------------------------------------------|
| `git ci`            | Enter commit message interactively                        |
| `git ps`            | Push the current branch to the remote                     |
| `git feature NAME`  | Switch a new branch from the current branch (`feature/NAME`) |
| `git fix NAME`      | `git switch -c fix/NAME`                                  |
| `git hotfix NAME`   | `git switch -c hotfix/NAME`                               |
| `git docs NAME`     | `git switch -c docs/NAME`                                 |
| `git style NAME`    | `git switch -c style/NAME`                                |
| `git refactor NAME` | `git switch -c refactor/NAME`                             |
| `git chore NAME`    | `git switch -c chore/NAME`                                |
| `git perf NAME`     | `git switch -c perf/NAME`                                 |
| `git style NAME`    | `git switch -c style/NAME`                                |
| `git release NAME`  | `git switch -c release/NAME`                              |


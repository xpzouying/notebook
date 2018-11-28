# notebook
Notebook for Memo

## git

- How to sync a fork repo with remote/origin repo

```bash
# clone your fork
git clone git@github.com:YOUR-USERNAME/YOUR-FORKED-REPO.git

# Add remote from original repository in your forked repository:
cd into/cloned/fork-repo
git remote add upstream git://github.com/ORIGINAL-DEV-USERNAME/REPO-YOU-FORKED-FROM.git
git fetch upstream

# Updating your fork from original repo to keep up with their changes:
git pull upstream master

# rebase
git checkout local-master
git rebase upstream/master
```

## bash ##
- Rename filename in batch way

```bash
for file in $1/*.dat ; do mv "$file" "${file%.*}.txt" ; done
```

## vim ##

- [Execute current line in bash from vim](https://stackoverflow.com/questions/19883917/execute-current-line-in-bash-from-vim)

    `:.w !bash`

    or we could execute some bash command text in vim,
    1. select the bash text in vim by `v` or select the whole line by `V`
    2. `:` and select range ```:`<,`>w !bash```


## Linux

- Find all ip in private network:
    > 1. ping broadcast in subnetwork, like `192.168.1.255`
    > 2. arp to find ip

```bash
ping 192.168.1.255
arp -a
```

## Programming

### protobuf

`protoc --go_out=. *.proto`

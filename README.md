# notebook
Notebook for Memo


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
    2. `:` and select range `:\`<,\`>w !bash`


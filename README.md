# terminal-truncate

truncate text (with ansi color) for terminal

## how to install
```
go get -u github.com/umaumax/terminal-truncate
```

## ISSUE
* tabが最大のスペース数で置換されている
  * ansi color codeを解釈しつつtabの数を置換する必要がある

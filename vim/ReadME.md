
vim-go插件
```
let g:go_fmt_command = "goimports" " 格式化将默认的 gofmt 替换
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"
let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1
let g:godef_split=2
```
#**************************************************************
#   my bash shell

# 环境变量
export GO111MODULE=auto
export GOPROXY="https://goproxy.cn"
export GOROOT=/usr/local/go
export MYSQL_PATH=/usr/local/mysql
export PKG_CONFIG_PATH=/usr/lib/pkgconfig
export LD_LIBRARY_PATH=/usr/local/Cellar/zeromq/4.3.4/lib

export GOPATH=/Users/halou/go
export REDIS_PATH=/Users/halou/redis
#export CLEOS_PATH=/Users/halou/eosio
# export CGO_ENABLED=0

export PATH=$PATH:$GOROOT/bin:$REDIS_PATH/src:$MYSQL_PATH/bin:$GOPATH/bin:$PKG_CONFIG_PATH

# 关闭Homebrew自动更新
export HOMEBREW_NO_AUTO_UPDATE=true

# Mysql相关环境变量 ############################
export MYSQL_BASE=/usr/local/mysql

# system
alias redis_start='sudo /usr/local/bin/redis-server /Users/halou/redis/redis.conf'

# 自定义快捷指令
alias cdgit='cd /Users/halou/work/ystar'
alias cdbg='cd /Users/halou/work/ystar/bingoo-gateway'
alias cdlog='cd /Users/halou/work/workLog'
alias cdexe='cd /Users/halou/work/Exercise'
alias cdnft='cd /Users/halou/work/ystar/backend-server'
alias cdwork='cd /Users/halou/work'
alias cdng='cd /usr/local/var/www'
alias gs='git status'
alias gcbw='git status;git checkout -b wangch;git status'
alias getnew='git checkout master;git pull'

alias gogrep='find . -path "./vendor" -prune -o -name "*.go" -print | xargs grep --color -n';

#**************************************************************

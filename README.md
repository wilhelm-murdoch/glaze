# glaze

```
NAME:
   glaze - easily manage tmux windows and panes

USAGE:
   glaze [global options] command [command options] [arguments...]

VERSION:
   dev

AUTHOR:
   Wilhelm Murdoch <wilhelm@devilmayco.de>

COMMANDS:
   apply    apply the specified glaze profile
   fmt      rewrites the target glaze profile file to a canonical format
   save     running this within a tmux session will save its current state to the specified glaze profile
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print only the version (default: false)

COPYRIGHT:
   (c) 2023 Wilhelm Codes ( https://wilhelm.codes )
```

```
go run main.go apply ../../.glaze                                    
new-session -d -s "Session One"
new-window -t "{name}" -n "Window One"
set-window-option -t "{name}" main-pain-width 50
split-window -t "{name}" -n "Terminal"
split-window -t "{name}" -n "Logs"
split-window -t "{name}" -n "Editor"
new-window -t "{name}" -n "Window Two"
set-window-option -t "{name}" main-pain-width 50
split-window -t "{name}" -n "One"
split-window -t "{name}" -n "Two"
split-window -t "{name}" -n "Three"
split-window -t "{name}" -n "Four"

```
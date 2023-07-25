# glaze

- reindex panes while respecting base index
- reindex windows while respecting base index
- remove default window and pane on creation
- `tmux show-option -g base-index`
- fix send keys to send to the appropriate `session:window.pane` target
- set active window / pane
- pass `Environment` or `EnvironmentFile` for session > window > pane cascading
- `--debug` flag or `glaze debug ...` which spits out the list of tmux commands used to spin up a profile
- use charm for menu-based selection of glaze profiles
- `client.GetOption` & `client.SetOption` global, session, window + pane?
- add support for `err` in `collection.Each`


```hcl
session "Glaze" {
  window "Workspace" {
    layout = "tiled"
    starting_directory = "~/Development/github/wilhelm-murdoch/glaze"

    pane "Terminal" {
      split = "vertical"
      focus = true
      commands = [
        "echo \"Starting...\"",
      ]

        environment "USING_GLAZE" {
            value = "yes"
        }
    }

    pane "Watcher" {
      split = "horizontal"
      starting_directory = "~/Development/github/wilhelm-murdoch/glaze/cmd/glaze/"
      commands = [
        "echo \"Watching...\"",
      ]
    }

    pane "Tester" {
      split = "vertical"
      starting_directory = "~/Development/github/wilhelm-murdoch/glaze/"
      commands = [
        "echo \"Testing...\"",
      ]
    }
  }
}
```
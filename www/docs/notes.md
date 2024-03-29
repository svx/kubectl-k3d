# Notes

## Cobra

### Version

- https://www.jvt.me/posts/2023/05/27/go-cobra-version/

## Go Documentation

- https://www.jetbrains.com/guide/go/tips/quick-documentation-go-doc-comments/
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
- https://gist.github.com/jgfrancisco/6610078040b15b7e9611
- https://www.reddit.com/r/golang/comments/r23jy9/getting_stdin_and_stdout_from_spawned_process_in/
- https://international.github.io/2017/01/09/11-45-golang_redirect_os.Stdout
- https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string/10476304#10476304
- https://www.socketloop.com/tutorials/golang-capture-stdout-of-a-child-process-and-act-according-to-the-result
- 

## Mage

- https://carolynvanslyck.com/blog/2021/01/mage-is-my-favorite-make/
- https://carolynvanslyck.com/blog/2022/07/magefiles-stop-on-first-error/
- https://cloudnativeislamabad.hashnode.dev/say-the-magic-words-and-go-mage-will-take-care-of-the-rest
- https://github.com/helm/chart-releaser/blob/main/magefile.go

## GolangCI

Installed via brew:

```shell
brew install golangci-lint
brew upgrade golangci-lint
```

- https://olegk.dev/go-linters-configuration-the-right-version

## Path stuff

- https://forum.golangbridge.org/t/help-with-simple-idiomatic-go/22308/3

## k3d

### Config file

- https://k3d.io/v5.4.1/usage/configfile/
- https://chipnibbles.com/create-your-own-kubernetes-cluster/

### Commands

```shell
k3d cluster create foo --config k3d-default.yaml
```

```shell
3d cluster create --config k3d-default.yaml
```

```shell
 k3d cluster create \
    "test1" \
    --agents 2 \
    --port="80:80@loadbalancer" \
    --port="443:443@loadbalancer" \
    --k3s-arg="--disable=traefik@server:0" \
  --image docker.io/rancher/k3s:v1.27.5-k3s1
```

## CLI UI

- https://medium.com/@nexidian/writing-an-interactive-cli-menu-in-golang-d6438b175fb6
- https://elewis.dev/charming-cobras-with-bubbletea-part-1
- https://github.com/elewis787/rkl
- https://medium.com/@originalrad50/building-ui-of-golang-cli-app-with-bubble-tea-68b61e25445e
- https://github.com/elewis787/boa
- https://earthly.dev/blog/tui-app-with-go/
- https://github.com/inngest/inngest
- https://github.com/charmbracelet/taskcli
- https://github.com/cqroot/prompt

## Goreleaser

- https://github.com/inngest/inngest/blob/main/.goreleaser.yml
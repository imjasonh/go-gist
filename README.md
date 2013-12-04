go-gist
=======

Ever see some tasty Go code in a [GitHub Gist][1] and want to try it out? What are you supposed to do, clone it into your own GitHub repo and go get it? Copy it locally? What are you a farmer?

Why can't you just `import "gist.github.com/UserName/11235813"`? Because Gists don't support that, unfortunately. :(

Which is why I made this silly hack. Instead of `import "gist.github.com/..."`, `import "go-gist.appspot.com/..."`, and this site will tell the Go tool to go fetch the contents of the Gist instead.

It's all described [in the docs][0]. Those guys thought of everything.

Try it out:
```
go get go-gist.appspot.com/ImJasonH/7791518
```

By the way, this would be totally unnecessary if the Gist HTML page was served with the necessary `<meta>` tag. If you work at GitHub or know someone who does, bug them to add a tag to their Gist HTML so I can turn this site off.

[0]: http://golang.org/cmd/go/#hdr-Remote_import_paths
[1]: http://gist.github.com

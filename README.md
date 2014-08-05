go-gist
=======

Ever see some tasty Go code in a [GitHub Gist][1] and want to try it out? What are you supposed to do, clone it into your own GitHub repo and `go get` it from there? Copy it locally? What are you? A farmer?

What you really want is to `import "gist.github.com/UserName/11235813"`, but Gists don't support that. :cry:

Which is why I made this simple app to solve a problem. Instead of `import "gist.github.com/..."`, `import "go-gist.appspot.com/..."`, and the site will tell the Go tool to go fetch the contents of the Gist.

It's all described [in the docs][0]. Those guys thought of everything.

Try it out:
```
go get go-gist.appspot.com/ImJasonH/7791518
```
or
```
go get go-gist.appspot.com/ImJasonH/7791518/readyset
```

Then:

```
import (
	//...
	readyset "go-gist.appspot.com/ImJasonH/7791518"
)
```
or

```
import (
	//...
	"go-gist.appspot.com/ImJasonH/7791518/readyset"
)
```

By the way, this would be totally unnecessary if the Gist HTML page was served with the necessary `<meta>` tag. If you work at GitHub or know someone who does, bug them to add a tag to their Gist HTML so I can turn this site off.

[0]: http://golang.org/cmd/go/#hdr-Remote_import_paths
[1]: http://gist.github.com


----------

License
-----

    Copyright 2014 Jason Hall

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

# This is

A simple "markdown to html" library.

For the first iteration it only supports bold text, italic and strikethrough text styles.

# Usage

```
make build
```

and then

```
./markdown_to_html "*bold* 
_italic_ 
~strikethroughed~ 
*surprisingly*_works_ 
* surprisingly also works* 
*and ~also~ _works_ like that*"
```
or
```
./markdown_to_html < your.md
```

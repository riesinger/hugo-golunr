# `hugo-golunr`, a golang alternative to [hugo-lunr](https://www.npmjs.com/package/hugo-lunr)

As you probably don't like installing node, npm and a ton of packages into your CI, which generates
a static hugo page, I created this golang implementation of `hugo-lunr`. It generates a lunrjs
search index from the current working directory. 

## Installing

`go get github.com/arial7/hugo-golunr`

## Usage 

```sh
cd /path/to/your/site
hugo-golunr
```

Pretty easy, huh? After running `hugo-golunr`, you'll see a `search_index.json` file in your
`./static` directory. Just load that in your theme.



# netlify-functions-zaimapi
zaim APIでデータ取得するサーバサイド関数

# setup history

```zsh
% go mod init example.com/hello
go: creating new go.mod: module example.com/hello
```


```zsh
% npm install netlify-cli -g

% /Users/yoshiki/.anyenv/envs/nodenv/versions/17.3.0/lib/node_modules/netlify-cli/bin/run login

% /Users/yoshiki/.anyenv/envs/nodenv/versions/17.3.0/lib/node_modules/netlify-cli/bin/run link

Adding local .netlify folder to .gitignore file...

netlify link will connect this folder to a site on Netlify

? How do you want to link this folder to a site? Choose from a list of your recently updated sites
Fetching recently updated sites...

? Which site do you want to link? chipper-liger-5cc95d

Directory Linked

Admin url: https://app.netlify.com/sites/chipper-liger-5cc95d
Site url:  https://chipper-liger-5cc95d.netlify.app

You can now run other `netlify` cli commands in this directory


% /Users/yoshiki/.anyenv/envs/nodenv/versions/17.3.0/lib/node_modules/netlify-cli/bin/run functions:create
? Select the type of function you'd like to create Serverless function (Node/Go)
◈ functions directory not specified in netlify.toml or UI settings
? Enter the path, relative to your site’s base directory in your repository, where your functions should live: netlify/functions
◈ updating site settings with netlify/functions
◈ functions directory netlify/functions updated in site settings
◈ functions directory netlify/functions does not exist yet, creating it...
◈ functions directory netlify/functions created
? Select the language of your function Go
? Pick a template go-hello-world
? Name your function: zaimapi
◈ Creating function zaimapi
◈ Created netlify/functions/zaimapi/go.mod
◈ Created netlify/functions/zaimapi/go.sum
◈ Created netlify/functions/zaimapi/main.go
```

# 参考サイト

https://qiita.com/lkeix_/items/19c17b2334c6481a18ba
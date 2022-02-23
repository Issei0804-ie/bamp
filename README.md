## What is bamp?

bamp is backup-software that made by Go.

And bamp can use quick that you just write config(You modify crontab if you need it).


## How to install 

If you don't install go, you should install [go](https://go.dev/).

After install go, enter below command to terminal.

`go install github.com/Issei0804-ie/bamp@latest`

## How to use it

First, you need to copy settings.json to your local machine.

```
$ wget "https://raw.githubusercontent.com/Issei0804-ie/bamp/main/settings-copy.json" -O settings.json 
```

Next, you modify settings.json.

example: settings.json
```
{
  "backup_dir": ["/home/issei/Desktop"],
  "store_dir" : "/mnt/backup/issei"
}
```

Finally, execute bamp.

Please set path of settings.json.

```
bamp settings.json
```

example: crontab

```
0 19 * * * /home/issei/go/bin/bamp /home/issei/my-script/settings.json >> /var/log/my-app/backup.log
```

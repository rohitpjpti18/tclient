# tclient - A go based torrent client

The client is able to download single file torrent files with tracker url.

## Usage
In order to download [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist)! from torrent file

Note: Program takes two argument
 - torrent file
 - download file (currently you need to create an empty file for download)
```sh
torrent-client debian-xxxxxxx-amd64-netinst.iso.torrent debian.iso
```

![screenshot](./docs/tclient-ss.gif)


Based on tutorial form https://blog.jse.li/posts/torrent/

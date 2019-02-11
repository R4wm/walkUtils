# walkUtils
65;5403;1cGolang helper lib to do things while walking directory (exercises)

Ex: 
```bash
stat main.go: no such file or directory
[@timeshot walkUtils]$ cd cmd/
[@timeshot cmd]$ 
[@timeshot cmd]$ 
[@timeshot cmd]$ 
[@timeshot cmd]$ go build -o md5Lister main.go
[@timeshot cmd]$ time ./md5Lister /tmp > ~/myMD5list.txt 
2019/02/10 17:19:20 Skip: /tmp is a directory -> tmp
2019/02/10 17:19:20 Skip: /tmp/.ICE-unix is a directory -> .ICE-unix
2019/02/10 17:19:20 Skip: /tmp/.Test-unix is a directory -> .Test-unix
2019/02/10 17:19:20 Skip: /tmp/.X11-unix is a directory -> .X11-unix
2019/02/10 17:19:20 Skip: Not a dir but size is 0, so skipping /tmp/.X11-unix/X0.
2019/02/10 17:19:20 Skip: /tmp/.XIM-unix is a directory -> .XIM-unix
2019/02/10 17:19:20 Skip: /tmp/.font-unix is a directory -> .font-unix
2019/02/10 17:19:20 Skip: /tmp/.org.chromium.Chromium.u0ubuB is a directory -> .org.chromium.Chromium.u0ubuB
2019/02/10 17:19:20 Skip: A Symlink: /tmp/.org.chromium.Chromium.u0ubuB/SingletonCookie -> Lrwxrwxrwx
2019/02/10 17:19:20 Skip: Not a dir but size is 0, so skipping /tmp/.org.chromium.Chromium.u0ubuB/SingletonSocket.
2019/02/10 17:19:20 Skip: /tmp/Temp-4d90137b-1068-4b47-a92f-04d93d8fd94d is a directory -> Temp-4d90137b-1068-4b47-a92f-04d93d8fd94d
2019/02/10 17:19:20 Skip: /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303 is a directory -> Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303
2019/02/10 17:19:20 Skip: /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303/mesa_shader_cache is a directory -> mesa_shader_cache
2019/02/10 17:19:20 Skip: /tmp/firefox_rmintz is a directory -> firefox_rmintz
2019/02/10 17:19:20 Skip: Not a dir but size is 0, so skipping /tmp/firefox_rmintz/.parentlock.
2019/02/10 17:19:20 Skip: /tmp/go-build810309424 is a directory -> go-build810309424
2019/02/10 17:19:20 Skip: /tmp/hey is a directory -> hey
2019/02/10 17:19:20 Skip: /tmp/hey/how is a directory -> how
2019/02/10 17:19:20 Skip: /tmp/hey/how/you is a directory -> you
2019/02/10 17:19:20 Skip: /tmp/hey/how/you/doin is a directory -> doin
2019/02/10 17:19:20 Skip: /tmp/mozilla_rmintz0 is a directory -> mozilla_rmintz0
2019/02/10 17:19:20 open /tmp/systemd-private-542ad9b5433947c4a89557678c133894-redis.service-VxVnUv: permission denied

real	0m0.065s
user	0m0.054s
sys	0m0.011s
[@timeshot cmd]$ 
[@timeshot cmd]$ 
[@timeshot cmd]$ bat ~/myMD5list.txt
───────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
       │ File: /home/rmintz/myMD5list.txt
───────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
   1   │ 05bff0451d3a3d9697cfbfa5f7811433 /tmp/.X0-lock
   2   │ ba36c9e2e9ed85422ebd5e63fbb7b363 /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303/mesa_shader_cache/index
   3   │ d9f47bd4b99cb789a4932c06990a1a98 /tmp/bashExample.txt
   4   │ b52a2b25528110bf50b82ca474962352 /tmp/goExample.txt
   5   │ d36f8f9425c4a8000ad9c4a97185aca5 /tmp/hey/how/you/doin/ok.txt
   6   │ 307c25d470e67483915acb6ee25e0c53 /tmp/md5List.txt
   7   │ e641c10e24edd14debf08b34c0c207ca /tmp/serverauth.umJFV2CDRt
───────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
[@timeshot cmd]$ 
```

Seems like go is faster, but its also not following symlinks at the moment, give or take 100 files to parse..

```bash
[@timeshot cmd]$ time find /home/rmintz/Downloads -type f -exec md5sum {} \; > /tmp/bashExample.txt

real	2m35.233s
user	1m38.843s
sys	0m29.940s
[@timeshot cmd]$ 


real	1m20.741s
user	0m46.146s
sys	0m11.033s
[@timeshot cmd]$ time ./md5Lister /home/rmintz/Downloads/ > /tmp/goExample.txt

[@timeshot cmd]$ wc -l /tmp/goExample.txt /tmp/bashExample.txt 
   80604 /tmp/goExample.txt
   80797 /tmp/bashExample.txt
  161401 total
[@timeshot cmd]$
```
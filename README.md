# walkUtils
Golang helper lib to do things while walking directory (exercises)

Ex: 
```bash

[@timeshot walkUtils]$ go get github.com/r4wm/walkUtils
[@timeshot walkUtils]$ 
[@timeshot walkUtils]$ 
[@timeshot walkUtils]$ go run cmd/main.go /tmp/ > myMD5List.txt
2019/02/10 16:46:19 Skip: /tmp/ is a directory -> tmp
2019/02/10 16:46:19 Skip: /tmp/.ICE-unix is a directory -> .ICE-unix
2019/02/10 16:46:19 Skip: /tmp/.Test-unix is a directory -> .Test-unix
2019/02/10 16:46:19 Skip: /tmp/.X11-unix is a directory -> .X11-unix
2019/02/10 16:46:19 Skip: Not a dir but size is 0, so skipping /tmp/.X11-unix/X0.
2019/02/10 16:46:19 Skip: /tmp/.XIM-unix is a directory -> .XIM-unix
2019/02/10 16:46:19 Skip: /tmp/.font-unix is a directory -> .font-unix
2019/02/10 16:46:19 Skip: /tmp/.org.chromium.Chromium.u0ubuB is a directory -> .org.chromium.Chromium.u0ubuB
2019/02/10 16:46:19 Skip: A Symlink: /tmp/.org.chromium.Chromium.u0ubuB/SingletonCookie -> Lrwxrwxrwx
2019/02/10 16:46:19 Skip: Not a dir but size is 0, so skipping /tmp/.org.chromium.Chromium.u0ubuB/SingletonSocket.
2019/02/10 16:46:19 Skip: /tmp/Temp-4d90137b-1068-4b47-a92f-04d93d8fd94d is a directory -> Temp-4d90137b-1068-4b47-a92f-04d93d8fd94d
2019/02/10 16:46:19 Skip: /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303 is a directory -> Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303
2019/02/10 16:46:19 Skip: /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303/mesa_shader_cache is a directory -> mesa_shader_cache
2019/02/10 16:46:19 Skip: /tmp/firefox_rmintz is a directory -> firefox_rmintz
2019/02/10 16:46:19 Skip: Not a dir but size is 0, so skipping /tmp/firefox_rmintz/.parentlock.
2019/02/10 16:46:19 Skip: /tmp/go-build128436224 is a directory -> go-build128436224
2019/02/10 16:46:19 Skip: /tmp/go-build128436224/b001 is a directory -> b001
2019/02/10 16:46:19 Skip: /tmp/go-build128436224/b001/exe is a directory -> exe
2019/02/10 16:46:19 Skip: /tmp/go-build828067301 is a directory -> go-build828067301
2019/02/10 16:46:19 Skip: /tmp/hey is a directory -> hey
2019/02/10 16:46:19 Skip: /tmp/hey/how is a directory -> how
2019/02/10 16:46:19 Skip: /tmp/hey/how/you is a directory -> you
2019/02/10 16:46:19 Skip: /tmp/hey/how/you/doin is a directory -> doin
2019/02/10 16:46:19 Skip: /tmp/mozilla_rmintz0 is a directory -> mozilla_rmintz0
2019/02/10 16:46:19 open /tmp/systemd-private-542ad9b5433947c4a89557678c133894-redis.service-VxVnUv: permission denied
[@timeshot walkUtils]$ 
[@timeshot walkUtils]$ 
[@timeshot walkUtils]$ bat myMD5List.txt
───────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
       │ File: myMD5List.txt
───────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
   1   │ /tmp/.X0-lock 05bff0451d3a3d9697cfbfa5f7811433
   2   │ /tmp/Temp-a820ba7c-22a8-42f2-a21f-8c43b1850303/mesa_shader_cache/index ba36c9e2e9ed85422ebd5e63fbb7b363
   3   │ /tmp/go-build128436224/b001/exe/main 75694d0fa7997abf0b18c195b5482f1b
   4   │ /tmp/go-build128436224/b001/importcfg.link 5ed8ab36d3d6c1e60b3590f8292516b5
   5   │ /tmp/hey/how/you/doin/ok.txt d36f8f9425c4a8000ad9c4a97185aca5
   6   │ /tmp/md5List.txt 307c25d470e67483915acb6ee25e0c53
   7   │ /tmp/serverauth.umJFV2CDRt e641c10e24edd14debf08b34c0c207ca
───────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
[@timeshot walkUtils]$ 
```

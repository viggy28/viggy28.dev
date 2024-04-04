---
date: 2023-11-20T18:00:08-04:00
description: "ICU library issue on MacOS when installing Postgres v16"
featured_image: ""
tags: ["postgres", "C"]
title: "Postgres v16 installation issues wrt ICU"
---

I was trying to play with Phil's [pgtam](https://github.com/eatonphil/pgtam) and the first step is that to install Postgres version 16. Sounds fairly innocous. When I ran `configure`, I was getting the below error:

```
checking for icu-uc icu-i18n... no
configure: error: ICU library not found
If you have ICU already installed, see config.log for details on the
failure.  It is possible the compiler isn't looking in the proper directory.
Use --without-icu to disable ICU support.
```

Obviously, one solution is to disable ICU support as recommended by the error message. However, that's an [anti-feature](https://www.postgresql.org/docs/current/install-make.html#CONFIGURE-OPTIONS-ANTI-FEATURES) as per the Postgres documentation.

So, my next question is how to build and install Postgres 16 with ICU support on Mac.

FWIW, I have built Postgres multiple times in the past and I haven't ran into this issue. So, it is likely with version 16. Also, I did a customary Google search and ran into this [Stackoverflow answer](https://stackoverflow.com/questions/52510499/the-following-icu-libraries-were-not-found-i18n-required) which mentions something to do with v16. So, I started checking PG16 [release log](https://www.postgresql.org/about/news/postgresql-16-released-2715) and noticed the below.

```
PostgreSQL 16 improves general support for text collations, which provide rules for how text is sorted. PostgreSQL 16 builds with ICU support by default, determines the default ICU locale from the environment, and allows users to define custom ICU collation rules.
```

That explained why I am noticing the error "NOW" and not in the past. PG16 builds with ICU support by default.

I started digging into how to pass the right path to Postgres while building it. A simple grep showed the files I should be playing with. Adding any debug message on configure.ac actually didn't print anything to my stdout. Probably that's for another day.

```
[16:25:52] viggy28:postgres git:(v16.1*) $ rg "icu-uc" --files-with-matches * 
meson.build
configure.ac
configure
config.log
```

In configure, I noticed the following assigned. 

```bash
pkg_cv_ICU_CFLAGS=`$PKG_CONFIG --cflags "icu-uc icu-i18n"`
```

I found the value of `PKG_CONFIG` to `/usr/local/bin/pkg-config`. So running `/usr/local/bin/pkg-config --cflags "icu-uc icu-i18n"` yielded the following error.

```
[16:34:47] viggy28:viggy28.dev git:(article/pg16-installation-icu*) $ /usr/local/bin/pkg-config --cflags "icu-uc icu-i18n"
Package icu-uc was not found in the pkg-config search path.
Perhaps you should add the directory containing `icu-uc.pc'
to the PKG_CONFIG_PATH environment variable
No package 'icu-uc' found
Package icu-i18n was not found in the pkg-config search path.
Perhaps you should add the directory containing `icu-i18n.pc'
to the PKG_CONFIG_PATH environment variable
No package 'icu-i18n' found
```

A simple find for the file `icu-uc.pc` resulted in the directory `/usr/local/Cellar/icu4c/73.2/lib/pkgconfig/`. Setting `export PKG_CONFIG_PATH=/usr/local/opt/icu4c/lib/pkgconfig/` and then running `configure` built the postgres with ICU support.

Note: This was an issue with MacOS intel chip. I ran into the same on Arm chip too. To fix this issue the path needs to be modified slighty. It could also be in other directory depends on where `icu4c` package is installed.

```
export PKG_CONFIG_PATH=/opt/homebrew/Cellar/icu4c/74.2/lib/pkgconfig/
```
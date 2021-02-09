---
date: 2019-03-23T18:00:08-04:00
description: "GPG in simple terms"
featured_image: "/images/gpg.png"
tags: ["firebase", "hugo", "hosting"]
title: "GPG in simple terms"
---

GPG in Simple Terms {.f2 .f1-l .fw2 .white-90 .mb0 .lh-title}
===================

Understanding the basics of GPG {.fw1 .f5 .f3-l .white-80 .measure-wide-l .center .lh-copy .mt3 .mb4}
-------------------------------

ARTICLES

GPG in Simple Terms {.f1 .athelas .mb1}
===================

March 9, 2020 - 2 minutes read - 342 words

When I first heard GPG it was little daunting. I haven’t used it
exactly, even though I was familar with the concept. So, I decided to
dig deeper and understand how to use it.

GnuPG (GPG) is an implementation of PGP

The simple idea is that you generate a key pair: secret or private key
and a public key. You can provide your public key to anyone. They can
then encrypt either a secret or file using your public key. You are the
only one who can decrypt the secret or file (assuming you aren’t sharing
your secret key with others, which is bad).

**Use case:**

1.  When sharing a password with someone
2.  When creating database backups

**Jargon Alert:**

***Keyring***: Public and secret keys are required in order to use GPG.
These keys are stored in two files called public and secret keyrings
under \~/.gnupg directory.

**Key Server**: Locating keys on the web or writing to the individual
asking them to transmit their public keys can be time consuming and
insecure. Key servers act as central repositories to alleviate the need
to individually transmit public keys and can act as the root of a chain
of trust.

**To list the public keys in your keyring:**

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --list-keys
```

**To list the secret keys in your keyring:**

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --list-secret-keys
```

**To import a public or private key to your keyring:**

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --import secret.key
gpg --import public.key
```

**To generate a new keypair**

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --generate-key
```

**To export a public in ASCII format**

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --armor --export vigneshravichandran28@gmail.com
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBF5nJ7oBCAC4f8devBv1o+Rpx4nRyytCyIMq2JvBKsYNxsc/NfCoH8IKkGBG
SpMGN0FTOmRg1V23rJHFHvVmF2jQ3gYeygydo3CdPX0Fjxbgmqg7GP2lbAmXlCix
Ty9XEaGRrFUoMwK4DxFYq25HgWF8PWRlALsX9jMDPXCHMabYrNdDFMWXP3VyXR2v
zEz0EC4OE3V7NbQP0w/kAiWPSbR99r2g1KG1XAkRR+ANWBch5k1Vsjezmc8/7G59
F6GRcANjvJ9T3GnRp7ka/AAj4q7/lIc/Zbr+ajtVq9kLdBSrCsB3IP9PoHIzulhH
0JrfzrgV+Y4IigrrfP3b03E++2Js8rLXZHN3ABEBAAG0NnZpZ25lc2ggcmF2aWNo
YW5kcmFuIDx2aWduZXNocmF2aWNoYW5kcmFuMjhAZ21haWwuY29tPokBVAQTAQgA
PhYhBCzrfMeNI66xBeaTWL4+f5R69VVbBQJeZye6AhsDBQkDwmcABQsJCAcCBhUK
CQgLAgQWAgMBAh4BAheAAAoJEL4+f5R69VVbPoMH/0YLcG+tuZIgKAw51m/a6ww1
Nwe3gFX5Jn4tMf4/w+BVxkXVAlekFsfHJWiiZ1fxaGF8plxmhdbM8b/rKJpE80dU
ibP5REN3Ph2QGWDn6jvTpBve6CFOFjE5eDMo+9n7fY9H9bNG3k7CkwRk0LhV/WUR
92Y7H6QOBswcEw7dD/MwdlP7neWrTSzaedAw0rXU9mYO5GLp3gTTsyd0TJhDa7Zd
zOGFeKItXSH1Uka0LUSdRM/pYYNoWd4XveQ7Y4+X00rjPAkL1WND8/TuFRHP4uBP
xykqU4bna+4keMvSIpUqdXqn72YFb4ieUB4+enF5Tzt1QeVb4YInw4GnePwrN+i5
AQ0EXmcnugEIALCO2N1YGfclt7tvJhYbVac2JTDjRYIPE9EAyMqmyeexdMvRxKo+
PuZjZ03Mt40ewWE8dwXNnlkgCU61gUTFbXspiDxAh66Bv9/G8vdycBl/VerU3njJ
ug+sCTeCPdr92uS2HtUUOJj8k/Kcq6ckuc5y+GnubUU/Q1FFHTd93OtFwvRYwmgR
umQmbuuRuGiTfL9/JMOpZ3IG44SKT26r3rOdR+u5e7+Kc4eGVXx0be5X3J/hBesh
VXcWTzz8yEMnsFpnQ4lDjutiVte3ralPeplRe9b9pDfQYSw9mMDQF2EpgxoIAeDF
iOmwQQtY+vp+lMXMoR+lcv9g3trCQAZ0N+8AEQEAAYkBPAQYAQgAJhYhBCzrfMeN
I66xBeaTWL4+f5R69VVbBQJeZye6AhsMBQkDwmcAAAoJEL4+f5R69VVbqawIAJmy
MAtp59XswbjMZ5ZuEPq77SwqffDrLmu/DjTKuU0aU3bbKWiesCH9D4TJSVjdZjBh
tugV9kUBluUH4T87JieOhJ/Kv07a1QEbU7As3gk9Ps9PzLTWPG4BENOpY8PCsl4G
SxGrfcK1WDaoMwZGiM2cXbWpeZza4Q3nCF9EdRQTfzUmASn+vY4qFadhB+GV+eVj
My2J6jGzcm6Pts2E+HRqFIHMmpdvN6flZPYf/AqUNR4P6q46gmJ54tLM30RYxEus
VgmgI8DkKt4MZjaRLSFNglfqzkxNX0QZcJnG/CgA2wH3q026UynEhWfy7ap5P4LJ
xSZE7x2vArQSEFh5XJY=
=+5TP
-----END PGP PUBLIC KEY BLOCK-----
```

**To verify that the public key belongs to the right person** use
fingerprint. It is like 2 factor authentication. You can verify that the
fingerprint matches over phone or in-person. This can get more
cumbersome if you don’t know that person personally. Thats where trust
factor comes in.

``` {style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"}
gpg --fingerprint
```

reference:
[https://www.gnupg.org/gph/en/manual/x334.html](https://www.gnupg.org/gph/en/manual/x334.html)
cheatsheet: [https://devhints.io/gnupg](https://devhints.io/gnupg)

-   [gpg](/tags/gpg)
-   [cryptography](/tags/cryptography)

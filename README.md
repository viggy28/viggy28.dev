# Source for my personal [website](<https://viggy28.dev>)

## How to publish a new post
1. Create a new markdown file in content directory
2. At the top of the file add an index. Example below
```
---
date: 2019-03-24T18:00:08-04:00
description: "GPG in simple terms"
featured_image: "/images/gpg.png"
tags: ["firebase", "hugo", "hosting"]
title: "GPG in simple terms"
---
```
3. Store any images in /public/images directory and refer them like "![whatever name](/images/filename.png)"


## To learn more about the hosting and development checkout the [article](<https://viggy28.dev/post/hosting-a-simple-website/>)

## Deployment instruction

### To run locally:
```
hugo server
```

### To build for production:

```
HUGO_ENV=production hugo
```

### To deploy to firebase (do fl aka. firebase login )
```
fd
```
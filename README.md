# FinalEndPoint

This is the application which is used to help guide FastGit UK users who live in the country which FastGit UK cannot
provide the service to.

## Introduction

Before all, we need figure out how FastGit detects the unavailable regions. Generally, we use the feature called ECS (
EDNS Client Subnet) in Aliyun DNS. Previously, if we find the user who is in the region which is unserviceable, we will
redirect the request to GitHub directly. But now we can redirect to CloudFlare workers.

To convert both of them, we will create a CloudFlare Worker on one website and run an endpoint in original url, which
can help FastGit UK redirect the users to the correct Workers link.

```plain
                         +------------------+
                         | Service EndPoint |
                         +--------+---------+
+------+    +---------+           | if serviceable
| User +----+ DNS ECS +-----------+
+------+    +---------+           | if unserviceable
                          +-------+--------+      +--------------------+
                          | Final EndPoint +------+ CloudFlare Workers |
                          +----------------+      +--------------------+
```
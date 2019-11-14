Recommended Base Images
==========================

For container development I have seen many base images. Everyone has their own preference. Here are few of mine base image recommendation for building your application on top of them.


* Debian 9 (Stretch) launcher.gcr.io/google/debian9 and gcr.io/google-appengine/debian9
* Ubuntu 16.04 (Xenial Xerus) launcher.gcr.io/google/ubuntu16_04 and gcr.io/gcp-runtimes/ubuntu_16_0_4
* Ubuntu 18.04 (Bionic Beaver) launcher.gcr.io/google/ubuntu18_04 and gcr.io/gcp-runtimes/ubuntu_18_0_4
* CentOS 7 launcher.gcr.io/google/centos7 and gcr.io/gcp-runtimes/centos7
* docker pull scratch

There are RedHat Universal Base Image which are also pretty good. I recommend using this if you are in a predominantly redhat eco-system.

UBI 8
=========================
```
registry.access.redhat.com/ubi8/ubi
registry.access.redhat.com/ubi8/ubi-minimal
registry.access.redhat.com/ubi8/ubi-init
```

UBI 7
=========================
```
registry.access.redhat.com/ubi7/ubi
registry.access.redhat.com/ubi7/ubi-minimal
registry.access.redhat.com/ubi7/ubi-init
```

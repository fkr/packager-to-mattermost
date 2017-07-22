# Adapter for Packager.io to Mattermost

This is an Adapter that translates Packager.io webhook packets to the
Mattermost API.

# Data

Typical packager.io Datapacket:

```
Content={
"event":"package","repository_uuid":"74036391",
"repository_slug":"idb-project/the-idb","filename":"the-idb_1.7.0-1500714459.05db224.xenial_amd64.deb",
"commit":"05db224210523c8a7cb3221dbaf393dd8ba4cbea",
"branch":"1.7",
"tag":"1.7.0",
"tagged":false,
"real_tag":null,
"distribution":"ubuntu-16.04",
"package_url":"https://pkgr-production-deb.s3.amazonaws.com/gh/idb-project/the-idb/pool/t/th/the-idb_1.7.0-1500714459.05db224.xenial_amd64.deb",
"upstream_url":"https://github.com/idb-project/the-idb/commit/05db224210523c8a7cb3221dbaf393dd8ba4cbea",
"build_url":"https://packager.io/gh/idb-project/the-idb/build_runs/5423112"}
```

# Author

Felix Kronlage <fkr@hazardous.org>


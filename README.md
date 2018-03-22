### ErpWorkers ERP Oversinerco ###

Required folder structure under gopath is:
`erpOversinerco/erpWorkers/{contentOfThiRepo}`

Required ENV:

`MONGODB_URL`: MongoDB url

`APP_PORT`: Http port where server listen


Required ENV for KUBERNETES CI-CD:

```
__ENV__APP_NAME=workers
__ENV__APP_PORT=50004
__ENV__BRANCH_NAME=master|develop
__ENV__MONGODB_URL=mongo:27017
__ENV__MONGODB_DB_NAME=workers
__ENV__MONGODB_COLLECTION_NAME=workers
__ENV__PUBSUB_TOPIC=erpoversinerco
__ENV__GCE_PROJECT=oversinerco-1
```
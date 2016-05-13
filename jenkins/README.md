Without Git
===========

```bash
cat jenkins/vars/dockerFlowWorkshop.groovy

mkdir -p /data/jenkins/workflow-libs/vars/

cp jenkins/vars/dockerFlowWorkshop.groovy \
    /data/jenkins/workflow-libs/vars/dockerFlow.groovy
```

With Git
========

```bash
cd /tmp

git clone http://10.100.198.200:8080/workflowLibs.git

cd workflowLibs

git checkout -b master

mkdir vars

cp ~/go-demo/jenkins/vars/dockerFlowWorkshop.groovy /tmp/workflowLibs/vars/dockerFlow.groovy

git add --all

git config --global user.name "vfarcic"

git commit -a -m "Docker Flow"

git push --set-upstream origin master
```
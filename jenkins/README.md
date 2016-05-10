```bash
git clone http://10.100.198.200:8080/workflowLibs.git

cd /tmp

git checkout -b master

mkdir vars

cp ~/go-demo/jenkins/vars/dockerFlowWorkshop.groovy /tmp/workflowLibs/vars/dockerFlow.groovy

git add --all

git commit -a -m "Docker Flow"

git push --set-upstream origin master
```
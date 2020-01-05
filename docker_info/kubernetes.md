# kubuernetnes

## 1 安装kubuernetnes

### 1.1 用minikube运行一个本地单节点kubernetes集群

#### 1.1.1 安装kubectl

```javascript
apt-get update && apt-get install -y apt-transport-https
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

apt-get update

apt-get install -y kubelet kubeadm kubectl
```

#### 1.1.2 安装minikube

```shell
Download and install minikube:
curl -Lo minikube https://github.com/kubernetes/minikube/releases/download/v1.5.2/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
sudo mv minikube /usr/bin/minikube

Verify that your system has virtualization support enabled:
egrep -q 'vmx|svm' /proc/cpuinfo && echo yes || echo no

因为是虚拟机启动，直接本机运行
minikube config set vm-driver none

关闭swap
swapoff -a

minikube start --image-mirror-country cn \
    --iso-url=https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.5.1.iso \
    --registry-mirror=https://xxxxxx.mirror.aliyuncs.com
    
报错：
	 Error restarting cluster: waiting for apiserver: apiserver process never appeared
解决：
	重启即可
```

#### 1.1.3 为kubectl配置别名和命令行补齐
```bash
vi ~/.bashrc
#别名
alias k='kubectl'
#命令行补全
#安装bashcompletion 
apt-get install bashcompletion
source /etc/bash_completion
source <(kubectl completion bash | sed s/kubectl/k/g)
```

## 2 在kubernetes上运行第一个应用

### 2.1 介绍pod

一个pod是一组紧密相关的容器，它们总是一起运行在同一个工作节点上，以及同一个linux命名空间中。每个pod就像一个独立的逻辑机器，拥有自己的IP、主机名、进程等，运行一个独立的应用程序。应用程序可以是单个进程，运行在单个容器中，也可以是一个主应用进程或者其他支持进程。每个进程都在自己的容器中运行。一个pod的所有容器都运行在同一个逻辑机器上，而其他pod中的容器，即使运行在同一个工作节点上，也会出现在不同的节点上。

### 2.2 访问Web应用

要让pod能够从外部访问，需要通过服务对象公开它，要创建一个特殊的LoadBalancer类型的服务。通过创建一个外部的负载均衡，可以通过负载均衡的公共IP访问pod。

#### 2.2.1 创建一个服务对象

要创建服务，需要告知kubernetes对外暴露之前创建的ReplicationController:

kubectl expose rc kubia --type=LoadBalancer --name kubia-http

其中rc是ReplicationController的简写，expose意思为暴露

#### 2.2.2 列出服务

通过kubectl get services命令创建新的服务对象

![image-20191226092716690](C:\Users\吕同昕\AppData\Roaming\Typora\typora-user-images\image-20191226092716690.png)

仔细查看kubia-http服务，他还没有外部ip地址，因为kubernetes运行的云基础设施创建负载均衡需要一段时间。负载均衡启动后，应该会显示服务的外部ip地址。

注意：minikube不支持LoadBalancer类型服务，因此服务不会有外部ip。但是可以通过外部端口访问服务

#### 2.2.3 使用外部ip访问服务

curl <外部ip>

注意：使用minikube时，可以运行minikube service kubia-http获取可以访问服务的ip和port

如果仔细观察，会发现应用将pod名称作为了他的主机名。如前所述，每个pod都像是一个独立的机器

###  2.3 系统的逻辑部分

#### 2.3.1 ReplicationController,pod和服务是如何组合在一起的

kubernetes的基础构件是pod。但是，你并没有真的创建出任何pod，至少不是直接创建。通过运行kubectl run命令，创建了一个ReplicationController，它用于创建pod实例。为了使该pod能够从集群外部访问，需要让Kubernetes将该ReplicationController管理的所有pod由一个服务对外暴露。

#### 2.3.2 pod和它的容器

在你的系统中最重要的组件是pod。他只包含一个容器，但是通常一个pod可以包含任意数量的容器。

#### 2.3.3 RepliacionController的角色

它确保始终存在一个运行中的pod实例。通常，ReplicationController用于复制pod并让它们保持运行。

#### 2.3.4 为什么需要服务

系统的第三个组件是服务。pod的存在是短暂的，一个pod可能会在任何时候消失，有可能发生故障，有可能会被人删除。当其中任何一种情况发生时，将会有新的pod替换他。新的pod和之前的ip地址会发生变化。我们需要对外暴露一个不变的地址，这就是为什么需要服务。

### 2.4 水平伸缩应用

使用kubernetes的一个主要好处是可以简单地扩展部署

#### 2.4.1 增加期望的副本数

kubectl scale rc kubia --replicas=3

告诉kubernetes需要确保pod始终有三个实例在运行

我们不会告诉kubernetes应该做什么，而是申明我们的期望状态，然后让kubernetes检查状态是否满足

#### 2.4.2 查看扩容的结果

#### 2.4.3 当切换到服务时请求切换到所有3个pod上

服务作为负载均衡挡在多个pod前面，请求会随机切换到不同的pod上

### 2.5 查看应用运行在哪个节点上

#### 2.5.1 列出pod时显示podIP和pod的节点

kubectl get pods -o wide

#### 2.5.2 使用kubectl describe查看pod的其他细节

### 2.6 介绍kubernetes dashboard

到目前为止，我们只是用了kubectl来探索kubernetes集群。但是kubernetes也提供了一个不错的web dashboard，图形界面。

在minikube中使用minikube dashboard

将在默认浏览器中打开

## 3 pod:运行在kubernetes中的容器

### 3.1 介绍pod

#### 3.1.1 为何需要pod

##### 为何多个容器比单个容器包含多个进程好？

容器被设计为只会将内容按照一个进程进行维护，如果在容器中运行多个不相关的进程，那么这些进程的运行和管理将由我们维护

#### 3.1.2 了解pod

由于不能将多个进程聚集在一个单独的容器中，我们需要另一种更高级的结构将容器绑定在一起，并将它们作为一个单元进行管理，这就是pod背后的根本原理。

##### 同一pod中容器之间的部分隔离

在上一章中，我们已经了解到容器之间彼此是完全隔离的，但此时我们期望的是隔离容器组，而不是单个容器，并让每个容器组内的容器共享一些资源，而不是全部。kubernetes通过配置docker来让一个pod内的所欲容器共享相同的linux命名空间，而不是每个容器都有自己的一组命名空间。

由于一个pod中的所有容器大偶在相同的network和UTS命名空间下运行，所以他们都共享相同的主机名和网络接口。同样的，这些容器也都在相同的IPC命名空间下运行，因此能够通过IPC进行通信。

在文件系统中，我们通过Volume来实现共享文件目录

##### 容器如何共享相同的IP和端口空间

会产生端口冲突

由于pod中所有的容器也都具有相同的loopback网络接口，因此可以通过localhost与同一pod中的其他容器进行通信

##### 介绍平坦pod间网络

#### 3.1.3 通过pod合理管理容器

例如，对于一个由前端应用服务器和后端数据库组成的多层应用程序，你认为应该将其配置为单个pod还是两个pod呢？

##### 将多层应用分散到多个pod中

将多层应用分散到多个pod中，尽可能利用多个节点的资源

##### 基于扩缩容考虑而分割到多个pod中

pod是扩缩容的基本单位，这意味着如果扩大pod的实例数量，就会有两个前端应用服务器和两个后端数据库服务器了

##### 如何在pod中使用多个容器

将多个容器添加到单个pod的主要原因是应用可能是由一个主进程和一个或者多个辅助进程组成

##### 决定何时在pod中使用多个容器

- 它们需要一起运行还是可以在不同的主机上运行
- 它们代表的是一个整体还是相互独立的组件
- 它们必须一起进行扩缩容还是可以分别进行

### 3.2 以YAML或JSON描述文件创建pod

pod个其他kuberbetes资源通常是通过向kubernetes REST API提供json或者YAML描述文件来创建的。此外还有其他更简单的创建资源的方法，比如在前一章使用的kubectl run 命令。通过YAML文件定义所有的Kubernetes对象之后，还可以将它们存储在版本控制系统中，还充分利用版本控制所带来的便利性。

为了配置每种类型资源的各种属性，我们需要了解并理解kubernetes API对象定义。通过本书学习各种资源类型时，我们将会了解其中的大部分内容。需要注意的是，我们不会解释每一个独立属性，因此在创建对象时还应参考http://kubernetes.io/docs/reference/ 中的参考文档.

#### 3.2.1 检查现有pod的YAML描述文档

kubectl get pod (podname) -o yaml

```yaml
apiVersion: v1		//kubernets API版本
kind: Pod			//资源类型
metadata:			//pod元数据（名称，标签和注解等）
  creationTimestamp: "2019-12-25T10:42:29Z"
  generateName: kubia-
  labels:
    run: kubia
  name: kubia-xbdtq
  namespace: default
  ownerReferences:
  - apiVersion: v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicationController
    name: kubia
    uid: 657bc8cc-38f6-46b5-83e7-8a61a66f5577
  resourceVersion: "201744"
  selfLink: /api/v1/namespaces/default/pods/kubia-xbdtq
  uid: 73859157-d800-47b7-a335-6e09d0a33a7f
spec:				//pod规格/内容（pod的容器列表、volume等）
  containers:
  - image: luksa/kubia
    imagePullPolicy: Always
    name: kubia
    ports:
    - containerPort: 8080
      protocol: TCP
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-t64bg
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: minikube
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-t64bg
    secret:
      defaultMode: 420
      secretName: default-token-t64bg
status:							//pod及其内部容器的详细状态
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2019-12-25T10:42:30Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2019-12-27T09:28:34Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2019-12-27T09:28:34Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2019-12-25T10:42:30Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://6b9f28408ad19e583de2771a976a32aa9378139a08f89ca261ef1c6545ec8870
    image: luksa/kubia:latest
    imageID: docker-pullable://luksa/kubia@sha256:3f28e304dc0f63dc30f273a4202096f0fa0d08510bd2ee7e1032ce600616de24
    lastState:
      terminated:
        containerID: docker://dda6aac95bbdda8a286623f43b1e6433bb7c52e27bffbc4a02a8f71b07f92f9e
        exitCode: 255
        finishedAt: "2019-12-27T09:18:52Z"
        reason: Error
        startedAt: "2019-12-25T10:44:32Z"
    name: kubia
    ready: true
    restartCount: 1
    started: true
    state:
      running:
        startedAt: "2019-12-27T09:28:33Z"
  hostIP: 192.168.2.149
  phase: Running
  podIP: 172.17.0.4
  podIPs:
  - ip: 172.17.0.4
  qosClass: BestEffort
  startTime: "2019-12-25T10:42:30Z"

```

##### 介绍pod定义的主要部分

pod定义由这么几个部分组成：首先是YAML中使用的Kubernetes API版本和YAML描述的资源类型；其次是几乎在所有kubernetes资源中都可找到的三大重要部分：

- metadata 包括名称、命名空间、标签会和关于该容器的其他信息
- spec 包含pod内容的实际说明，例如pod的容器、卷和其他数据
- status包含运行中的pod的当前信息，例如pod所处的条件，每个容器的描述和状态

#### 3.2.2 为pod创建一个简单的YAML描述文件

我们将创建一个名为kubia-manual.yaml的文件

```yaml
#描述文件遵循v1版本的kubernetes API
apiVersion: v1
#我们描述的资源是一个Pod
kind: Pod

metadata:
        #pod的名称
        name: kubia-manual
spec:
        containers:
                #创建容器所用的镜像？
                - image: luksa/kubia
                  #容器的名称
                  name: kubia
                  #容器监听的端口
                  ports:
                          - containerPort: 8080
                            protocol: TCP
```

##### 指定容器端口

在pod定义中指定端口纯粹是展示性的。忽略它们对于客户端是否可以通过端口连接到pod不会带来任何影响。

#### 3.2.3 使用kubctl create来创建pod

我们使用kubectl -f kubia-manual.yaml命令从YAML创建pod：

##### 得到运行中pod的完整定义

kubectl get pod (name) -o yaml

kubectl get pod (name) -o json

##### 在pod列表中查看新创建的pod

kubectl get pods

#### 3.2.4 查看应用程序日志

小型Node.js应用将日志记录到进程的标准输出。容器化的应用程序通常会将日志记录到标准输出和标准错误流，而不是将其写入文件，这就允许用户可以简单、标准的方式查看不同应用程序的日志

容器运行时将这些流重定向到文件，并允许我们运行以下的命令来获取容器的日志:

以docker为例

docker logs <container id>

##### 使用kubectl logs命令获取pod日志

为了查看pod的日志，只需要在本地机器上运行以下命令：

kubectl logs (name)

**注意** 每天或者每次日志文件达到10MB大小时。容器日志都会自动轮替。kubectl pods命令仅显示最后一次轮替后的日志条目

##### 获取多容器pod日志时指定容器名称

如果我们的pod包含读个容器，在运行kubectl logs命令时则必须要通过kubectl logs (pod_name) -c kubia

容器名称则在YAML文件中指定过了，获取很容易

这里我们只能获取仍然存在的pod的日志。

#### 3.2.5 向pod发送请求

前面提到了是用kubectl expose方法创建一个service以便在外部访问该pod。这次我们是用端口转发进行通信

##### 将本地网络端口转发到pod中的端口

kubectl port-forward (pod_name) 8888:8080

将本地端口8888转发到我们的pod的8080端口

##### 通过端口转发连接到pod

curl localhost:8888

### 3.3 使用标签组织pod

通过标签来组织pod和所有其他kubernetes对象

#### 3.3.1 介绍标签

标签是一种简单却功能强大的kubernetes特性，不仅可以组织pod，也可以组织所有其他的kubernetes资源。

#### 3.3.2 创建pod时指定标签

```yaml
#描述文件遵循v1版本的kubernetes API
apiVersion: v1
#我们描述的资源是一个Pod
kind: Pod

metadata:
        #pod的名称
        name: kubia-manual-v2
        #附加两个标签
        labels:
                creation_method: manual
                env: prod
spec:
        containers:
                #创建容器所用的镜像？
                - image: luksa/kubia
                  #容器的名称
                  name: kubia
                  #容器监听的端口
                  ports:
                          - containerPort: 8080
                            protocol: TCP
```

使用--show-labels选项来展示标签

kubectl get pods --show-labels

使用-L选项指定感兴趣的标签

kubectl get pods -L (label_name)

#### 3.3.3 修改现有pod的标签

新增标签

kubectl label pod (pod_name) env=v1

修改现有标签值

kubectl label pod (pod_name) env=v2 --overwrite

### 3.4 通过标签选择器列出pod子集

标签要与标签选择器结合在一起。

标签选择器根据资源的以下条件来选择资源：

- 包含使用特定键的标签
- 包含具有特定键和值的标签
- 包含具有特定键的标签，但是值与我们指定的不同

#### 3.4.1 使用标签选择器列出pod

kubectl get pod -l creation=manual

kubectl get pod -l env

kubectl get pod -l '!env'

kubectl get pod -l env!=v1

kubectl get pod -l env in (v1,v2)

kubectl get pod -l env notin (v1,v2)

#### 3.4.2 在标签选择器中使用多个条件

多个条件使用逗号分隔

### 3.5 使用标签和选择器来约束pod调度

如果你想对一个pod应该调度到哪里拥有发言权，那就不应该指定一个确切的节点，而应该用每种方式描述对节点的的需求，使kubernetes选择一个符合这些需求的节点。这恰恰可以通过节点标签和节点标签选择器完成

#### 3.5.1 使用标签分类工作节点

如前所述，pod并不是唯一可以附加标签的资源。节点也是资源之一。

假设我们集群中的一个节点刚添加完成，它包含一个用于通用GPU计算的GPU。我们希望向节点添加便签来展示这个功能特性，可以通过将标签gpu=true添加到其中一个节点来实现

kubectl label node (node_name) gpu=rue

#### 3.5.2 将pod调度到特定节点

```yaml
#描述文件遵循v1版本的kubernetes API
apiVersion: v1
#我们描述的资源是一个Pod
kind: Pod

metadata:
        #pod的名称
        name: kubia-gpu
spec:
		#节点调度条件
        nodeSelector:
                gpu: "true"
        containers:
                #创建容器所用的镜像？
                - image: luksa/kubia
                  #容器的名称
                  name: kubia
                  #容器监听的端口
                  ports:
                          - containerPort: 8080
                            protocol: TCP
```

### 3.6 注解pod

除标签外，pod和其他对象还可以包含注解

可以理解为注释

#### 3.6.1 查找对象的注解

为了查看注解，我们需要获取pod的完整YAML文件或者使用kubernetes describe命令

Annotations字段

#### 3.6.2 添加和修改注解

使用kubectl annotate pod (pod_name) version="v1"

注意保证键值不冲突，因为和标签不同，注解会被这个命令直接覆盖

### 3.7 使用命名空间对资源进行分组

#### 3.7.1 了解对命名空间的需求

#### 3.7.2 发现其他命名空间及其pod

首先，让我们列出集群中的所有命名空间

kubernetes get ns

到目前为止，我们只在default命名空间中进行操作，因为我们没有指定特定的命名空间，所以默认使用default命名空间

kubectl get pod --namespace kube-system

除了隔离资源，命名空间还用于允许某些用户访问某些特定资源

#### 3.7.3 创建一个命名空间

命名空间也是kubernetes的资源，因此也可以通过yaml文件进行创建

##### 从YAML文件创建命名空间

```yaml
apiVersion: v1
kind: Namespace
metadata:
        name: custom-namespace
```

##### 使用kubectl create namespace命令创建命名空间

#### 3.7.4 管理其他命名空间中的对象

如果想要在刚创建的命名空间中创建资源，需要在yaml文件中的metadata中添加namespace: value。也可以在创建时指定命名空间

kubectl create -f (yaml_name) -n (ns_name)

提示 想要快速切换到不同的命名空间，可以设置以下别名：

alias kcd='kubectl config set-context $(kubectl config current-context) --namespace'

然后使用kcd some-namespace在命名空间之间进行切换

#### 3.7.5 命名空间提供的隔离

命名空间提供的隔离是管理的隔离，对实际资源并没有进行隔离

### 3.8 停止和移除pod

#### 3.8.1 按名称删除pod

kubectl delete pod (pod_name) (pod_name1)

通过向进程发送SIGTERM信号来实现关闭进程，如果没有及时关闭，再发送SIGKILL

#### 3.8.2 使用标签选择器删除pod

kubectl delete pod -l ()

#### 3.8.3 通过删除整个命名空间来删除pod

kubectl delete ns (ns_name)

#### 3.8.4 删除命名空间中的所有pod，但保留命名空间

kebectl delete pod --all

#### 3.8.5 删除命名空间中的（几乎）所有资源

kubectl delete all --all

## 4 副本机制和其他控制器：部署托管的pod

### 4.1 保持pod健康

使用kubernetes的一个主要好处是，可以给其一个容器列表来由其保持容器在集群中的运行。可以通过让kubernetes创建pod资源，为其选择一个工作节点并在该节点上运行该pod的容器来完成此操作

但是如果其中一个容器终止，怎么办？

#### 4.1.1 介绍存活探针

kubernetes可以通过存活探针检查容器是否还在运行。可以为pod中的每个容器单独指定存活探针。如果探测失败，kubernetes将定期执行探针并重新启动容器。

kubernetes有以下3中探测容器的机制：

HTTP GET探针针对容器的IP地址执行HTTP GET请求

TCP连接

Exec探针在容器内执行任意命令，并检查命令的退出状态码

#### 4.1.2 创建HTTP的存活探针

```yaml
metadata:
        name: kubia-liveness
spec:   
        containers:
                - image: luksa/kubia-unhealthy
                  name: kubia
                  livenessProbe:
                          httpGet:
                                  path: /
                                  port: 8080
```

#### 4.1.3 使用存活探针

kubectl logs (pod_name) --previous

--previous选项可以查看前一个容器的日志

kubectl describe pod (pod_name) 可以查看错误的描述

#### 4.1.4 配置存活探针的附加属性

```yaml
apiVersion: v1
kind: Pod
metadata:
        name: kubia-liveness
spec:
        containers:
        - image: luksa/kubia-unhealthy
          name: kubia
          livenessProbe:
                  httpGet:
                          path: /
                          port: 8080
                  #kuberbetes会在第一次探测前等待15秒
                  initialDelaySeconds: 15
```

现实中有很多容器启动失败是因为探测延迟时间过短造成的，在这段时间内容器并没有做好接受请求的准备就被迫关闭了

#### 4.1.5 创建有效的存活探针

对于在生产中运行的pod，一定要定义一个存活探针。没有探针的话kubernetes无法知道你的应用是否还活着。只要进程还在运行，kubernetes会认为容器是健康的。

##### 存活探针应该检查什么

简单的存活探针仅仅检查了服务器是否响应。

为了更好地进行存活检查，需要将探针配置为请求特定的URL路径例如/health，并让应用从内部对内部运行的所有重要组件执行状态检查，以确保他们都没有终止或停止响应。

##### 保持探针轻量

存活探针不应消耗太多计算资源，并且运行不应该花太长时间。

##### 无须在探针中实现重试循环

##### 存活探针小结

kubernetes会在你的容器崩溃或其存活探针失败时，通过重启容器来保持运行。这项任务由承载pod的节点上的kubelet执行，主服务器上的KCP不会参与此过程。这意味着如果节点本身崩溃，那么该pod将不能重启。为了确保你的应用程序在另一个节点上重新启动，需要使用ReplicationController

### 4.2 了解ReplicationController

ReplicationController是一种kubernetes资源，可确保他的pod始终保持运行状态。

如果pod因任何原因消失， 则ReplicationController会注意到缺少pod并创建替代pod。

#### 4.2.1 ReplicationController的操作

ReplicationController会持续监控正在运行的pod列表，并保证相应“类型”pod的数目与期望相符。

pod的类型说法其实是不正确的，ReplicationController其实是根据pod是否匹配某个标签选择器来执行操作的

##### 介绍控制器的协调流程

ReplicationController的工作时确保pod的数量始终与其标签选择器匹配。如果不匹配，则ReplicationController将根据所需，采取适当的操作来协调pod的数量

##### 了解ReplicationController的三部分

- label selector（标签选择器）
- replica count（副本个数）
- pod template（pod 模板）

##### 更改控制器的标签选择器或pod模板的效果

更改标签选择器和pod模板对现有的pod没有影响。更改标签选择器会使现有的pod脱离ReplicationController的范围，因此控制器会停止关注他们

##### 使用ReplicationController的好处

尽管是一个令人难以置信的简单概念，却提供或启用了以下强大功能：

- 确保一个pod（或多个pod副本）持续运行
- 集群节点发生故障时，它将为故障节点上运行的所有pod创建替代副本
- 能轻松实现pod的水平伸缩

#### 4.2.2 创建一个ReplicationController

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
        name: kubia
spec:
        replicas: 3
        #让kubernetes自己从模板中提取选择器使用的标签
        #selector:
        #        app: kubia
        template:
                metadata:
                        labels:
                                app: kubia
                spec:
                        containers:
                                - name: kubia
                                image: luksa/kubia
                                ports:
                                        - containerPort: 8080
```



#### 4.2.3 使用ReplicationController

##### 查看ReplicationController对已删除的pod的响应

kubectl delete pod (pod_name)

kubectl get pods

##### 获取有关ReplicationController的信息

kubectl get rc

kubectl describe rc (rc_name)

##### ReplicationController如何创建新的pod

控制器通过创建一个新的替代pod来响应pod的删除操作。从技术上讲，它并没对删除本身作出反应，而是针对由此产生的状态——pod数量不足。

##### 应对节点故障

#### 4.2.4 将pod移入或者移出ReplicationController的作用域

由ReplicationController创建的pod并不是绑定到ReplicationController的作用域。在任何时候，ReplicationController管理与标签选择器匹配的pod。通过更改pod的标签，可以将它从ReplicationController的作用域中添加或删除。它甚至可以从一个ReplicationController中移动到另一个。

如果你更改了一个pod的标签，使它不在与ReplicationController的标签选择器相匹配，那么改pod就会变得和其他手动创建的pod一样了。

##### 给ReplicationController管理的pod加标签

kubectl label pod (pod_name) type=special

如果你向ReplicationController管理的pod添加其他标签，它并不关心

##### 更改已托管的pod的标签

kubectl label pod (pod_name) app=foo --overwrite

##### 从控制器删除pod

如果想要操作特定的pod时，将其从ReplicationController的管理范围中移除，然后再删除

##### 更改ReplicationController的标签选择器

#### 4.2.5 修改pod模板

ReplicationController的pod模板可以随时修改，更改pod模板只会影响你之后创建的pod，而不会影响现有的pod。

kubectl edit rc (rc_name)

#### 4.2.6 水平缩放pod

#### ReplicationController扩缩容

1. 通过scale命令

kubectl scale rc (rc_name) --replicas=10

2. 通过编辑yaml文件

kubectl edit rc (rc_name)

#### 4.2.7 删除一个ReplicationController

当你通过kubectl delete 删除ReplicationController时，pod也会被删除。

如果想在删除ReplicationController时保持pod的运行

kubectl delete rc (rc_name) --cascade=false

### 4.3 使用ReplicaSet而不是ReplicationController

最初，ReplicationController是用于复制和在异常时重新调度节点的唯一kubernetes组件，后来引入了一个名为ReplicaSet的类似资源，他是新一代的ReplicationController。

#### 4.3.1 比较ReplicaSet和ReplicationController

虽然两者的行为相同，但是ReplicaSet的pod选择器的表达能力更强。虽然ReplicationController的标签选择器只允许包含某个标签的匹配pod，ReplicaSet的选择器还允许匹配缺少某个标签的pod，或包含特定标签名的pod不管其值为何

另外ReplicationController无法同事匹配两组标签，但是ReplicaSet可以

#### 4.3.2 定义ReplicaSet

```yaml
#ReplicaSet所属版本
apiVersion: apps/v1
kind: ReplicaSet
metadata:
        name: kubia
spec:
        replicas: 3
        selector:
                matchLabels:
                        app: kubia
        template:
                metadata:
                        labels:
                                app: kubia
                spec:
                        containers:
                                - name: kubia
                                  image: luksa/kubia

```

#### 4.3.3 创建和检查ReplicaSet

#### 4.3.4 使用ReplicaSet的更富表现力的标签选择器

```yaml
#ReplicaSet所属版本
apiVersion: apps/v1
kind: ReplicaSet
metadata:
        name: kubia
spec:
        replicas: 3
        selector:
                matchExpressions:
                		#每个表达式必须包含一个key，operator，可能有一个value列表
                        - key: app
                        #In表示必须是values中的一个
                        #NotIn表示必须不是values中的任意一个
                        #Exists表示pod必须含有一个指定key的标签，值不重要，这时不应该由value列表
                        #DoesNotExist表示pod不能包含指定key的标签，value不得指定
                          operator: In
                          values:
                                  - kubia
        template:
                metadata:
                        labels:
                                app: kubia
                spec:
                        containers:
                                - name: kubia
                                  image: luksa/kubia
```

#### 4.3.5 ReplicaSet小结

### 4.4 使用DaemonSet在节点上运行pod

使用ReplicationController和ReplicaSet都用于在kubernetes集群上运行部署特定数量的pod。但是，当你希望pod在集群的每个节点上运行时（并且每个节点都需要正好一个运行的pod实例），就会出现某些情况。

#### 4.4.1 使用DaemonSet在每个节点上运行一个pod

#### 4.4.2 使用DaemonSet只在特定的节点上运行pod

DaemonSet将pod部署到集群中的所有节点上，除非指定这些pod只在部分节点上运行。这是通过pod模板中的nodeSelector属性指定的，这是DaemonSet定义的一部分。

##### 用一个例子来解释DaemonSet

##### 创建一个DaemonSet YAML定义文件

````yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
        name: ssd-monitor
spec:
        selector:
                matchLabels:
                        app: ssd-monitor
        template:
                metadata:
                        labels:
                                app: ssd-monitor
                spec:
                        nodeSelector:
                                disk: ssd
                        containers:
                                - name: main
                                  image: luksa/ssd-monitor
````



##### 创建DaemonSet

创建一个DaemonSet就像从YAML文件创建资源那样

##### 向节点上添加所需的标签

kubectl get node

kubectl label node (node_name) disk=ssd

##### 从节点上删除所需的标签

重命名标签值就可以做到

### 4.5 运行执行单个任务的pod

目前为止，我们只讨论了需要持续运行的pod。你会遇到只想运行完成工作后就终止任务的情况

#### 4.5.1 介绍Job资源

#### 4.5.2 定义Job资源

```yaml
apiVersion: batch/v1
kind: Job
metadata:
        name: batch-job
spec:
        template:
                metadata:
                        labels:
                                app: batch-job
                spec:
                		#设置重启策略
                        restartPolicy: OnFailure
                        containers:
                                - name: main
                                  image: luksa/batch-job
```

#### 4.5.3 看Job运行一个pod

kubectl get jobs

kubectl get pods

完成后pod未被删除的原因是允许你查看其日志

kubectl logs (pod_name)

#### 4.5.4 在Job中运行多个pod实例

作业可以配置为创建多个pod实例，并以并行或者串行方式运行它们。这是通过在Job设置中设置completions和parallelism属性来完成的

##### 顺序运行Job pod

如果你需要一个Job运行多次，则可以将completions设为你希望作业的pod运行多少次

```yaml
apiVersion: batch/v1
kind: Job
metadata:
        name: multi-completion-batch-job
spec:
		#将有5个pod完成任务
        completions: 5
        template:
                metadata:
                        labels:
                                app:
                                        batch-job
                spec:
                        restartPolicy: OnFailure
                        containers:
                                - name: main
                                  image: luksa/batch-job
```

##### 并行运行Job pod

通过属性parallelism进行配置

##### 限制Job pod的完成任务的时间

通过属性activeDeadlineSeconds进行配置

### 4.6 安排Job定期运行或在将来运行一次

CronJob

在配置的时间，kubernetes将根据在CronJob对象中配置的Job模板创建CronJob资源。创建Job资源时，将根据任务的pod模板创建并启动一个或者多个pod副本。

#### 4.6.1 创建一个CronJob

##### 配置时间表安排

##### 配置Job模板

#### 4.6.2 了解计划任务的运行方式

## 5 服务：让客户端发现pod并与之通信

### 5.1 介绍服务

服务是一种为一组功能相同的pod提供单一不变的接入点的资源。

#### 结合实例解释服务

#### 5.1.1 创建服务

服务器的后端可以不知一个pod。服务的连接对所有的后端pod是负载均衡的。

在前面的章节中通过创建rc运行了三个node.js应用的pod。现在我们通过这三个pod进行演示，为这三个pod创建一个服务

##### 通过kubectl expose创建服务

创建服务的最简单方法是通过kubectl expose

##### 通过YAML描述文件来创建服务

```yaml
apiVersion: v1
kind: Service 
metadata:
        name: kubia
spec:           
        ports:    
                - port: 80
                  targetPort: 8080
        #具有app=kubia标签的pod都属于该服务
        selector:
                app: kubia
```

##### 检测新的服务

kubectl get svc

##### 从内部集群测试服务

- 在运行的容器中远程执行命令

可以使用kubectl exec命令远程在一个已存在的pod容器上执行任何命令。

用kubectl get pod命令列出所有的pod，并且选择其中一个作为exec命令的执行目标。

kubectl exec kubia-626jj -- curl -s http://10.111.249.153

##### 配置服务上的会话亲和性

如果多次执行同样的命令，每次调用执行应该在不同的pod上。因为服务代理通常将每个连接随机指向选中的后端pod中的一个，即使连接来自于同一个客户端。

另一方面，如果希望特定客户端产生的所有请求每次都指向同一个pod，可以设置服务的sessionAffinity属性为ClientIP。这种方式会使服务代理把来自同一个clientIP的所有请求每次都指向同一个pod。

##### 同一个服务暴露多个端口

创建的服务可以暴露一个端口，也可以暴露多个端口。

比如，你的pod监听两个端口，比如HTTP监听8080，HTTPS监听8443，可以使用一个服务从端口80和443转发至8080和8443.在这种情况下，无需创建两个不同的服务。

```yaml
apiVersion: v1
kind: Service
metadata:
        name: kubia
spec:
        sessionAffinity: ClientIP
        ports:
                - name: http
                  port: 80
                  targetPort: 8080
                - name: https
                  port: 443
                  targetPort: 8443
        #具有app=kubia标签的pod都属于该服务
        selector:
                app: kubia
```

##### 使用命名的端口

```yaml
#在pod中命名端口
apiVersion: v1
kind: Pod
spec:
	containers:
	- name: kubia
	  ports:
	  - name: http
	    containerPort: 8080
	  - name: https
	    containerPort: 8443
```

```yaml
#在服务中使用命名的端口
apiVersion: v1
kind: Service
metadata:
        name: kubia
spec:
        sessionAffinity: ClientIP
        ports:
                - name: http
                  port: 80
                  targetPort: http
                - name: https
                  port: 443
                  targetPort: https
        #具有app=kubia标签的pod都属于该服务
        selector:
                app: kubia
```

好处是即使更改pod的端口，服务也不需要修改

#### 5.1.2 服务发现

通过创建服务，现在就可以通过一个单一稳定的IP地址访问到pod。在服务整个生命周期内这个地址保持不变。

但是客户端pod如何才能发现服务呢。

##### 通过环境变量发现服务

在pod开始运行的时候，kubernetes会初始化一系列的环境变量指向现在存在的服务。如果你创建的服务早于客户端pod的创建，pod上的进程可以通过环境变量获取服务的IP地址和端口号

kubectl exec (pod_name) -- env

##### 通过DNS发现服务

pod是否使用内部的DNS服务器是根据pod中spec的dnsPolicy属性来决定的。

每个服务从内部DNS服务器中获得一个DNS条目，客户端的pod在知道名称的情况下可以通过全限定域名来访问，而不是诉诸于环境变量。

##### 通过FQDN连接服务

尝试一下通过FQDN来代替IP去访问kubia服务

在pod容器中运行shell

kubectl exec -it (pod_name)

##### 无法ping通服务IP的原因

这是因为服务的集群IP是一个虚拟IP并且只有与服务端口结合时才有意义

### 5.2 连接集群外部的服务

将服务重定向到外部IP和端口。在集群中运行的客户端可以像连接到内部服务一样连接到外部服务。

#### 5.2.1 介绍服务endpoint

服务并不是和pod直接相连的。相反，有一种资源介于两者之间，它就是Endpoint资源。

kubectl describe svc (svc_name)

```yaml
Name:              kubia
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=kubia
Type:              ClusterIP
IP:                10.96.223.83
Port:              http  80/TCP
TargetPort:        8080/TCP
#代表服务endpoint的pod的IP和端口列表
Endpoints:         172.17.0.11:8080,172.17.0.12:8080,172.17.0.13:8080
Port:              https  443/TCP
TargetPort:        8443/TCP
Endpoints:         172.17.0.11:8443,172.17.0.12:8443,172.17.0.13:8443
Session Affinity:  ClientIP
Events:            <none>
```

kubectl get endpoints (endpoint_name)

尽管在spec服务中定义了pod选择器，但在重定向传入连接时不会直接使用它。相反，选择器用于构建IP和端口列表，然后存储在Endpoint资源中。当客户端连接到服务时，服务代理选择这些IP和端口队中的一个，并将传入连接重定向到在该位置监听的服务器。

#### 5.2.2 手动配置服务endpoint

服务的endpoint与服务解耦后，可以分别手动配置和更新他们。

如果创建了不包含pod选择器的服务，kubernetes将不会创建endpoint资源。这样就需要创建Endpoint资源来指定该服务的endpoint表。

##### 创建没有选择器的服务

首先为服务创建一个YAML文件

```yaml
apiVersion: v1
kind: Service
metadata:
        name: external-service
spec:
		#服务中没有定义pod选择器
        ports:
        - port: 80
```

##### 为没有选择器的服务创建endpoint资源

```yaml
apiVersion: v1
kind: Endpoints
metadata:
        name: external-service
subsets:
        - addresses:
          - ip: 11.11.11.11
          - ip: 22.22.22.22
          ports:
          - port: 80
```

#### 5.2.3 为外部服务创建别名

除了手动配置服务的Endpoint来代替公开外部服务的方法，有一种更简单的方法。就是通过其完全限定域名FQDN来访问外部服务

##### 创建ExternalName类型的服务

要创建一个具有别名的外部服务的服务时，要创建服务资源的一个type字段设置为ExternalName。

### 5.3 将服务暴露给外部客户端

有几种方式可以在外部访问服务：

- 将服务的类型设置成NodePort——每个集群节点都会在节点上打开一个端口，对于NodePort服务，每个集群节点在节点本身上打开一个端口，并将在该端口上接收到的流量重定向到基础服务。该服务仅在服务集群IP和端口上才可以访问，但也可以通过所有节点上的专用端口访问。
- 将服务的类型设置成LoadBalance，这使得服务可以通过一个专用的负载均衡器来进行访问。
- 创建一个Ingress资源，通过一个IP地址公开多个服务——它运行在HTTP层上可以提供更多的功能

#### 5.3.1 使用NodePort类型的服务

将一组pod公开给外部客户端的第一种方法是创建一个服务并将其类型设置为NodePort。通过创建NodePort服务，可以让kubernetes在其所有节点上保留一个端口（所有节点上都是用同一个端口），并将传入的连接转发给作为服务部分的pod。

当尝试和NodePort服务交互时，意义更加重大。

##### 创建NodePort类型的服务

```yaml
apiVersion: v1
kind: Servuce
metadata:
        name: kubia-nodeport
spec:
		#设置服务类型为NodePort
        type: NodePort
        ports:
        		#服务集群ip的端口号
                - port: 80
                  #背后pod的目标端口号
                  targetPort: 8080
                  #通过集群节点的30123端口可以访问该服务
                  nodePort: 30123
        selector:
                app: kubia
```

##### 查看NodePort类型的服务

```yaml
Name:                     kubia-nodeport
Namespace:                default
Labels:                   <none>
Annotations:              <none>
Selector:                 app=kubia
Type:                     NodePort
IP:                       10.96.234.68
Port:                     <unset>  80/TCP
TargetPort:               8080/TCP
NodePort:                 <unset>  30123/TCP
Endpoints:                172.17.0.11:8080,172.17.0.12:8080,172.17.0.13:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

更改防火墙规则，允许外界连接到暴露的端口

#### 5.3.2 通过负载均衡器将服务暴露出来

在云提供商上运行的kubernetes集群通常支持从云基础架构自动提供负载均衡器。所有需要做的就是设置服务类型为Load Badancer而不是NodePort。

#### 5.3.3 了解外部连接的特性

##### 了解并防止不必要的网络跳数

##### 记住客户端IP是不记录的

### 5.4 通过Ingerss暴露服务


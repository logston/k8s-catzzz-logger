# Kubernetes Custom Controller - Custom Resource Handling

**Note**: the source code is _verbosely_ commented, so the source is meant to
be read and to teach

## What is this?

An example of a custom Kubernetes controller that's only purpose is to watch
for the creation, updating, or deletion of all custom resource of type
`CatzzzLogger` (in the all namespaces). This was created as an exercise to
understand how Kubernetes controllers work and interact with the cluster
and resources.

## Running

```
$ git clone https://github.com/trstringer/k8s-controller-custom-resource
$ cd k8s-controller-custom-resource $ go run *.go
```

## Inspecting resources in the handler

You are welcome to dump the resources themselves in handler but logging would
be extremely verbose (and not interactive). I recommend you use a debugger...

```
$ dlv debug
(dlv) b main.ObjectCreated
(dlv) c
```

You can then trigger an event by creating a deployment of nginx...

``` $ kubectl apply -f example/example-catzzzlogger.yaml ```

The breakpoint should be hit and you can analyze in the debugger the object...

``` (dlv) p obj```


```
ROOT_PACKAGE="github.com/logston/k8s-catzzz-logger" ./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "catzzzlogger:v1"
```

```
INFO[0114] Controller.runWorker: processing next item   
INFO[0114] Controller.processNextItem: start            
INFO[0151] Add myresource: default/example-catzzzlogger 
INFO[0151] Controller.processNextItem: object created detected: default/example-catzzzlogger 
INFO[0151] CatzzzLoggerHandler.ObjectCreated            
INFO[0151] 
/}_{\           /.-'
( a a )-.___...-'/
==._.==         ;
     \ i _..._ /,
     {_;/   {_// 
INFO[0151] 
/}_{\           /.-'
( a a )-.___...-'/
==._.==         ;
     \ i _..._ /,
     {_;/   {_// 
INFO[0151] 
/}_{\           /.-'
( a a )-.___...-'/
==._.==         ;
     \ i _..._ /,
     {_;/   {_// 
INFO[0151] Controller.runWorker: processing next item   
INFO[0151] Controller.processNextItem: start
```

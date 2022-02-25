#!/bin/bash
set -e
cd $(dirname $0)

SERVICE_NODE=$1
SERVICE_NODE=${SERVICE_NODE:=node.json}
echo $SERVICE_NODE
# id, address, checks.http should map with existing service nodes
curl -XPUT localhost:8500/v1/agent/service/register -H "Content-Type:application/json" --data-binary @${SERVICE_NODE}
curl http://localhost:8500/v1/catalog/service/node-exporter | jq .
# 查询指定节点以及指定的服务信息
#[root@iZ2zejaz33icbod2k4cvy6Z ~]# curl http://182.92.219.202:8500/v1/catalog/service/redis
#[{"ID":"9d76becb-c557-e605-de13-a906ef32497c","Node":"ndoe5","Address":"172.20.0.6","Datacenter":"dc1","TaggedAddresses":{"lan":"172.20.0.6","lan_ipv4":"172.20.0.6","wan":"172.20.0.6","wan_ipv4":"172.20.0.6"},"NodeMeta":{"consul-network-segment":""},"ServiceKind":"","ServiceID":"redis","ServiceName":"redis","ServiceTags":["service"],"ServiceAddress":"182.92.219.202","ServiceTaggedAddresses":{"lan_ipv4":{"Address":"182.92.219.202","Port":9121},"wan_ipv4":{"Address":"182.92.219.202","Port":9121}},"ServiceWeights":{"Passing":1,"Warning":1},"ServiceMeta":{},"ServicePort":9121,"ServiceEnableTagOverride":false,"ServiceProxy":{"MeshGateway":{},"Expose":{}},"ServiceConnect":{},"CreateIndex":458,"ModifyIndex":458}][root@iZ2zejaz33icbod2k4cvy6Z ~]# 


#删除指定服务 redis为要删除服务的id
#curl -X PUT  http://localhost:8500/v1/agent/service/deregister/redis

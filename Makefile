clickhouse-init:
	docker run -d --name clickhouse --network k8s-hub-network -p 8123:8123 -p 9000:9000 --ulimit nofile=262144:262144 \
		-e CLICKHOUSE_USER=${CLICKHOUSE_USER} \
        -e CLICKHOUSE_PASSWORD=${CLICKHOUSE_PASSWORD} \
        -e CLICKHOUSE_DB=${CLICKHOUSE_DATABASE} \
        yandex/clickhouse-server

kafka-ui-init:
	docker run -d --name kafka-ui --network k8s-hub-network -p 8090:8080 \
        -e KAFKA_CLUSTERS_0_NAME=local \
        -e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka:9092 \
        -e KAFKA_CLUSTERS_0_ZOOKEEPER=zookeeper:2181 \
        provectuslabs/kafka-ui:latest

kafka-local-init:
	docker run -d --name kafka-local --network k8s-hub-network -p 9092:9092 \
		-e KAFKA_ZOOKEEPER_CONNECT=zookeeper-local:2181 \
		-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092  \
		-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
		-e KAFKA_BROKER_ID=1 \
		-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
		-e KAFKA_LOG_DIRS=/var/lib/kafka/data \
		-e KAFKA_CREATE_TOPICS="k8s-hub:3:1" \
		confluentinc/cp-kafka:latest

kafka-init:
	docker run -d --name kafka --network k8s-hub-network -p 9093:9092 \
		-e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
		-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092  \
		-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
		-e KAFKA_BROKER_ID=1 \
		-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
		-e KAFKA_LOG_DIRS=/var/lib/kafka/data \
		-e KAFKA_CREATE_TOPICS="k8s-hub:3:1" \
		confluentinc/cp-kafka:latest

zookeeper-local-init:
	docker run -d --name zookeeper-local --network k8s-hub-network -p 2181:2181 \
		-e ZOOKEEPER_CLIENT_PORT=2181 \
		confluentinc/cp-zookeeper:latest

zookeeper-init:
	docker run -d --name zookeeper --network k8s-hub-network -p 2182:2181 \
		-e ZOOKEEPER_CLIENT_PORT=2181 \
		confluentinc/cp-zookeeper:latest

lint:
	golangci-lint -v run ./...

gofumpt:
	gofumpt -l -w .

migrate-up:
	goose -dir ./migrations clickhouse "tcp://${CLICKHOUSE_USER}:${CLICKHOUSE_PASSWORD}@localhost:9000" up
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:9000" up
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:30001" up

migrate-down:
	goose -dir ./migrations clickhouse "tcp://${CLICKHOUSE_USER}:${CLICKHOUSE_PASSWORD}@localhost:9000" down
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:9000" down
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:30001" down


# cool commands
# kubectl run -i --tty --rm debug --image=busybox --restart=Never -- sh
# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx (Добавляет ingress nginx в helm repo - "helm repo list")
# helm install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace
# kubectl get deployment -n ingress-nginx
# kubectl get svc -n ingress-nginx
# kubectl get ingressclass -n ingress-nginx

# helm uninstall ingress-nginx -n ingress-nginx

# K8s-hub

Цель проекта развернуть сервисы и их зависимости в кубере (Сугубо в научных целях). Реализованы три сервиса + зависимости clickhouse, kafka:
 - telegram-bot: Сохранение сообщений пользователей в бд (hdl - telegram api)
 - admin-panel: Анализ пользовательских сообщений (hdl - http)
 - statistics-sender: Отправка статистики пользователям в указанные таймера (hdl - parsed timers (not cron))   

## Для запуска проекта

1) Запуск конфиг + секрет манифестов:   
``` kubectl apply -f .\k8s\config.yaml -f .\k8s\secret.yaml  ```
2) Запуск ingress:  
``` helm install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace ```  
``` kubectl apply -f .\k8s\ingress_nginx.yaml ```
3) Запуск clickhouse (миграции там же по джобе накатываются)  
``` kubectl apply -f .\k8s\clickhouse.yaml ```
4) Запуск zookeeper + kafka:  
``` kubectl apply -f .\k8s\zookeeper.yaml -f .\k8s\kafka.yaml ```
5) Запуск telegram-bot:  
``` kubectl apply -f .\k8s\telegram-bot.yaml ```
6) Запуск admin-panel:  
``` kubectl apply -f .\k8s\admin-panel.yaml ```
7) Запуск statistics-sender:  
``` kubectl apply -f .\k8s\statistics-sender.yaml ```

PS. Также вписать .env по .env.example

## Также в проекте

Было добавлено:
 - grpc между сервисами, fiber http клиент
 - swagger для admin-panel
 - линтер, мейкфайлы, докерфайлы
 - парсинг конфигов через енвы, ямлы
 - вспомогательные модули - логгер, валидатор

## Нюансы проекта

Что было реализовано на скорую руку, чтобы скорее подойти к самому развертыванию в кубере:
 - Парсинг таймеров для отправки статистики пользователям по енвам в формате hh:mm:ss -> Нельзя ни установить день, ни периодичность и тд
 - Максимально сокращенный функционал (не более 2 ручки на каждый сервис)
 - Развертывание кафки и бд в самом кластере
 - Сервис закидывающий сообщения в кафку и читающий из нее в одном модуле
 - Отсутствие юнит тестов, комментов к коду, некоторые недоделки из-за развертывания сначала на локале, затем в докере и наконец в кубере
 - Нет initcontainer, readinessprobe ... 

Почта - 1pyankov.d.s@gmail.com  
Телега - @Xonesent  
---
layout: post
title: "Dependencias del proyecto"
---

La arquitectura del sistema definida anteriormente se puede encontrar en [este *post*](https://varrrro.github.io/pay-up/2019/10/28/system-architecture.html). Aunque ahí ya indicamos algunas de las herramientas que íbamos a utilizar, a continuación vamos a describirlas en mayor profundidad.

Para la implementación de los microservicios de las distintas entidades y del API Gateway, usaremos el [paquete `mux` del Gorilla web toolkit](https://github.com/gorilla/mux).

Para agilizar las consultas sobre el balance en un grupo, implementaremos una caché con [Memcached](https://memcached.org/), usando el paquete [gomemcache](https://github.com/bradfitz/gomemcache) para trabajar con ésta desde nuestro código.

En el anterior *post* indicamos que íbamos a usar [PostgreSQL](https://www.postgresql.org/) para el almacenamiento persistente en un nuestro sistema. Sin embargo, para abstraer nuestra implementación de la base de datos utilizada, usaremos como ORM el paquete [GORM](https://github.com/jinzhu/gorm).

Usaremos [Zookeeper](https://zookeeper.apache.org/) para la configuración remota, con el cliente nativo de Go [go-zookeeper](https://github.com/samuel/go-zookeeper).

Los logs producidos por los distintos servicios serán estructurados con el paquete [Logrus](https://github.com/sirupsen/logrus).

Para comunicarnos con la cola de mensajes [RabbitMQ](https://www.rabbitmq.com/), usaremos el protocolo AMQP con el paquete [amqp](https://github.com/streadway/amqp).

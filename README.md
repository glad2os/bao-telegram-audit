# Интеграция c OpenBao

Это приложение предназначено для приёма и обработки логов аудита от **OpenBao** через TCP-сокет. Оно позволяет собирать и фильтровать аудиторские записи, передаваемые от OpenBao, ретранслируя их в Telegram

## Включение аудита

Для того чтобы настроить OpenBao на отправку логов аудита в ваше приложение, необходимо включить аудиторское устройство с использованием **TCP-сокета**.

1. **Включение аудита Kubernetes**:

   ```shell
   bao audit enable socket address=bao-telegram-svc.default.svc.cluster.local:9090 socket_type=tcp
   ```

2. **Проверка активации**:

   ```shell
   bao audit list
   ```

## Отладка

1. Запуск OpenBao в **dev**-режиме:

   ```shell
   bao server -dev -dev-root-token-id="dev-only-token"
   ```

2. Включение аудита для тестирования через сокет:

   ```shell
   bao audit enable socket address=localhost:9090 socket_type=tcp
   ```

3. Активация **KV Secret Engine**:

   ```shell
   bao secrets enable -path=secret kv
   ```

4. Проверка аудита:

   ```shell
   bao kv put secret/foo bar=baz
   bao kv get secret/foo
   ```

5. Логи аудита будут отправлены на локальный сокет `localhost:9090` и обрабатываться Go приложением.

## Как запустить сервер

### Локально

1. Соберите и запустите Go TCP-сервер, который будет прослушивать входящие аудиторские записи:

   ```shell
   go run main.go
   ```

   Сервер будет слушать порт `9090` и принимать все соединения от OpenBao.

2. Проверьте, что приложение запущено и выводит логи аудита в консоль.

### Kubernetes
#### Текущая реализация имеет default Service Account. Не выпускать в продакшн без RBAC
1. Примените манифесты **Deployment** и **Service**.
2. После развертывания убедитесь, что приложение работает корректно, и OpenBao может отправлять аудиторские записи на указанный адрес.
3. Helm chart допишу позже 🤡

#!/bin/sh
set -e

# Проверка, заданы ли переменные окружения
if [ -z "$DB_HOST" ] || [ -z "$DB_PORT" ] || [ -z "$DB_USER" ]; then
  echo "❌ Ошибка: Не заданы переменные окружения DB_HOST, DB_PORT или DB_USER"
  exit 1
fi

echo "⏳ Ожидание PostgreSQL на ${DB_HOST}:${DB_PORT} от пользователя ${DB_USER}..."

# Ожидание готовности PostgreSQL
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "✅ PostgreSQL доступен!"

# Запуск основного процесса
exec "$@"

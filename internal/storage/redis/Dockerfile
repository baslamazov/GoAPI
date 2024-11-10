# Используем официальный образ Redis  
FROM redis:7.4.1



# Устанавливаем рабочую директорию  
WORKDIR /data  

# Копируем конфигурационный файл Redis (если у вас есть свой конфиг)  
# COPY redis.conf /usr/local/etc/redis/redis.conf  

# Открываем порт для Redis  
EXPOSE 6379  

# Запускаем Redis с конфигурационным файлом (если вы его скопировали)  
# CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]  

# Если конфигурационный файл не нужен, просто запускаем Redis  
CMD ["redis-server"]  
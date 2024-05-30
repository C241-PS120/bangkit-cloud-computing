FROM python:3.12.3
WORKDIR /app
COPY . .
COPY .env .
RUN "pip" "install" "-r" "requirement.txt"
CMD ["python", "app.py",]
EXPOSE 9000
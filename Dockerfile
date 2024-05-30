FROM python:3.12.3-alpine
WORKDIR /app
COPY . .
COPY .env .
COPY requirements.txt requirements.txt
RUN pip install -r requirement.txt
CMD ["flask", "--app", "app", "run"]
EXPOSE 9000
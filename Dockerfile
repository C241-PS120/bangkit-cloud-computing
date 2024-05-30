FROM python:3.12.3-alpine
WORKDIR /app
COPY . .
COPY .env .
COPY requirements.txt requirements.txt
RUN pip3 install -r requirement.txt
CMD ["python3", "-m", "flask", "--app", "app", "run"]
EXPOSE 9000
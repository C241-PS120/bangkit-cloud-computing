FROM python:3.12.3-bookworm
WORKDIR /app
COPY . .
COPY .env .
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt
CMD ["python3", "-m", "flask", "--app", "app", "run"]
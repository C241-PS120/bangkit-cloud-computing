FROM python:3.12.3-bookworm
WORKDIR /app
COPY . .
COPY .env .
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt
EXPOSE 5000
CMD ["python3", "app.py"]
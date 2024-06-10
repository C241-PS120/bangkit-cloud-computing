# Stage 1: Build the application
FROM python:3.12-slim as builder

ENV PYTHONUNBUFFERED=1 \
    PYTHONHASHSEED=random

# Install dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    gcc \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt requirements.txt

# Install the dependencies in a virtual environment
RUN python -m venv /opt/venv && \
    . /opt/venv/bin/activate && \
    pip install --no-cache-dir -r requirements.txt

# Stage 2: Create the final image
FROM python:3.12-slim as realease

# Copy the virtual environment from the builder stage
COPY --from=builder /opt/venv /opt/venv

# Ensure the virtual environment is used
ENV PATH="/opt/venv/bin:$PATH"

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["python", "app.py"]

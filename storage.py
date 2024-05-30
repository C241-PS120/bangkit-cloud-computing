from google.cloud import storage

import os

from pathlib import Path

import uuid

from PIL import Image

from io import BytesIO

class Storage:
    _instances = {}
    storageType: str
    client: storage.Client
    bucket: storage.Bucket
    
    #singleton class
    def __new__(cls, storageType: str):
        if storageType not in cls._instances:
            cls._instances[storageType] = super().__new__(cls)
            cls._instances[storageType].initialize(storageType)
        return cls._instances[storageType]

    def initialize(self, storageType: str):
        self.storageType = storageType
        
        if storageType == "model":
            self.client = storage.Client.from_service_account_json(Path("keys/" + os.environ["MODEL_KEY_FILENAME"]).resolve())
            self.bucket = self.client.bucket(os.environ["STORAGE_BUCKET_MODEL"])
        elif storageType == "photo":
            self.client = storage.Client.from_service_account_json(Path("keys/" + os.environ["PHOTO_KEY_FILENAME"]).resolve())
            self.bucket = self.client.bucket(os.environ["STORAGE_BUCKET_PHOTO"])
        else:
            raise ValueError(f"Unknown storageType: {storageType}")

    def upload(self, image: Image):
        imageName = f"{str(uuid.uuid4())}.jpg"

        blob = self.bucket.blob(f"upload/{imageName}")

        image_bytes = BytesIO()

        image.save(image_bytes, format="JPEG")

        image_bytes.seek(0)

        blob.upload_from_file(image_bytes,  content_type='image/jpeg')

        return blob.public_url

    def download(self, blobName: str, destinationFileName: str):
        try:
            blob = self.bucket.blob(blob_name=blobName)

            path = Path("temp/model/" + destinationFileName).resolve()

            os.remove(path)
            blob.download_to_filename(path)
        except FileNotFoundError:
            pass
        except Exception as e:
            print("Error downloading blob:", e)
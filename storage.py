from google.cloud import storage

import os

from pathlib import Path

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
            self.setStorageClient(Path("keys/" + os.environ["MODEL_KEY_FILENAME"]).resolve())
            self.bucket = self.client.bucket(os.environ["STORAGE_BUCKET_MODEL"])
        elif storageType == "photo":
            self.setStorageClient(Path("keys/" + os.environ["PHOTO_KEY_FILENAME"]).resolve())
            self.bucket = self.client.bucket(os.environ["STORAGE_BUCKET_PHOTO"])
        else:
            raise ValueError(f"Unknown storageType: {storageType}")

    def setStorageClient(self, path: Path = None):
        prod = os.environ.get("PRODUCTION")
        print("=====================================")
        print(prod, type(prod))
        production = bool(os.environ.get("PRODUCTION", "False") == "True")
        if production:
            self.client = storage.Client()
        else:
            self.client = storage.Client.from_service_account_json(path)

    def upload(self, id: str, image: Image.Image):
        imageName = f"{id}.{"jpg"}"

        blob = self.bucket.blob(f"upload/{imageName}")

        image_bytes = BytesIO()

        image.save(image_bytes, format="JPEG")

        image_bytes.seek(0)

        blob.upload_from_file(image_bytes, content_type=f'image/jpeg')

        return blob.public_url

    def download(self, blobName: str, destinationFileName: str):
        try:
            blob = self.bucket.blob(blob_name=blobName)

            if os.path.isfile(destinationFileName):
                os.remove(destinationFileName)
            blob.download_to_filename(destinationFileName)
        except FileNotFoundError:
            pass
        except Exception as e:
            print("Error downloading blob:", e)
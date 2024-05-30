import tensorflow as tf
import numpy as np
from PIL import Image

from keras import models, preprocessing

from storage import Storage

import os

from io import BytesIO

from pathlib import Path

class Model:
    label: list[str]

    #singleton class
    def __new__(cls):
        if not hasattr(cls, 'instance'):
            cls.instance = super(Model, cls).__new__(cls)
        return cls.instance

    #load the model from cloud storage to this instance
    def loadModel(self):
        storage_instance = Storage("model")
        storage_instance.download(os.environ['MODEL_NAME'], os.environ['MODEL_NAME'])

        self.model = models.load_model(Path(f"temp/model/{os.environ['MODEL_NAME']}").resolve())

    #load the label from local 
    def loadLabel(self):
        f = open(Path("label/label.txt").resolve())
        self.label = f.readlines()
        f.close()

    #preprocessing the image
    def preprocessImage(self, image: Image):
        image_bytes = BytesIO()

        image.save(image_bytes, format="JPEG")

        image_bytes.seek(0)

        processedImage = preprocessing.image.load_img(image_bytes, color_mode= "rgb", target_size=[224, 224])
        processedImage = preprocessing.image.img_to_array(processedImage) / 255.0
        return np.expand_dims(processedImage, axis=0)
    
    #predict given image with this model
    def predict(self, data):
        predict = self.model.predict(data)
        index_max = np.argmax(predict[0])
        return self.label[index_max], predict[0][index_max] * 100
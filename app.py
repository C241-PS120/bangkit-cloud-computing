from flask import Flask, request, jsonify

from dotenv import load_dotenv
import os
from storage import Storage
from model import Model
from pathlib import Path

from PIL import Image
from rembg import remove, new_session
import uuid

from waitress import serve

# initialization
load_dotenv()
os.environ["U2NET_HOME"] = str(Path("temp/model").resolve())
Storage("model")
Storage("photo")
m = Model()
m.loadModel()
m.loadLabel()
app = Flask(__name__)
rembgSession = new_session()

# constant
ALLOWED_MIME_TYPES = {"image/jpeg", "image/png"}

#
def allowed_file(filename):
    allowed_extensions = {"jpg", "jpeg", "png"}
    return "." in filename and filename.rsplit(".", 1)[1].lower() in allowed_extensions


def responseSuccess(data, message):
    return {"status": "success", "message": message, "data": data}


def responseFail(message):
    return {"status": "error", "message": message}


# route
@app.route("/")
def index():
    return "Hello!", 200


@app.route("/api/v1/predict", methods=["POST"])
def predict():
    if "photo" not in request.files or request.files["photo"].filename == "":
        return jsonify(responseFail("No Image given")), 400

    image = request.files["photo"]

    if image.mimetype not in ALLOWED_MIME_TYPES or not allowed_file(image.filename):
        return jsonify(responseFail("File is not an allowed image format")), 415

    openImage = Image.open(image)

    removedBgImage = remove(openImage, session=rembgSession, bgcolor=(0, 0, 0, 255))
    if removedBgImage.mode == "RGBA":
        removedBgImage = removedBgImage.convert("RGB")

    m = Model()

    data = m.preprocessImage(image=removedBgImage)
    label, confidentScore = m.predict(data)

    responseData = None

    if confidentScore >= 70.0:
        id = str(uuid.uuid4())

        fileUrl = Storage("photo").upload(id, image=removedBgImage)

        responseData = {
            "id": id,
            "label": label,
            "suggestion": f"Menurut hasil prediksi, tumbuhan kopimu dalam kondisi {'sehat!' if label == 'Healthy' else f'mengalami penyakit {label}, segeralah cari pestisida sebelum tanamanmu rusak!.'}",
            "search": label == "Healthy" if None else label.casefold(),
            "imageUrl": fileUrl,
        }

        return (
            jsonify(responseSuccess(responseData, "Model is predicted successfully")),
            201,
        )
    return jsonify(responseFail("Please try again with a clearer image.")), 400


# server run
if __name__ == "__main__":
    production = os.environ.get("PRODUCTION", "False").lower() == "true"

    if production:
        serve(app, host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))
    else:
        app.run(host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))

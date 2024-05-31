from flask import Flask, render_template, request, jsonify


from dotenv import load_dotenv
import os
from storage import Storage
from model import Model

from PIL import Image
import uuid

from waitress import serve

#initialization
load_dotenv()
Storage("model")
Storage("photo")
m = Model()
m.loadModel()
m.loadLabel()
app = Flask(__name__)

#constant
ALLOWED_MIME_TYPES = {'image/jpeg', 'image/png'}

#
def allowed_file(filename):
    allowed_extensions = {'jpg', 'jpeg', 'png'}
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in allowed_extensions

def responseSuccess(data, message):
    return {
        "status": "success",
        "message": message,
        "data": data
    }

def responseFail(message):
    return{
        "status": "error",
        "message": message
    }

#route
@app.route("/")
def index():
    return "Hello!", 200

@app.route("/api/v1/predict", methods=["POST"])
def predict():
    if 'photo' not in request.files or request.files['photo'].filename == '':
        return jsonify(responseFail("No Image given")), 400
    
    image = request.files['photo']

    if image.mimetype not in ALLOWED_MIME_TYPES or not allowed_file(image.filename):
        return jsonify(responseFail("File is not an allowed image format")), 415

    openImage = Image.open(image)

    if openImage.mode == 'RGBA':
        openImage = openImage.convert("RGB")

    m = Model()

    data = m.preprocessImage(image=openImage)
    label, confidentScore = m.predict(data)

    responseData = None 

    if(confidentScore >= 90.0):
        id = str(uuid.uuid4())

        fileUrl = Storage("photo").upload(id, image=openImage)

        responseData = {
            "id": id,
            "label": label, 
            "suggestion": f"Menurut hasil prediksi, tumbuhan kopimu dalam kondisi {'sehat!' if label == 'Healthy' else f'mengalami penyakit {label}'}, segeralah cari pestisida sebelum tanamanmu rusak!.",
            "search": label == "Healthy" if None else label.casefold(),
            "imageUrl": fileUrl
        }
        
        return jsonify(responseSuccess(responseData, "Model is predicted successfully")), 201
    return jsonify(responseFail("An error occurred in detecting the image, please try again with a clearer image.")), 400



#server run
if __name__ ==  '__main__':
    production = bool(os.environ.get("PRODUCTION", "False") == "True")

    if production:    
        serve(app, host='0.0.0.0', port=int(os.environ.get("PORT", 5000)))
    else:
        app.run(host='0.0.0.0', port=int(os.environ.get("PORT", 5000)))
rust = [
    "Munculnya bintik-bintik kuning kecil pada permukaan daun bagian atas.", 
    "Bintik-bintik tersebut berkembang menjadi bercak oranye kekuningan di bagian bawah daun.",
    "Daun yang terinfeksi parah menguning dan gugur."
    ]

miner = [
    "Munculnya jejak berliku-liku berwarna putih atau cokelat pada permukaan daun.", 
    "Adanya garis-garis transparan atau pucat di bawah epidermis daun kopi.",
    "Daun yang terinfeksi parah mengering dan gugur lebih cepat dari biasanya."
    ]

phoma = [
    "Munculnya bercak-bercak kecil berwarna cokelat atau hitam pada daun.", 
    "Bercak-bercak tersebut berbentuk bulat atau oval dengan tepi yang jelas.",
    "Daun yang terinfeksi parah menguning, mengalami nekrosis (kematian sel), dan akhirnya gugur."
    ]


def getSymptoms(label: str):
    if label == "Rust":
        return rust
    if label == "Miner":
        return miner
    if label == "Phoma": 
        return phoma
    if label == "Healthy":
        return []
CREATE TABLE IF NOT EXISTS label (
    label_id INT AUTO_INCREMENT PRIMARY KEY,
    label_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS disease (
    disease_id INT AUTO_INCREMENT PRIMARY KEY,
    disease_name VARCHAR(100) NOT NULL UNIQUE,
    cause TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS article (
    article_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url VARCHAR(255),
    symptom_summary TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    disease_id INT,
    label_id INT,
    FOREIGN KEY (disease_id) REFERENCES disease(disease_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS plant (
    plant_id INT AUTO_INCREMENT PRIMARY KEY,
    plant_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS plant_disease (
    plant_id INT,
    disease_id INT,
    PRIMARY KEY (plant_id, disease_id),
    FOREIGN KEY (plant_id) REFERENCES plant(plant_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (disease_id) REFERENCES disease(disease_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS symptom (
    symptom_id INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT,
    symptom_description TEXT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS prevention (
    prevention_id INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT,
    prevention_description TEXT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS treatment (
    treatment_id INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT,
    treatment_description TEXT NOT NULL,
    treatment_type ENUM('organic', 'chemical') NOT NULL CHECK (treatment_type IN ('organic', 'chemical')),
    FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_article_title ON article (title);
CREATE INDEX idx_disease_name ON disease (disease_name);
CREATE INDEX idx_label_name ON label (label_name);

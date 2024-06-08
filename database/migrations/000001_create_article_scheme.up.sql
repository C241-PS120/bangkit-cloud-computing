CREATE TABLE category (
                          category_id INT AUTO_INCREMENT PRIMARY KEY,
                          category_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE article (
                         article_id INT AUTO_INCREMENT PRIMARY KEY,
                         title VARCHAR(255) NOT NULL,
                         content TEXT NOT NULL,
                         cause TEXT NOT NULL,
                         image_url VARCHAR(255),
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         category_id INT,
                         FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE SET NULL
);

CREATE TABLE symptom (
                         symptom_id INT AUTO_INCREMENT PRIMARY KEY,
                         article_id INT,
                         symptom_description TEXT NOT NULL,
                         FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
);

CREATE TABLE prevention (
                            prevention_id INT AUTO_INCREMENT PRIMARY KEY,
                            article_id INT,
                            prevention_description TEXT NOT NULL,
                            FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
);

CREATE TABLE treatment (
                       treatment_id INT AUTO_INCREMENT PRIMARY KEY,
                       article_id INT,
                       treatment_description TEXT NOT NULL,
                       treatment_type ENUM('organic', 'chemical') NOT NULL CHECK (treatment_type IN ('organic', 'chemical')),
                       FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
);

CREATE INDEX idx_article_title ON article (title);
CREATE INDEX idx_category_name ON category (category_name);

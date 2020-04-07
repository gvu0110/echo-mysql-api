

```
docker pull mysql/mysql-server:lastest
docker run -d -it --name mysql -p 3306:3306 mysql/mysql-server:latest

# Get the temporary password
docker logs mysql

docker exec -it mysql mysql -u root -p

# Change password
ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';

# Allow remote connection
USE mysql;
UPDATE mysql.user SET host='%' WHERE user='root'
FLUSH PRIVILEGES;

CREATE DATABASE foo;
USE foo;
CREATE TABLE Employees (
id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
employee_name NVARCHAR(100) NOT NULL,
employee_age INT NOT NULL,
employee_salary FLOAT NOT NULL
);
INSERT INTO Employees (id, employee_name, employee_salary, employee_age)
VALUES (1, "Adam Vu", 1234.56, 25),
(2, "Mikhaila Santos", 2345.67, 24),
(3, "Robert Robertson", 3456.78, 30);
```
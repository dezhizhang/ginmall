# ginApi
ginApi

### user表
```
CREATE TABLE users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `age` INT NOT NULL,
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
### role角色表
```
CREATE TABLE roles(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    status INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    add_time INT NOT NULL,
    PRIMARY KEY (id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
### 管吾员表
```
CREATE TABLE managers(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    mobile VARCHAR(35) NOT NULL,
    email VARCHAR(255) NOT NULL,
    status INT NOT NULL,
    role_id INT NOT NULL,
    add_time INT NOT NULL,
    is_super INT NOT NULL,
    PRIMARY KEY (id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


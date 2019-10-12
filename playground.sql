CREATE TABLE IF NOT EXISTS students(
   id serial PRIMARY KEY,
   firstName VARCHAR (50) NOT NULL,
   secondName VARCHAR (50) NOT NULL,
   email VARCHAR (355) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS courses(
   id serial PRIMARY KEY,
   title VARCHAR (50) NOT NULL
);

INSERT INTO students (firstName, secondName, email) VALUES ('student', 'one', 'student_one@gmail.com');
INSERT INTO students (firstName, secondName, email) VALUES ('student', 'two', 'student_two@gmail.com');

INSERT INTO courses (title) VALUES ('go');
INSERT INTO courses (title) VALUES ('python');
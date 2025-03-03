## ERD Title: Company & Employees Management System

1. Entities and Their Attributes:
Entity: Companies, Departments, Employees, Projects

## Attributes:

A. Entity: companies

- Attributes:
- id INT (PK, AI)
- name Varchar(15) NOT NULL
- address Varchar(100) NOT NULL

B. Entity: departments

- Attributes:
- id INT (PK, AI)
- name Varchar(15) NOT NULL
- company_id INT (FK) NOT NULL

C. Entity: employees

- Attributes:
- id INT (PK, AI)
- name Varchar(15) NOT NULL
- company_id INT (FK) NOT NULL
- department_id INT (FK) NOT NULL

D. Entity: projects

- Attributes:
- id INT (PK, AI)
- name Varchar(128) NOT NULL

E. Entity: employee_projects

- Attributes:
- id INT (PK, AI)
- employee_id INT (FK) NOT NULL
- projects_id INT (FK) NOT NULL


## Relationships:
- companies to employees:

A. Type: One to Many
- Description: One company can have many employees, but each employee works for one company.


- employees to projects:

B. Type: Many to Many
- Description: One employee can work on multiple projects, and each project can have multiple employees working on it.


- companies to departments:

C. Type: One to Many
- Description: One company can have many departments.

- departments to employees:

D. Type: One to Many
- Description: One department can have many employees working on it, but each employee works in one department.

## Integrity Constraints:
 • Non-Null Constraints: all attributes cannot be null.
 • Foreign Key Constraints: Foreign keys (e.g., company_id in departments, employees; department_id in employees) must match values in their referenced tables to maintain referential integrity.
 • Auto Increment: id attributes in all entity should always be positive and unique.
 • Unique Constraints: Each employee_id-project_id combination in the employee_projects table should be unique to prevent duplicate assignments.

## Additional Notes:

• The employee_projects table resolves the many-to-many relationship between Employee and Project.
 • The system can handle multiple departments, employees, and projects for each company, providing flexibility in management.
 • Employees can be assigned specific roles on each project, helping in project resource planning.
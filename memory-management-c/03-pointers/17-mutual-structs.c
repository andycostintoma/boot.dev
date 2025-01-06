#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

typedef struct Employee employee_t;
typedef struct Department department_t;

typedef struct Employee {
    int id;
    char *name;
    department_t *department;
} employee_t;

typedef struct Department {
    char *name;
    employee_t *manager;
} department_t;

employee_t create_employee(int id, char *name) {
    employee_t emp = {
        .id = id,
        .name = name,
        .department = NULL
    };
    return emp;
}

department_t create_department(char *name) {
    department_t dept = {
        .name = name,
        .manager = NULL
    };
    return dept;
}

void assign_employee(employee_t *emp, department_t *department) {
    emp->department = department;
}

void assign_manager(department_t *dept, employee_t *manager) {
    dept->manager = manager;
}

void test_assign_employee1() {
    employee_t emp = create_employee(2, "CEO Dax");
    department_t dept = create_department("C Suite");
    assign_employee(&emp, &dept);
    assert(strcmp(emp.department->name, "C Suite") == 0);
}

void test_assign_employee2() {
    employee_t emp = create_employee(4, "Vegan Intern Adam");
    department_t dept = create_department("Marketing");
    assign_employee(&emp, &dept);
    assert(strcmp(emp.department->name, "Marketing") == 0);
}

void test_assign_manager1() {
    employee_t manager = create_employee(3, "Influencer Prime");
    department_t dept = create_department("Marketing");
    assign_manager(&dept, &manager);
    assert(strcmp(dept.manager->name, "Influencer Prime") == 0);
}

void test_assign_manager2() {
    employee_t manager = create_employee(5, "CDO David");
    department_t dept = create_department("C Suite");
    assign_manager(&dept, &manager);
    assert(strcmp(dept.manager->name, "CDO David") == 0);
    assert(manager.id == 5);
}

int main() {
    test_assign_employee1();
    test_assign_employee2();
    test_assign_manager1();
    test_assign_manager2();

    printf("All tests passed!\n");

    return 0;
}

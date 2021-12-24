using System;

namespace Test;

public static void Main() {
    var emp = new Employee();
    emp.Grade = 1;

    Calculate(emp);
    Console.WriteLine($"Grade : {emp.Grade}");
}

private static void Calculate(Employee employee)
{
    employee.Grade = 2;
}

public class Employee 
{
    public int Grade {get;set;}
}
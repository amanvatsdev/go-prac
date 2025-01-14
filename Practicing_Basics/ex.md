### Exercise: **Student Grade Calculator**

#### Description:  
Write a Go program to manage a list of students and calculate their average grades. Each student has a name, roll number, and a list of grades for three subjects. The program should:  

1. Use **variables** to store and manipulate data.  
2. Use a **struct** to define the student details (name, roll number, grades).  
3. Use a **slice** to store multiple students.  
4. Use a **loop** to iterate through the students and calculate their average grades.  
5. Use a **function** to perform the grade calculation.

---

### Instructions:

1. **Define a Struct**:
   Create a `Student` struct with the following fields:
   - `RollNumber` (int)
   - `Name` (string)
   - `Grades` (slice of `float64` for storing grades of three subjects)

2. **Add Data**:
   Create at least 3 students with their respective roll numbers, names, and grades.

3. **Write a Function**:
   Create a function `calculateAverage` that takes a slice of `float64` (grades) as input and returns the average grade as a `float64`.

4. **Use a Loop**:
   Use a `for` loop to iterate over the slice of students, calculate their average grades using the function, and print their names along with their average grades.

5. **Use Constants**:
   Define a constant `PassMark` (e.g., 40). Check in the loop if a student’s average grade is greater than or equal to `PassMark`. Print whether the student has **passed** or **failed**.

---

### Expected Output:
For students like:
```go
[
  {RollNumber: 1, Name: "Alice", Grades: [85, 78, 92]},
  {RollNumber: 2, Name: "Bob", Grades: [40, 35, 50]},
  {RollNumber: 3, Name: "Charlie", Grades: [30, 25, 20]}
]
```

The program should output something like:
```plaintext
Student: Alice, Average Grade: 85.0 - Passed
Student: Bob, Average Grade: 41.67 - Passed
Student: Charlie, Average Grade: 25.0 - Failed

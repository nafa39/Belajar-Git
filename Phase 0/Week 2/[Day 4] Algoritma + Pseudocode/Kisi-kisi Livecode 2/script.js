//Soal 1

//declare variable array 2 dimensi

let facultyGrades1 = [
    [80, 86, 90, 78, 88],
    [70, 76, 80, 72, 79],
    [65, 68, 71, 66, 67]
];

let facultyGrades2 = [
    [82, 88, 89, 80, 85],
    [71, 73, 75, 72, 70],
    [65, 70, 68, 67, 69]
];

//declare function calculate average

function calculateAverage(grade){
    let sum = [];
    let a = [];
    let avg = [];
    for (let i=0; i < grade.length; i++){
        sum[i] = 0;
        a[i] = grade[i].length;
        for (let j=0; j<grade[i].length; j++){
            sum[i] += grade[i][j];
        }
        sum.push(sum[i]);
        sum.pop();
        a.push(a[i])
        a.pop();
    }
    for (let x = 0; x< sum.length; x++){
        avg[x] = sum[x]/a[x]
        avg.push(avg[x]);
        avg.pop();
    }
    console.log(avg);
}

console.log(calculateAverage(facultyGrades2));

//Soal 2


/*let n = prompt('Masukkan bilangan:');

function prima(n){
    if(n<=1){
        return false;
    }
    for (let i=2; i<=n-1;i++){
        if (n % i === 0){
            return false;
        }
    }
    return true;
}

if (prima(n)){
    console.log('Ini bilangan prima');
}
else{
    console.log('Ini bukan bilangan prima');
}*/

//let list = [];

// let lowLimit = parseInt(prompt('Masukkan limit bawah:'));
// let highLimit = parseInt(prompt('Masukkan limit atas:'));

// function bilPrima(n){
//     if(n<=1){
//         return false;
//     }
//     for (let i=2; i<=n-1;i++){
//         if (n % i === 0){
//             return false;
//         }
//     }
//     return true;
// }

// while (lowLimit<=highLimit){
//     if (bilPrima(lowLimit)){
//         console.log(lowLimit);
//     }    
//     lowLimit++;
// }
    

//console.log(list);

//MATERI LATIHAN P0W2D3

var arr = [1, 2, 15];
var arrChar = ['Tono', 'Budi', 'Charlie', 'Ahmad']

//Push()
// arr.push(4);

//Pop()
// arr.pop();

//UnShift()
// arr.unshift(3);

//shift()
// arr.shift();

//sort()
// arr.sort();
// arrChar.sort();
// arr.sort(function(value1, value2){return value1 > value2});

//splice()
// var arr1 = ['buku', 'laptop', 'komputer'];

// arr1.splice(2, 0, 'televisi'); //mulai dari index 2, menghapus 0 data, menambahkan nilai televisi
// console.log(arr1);

// arr1.splice(0, 2); //mulai dari index 0, menghapus 2 data
// console.log(arr1);

// arr1.splice(0, 1, 'majalah', 'koran'); //mulai dari index 0, menghapus 1 data, menambahakan majalah dan koran
// console.log(arr1);

//split()
// var kalimat = 'saya adalah full-stack javascript programmer!';
// var kata = kalimat.split(' ');
// console.log(kata);

// console.log(arr);
// console.log(arrChar);

//object()
const supermanObj = {
    id: '001',
    name: 'superman',
    address: 'Cryton street',
    age: 32,
    stillAlive: true,
    hobbies: [
        'saving planet earth',
        'reading newspaper',
    ],
    favoriteFoods: [
        'bakso granat',
        'pecel lele'
    ]
}

const person = {
    name: 'John',
    age: 23
}

// property re assignment

person.name = 'Mike';
person.age = 40;

//modular function


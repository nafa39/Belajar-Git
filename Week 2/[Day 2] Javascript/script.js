var index = 0;

/*while (index < 10){
    console.log("Perulangan ke:", index);
    index++
}*/

do{
    console.log('Perulangan ke:', index)
    index++
}
while(index<10);

//forEach loop v1
var students = ["John", "Mark", "Elisa"];
students.forEach(function students(value, index){
    console.log(index,': ', value)
});

//forEach loop v2
students.forEach((value, index) => {console.log(index,': ', value)});

//repeat loop
var name= "john doe";
console.log(name.repeat(5));

//for loop
for (var i= 3; i <= 9; i++){
    console.log('Perulangan ke: ', i);
}

//nested loop

for (var i=0; i < 10; i++){
    console.log('perulangan pertama: ',i);
    for(var j = 3; j<=i; j++){
        console.log('perulangan kedua: ',j);
    }
}

//last challenge 1
let serialNumber = 1;
while (serialNumber<=20){
    console.log('Halo, saya orang ke', serialNumber);
    serialNumber++;
}

//last challenge 1 with for
for (serialNumber = 1; serialNumber<=20;serialNumber++){
    console.log('Halo, saya orang ke', serialNumber);
}

//last challenge 2
let b = prompt('Masukkan batas jumlah huruf O:');
for (let a = 1; a<=b; a++){
    if(a % 2 === 1){
        console.log('O'.repeat(a));
    }
}

//NGC 5

//conditional bandingkan angka
let angka1 = prompt('Masukkan angka pertama:');
let angka2 = prompt('Masukkan angka kedua:');

if (angka1 < angka2){
    console.log('true');
}
else if (angka1 > angka2){
    console.log('false');
}
else if (angka1 === angka2){
    console.log('-1');
}

//conditional graduate
let nama = prompt('masukkan nama murid anda:');
let nilai = prompt('berapa nilai murid anda:')
let absen = prompt('berapa kali murid ini absen:');

if (nilai >= 70){
    if (absen<5){
        console.log('Murid dengan nama: ', nama, 'dinyatakan LULUS');
    }
    else{
        console.log('Murid dengan nama: ', nama, 'dinyatakan TIDAK LULUS')
    }
}
else{
    console.log('Murid dengan nama: ', nama, 'dinyatakan TIDAK LULUS')
}
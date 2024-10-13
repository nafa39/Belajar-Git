//CASE 1 - DISNEY ISLAND

//declare variable and let customer input
/*let wahana = prompt('Wahana mana yang ingin anda kunjungi?\n');
let usia = parseInt(prompt('Berapa usia anda?'));
let saldo = parseInt(prompt('Mohon masukkan saldo anda'));

//declare function for dcide the tarif
function disneyIsland(){
    let tarif = 0;
    //conditional foe decide the wahana
    if (wahana === "Wahana Utara"){
        //nested if/else for decide the age
        if(usia<=1){
            //to describe that value is false
            console.log('Dilarang Masuk');
            return false;
        }
        else if (usia >=2 && usia <=12){
            //input the price
            tarif = 85000;
        }
        else if (usia >=13){
            tarif = 125000;
        }
    }
    else if (wahana === "Wahana Selatan"){
        if(usia<=1){
            console.log('Dilarang Masuk');
            return false;
        }
        else if (usia >=2 && usia <=12){
            tarif = 143000;
        }
        else if (usia >=13){
            tarif = 165000;
        }
    }
    else{
        console.log('Tket tidak ditemukan');
        return false;
    }
    //return to the price when the value is true
    return tarif;
}

//declare total and tarif
let tarif = disneyIsland();
let total = saldo - tarif;

//conditional if the the total remain or not
if (tarif !== false){
    if (total<0){
        //declare the minus to absolute
        let kurang = Math.abs(total);
        console.log(`Saldo anda kurang RP ${kurang},00. Tidak cukup untuk membeli tiket.`)
    }
    else{
        console.log(`Sisa saldo anda adalah RP ${total},00. Selamat bermain.`)
    }
        
}

//CASE 2 - AYO JOGET

//declare input variable
var excercise = prompt('Masukkan gerakan latihan');
var userInput = prompt('Masukkan gerakan yang dilakukan user');

//declare the length of string input variable
let amountExercise = excercise.length;
let amountUser = userInput.length;

//declare array of input variable
let arrayExercise = excercise.split("");
let arrayUser = userInput.split("");

//declare total score must be reach
let totalScore = 10 * amountExercise;

//conditional decision which the length of execise and user same or not
if (amountExercise === amountUser){

    //loop for count the true exercise
    let trueCount = 0;
    for (let i = 0; i < amountUser; i++){
        if (arrayExercise[i] === arrayUser[i]){
            trueCount++;
        }
    }

    //declare the score that get
    let score = 10 * trueCount;

    //declare the percentage
    let percentage = Math.floor((trueCount/amountExercise)*100);

    //declare kategori and conditional if/else for the kategori
    let Kategori = '';
    if (percentage === 100){
        Kategori = "Perfect";
    }
    else if (percentage<= 99 && percentage >= 80){
        Kategori = "Great";
    }
    else if (percentage>= 60 && percentage <= 79){
        Kategori = "Good";
    }
    else if(percentage>= 0 && percentage<= 59){
        Kategori = "Bad";
    }

    console.log(`Anda mendapatkan score ${score} / ${totalScore}. Persentase: ${percentage}%, Kategori: ${Kategori}`)
}
else{
    console.log('Input yang anda masukkan tidak lengkap!');
}*/


//CASE 3 - ASTERIX IN THE BOX

//declare the variable of array for row
let rowArray = [];

//declare variable for column number
let col = 5;

//declare variable content hashtag
let hashtag = "#";

//declare variable for row number as input
let row = parseInt(prompt('Masukkan jumlah row'));

//create function to resulting the array for 1 row with 5 column
function columns(){
    let colArray = [];
    for (let i = 0; i <= col - 1; i++){
        colArray[i] = hashtag;
    }
    return colArray;
}

//create function to resulting the 2 dimension array based on row number
function rows(){
    for (let j = 0; j <= row-1; j++){
        rowArray[j] = columns();
    }
    //hide the result of matriks
    //rowArray.forEach(row => console.log(row));
}

//call the function of matrix
rows();

//declare variable star, input coordinate and matrix coordinate
let star = "*";
let coordinate = prompt('Masukkan koordinat bintang');
let rowCoordinate = parseInt(coordinate[0])-1;
let colCoordinate = parseInt(coordinate[1])-1;

//declare replacing hashtag with star in the coordinate
rowArray[rowCoordinate][colCoordinate] = star;

//display the result
rowArray.forEach(row => console.log(row));

let n = prompt('Masukkan bilangan:');

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
}

let list = [];

let lowLimit = prompt('Masukkan limit bawah:');
let highLimit = prompt('Masukkan limit atas:');

function bilPrima(lowLimit){
    if(lowLimit<=1){
        return false;
    }
    for (let i=2; i<=lowLimit-1;i++){
        if (lowLimit % i === 0){
            return false;
        }
    }
    return true;
}

while (lowLimit<=highLimit){
    if (bilPrima(lowLimit)){
        list.push(lowLimit);
    }    
    lowLimit++;
}
    

console.log(list);
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

function prima(n){
    if (n <= 1){
        return false;
    }
    for (let i = 2; i <= Math.sqrt(n); i++){
        if (n % i === 0){
            return false;
        }
    }
    return true;
}

if (prima(n)){
    console.log('ini bilangan Prima')
}
else{
    console.log('ini bukan prima')
}

// a adalah batas bawah
//b adalah batas atas

function bilanganPrima(a, b){
    if (b <= a){
        return true;
    }
    for ( let i = 2; i<= b; 1++){
        if (a % i === 0 ){
            return false;
        }
    }
    return true;
}

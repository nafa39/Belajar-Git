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

calculateAverage(facultyGrades2);

//Soal 2

function checkKeyword(keyword){
    let word = keyword;
    let arrayKeyword = word.split(' ');
    if (arrayKeyword.length<2){
        return false;
    }
    else{
        let chara = [];
        for (let i=0; i<arrayKeyword.length; i++){
            let character = arrayKeyword[i].split("");
            chara.push(character);
            }
        
        for (let j=0; j<chara.length;j++){
            if (chara[j].length<4){
                return false;
            }
        }
    }
    for (let x=0; x<arrayKeyword.length;x++){
        const pattern = /[0-9~`!@#$%\^&*()+=\-\[\]\\';,/{}|\\":<>\?¬£.]+/;
        if (pattern.test(arrayKeyword[x])){
            //console.log('contains special chara');
            return false;
        }else{
            //console.log('not contains special chara');
            return true;
        }
    }
    
}

//console.log(checkKeyword("Harry Pottter"));

const libraryData = {
    publisher: "Bloomsburry",
    books: [
        "Harry Potter and the Philosoper's Stone",
        "Harry Potter and the Chamber of Secrets",
        "Harry Potter and the Prisoner of Azkaban",
        "The Casual Vacancy",
        "A Cuckoo's Calling",
        "The Silkworm"
    ]
}

function bookFinder(libraryData, keyword){
    checkKeyword(keyword);
    if (checkKeyword(keyword)){
        let key = keyword.toLowerCase().split(' ');
        let foundBooks = [];
        for (let i=0; i<libraryData.books.length; i++){
            let bookWords = libraryData.books[i].toLowerCase().split(' ');
            let match = key.every(word => bookWords.includes(word));
            if (match){
                foundBooks.push(libraryData.books[i]);
            }
        }

        console.log ([libraryData.publisher, [foundBooks]]);
    }else{
        console.log('Keyword not suitable');
    }
}

bookFinder(libraryData, "Harry Potter");

//Soal 3

const armorsAvailable = [
    "Dragon Scale",
    "Elven Leather",
    "Chainmail"
]

function getArmorAvailability (armors){
    let armorsAvail = [];
    for (let i=0; i<armorsAvailable.length; i++){
        if (armors === armorsAvailable[i]){
            armorsAvail.push(armorsAvailable[i]);
            console.log(`"${armorsAvail}": true`);
            return armorsAvail;
        }
    }
    console.log(`"${armors}": false`);
}
getArmorAvailability("Chainmail");

const armors = [
    ['Dragon Scale', 'Heavy', 50],
    ['Elven Leather', 'Light', 30],
    ['Wizard Robe', 'Cloth', 15],
    ['Chainmail', 'Medium', 40]
]

const character1 = {
    name: 'Lorian',
    classType: 'Knight',
    baseDefense: 100,
    armorUsed: ['Elven Leather', 'Chainmail']
}

function getArmorAv1ailability1 (armors1){
    let armorsAvail = [];
    for (let i=0; i<armors.length; i++){
        if (armors1 === armors[i][0]){
            armorsAvail.push(armors[i][0]);
            console.log(`"${armorsAvail}": ${true}`);
            return armors[i][2];
        }
    }
    console.log(`"${armors1}": ${false}`);
    //return index;
}

function getTotalDefense(character){
    let total = character.baseDefense;
    for (let i=0; i<character.armorUsed.length; i++){
        let addedDefense = getArmorAv1ailability1(character.armorUsed[i])
        total += addedDefense;
    }
    console.log(total);
}

getTotalDefense(character1);

const inputCharacters = [
    {
        character: {
            name: 'Lorian',
            classType: 'Knight',
            baseDefense: 100,
            armorUsed: []
        },
        armorToEquip: ['Elven Leather', 'Chainmail']
    },
];

const armorsInArmory = [
    "Dragon Scale",
    "Elven Leather",
    "Chainmail",
    "Wizard Robe"
]

function getArmorAvailability2 (armors){
    let armorsAvail = [];
    for (let i=0; i<armors.length; i++){
        if (armors1 === armors[i][0]){
            armorsAvail.push(armors[i][0]);
            console.log(`"${armorsAvail}": ${true}`);
            return armors[i][2];
        }
    }
    console.log(`"${armors1}": ${false}`);
}

function getTotalDefense(character){
    let total = character.baseDefense;
    for (let i=0; i<character.armorUsed.length; i++){
        let addedDefense = getArmorAv1ailability2(character.armorUsed[i])
        total += addedDefense;
    }
    console.log(total);
}

function getDefenseSummary (inputCharacters, armorsInArmory){
    for (let i=0; i<inputCharacters.length; i++){
        for (let j=0; j< inputCharacters[i].armorToEquip.length; j++){
            let arUsed = getArmorAvailability2 (inputCharacters[i].armorToEquip[j])
            inputCharacters[i].character[j].armorUsed.push(arUsed);
            return arUsed;
        }
        getTotalDefense(inputCharacters[i].character);
    }
}

getDefenseSummary(inputCharacters, armorsAvailable);

function addTwoNumbers (l1, l2) {
    let rl1 = ""
    let rl2 = ""
    for (let i = l1.length - 1; i >= 0; i--) {
        rl1 += l1[i]
    }
    for (let i = l2.length - 1; i >= 0; i--) {
        rl2 += l2[i]
    }

    let sum = Number(rl1) + Number(rl2)
    let result = Array.from(String(sum), Number)
    console.log(result)

};

addTwoNumbers([2,4,3], [5,6,4])

addTwoNumbers([0], [0])

addTwoNumbers([9,9,9,9,9,9,9], [9,9,9,9])




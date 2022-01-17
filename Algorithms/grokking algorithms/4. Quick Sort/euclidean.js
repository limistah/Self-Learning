const euclidean = (a, b) => {
    if (a === 0 || b === 0) {
        return 0
    }
    const remainder = a%b
    if (remainder === 0){
        return b
    }
    return euclidean(b, remainder)
}
console.log(euclidean(270, 192))
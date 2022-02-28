
function merge(arr, start, mid, end) {
    let l1 = mid - start + 1
    let l2 = end - mid
    const L = Array.from({length: l1 + 1}).fill(0).map((val, index) => {
        return arr[start+index - 1]
    })
    const R = Array.from({length: l1 + 1}).fill(0).map((val, index) => {
        return arr[mid+index]
    })

    console.log({L, R})
    return arr
}

function mergeSort (arr, start, end) {
    if ( start < end) {
        const mid = Math.ceil((start + end) / 2)
        mergeSort(arr, start, mid)
        mergeSort(arr, mid+1, end)
        merge(arr, start, mid, end)
    }
    return arr
}

let arr = [31, 41, 59, 26, 42, 58, 4, 1, 3, 0, 9]
const sorted = mergeSort(arr, 0, arr.length-1);

console.log(arr)
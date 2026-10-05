fn insertion_sort(arr: &mut [i32]) {
    for i in 1..arr.len() {
        let key = arr[i];
        let mut j = i;

        while j > 0 && arr[j - 1] > key {
            arr[j] = arr[j - 1];
            j -= 1;
        }

        arr[j] = key;
    }
}

fn main() {
    let mut arr = [3, 43, 21, 54, 75, 22, 21];
    println!("Before : {:?}", arr);

    insertion_sort(&mut arr);

    println!("After : {:?}", arr);
}

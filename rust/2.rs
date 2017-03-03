fn main() {
    let (mut a, mut b, mut acc) = (1, 1, 0);
    while a < 4_000_000 {
        if a % 2 == 0 {
            acc += a
        }
        let tmp = a;
        a = b;
        b += tmp;
    }
    println!("{}", acc);
}

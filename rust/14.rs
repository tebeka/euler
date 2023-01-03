use std::time::SystemTime;

fn collatz(n: u64) -> u64 {
    if n % 2 == 0 {
        return n / 2;
    }
    return n * 3 + 1;
}

fn collatz_len(n: u64) -> u64 {
    let mut len = 1;
    let mut n = n;

    while n != 1 {
        len += 1;
        n = collatz(n);
    }

    return len;
}

fn main() {
    let now = SystemTime::now();
    let (mut n, mut nc) = (0, 0);
    for i in 1..1_000_001 {
        let c = collatz_len(i);
        if c > nc {
            n = i;
            nc = c;
        }
    }
    let duration = now.elapsed().unwrap();
    println!("{n} {duration:?}")
}

fn is_ok(n: u64) -> bool {
    for x in 2..20 {
        if n % x != 0 {
            return false;
        }
    }
    return true;
}

fn main() {
    for n in 20..1_000_000_000 {
        if is_ok(n) {
            println!("{}", n);
            break;
        }
    }
}

fn is_palinrome(n: u32) -> bool {
    let s: String = n.to_string();

    return s ==  s.chars().rev().collect::<String>();
}

fn main() {
    let mut max = 0;
    for x in 100..1000 {
        for y in 100..1000 {
            let mul = x * y;
            if mul > max && is_palinrome(mul) {
                max = mul;
            }
        }
    }
    println!("{}", max);
}

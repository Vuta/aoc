use std::fs;

fn main() {
    let contents = fs::read_to_string("input").expect("input not found");
    let mut result = 0;

    for line in contents.lines() {
        let mut i = 0;
        let mut j = line.len() - 1;
        let chars: Vec<char> = line.chars().collect();
        let mut first = 'z';
        let mut last = 'z';

        while i < line.len() {
            if first != 'z' && last != 'z' {
                break;
            }

            let f = chars[i];
            let l = chars[j];

            if first == 'z' && f >= '0' && f <= '9' {
                first = f;
            } else {
                i = i + 1;
            }

            if last == 'z' && l >= '0' && l <= '9' {
                last = l;
            } else {
                j = j - 1;
            }
        }

        result += first.to_digit(10).expect("") * 10 + last.to_digit(10).expect("");

    }

    println!("{}", result);
}

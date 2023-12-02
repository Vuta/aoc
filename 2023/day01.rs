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

            if first == 'z' {
                if single_digit(f) {
                    first = f;
                } else if one(i, &chars) {
                    first = '1';
                } else if two(i, &chars) {
                    first = '2';
                } else if three(i, &chars) {
                    first = '3';
                } else if four(i, &chars) {
                    first = '4';
                } else if five(i, &chars) {
                    first = '5';
                } else if six(i, &chars) {
                    first = '6';
                } else if seven(i, &chars) {
                    first = '7';
                } else if eight(i, &chars) {
                    first = '8';
                } else if nine(i, &chars) {
                    first = '9';
                } else {
                    i += 1;
                }
            }

            if last == 'z' {
                if single_digit(l) {
                    last = l;
                } else if one_r(j, &chars) {
                    last = '1';
                } else if two_r(j, &chars) {
                    last = '2';
                } else if three_r(j, &chars) {
                    last = '3';
                } else if four_r(j, &chars) {
                    last = '4';
                } else if five_r(j, &chars) {
                    last = '5';
                } else if six_r(j, &chars) {
                    last = '6';
                } else if seven_r(j, &chars) {
                    last = '7';
                } else if eight_r(j, &chars) {
                    last = '8';
                } else if nine_r(j, &chars) {
                    last = '9';
                } else {
                    j -= 1;
                }
            }
        }

        result += first.to_digit(10).expect("") * 10 + last.to_digit(10).expect("");

    }

    println!("{}", result);
}

fn single_digit(c: char) -> bool {
    c >= '0' && c <= '9'
}

fn one(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'o' && chars[i + 1] == 'n' && chars[i + 2] == 'e'
}

fn two(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 't' && chars[i + 1] == 'w' && chars[i + 2] == 'o'
}

fn three(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 't' && chars[i + 1] == 'h' && chars[i + 2] == 'r' && chars[i + 3] == 'e' && chars[i + 4] == 'e'
}

fn four(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'f' && chars[i + 1] == 'o' && chars[i + 2] == 'u' && chars[i + 3] == 'r'
}

fn five(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'f' && chars[i + 1] == 'i' && chars[i + 2] == 'v' && chars[i + 3] == 'e'
}

fn six(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 's' && chars[i + 1] == 'i' && chars[i + 2] == 'x'
}

fn seven(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 's' && chars[i + 1] == 'e' && chars[i + 2] == 'v' && chars[i + 3] == 'e' && chars[i + 4] == 'n'
}

fn eight(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'e' && chars[i + 1] == 'i' && chars[i + 2] == 'g' && chars[i + 3] == 'h' && chars[i + 4] == 't'
}

fn nine(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'n' && chars[i + 1] == 'i' && chars[i + 2] == 'n' && chars[i + 3] == 'e'
}

fn one_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'e' && chars[i - 1] == 'n' && chars[i - 2] == 'o'
}

fn two_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'o' && chars[i - 1] == 'w' && chars[i - 2] == 't'
}

// three -> eerht
fn three_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'e' && chars[i - 1] == 'e' && chars[i - 2] == 'r' && chars[i - 3] == 'h' && chars[i - 4] == 't'
}

// four -> ruof
fn four_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'r' && chars[i - 1] == 'u' && chars[i - 2] == 'o' && chars[i - 3] == 'f'
}

// five -> evif
fn five_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'e' && chars[i - 1] == 'v' && chars[i - 2] == 'i' && chars[i - 3] == 'f'
}

fn six_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'x' && chars[i - 1] == 'i' && chars[i - 2] == 's'
}

// seven -> neves
fn seven_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'n' && chars[i - 1] == 'e' && chars[i - 2] == 'v' && chars[i - 3] == 'e' && chars[i - 4] == 's'
}

// eight -> thgie
fn eight_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 't' && chars[i - 1] == 'h' && chars[i - 2] == 'g' && chars[i - 3] == 'i' && chars[i - 4] == 'e'
}

// nine -> enin
fn nine_r(i: usize, chars: &Vec<char>) -> bool {
    chars[i] == 'e' && chars[i - 1] == 'n' && chars[i - 2] == 'i' && chars[i - 3] == 'n'
}

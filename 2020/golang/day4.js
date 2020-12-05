const fs = require('fs')

const file = fs.readFileSync('../input/day4.txt', 'utf-8')

let lines = file.split("\n")

console.log(lines.reduce((acc, l) => {
  if (l === "") {
    acc.lines.push(acc.line)
    acc.line = ""
  } else {
    acc.line += ` ${l}`
  }
  return acc
}, {line: "", lines: []})
  .lines
  .filter(l => ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'].every((w) => l.includes(w)))
  .filter(l => validate(l))
  .map((p) => { console.log({p}); return p })
  .length)

function validate(line) {
  const passport = line.split(" ").reduce((acc, attr) => {
    const [key, val] = attr.split(":")
    acc[key] = val
    return acc
  }, {})
  return [validateByr, validateIyr, validateEyr, validateHgt, validateHcl, validateEcl, validatePid].every(f => f(passport))
}

function validateByr(passport) {
  const byr = parseInt(passport.byr, 10)
  return (1920 <= byr && byr <= 2002)
}
function validateIyr(passport) {
  const iyr = parseInt(passport.iyr, 10)
  return (2010 <= iyr && iyr <= 2020)
}
function validateEyr(passport) {
  const eyr = parseInt(passport.eyr, 10)
  return (2020 <= eyr && eyr <= 2030)
}
function validateHgt(passport) {
  const result = /^([0-9]+)(cm|in)$/.exec(passport.hgt)
  if (!result) {
    return false
  }
  const size = parseInt(result[1], 10)
  switch(result[2]) {
    case "cm":
      return (150 <= size && size <= 193)
    case "in":
      return (59 <= size && size <= 76)
  }
}
function validateHcl(passport) {
  return /^#[0-9a-f]{6}$/.test(passport.hcl)
}
function validateEcl(passport) {
  return ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].includes(passport.ecl)
}
function validatePid(passport) {
  return /^[0-9]{9}$/.test(passport.pid)
}


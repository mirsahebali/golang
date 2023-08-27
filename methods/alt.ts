function r(l: number, b: number) {
  const area = () => l * b;
  const perim = () => l * 2 + b * 2;
  return { area, perim };
}

const a = r(12, 14);
console.log(a.perim());
console.log(a.area());

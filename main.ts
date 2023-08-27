type Color = "Red" | "Green" | "Blue";

type Thing = {
  foo: Function;
};

function smth(str: Function): Thing {
  return {
    foo: str,
  };
}
console.log(smth(() => console.log("hello")));

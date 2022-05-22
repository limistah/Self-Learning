function reverse(str) { 
  if (!str || str.length < 2 || typeof str !== "string") {
    return undefined;
  }

  let backwards = "";
  const totalItems = str.length - 1;
  for (let i = totalItems; i >= 0; i--) {
    backwards = `${backwards}${str[i]}`;
  }
  return backwards;
}

console.log(reverse("Hello, Aleem!"));
